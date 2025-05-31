package library_models

import (
	library "backend/protos/gen/go/library"
	"fmt"
	"strings"
)

type ListBooksParams struct {
	ClientID  string           `json:"client_id"`
	PageSize  int32            `json:"page_size"`
	PageToken string           `json:"page_token"`
	Filter    string           `json:"filter"`
	Cursor    *ListBooksCursor `json:"cursor,omitempty"`
}

func ListParamsFromProto(req *library.ListBooksRequest) (*ListBooksParams, error) {

	if req.GetClientId() == "" {
		return nil, ErrClientIDRequired
	}

	params := &ListBooksParams{
		ClientID: req.GetClientId(),
		Filter:   req.GetFilter(),
	}

	if req.PageSize != nil {
		params.PageSize = *req.PageSize
	}
	if req.PageToken != nil {
		params.PageToken = *req.PageToken
		if cursor, err := decodeCursor(*req.PageToken); err == nil {
			params.Cursor = cursor
		}
	}

	if err := params.Validate(); err != nil {
		return nil, err
	}

	return params, nil
}

type ListBooksResult struct {
	Books         []*Book `json:"books"`
	NextPageToken string  `json:"next_page_token"`
	TotalCount    int32   `json:"total_count"`
}

func (r *ListBooksResult) ToListResponseProto() *library.ListBooksResponse {
	resp := &library.ListBooksResponse{
		Books:         make([]*library.Book, len(r.Books)),
		NextPageToken: r.NextPageToken,
		TotalCount:    r.TotalCount,
	}

	for i, book := range r.Books {
		resp.Books[i] = BookToProto(book)
	}

	return resp
}

func (p *ListBooksParams) Validate() error {
	if p.ClientID == "" {
		return ErrClientIDRequired
	}
	if p.PageSize < 0 {
		return fmt.Errorf("%w: %d", ErrInvalidPageSize, p.PageSize)
	}
	return nil
}

func (p *ListBooksParams) Sanitize() {
	p.Filter = strings.TrimSpace(p.Filter)
}
