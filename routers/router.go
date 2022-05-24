package routers

import (
	"dev_test/infrastructure/middleware"
	"dev_test/internal/user/adapter/handler"
	. "dev_test/pkg/logx"
	"github.com/gin-gonic/gin"
	"time"
)

func NewGinRouter(r *gin.Engine){
	//

	r.Use(middleware.Ginzap(Logger,time.RFC3339,true))
	r.Use(middleware.RecoveryWithZap(Logger,true))

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
