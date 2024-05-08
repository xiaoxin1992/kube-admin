package statefulset

import (
	"github.com/gin-gonic/gin"
	models "github.com/xiaoxin1992/kube-admin/models/statefulset"
	"github.com/xiaoxin1992/kube-admin/pkg/logger"
	"github.com/xiaoxin1992/kube-admin/services/statefulset"
	"net/http"
)

func DetailStateFulSet(ctx *gin.Context) {
	req := models.DetailQuery{}
	response := models.Response{}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.Message = "缺少参数"
		logger.GetLogger().S("http").Named("statefulset").Errorf("获取参数错误 %v", err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response = statefulset.NewServices().DetailStateFulSet(ctx, req)
	ctx.JSON(http.StatusOK, response)
	return
}
