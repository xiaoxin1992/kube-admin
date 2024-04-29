package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaoxin1992/kube-admin/pkg/jwt_utils"
	"go.uber.org/zap"
	"net/http"
)

func JWTAuth(zapLog *zap.SugaredLogger) func(c *gin.Context) {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")
		result := struct {
			Code    int    `json:"code"`
			Message string `json:"Message"`
		}{
			Code:    http.StatusUnauthorized,
			Message: "未登陆，请登登录后重试",
		}
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, result)
			ctx.Abort()
			return
		}
		claims, err := jwt_utils.ParseToken(token)
		if err != nil {
			zapLog.Errorf("authorization error %s", err.Error())
			ctx.JSON(http.StatusUnauthorized, result)
			ctx.Abort()
			return
		}
		ctx.Set("auth", claims)
		ctx.Next()
	}
}
