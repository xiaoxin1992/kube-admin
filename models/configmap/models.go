package configmap

import v1 "k8s.io/api/core/v1"

type CreateConfigMap struct {
	Zone      string       `form:"zone" json:"zone" binding:"required"`
	ConfigMap v1.ConfigMap `form:"configmap" json:"configmap" binding:"required"`
}

type QueryList struct {
	Zone      string `json:"zone" form:"zone" binding:"required"`
	Namespace string `form:"namespace" json:"namespace" binding:"required"`
	Size      int    `json:"size,omitempty" form:"size" binding:"required"`
	Page      int    `json:"page,omitempty" form:"page" binding:"required"`
}

type ListConfigMap struct {
	Name        string            `json:"name"`
	Namespace   string            `json:"namespace"`
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
	CreateTime  string            `json:"createTime"`
}

type DetailQuery struct {
	Zone      string `form:"zone" json:"zone" binding:"required"`
	Namespace string `form:"namespace" json:"namespace" binding:"required"`
	Name      string `form:"name" json:"name" binding:"required"`
}

type UpdateConfigmap struct {
	Zone      string       `form:"zone" json:"zone" binding:"required"`
	ConfigMap v1.ConfigMap `form:"configmap" json:"configMap" binding:"required"`
}

type DeleteConfigmap struct {
	Zone      string `form:"zone" json:"zone" binding:"required"`
	Namespace string `form:"namespace" json:"namespace" binding:"required"`
	Name      string `form:"name" json:"name" binding:"required"`
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
