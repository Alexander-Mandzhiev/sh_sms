package sso_models

type PermissionFilters struct {
	CodeFilter *string `json:"code_filter,omitempty"`
	Category   *string `json:"category,omitempty"`
	ActiveOnly *bool   `json:"active_only,omitempty"`
}
