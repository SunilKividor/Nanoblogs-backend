package models

import "github.com/google/uuid"

type AuthReqModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResModel struct {
	Userid       uuid.UUID `json:"id"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
}

type RefreshreqModel struct {
	RefreshToken string `json:"refresh_token"`
}
