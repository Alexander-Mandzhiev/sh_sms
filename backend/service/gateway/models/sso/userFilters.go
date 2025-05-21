package sso_models

type UserFilters struct {
	Email      *string `json:"email,omitempty"`
	Phone      *string `json:"phone,omitempty"`
	ActiveOnly *bool   `json:"active_only,omitempty"`
}
