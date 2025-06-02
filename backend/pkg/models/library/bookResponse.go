package library_models

import (
	"time"

	library "backend/protos/gen/go/library"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type BookResponse struct {
	ID          int64     `json:"id"`
	ClientID    string    `json:"client_id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Description string    `json:"description"`
	SubjectName string    `json:"subject_name"`
	Grade       int32     `json:"grade"`
	CreatedAt   time.Time `json:"created_at"`
}

func BookResponseFromProto(pbResp *library.BookResponse) *BookResponse {
	if pbResp == nil {
		return nil
	}

	return &BookResponse{
		ID:          pbResp.GetId(),
		ClientID:    pbResp.GetClientId(),
		Title:       pbResp.GetTitle(),
		Author:      pbResp.GetAuthor(),
		Description: pbResp.GetDescription(),
		SubjectName: pbResp.GetSubjectName(),
		Grade:       pbResp.GetGrade(),
		CreatedAt:   pbResp.GetCreatedAt().AsTime(),
	}
}

func (b *BookResponse) BookResponseToProto() *library.BookResponse {
	return &library.BookResponse{
		Id:          b.ID,
		ClientId:    b.ClientID,
		Title:       b.Title,
		Author:      b.Author,
		Description: b.Description,
		SubjectName: b.SubjectName,
		Grade:       b.Grade,
		CreatedAt:   timestamppb.New(b.CreatedAt),
	}
}
