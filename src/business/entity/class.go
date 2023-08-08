package entity

import "gorm.io/gorm"

type Class struct {
	gorm.Model
	Name        string  `json:"name" gorm:"type:varchar(50)"`
	Course_id   uint    `json:"course_id" gorm:"type:uint"`
	Participant int     `json:"participant" gorm:"type:int;default:0"`
	ClassCode   string  `json:"class_code" gorm:"type:varchar(20)"`
	Course      Course  `json:"course"`
	Users       []*User `json:"users" gorm:"many2many:user_classes"`
}

type ClassApi struct {
	Name        string `json:"name"`
	Course_id   uint   `json:"course_id"`
	Participant int    `json:"participant"`
	ClassCode   string `json:"class_code"`
}
