package service

import (
	"github.com/gin-gonic/gin"
	models "github.com/xiaoxin1992/kube-admin/models/service"
	"github.com/xiaoxin1992/kube-admin/pkg/logger"
	"github.com/xiaoxin1992/kube-admin/services/service"
	"net/http"
)

func CreateService(ctx *gin.Context) {
	response := models.Response{}
	req := models.CreateService{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Message = "缺少参数"
		logger.GetLogger().S("http").Named("service").Errorf("获取参数错误 %v", err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response = service.NewServices().CreateService(ctx, req)
	ctx.JSON(http.StatusOK, response)
	return
}
