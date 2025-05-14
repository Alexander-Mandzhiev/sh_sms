package models

type ListRolesRequest struct {
	ClientID   string      `json:"client_id" binding:"required"`
	AppID      int         `json:"app_id" binding:"required,min=1"`
	Filters    RoleFilters `json:"filters,omitempty"`
	Pagination Pagination  `json:"pagination" binding:"required"`
}
