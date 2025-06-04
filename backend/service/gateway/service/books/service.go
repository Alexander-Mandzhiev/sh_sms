package books_service

import (
	library_models "backend/pkg/models/library"
	library "backend/protos/gen/go/library"
	"context"
	"log/slog"
)

type BooksService interface {
	CreateBook(ctx context.Context, params *library_models.CreateBookParams) (*library.BookResponse, error)
	GetBook(ctx context.Context, id int64, clientID string) (*library.BookResponse, error)
	UpdateBook(ctx context.Context, params *library_models.UpdateBookRequest) (*library.BookResponse, error)
	DeleteBook(ctx context.Context, id int64, clientID string) error
	ListBooks(ctx context.Context, req *library.ListBooksRequest) (*library.ListBooksResponse, error)
}

type booksService struct {
	client library.BookServiceClient
	logger *slog.Logger
}

func NewBooksService(client library.BookServiceClient, logger *slog.Logger) BooksService {
	return &booksService{
		client: client,
		logger: logger.With("service", "books"),
	}
}

func (s *booksService) CreateBook(ctx context.Context, params *library_models.CreateBookParams) (*library.BookResponse, error) {
	s.logger.Debug("Creating book", "client_id", params.ClientID, "title", params.Title)
	req := params.ToCreateRequestProto()
	return s.client.CreateBook(ctx, req)
}

func (s *booksService) GetBook(ctx context.Context, id int64, clientID string) (*library.BookResponse, error) {
	s.logger.Debug("Getting book", "id", id, "client_id", clientID)

	return s.client.GetBook(ctx, &library.GetBookRequest{Id: id, ClientId: clientID})
}

func (s *booksService) UpdateBook(ctx context.Context, params *library_models.UpdateBookRequest) (*library.BookResponse, error) {
	s.logger.Debug("Updating book", "id", params.ID, "client_id", params.ClientID)

	req := params.UpdateBookRequestToProto()
	return s.client.UpdateBook(ctx, req)
}

func (s *booksService) DeleteBook(ctx context.Context, id int64, clientID string) error {
	s.logger.Debug("Deleting book", "id", id, "client_id", clientID)
	_, err := s.client.DeleteBook(ctx, &library.DeleteBookRequest{Id: id, ClientId: clientID})
	return err
}

func (s *booksService) ListBooks(ctx context.Context, req *library.ListBooksRequest) (*library.ListBooksResponse, error) {
	s.logger.Debug("Listing books", "client_id", req.ClientId, "count", req.Count, "cursor", req.Cursor, "filter", req.Filter)
	return s.client.ListBooks(ctx, req)
}
