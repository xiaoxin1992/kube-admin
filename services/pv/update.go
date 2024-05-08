package pv

import (
	"context"
	"fmt"
	models "github.com/xiaoxin1992/kube-admin/models/pv"
	"github.com/xiaoxin1992/kube-admin/services/k8s"
	k8sError "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

func (s *Services) UpdatePV(ctx context.Context, req models.UpdatePV) models.Response {
	response := models.Response{}
	client, err := k8s.NewService().GetClient(ctx, req.Zone)
	if err != nil {
		s.logger.Errorf("get k8s client err: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = fmt.Sprintf("获取k8s客户端出错%v!", err)
		return response
	}
	_, err = client.CoreV1().PersistentVolumes().Get(ctx, req.PV.Name, metav1.GetOptions{})
	if err != nil {
		response.Code = http.StatusBadRequest
		if k8sError.IsNotFound(err) {
			response.Message = fmt.Sprintf("pv %s 不存在", req.PV.Name)
		} else {
			response.Message = fmt.Sprintf("pv %s 查询出错 %v", req.PV.Name, err)
		}
		return response
	}
	_, err = client.CoreV1().PersistentVolumes().Update(ctx, &req.PV, metav1.UpdateOptions{DryRun: []string{metav1.DryRunAll}})
	if err != nil {
		s.logger.Errorf(" pv %s syntax error: %v", req.PV.Name, err)
		response.Code = http.StatusBadRequest
		response.Message = fmt.Sprintf("pv语法错误%v!", err)
		return response
	}
	_, err = client.CoreV1().PersistentVolumes().Update(ctx, &req.PV, metav1.UpdateOptions{})
	if err != nil {
		s.logger.Errorf("update pv %s error: %v", req.PV.Name, err)
		response.Code = http.StatusBadRequest
		response.Message = fmt.Sprintf("更新pv %s 出错 %v", req.PV.Name, err)
	} else {
		response.Code = http.StatusOK
		response.Message = fmt.Sprintf("更新pv %s 完成", req.PV.Name)
	}
	return response
}
