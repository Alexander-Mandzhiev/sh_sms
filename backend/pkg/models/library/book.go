package library_models

import (
	library "backend/protos/gen/go/library"
	"github.com/google/uuid"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Book struct {
	ID          int64     `json:"id"`
	ClientID    uuid.UUID `json:"client_id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Description string    `json:"description"`
	SubjectID   int       `json:"subject_id"`
	ClassID     int       `json:"class_id"`
	CreatedAt   time.Time `json:"created_at"`
}

func (b *Book) BookToProto() *library.Book {
	return &library.Book{
		Id:          b.ID,
		ClientId:    b.ClientID.String(),
		Title:       b.Title,
		Author:      b.Author,
		Description: b.Description,
		SubjectId:   int32(b.SubjectID),
		ClassId:     int32(b.ClassID),
		CreatedAt:   timestamppb.New(b.CreatedAt),
	}
}

func BookFromProto(pbBook *library.Book) (*Book, error) {
	clientID, err := uuid.Parse(pbBook.GetClientId())
	if err != nil {
		return nil, err
	}
	return &Book{
		ID:          pbBook.GetId(),
		ClientID:    clientID,
		Title:       pbBook.GetTitle(),
		Author:      pbBook.GetAuthor(),
		Description: pbBook.GetDescription(),
		SubjectID:   int(pbBook.GetSubjectId()),
		ClassID:     int(pbBook.GetClassId()),
		CreatedAt:   pbBook.GetCreatedAt().AsTime(),
	}, nil
}
