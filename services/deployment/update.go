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

func (s *Services) UpdateDeployment(ctx context.Context, req models.UpdateDeployment) models.Response {
	response := models.Response{}
	client, err := k8s.NewService().GetClient(ctx, req.Zone)
	if err != nil {
		s.logger.Errorf("get k8s client err: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = fmt.Sprintf("获取k8s客户端出错%v!", err)
		return response
	}
	_, err = client.AppsV1().Deployments(req.Deployment.Namespace).Get(ctx, req.Deployment.Name, metav1.GetOptions{})
	if err != nil {
		response.Code = http.StatusBadRequest
		if k8sError.IsNotFound(err) {
			response.Message = fmt.Sprintf("deployment %s 不存在", req.Deployment.Name)
		} else {
			response.Message = fmt.Sprintf("deployment %s 查询出错 %v", req.Deployment.Name, err)
		}
		return response
	}
	_, err = client.AppsV1().Deployments(req.Deployment.Namespace).Update(ctx, &req.Deployment, metav1.UpdateOptions{DryRun: []string{metav1.DryRunAll}})
	if err != nil {
		s.logger.Errorf(" deployment %s syntax error: %v", req.Deployment.Name, err)
		response.Code = http.StatusBadRequest
		response.Message = fmt.Sprintf("deployment语法错误%v!", err)
		return response
	}
	_, err = client.AppsV1().Deployments(req.Deployment.Namespace).Update(ctx, &req.Deployment, metav1.UpdateOptions{})
	if err != nil {
		s.logger.Errorf("update configmap %s error: %v", req.Deployment.Name, err)
		response.Code = http.StatusBadRequest
		response.Message = fmt.Sprintf("更新deployment %s 出错 %v", req.Deployment.Name, err)
	} else {
		response.Code = http.StatusOK
		response.Message = fmt.Sprintf("更新deployment %s 完成", req.Deployment.Name)
	}
	return response
}
