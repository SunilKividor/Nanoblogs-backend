package models

type User struct {
	Id           int    `json:"id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
