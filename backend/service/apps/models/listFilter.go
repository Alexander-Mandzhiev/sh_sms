package models

type ListFilter struct {
	Page     int
	Count    int
	ClientID *string
	AppID    *int32
	IsActive *bool
}
