package deployment

import (
	"github.com/gin-gonic/gin"
	models "github.com/xiaoxin1992/kube-admin/models/deployment"
	"github.com/xiaoxin1992/kube-admin/pkg/logger"
	"github.com/xiaoxin1992/kube-admin/services/deployment"
	"net/http"
)

func DeleteDeployment(ctx *gin.Context) {
	response := models.Response{}
	req := models.DeleteDeployment{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Message = "缺少参数"
		logger.GetLogger().S("http").Named("deployment").Errorf("获取参数错误 %v", err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response = deployment.NewServices().DeleteDeployment(ctx, req)
	ctx.JSON(http.StatusOK, response)
	return
}
