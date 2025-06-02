package books_handle

import (
	sl "backend/pkg/logger"
	"backend/pkg/models/library"
	library "backend/protos/gen/go/library"
	"context"
	"log/slog"
)

func (s *serverAPI) ListBooks(ctx context.Context, req *library.ListBooksRequest) (*library.ListBooksResponse, error) {
	const op = "grpc.Library.Books.ListBooks"
	logger := s.logger.With(slog.String("op", op))
	logger.Debug("List books called")

	params, err := library_models.ListBooksRequestFromProto(req)
	if err != nil {
		logger.Warn("Invalid list parameters", sl.Err(err, true))
		return nil, s.convertError(err)
	}

	result, err := s.service.ListBooks(ctx, params)
	if err != nil {
		logger.Error("Failed to list books", sl.Err(err, true))
		return nil, s.convertError(err)
	}

	response := library_models.ListBooksResponseToProto(result)
	logger.Info("Books listed", slog.Int("count", len(response.Books)))
	return response, nil
}
