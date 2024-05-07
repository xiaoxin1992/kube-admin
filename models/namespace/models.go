package namespace

import v1 "k8s.io/api/core/v1"

type Namespace struct {
	Name  string            `json:"name" form:"name" json:"name" binding:"required"`
	Label map[string]string `json:"label" form:"label"`
}

type ListNamespace struct {
	Namespace
	Status     string `json:"status"`
	CreateTime string `json:"createTime"`
}

type CreateNamespace struct {
	Zone      string       `json:"zone" form:"zone" binding:"required"`
	Namespace v1.Namespace `form:"namespace" binding:"required"`
}

type DeleteNamespace struct {
	Zone string `json:"zone" form:"zone" binding:"required"`
	Name string `json:"name" form:"name" binding:"required"`
}

type QueryList struct {
	Zone string `json:"zone" form:"zone" binding:"required"`
	Size int    `json:"size,omitempty" form:"size" binding:"required"`
	Page int    `json:"page,omitempty" form:"page" binding:"required"`
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
