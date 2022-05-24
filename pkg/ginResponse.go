package pkg

import (

	"github.com/gin-gonic/gin"
)

func Response(ctx *gin.Context,code int,msg string,data interface{}){
	ctx.AbortWithStatusJSON(code,gin.H{"msg":msg,"data":data})
	return
}
