package users

import (
	"github.com/xiaoxin1992/kube-admin/models"
)

const (
	ADMIN = 1 // 管理员
	USER  = 2 // 普通用户
)

type User struct {
	Username    string `json:"username" form:"username" binding:"required"`
	DisplayName string `json:"display_name" form:"display_name" binding:"required"`
	Password    string `json:"password,omitempty" form:"password,omitempty" binding:"required"`
	// 角色: 1 表示管理员, 2表示普通用户, 默认是2
	Role  int    `json:"role" form:"role" binding:"required"`
	Email string `json:"email" form:"email"`
	Phone string `json:"phone" form:"phone"`
}

type QueryUser struct {
	Id int64 `json:"id"`
	User
	CreateTime models.CNTime `json:"create_time"`
	UpdateTime models.CNTime `json:"update_time"`
}

type QueryList struct {
	Size  int    `json:"size" form:"size" binding:"required"`
	Page  int    `json:"page" form:"page" binding:"required"`
	Query string `json:"query" form:"query"`
}

type UpdateUser struct {
	Username    string `json:"username" form:"username" binding:"required"`
	DisplayName string `json:"display_name" form:"display_name" binding:"required"`
	Email       string `json:"email" form:"email"`
	Phone       string `json:"phone" form:"phone"`
}

type ResetPassword struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type DeleteUser struct {
	Username string `json:"username"`
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type ResponseUserList struct {
	Size int         `json:"size"`
	Page int         `json:"page"`
	Data interface{} `json:"data"`
}
