package groups_models

import "errors"

var (
	ErrInvalidGroupID        = errors.New("invalid group ID format")
	ErrInvalidClientID       = errors.New("invalid client ID format")
	ErrInvalidCuratorID      = errors.New("invalid curator ID format")
	ErrGroupNameRequired     = errors.New("group name is required")
	ErrGroupNameTooLong      = errors.New("group name exceeds maximum length (100 characters)")
	ErrInvalidPageSize       = errors.New("page size must be between 1 and 100")
	ErrFilterValueTooLong    = errors.New("filter value too long, max 100 characters")
	ErrDuplicateGroupName    = errors.New("group name already exists for this client")
	ErrCreateFailed          = errors.New("failed to create group")
	ErrGroupNotFound         = errors.New("group not found")
	ErrGetFailed             = errors.New("failed to get group")
	ErrUpdateFailed          = errors.New("failed to update group")
	ErrDependentRecordsExist = errors.New("dependent records already exist")
	ErrForeignKeyViolation   = errors.New("foreign key violation")
	ErrDeleteFailed          = errors.New("failed to delete group")
	ErrListFailed            = errors.New("failed to list group")
	ErrInvalidCursor         = errors.New("invalid cursor value")
)
