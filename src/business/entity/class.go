package entity

import "gorm.io/gorm"

type Class struct {
	gorm.Model
	Name        string  `json:"name" gorm:"type:varchar(50);unique"`
	Course_id   uint    `json:"course_id" gorm:"type:uint"`
	Participant int     `json:"participant" gorm:"type:int;default:0"`
	ClassCode   string  `json:"class_code" gorm:"type:varchar(20);unique"`
	Course      Course  `json:"course"`
	Users       []*User `json:"users" gorm:"many2many:user_classes"`
}

type ClassApi struct {
	Name        string `json:"name"`
	Course_id   uint   `json:"course_id"`
	Participant int    `json:"participant"`
	ClassCode   string `json:"class_code"`
}

type CreateUpdateClassBind struct {
	Name      string `json:"name" binding:"required"`
	Course_id uint   `json:"course_id" binding:"required,numeric"`
}

type CreateUpdateClassApi struct {
	Name        string `json:"name"`
	Course_id   uint   `json:"course_id"`
	Participant int    `json:"participant"`
	ClassCode   string `json:"class_code"`
	Course      struct {
		Name   string `json:"name"`
		Credit int    `json:"credit"`
	} `json:"Course"`
}
