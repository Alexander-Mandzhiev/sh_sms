package library_models

import (
	library "backend/protos/gen/go/library"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Book struct {
	ID          int64     `json:"id"`
	ClientID    string    `json:"client_id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Description string    `json:"description"`
	SubjectID   int32     `json:"subject_id"`
	ClassID     int32     `json:"class_id"`
	CreatedAt   time.Time `json:"created_at"`
}

func ProtoToBook(pb *library.Book) *Book {
	return &Book{
		ID:          pb.GetId(),
		ClientID:    pb.GetClientId(),
		Title:       pb.GetTitle(),
		Author:      pb.GetAuthor(),
		Description: pb.GetDescription(),
		SubjectID:   pb.GetSubjectId(),
		ClassID:     pb.GetClassId(),
		CreatedAt:   pb.GetCreatedAt().AsTime(),
	}
}

func BookToProto(b *Book) *library.Book {
	return &library.Book{
		Id:          b.ID,
		ClientId:    b.ClientID,
		Title:       b.Title,
		Author:      b.Author,
		Description: b.Description,
		SubjectId:   b.SubjectID,
		ClassId:     b.ClassID,
		CreatedAt:   timestamppb.New(b.CreatedAt),
	}
}
