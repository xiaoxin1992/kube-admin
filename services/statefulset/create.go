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

func (s *Services) CreateStateFulSet(ctx context.Context, req models.CreateStateFulSet) models.Response {
	response := models.Response{}
	client, err := k8s.NewService().GetClient(ctx, req.Zone)
	if err != nil {
		s.logger.Errorf("get k8s client err: %+v", err)
		response.Message = fmt.Sprintf("获取k8s客户端出错: %v", err)
		response.Code = http.StatusBadRequest
		return response
	}
	_, err = client.AppsV1().StatefulSets(req.StateFulSet.Namespace).Get(ctx, req.StateFulSet.Name, metav1.GetOptions{})
	if err == nil {
		response.Message = fmt.Sprintf("StatefulSet %s 已经存在", req.StateFulSet.Name)
		response.Code = http.StatusBadRequest
		return response
	} else {
		if !k8sError.IsNotFound(err) {
			s.logger.Errorf("get StatefulSet err: %v", err)
			response.Message = fmt.Sprintf("查询StatefulSet %s 出错: %v", req.StateFulSet.Name, err)
			response.Code = http.StatusBadRequest
			return response
		}
	}
	_, err = client.AppsV1().StatefulSets(req.StateFulSet.Namespace).Create(ctx, &req.StateFulSet, metav1.CreateOptions{DryRun: []string{metav1.DryRunAll}})
	if err != nil {
		s.logger.Errorf("create StatefulSet syntax err: %v", err)
		response.Message = fmt.Sprintf("创建StatefulSet %s 语法错误: %v", req.StateFulSet.Name, err)
		response.Code = http.StatusBadRequest
		return response
	}
	_, err = client.AppsV1().StatefulSets(req.StateFulSet.Namespace).Create(ctx, &req.StateFulSet, metav1.CreateOptions{})
	if err != nil {
		s.logger.Errorf("create StatefulSet err: %v", err)
		response.Message = fmt.Sprintf("创建StatefulSet %s 出错: %v", req.StateFulSet.Name, err)
		response.Code = http.StatusBadRequest
		return response
	}
	response.Message = fmt.Sprintf("创建StatefulSet %s 完成", req.StateFulSet.Name)
	response.Code = http.StatusOK
	return response
}
