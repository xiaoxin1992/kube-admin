package cluster

import "github.com/xiaoxin1992/kube-admin/models"

type Cluster struct {
	Zone   string `json:"zone" form:"id" binding:"required"`
	Host   string `json:"host" form:"id" binding:"required"`
	Token  string `json:"token" form:"id" binding:"required"`
	Remark string `json:"remark" form:"id"`
}

type CreateCluster struct {
	Cluster
	Version string `json:"version"`
}

type QueryCluster struct {
	Id int64 `json:"id"`
	Cluster
	Version    string        `json:"version"`
	CreateTime models.CNTime `json:"create_time"`
	UpdateTime models.CNTime `json:"update_time"`
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type DeleteCluster struct {
	Zone string `json:"zone" form:"zone" binding:"required"`
}

type QueryList struct {
	Size  int    `json:"size" form:"size" binding:"required"`
	Page  int    `json:"page" form:"page" binding:"required"`
	Query string `json:"query" form:"query"`
}

type ResponseClusterList struct {
	Size int         `json:"size"`
	Page int         `json:"page"`
	Data interface{} `json:"data"`
}

type UpdateCluster struct {
	Zone   string `json:"zone" form:"id" binding:"required"`
	Host   string `json:"host" form:"id" binding:"required"`
	Token  string `json:"token" form:"id" binding:"required"`
	Remark string `json:"remark" form:"id" binding:"required"`
}
