package entity

import "gorm.io/gorm"

type Class struct {
	gorm.Model
	Name        string  `json:"name" gorm:"type:varchar(50)"`
	Course_id   uint    `json:"course_id" gorm:"type:uint"`
	Participant int     `json:"participant" gorm:"type:int;default:0"`
	Users       []*User `json:"users" gorm:"many2many:user_classes"`
}

type ClassResponse struct {
	Name        string `json:"name"`
	Course_id   uint   `json:"course_id"`
	Participant int    `json:"participant"`
}
