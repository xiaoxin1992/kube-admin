package pvc

import (
	"context"
	"fmt"
	models "github.com/xiaoxin1992/kube-admin/models/pvc"
	"github.com/xiaoxin1992/kube-admin/services/k8s"
	k8sError "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

func (s *Services) CreatePersistentVolumeClaim(ctx context.Context, req models.CreatePersistentVolumeClaim) models.Response {
	response := models.Response{}
	client, err := k8s.NewService().GetClient(ctx, req.Zone)
	if err != nil {
		s.logger.Errorf("get k8s client err: %v", err)
		response.Message = fmt.Sprintf("获取k8s客户端出错: %v!", err)
		response.Code = http.StatusBadRequest
		return response
	}
	_, err = client.CoreV1().PersistentVolumeClaims(req.PersistentVolumeClaim.Namespace).Get(ctx, req.PersistentVolumeClaim.Name, metav1.GetOptions{})
	if err == nil {
		response.Code = http.StatusBadRequest
		response.Data = fmt.Sprintf("pvc %s 已经存在", req.PersistentVolumeClaim.Name)
		return response
	} else {
		if !k8sError.IsNotFound(err) {
			s.logger.Errorf("get k8s pvc err: %v", err)
			response.Code = http.StatusBadRequest
			response.Message = fmt.Sprintf("检查pvc %s 信息出错", req.PersistentVolumeClaim.Name)
		}
	}
	_, err = client.CoreV1().PersistentVolumeClaims(req.PersistentVolumeClaim.Namespace).Create(ctx, &req.PersistentVolumeClaim, metav1.CreateOptions{DryRun: []string{metav1.DryRunAll}})
	if err != nil {
		s.logger.Errorf("create k8s syntax pvc err: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = fmt.Sprintf("创建pvc语法错误: %v!", err)
		return response
	}
	_, err = client.CoreV1().PersistentVolumeClaims(req.PersistentVolumeClaim.Namespace).Create(ctx, &req.PersistentVolumeClaim, metav1.CreateOptions{})
	if err != nil {
		s.logger.Errorf("create k8s pvc err: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = fmt.Sprintf("创建pvc %s 出错: %v!", req.PersistentVolumeClaim.Name, err)
		return response
	}
	response.Code = http.StatusOK
	response.Message = fmt.Sprintf("pvc %s 创建完成!", req.PersistentVolumeClaim.Name)
	return response
}
