package pv

import (
	"context"
	"fmt"
	models "github.com/xiaoxin1992/kube-admin/models/pv"
	"github.com/xiaoxin1992/kube-admin/pkg"
	"github.com/xiaoxin1992/kube-admin/services/k8s"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

func (s *Services) ListPV(ctx context.Context, req models.QueryList) models.Response {
	response := models.Response{}
	PVSlice := make([]models.PV, 0)
	client, err := k8s.NewService().GetClient(ctx, req.Zone)
	if err != nil {
		s.logger.Errorf("get k8s client err: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = fmt.Sprintf("获取 k8s 客户端出错: %v!", err)
		return response
	}
	total, err := client.CoreV1().PersistentVolumes().List(ctx, metav1.ListOptions{})
	if err != nil {
		s.logger.Errorf("get k8s total pv err: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = fmt.Sprintf("获取 pv总数出错: %v!", err)
		return response
	}
	opts := metav1.ListOptions{
		Limit: int64(req.Size * req.Page),
	}
	pvs, err := client.CoreV1().PersistentVolumes().List(ctx, opts)
	if err != nil {
		s.logger.Errorf("get k8s pv err: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = fmt.Sprintf("获取 pv 出错: %v!", err)
		return response
	}
	pvItems := pvs.Items
	offset, limits := pkg.Page(req.Page, req.Size, len(total.Items))
	pvItems = pvs.Items[offset:limits]
	for _, pv := range pvItems {
		fmt.Println(pv)
		pvTemp := models.PV{
			Name:        pv.Name,
			Labels:      make(map[string]string),
			Capacity:    pv.Spec.Capacity,
			AccessModes: pv.Spec.AccessModes,
			VolumeMode:  pv.Spec.VolumeMode,
			Phase:       pv.Status.Phase,
			CreateTime:  pv.CreationTimestamp.Format("2006-01-02 15:04:05"),
		}
		if len(pv.Labels) > 0 {
			pvTemp.Labels = pv.Labels
		}
		PVSlice = append(PVSlice, pvTemp)
	}

	response.Code = http.StatusOK
	response.Data = map[string]interface{}{
		"pv":    PVSlice,
		"page":  req.Page,
		"size":  req.Size,
		"total": len(total.Items),
	}
	return response
}
