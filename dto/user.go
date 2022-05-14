package dto

import "gorm.io/gorm"

type Login struct {
	Email string `json:"email" binding:"required,email"`
	Pwd string `json:"password" binding:"required,min=6,max=16,alphanum"`
}
type ProfileMsg struct {
	FirstName string `json:"first_name" binding:"required,min=2"`
	LastName string `json:"last_name" binding:"required,min=2"`
	Email string `json:"email" binding:"required,email"`
}
type ProfileUpdate struct {
	gorm.Model
	FirstName string `json:"first_name" binding:"required,min=2"`
	LastName string `json:"last_name" binding:"required,min=2"`
}

