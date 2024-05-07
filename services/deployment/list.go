package deployment

import (
	"context"
	models "github.com/xiaoxin1992/kube-admin/models/deployment"
	"github.com/xiaoxin1992/kube-admin/services/k8s"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

func (s *Services) ListDeployment(ctx context.Context, req models.QueryList) models.Response {
	response := models.Response{}
	deploymnetList := make([]models.Deployment, 0)
	client, err := k8s.NewService().GetClient(ctx, req.Zone)
	if err != nil {
		s.logger.Errorf("get k8s client error: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = "获取k8s客户端出错!"
		return response
	}
	total, err := client.AppsV1().Deployments(req.Namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		s.logger.Errorf("get deployment total list error: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = "获取deployment总数出错"
		return response
	}
	opts := metav1.ListOptions{
		Limit: int64(req.Page * req.Size),
	}
	deployments, err := client.AppsV1().Deployments(req.Namespace).List(ctx, opts)
	if err != nil {
		s.logger.Errorf("get deployment list error: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = "获取deployment列表出错!"
		return response
	}

	for _, deployment := range deployments.Items {
		deploy := models.Deployment{
			Name:              deployment.Name,
			Namespace:         deployment.Namespace,
			Labels:            deployment.Labels,
			Replicas:          deployment.Status.Replicas,
			UpdatedReplicas:   deployment.Status.UpdatedReplicas,
			AvailableReplicas: deployment.Status.AvailableReplicas,
			CreateTime:        deployment.CreationTimestamp.Time.Format("2006-01-02 15:04:05"),
		}
		deploymnetList = append(deploymnetList, deploy)
	}
	response.Code = http.StatusOK
	response.Data = map[string]interface{}{
		"configmaps": deploymnetList,
		"page":       req.Page,
		"size":       req.Size,
		"total":      len(total.Items),
	}
	return response
}
