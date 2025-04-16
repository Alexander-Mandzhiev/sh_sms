package models

type ListFilter struct {
	Page     int
	Count    int
	ClientID *string
	AppID    *int
	IsActive *bool
}
