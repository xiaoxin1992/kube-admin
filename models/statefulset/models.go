package statefulset

import (
	appsv1 "k8s.io/api/apps/v1"
)

type StateFulSet struct {
	Name          string            `json:"name"`
	Namespace     string            `json:"namespace"`
	Labels        map[string]string `json:"labels"`
	Replicas      int32             `json:"replicas"`
	ReadyReplicas int32             `json:"readyReplicas"`
	CreateTime    string            `json:"createTime"`
}
type CreateStateFulSet struct {
	Zone        string             `json:"Zone" form:"Zone" binding:"required"`
	StateFulSet appsv1.StatefulSet `json:"sts" form:"sts" binding:"required"`
}

type QueryList struct {
	Zone      string `json:"zone" form:"zone" binding:"required"`
	Namespace string `json:"namespace" form:"namespace" binding:"required"`
	Size      int    `json:"size,omitempty" form:"size" binding:"required"`
	Page      int    `json:"page,omitempty" form:"page" binding:"required"`
}

type DeleteStateFulSet struct {
	Zone      string `json:"zone" form:"zone" binding:"required"`
	Namespace string `json:"namespace" form:"namespace" binding:"required"`
	Name      string `json:"name" form:"name" binding:"required"`
}

type DetailQuery struct {
	Zone      string `form:"zone" json:"zone" binding:"required"`
	Namespace string `form:"namespace" json:"namespace" binding:"required"`
	Name      string `form:"name" json:"name" binding:"required"`
}

type UpdateStateFulSet struct {
	Zone        string             `form:"zone" json:"zone" binding:"required"`
	StateFulSet appsv1.StatefulSet `form:"sts" json:"sts" binding:"required"`
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
