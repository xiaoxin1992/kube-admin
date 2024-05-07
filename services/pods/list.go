package pods

import (
	"context"
	models "github.com/xiaoxin1992/kube-admin/models/pods"
	pkgK8s "github.com/xiaoxin1992/kube-admin/pkg/k8s"
	"github.com/xiaoxin1992/kube-admin/services/k8s"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

func (s *Services) ListPods(ctx context.Context, req models.QueryList) models.Response {
	// 如果namespace 不传递则返回集群内所有的Pod列表
	response := models.Response{}
	k := k8s.NewService()
	client, err := k.GetClient(ctx, req.Zone)
	if err != nil {
		s.logger.Errorf("get k8s client error: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = "获取集群客户端出错"
		return response
	}
	total, err := client.CoreV1().Pods(req.Namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		s.logger.Errorf("get pods list total error: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = "获取pod总数出错!"
		return response
	}
	opts := metav1.ListOptions{
		Limit: int64(req.Page * req.Size),
	}
	pods, err := client.CoreV1().Pods(req.Namespace).List(ctx, opts)
	if err != nil {
		s.logger.Errorf("get pods error: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = "获取Pod列表出错!"
		return response
	}
	offset := (req.Page - 1) * req.Size
	podItems := pods.Items
	if req.Page*req.Size < len(total.Items) {
		podItems = pods.Items[offset:]
	}
	podList := make([]interface{}, 0)
	for _, pod := range podItems {
		p := pkgK8s.PodAnalysis(&pod)
		podList = append(podList, p)
	}
	response.Code = http.StatusOK
	response.Data = map[string]interface{}{
		"pods":  podList,
		"page":  req.Page,
		"size":  req.Size,
		"total": len(total.Items),
	}
	return response
}
