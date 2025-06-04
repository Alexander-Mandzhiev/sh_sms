package library_models

import (
	library "backend/protos/gen/go/library"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CreateAttachmentRequest struct {
	BookID  int64  `json:"book_id"`
	Format  string `json:"format"`
	FileURL string `json:"file_url"`
}

func (a *CreateAttachmentRequest) CreateAttachmentRequestToProto() (*library.CreateAttachmentRequest, error) {
	if err := a.Validate(); err != nil {
		return nil, err
	}
	return &library.CreateAttachmentRequest{
		BookId:  a.BookID,
		Format:  a.Format,
		FileUrl: a.FileURL,
	}, nil
}

func CreateAttachmentRequestFromProto(a *library.CreateAttachmentRequest) *CreateAttachmentRequest {
	return &CreateAttachmentRequest{
		BookID:  a.BookId,
		Format:  a.Format,
		FileURL: a.FileUrl,
	}
}

func (a *CreateAttachmentRequest) Validate() error {
	if err := ValidateBookID(a.BookID); err != nil {
		return status.Error(codes.InvalidArgument, "invalid book ID")
	}
	if err := ValidateAttachmentFormat(a.Format); err != nil {
		return status.Error(codes.InvalidArgument, "file URL is required")
	}

	if err := ValidateFileURL(a.FileURL); err != nil {
		return status.Error(codes.InvalidArgument, "file URL is required")
	}

	return nil
}

func ValidateBookID(bookID int64) error {
	if bookID <= 0 {
		return status.Error(codes.InvalidArgument, "invalid book ID")
	}
	return nil
}

func ValidateAttachmentFormat(format string) error {
	if format == "" {
		return status.Error(codes.InvalidArgument, "format is required")
	}
	if len(format) > 10 {
		return status.Error(codes.InvalidArgument, "format exceeds maximum length (10 characters)")
	}
	return nil
}

func ValidateFileURL(url string) error {
	if url == "" {
		return status.Error(codes.InvalidArgument, "file URL is required")
	}
	return nil
}
