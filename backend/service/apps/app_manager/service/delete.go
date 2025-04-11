package service

import (
	"context"
	"log/slog"
)

func (s *Service) Delete(ctx context.Context, id int32) error {
	const op = "service.Delete"
	logger := s.logger.With(slog.String("op", op))

	if id <= 0 {
		logger.Error("Invalid ID")
		return ErrInvalidID
	}

	if err := s.provider.Delete(ctx, id); err != nil {
		logger.Error("Delete failed", slog.Int("id", int(id)), slog.Any("error", err))
		return err
	}

	logger.Info("App deleted", slog.Int("id", int(id)))
	return nil
}
