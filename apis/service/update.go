package service

import (
	"github.com/gin-gonic/gin"
	models "github.com/xiaoxin1992/kube-admin/models/service"
	"github.com/xiaoxin1992/kube-admin/pkg/logger"
	"github.com/xiaoxin1992/kube-admin/services/service"
	"net/http"
)

func UpdateService(ctx *gin.Context) {
	response := models.Response{}
	req := models.UpdateService{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Message = "缺少参数"
		logger.GetLogger().S("http").Named("service").Errorf("get args error %v", err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response = service.NewServices().UpdateService(ctx, req)
	ctx.JSON(http.StatusOK, response)
	return
}
