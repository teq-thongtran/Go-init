package model

import (
	"gorm.io/gorm"
	"time"
)

type Card struct {
	ID        int64           `json:"id"`
	NameCard  string          `json:"name_card"`
	CardType  string          `json:"card_type"`
	UserId    int64           `json:"user_id"`
	User      User            `json:"user"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"-"`
}
