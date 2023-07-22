package repository

import "gorm.io/gorm"

type Repository struct {

}

func InjectRepository(db *gorm.DB) *Repository {
	return &Repository{}
}