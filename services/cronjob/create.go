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

func (s *Services) CreateCronjob(ctx context.Context, req models.CreateCronjob) models.Response {
	response := models.Response{}
	client, err := k8s.NewService().GetClient(ctx, req.Zone)
	if err != nil {
		s.logger.Errorf("get k8s client err: %+v", err)
		response.Message = fmt.Sprintf("获取k8s客户端出错: %v", err)
		response.Code = http.StatusBadRequest
		return response
	}
	_, err = client.BatchV1().CronJobs(req.Cronjob.Namespace).Get(ctx, req.Cronjob.Name, metav1.GetOptions{})
	if err == nil {
		response.Message = fmt.Sprintf("cronjob %s 已经存在", req.Cronjob.Name)
		response.Code = http.StatusBadRequest
		return response
	} else {
		if !k8sError.IsNotFound(err) {
			s.logger.Errorf("get cronjob err: %v", err)
			response.Message = fmt.Sprintf("查询cronjob %s 出错: %v", req.Cronjob.Name, err)
			response.Code = http.StatusBadRequest
			return response
		}
	}
	_, err = client.BatchV1().CronJobs(req.Cronjob.Namespace).Create(ctx, &req.Cronjob, metav1.CreateOptions{DryRun: []string{metav1.DryRunAll}})
	if err != nil {
		s.logger.Errorf("create cronjob syntax err: %v", err)
		response.Message = fmt.Sprintf("创建cronjob %s 语法错误: %v", req.Cronjob.Name, err)
		response.Code = http.StatusBadRequest
		return response
	}
	_, err = client.BatchV1().CronJobs(req.Cronjob.Namespace).Create(ctx, &req.Cronjob, metav1.CreateOptions{})
	if err != nil {
		s.logger.Errorf("create cronjob err: %v", err)
		response.Message = fmt.Sprintf("创建cronjob %s 出错: %v", req.Cronjob.Name, err)
		response.Code = http.StatusBadRequest
		return response
	}
	response.Message = fmt.Sprintf("创建cronjob %s 完成", req.Cronjob.Name)
	response.Code = http.StatusOK
	return response
}
