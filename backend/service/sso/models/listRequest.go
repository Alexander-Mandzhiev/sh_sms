package models

import (
	"github.com/google/uuid"
)

type ListRequest struct {
	Page        int
	Count       int
	ClientID    *uuid.UUID
	EmailFilter *string
	PhoneFilter *string
	ActiveOnly  *bool
}
