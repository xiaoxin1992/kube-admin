package cluster

import (
	"github.com/gin-gonic/gin"
	models "github.com/xiaoxin1992/kube-admin/models/cluster"
	"github.com/xiaoxin1992/kube-admin/pkg/logger"
	"github.com/xiaoxin1992/kube-admin/services/cluster"
	"net/http"
)

func DeleteCluster(ctx *gin.Context) {
	response := models.Response{
		Code:    http.StatusBadRequest,
		Message: "",
	}
	deleteCluster := models.DeleteCluster{}
	if err := ctx.ShouldBind(&deleteCluster); err != nil {
		response.Message = "缺少参数"
		logger.GetLogger().S("http").Named("cluster").Errorf("获取参数错误 %v", err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response = cluster.NewCluster().DeleteCluster(ctx, deleteCluster)
	ctx.JSON(http.StatusOK, response)
	return
}
