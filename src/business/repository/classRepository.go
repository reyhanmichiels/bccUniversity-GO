package repository

import (
	"bcc-university/src/business/entity"

	"gorm.io/gorm"
)

type ClassRepository interface {
	FindAllClass(allClass interface{}) error
	ELFindClassByCondition(class interface{}, condition string, value interface{}) error
	FindClassByCondition(class interface{}, condition string, value interface{}) error
	CreateClass(class *entity.Class) error
	Updateclass(class *entity.Class, updateData interface{}) error
	DeleteClass(class *entity.Class) error
}

type classRepository struct {
	db *gorm.DB
}

func NewClassRepository(db *gorm.DB) ClassRepository {

	return &classRepository{
		db: db,
	}

}

func (classRepository *classRepository) FindAllClass(allClass interface{}) error {

	err := classRepository.db.Model(&entity.Class{}).Find(allClass).Error
	if err != nil {

		return err

	}

	return nil

}

func (classRepository *classRepository) ELFindClassByCondition(class interface{}, condition string, value interface{}) error {

	err := classRepository.db.Model(&entity.Class{}).Preload("Course").First(class, condition, value).Error
	if err != nil {

		return err

	}

	return nil

}

func (classRepository *classRepository) FindClassByCondition(class interface{}, condition string, value interface{}) error {

	err := classRepository.db.First(class, condition, value).Error
	if err != nil {

		return err

	}

	return nil

}

func (classRepository *classRepository) CreateClass(class *entity.Class) error {

	err := classRepository.db.Create(class).Error
	if err != nil {

		return err

	}

	return nil

}

func (classRepository *classRepository) Updateclass(class *entity.Class, updateData interface{}) error {

	err := classRepository.db.Model(class).Updates(updateData).Error
	if err != nil {

		return err

	}

	return nil

}

func (classRepository *classRepository) DeleteClass(class *entity.Class) error {

	err := classRepository.db.Delete(class).Error
	if err != nil {

		return err

	}

	return nil

}
