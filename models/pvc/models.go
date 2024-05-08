package pvc

import (
	corev1 "k8s.io/api/core/v1"
)

type PersistentVolumeClaim struct {
	Name             string                              `json:"name"`
	Labels           map[string]string                   `json:"labels"`
	StorageClassName string                              `json:"storageClassName"`
	VolumeMode       *corev1.PersistentVolumeMode        `json:"volumeMode"`
	Phase            corev1.PersistentVolumeClaimPhase   `json:"phase"`
	AccessModes      []corev1.PersistentVolumeAccessMode `json:"accessModes"`
	Resources        corev1.VolumeResourceRequirements   `json:"resources"`
	CreateTime       string                              `json:"CreateTime"`
}

type CreatePersistentVolumeClaim struct {
	Zone                  string                       `json:"Zone" form:"Zone" binding:"required"`
	PersistentVolumeClaim corev1.PersistentVolumeClaim `json:"pvc" form:"pvc" binding:"required"`
}

type QueryList struct {
	Zone     string `json:"zone" form:"zone" binding:"required"`
	Namspace string `json:"namespace" form:"namespace" binding:"required"`
	Size     int    `json:"size,omitempty" form:"size" binding:"required"`
	Page     int    `json:"page,omitempty" form:"page" binding:"required"`
}

type DeletePersistentVolumeClaim struct {
	Zone      string `json:"zone" form:"zone" binding:"required"`
	Namespace string `json:"namespace" form:"namespace" binding:"required"`
	Name      string `json:"name" form:"name" binding:"required"`
}

type DetailQuery struct {
	Zone      string `form:"zone" json:"zone" binding:"required"`
	Namespace string `form:"namespace" json:"namespace" binding:"required"`
	Name      string `form:"name" json:"name" binding:"required"`
}

type UpdatePersistentVolumeClaim struct {
	Zone                  string                       `form:"zone" json:"zone" binding:"required"`
	PersistentVolumeClaim corev1.PersistentVolumeClaim `form:"pvc" json:"pvc" binding:"required"`
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
