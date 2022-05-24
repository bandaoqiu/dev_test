package repository

import (
	"dev_test/internal/user/domain/dto"
	"dev_test/internal/user/domain/entity"
)

type UserI interface {
	AddUser(user *entity.User)error
	SingleByEmail(email string)(*entity.User,error)
	UpdateByEmail(email string,update dto.ProfileUpdate)error
}
