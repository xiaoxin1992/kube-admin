package users

import (
	"context"
	"github.com/xiaoxin1992/kube-admin/dao/users"
	models "github.com/xiaoxin1992/kube-admin/models/users"
	"net/http"
)

func (s *service) ListUser(ctx context.Context, query models.QueryList) models.Response {
	dao := users.NewDao()
	response := models.Response{}
	total, err := dao.ListUserCount(ctx, query)
	if err != nil {
		response.Message = "获取用户列表出错"
		response.Code = http.StatusBadRequest
		s.logger.Errorf("list user count error: %s", err.Error())
		return response
	}
	userList, err := dao.ListUsers(ctx, query)
	if err != nil {
		response.Message = "获取用户列表出错"
		response.Code = http.StatusBadRequest
		s.logger.Errorf("get user list error %v", err)
		return response
	}
	response.Data = map[string]interface{}{
		"total": total,
		"size":  query.Size,
		"page":  query.Page,
		"data":  userList,
	}
	response.Code = http.StatusOK
	return response
}
