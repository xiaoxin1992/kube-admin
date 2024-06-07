package pods

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/xiaoxin1992/kube-admin/pkg/logger"
	"go.uber.org/zap"
	"net/http"
)

const defaultName = "default"

func NewServices() *Services {
	return &Services{
		logger: logger.GetLogger().S("services").Named("pods"),
		upgrade: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

type Services struct {
	logger  *zap.SugaredLogger
	upgrade *websocket.Upgrader
}

func (s *Services) WebSocketMessagePing(data string) int {
	d := struct {
		Type string `json:"type"`
	}{}
	err := json.Unmarshal([]byte(data), &d)
	if err != nil {
		return 1
	}
	if d.Type == "ping" {
		return 9
	}
	return 1
}
