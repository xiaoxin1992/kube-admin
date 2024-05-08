package pv

import (
	"context"
	"fmt"
	models "github.com/xiaoxin1992/kube-admin/models/pv"
	"github.com/xiaoxin1992/kube-admin/services/k8s"
	k8sError "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

func (s *Services) DetailPV(ctx context.Context, req models.DetailQuery) models.Response {
	response := models.Response{}
	client, err := k8s.NewService().GetClient(ctx, req.Zone)
	if err != nil {
		s.logger.Errorf("get k8s client err: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = fmt.Sprintf("获取k8s客户端出错 %v", err)
		return response
	}
	pvDetail, err := client.CoreV1().PersistentVolumes().Get(ctx, req.Name, metav1.GetOptions{})
	if err != nil {
		response.Code = http.StatusNotFound
		if k8sError.IsNotFound(err) {
			response.Message = fmt.Sprintf("pv %s 不存在", req.Name)
		} else {
			response.Message = fmt.Sprintf("pv %s 查询出错 %v", req.Name, err)
		}
		return response
	}
	response.Code = http.StatusOK
	response.Data = map[string]interface{}{
		"pv": pvDetail,
	}
	return response
}
