package sso_models

type RoleFilters struct {
	Name       *string `json:"name,omitempty"`
	Level      *int32  `json:"level,omitempty"`
	ActiveOnly *bool   `json:"active_only,omitempty"`
}
