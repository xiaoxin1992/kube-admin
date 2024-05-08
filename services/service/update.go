package service

import (
	"context"
	"fmt"
	models "github.com/xiaoxin1992/kube-admin/models/service"
	"github.com/xiaoxin1992/kube-admin/services/k8s"
	k8sError "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

func (s *Services) UpdateService(ctx context.Context, req models.UpdateService) models.Response {
	response := models.Response{}
	client, err := k8s.NewService().GetClient(ctx, req.Zone)
	if err != nil {
		s.logger.Errorf("get k8s client err: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = fmt.Sprintf("获取k8s客户端出错%v!", err)
		return response
	}
	_, err = client.CoreV1().Services(req.Service.Namespace).Get(ctx, req.Service.Name, metav1.GetOptions{})
	if err != nil {
		response.Code = http.StatusBadRequest
		if k8sError.IsNotFound(err) {
			response.Message = fmt.Sprintf("service %s 不存在", req.Service.Name)
		} else {
			response.Message = fmt.Sprintf("service %s 查询出错 %v", req.Service.Name, err)
		}
		return response
	}
	_, err = client.CoreV1().Services(req.Service.Namespace).Update(ctx, &req.Service, metav1.UpdateOptions{DryRun: []string{metav1.DryRunAll}})
	if err != nil {
		s.logger.Errorf(" service %s syntax error: %v", req.Service.Name, err)
		response.Code = http.StatusBadRequest
		response.Message = fmt.Sprintf("service语法错误%v!", err)
		return response
	}
	_, err = client.CoreV1().Services(req.Service.Namespace).Update(ctx, &req.Service, metav1.UpdateOptions{})
	if err != nil {
		s.logger.Errorf("update service %s error: %v", req.Service.Name, err)
		response.Code = http.StatusBadRequest
		response.Message = fmt.Sprintf("更新service %s 出错 %v", req.Service.Name, err)
	} else {
		response.Code = http.StatusOK
		response.Message = fmt.Sprintf("更新service %s 完成", req.Service.Name)
	}
	return response
}
