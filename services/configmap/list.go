package configmap

import (
	"context"
	models "github.com/xiaoxin1992/kube-admin/models/configmap"
	"github.com/xiaoxin1992/kube-admin/pkg"
	"github.com/xiaoxin1992/kube-admin/services/k8s"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

func (s *Services) ListConfigmap(ctx context.Context, req models.QueryList) models.Response {
	response := models.Response{}
	//configmapList := make([]models.ListConfigMap, 0)
	client, err := k8s.NewService().GetClient(ctx, req.Zone)
	if err != nil {
		s.logger.Errorf("get k8s client error: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = "获取k8s客户端出错!"
		return response
	}
	total, err := client.CoreV1().ConfigMaps(req.Namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		s.logger.Errorf("get configmap total list error: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = "获取configmap总数出错"
		return response
	}
	opts := metav1.ListOptions{
		Limit: int64(req.Page * req.Size),
	}
	configmaps, err := client.CoreV1().ConfigMaps(req.Namespace).List(ctx, opts)
	if err != nil {
		s.logger.Errorf("get configmap list error: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = "获取configmap列表出错!"
		return response
	}
	configmapItems := configmaps.Items
	offset, limits := pkg.Page(req.Page, req.Size, len(total.Items))
	configmapItems = configmaps.Items[offset:limits]
	//for _, configmap := range configmapItems {
	//	cmp := models.ListConfigMap{
	//		Name:        configmap.Name,
	//		Namespace:   configmap.Namespace,
	//		Labels:      make(map[string]string),
	//		Annotations: make(map[string]string),
	//		CreateTime:  configmap.CreationTimestamp.Time.Format("2006-01-02 15:04:05"),
	//	}
	//	if len(configmap.Labels) > 0 {
	//		cmp.Labels = configmap.Labels
	//	}
	//	if len(configmap.Annotations) > 0 {
	//		cmp.Annotations = configmap.Annotations
	//	}
	//	configmapList = append(configmapList, cmp)
	//}
	response.Code = http.StatusOK
	response.Data = map[string]interface{}{
		"configmaps": configmapItems,
		"page":       req.Page,
		"size":       req.Size,
		"total":      len(total.Items),
	}
	return response
}
