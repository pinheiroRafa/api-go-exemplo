package entities

type UserToken struct {
	Id     string `json:"id"`
	Email  string `json:"email"`
	Status int8   `json:"status"`
}
