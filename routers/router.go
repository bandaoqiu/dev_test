package routers

import (
	"dev_test/infrastructure/middleware"
	"dev_test/internal/user/adapter/handler"
	"github.com/gin-gonic/gin"

)

func NewGinRouter(r *gin.Engine){
	//
	middleware.InitLogger()
	r.Use(middleware.GinLogger())
	r.Use(middleware.GinRecovery(true))

	user := handler.NewUser()
	v1 := r.Group("/v1")
	{

		v1.POST("/signup", user.Signup)
		v1.POST("/signin", user.Signin)
		v1.Use(middleware.Jwt())
		{
			v1.GET("/profile", user.Profile)
			v1.PUT("/profile", user.ProfileUpdate)
		}
	}
}
