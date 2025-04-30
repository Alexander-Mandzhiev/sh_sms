package service

import (
	"context"
	"fmt"
	"log/slog"
)

func (s *Service) Delete(ctx context.Context, id int, permanent bool) error {
	const op = "service.ClientType.Delete"
	logger := s.logger.With(slog.String("op", op), slog.Int("id", id), slog.Bool("permanent", permanent))
	logger.Debug("processing delete request")

	if id <= 0 {
		err := fmt.Errorf("%w: invalid ID format", ErrInvalidArgument)
		logger.Warn("validation failed", slog.Any("error", err))
		return err
	}

	if err := s.provider.Delete(ctx, id, permanent); err != nil {
		logger.Error("delete operation failed", slog.Any("error", err), slog.Bool("permanent", permanent))
		return fmt.Errorf("%w: %v", ErrInternal, err)
	}

	logger.Info("client type deleted successfully", slog.Bool("permanent", permanent))
	return nil
}
