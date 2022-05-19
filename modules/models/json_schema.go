package models

type ReqLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// Token struct
type ResToken struct {
	TokenType string  `json:"token_type"`
	Duration  float64 `json:"duration"`
	Token     string  `json:"access_token"`
}
