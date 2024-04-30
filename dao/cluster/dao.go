package cluster

import (
	"github.com/xiaoxin1992/kube-admin/pkg/logger"
	"go.uber.org/zap"
)

func NewDao() *Dao {
	return &Dao{
		logger: logger.GetLogger().S("MySQL").Named("cluster"),
	}
}

type Dao struct {
	logger *zap.SugaredLogger
}
