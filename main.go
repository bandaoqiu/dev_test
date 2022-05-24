package main

import (
	"dev_test/infrastructure/mysqlx"
	"dev_test/internal/user/domain/entity"
	"dev_test/pkg/configx"
	"dev_test/routers"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)
func init(){
	db := mysqlx.InitMysql()
	if err := db.AutoMigrate(&entity.User{});err != nil{
		panic(err)
	}
}
func main(){
	r := gin.New()
	gin.SetMode("debug")
	_ = r.SetTrustedProxies([]string{"127.0.0.1"})
	routers.NewGinRouter(r)
	s := &http.Server{
		Addr:              configx.Cfg.App.Port,
		Handler:       r   ,
		ReadTimeout:       15*time.Second,
		WriteTimeout:      15*time.Second,
		MaxHeaderBytes:    1<< 20,
	}
	_ = s.ListenAndServe()
}
