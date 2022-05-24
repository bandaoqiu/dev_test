package pkg

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(ctx *gin.Context,code int,msg string,data interface{}){
	ctx.AbortWithStatusJSON(http.StatusOK,gin.H{"code":code,"msg":msg,"data":data})
	return
}
