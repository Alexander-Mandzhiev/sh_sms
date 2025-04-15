package models

import "time"

type App struct {
	ID          int
	Code        string
	Name        string
	Description string
	IsActive    bool
	Version     int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
