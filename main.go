package main

import (
	"dev_test/infrastructure/mysqlx"
	"dev_test/internal/user/domain/entity"
	"dev_test/pkg/configx"
	"dev_test/routers"
	"github.com/gin-gonic/gin"
)
func init(){
	db := mysqlx.InitMysql()
	if err := db.AutoMigrate(&entity.User{});err != nil{
		panic(err)
	}
}
func main(){
	r := gin.New()
	//r.SetTrustedProxies([]string{"127.0.0.1"})
	routers.NewGinRouter(r)
	r.Run(configx.Cfg.App.Port)
}
