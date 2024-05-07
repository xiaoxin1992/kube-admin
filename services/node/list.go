package node

import (
	"context"
	models "github.com/xiaoxin1992/kube-admin/models/node"
	"github.com/xiaoxin1992/kube-admin/services/k8s"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

func (s *Services) ListNode(ctx context.Context, req models.QueryList) models.Response {
	response := models.Response{}
	nodeList := make([]models.ListNode, 0)
	client, err := k8s.NewService().GetClient(ctx, req.Zone)
	if err != nil {
		s.logger.Errorf("get k8s client error: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = "获取k8s客户端出错!"
		return response
	}
	total, err := client.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		s.logger.Errorf("get k8s node total list error: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = "获取node总数出错"
		return response
	}
	opts := metav1.ListOptions{
		Limit: int64(req.Page * req.Size),
	}
	nodes, err := client.CoreV1().Nodes().List(ctx, opts)
	if err != nil {
		s.logger.Errorf("get namespace list error: %v", err)
		response.Code = http.StatusInternalServerError
		response.Message = "获取namespace列表出错!"
		return response
	}

	for _, node := range nodes.Items {
		nList := models.ListNode{
			Name:       node.Name,
			Labels:     node.GetLabels(),
			Taints:     make([]v1.Taint, 0),
			CreateTime: node.CreationTimestamp.Time.Format("2006-01-02 15:04:05"),
		}
		if len(node.Spec.Taints) > 0 {
			nList.Taints = node.Spec.Taints
		}
		for _, addr := range node.Status.Addresses {
			switch addr.Type {
			case v1.NodeInternalIP:
				nList.Address = addr.Address
			case v1.NodeHostName:
				nList.HostName = addr.Address
			}
		}
		for _, condition := range node.Status.Conditions {
			if condition.Type == v1.NodeReady {
				nList.Status = condition.Status
			}
		}
		nodeList = append(nodeList, nList)
	}
	response.Code = http.StatusOK
	response.Data = map[string]interface{}{
		"nodes": nodeList,
		"page":  req.Page,
		"size":  req.Size,
		"total": len(total.Items),
	}
	return response
}
