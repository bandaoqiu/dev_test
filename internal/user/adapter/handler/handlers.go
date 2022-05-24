package handler

import (
	"dev_test/internal/user/domain/dto"
	"dev_test/internal/user/domain/entity"
	"dev_test/internal/user/domain/service"
	"dev_test/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)
type User struct {
	srv *service.UserSrv
}
func NewUser()*User{
	return &User{srv: service.NewUserSrv()}
}
//注册
func (t *User)Signup(ctx *gin.Context){
	var user entity.User
	if err := ctx.ShouldBindJSON(&user);err != nil{
		pkg.Response(ctx,http.StatusBadRequest,err.Error(),nil)

		return
	}
	cmd,err :=  t.srv.Add(user)
	if err != nil{
		pkg.Response(ctx,http.StatusBadRequest,err.Error(),nil)
		return
	}
	if err = cmd.AddUser();err!= nil{

		pkg.Response(ctx,http.StatusBadRequest,err.Error(),nil)
		return
	}
	pkg.Response(ctx,http.StatusOK,"success",gin.H{"first_name":user.FirstName,"last_name":user.LastName,"email":user.Email})

	return
}
//登录

func (t *User)Signin(ctx *gin.Context){
	var login dto.Login
	if err := ctx.ShouldBindJSON(&login);err != nil{
		pkg.Response(ctx,http.StatusBadRequest,err.Error(),nil)

		return
	}
	cmd,err:=  t.srv.FindOne(login.Email,login.Pwd)
	if err != nil{
		pkg.Response(ctx,http.StatusBadRequest,err.Error(),nil)

		return
	}
	token,err := cmd.CheckLogin()
	if err != nil{
		pkg.Response(ctx,http.StatusBadRequest,err.Error(),nil)
		return
	}

	pkg.Response(ctx,http.StatusOK,"success",token)

	return
}
//信息

func (t *User)Profile(ctx *gin.Context){
	email := ctx.GetString("email")

	cmd,err := t.srv.FindOne(email,"")
	if err != nil{
		pkg.Response(ctx,http.StatusBadRequest,err.Error(),nil)
		return
	}
	data,err := cmd.Profile()
	if err != nil{
		pkg.Response(ctx,http.StatusBadRequest,err.Error(),nil)
		return
	}
	pkg.Response(ctx,http.StatusOK,"success",data)

	return
}
//更新信息
func (t *User)ProfileUpdate(ctx *gin.Context){
	var update dto.ProfileUpdate
	if err := ctx.ShouldBindJSON(&update);err!=nil{
		pkg.Response(ctx,http.StatusBadRequest,err.Error(),nil)

		return
	}
	email := ctx.GetString("email")
	cmd,err:=  t.srv.UpdateUser(email)
	if err != nil{
		pkg.Response(ctx,http.StatusBadRequest,err.Error(),nil)
		return
	}
	if err =  cmd.Update(update);err!= nil{
		pkg.Response(ctx,http.StatusInternalServerError,err.Error(),nil)
		return
	}
	pkg.Response(ctx,http.StatusOK,"success",gin.H{
		"first_name":update.FirstName,
		"last_name":update.LastName,
		"email":email,
	})

	return
}
