package namespace

import (
	"github.com/gin-gonic/gin"
	models "github.com/xiaoxin1992/kube-admin/models/namespace"
	"github.com/xiaoxin1992/kube-admin/pkg/logger"
	"github.com/xiaoxin1992/kube-admin/services/namespace"
	"net/http"
)

func DeleteNamespace(ctx *gin.Context) {
	response := models.Response{}
	req := models.DeleteNamespace{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Message = "缺少参数"
		logger.GetLogger().S("http").Named("namespace").Errorf("获取参数错误 %v", err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response = namespace.NewServices().DeleteNamespace(ctx, req)
	ctx.JSON(http.StatusOK, response)
	return
}
