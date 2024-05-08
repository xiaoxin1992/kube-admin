package pv

import (
	corev1 "k8s.io/api/core/v1"
)

type PV struct {
	Name        string                              `json:"name"`
	Labels      map[string]string                   `json:"labels"`
	Capacity    corev1.ResourceList                 `json:"capacity"`
	AccessModes []corev1.PersistentVolumeAccessMode `json:"accessModes"`
	VolumeMode  *corev1.PersistentVolumeMode        `json:"volumeMode"`
	Phase       corev1.PersistentVolumePhase        `json:"phase"`
	CreateTime  string                              `json:"CreateTime"`
}
type CreatePV struct {
	Zone string                  `json:"Zone" form:"Zone" binding:"required"`
	PV   corev1.PersistentVolume `json:"pv" form:"pv" binding:"required"`
}

type QueryList struct {
	Zone string `json:"zone" form:"zone" binding:"required"`
	Size int    `json:"size,omitempty" form:"size" binding:"required"`
	Page int    `json:"page,omitempty" form:"page" binding:"required"`
}

type DeletePV struct {
	Zone string `json:"zone" form:"zone" binding:"required"`
	Name string `json:"name" form:"name" binding:"required"`
}

type DetailQuery struct {
	Zone string `form:"zone" json:"zone" binding:"required"`
	Name string `form:"name" json:"name" binding:"required"`
}

type UpdatePV struct {
	Zone string                  `form:"zone" json:"zone" binding:"required"`
	PV   corev1.PersistentVolume `form:"pv" json:"pv" binding:"required"`
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
