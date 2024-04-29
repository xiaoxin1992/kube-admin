package users

import (
	"github.com/gin-gonic/gin"
	models "github.com/xiaoxin1992/kube-admin/models/users"
	"github.com/xiaoxin1992/kube-admin/pkg/logger"
	"github.com/xiaoxin1992/kube-admin/services/users"
	"net/http"
)

func ListUser(ctx *gin.Context) {
	request := models.QueryList{}
	response := &models.Response{
		Code:    http.StatusBadRequest,
		Message: "",
	}
	if err := ctx.ShouldBindQuery(&request); err != nil {
		response.Message = "缺少参数"
		logger.GetLogger().S("http").Named("user").Errorf("获取参数错误 %v", err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	userList := users.NewService().ListUser(ctx, request)
	ctx.JSON(http.StatusOK, userList)
	return
}
