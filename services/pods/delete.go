package pods

import (
	"context"
	"fmt"
	models "github.com/xiaoxin1992/kube-admin/models/pods"
	"github.com/xiaoxin1992/kube-admin/services/k8s"
	k8sError "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

func (s *Services) DeletePod(ctx context.Context, req models.DeletePod) models.Response {
	response := models.Response{}
	k, err := k8s.NewService().GetClient(ctx, req.Zone)
	if err != nil {
		s.logger.Errorf("get k8s client err: %+v", err)
		response.Message = "获取集群信息出错!"
		response.Code = http.StatusBadRequest
		return response
	}
	_, err = k.CoreV1().Pods(req.Namespace).Get(ctx, req.Name, metav1.GetOptions{})
	if err != nil {
		response.Code = http.StatusBadRequest
		if k8sError.IsNotFound(err) {
			response.Message = fmt.Sprintf("pod %s 不存在!", req.Name)
		} else {
			s.logger.Errorf("get k8s pod err: %v", err)
			response.Message = fmt.Sprintf("pod %s 查询出错!", req.Name)
		}
		return response
	}
	propagationPolicy := metav1.DeletePropagationBackground
	opts := metav1.DeleteOptions{
		PropagationPolicy: &propagationPolicy,
	}
	err = k.CoreV1().Pods(req.Namespace).Delete(ctx, req.Name, opts)
	if err != nil {
		s.logger.Errorf("delete pod err: %+v", err)
		response.Message = fmt.Sprintf("删除pod %s 出错!", req.Name)
		response.Code = http.StatusBadRequest
		return response
	}
	response.Code = http.StatusOK
	response.Message = fmt.Sprintf("删除Pod %s 完成!", req.Name)
	return response
}
