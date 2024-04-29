package users

import (
	"context"
	"fmt"
	"github.com/xiaoxin1992/kube-admin/dao/users"
	models "github.com/xiaoxin1992/kube-admin/models/users"
	"net/http"
)

func (s *service) CreateUser(ctx context.Context, user *models.User) *models.Response {
	response := models.Response{
		Code:    http.StatusOK,
		Message: "",
	}
	dao := users.NewDao()
	// 判断用户是否存在了
	ok, err := dao.ExistsUsername(ctx, user.Username)
	if err != nil {
		s.logger.Errorf("query user info error: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = "查询用户出错!"
		return &response
	}
	if ok {
		response.Code = http.StatusConflict
		response.Message = fmt.Sprintf("用户 %s 已经存在!", user.Username)
		return &response
	}
	// 创建用户
	err = dao.CreateUser(ctx, user)
	if err != nil {
		s.logger.Errorf("create user error: %v", err)
		response.Code = http.StatusBadRequest
		response.Message = "创建用户失败!"
	} else {
		response.Code = http.StatusOK
		response.Message = fmt.Sprintf("用户 %s 创建完成!", user.Username)
	}
	return &response
}
