package cronjob

import (
	batchv1 "k8s.io/api/batch/v1"
)

type Cronjob struct {
	Name               string            `json:"name"`
	Namespace          string            `json:"namespace"`
	Labels             map[string]string `json:"labels"`
	Schedule           string            `json:"schedule"`
	LastScheduleTime   string            `json:"lastScheduleTime"`
	LastSuccessfulTime string            `json:"lastSuccessfulTime"`
	CreateTime         string            `json:"createTime"`
}
type CreateCronjob struct {
	Zone    string          `json:"Zone" form:"Zone" binding:"required"`
	Cronjob batchv1.CronJob `json:"cronjob" form:"cronjob" binding:"required"`
}

type QueryList struct {
	Zone      string `json:"zone" form:"zone" binding:"required"`
	Namespace string `json:"namespace" form:"namespace" binding:"required"`
	Size      int    `json:"size,omitempty" form:"size" binding:"required"`
	Page      int    `json:"page,omitempty" form:"page" binding:"required"`
}

type DeleteCronjob struct {
	Zone      string `json:"zone" form:"zone" binding:"required"`
	Namespace string `json:"namespace" form:"namespace" binding:"required"`
	Name      string `json:"name" form:"name" binding:"required"`
}

type DetailQuery struct {
	Zone      string `form:"zone" json:"zone" binding:"required"`
	Namespace string `form:"namespace" json:"namespace" binding:"required"`
	Name      string `form:"name" json:"name" binding:"required"`
}

type UpdateCronjob struct {
	Zone    string          `form:"zone" json:"zone" binding:"required"`
	Cronjob batchv1.CronJob `form:"cronjob" json:"cronjob" binding:"required"`
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
