package users

import (
	"context"
	"fmt"
	"github.com/xiaoxin1992/kube-admin/dao/users"
	models "github.com/xiaoxin1992/kube-admin/models/users"
	"net/http"
)

func (s *service) UpdateUser(ctx context.Context, updateUser models.UpdateUser) models.Response {
	db := users.NewDao()
	ok, err := db.ExistsUsername(ctx, updateUser.Username)
	if err != nil {
		s.logger.Errorf("query user info err %v", err)
		return models.Response{
			Code:    http.StatusBadRequest,
			Message: "查询用户信息出错!",
		}
	}
	if !ok {
		return models.Response{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("用户 %s 不存在", updateUser.Username),
		}
	}
	err = db.UpdateUser(ctx, updateUser)
	if err != nil {
		s.logger.Errorf("update user err: %v", err)
		return models.Response{
			Code:    http.StatusBadRequest,
			Message: "更新用户信息失败!",
		}
	}
	return models.Response{
		Code:    http.StatusOK,
		Message: "更新用户信息完成!",
	}
}

func (s *service) ResetPassword(ctx context.Context, resetPassword models.ResetPassword) models.Response {
	db := users.NewDao()
	ok, err := db.ExistsUsername(ctx, resetPassword.Username)
	if err != nil {
		s.logger.Errorf("query user info err %v", err)
		return models.Response{
			Code:    http.StatusBadRequest,
			Message: "查询用户信息出错!",
		}
	}
	if !ok {
		return models.Response{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("用户 %s 不存在!", resetPassword.Username),
		}
	}
	err = db.ResetPassword(ctx, resetPassword)
	if err != nil {
		s.logger.Errorf("update user err: %v", err)
		return models.Response{
			Code:    http.StatusBadRequest,
			Message: "密码重置出错!",
		}
	}
	return models.Response{
		Code:    http.StatusOK,
		Message: "密码重置完成!",
	}
}
