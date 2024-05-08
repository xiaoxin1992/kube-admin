package pv

import (
	"github.com/gin-gonic/gin"
	models "github.com/xiaoxin1992/kube-admin/models/pv"
	"github.com/xiaoxin1992/kube-admin/pkg/logger"
	"github.com/xiaoxin1992/kube-admin/services/pv"
	"net/http"
)

func ListPV(ctx *gin.Context) {
	req := models.QueryList{}
	response := models.Response{}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.Message = "缺少参数"
		logger.GetLogger().S("http").Named("pv").Errorf("获取参数错误 %v", err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response = pv.NewServices().ListPV(ctx, req)
	ctx.JSON(http.StatusOK, response)
	return
}
