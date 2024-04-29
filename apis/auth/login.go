package auth

import (
	"github.com/gin-gonic/gin"
	models "github.com/xiaoxin1992/kube-admin/models/auth"
	"github.com/xiaoxin1992/kube-admin/pkg/logger"
	"github.com/xiaoxin1992/kube-admin/services/login"
	"net/http"
)

func Login(ctx *gin.Context) {
	userInfo := models.Login{}
	response := models.Response{}
	if err := ctx.ShouldBind(&userInfo); err != nil {
		response.Message = "用户名密码错误"
		logger.GetLogger().S("http").Named("login").Errorf("get user info error: %v", err)
		response.Code = http.StatusUnauthorized
		ctx.JSON(http.StatusUnauthorized, response)
		return
	}
	response = login.NewService().Login(ctx, userInfo.Username, userInfo.Password)
	ctx.JSON(http.StatusOK, response)
}
