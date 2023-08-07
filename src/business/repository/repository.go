package repository

import "gorm.io/gorm"

type Repository struct {
	User    UserRepository
	Student StudentRepository
	Class   ClassRepository
}

func InjectRepository(db *gorm.DB) *Repository {
	return &Repository{
		User:    NewUserRepository(db),
		Student: NewStudentRepository(db),
		Class:   NewClassRepository(db),
	}
}
