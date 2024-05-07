package pods

import (
	"github.com/xiaoxin1992/kube-admin/pkg/logger"
	"go.uber.org/zap"
)

const defaultName = "default"

func NewServices() *Services {
	return &Services{
		logger: logger.GetLogger().S("services").Named("pods"),
	}
}

type Services struct {
	logger *zap.SugaredLogger
}
