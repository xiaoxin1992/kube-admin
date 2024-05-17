package secret

import (
	"context"
	models "github.com/xiaoxin1992/kube-admin/models/secret"
	"github.com/xiaoxin1992/kube-admin/pkg"
	"github.com/xiaoxin1992/kube-admin/services/k8s"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

func (s *Services) ListSecret(ctx context.Context, req models.QueryList) models.Response {
	response := models.Response{}
	client, err := k8s.NewService().GetClient(ctx, req.Zone)
	if err != nil {
		s.logger.Errorf("get k8s client error: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = "获取k8s客户端出错!"
		return response
	}
	total, err := client.CoreV1().Secrets(req.Namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		s.logger.Errorf("get secret total list error: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = "获取secret总数出错"
		return response
	}
	opts := metav1.ListOptions{
		Limit: int64(req.Page * req.Size),
	}
	secrets, err := client.CoreV1().Secrets(req.Namespace).List(ctx, opts)
	if err != nil {
		s.logger.Errorf("get secret list error: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = "获取secret列表出错!"
		return response
	}
	secretsItems := secrets.Items
	offset, limits := pkg.Page(req.Page, req.Size, len(total.Items))
	secretsItems = secrets.Items[offset:limits]
	response.Code = http.StatusOK
	response.Data = map[string]interface{}{
		"secrets": secretsItems,
		"page":    req.Page,
		"size":    req.Size,
		"total":   len(total.Items),
	}
	return response
}
