package models

type CreateParams struct {
	Code        string
	Name        string
	Description *string
	IsActive    *bool
}
