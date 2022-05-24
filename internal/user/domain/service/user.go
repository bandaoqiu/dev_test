package service

import (
	"dev_test/infrastructure/mysqlx"
	"dev_test/internal/user/domain/aggregate"
	"dev_test/internal/user/domain/entity"
	"dev_test/internal/user/domain/persistence/mysqlp"
	"dev_test/internal/user/domain/repository"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UserSrv struct {
	userI repository.UserI
}
func NewUserSrv()*UserSrv{
	userDb := mysqlp.User{Db: mysqlx.InitMysql()}
	return &UserSrv{
		userI: &userDb,
	}
}
func (t *UserSrv)Add(user entity.User)(*aggregate.User,error){
	u,err :=  t.userI.SingleByEmail(user.Email)
	if err != nil{
		return nil, err
	}
	if u != nil{
		return nil,errors.New("用户已存在")
	}
	pwd,err := bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
	if err != nil{
		return nil, err
	}
	user.Password = string(pwd)
	return aggregate.NewUser("","",&user,t.userI),nil
}
func (t *UserSrv)FindOne(email string,password string)(*aggregate.User,error){
	if email == ""{
		return nil,errors.New("邮箱不能为空")
	}
	return aggregate.NewUser(email,password,nil,t.userI),nil
}
func (t *UserSrv)UpdateUser(email string)(*aggregate.User,error){
	if email == ""{
		return nil,errors.New("邮箱不能为空")
	}
	return aggregate.NewUser(email,"",nil,t.userI),nil
}

