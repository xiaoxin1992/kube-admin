package daemonset

import (
	"context"
	"fmt"
	models "github.com/xiaoxin1992/kube-admin/models/daemonset"
	"github.com/xiaoxin1992/kube-admin/services/k8s"
	k8sError "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

func (s *Services) CreateDaemonSet(ctx context.Context, req models.CreateDaemonSet) models.Response {
	response := models.Response{}
	client, err := k8s.NewService().GetClient(ctx, req.Zone)
	if err != nil {
		s.logger.Errorf("get k8s client err: %+v", err)
		response.Message = fmt.Sprintf("获取k8s客户端出错: %v", err)
		response.Code = http.StatusBadRequest
		return response
	}
	_, err = client.AppsV1().DaemonSets(req.DaemonSet.Namespace).Get(ctx, req.DaemonSet.Name, metav1.GetOptions{})
	if err == nil {
		response.Message = fmt.Sprintf("daemonSet %s 已经存在", req.DaemonSet.Name)
		response.Code = http.StatusBadRequest
		return response
	} else {
		if !k8sError.IsNotFound(err) {
			s.logger.Errorf("get daemonSet err: %v", err)
			response.Message = fmt.Sprintf("查询daemonSet %s 出错: %v", req.DaemonSet.Name, err)
			response.Code = http.StatusBadRequest
			return response
		}
	}
	_, err = client.AppsV1().DaemonSets(req.DaemonSet.Namespace).Create(ctx, &req.DaemonSet, metav1.CreateOptions{DryRun: []string{metav1.DryRunAll}})
	if err != nil {
		s.logger.Errorf("create daemonSet syntax err: %v", err)
		response.Message = fmt.Sprintf("创建daemonSet %s 语法错误: %v", req.DaemonSet.Name, err)
		response.Code = http.StatusBadRequest
		return response
	}
	_, err = client.AppsV1().DaemonSets(req.DaemonSet.Namespace).Create(ctx, &req.DaemonSet, metav1.CreateOptions{})
	if err != nil {
		s.logger.Errorf("create daemonSet err: %v", err)
		response.Message = fmt.Sprintf("创建daemonSet %s 出错: %v", req.DaemonSet.Name, err)
		response.Code = http.StatusBadRequest
		return response
	}
	response.Message = fmt.Sprintf("创建daemonSet %s 完成", req.DaemonSet.Name)
	response.Code = http.StatusOK
	return response
}
