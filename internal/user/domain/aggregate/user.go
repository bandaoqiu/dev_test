package aggregate

import (
	"dev_test/infrastructure/redisx"
	"dev_test/internal/user/domain/dto"
	"dev_test/internal/user/domain/entity"
	"dev_test/internal/user/domain/repository"
	"dev_test/pkg/jwtx"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	email string
	password string
	user *entity.User
	userDb repository.UserI
}
func NewUser(email ,password string,user *entity.User,i repository.UserI)*User{
	return &User{
		email:  email,
		password: password,
		user:   user,
		userDb: i,
	}
}
func (t *User)AddUser()error{
	return t.userDb.AddUser(t.user)
}
func (t *User)CheckLogin()(string,error){
	user,err := t.userDb.SingleByEmail(t.email)
	if err != nil{
		return "",err
	}
	if user == nil{
		return "",errors.New("用户不存在")
	}
	if err =  bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(t.password));err!=nil{
		return "", errors.New("密码错误")
	}
	token,err := jwtx.MakeToken(t.email)
	if err != nil{
		return "",err
	}
	client := redisx.InitRedis()
	defer client.Close()
	client.Set(t.email,token,time.Hour*2)
	return token,nil
}
func (t *User)Profile()(*dto.ProfileMsg,error){
	user,err := t.userDb.SingleByEmail(t.email)
	if err != nil{
		return nil,err
	}
	return &dto.ProfileMsg{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	},nil
}
func (t *User)Update(update dto.ProfileUpdate)error{
	return t.userDb.UpdateByEmail(t.email,update)
}
