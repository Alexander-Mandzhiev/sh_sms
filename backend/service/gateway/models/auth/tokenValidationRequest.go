package auth_models

type TokenValidationRequest struct {
	TokenTypeHint string `json:"token_type_hint"`
}
