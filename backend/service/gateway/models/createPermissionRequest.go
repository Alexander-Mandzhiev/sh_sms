package models

type CreatePermissionRequest struct {
	Code        string `json:"code" binding:"required,alphanum,max=50"`
	Description string `json:"description" binding:"max=255"`
	Category    string `json:"category" binding:"required,max=50"`
	AppID       int32  `json:"app_id" binding:"required,min=1"`
}
