package pv

import (
	"github.com/xiaoxin1992/kube-admin/pkg/logger"
	"go.uber.org/zap"
)

func NewServices() *Services {
	return &Services{
		logger: logger.GetLogger().S("services").Named("pv"),
	}
}

type Services struct {
	logger *zap.SugaredLogger
}
