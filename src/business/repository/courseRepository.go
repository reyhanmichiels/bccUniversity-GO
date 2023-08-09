package repository

import (
	"bcc-university/src/business/entity"

	"gorm.io/gorm"
)

type CourseRepository interface {
	FindCourseByCondition(course interface{}, condition string, value interface{}) error
}

type courseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) CourseRepository {

	return &courseRepository{
		db: db,
	}

}

func (courseRepository *courseRepository) FindCourseByCondition(course interface{}, condition string, value interface{}) error {

	err := courseRepository.db.Model(&entity.Course{}).First(course, condition, value).Error
	if err != nil {

		return err

	}

	return nil

}
