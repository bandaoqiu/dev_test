package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `json:"first_name" binding:"required,min=2" gorm:"column:first_name;type:varchar(255);not null;"`
	LastName  string `json:"last_name" binding:"required,min=2" gorm:"column:last_name;type:varchar(255);not null"`
	Email     string `json:"email" binding:"required,email" gorm:"column:email;type:varchar(200);not null;uniqueIndex"`
	Password  string `json:"password" binding:"required,min=6,max=16,alphanum" gorm:"column:password;type:varchar(100);not null;"`
}
