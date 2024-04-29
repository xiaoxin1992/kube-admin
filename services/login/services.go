package login

import (
	"context"
	"github.com/xiaoxin1992/kube-admin/dao/users"
	models "github.com/xiaoxin1992/kube-admin/models/auth"
	"github.com/xiaoxin1992/kube-admin/pkg/jwt_utils"
	"github.com/xiaoxin1992/kube-admin/pkg/logger"
	"go.uber.org/zap"
	"net/http"
)

func NewService() *service {
	return &service{
		logger: logger.GetLogger().S("service").Named("login"),
	}
}

type service struct {
	logger *zap.SugaredLogger
}

func (s *service) Login(ctx context.Context, username string, password string) models.Response {
	dao := users.NewDao()
	response := models.Response{}
	queryUser, err := dao.VerifyPassword(ctx, username, password)
	if err != nil {
		response.BaseResponse.Code = http.StatusUnauthorized
		response.BaseResponse.Message = "用户名或密码错误"
		return response
	}
	token, err := jwt_utils.GenerateToken(queryUser.Username)
	if err != nil {
		s.logger.Errorf("generate token error: %v", err)
		response.BaseResponse.Code = http.StatusUnauthorized
		response.BaseResponse.Message = "用户名或密码错误"
		return response
	}
	response.BaseResponse.Code = http.StatusOK
	response.Data = map[string]interface{}{
		"token":    token,
		"username": queryUser.Username,
		"role":     queryUser.Role,
	}
	return response
}
