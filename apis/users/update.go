package users

import (
	"github.com/gin-gonic/gin"
	models "github.com/xiaoxin1992/kube-admin/models/users"
	"github.com/xiaoxin1992/kube-admin/pkg/logger"
	"github.com/xiaoxin1992/kube-admin/services/users"
	"net/http"
)

func UpdateUser(ctx *gin.Context) {
	response := models.Response{
		Code:    http.StatusBadRequest,
		Message: "",
	}
	update := models.UpdateUser{}
	if err := ctx.ShouldBind(&update); err != nil {
		response.Message = "缺少参数"
		logger.GetLogger().S("http").Named("user").Errorf("获取参数错误 %v", err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response = users.NewService().UpdateUser(ctx, update)
	ctx.JSON(http.StatusOK, response)
	return
}

func ResetPassword(ctx *gin.Context) {
	response := models.Response{
		Code:    http.StatusBadRequest,
		Message: "",
	}
	update := models.ResetPassword{}
	if err := ctx.ShouldBind(&update); err != nil {
		response.Message = "缺少参数"
		logger.GetLogger().S("http").Named("user").Errorf("获取参数错误 %v", err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response = users.NewService().ResetPassword(ctx, update)
	ctx.JSON(http.StatusOK, response)
	return

}
