package repository

import (
	"bcc-university/src/business/entity"

	"gorm.io/gorm"
)

type StudentRepository interface {
	GetLastStudentNumber() string
	CreateStudent(student entity.Student) error
}

type studentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &studentRepository{
		db: db,
	}
}

func (studentRepository *studentRepository) GetLastStudentNumber() string {

	var student entity.Student
	
	studentRepository.db.Last(&student)

	return student.Student_id_number

}

func (studentRepository *studentRepository) CreateStudent(student entity.Student) error {

	err := studentRepository.db.Create(&student).Error

	return err

}
