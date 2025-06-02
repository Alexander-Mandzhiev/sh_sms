package library_models

import (
	library "backend/protos/gen/go/library"
	"fmt"
	"github.com/google/uuid"
	"strings"
)

const (
	maxPageSize     = 100
	defaultPageSize = 10
)

type ListBooksRequest struct {
	ClientID uuid.UUID `json:"client_id"`
	Count    int32     `json:"count"`
	Cursor   *int64    `json:"cursor"`
	Filter   *string   `json:"filter"`
}

func ListBooksRequestFromProto(pbReq *library.ListBooksRequest) (*ListBooksRequest, error) {
	clientID, err := uuid.Parse(pbReq.GetClientId())
	if err != nil {
		return nil, fmt.Errorf("invalid client ID: %w", err)
	}

	count := defaultPageSize
	if pbReq.Count != nil {
		count = int(*pbReq.Count)
		if count <= 0 {
			count = defaultPageSize
		} else if count > maxPageSize {
			count = maxPageSize
		}
	}

	req := &ListBooksRequest{
		ClientID: clientID,
		Count:    int32(count),
	}

	if pbReq.Cursor != nil {
		if *pbReq.Cursor <= 0 {
			return nil, fmt.Errorf("invalid cursor value: %d", *pbReq.Cursor)
		}
		req.Cursor = pbReq.Cursor
	}

	if pbReq.Filter != nil {
		filter := strings.TrimSpace(*pbReq.Filter)
		if len(filter) > 100 {
			filter = filter[:100]
		}
		req.Filter = &filter
	}

	return req, nil
}

func (r *ListBooksRequest) ListBooksRequestToProto() *library.ListBooksRequest {
	pbReq := &library.ListBooksRequest{
		ClientId: r.ClientID.String(),
		Count:    &r.Count,
	}

	if r.Cursor != nil {
		pbReq.Cursor = r.Cursor
	}

	if r.Filter != nil {
		pbReq.Filter = r.Filter
	}

	return pbReq
}
