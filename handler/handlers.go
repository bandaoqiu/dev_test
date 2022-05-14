package handler

import (
	"dev_test/dto"
	"dev_test/model"
	"dev_test/pkg"
	"dev_test/pkg/jwtx"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"time"
)
var db *gorm.DB
func init(){
	db = pkg.InitMysql()
}
//注册
func Signup(ctx *gin.Context){
	var user model.User
	if err := ctx.ShouldBindJSON(&user);err != nil{
		ctx.JSON(http.StatusOK,gin.H{"code":400,"msg":err.Error()})
		return
	}
	pwd,err := bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
	if err != nil{
		ctx.JSON(http.StatusOK,gin.H{"code":500,"msg":err.Error()})
		return
	}
	user.Password = string(pwd)

	tx := db.Create(&user)
	if err := tx.Error;err != nil{
		ctx.JSON(http.StatusOK,gin.H{"code":500,"msg":err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,
		gin.H{"code":200,"msg":"success",
			"data":gin.H{"first_name":user.FirstName,"last_name":user.LastName,"email":user.Email}})
	return
}
//登录

func Signin(ctx *gin.Context){
	var login dto.Login
	if err := ctx.ShouldBindJSON(&login);err != nil{
		ctx.JSON(200,gin.H{"code":400,"msg":err.Error()})
		return
	}
	var user model.User
	tx := db.Where("email=?",login.Email).First(&user)
	if err := tx.Error;err!= nil{
		ctx.JSON(200,gin.H{"code":500,"msg":err.Error()})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(login.Pwd));err!= nil{
		ctx.JSON(200,gin.H{"code":500,"msg":"密码错误"})
		return
	}
	token,err := jwtx.MakeToken(user.Email)
	if err != nil{
		ctx.JSON(200,gin.H{"code":400,"msg":err.Error()})
		return
	}

	client := redis.NewClient(&redis.Options{
		Network:            "",
		Addr:               pkg.Cfg.Redis.Addr,
		Password:           pkg.Cfg.Redis.PassWord,
		DB:                 0,

	})
	defer client.Close()
	client.Set(user.Email,token,time.Hour*2)

	ctx.JSON(200,gin.H{"code":200,"msg":"success","data":token})
	return
}
//信息

func Profile(ctx *gin.Context){
	email := ctx.GetString("email")

	var user model.User
	var profile dto.ProfileMsg
	tx := db.Model(&user).Where("email=?",email).First(&profile)
	if err := tx.Error;err!= nil{
		ctx.JSON(200,gin.H{"code":500,"msg":err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,
		gin.H{
		"code":http.StatusOK,
		"msg":"success",
		"data":profile})
	return
}
//更新信息
func ProfileUpdate(ctx *gin.Context){
	var update dto.ProfileUpdate
	if err := ctx.ShouldBindJSON(&update);err!=nil{
		ctx.JSON(200,gin.H{"code":400,"msg":err.Error()})
		return
	}
	email := ctx.GetString("email")
	var user model.User
	tx := db.Model(&user).Where("email=?",email).Updates(update)
	if tx.Error != nil{
		ctx.JSON(200,gin.H{"code":500,"msg":tx.Error})
		return
	}
	ctx.JSON(200,gin.H{
		"code":200,
		"msg":"success",
		"data":gin.H{
			"first_name":update.FirstName,
		"last_name":update.LastName,
		"email":email}})
	return
}
