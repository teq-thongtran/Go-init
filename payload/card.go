package payload

type CreateCardRequest struct {
	NameCard string `json:"name_card"`
	CardType string `json:"card_type"`
	UserId   int64  `json:"user_id"`
}

type UpdateCardRequest struct {
	ID       int64   `json:"-"`
	CardType *string `json:"card_type"`
	NameCard *string `json:"name_card"`
	UserId   int64   `json:"user_id"`
}
