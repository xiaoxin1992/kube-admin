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

func (s *Services) UpdateConfigmap(ctx context.Context, req models.UpdateConfigmap) models.Response {
	response := models.Response{}
	client, err := k8s.NewService().GetClient(ctx, req.Zone)
	if err != nil {
		s.logger.Errorf("get k8s client err: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = fmt.Sprintf("获取k8s客户端出错%v!", err)
		return response
	}
	_, err = client.CoreV1().ConfigMaps(req.ConfigMap.Namespace).Get(ctx, req.ConfigMap.Name, metav1.GetOptions{})
	if err != nil {
		response.Code = http.StatusBadRequest
		if k8sError.IsNotFound(err) {
			response.Message = fmt.Sprintf("configmap %s 不存在", req.ConfigMap.Name)
		} else {
			response.Message = fmt.Sprintf("configmap %s 查询出错 %v", req.ConfigMap.Name, err)
		}
		return response
	}
	_, err = client.CoreV1().ConfigMaps(req.ConfigMap.Namespace).Update(ctx, &req.ConfigMap, metav1.UpdateOptions{DryRun: []string{metav1.DryRunAll}})
	if err != nil {
		s.logger.Errorf(" configmap %s syntax error: %v", req.ConfigMap.Name, err)
		response.Code = http.StatusBadRequest
		response.Message = fmt.Sprintf("configmap语法错误%v!", err)
		return response
	}
	_, err = client.CoreV1().ConfigMaps(req.ConfigMap.Namespace).Update(ctx, &req.ConfigMap, metav1.UpdateOptions{})
	if err != nil {
		s.logger.Errorf("update configmap %s error: %v", req.ConfigMap.Name, err)
		response.Code = http.StatusBadRequest
		response.Message = fmt.Sprintf("更新configmap %s 出错 %v", req.ConfigMap.Name, err)
	} else {
		response.Code = http.StatusOK
		response.Message = fmt.Sprintf("更新configmap %s 完成", req.ConfigMap.Name)
	}
	return response
}
