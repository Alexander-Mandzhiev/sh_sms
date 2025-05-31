package books_handle

import (
	sl "backend/pkg/logger"
	"backend/pkg/models/library"
	library "backend/protos/gen/go/library"
	"context"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
	"log/slog"
)

func (s *serverAPI) DeleteBook(ctx context.Context, req *library.DeleteBookRequest) (*emptypb.Empty, error) {
	const op = "grpc.Library.Books.DeleteBook"
	logger := s.logger.With(slog.String("op", op), slog.Int64("id", req.GetId()))
	logger.Debug("Delete book called")

	if req.GetId() <= 0 {
		logger.Warn("Invalid book ID", slog.Int64("id", req.GetId()))
		return nil, s.convertError(library_models.ErrInvalidID)
	}

	clientID, err := uuid.Parse(req.GetClientId())
	if err != nil {
		logger.Warn("Invalid client ID format", slog.String("client_id", req.GetClientId()))
		return nil, s.convertError(library_models.ErrBookInvalidClientID)
	}

	err = s.service.DeleteBook(ctx, req.GetId(), clientID)
	if err != nil {
		logger.Error("Failed to delete book", sl.Err(err, true))
		return nil, s.convertError(err)
	}

	logger.Info("Book deleted", slog.Int64("id", req.GetId()))
	return &emptypb.Empty{}, nil
}
