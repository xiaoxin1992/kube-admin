package pods

import (
	"github.com/gin-gonic/gin"
	models "github.com/xiaoxin1992/kube-admin/models/pods"
	"github.com/xiaoxin1992/kube-admin/pkg/logger"
	"github.com/xiaoxin1992/kube-admin/services/pods"
	"net/http"
)

func DeletePod(ctx *gin.Context) {
	response := models.Response{}
	req := models.DeletePod{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Message = "缺少参数"
		logger.GetLogger().S("http").Named("user").Errorf("获取参数错误 %v", err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response = pods.NewServices().DeletePod(ctx, req)
	ctx.JSON(http.StatusOK, response)
	return
}
