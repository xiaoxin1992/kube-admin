package cronjob

import (
	"github.com/gin-gonic/gin"
	models "github.com/xiaoxin1992/kube-admin/models/cronjob"
	"github.com/xiaoxin1992/kube-admin/pkg/logger"
	"github.com/xiaoxin1992/kube-admin/services/cronjob"
	"net/http"
)

func DetailCronjob(ctx *gin.Context) {
	req := models.DetailQuery{}
	response := models.Response{}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.Message = "缺少参数"
		logger.GetLogger().S("http").Named("cronjob").Errorf("获取参数错误 %v", err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response = cronjob.NewServices().DetailCronjob(ctx, req)
	ctx.JSON(http.StatusOK, response)
	return
}
