package repository

import (
	"bcc-university/src/business/entity"

	"gorm.io/gorm"
)

type ClassRepository interface {
	FindAllClass(allClass *[]entity.ClassResponse) error
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
