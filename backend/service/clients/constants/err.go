package constants

import "errors"

var (
	ErrNotFound        = errors.New("not found")
	ErrInvalidArgument = errors.New("invalid argument")
	ErrAlreadyExists   = errors.New("already exists")
	ErrInternal        = errors.New("internal error")
)
