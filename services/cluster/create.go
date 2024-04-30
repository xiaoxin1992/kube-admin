package cluster

import (
	"context"
	"fmt"
	"github.com/xiaoxin1992/kube-admin/dao/cluster"
	models "github.com/xiaoxin1992/kube-admin/models/cluster"
	"net/http"
)

func (s *Service) CreateCluster(ctx context.Context, req models.Cluster) models.Response {
	response := models.Response{}
	dao := cluster.NewDao()
	// 判断集群是否存在
	ok, err := dao.ExistsZone(ctx, req.Zone)
	if err != nil {
		s.logger.Errorf("get cluster info error %v", err)
		response.Message = fmt.Sprintf("创建集群 %s 出错!", req.Zone)
		response.Code = http.StatusBadRequest
		return response
	}
	if ok {
		response.Message = fmt.Sprintf("集群 %s 已经存在!", req.Zone)
		response.Code = http.StatusBadRequest
		return response
	}
	err = dao.CreateUser(ctx, &req)
	if err != nil {
		response.Message = fmt.Sprintf("创建集群 %s 出错!", req.Zone)
		response.Code = http.StatusBadRequest
		return response
	}
	response.Code = http.StatusOK
	response.Message = fmt.Sprintf("创建集群 %s 完成!", req.Zone)
	return response
}
