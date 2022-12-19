package payload

import (
	"strings"
)

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Score    int    `json:"score"`
}

type GetByIDRequest struct {
	ID int64 `json:"-"`
}

var orderByUser = []string{"id", "name", "created_by", "updated_by"}

type GetListUserRequest struct {
	Page    int    `json:"page" query:"page"`
	Limit   int    `json:"limit" query:"limit"`
	OrderBy string `json:"order_by,omitempty" query:"order_by"`
	Search  string `json:"search,omitempty" query:"search"`
}

func (g *GetListUserRequest) Format() {
	g.Search = strings.TrimSpace(g.Search)
	g.OrderBy = strings.ToLower(strings.TrimSpace(g.OrderBy))

	for i := range orderByUser {
		if g.OrderBy == orderByUser[i] {
			return
		}
	}

	g.OrderBy = ""
}

type UpdateUserRequest struct {
	ID       int64   `json:"-"`
	Name     *string `json:"name"`
	Score    *int    `json:"score"`
	Username *string `json:"username"`
	Email    *string `json:"email"`
}

type DeleteUserRequest struct {
	ID int64 `json:"-"`
}
