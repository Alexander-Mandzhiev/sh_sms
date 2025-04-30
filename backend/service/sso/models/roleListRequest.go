package models

import "github.com/google/uuid"

type RoleListRequest struct {
	RoleID     uuid.UUID
	ClientID   uuid.UUID
	AppID      int
	ActiveOnly bool
	Page       int
	Count      int
}
