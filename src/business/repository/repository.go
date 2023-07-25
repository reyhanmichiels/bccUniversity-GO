package repository

import "gorm.io/gorm"

type Repository struct {
	User UserRepository
}

func InjectRepository(db *gorm.DB) *Repository {
	return &Repository{User: NewUserRepository(db)}
}