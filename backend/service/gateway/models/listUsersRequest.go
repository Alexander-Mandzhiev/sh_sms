package models

type ListUsersRequest struct {
	ClientID   string      `json:"client_id" binding:"required"`
	Filters    UserFilters `json:"filters,omitempty"`
	Pagination Pagination  `json:"pagination" binding:"required"`
}
