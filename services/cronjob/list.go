package cronjob

import (
	"context"
	models "github.com/xiaoxin1992/kube-admin/models/cronjob"
	"github.com/xiaoxin1992/kube-admin/pkg"
	"github.com/xiaoxin1992/kube-admin/services/k8s"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

func (s *Services) ListCronjob(ctx context.Context, req models.QueryList) models.Response {
	response := models.Response{}
	cronjobList := make([]models.Cronjob, 0)
	client, err := k8s.NewService().GetClient(ctx, req.Zone)
	if err != nil {
		s.logger.Errorf("get k8s client error: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = "获取k8s客户端出错!"
		return response
	}
	total, err := client.BatchV1().CronJobs(req.Namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		s.logger.Errorf("get cronjob total list error: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = "获取cronjob总数出错"
		return response
	}
	opts := metav1.ListOptions{
		Limit: int64(req.Page * req.Size),
	}
	cronjobs, err := client.BatchV1().CronJobs(req.Namespace).List(ctx, opts)
	if err != nil {
		s.logger.Errorf("get cronjob list error: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = "获取cronjob列表出错!"
		return response
	}
	cronjobItems := cronjobs.Items
	offset, limits := pkg.Page(req.Page, req.Size, len(total.Items))
	cronjobItems = cronjobs.Items[offset:limits]
	for _, cronjob := range cronjobItems {
		cron := models.Cronjob{
			Name:               cronjob.Name,
			Namespace:          cronjob.Namespace,
			Labels:             map[string]string{},
			Schedule:           cronjob.Spec.Schedule,
			LastScheduleTime:   cronjob.Status.LastScheduleTime.Time.Format("2006-01-02 15:04:05"),
			LastSuccessfulTime: cronjob.Status.LastSuccessfulTime.Time.Format("2006-01-02 15:04:05"),
			CreateTime:         cronjob.CreationTimestamp.Time.Format("2006-01-02 15:04:05"),
		}
		if len(cronjob.Labels) > 0 {
			cron.Labels = cronjob.Labels
		}
		cronjobList = append(cronjobList, cron)
	}
	response.Code = http.StatusOK
	response.Data = map[string]interface{}{
		"cronjob": cronjobList,
		"page":    req.Page,
		"size":    req.Size,
		"total":   len(total.Items),
	}
	return response
}
