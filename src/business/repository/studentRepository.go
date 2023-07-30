package repository

import (
	"bcc-university/src/business/entity"

	"gorm.io/gorm"
)

type StudentRepository interface {
	GetLastStudentNumber() (string, error)
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

func (studentRepository *studentRepository) GetLastStudentNumber() (string, error) {

	var student entity.Student
	err := studentRepository.db.Last(&student)

	return student.Student_id_number, err.Error

}

func (studentRepository *studentRepository) CreateStudent(student entity.Student) error {

	err := studentRepository.db.Create(&student)

	return err.Error

}
