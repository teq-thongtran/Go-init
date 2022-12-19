package usecase

import (
	"myapp/repository"
	"myapp/usecase/card"
	"myapp/usecase/user"
)

type UseCase struct {
	User user.UserUserCase
	Card card.CardCardCase
}

func New(repo *repository.Repository) *UseCase {
	return &UseCase{
		User: user.New(repo),
		Card: card.New(repo),
	}
}
