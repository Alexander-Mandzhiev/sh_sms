package private_school_models

import (
	"errors"
	"fmt"
)

var (
	ErrEmptySubjectName      = errors.New("subject name cannot be empty")
	ErrInvalidSubjectID      = errors.New("invalid subject ID")
	ErrNotFoundSubjectName   = errors.New("subject name not found")
	ErrDuplicateSubjectName  = errors.New("subject name already exists")
	ErrDeleteSubjectConflict = errors.New("subject cannot be deleted due to existing references")
	ErrPermissionDenied      = errors.New("permission denied")

	ErrEmptyFullName         = fmt.Errorf("full name is required")
	ErrInvalidClientID       = fmt.Errorf("client_id is not a valid UUID")
	ErrInvalidTeacherID      = fmt.Errorf("teacher_id is not a valid UUID")
	ErrInvalidPhone          = errors.New("phone number is not valid")
	ErrInvalidEmail          = errors.New("email address is not valid")
	ErrTeacherNotFound       = errors.New("teacher not found")
	ErrDuplicateTeacher      = errors.New("teacher already exists")
	ErrDeleteTeacherConflict = errors.New("teacher cannot be deleted due to existing references")
	ErrInvalidLimitRange     = errors.New("limit out of range")
	ErrCreateFailed          = errors.New("create failed due to error")
	ErrInvalidClient         = errors.New("client is invalid")
)
