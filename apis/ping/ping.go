package ping

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func List(ctx *gin.Context) {
	// ping 接口测试
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "time": nowTime, "msg": "pong"})
}
