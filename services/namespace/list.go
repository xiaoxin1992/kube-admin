package namespace

import (
	"context"
	models "github.com/xiaoxin1992/kube-admin/models/namespace"
	"github.com/xiaoxin1992/kube-admin/services/k8s"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

func (s *Services) ListNamespaces(ctx context.Context, req models.QueryList) models.Response {
	response := models.Response{}
	nsList := make([]models.ListNamespace, 0)
	client, err := k8s.NewService().GetClient(ctx, req.Zone)
	if err != nil {
		s.logger.Errorf("get k8s client err: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = "后去集群客户端出错!"
		return response
	}
	total, err := client.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		s.logger.Errorf("get namespaces total err: %v", err)
		response.Code = http.StatusInternalServerError
		response.Message = "获取namespace列表出错!"
		return response
	}
	opts := metav1.ListOptions{
		Limit: int64(req.Page * req.Size),
	}
	namespaces, err := client.CoreV1().Namespaces().List(ctx, opts)
	if err != nil {
		s.logger.Errorf("get namespace list error: %v", err)
		response.Code = http.StatusInternalServerError
		response.Message = "获取namespace列表出错!"
		return response
	}
	offset := (req.Page - 1) * req.Size
	namespacesItems := namespaces.Items
	if req.Page*req.Size <= len(total.Items) {
		namespacesItems = namespaces.Items[offset:]
	}
	for _, ns := range namespacesItems {
		nsList = append(nsList, models.ListNamespace{
			Namespace: models.Namespace{
				Name:  ns.GetName(),
				Label: ns.GetLabels(),
			},
			Status:     string(ns.Status.Phase),
			CreateTime: ns.CreationTimestamp.Time.Format("2006-01-02 15:04:05"),
		})
	}
	response.Code = http.StatusOK
	response.Data = map[string]interface{}{
		"namespace": nsList,
		"page":      req.Page,
		"size":      req.Size,
		"total":     len(total.Items),
	}
	return response
}
