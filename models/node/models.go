package node

import v1 "k8s.io/api/core/v1"

type QueryList struct {
	Size int    `json:"size,omitempty" form:"size" binding:"required"`
	Page int    `json:"page,omitempty" form:"page" binding:"required"`
	Zone string `json:"zone" form:"zone" binding:"required"`
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type ListNode struct {
	Name       string             `json:"name"`
	Labels     map[string]string  `json:"labels"`
	Taints     []v1.Taint         `json:"taints"`
	Address    string             `json:"address"`
	HostName   string             `json:"hostName"`
	Status     v1.ConditionStatus `json:"status"`
	CreateTime string             `json:"create_time"`
}
