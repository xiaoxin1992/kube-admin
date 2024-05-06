package cluster

import (
	"context"
	"fmt"
	"github.com/xiaoxin1992/kube-admin/dao/cluster"
	models "github.com/xiaoxin1992/kube-admin/models/cluster"
	"github.com/xiaoxin1992/kube-admin/pkg/k8s"
	"net/http"
)

func (s *Service) CreateCluster(ctx context.Context, req models.CreateCluster) models.Response {
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
	k, err := k8s.NewClient(req.Host, req.Token).GetClientSet()
	if err != nil {
		s.logger.Errorf("connect cluster error %v", err)
		response.Message = fmt.Sprintf("检测 %s 集群链接出错!", req.Zone)
		response.Code = http.StatusBadRequest
		return response
	}
	version, err := k.ServerVersion()
	if err != nil {
		req.Version = ""
		s.logger.Errorf("get server version error %v", err)
		response.Message = fmt.Sprintf("获取 %s 集群版本出错!", req.Zone)
		response.Code = http.StatusBadRequest
		return response
	} else {
		req.Version = version.String()
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
