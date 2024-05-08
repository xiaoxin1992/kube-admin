package secret

import (
	"github.com/gin-gonic/gin"
	models "github.com/xiaoxin1992/kube-admin/models/secret"
	"github.com/xiaoxin1992/kube-admin/pkg/logger"
	"github.com/xiaoxin1992/kube-admin/services/secret"
	"net/http"
)

func UpdateSecret(ctx *gin.Context) {
	response := models.Response{}
	req := models.UpdateSecret{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Message = "缺少参数"
		logger.GetLogger().S("http").Named("secret").Errorf("get args error %v", err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response = secret.NewServices().UpdateSecret(ctx, req)
	ctx.JSON(http.StatusOK, response)
	return
}
