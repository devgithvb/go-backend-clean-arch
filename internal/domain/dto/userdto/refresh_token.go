package userdto

type RefreshTokenRequest struct {
	RefreshToken string `form:"refresh_token" json:"refresh_token"`
}

type RefreshTokenResponse struct {
	Tokens Tokens `json:"tokens"`
}
