package models

type ListPermissionsRequest struct {
	AppID      int32             `json:"app_id" binding:"required,min=1"`
	Filters    PermissionFilters `json:"filters,omitempty"`
	Pagination Pagination        `json:"pagination" binding:"required"`
}
