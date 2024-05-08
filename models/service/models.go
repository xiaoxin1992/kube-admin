package service

import (
	corev1 "k8s.io/api/core/v1"
)

type Service struct {
	Name       string               `json:"name"`
	Namespace  string               `json:"namespace"`
	Labels     map[string]string    `json:"labels"`
	Port       []corev1.ServicePort `json:"port"`
	Type       corev1.ServiceType   `json:"type"`
	CreateTime string               `json:"CreateTime"`
}
type CreateService struct {
	Zone    string         `json:"Zone" form:"Zone" binding:"required"`
	Service corev1.Service `json:"service" form:"service" binding:"required"`
}

type QueryList struct {
	Zone      string `json:"zone" form:"zone" binding:"required"`
	Namespace string `json:"namespace" form:"namespace" binding:"required"`
	Size      int    `json:"size,omitempty" form:"size" binding:"required"`
	Page      int    `json:"page,omitempty" form:"page" binding:"required"`
}

type DeleteService struct {
	Zone      string `json:"zone" form:"zone" binding:"required"`
	Namespace string `json:"namespace" form:"namespace" binding:"required"`
	Name      string `json:"name" form:"name" binding:"required"`
}

type DetailQuery struct {
	Zone      string `form:"zone" json:"zone" binding:"required"`
	Namespace string `form:"namespace" json:"namespace" binding:"required"`
	Name      string `form:"name" json:"name" binding:"required"`
}

type UpdateService struct {
	Zone    string         `form:"zone" json:"zone" binding:"required"`
	Service corev1.Service `form:"service" json:"service" binding:"required"`
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
