package library_models

import (
	library "backend/protos/gen/go/library"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Attachment struct {
	BookID    int64     `json:"book_id"`
	Format    string    `json:"format"`
	FileID    string    `json:"file_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (a *Attachment) AttachmentToProto() *library.Attachment {
	att := &library.Attachment{
		BookId:    a.BookID,
		Format:    a.Format,
		FileId:    a.FileID,
		UpdatedAt: timestamppb.New(a.UpdatedAt),
		CreatedAt: timestamppb.New(a.CreatedAt),
	}
	return att
}

func AttachmentFromProto(a *library.Attachment) *Attachment {
	att := &Attachment{
		BookID:    a.BookId,
		Format:    a.Format,
		FileID:    a.FileId,
		CreatedAt: a.CreatedAt.AsTime(),
		UpdatedAt: a.UpdatedAt.AsTime(),
	}
	return att
}
