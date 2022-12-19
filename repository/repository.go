package repository

import (
	"context"
	"myapp/repository/card"

	"gorm.io/gorm"

	"myapp/repository/user"
)

type Repository struct {
	User user.Repository
	Card card.Repository
}

func New(getClient func(ctx context.Context) *gorm.DB) *Repository {
	return &Repository{
		User: user.NewPG(getClient),
		Card: card.NewPG(getClient),
	}
}
