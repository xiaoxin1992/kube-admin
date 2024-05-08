package secret

import (
	"context"
	models "github.com/xiaoxin1992/kube-admin/models/secret"
	"github.com/xiaoxin1992/kube-admin/services/k8s"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

func (s *Services) ListSecret(ctx context.Context, req models.QueryList) models.Response {
	response := models.Response{}
	secretList := make([]models.Secret, 0)
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
	offset := (req.Page - 1) * req.Size
	secretsItems := secrets.Items
	if req.Page*req.Size < len(total.Items) {
		secretsItems = secrets.Items[offset:]
	}
	for _, secret := range secretsItems {
		sec := models.Secret{
			Name:       secret.Name,
			Namespace:  secret.Namespace,
			Labels:     make(map[string]string),
			Type:       secret.Type,
			CreateTime: secret.CreationTimestamp.Time.Format("2006-01-02 15:04:05"),
		}
		if len(secret.Labels) > 0 {
			sec.Labels = secret.Labels
		}
		secretList = append(secretList, sec)
	}
	response.Code = http.StatusOK
	response.Data = map[string]interface{}{
		"secrets": secretList,
		"page":    req.Page,
		"size":    req.Size,
		"total":   len(total.Items),
	}
	return response
}
