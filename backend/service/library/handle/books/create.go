package books_handle

import (
	sl "backend/pkg/logger"
	"backend/pkg/models/library"
	library "backend/protos/gen/go/library"
	"context"
	"log/slog"
)

func (s *serverAPI) CreateBook(ctx context.Context, req *library.CreateBookRequest) (*library.Book, error) {
	const op = "grpc.Library.Books.CreateBook"
	logger := s.logger.With(slog.String("op", op))
	logger.Debug("Create book called")

	params, err := library_models.CreateParamsFromProto(req)
	if err != nil {
		logger.Warn("Invalid create book parameters", sl.Err(err, true))
		return nil, s.convertError(err)
	}

	book, err := s.service.CreateBook(ctx, params)
	if err != nil {
		logger.Error("Failed to create book", sl.Err(err, true))
		return nil, s.convertError(err)
	}

	protoBook := library_models.BookToProto(book)
	logger.Info("Book created", slog.Int64("id", protoBook.GetId()), slog.String("title", protoBook.GetTitle()))

	return protoBook, nil
}
