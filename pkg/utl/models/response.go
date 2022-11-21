package models

type LoginResponse struct {
	Status       int    `json:"status"`
	Message      string `json:"message"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    int    `json:"expiresIn"`
	// Data         *User
}

type ResetPasswordResponse struct {
	Status  int  `json:"status"`
	Message bool `json:"message"`
	//Data	*User
}
