package pods

import (
	v1 "k8s.io/api/core/v1"
	"time"
)

type Conditions struct {
}

type Status struct {
	Phase     string    `json:"phase"`
	StartTime time.Time `json:"startTime"`
	Ready     string    `json:"ready"`
	QosClass  string    `json:"qosClass"`
}
type Pod struct {
	NodeName      string            `json:"nodeName"`
	Namespace     string            `json:"namespace"`
	Name          string            `json:"name"`
	Labels        map[string]string `json:"labels"`
	RestartPolicy string            `json:"restartPolicy"`
	Conditions    Conditions        `json:"conditions"`
	Status        Status            `json:"status"`
	Age           string            `json:"age"`
}

type QueryList struct {
	Pod
	Size      int    `json:"size,omitempty" form:"size" binding:"required"`
	Page      int    `json:"page,omitempty" form:"page" binding:"required"`
	Zone      string `json:"zone" form:"zone" binding:"required"`
	Namespace string `json:"namespace" form:"namespace"`
}

type CreatePod struct {
	Zone string `json:"zone" form:"zone" binding:"required"`
	Pod  v1.Pod `json:"pod" form:"pod" binding:"required"`
}

type DeletePod struct {
	Zone      string `json:"zone" form:"zone" binding:"required"`
	Namespace string `json:"namespace" form:"namespace" binding:"required"`
	Name      string `json:"name" form:"name" binding:"required"`
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
