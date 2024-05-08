package configmap

import (
	"context"
	"fmt"
	models "github.com/xiaoxin1992/kube-admin/models/configmap"
	"github.com/xiaoxin1992/kube-admin/services/k8s"
	k8sError "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

func (s *Services) DeleteConfigmap(ctx context.Context, req models.DeleteConfigmap) models.Response {
	response := models.Response{}
	client, err := k8s.NewService().GetClient(ctx, req.Zone)
	if err != nil {
		s.logger.Errorf("get k8s client err: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = fmt.Sprintf("获取k8s客户端出错%v!", err)
		return response
	}
	_, err = client.CoreV1().ConfigMaps(req.Namespace).Get(ctx, req.Name, metav1.GetOptions{})
	if err != nil {
		response.Code = http.StatusBadRequest
		if k8sError.IsNotFound(err) {
			response.Message = fmt.Sprintf("configmap %s 不存在", req.Name)
		} else {
			s.logger.Errorf("get configmap %s err: %v", req.Name, err)
			response.Message = fmt.Sprintf("configmap %s 查询出错: %v", req.Name, err)
		}
		return response
	}
	err = client.CoreV1().ConfigMaps(req.Namespace).Delete(ctx, req.Name, metav1.DeleteOptions{})
	if err != nil {
		s.logger.Errorf("delete configmap %s error: %v", req.Name, err)
		response.Code = http.StatusBadRequest
		response.Message = fmt.Sprintf("删除configmap %s 出错 %v", req.Name, err)
		return response
	} else {
		response.Code = http.StatusOK
		response.Message = fmt.Sprintf("删除configmap %s 完成", req.Name)
	}
	return response
}
