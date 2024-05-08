package pv

import (
	"github.com/gin-gonic/gin"
	models "github.com/xiaoxin1992/kube-admin/models/pv"
	"github.com/xiaoxin1992/kube-admin/pkg/logger"
	"github.com/xiaoxin1992/kube-admin/services/pv"
	"net/http"
)

func UpdatePV(ctx *gin.Context) {
	response := models.Response{}
	req := models.UpdatePV{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Message = "缺少参数"
		logger.GetLogger().S("http").Named("pv").Errorf("get args error %v", err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response = pv.NewServices().UpdatePV(ctx, req)
	ctx.JSON(http.StatusOK, response)
	return
}
