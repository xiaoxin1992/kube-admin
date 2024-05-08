package cronjob

import (
	"context"
	"fmt"
	models "github.com/xiaoxin1992/kube-admin/models/cronjob"
	"github.com/xiaoxin1992/kube-admin/services/k8s"
	k8sError "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

func (s *Services) UpdateCronjob(ctx context.Context, req models.UpdateCronjob) models.Response {
	response := models.Response{}
	client, err := k8s.NewService().GetClient(ctx, req.Zone)
	if err != nil {
		s.logger.Errorf("get k8s client err: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = fmt.Sprintf("获取k8s客户端出错%v!", err)
		return response
	}
	_, err = client.BatchV1().CronJobs(req.Cronjob.Namespace).Get(ctx, req.Cronjob.Name, metav1.GetOptions{})
	if err != nil {
		response.Code = http.StatusBadRequest
		if k8sError.IsNotFound(err) {
			response.Message = fmt.Sprintf("cronjob %s 不存在", req.Cronjob.Name)
		} else {
			response.Message = fmt.Sprintf("cronjob %s 查询出错 %v", req.Cronjob.Name, err)
		}
		return response
	}
	_, err = client.BatchV1().CronJobs(req.Cronjob.Namespace).Update(ctx, &req.Cronjob, metav1.UpdateOptions{DryRun: []string{metav1.DryRunAll}})
	if err != nil {
		s.logger.Errorf("cronjob %s syntax error: %v", req.Cronjob.Name, err)
		response.Code = http.StatusBadRequest
		response.Message = fmt.Sprintf("cronjob语法错误%v!", err)
		return response
	}
	_, err = client.BatchV1().CronJobs(req.Cronjob.Namespace).Update(ctx, &req.Cronjob, metav1.UpdateOptions{})
	if err != nil {
		s.logger.Errorf("update cronjob %s error: %v", req.Cronjob.Name, err)
		response.Code = http.StatusBadRequest
		response.Message = fmt.Sprintf("更新cronjob %s 出错 %v", req.Cronjob.Name, err)
	} else {
		response.Code = http.StatusOK
		response.Message = fmt.Sprintf("更新cronjob%s 完成", req.Cronjob.Name)
	}
	return response
}
