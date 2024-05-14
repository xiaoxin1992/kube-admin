package statefulset

import (
	"context"
	models "github.com/xiaoxin1992/kube-admin/models/statefulset"
	"github.com/xiaoxin1992/kube-admin/services/k8s"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

func (s *Services) ListStateFulSet(ctx context.Context, req models.QueryList) models.Response {
	response := models.Response{}
	stsList := make([]models.StateFulSet, 0)
	client, err := k8s.NewService().GetClient(ctx, req.Zone)
	if err != nil {
		s.logger.Errorf("get k8s client error: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = "获取k8s客户端出错!"
		return response
	}
	total, err := client.AppsV1().StatefulSets(req.Namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		s.logger.Errorf("get StatefulSet total list error: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = "获取StatefulSet总数出错"
		return response
	}
	opts := metav1.ListOptions{
		Limit: int64(req.Page * req.Size),
	}
	stss, err := client.AppsV1().StatefulSets(req.Namespace).List(ctx, opts)
	if err != nil {
		s.logger.Errorf("get StatefulSet list error: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = "获取StatefulSet列表出错!"
		return response
	}
	offset := (req.Page - 1) * req.Size
	stsItems := stss.Items
	if req.Page*req.Size <= len(total.Items) {
		stsItems = stss.Items[offset:]
	}
	for _, sts := range stsItems {
		st := models.StateFulSet{
			Name:          sts.Name,
			Namespace:     sts.Namespace,
			Labels:        make(map[string]string),
			Replicas:      sts.Status.Replicas,
			ReadyReplicas: sts.Status.ReadyReplicas,
			CreateTime:    sts.CreationTimestamp.Time.Format("2006-01-02 15:04:05"),
		}
		if len(sts.Labels) > 0 {
			st.Labels = sts.Labels
		}
		stsList = append(stsList, st)
	}
	response.Code = http.StatusOK
	response.Data = map[string]interface{}{
		"StatefulSet": stsList,
		"page":        req.Page,
		"size":        req.Size,
		"total":       len(total.Items),
	}
	return response
}
