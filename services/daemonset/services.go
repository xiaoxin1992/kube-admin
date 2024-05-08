package daemonset

import (
	"github.com/xiaoxin1992/kube-admin/pkg/logger"
	"go.uber.org/zap"
)

func NewServices() *Services {
	return &Services{
		logger: logger.GetLogger().S("services").Named("daemonSet"),
	}
}

type Services struct {
	logger *zap.SugaredLogger
}
