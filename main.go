package main

import (
	"dev_test/model"
	"dev_test/pkg"
	"dev_test/routers"
	"github.com/gin-gonic/gin"
)
func init(){
	db := pkg.InitMysql()
	db.AutoMigrate(&model.User{})
}
func main(){
	r := gin.New()
	r.SetTrustedProxies([]string{"127.0.0.1"})
	routers.NewGinRouter(r)
	r.Run(":8080")
}
