package secret

import (
	"github.com/gin-gonic/gin"
	models "github.com/xiaoxin1992/kube-admin/models/secret"
	"github.com/xiaoxin1992/kube-admin/pkg/logger"
	"github.com/xiaoxin1992/kube-admin/services/secret"
	"net/http"
)

func DetailSecret(ctx *gin.Context) {
	req := models.DetailQuery{}
	response := models.Response{}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.Message = "缺少参数"
		logger.GetLogger().S("http").Named("secret").Errorf("get args error %v", err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response = secret.NewServices().DetailSecret(ctx, req)
	ctx.JSON(http.StatusOK, response)
	return
}
