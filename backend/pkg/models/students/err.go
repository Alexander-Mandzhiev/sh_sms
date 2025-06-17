package students_models

import "errors"

var (
	ErrEmptyEmail          = errors.New("empty email")
	ErrEmailTooLong        = errors.New("email too long")
	ErrEmptyPhone          = errors.New("empty phone")
	ErrEmptyFullName       = errors.New("empty full name")
	ErrEmptyContractNumber = errors.New("empty contract number")

	ErrFullNameTooLong       = errors.New("full name too long")
	ErrContractNumberTooLong = errors.New("contract number too long")
	ErrPhoneTooLong          = errors.New("phone too long")
	ErrPhoneTooShort         = errors.New("phone number too short")
	ErrFilterValueTooLong    = errors.New("filter value too long, max 100 characters")

	ErrInvalidClientID   = errors.New("invalid client id")
	ErrInvalidPhone      = errors.New("invalid phone format")
	ErrInvalidEmail      = errors.New("invalid email format")
	ErrInvalidStudentID  = errors.New("invalid student id")
	ErrInvalidCursor     = errors.New("invalid cursor format")
	ErrInvalidCountValue = errors.New("invalid count value, must be between 1 and 1000")

	ErrStudentNotFound       = errors.New("student not found")
	ErrStudentAlreadyDeleted = errors.New("student already deleted")
	ErrStudentNotDeleted     = errors.New("student not deleted")

	ErrDuplicateContract = errors.New("duplicate contract")
	ErrRepositoryFailure = errors.New("repository operation failed")
	ErrUnexpectedError   = errors.New("unexpected error occurred")
	ErrCreateFailed      = errors.New("create student failed")
	ErrUpdateFailed      = errors.New("update student failed")
)
