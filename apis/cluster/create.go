package cluster

import (
	"github.com/gin-gonic/gin"
	models "github.com/xiaoxin1992/kube-admin/models/cluster"
	"github.com/xiaoxin1992/kube-admin/pkg/logger"
	"github.com/xiaoxin1992/kube-admin/services/cluster"
	"net/http"
)

func CrateCluster(ctx *gin.Context) {
	data := models.CreateCluster{}
	response := models.Response{Code: http.StatusBadRequest}
	if err := ctx.ShouldBind(&data); err != nil {
		response.Message = "缺少参数"
		logger.GetLogger().S("http").Named("user").Errorf("获取参数错误 %v", err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response = cluster.NewCluster().CreateCluster(ctx, data)
	ctx.JSON(http.StatusOK, response)
	return
}
