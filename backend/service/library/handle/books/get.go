package books_handle

import (
	sl "backend/pkg/logger"
	"backend/pkg/models/library"
	library "backend/protos/gen/go/library"
	"context"
	"github.com/google/uuid"
	"log/slog"
)

func (s *serverAPI) GetBook(ctx context.Context, req *library.GetBookRequest) (*library.Book, error) {
	const op = "grpc.Library.Books.GetBook"
	logger := s.logger.With(slog.String("op", op), slog.Int64("id", req.GetId()), slog.String("client_id", req.GetClientId()))
	logger.Debug("Get book called")

	if req.GetId() <= 0 {
		logger.Warn("Invalid book ID")
		return nil, s.convertError(library_models.ErrInvalidID)
	}

	clientID, err := uuid.Parse(req.GetClientId())
	if err != nil {
		logger.Warn("Invalid client ID format", slog.String("client_id", req.GetClientId()))
		return nil, s.convertError(library_models.ErrBookInvalidClientID)
	}

	book, err := s.service.GetBook(ctx, req.GetId(), clientID)
	if err != nil {
		logger.Error("Failed to get book", sl.Err(err, true))
		return nil, s.convertError(err)
	}

	protoBook := library_models.BookToProto(book)
	logger.Info("Book retrieved", slog.Int64("id", protoBook.GetId()))
	return protoBook, nil
}
