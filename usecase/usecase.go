package usecase

import (
	"myapp/repository"
	"myapp/usecase/user"
)

type UseCase struct {
	User user.IUseCase
}

func New(repo *repository.Repository) *UseCase {
	return &UseCase{
		User: user.New(repo),
	}
}
