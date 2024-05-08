package statefulset

import (
	"context"
	"fmt"
	models "github.com/xiaoxin1992/kube-admin/models/statefulset"
	"github.com/xiaoxin1992/kube-admin/services/k8s"
	k8sError "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

func (s *Services) UpdateStateFulSet(ctx context.Context, req models.UpdateStateFulSet) models.Response {
	response := models.Response{}
	client, err := k8s.NewService().GetClient(ctx, req.Zone)
	if err != nil {
		s.logger.Errorf("get k8s client err: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = fmt.Sprintf("获取k8s客户端出错%v!", err)
		return response
	}
	_, err = client.AppsV1().StatefulSets(req.StateFulSet.Namespace).Get(ctx, req.StateFulSet.Name, metav1.GetOptions{})
	if err != nil {
		response.Code = http.StatusBadRequest
		if k8sError.IsNotFound(err) {
			response.Message = fmt.Sprintf("StatefulSet %s 不存在", req.StateFulSet.Name)
		} else {
			response.Message = fmt.Sprintf("StatefulSet %s 查询出错 %v", req.StateFulSet.Name, err)
		}
		return response
	}
	_, err = client.AppsV1().StatefulSets(req.StateFulSet.Namespace).Update(ctx, &req.StateFulSet, metav1.UpdateOptions{DryRun: []string{metav1.DryRunAll}})
	if err != nil {
		s.logger.Errorf("StatefulSet %s syntax error: %v", req.StateFulSet.Name, err)
		response.Code = http.StatusBadRequest
		response.Message = fmt.Sprintf("StatefulSet语法错误%v!", err)
		return response
	}
	_, err = client.AppsV1().StatefulSets(req.StateFulSet.Namespace).Update(ctx, &req.StateFulSet, metav1.UpdateOptions{})
	if err != nil {
		s.logger.Errorf("update daemonSet %s error: %v", req.StateFulSet.Name, err)
		response.Code = http.StatusBadRequest
		response.Message = fmt.Sprintf("更新StatefulSet %s 出错 %v", req.StateFulSet.Name, err)
	} else {
		response.Code = http.StatusOK
		response.Message = fmt.Sprintf("更新StatefulSet%s 完成", req.StateFulSet.Name)
	}
	return response
}
