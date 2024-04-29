package users

import (
	"github.com/xiaoxin1992/kube-admin/pkg/logger"
	"go.uber.org/zap"
)

func NewService() *service {
	return &service{
		logger: logger.GetLogger().S("services").Named("user"),
	}
}

type service struct {
	logger *zap.SugaredLogger
}
