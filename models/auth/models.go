package auth

import (
	"github.com/xiaoxin1992/kube-admin/models"
)

type Login struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type Response struct {
	models.BaseResponse
	Data interface{} `json:"data,omitempty"`
}
