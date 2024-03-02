package models

type AuthReqModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResModel struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserDetails struct {
	Id           int    `json:"id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Refreshreq struct {
	RefreshToken string `json:"refresh_token"`
}
