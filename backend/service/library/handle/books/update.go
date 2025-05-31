package books_handle

import (
	sl "backend/pkg/logger"
	"backend/pkg/models/library"
	library "backend/protos/gen/go/library"
	"context"
	"log/slog"
)

func (s *serverAPI) UpdateBook(ctx context.Context, req *library.UpdateBookRequest) (*library.Book, error) {
	const op = "grpc.Library.Books.UpdateBook"
	logger := s.logger.With(slog.String("op", op), slog.Int64("id", req.GetId()), slog.String("client_id", req.GetClientId()))
	logger.Debug("Update book called")

	params, err := library_models.UpdateParamsFromProto(req)
	if err != nil {
		logger.Warn("Invalid update parameters", sl.Err(err, true))
		return nil, s.convertError(err)
	}

	params.Sanitize()

	book, err := s.service.UpdateBook(ctx, params)
	if err != nil {
		logger.Error("Failed to update book", sl.Err(err, true))
		return nil, s.convertError(err)
	}

	protoBook := library_models.BookToProto(book)
	logger.Info("Book updated", slog.Int64("id", protoBook.GetId()))
	return protoBook, nil
}
