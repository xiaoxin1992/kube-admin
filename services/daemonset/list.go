package daemonset

import (
	"context"
	models "github.com/xiaoxin1992/kube-admin/models/daemonset"
	"github.com/xiaoxin1992/kube-admin/pkg"
	"github.com/xiaoxin1992/kube-admin/services/k8s"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

func (s *Services) ListDaemonSet(ctx context.Context, req models.QueryList) models.Response {
	response := models.Response{}
	daemonSetList := make([]models.DaemonSet, 0)
	client, err := k8s.NewService().GetClient(ctx, req.Zone)
	if err != nil {
		s.logger.Errorf("get k8s client error: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = "获取k8s客户端出错!"
		return response
	}
	total, err := client.AppsV1().DaemonSets(req.Namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		s.logger.Errorf("get daemonSet total list error: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = "获取daemonSet总数出错"
		return response
	}
	opts := metav1.ListOptions{
		Limit: int64(req.Page * req.Size),
	}
	daemonSets, err := client.AppsV1().DaemonSets(req.Namespace).List(ctx, opts)
	if err != nil {
		s.logger.Errorf("get daemonSet list error: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = "获取daemonSet列表出错!"
		return response
	}
	daemonSetItems := daemonSets.Items
	offset, limits := pkg.Page(req.Page, req.Size, len(total.Items))
	daemonSetItems = daemonSets.Items[offset:limits]
	for _, ds := range daemonSetItems {
		d := models.DaemonSet{
			Name:                   ds.Name,
			Namespace:              ds.Namespace,
			Labels:                 make(map[string]string),
			CurrentNumberScheduled: ds.Status.CurrentNumberScheduled,
			CreateTime:             ds.CreationTimestamp.Time.Format("2006-01-02 15:04:05"),
		}
		if len(ds.Labels) > 0 {
			d.Labels = ds.Labels
		}
		daemonSetList = append(daemonSetList, d)
	}
	response.Code = http.StatusOK
	response.Data = map[string]interface{}{
		"daemonsets": daemonSetList,
		"page":       req.Page,
		"size":       req.Size,
		"total":      len(total.Items),
	}
	return response
}
