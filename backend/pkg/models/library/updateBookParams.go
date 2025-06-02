package library_models

import (
	library "backend/protos/gen/go/library"
	"github.com/google/uuid"
)

type UpdateBookRequest struct {
	ID          int64     `json:"id"`
	ClientID    uuid.UUID `json:"client_id"`
	Title       *string   `json:"title"`
	Author      *string   `json:"author"`
	Description *string   `json:"description"`
	SubjectID   *int32    `json:"subject_id"`
	ClassID     *int32    `json:"class_id"`
}

func UpdateBookRequestFromProto(pbReq *library.UpdateBookRequest) (*UpdateBookRequest, error) {
	clientID, err := uuid.Parse(pbReq.GetClientId())
	if err != nil {
		return nil, err
	}

	req := &UpdateBookRequest{
		ID:       pbReq.GetId(),
		ClientID: clientID,
	}

	if pbReq.Title != nil {
		req.Title = pbReq.Title
	}
	if pbReq.Author != nil {
		req.Author = pbReq.Author
	}
	if pbReq.Description != nil {
		req.Description = pbReq.Description
	}
	if pbReq.SubjectId != nil {
		req.SubjectID = pbReq.SubjectId
	}
	if pbReq.ClassId != nil {
		req.ClassID = pbReq.ClassId
	}

	return req, nil
}

func (r *UpdateBookRequest) UpdateBookRequestToProto() *library.UpdateBookRequest {
	pbReq := &library.UpdateBookRequest{
		Id:       r.ID,
		ClientId: r.ClientID.String(),
	}

	if r.Title != nil {
		pbReq.Title = r.Title
	}
	if r.Author != nil {
		pbReq.Author = r.Author
	}
	if r.Description != nil {
		pbReq.Description = r.Description
	}
	if r.SubjectID != nil {
		pbReq.SubjectId = r.SubjectID
	}
	if r.ClassID != nil {
		pbReq.ClassId = r.ClassID
	}

	return pbReq
}
