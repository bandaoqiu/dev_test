package middleware

import (
	"dev_test/pkg/jwtx"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Jwt()gin.HandlerFunc{
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization") // 获取请求头中的数据
		if authHeader == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"code": http.StatusBadRequest,
				"msg":  "请求头中auth为空",
			})
			ctx.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			ctx.JSON(http.StatusOK, gin.H{
				"code": http.StatusBadRequest,
				"msg":  "请求头中auth格式有误",
			})
			ctx.Abort()
			return
		}
		mc, err := jwtx.ParseToken(parts[1])
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": http.StatusBadRequest,
				"msg":  "无效的Token",
			})
			ctx.Abort()
			return
		}
		ctx.Set("email", mc.Email)
		ctx.Next()
	}
}
