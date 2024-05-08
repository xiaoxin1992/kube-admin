package secret

import (
	corev1 "k8s.io/api/core/v1"
)

type Secret struct {
	Name       string            `json:"name"`
	Namespace  string            `json:"namespace"`
	Labels     map[string]string `json:"labels"`
	Type       corev1.SecretType `json:"type"`
	CreateTime string            `json:"createTime"`
}
type CreateSecret struct {
	Zone   string        `json:"Zone" form:"Zone" binding:"required"`
	Secret corev1.Secret `json:"secret" form:"secret" binding:"required"`
}

type QueryList struct {
	Zone      string `json:"zone" form:"zone" binding:"required"`
	Namespace string `json:"namespace" form:"namespace" binding:"required"`
	Size      int    `json:"size,omitempty" form:"size" binding:"required"`
	Page      int    `json:"page,omitempty" form:"page" binding:"required"`
}

type DeleteSecret struct {
	Zone      string `json:"zone" form:"zone" binding:"required"`
	Namespace string `json:"namespace" form:"namespace" binding:"required"`
	Name      string `json:"name" form:"name" binding:"required"`
}

type DetailQuery struct {
	Zone      string `form:"zone" json:"zone" binding:"required"`
	Namespace string `form:"namespace" json:"namespace" binding:"required"`
	Name      string `form:"name" json:"name" binding:"required"`
}

type UpdateSecret struct {
	Zone   string        `form:"zone" json:"zone" binding:"required"`
	Secret corev1.Secret `form:"secret" json:"secret" binding:"required"`
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
