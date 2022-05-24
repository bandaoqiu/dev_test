package mysqlp

import (
	"dev_test/internal/user/domain/dto"
	"dev_test/internal/user/domain/entity"
	"dev_test/internal/user/domain/repository"
	"errors"
	"gorm.io/gorm"
)

type User struct {
	Db *gorm.DB
}
var _ repository.UserI = new(User)
func (t *User)AddUser(user *entity.User)error{
	tx :=  t.Db.Create(&user)
	return tx.Error
}
func (t *User)SingleByEmail(email string)(user *entity.User,err error){
	tx := t.Db.Where("email = ?",email).First(&user)
	if tx.Error != nil{
		if errors.Is(tx.Error,gorm.ErrRecordNotFound){
			return nil,nil
		}
		return nil,tx.Error
	}
	return
}
func (t *User)UpdateByEmail(email string,update dto.ProfileUpdate)error{
	tx := t.Db.Model(&entity.User{})
	tx = tx.Where("email = ?",email)
	tx = tx.Updates(update)
	return tx.Error
}
