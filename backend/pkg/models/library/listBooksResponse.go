package library_models

import (
	library "backend/protos/gen/go/library"
)

type ListBooksResponse struct {
	Books      []*BookResponse `json:"books"`
	NextCursor int64           `json:"next_cursor,omitempty"`
	TotalCount int32           `json:"total_count"`
}

func ListBooksResponseFromProto(pbResp *library.ListBooksResponse) *ListBooksResponse {
	resp := &ListBooksResponse{
		Books:      make([]*BookResponse, 0, len(pbResp.GetBooks())),
		NextCursor: pbResp.NextCursor,
	}

	for _, pbBook := range pbResp.GetBooks() {
		resp.Books = append(resp.Books, BookResponseFromProto(pbBook))
	}

	return resp
}

func ListBooksResponseToProto(req *ListBooksResponse) *library.ListBooksResponse {
	booksResponse := make([]*library.BookResponse, 0, len(req.Books))
	for _, pbBook := range req.Books {
		book := pbBook.BookResponseToProto()
		booksResponse = append(booksResponse, book)
	}

	return &library.ListBooksResponse{
		Books:      booksResponse,
		NextCursor: req.NextCursor,
		TotalCount: req.TotalCount,
	}
}
