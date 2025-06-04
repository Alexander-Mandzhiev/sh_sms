package books_service

import (
	"context"
	"github.com/google/uuid"
	"log/slog"
)

func (s *Service) DeleteBook(ctx context.Context, id int64, clientID uuid.UUID) error {
	const op = "service.Library.Books.Delete"
	logger := s.logger.With(slog.String("op", op), slog.Int64("id", id), slog.String("client_id", clientID.String()))
	logger.Debug("Getting book")

	if err := s.provider.DeleteBook(ctx, id, clientID); err != nil {
		logger.Error("Failed to get book", slog.Any("error", err))
		return err
	}

	logger.Debug("Book deleted")
	return nil
}
