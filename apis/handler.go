package apis

import (
	"dev_test/engine"
	"dev_test/entity"
)

func Signup(ctx *engine.Context) {

	firstName := ctx.PostForm("firstName")
	lastName := ctx.PostForm("lastName")
	email := ctx.PostForm("email")
	pwd := ctx.PostForm("password")

	user := entity.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  pwd,
	}
	if err := user.Validate();err!= nil{
		ctx.Json(engine.H{"msg":err.Error(),"code":engine.NoValidate})
		return
	}
	ctx.Json(engine.H{"msg":"success","code":engine.Success})
	return
}
func Signin(ctx *engine.Context) {

}
func Profile(ctx *engine.Context) {
	ctx.Write(engine.Success, []byte("bbbb"))
}
func ProfileUpdate(ctx *engine.Context) {

}
