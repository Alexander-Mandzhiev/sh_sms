package library_models

import "errors"

var (
	ErrInvalidID        = errors.New("invalid id")
	ErrNotFound         = errors.New("not found")
	ErrPermissionDenied = errors.New("permission denied")
	ErrClientIDRequired = errors.New("client ID is required")
)

var (
	ErrBookInvalidClientID  = errors.New("invalid client id format")
	ErrBookEmptyTitle       = errors.New("title cannot be empty")
	ErrBookEmptyAuthor      = errors.New("author cannot be empty")
	ErrBookInvalidSubjectID = errors.New("subject ID must be positive")
	ErrBookInvalidClassID   = errors.New("class ID must be between 1 and 11")
	ErrBookDescriptionLong  = errors.New("description too long")
	ErrInvalidPageSize      = errors.New("invalid page size")
	ErrBookInvalidTitle     = errors.New("invalid title")
	ErrBookInvalidAuthor    = errors.New("invalid author")
)

var (
	ErrEmptyName        = errors.New("subject name cannot be empty")
	ErrInvalidSubjectID = errors.New("invalid subject ID")

	ErrDuplicateName  = errors.New("subject name already exists")
	ErrDeleteConflict = errors.New("subject cannot be deleted due to existing references")
)

var (
	ErrClassInvalidGrade = errors.New("invalid class grade (must be 1-11)")
)
