package models

type RoleUsersListRequest struct {
	RoleID   string `json:"role_id" binding:"required,uuid"`
	ClientID string `json:"client_id" binding:"required,uuid"`
	AppID    int32  `json:"app_id" binding:"required,min=1"`
	Page     int32  `json:"page" binding:"min=1"`
	Count    int32  `json:"count" binding:"min=1,max=100"`
}
