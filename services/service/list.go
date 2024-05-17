package service

import (
	"context"
	"fmt"
	models "github.com/xiaoxin1992/kube-admin/models/service"
	"github.com/xiaoxin1992/kube-admin/pkg"
	"github.com/xiaoxin1992/kube-admin/services/k8s"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

func (s *Services) ListService(ctx context.Context, req models.QueryList) models.Response {
	response := models.Response{}
	serviceSlice := make([]models.Service, 0)
	client, err := k8s.NewService().GetClient(ctx, req.Zone)
	if err != nil {
		s.logger.Errorf("get k8s client err: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = fmt.Sprintf("获取 k8s 客户端出错: %v!", err)
		return response
	}
	total, err := client.CoreV1().Services(req.Namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		s.logger.Errorf("get k8s total service err: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = fmt.Sprintf("获取 service总数出错: %v!", err)
		return response
	}
	opts := metav1.ListOptions{
		Limit: int64(req.Size * req.Page),
	}
	serviceList, err := client.CoreV1().Services(req.Namespace).List(ctx, opts)
	if err != nil {
		s.logger.Errorf("get k8s service err: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = fmt.Sprintf("获取 service 出错: %v!", err)
		return response
	}
	serviceItems := serviceList.Items
	offset, limits := pkg.Page(req.Page, req.Size, len(total.Items))
	serviceItems = serviceList.Items[offset:limits]
	for _, service := range serviceItems {
		si := models.Service{
			Name:       service.Name,
			Namespace:  service.Namespace,
			Labels:     make(map[string]string),
			Port:       make([]corev1.ServicePort, 0),
			Type:       service.Spec.Type,
			CreateTime: service.CreationTimestamp.Format("2006-01-02 15:04:05"),
		}
		if len(service.Labels) > 0 {
			si.Labels = service.Labels
		}
		if len(service.Spec.Ports) != 0 {
			si.Port = service.Spec.Ports
		}
		serviceSlice = append(serviceSlice, si)
	}

	response.Code = http.StatusOK
	response.Data = map[string]interface{}{
		"services": serviceSlice,
		"page":     req.Page,
		"size":     req.Size,
		"total":    len(total.Items),
	}
	return response
}
