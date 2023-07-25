package usecase

import "bcc-university/src/business/repository"

type UseCase struct {
	User UserUseCase
}

func InjectUseCase(r *repository.Repository) *UseCase {
	return &UseCase{
		User: NewUserUseCase(r.User),
	}
}