package secret

import (
	"context"
	"fmt"
	models "github.com/xiaoxin1992/kube-admin/models/secret"
	"github.com/xiaoxin1992/kube-admin/services/k8s"
	k8sError "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

func (s *Services) CreateSecret(ctx context.Context, req models.CreateSecret) models.Response {
	response := models.Response{}
	client, err := k8s.NewService().GetClient(ctx, req.Zone)
	if err != nil {
		s.logger.Errorf("get k8s client err: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = fmt.Sprintf("获取k8s客户端出错%v!", err)
		return response
	}
	_, err = client.CoreV1().Secrets(req.Secret.Namespace).Get(ctx, req.Secret.Name, metav1.GetOptions{})
	if err == nil {
		response.Code = http.StatusBadRequest
		response.Message = fmt.Sprintf("secret %s 已经存在", req.Secret.Name)
		return response
	} else {
		if !k8sError.IsNotFound(err) {
			s.logger.Errorf("get secret err: %v", err)
			response.Code = http.StatusBadRequest
			response.Message = fmt.Sprintf("检查secret %s 信息出错", req.Secret.Name)
			return response
		}
	}
	_, err = client.CoreV1().Secrets(req.Secret.Namespace).Create(ctx, &req.Secret, metav1.CreateOptions{DryRun: []string{metav1.DryRunAll}})
	if err != nil {
		s.logger.Errorf(" secret %s syntax error: %v", req.Secret.Name, err)
		response.Code = http.StatusBadRequest
		response.Message = fmt.Sprintf("secret语法错误%v!", err)
		return response
	}
	_, err = client.CoreV1().Secrets(req.Secret.Namespace).Create(ctx, &req.Secret, metav1.CreateOptions{})
	if err != nil {
		s.logger.Errorf("create secret %s error: %v", req.Secret.Name, err)
		response.Code = http.StatusBadRequest
		response.Message = fmt.Sprintf("创建secret %s 出错 %v", req.Secret.Name, err)
	} else {
		response.Code = http.StatusOK
		response.Message = fmt.Sprintf("创建secret %s 完成", req.Secret.Name)
	}
	return response
}
