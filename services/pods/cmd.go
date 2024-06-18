package pods

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	models "github.com/xiaoxin1992/kube-admin/models/pods"
	"github.com/xiaoxin1992/kube-admin/services/k8s"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
)

// 消息结构体， 解析从websocket接受的消息
type message struct {
	MsgType string `json:"msgType"`
	Rows    uint16 `json:"rows"`
	Cols    uint16 `json:"cols"`
	Data    string `json:"data"`
}

// 封装websocket， 用于remoteCommand之间的数据交换

type WebsocketCommand struct {
	ws *websocket.Conn
	// 类型通道， 用于传输窗口大小事件
	resize chan remotecommand.TerminalSize
	cancel context.CancelFunc
}

// 定义Read方法，实现io.Reader接口， 从websocket读取数据
func (w *WebsocketCommand) Read(p []byte) (int, error) {
	_, msgContent, err := w.ws.ReadMessage()
	if err != nil {
		if websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure) {
			closeTExt := "exit\r"
			copy(p, closeTExt)
			w.cancel()
			return len(closeTExt), nil
		}
		return 0, err
	}
	var msg message
	if err = json.Unmarshal(msgContent, &msg); err != nil {
		return 0, err
	}
	switch msg.MsgType {
	case "resize":
		winSize := remotecommand.TerminalSize{Width: msg.Cols, Height: msg.Rows}
		w.resize <- winSize
		return 0, nil
	case "input":
		copy(p, msg.Data)
		return len(msg.Data), nil
	}
	return 0, nil
}

// 定义Write方法， 实现io.Writer接口, 将数据写入websocket
func (w *WebsocketCommand) Write(p []byte) (n int, err error) {
	err = w.ws.WriteMessage(websocket.TextMessage, p)
	return len(p), err
}

// Next Next方法，用于resize通道获取下一个terminalSize事件
func (w *WebsocketCommand) Next() *remotecommand.TerminalSize {
	size := <-w.resize
	return &size
}

func (s *Services) terminalCreate(ctx context.Context, wsClient *WebsocketCommand, config *rest.Config, request *rest.Request) (err error) {
	exec, err := remotecommand.NewSPDYExecutor(config, "POST", request.URL())
	if err != nil {
		s.logger.Errorf("Failed to create executor: %v", err)
		return
	}
	err = exec.StreamWithContext(ctx, remotecommand.StreamOptions{
		Stderr:            wsClient,
		Stdout:            wsClient,
		Stdin:             wsClient,
		Tty:               true,
		TerminalSizeQueue: wsClient,
	})
	if err != nil {
		s.logger.Errorf("Failed to stream: %v", err)
	}
	return err
}
func (s *Services) CmdPod(ctx *gin.Context, req models.CmdPod) models.Response {
	clientConfig, err := k8s.NewService().GetConfig(ctx, req.Zone)
	if err != nil {
		s.logger.Errorf("get k8s client err: %+v", err)
		return models.Response{}
	}
	ws, err := s.upgrade.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		s.logger.Errorf("upgrade websocket err: %+v", err)
		return models.Response{}
	}
	defer func() {
		if wsErr := ws.Close(); wsErr != nil {
			s.logger.Errorf("close websocket err: %+v", wsErr)
		}
	}()
	ws.SetCloseHandler(func(code int, text string) error {
		s.logger.Errorf("close websocket code: %d, text: %s", code, text)
		return nil
	})
	clientSet, err := kubernetes.NewForConfig(clientConfig)
	if err != nil {
		s.logger.Errorf("new k8s client err: %+v", err)
		return models.Response{}
	}
	request := clientSet.CoreV1().RESTClient().Post().
		Resource("pods").
		Namespace(req.Namespace).
		Name(req.Name).
		SubResource("exec").
		VersionedParams(&v1.PodExecOptions{
			Container: req.Container,
			Command:   []string{req.Cmd},
			Stdout:    true,
			Stdin:     true,
			Stderr:    true,
			TTY:       true,
		}, scheme.ParameterCodec)
	cmdCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	err = s.terminalCreate(cmdCtx, &WebsocketCommand{
		ws:     ws,
		resize: make(chan remotecommand.TerminalSize),
		cancel: cancel,
	}, clientConfig, request)
	if err != nil {
		s.logger.Errorf("cmd exec err: %+v", err)
		err = ws.WriteMessage(websocket.TextMessage, []byte(err.Error()))
		if err != nil {
			s.logger.Errorf("write cmd info error: %+v", err)
		}
	}
	s.logger.Warn("websocket exit...")
	return models.Response{}
}
