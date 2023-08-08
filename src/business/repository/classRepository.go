package repository

import (
	"bcc-university/src/business/entity"

	"gorm.io/gorm"
)

type ClassRepository interface {
	FindAllClass(allClass *[]entity.ClassResponse) error
	ELFindClassByClassCode(inputClass *entity.Class, classCode string) error
	FindClassById(class *entity.Class, classId uint) error
}

type classRepository struct {
	db *gorm.DB
}

func NewClassRepository(db *gorm.DB) ClassRepository {

	return &classRepository{
		db: db,
	}

}

func (classRepository *classRepository) FindAllClass(allClass *[]entity.ClassResponse) error {

	err := classRepository.db.Model(&entity.Class{}).Find(&allClass).Error
	if err != nil {
		return err
	}

	return nil

}

func (classRepository *classRepository) ELFindClassByClassCode(inputClass *entity.Class, classCode string) error {

	err := classRepository.db.Model(&entity.Class{}).Preload("Course").First(inputClass, "class_code = ?", classCode).Error
	if err != nil {
		return err
	}

	return nil

}

func (classRepository *classRepository) FindClassById(class *entity.Class, classId uint) error {

	err := classRepository.db.First(class, classId).Error
	if err != nil {
		return err
	}

	return nil

}
