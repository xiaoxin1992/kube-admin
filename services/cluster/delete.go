package cluster

import (
	"context"
	"fmt"
	"github.com/xiaoxin1992/kube-admin/dao/cluster"
	models "github.com/xiaoxin1992/kube-admin/models/cluster"
	"net/http"
)

func (s *Service) DeleteCluster(ctx context.Context, req models.DeleteCluster) models.Response {
	dao := cluster.NewDao()
	response := models.Response{
		Code:    http.StatusBadRequest,
		Message: "",
	}
	ok, err := dao.ExistsZone(ctx, req.Zone)
	if err != nil {
		s.logger.Errorf("get cluster info error %v", err)
		response.Message = fmt.Sprintf("获取集群 %s 信息出错!", req.Zone)
		return response
	}
	if !ok {
		response.Message = fmt.Sprintf("集群 %s 不存在!", req.Zone)
		return response
	}
	err = dao.DeleteCluster(ctx, &req)
	if err != nil {
		s.logger.Errorf("delete user error %v", err)
		response.Message = fmt.Sprintf("集群 %s 删除失败", req.Zone)
		return response
	}
	response.Code = http.StatusOK
	response.Message = fmt.Sprintf("集群 %s 删除完成!", req.Zone)
	return response
}
