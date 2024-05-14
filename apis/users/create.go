package users

import (
	"github.com/gin-gonic/gin"
	models "github.com/xiaoxin1992/kube-admin/models/users"
	"github.com/xiaoxin1992/kube-admin/pkg/logger"
	"github.com/xiaoxin1992/kube-admin/services/users"
	"net/http"
)

/*
创建一个用户
*/

func CreateUser(ctx *gin.Context) {
	user := models.User{}
	response := &models.Response{
		Code:    http.StatusBadRequest,
		Message: "",
		Data:    make(map[string]interface{}),
	}
	if err := ctx.ShouldBind(&user); err != nil {
		response.Message = "缺少参数"
		logger.GetLogger().S("http").Named("user").Errorf("获取参数错误 %v", err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	svc := users.NewService()
	response = svc.CreateUser(ctx, &user)
	ctx.JSON(http.StatusOK, response)
}
