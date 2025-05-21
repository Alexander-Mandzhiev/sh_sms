package sso_models

import "backend/service/gateway/models"

type ListUsersRequest struct {
	ClientID   string            `json:"client_id" binding:"required"`
	Filters    UserFilters       `json:"filters,omitempty"`
	Pagination models.Pagination `json:"pagination" binding:"required"`
}
