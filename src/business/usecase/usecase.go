package usecase

import "bcc-university/src/business/repository"

type UseCase struct {
	User    UserUseCase
	Student StudentUseCase
	Class   ClassUseCase
}

func InjectUseCase(r *repository.Repository) *UseCase {
	return &UseCase{
		User:    NewUserUseCase(r.User, r.Class),
		Student: NewStudentUseCase(r.Student),
		Class:   NewClassUseCase(r.Class, r.User),
	}
}
