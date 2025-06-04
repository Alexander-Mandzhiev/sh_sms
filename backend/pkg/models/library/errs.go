package library_models

import "errors"

var (
	ErrInvalidID        = errors.New("invalid id")
	ErrNotFound         = errors.New("not found")
	ErrPermissionDenied = errors.New("permission denied")
	ErrClientIDRequired = errors.New("client ID is required")
)

var (
	ErrInvalidClientID = errors.New("invalid client ID format")
	ErrEmptyTitle      = errors.New("title cannot be empty")
	ErrEmptyAuthor     = errors.New("author cannot be empty")
	ErrInvalidSubject  = errors.New("subject ID must be positive")
	ErrInvalidClass    = errors.New("class ID must be between 1 and 11")
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
	ErrAttachmentNotFound         = errors.New("attachment not found")
	ErrAttachmentAlreadyExists    = errors.New("attachment already exists")
	ErrAttachmentExistsButDeleted = errors.New("attachment exists but deleted")
	ErrAttachmentAlreadyActive    = errors.New("attachment is already active")
	ErrInvalidAttachmentFormat    = errors.New("invalid attachment format")
	ErrEmptyFileURL               = errors.New("file URL cannot be empty")
	ErrAttachmentInvalidBookID    = errors.New("invalid book ID")
	ErrAttachmentRestoreConflict  = errors.New("active attachment already exists, cannot restore")
	ErrAttachmentUpdateConflict   = errors.New("conflict during attachment update")
	ErrAttachmentAlreadyDeleted   = errors.New("attachment is already deleted")
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
