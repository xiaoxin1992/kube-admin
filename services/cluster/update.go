package cluster

import (
	"context"
	"fmt"
	"github.com/xiaoxin1992/kube-admin/dao/cluster"
	models "github.com/xiaoxin1992/kube-admin/models/cluster"
	"net/http"
)

func (s *Service) UpdateCluster(ctx context.Context, updateCluster models.UpdateCluster) models.Response {
	db := cluster.NewDao()
	ok, err := db.ExistsZone(ctx, updateCluster.Zone)
	if err != nil {
		s.logger.Errorf("query cluster info err %v", err)
		return models.Response{
			Code:    http.StatusBadRequest,
			Message: "查询集群信息出错!",
		}
	}
	if !ok {
		return models.Response{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("集群 %s 不存在", updateCluster.Zone),
		}
	}
	err = db.UpdateCluster(ctx, updateCluster)
	if err != nil {
		s.logger.Errorf("update cluster err: %v", err)
		return models.Response{
			Code:    http.StatusBadRequest,
			Message: "更新集群信息失败!",
		}
	}
	return models.Response{
		Code:    http.StatusOK,
		Message: "更新集群信息完成!",
	}
}
