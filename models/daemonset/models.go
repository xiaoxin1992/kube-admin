package daemonset

import (
	appsv1 "k8s.io/api/apps/v1"
)

type DaemonSet struct {
	Name                   string            `json:"name"`
	Namespace              string            `json:"namespace"`
	Labels                 map[string]string `json:"labels"`
	CurrentNumberScheduled int32             `json:"currentNumberScheduled"`
	CreateTime             string            `json:"createTime"`
}
type CreateDaemonSet struct {
	Zone      string           `json:"Zone" form:"Zone" binding:"required"`
	DaemonSet appsv1.DaemonSet `json:"daemonset" form:"daemonset" binding:"required"`
}

type QueryList struct {
	Zone      string `json:"zone" form:"zone" binding:"required"`
	Namespace string `json:"namespace" form:"namespace" binding:"required"`
	Size      int    `json:"size,omitempty" form:"size" binding:"required"`
	Page      int    `json:"page,omitempty" form:"page" binding:"required"`
}

type DeleteDaemonSet struct {
	Zone      string `json:"zone" form:"zone" binding:"required"`
	Namespace string `json:"namespace" form:"namespace" binding:"required"`
	Name      string `json:"name" form:"name" binding:"required"`
}

type DetailQuery struct {
	Zone      string `form:"zone" json:"zone" binding:"required"`
	Namespace string `form:"namespace" json:"namespace" binding:"required"`
	Name      string `form:"name" json:"name" binding:"required"`
}

type UpdateDaemonSet struct {
	Zone      string           `form:"zone" json:"zone" binding:"required"`
	DaemonSet appsv1.DaemonSet `form:"daemonset" json:"daemonset" binding:"required"`
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
