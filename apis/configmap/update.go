package configmap

import (
	"github.com/gin-gonic/gin"
	models "github.com/xiaoxin1992/kube-admin/models/configmap"
	"github.com/xiaoxin1992/kube-admin/pkg/logger"
	"github.com/xiaoxin1992/kube-admin/services/configmap"
	"net/http"
)

func UpdateConfigmap(ctx *gin.Context) {
	response := models.Response{}
	req := models.UpdateConfigmap{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Message = "缺少参数"
		logger.GetLogger().S("http").Named("configmap").Errorf("获取参数错误 %v", err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response = configmap.NewServices().UpdateConfigmap(ctx, req)
	ctx.JSON(http.StatusOK, response)
	return
}
