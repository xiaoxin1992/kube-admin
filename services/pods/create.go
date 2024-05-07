package pods

import (
	"context"
	"fmt"
	models "github.com/xiaoxin1992/kube-admin/models/pods"
	"github.com/xiaoxin1992/kube-admin/services/k8s"
	k8sError "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
	"strings"
)

func (s *Services) CreatePod(ctx context.Context, req models.CreatePod) models.Response {
	response := models.Response{}
	k, err := k8s.NewService().GetClient(ctx, req.Zone)
	if err != nil {
		response.Code = http.StatusBadRequest
		response.Message = "获取集群信息出错!"
		return response
	}
	_, err = k.CoreV1().Pods(req.Pod.Namespace).Get(ctx, req.Pod.Name, metav1.GetOptions{})
	if err == nil {
		response.Code = http.StatusBadRequest
		response.Message = fmt.Sprintf("pod %s 已经存在!", req.Pod.Name)
		return response
	} else {
		if !k8sError.IsNotFound(err) {
			s.logger.Errorf("get pod %s error: %v", req.Pod.Name, err)
			response.Code = http.StatusBadRequest
			response.Message = fmt.Sprintf("pod %s 获取出错!", req.Pod.Name)
			return response
		}
	}
	if strings.TrimSpace(req.Pod.Namespace) == "" {
		req.Pod.Namespace = defaultName
	}
	// 通过dry run检查语法是否正常
	_, err = k.CoreV1().Pods(req.Pod.Namespace).Create(ctx, &req.Pod, metav1.CreateOptions{DryRun: []string{metav1.DryRunAll}})
	if err != nil {
		s.logger.Errorf(" pod %s syntax error: %v", req.Pod.Name, err)
		response.Message = "创建Pod语法错误!"
		response.Code = http.StatusBadRequest
		return response
	}
	_, err = k.CoreV1().Pods(req.Pod.Namespace).Create(ctx, &req.Pod, metav1.CreateOptions{})
	if err != nil {
		s.logger.Errorf(" create pod %s error: %v", req.Pod.Name, err)
		response.Code = http.StatusBadRequest
		response.Message = fmt.Sprintf("创建pod %s 出错!", req.Pod.Name)
		return response
	}
	response.Code = http.StatusOK
	response.Message = fmt.Sprintf("创建Pod %s 完成", req.Pod.Name)
	return response
}
