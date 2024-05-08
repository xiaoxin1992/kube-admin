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

func (s *Services) CreateService(ctx context.Context, req models.CreateService) models.Response {
	response := models.Response{}
	client, err := k8s.NewService().GetClient(ctx, req.Zone)
	if err != nil {
		s.logger.Errorf("get k8s client err: %v", err)
		response.Message = fmt.Sprintf("获取k8s客户端出错: %v!", err)
		response.Code = http.StatusBadRequest
		return response
	}
	_, err = client.CoreV1().Services(req.Service.Namespace).Get(ctx, req.Service.Name, metav1.GetOptions{})
	if err == nil {
		response.Code = http.StatusBadRequest
		response.Data = fmt.Sprintf("service %s 已经存在", req.Service.Name)
		return response
	} else {
		if !k8sError.IsNotFound(err) {
			s.logger.Errorf("get k8s service err: %v", err)
			response.Code = http.StatusBadRequest
			response.Message = fmt.Sprintf("检查service %s 信息出错", req.Service.Name)
		}
	}
	_, err = client.CoreV1().Services(req.Service.Namespace).Create(ctx, &req.Service, metav1.CreateOptions{DryRun: []string{metav1.DryRunAll}})
	if err != nil {
		s.logger.Errorf("create k8s syntax service err: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = fmt.Sprintf("创建service语法错误: %v!", err)
		return response
	}
	_, err = client.CoreV1().Services(req.Service.Namespace).Create(ctx, &req.Service, metav1.CreateOptions{})
	if err != nil {
		s.logger.Errorf("create k8s service err: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = fmt.Sprintf("创建service %s 出错: %v!", req.Service.Name, err)
		return response
	}
	response.Code = http.StatusOK
	response.Message = fmt.Sprintf("Service %s 创建完成!", req.Service.Name)
	return response
}
