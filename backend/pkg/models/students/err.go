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

	ErrInvalidClientID  = errors.New("invalid client id")
	ErrInvalidPhone     = errors.New("invalid phone format")
	ErrInvalidEmail     = errors.New("invalid email format")
	ErrInvalidStudentID = errors.New("invalid student id")

	ErrStudentNotFound = errors.New("student not found")
)
