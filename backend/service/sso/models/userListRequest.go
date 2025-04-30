package models

import "github.com/google/uuid"

type UserListRequest struct {
	UserID     uuid.UUID
	ClientID   uuid.UUID
	AppID      int
	ActiveOnly bool
	Page       int
	Count      int
}
