package routers

import (
	"dev_test/handler"
	"dev_test/middleware"
	"github.com/gin-gonic/gin"
)

func NewGinRouter(r *gin.Engine){
	r.Use(gin.Logger())
	v1 := r.Group("/v1")
	{

		v1.POST("/signup",handler.Signup)
		v1.POST("/signin",handler.Signin)
		v1.Use(middleware.Jwt())
		{
			v1.GET("/profile",handler.Profile)
			v1.POST("/profile/update",handler.ProfileUpdate)
		}
	}
}
