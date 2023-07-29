package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string  `json:"name" gorm:"type:varchar(50)"`
	Username string  `json:"username" gorm:"type:varchar(20)"`
	Email    string  `json:"email" gorm:"type:varchar(50);unique"`
	Password string  `json:"password" gorm:"type:varchar(100)"`
	Role     string  `json:"role" gorm:"type:enum('admin', 'user');default:user"`
	Student  Student `json:"student"`
	Classes  []Class `json:"classes" gorm:"many2many:user_classes"`
}

type CreateUser struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserRegistResponse struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
