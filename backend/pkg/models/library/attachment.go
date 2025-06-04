package library_models

import (
	library "backend/protos/gen/go/library"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Attachment struct {
	BookID    int64      `json:"book_id"`
	Format    string     `json:"format"`
	FileURL   string     `json:"file_url"`
	DeletedAt *time.Time `json:"deleted_at"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func (a *Attachment) AttachmentToProto() *library.Attachment {
	att := &library.Attachment{
		BookId:    a.BookID,
		Format:    a.Format,
		FileUrl:   a.FileURL,
		UpdatedAt: timestamppb.New(a.UpdatedAt),
		CreatedAt: timestamppb.New(a.CreatedAt),
	}

	if a.DeletedAt != nil {
		att.DeletedAt = timestamppb.New(*a.DeletedAt)
	}
	return att
}

func AttachmentFromProto(a *library.Attachment) *Attachment {
	att := &Attachment{
		BookID:    a.BookId,
		Format:    a.Format,
		FileURL:   a.FileUrl,
		CreatedAt: a.CreatedAt.AsTime(),
		UpdatedAt: a.UpdatedAt.AsTime(),
	}

	if a.DeletedAt != nil {
		deletedAt := a.DeletedAt.AsTime()
		att.DeletedAt = &deletedAt
	}

	return att
}
