package books_handle

import (
	sl "backend/pkg/logger"
	"backend/pkg/models/library"
	library "backend/protos/gen/go/library"
	"context"
	"github.com/google/uuid"
	"log/slog"
)

func (s *serverAPI) ListBooks(ctx context.Context, req *library.ListBooksRequest) (*library.ListBooksResponse, error) {
	const op = "grpc.Library.Books.ListBooks"
	logger := s.logger.With(slog.String("op", op))
	logger.Debug("List books called")

	if _, err := uuid.Parse(req.GetClientId()); err != nil {
		logger.Warn("Invalid client ID format", slog.String("client_id", req.GetClientId()))
		return nil, s.convertError(library_models.ErrBookInvalidClientID)
	}

	params, err := library_models.ListParamsFromProto(req)
	if err != nil {
		logger.Warn("Invalid list parameters", sl.Err(err, true))
		return nil, s.convertError(err)
	}

	params.Sanitize()

	const maxPageSize = 100
	if params.PageSize == 0 || params.PageSize > maxPageSize {
		params.PageSize = maxPageSize
	}

	result, err := s.service.ListBooks(ctx, params)
	if err != nil {
		logger.Error("Failed to list books", sl.Err(err, true))
		return nil, s.convertError(err)
	}

	response := result.ToListResponseProto()
	logger.Info("Books listed", slog.Int("count", len(response.Books)), slog.Bool("has_next_page", response.NextPageToken != ""))
	return response, nil
}
