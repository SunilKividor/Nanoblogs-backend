package models

type AuthLoginReqModel struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthSignupReqModel struct {
	Name     string   `json:"name" validate:"required"`
	Username string   `json:"username" validate:"required"`
	Password string   `json:"password" validate:"required"`
	Category []string `json:"category" validate:"required"`
}

type AuthResModel struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshreqModel struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}
