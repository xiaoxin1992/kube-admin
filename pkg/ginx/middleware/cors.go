package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return corsHandle
}

func corsHandle(ctx *gin.Context) {
	method := ctx.Request.Method
	// 必须，接受指定域的请求，可以使用*不加以限制，但不安全
	ctx.Header("Access-Control-Allow-Origin", "*")
	//ctx.Header("Access-Control-Allow-Origin", context.GetHeader("Origin"))
	// 必须，设置服务器支持的所有跨域请求的方法
	ctx.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
	// 服务器支持的所有头信息字段，不限于浏览器在"预检"中请求的字段
	//ctx.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Token")
	ctx.Header("Access-Control-Allow-Headers", "*")
	// 可选，设置XMLHttpRequest的响应对象能拿到的额外字段
	//ctx.Header("Access-Control-Expose-Headers", "Access-Control-Allow-Headers, Token")
	ctx.Header("Access-Control-Expose-Headers", "*")
	// 可选，是否允许后续请求携带认证信息Cookir，该值只能是true，不需要则不设置
	ctx.Header("Access-Control-Allow-Credentials", "true")
	// 放行所有OPTIONS方法
	if method == "OPTIONS" {
		ctx.AbortWithStatus(http.StatusNoContent)
		return
	}
	ctx.Next()
}
