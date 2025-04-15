package models

type UpdateApp struct {
	Code        *string
	Name        *string
	Description *string
	IsActive    *bool
}

func (ua *UpdateApp) HasUpdates() bool {
	return ua.Code != nil ||
		ua.Name != nil ||
		ua.Description != nil ||
		ua.IsActive != nil
}
