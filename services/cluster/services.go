package cluster

import (
	"github.com/xiaoxin1992/kube-admin/pkg/logger"
	"go.uber.org/zap"
)

func NewCluster() *Service {
	return &Service{
		logger: logger.GetLogger().S("services").Named("cluster"),
	}
}

type Service struct {
	logger *zap.SugaredLogger
}
