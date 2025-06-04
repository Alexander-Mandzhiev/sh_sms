package library_models

import (
	library "backend/protos/gen/go/library"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UpdateAttachmentRequest struct {
	BookID     int64  `json:"book_id"`
	Format     string `json:"format"`
	NewFileURL string `json:"new_file_url"`
}

func (a *UpdateAttachmentRequest) UpdateAttachmentRequestToProto() *library.UpdateAttachmentRequest {
	return &library.UpdateAttachmentRequest{
		BookId:     a.BookID,
		Format:     a.Format,
		NewFileUrl: a.NewFileURL,
	}
}

func UpdateAttachmentRequestFromProto(a *library.UpdateAttachmentRequest) (*UpdateAttachmentRequest, error) {
	if err := ValidateBookID(a.BookId); err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid book ID")
	}
	if err := ValidateAttachmentFormat(a.Format); err != nil {
		return nil, status.Error(codes.InvalidArgument, "format is required")
	}
	if err := ValidateFileURL(a.NewFileUrl); err != nil {
		return nil, status.Error(codes.InvalidArgument, "new file URL is required")
	}

	return &UpdateAttachmentRequest{
		BookID:     a.BookId,
		Format:     a.Format,
		NewFileURL: a.NewFileUrl,
	}, nil
}
func (a *UpdateAttachmentRequest) Validate() error {
	if err := ValidateBookID(a.BookID); err != nil {
		return status.Error(codes.InvalidArgument, "invalid book ID")
	}
	if err := ValidateAttachmentFormat(a.Format); err != nil {
		return status.Error(codes.InvalidArgument, "file URL is required")
	}

	if err := ValidateFileURL(a.NewFileURL); err != nil {
		return status.Error(codes.InvalidArgument, "file URL is required")
	}

	return nil
}
