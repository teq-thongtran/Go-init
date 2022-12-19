package presenter

import (
	"myapp/model"
)

type CardResponseWrapper struct {
	Card *model.Card `json:"card"`
}

type ListCardResponseWrapper struct {
	Cards []model.Card `json:"cards"`
	Meta  interface{}  `json:"meta"`
}
