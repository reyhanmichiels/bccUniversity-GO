package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string  `json:"name" gorm:"type:varchar(20)"`
	Username string  `json:"username" gorm:"type:varchar(20)"`
	Email    string  `json:"email" gorm:"type:varchar(20)"`
	Password string  `json:"password" gorm:"type:varchar(20)"`
	Role     string  `json:"role" gorm:"type:enum('admin', 'user')"`
	Student  Student `json:"student"`
	Classes  Class   `json:"classes" gorm:"many2many:users_classes"`
}
