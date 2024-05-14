package namespace

import (
	"context"
	"fmt"
	models "github.com/xiaoxin1992/kube-admin/models/namespace"
	"github.com/xiaoxin1992/kube-admin/services/k8s"
	k8sError "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

func (s *Services) DeleteNamespace(ctx context.Context, req models.DeleteNamespace) models.Response {
	response := models.Response{}
	client, err := k8s.NewService().GetClient(ctx, req.Zone)
	if err != nil {
		s.logger.Errorf("get k8s client err: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = "获取集群信息出错!"
		return response
	}

	_, err = client.CoreV1().Namespaces().Get(ctx, req.Name, metav1.GetOptions{})
	if err != nil {
		response.Code = http.StatusBadRequest
		if k8sError.IsNotFound(err) {
			response.Message = fmt.Sprintf("namespace %s 不存在", req.Name)
		} else {
			response.Message = fmt.Sprintf("检查 namespace %s 出错: %v", req.Name, err)
		}
		return response
	}
	propagationPolicy := metav1.DeletePropagationBackground
	opts := metav1.DeleteOptions{
		PropagationPolicy: &propagationPolicy,
	}
	err = client.CoreV1().Namespaces().Delete(ctx, req.Name, opts)
	if err != nil {
		s.logger.Errorf("delete namespace err: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = fmt.Sprintf("删除 namespace %s 出错: %v", req.Name, err)
	}
	response.Code = http.StatusOK
	response.Message = fmt.Sprintf("删除 namespace %s 完成", req.Name)
	return response
}
