package entity

import "gorm.io/gorm"

type Class struct {
	gorm.Model
	Name        string  `json:"name" gorm:"type:varchar(20)"`
	Course_id   uint    `json:"course_id" gorm:"type:uint"`
	Participant int     `json:"participant" gorm:"type:int"`
	Users       []*User `json:"users" gorm:"many2many:users_classes"`
}
