package subjects_models

import (
	"errors"
)

var (
	ErrEmptySubjectName      = errors.New("subject name cannot be empty")
	ErrInvalidSubjectID      = errors.New("invalid subject ID")
	ErrNotFoundSubjectName   = errors.New("subject name not found")
	ErrDuplicateSubjectName  = errors.New("subject name already exists")
	ErrDeleteSubjectConflict = errors.New("subject cannot be deleted due to existing references")
	ErrPermissionDenied      = errors.New("permission denied")
)
