package usecase

import "bcc-university/src/business/repository"

type UseCase struct {

}

func InjectUseCase(r *repository.Repository) *UseCase {
	return &UseCase{}
}