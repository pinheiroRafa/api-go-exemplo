package user_handler

type RequestLoginUser struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}
