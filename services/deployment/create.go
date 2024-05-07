package deployment

import (
	"context"
	"fmt"
	models "github.com/xiaoxin1992/kube-admin/models/deployment"
	"github.com/xiaoxin1992/kube-admin/services/k8s"
	k8sError "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

func (s *Services) CreateDeployment(ctx context.Context, req models.CreateDeployment) models.Response {
	response := models.Response{}
	client, err := k8s.NewService().GetClient(ctx, req.Zone)
	if err != nil {
		s.logger.Errorf("get k8s client err: %+v", err)
		response.Message = fmt.Sprintf("获取k8s客户端出错: %v", err)
		response.Code = http.StatusBadRequest
		return response
	}
	_, err = client.AppsV1().Deployments(req.Deployment.Namespace).Get(ctx, req.Deployment.Name, metav1.GetOptions{})
	if err == nil {
		response.Message = fmt.Sprintf("deployment %s 已经存在", req.Deployment.Name)
		response.Code = http.StatusBadRequest
		return response
	} else {
		if !k8sError.IsNotFound(err) {
			s.logger.Errorf("get k8s deployment err: %v", err)
			response.Message = fmt.Sprintf("查询deployment %s 出错: %v", req.Deployment.Name, err)
			response.Code = http.StatusBadRequest
			return response
		}
	}
	_, err = client.AppsV1().Deployments(req.Deployment.Namespace).Create(ctx, &req.Deployment, metav1.CreateOptions{DryRun: []string{metav1.DryRunAll}})
	if err != nil {
		s.logger.Errorf("create k8s deployment syntax err: %v", err)
		response.Message = fmt.Sprintf("创建deployment %s 语法错误: %v", req.Deployment.Name, err)
		response.Code = http.StatusBadRequest
		return response
	}
	_, err = client.AppsV1().Deployments(req.Deployment.Namespace).Create(ctx, &req.Deployment, metav1.CreateOptions{})
	if err != nil {
		s.logger.Errorf("create k8s deployment err: %v", err)
		response.Message = fmt.Sprintf("创建deployment %s 出错: %v", req.Deployment.Name, err)
		response.Code = http.StatusBadRequest
		return response
	}
	response.Message = fmt.Sprintf("创建deployment %s 完成", req.Deployment.Name)
	response.Code = http.StatusOK
	return response
}
