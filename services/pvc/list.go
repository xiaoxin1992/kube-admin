package pvc

import (
	"context"
	"fmt"
	models "github.com/xiaoxin1992/kube-admin/models/pvc"
	"github.com/xiaoxin1992/kube-admin/services/k8s"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

func (s *Services) ListPersistentVolumeClaim(ctx context.Context, req models.QueryList) models.Response {
	response := models.Response{}
	pvcSlice := make([]models.PersistentVolumeClaim, 0)
	client, err := k8s.NewService().GetClient(ctx, req.Zone)
	if err != nil {
		s.logger.Errorf("get k8s client err: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = fmt.Sprintf("获取 k8s 客户端出错: %v!", err)
		return response
	}
	total, err := client.CoreV1().PersistentVolumeClaims(req.Namspace).List(ctx, metav1.ListOptions{})
	if err != nil {
		s.logger.Errorf("get k8s total pvc err: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = fmt.Sprintf("获取 pvc总数出错: %v!", err)
		return response
	}
	opts := metav1.ListOptions{
		Limit: int64(req.Size * req.Page),
	}
	pvcs, err := client.CoreV1().PersistentVolumeClaims(req.Namspace).List(ctx, opts)
	if err != nil {
		s.logger.Errorf("get k8s pvc err: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = fmt.Sprintf("获取 pvc 出错: %v!", err)
		return response
	}
	offset := (req.Page - 1) * req.Size
	pvcItems := pvcs.Items
	if req.Page*req.Size <= len(total.Items) {
		pvcItems = pvcs.Items[offset:]
	}
	for _, pvc := range pvcItems {
		pvcTmp := models.PersistentVolumeClaim{
			Name:        pvc.Name,
			Labels:      make(map[string]string),
			AccessModes: pvc.Spec.AccessModes,
			VolumeMode:  pvc.Spec.VolumeMode,
			Phase:       pvc.Status.Phase,
			Resources:   pvc.Spec.Resources,
			CreateTime:  pvc.CreationTimestamp.Format("2006-01-02 15:04:05"),
		}
		if len(pvc.Labels) > 0 {
			pvcTmp.Labels = pvc.Labels
		}
		pvcSlice = append(pvcSlice, pvcTmp)
	}

	response.Code = http.StatusOK
	response.Data = map[string]interface{}{
		"pvc":   pvcSlice,
		"page":  req.Page,
		"size":  req.Size,
		"total": len(total.Items),
	}
	return response
}
