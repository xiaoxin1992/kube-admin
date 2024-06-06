package pods

import (
	"bufio"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	models "github.com/xiaoxin1992/kube-admin/models/pods"
	"github.com/xiaoxin1992/kube-admin/services/k8s"
	v1 "k8s.io/api/core/v1"
	"time"
)

func (s *Services) LogsPod(ctx *gin.Context, req models.LogsPod) models.Response {
	result := models.Response{}
	k, err := k8s.NewService().GetClient(ctx, req.Zone)
	if err != nil {
		s.logger.Errorf("get k8s client err: %+v", err)
		result.Code = 400
		result.Message = "获取k8s客户端出错!"
		return result
	}
	if req.TailLines == 0 {
		req.TailLines = 100
	}
	opt := v1.PodLogOptions{
		Container: req.Container,
		Follow:    true,
		TailLines: &req.TailLines,
	}
	logReq := k.CoreV1().Pods(req.Namespace).GetLogs(req.Name, &opt)
	logCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	stream, err := logReq.Stream(logCtx)
	if err != nil {
		s.logger.Errorf("create k8s  logs stream err: %+v", err)
		result.Code = 400
		result.Message = "获取logs出错!"
		return result
	}
	defer stream.Close()
	ws, err := s.upgrade.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		s.logger.Errorf("upgrade websocket error : %+v", err)
		result.Code = 400
		result.Message = "升级websocket出错!"
		return result
	}
	defer ws.Close()
	ws.SetPingHandler(func(appData string) error {
		err = ws.WriteControl(websocket.PongMessage, []byte{}, time.Now().Add(time.Second*5))
		if err != nil {
			fmt.Println("write pong error:", err)
		}
		return err
	})
	var exitChan = make(chan bool)
	go func() {
		for {
			var msgErr error
			_, _, msgErr = ws.ReadMessage()
			if msgErr != nil {
				if websocket.IsUnexpectedCloseError(msgErr, websocket.CloseNormalClosure) {
					break
				} else {
					s.logger.Errorf("read message err: %+v", msgErr)
				}
				break
			}
		}

		defer func() {
			exitChan <- true
		}()
	}()
	reader := bufio.NewScanner(stream)
	result.Code = 200
	result.Message = "获取日志结束了"
	for reader.Scan() {
		select {
		case <-exitChan:
			fmt.Println("exit websocket")
			return result
		default:
			msg := reader.Text()
			fmt.Println(msg, "logs----")
			err = ws.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				s.logger.Errorf("write logs err: %+v", err)
				return result
			}
		}
	}
	return result
}
