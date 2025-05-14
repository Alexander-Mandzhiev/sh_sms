package models

type Pagination struct {
	Page  int `json:"page" binding:"required,min=1"`
	Count int `json:"count" binding:"required,min=1,max=1000"`
}
