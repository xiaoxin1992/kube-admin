package users

import (
	"github.com/gin-gonic/gin"
	models "github.com/xiaoxin1992/kube-admin/models/users"
	"github.com/xiaoxin1992/kube-admin/pkg/logger"
	"github.com/xiaoxin1992/kube-admin/services/users"
	"net/http"
)

func DeleteUser(ctx *gin.Context) {
	response := models.Response{
		Code:    http.StatusBadRequest,
		Message: "",
	}
	userinfo := models.DeleteUser{}
	if err := ctx.ShouldBind(&userinfo); err != nil {
		response.Message = "缺少参数"
		logger.GetLogger().S("http").Named("user").Errorf("获取参数错误 %v", err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response = users.NewService().DeleteUser(ctx, userinfo)
	ctx.JSON(http.StatusOK, response)
	return
}
