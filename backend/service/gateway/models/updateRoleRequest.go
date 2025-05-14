package models

type UpdateRoleRequest struct {
	Name        *string `json:"name,omitempty" binding:"omitempty,min=3,max=64"`
	Description *string `json:"description,omitempty" binding:"omitempty,max=255"`
	Level       *int32  `json:"level,omitempty" binding:"omitempty,min=1,max=1000"`
}
