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

func (s *Services) CreateNamespace(ctx context.Context, req models.CreateNamespace) models.Response {
	response := models.Response{}
	client, err := k8s.NewService().GetClient(ctx, req.Zone)
	if err != nil {
		s.logger.Errorf("get k8s client err: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = "获取k8s客户端出错!"
		return response
	}
	_, err = client.CoreV1().Namespaces().Get(ctx, req.Namespace.Name, metav1.GetOptions{})
	if err == nil {
		response.Code = http.StatusBadRequest
		response.Message = fmt.Sprintf("namespace %s 已经存在", req.Namespace.GetName())
		return response
	} else {
		if !k8sError.IsNotFound(err) {
			s.logger.Errorf("check namespace %s err: %v", req.Namespace.GetName(), err)
			response.Code = http.StatusBadRequest
			response.Message = "检查namespace出错"
			return response
		}
	}
	// 检查创建namespace语法是否正确
	_, err = client.CoreV1().Namespaces().Create(ctx, &req.Namespace, metav1.CreateOptions{DryRun: []string{metav1.DryRunAll}})
	if err != nil {
		s.logger.Errorf("check namespace syntax error %v", err)
		response.Code = http.StatusBadRequest
		response.Message = fmt.Sprintf("namespace %s 语法错误 %v", req.Namespace.GetName(), err)
		return response
	}
	_, err = client.CoreV1().Namespaces().Create(ctx, &req.Namespace, metav1.CreateOptions{})
	if err != nil {
		s.logger.Errorf("create namespace %s err: %v", req.Namespace.Name, err)
		response.Message = fmt.Sprintf("创建 namespace %s 出错: %v", req.Namespace.Name, err)
		return response
	}
	response.Code = http.StatusOK
	response.Message = fmt.Sprintf("namespace %s 创建完成", req.Namespace.Name)
	return response
}
