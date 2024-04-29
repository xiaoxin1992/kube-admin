package users

import (
	"context"
	"fmt"
	"github.com/xiaoxin1992/kube-admin/dao/users"
	models "github.com/xiaoxin1992/kube-admin/models/users"
	"net/http"
)

func (s *service) DeleteUser(ctx context.Context, userinfo models.DeleteUser) models.Response {
	dao := users.NewDao()
	response := models.Response{
		Code:    http.StatusBadRequest,
		Message: "",
	}
	ok, err := dao.ExistsUsername(ctx, userinfo.Username)
	if err != nil {
		s.logger.Errorf("get user info error %v", err)
		response.Message = fmt.Sprintf("获取用户 %s 信息出错!", userinfo.Username)
		return response
	}
	if !ok {
		response.Message = fmt.Sprintf("用户 %s 不存在!", userinfo.Username)
		return response
	}
	err = dao.DeleteUser(ctx, &userinfo)
	if err != nil {
		s.logger.Errorf("delete user error %v", err)
		response.Message = fmt.Sprintf("用户 %s 删除失败", userinfo.Username)
		return response
	}
	response.Code = http.StatusOK
	response.Message = fmt.Sprintf("用户 %s 删除完成!", userinfo.Username)
	return response
}
