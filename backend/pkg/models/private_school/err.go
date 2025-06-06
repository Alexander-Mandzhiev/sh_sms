package private_school_models

import "errors"

var (
	ErrEmptyName        = errors.New("subject name cannot be empty")
	ErrInvalidSubjectID = errors.New("invalid subject ID")

	ErrDuplicateName  = errors.New("subject name already exists")
	ErrDeleteConflict = errors.New("subject cannot be deleted due to existing references")
)
