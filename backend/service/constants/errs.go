package constants

import "errors"

var (
	ErrInvalidArgument     = errors.New("invalid argument")
	ErrInternal            = errors.New("internal server error")
	ErrPermissionDenied    = errors.New("permission denied")
	ErrInvalidID           = errors.New("invalid id")
	ErrEmptyName           = errors.New("name cannot be empty")
	ErrEmptyCode           = errors.New("code cannot be empty")
	ErrInvalidName         = errors.New("name length invalid ")
	ErrInvalidCode         = errors.New("code length invalid ")
	ErrConflictParams      = errors.New("conflicting parameters")
	ErrNoUpdateFields      = errors.New("no fields to update")
	ErrInvalidPagination   = errors.New("invalid pagination parameters")
	ErrIdentifierRequired  = errors.New("either id or code must be provided")
	ErrAlreadyExists       = errors.New("application already exists")
	ErrSecretAlreadyExists = errors.New("secret already exists")
	ErrNotFound            = errors.New("application not found")
	ErrUpdateConflict      = errors.New("update conflict")
	ErrVersionConflict     = errors.New("version conflict")
	ErrInvalidAppId        = errors.New("invalid app id")
	ErrAlreadyRevoked      = errors.New("already revoked")
	ErrRotationTooFrequent = errors.New("rotation allowed no more than once per 24 hours")
	ErrInvalidPassword     = errors.New("invalid password")
	ErrUnauthenticated     = errors.New("unauthenticated")
)
