package entity

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	User_id           uint   `json:"user_id" gorm:"type:uint"`
	Student_id_number string `json:"student_id_number" gorm:"type:varchar(20)"`
}

type ClaimStudentNumberApi struct {
	Student_id_number string `json:"student_id_number"`
}
