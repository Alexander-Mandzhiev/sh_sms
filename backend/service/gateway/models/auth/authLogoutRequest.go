package auth_models

type AuthLogoutRequest struct {
	AccessToken string `json:"access_token" binding:"required"`
}
