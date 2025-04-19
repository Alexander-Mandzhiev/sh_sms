package models

import "time"

type UserUpdate struct {
	Email     *string   `json:"email,omitempty"`
	FullName  *string   `json:"full_name,omitempty"`
	Phone     *string   `json:"phone,omitempty"`
	IsActive  *bool     `json:"is_active"`
	UpdatedAt time.Time `json:"-"`
}
