package payload

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Score    int    `json:"score"`
}

type UpdateUserRequest struct {
	ID       int64   `json:"-"`
	Name     *string `json:"name"`
	Score    *int    `json:"score"`
	Username *string `json:"username"`
	Email    *string `json:"email"`
}
