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

func (s *Services) UpdateDaemonSet(ctx context.Context, req models.UpdateDaemonSet) models.Response {
	response := models.Response{}
	client, err := k8s.NewService().GetClient(ctx, req.Zone)
	if err != nil {
		s.logger.Errorf("get k8s client err: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = fmt.Sprintf("获取k8s客户端出错%v!", err)
		return response
	}
	_, err = client.AppsV1().DaemonSets(req.DaemonSet.Namespace).Get(ctx, req.DaemonSet.Name, metav1.GetOptions{})
	if err != nil {
		response.Code = http.StatusBadRequest
		if k8sError.IsNotFound(err) {
			response.Message = fmt.Sprintf("daemonSet %s 不存在", req.DaemonSet.Name)
		} else {
			response.Message = fmt.Sprintf("daemonSet %s 查询出错 %v", req.DaemonSet.Name, err)
		}
		return response
	}
	_, err = client.AppsV1().DaemonSets(req.DaemonSet.Namespace).Update(ctx, &req.DaemonSet, metav1.UpdateOptions{DryRun: []string{metav1.DryRunAll}})
	if err != nil {
		s.logger.Errorf("daemonSet %s syntax error: %v", req.DaemonSet.Name, err)
		response.Code = http.StatusBadRequest
		response.Message = fmt.Sprintf("daemonSet语法错误%v!", err)
		return response
	}
	_, err = client.AppsV1().DaemonSets(req.DaemonSet.Namespace).Update(ctx, &req.DaemonSet, metav1.UpdateOptions{})
	if err != nil {
		s.logger.Errorf("update daemonSet %s error: %v", req.DaemonSet.Name, err)
		response.Code = http.StatusBadRequest
		response.Message = fmt.Sprintf("更新daemonSet %s 出错 %v", req.DaemonSet.Name, err)
	} else {
		response.Code = http.StatusOK
		response.Message = fmt.Sprintf("更新daemonSet%s 完成", req.DaemonSet.Name)
	}
	return response
}
