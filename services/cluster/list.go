package cluster

import (
	"context"
	"github.com/xiaoxin1992/kube-admin/dao/cluster"
	models "github.com/xiaoxin1992/kube-admin/models/cluster"
	"net/http"
)

func (s *Service) ClusterList(ctx context.Context, req models.QueryList) models.Response {
	dao := cluster.NewDao()
	response := models.Response{}
	total, err := dao.ListClusterCount(ctx, req)
	if err != nil {
		s.logger.Errorf("list cluster count error: %s", err.Error())
		response.Code = http.StatusBadRequest
		response.Message = "获取集群列表出错!"
		return response
	}
	clusterList, err := dao.ListCluster(ctx, req)
	if err != nil {
		s.logger.Errorf("get cluster list error %v", err)
		response.Code = http.StatusBadRequest
		response.Message = "获取集群列表出错!"
		return response
	}
	response.Code = http.StatusOK
	response.Data = map[string]interface{}{
		"total":    total,
		"clusters": clusterList,
	}
	return response
}
