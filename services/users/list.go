package users

import (
	"context"
	"github.com/xiaoxin1992/kube-admin/dao/users"
	models "github.com/xiaoxin1992/kube-admin/models/users"
)

func (s *service) ListUser(ctx context.Context, query models.QueryList) models.ResponseUserList {
	dao := users.NewDao()
	response := models.ResponseUserList{
		Size: 0,
		Page: 0,
		Data: make(map[string]interface{}, 0),
	}
	total, err := dao.ListUserCount(ctx, query)
	if err != nil {
		s.logger.Errorf("list user count error: %s", err.Error())
		return response
	}
	userList, err := dao.ListUsers(ctx, query)
	if err != nil {
		s.logger.Errorf("get user list error %v", err)
		return response
	}
	response.Size = query.Size
	response.Page = query.Page
	response.Data = map[string]interface{}{
		"total": total,
		"users": userList,
	}
	return response
}
