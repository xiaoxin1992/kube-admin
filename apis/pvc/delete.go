package pvc

import (
	"github.com/gin-gonic/gin"
	models "github.com/xiaoxin1992/kube-admin/models/pvc"
	"github.com/xiaoxin1992/kube-admin/pkg/logger"
	"github.com/xiaoxin1992/kube-admin/services/pvc"
	"net/http"
)

func DeletePersistentVolumeClaim(ctx *gin.Context) {
	response := models.Response{}
	req := models.DeletePersistentVolumeClaim{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Message = "缺少参数"
		logger.GetLogger().S("http").Named("pvc").Errorf("get args error %v", err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response = pvc.NewServices().DeletePersistentVolumeClaim(ctx, req)
	ctx.JSON(http.StatusOK, response)
	return
}
