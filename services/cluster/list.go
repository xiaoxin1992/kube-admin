package cluster

import (
	"context"
	"github.com/xiaoxin1992/kube-admin/dao/cluster"
	models "github.com/xiaoxin1992/kube-admin/models/cluster"
)

func (s *Service) ClusterList(ctx context.Context, req models.QueryList) models.ResponseClusterList {
	dao := cluster.NewDao()
	response := models.ResponseClusterList{
		Size: 0,
		Page: 0,
		Data: make(map[string]interface{}, 0),
	}
	total, err := dao.ListClusterCount(ctx, req)
	if err != nil {
		s.logger.Errorf("list cluster count error: %s", err.Error())
		return response
	}
	userList, err := dao.ListCluster(ctx, req)
	if err != nil {
		s.logger.Errorf("get cluster list error %v", err)
		return response
	}
	response.Size = req.Size
	response.Page = req.Page
	response.Data = map[string]interface{}{
		"total":    total,
		"clusters": userList,
	}
	return response
}
