package cluster

import (
	"context"
	"github.com/xiaoxin1992/kube-admin/dao/cluster"
	models "github.com/xiaoxin1992/kube-admin/models/cluster"
)

func (s *Service) detail(ctx context.Context, zone string) (*models.QueryCluster, error) {
	dao := cluster.NewDao()
	cl, err := dao.QueryByZone(ctx, zone)
	if err != nil {
		s.logger.Errorf("query cluster by zone fail: %v", err)
	}
	return cl, err
}
