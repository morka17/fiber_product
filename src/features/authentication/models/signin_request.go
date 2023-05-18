package models

type SigninRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SigninResponse struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}
