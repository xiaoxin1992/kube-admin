package pods

import (
	"context"
	"github.com/xiaoxin1992/kube-admin/pkg/logger"
	"go.uber.org/zap"
	k8sError "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
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

func (s *Services) ExistsByPod(ctx context.Context, client *kubernetes.Clientset, namespace, name string) (bool, error) {
	_, err := client.CoreV1().Pods(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		if !k8sError.IsNotFound(err) {
			return false, err
		}
	} else {
		return true, nil
	}
	return false, nil
}
