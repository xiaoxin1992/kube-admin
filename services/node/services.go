package node

import (
	"github.com/xiaoxin1992/kube-admin/pkg/logger"
	"go.uber.org/zap"
)

func NewService() *Services {
	return &Services{
		logger: logger.GetLogger().S("services").Named("node"),
	}
}

type Services struct {
	logger *zap.SugaredLogger
}
