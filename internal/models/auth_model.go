package models

type AuthLoginReqModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthSignupReqModel struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResModel struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshreqModel struct {
	RefreshToken string `json:"refresh_token"`
}
