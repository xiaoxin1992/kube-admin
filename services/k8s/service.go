package k8s

import (
	"context"
	"github.com/xiaoxin1992/kube-admin/dao/cluster"
	"github.com/xiaoxin1992/kube-admin/pkg/k8s"
	"k8s.io/client-go/kubernetes"
)

func NewService() *Service {
	return &Service{}
}

type Service struct {
}

func (s *Service) GetClient(ctx context.Context, zone string) (*kubernetes.Clientset, error) {
	dao := cluster.NewDao()
	cs, err := dao.QueryByZone(ctx, zone)
	if err != nil {
		return nil, err
	}
	k := k8s.NewClient(cs.Host, cs.Token)
	return k.GetClientSet()
}
