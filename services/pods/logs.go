package pods

import (
	"bufio"
	"context"
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
		s.logger.Errorf("get k8s client err: %v", err)
		return result
	}
	if req.TailLines == 0 {
		req.TailLines = 50
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
		s.logger.Errorf("get logs stream err: %v", err)
		return result
	}
	defer func() {
		if StreamErr := stream.Close(); StreamErr != nil {
			s.logger.Errorf("close stream err: %v", StreamErr)
		}
		s.logger.Debugf("close stream")
	}()
	conn, err := s.upgrade.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		s.logger.Errorf("http to websocket upgrade err: %v", err)
		return result
	}
	defer func(conn *websocket.Conn) {
		if connErr := conn.Close(); connErr != nil {
			s.logger.Errorf("close websocket err: %v", connErr)
		}
	}(conn)
	err = conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		s.logger.Errorf("set read deadline err: %v", err)
	}
	conn.SetPongHandler(func(string) error {
		err = conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		if err != nil {
			s.logger.Errorf("set pong err: %v", err)
		}
		return nil
	})
	conn.SetCloseHandler(func(code int, text string) error {
		s.logger.Debugf("websocket closed\n")
		cancel()
		return nil
	})
	go func() {
		for {
			msgType, content, msgErr := conn.ReadMessage()
			if msgErr != nil {
				if !websocket.IsUnexpectedCloseError(msgErr, websocket.CloseNormalClosure) {
					s.logger.Errorf("read message err: %v", err)
				}
				break
			}
			if msgType == websocket.TextMessage && s.WebSocketMessagePing(string(content)) == websocket.PingMessage {
				if err = conn.WriteMessage(websocket.PingMessage, nil); err != nil {
					s.logger.Errorf("write ping message err: %v", err)
					break
				}
			}
		}
		defer cancel()
	}()
	inputChan := make(chan string)
	go func() {
		// 读取日志数据
		reader := bufio.NewScanner(stream)
		for reader.Scan() {
			inputChan <- reader.Text()
		}
		defer close(inputChan)
	}()
	for {
		select {
		case <-logCtx.Done():
			s.logger.Debugf("exit websocket logs")
			return result
		case msg, ok := <-inputChan:
			if !ok {
				s.logger.Debugf("input channel closed")
				return result
			}
			err = conn.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				s.logger.Errorf("write message err: %v", err)
				return result
			}
		default:
			continue
		}
	}
}
