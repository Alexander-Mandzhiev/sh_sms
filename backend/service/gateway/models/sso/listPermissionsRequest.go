package sso_models

import "backend/service/gateway/models"

type ListPermissionsRequest struct {
	AppID      int32             `json:"app_id" binding:"required,min=1"`
	Filters    PermissionFilters `json:"filters,omitempty"`
	Pagination models.Pagination `json:"pagination" binding:"required"`
}
