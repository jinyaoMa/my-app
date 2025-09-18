package schema

import "time"

type AuthLoginRequest struct {
	Account  string `json:"account" doc:"User Account"`
	Password string `json:"password" doc:"User Password"`
}

type AuthLogin struct {
	AccessToken  string    `json:"accessToken" doc:"Access Token"`
	RefreshToken string    `json:"refreshToken" doc:"Refresh Token"`
	ExpiredAt    time.Time `json:"expiredAt" doc:"Expired At"`
}
