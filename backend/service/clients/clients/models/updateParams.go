package models

type UpdateParams struct {
	ID          string
	Name        *string
	Description *string
	TypeID      *int
	Website     *string
}
