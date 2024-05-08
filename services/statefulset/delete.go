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

func (s *Services) DeleteStateFulSet(ctx context.Context, req models.DeleteStateFulSet) models.Response {
	response := models.Response{}
	client, err := k8s.NewService().GetClient(ctx, req.Zone)
	if err != nil {
		s.logger.Errorf("get k8s client err: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = fmt.Sprintf("获取 k8s 客户端出错: %v!", err)
		return response
	}
	_, err = client.AppsV1().StatefulSets(req.Namespace).Get(ctx, req.Name, metav1.GetOptions{})
	if err != nil {
		response.Code = http.StatusNotFound
		if k8sError.IsNotFound(err) {
			response.Message = fmt.Sprintf("StatefulSet %s 不存在!", req.Name)
		} else {
			response.Message = fmt.Sprintf("查询StatefulSet %s 出错: %v!", req.Name, err)
		}
		return response
	}
	err = client.AppsV1().StatefulSets(req.Namespace).Delete(ctx, req.Name, metav1.DeleteOptions{})
	if err != nil {
		s.logger.Errorf("delete k8s StatefulSet err: %v!", err)
		response.Code = http.StatusNotFound
		response.Message = fmt.Sprintf("删除StatefulSet %s 出错: %v!", req.Name, err)
		return response
	}
	response.Code = http.StatusOK
	response.Message = fmt.Sprintf("删除StatefulSet %s 完成!", req.Name)
	return response
}
