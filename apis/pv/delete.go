package pv

import (
	"github.com/gin-gonic/gin"
	models "github.com/xiaoxin1992/kube-admin/models/pv"
	"github.com/xiaoxin1992/kube-admin/pkg/logger"
	"github.com/xiaoxin1992/kube-admin/services/pv"
	"net/http"
)

func DeletePV(ctx *gin.Context) {
	response := models.Response{}
	req := models.DeletePV{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Message = "缺少参数"
		logger.GetLogger().S("http").Named("pv").Errorf("get args error %v", err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response = pv.NewServices().DeletePV(ctx, req)
	ctx.JSON(http.StatusOK, response)
	return
}
