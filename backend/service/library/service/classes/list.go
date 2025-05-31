package classes_service

import (
	"backend/pkg/models/library"
	"context"
	"log/slog"
)

func (s *Service) List(ctx context.Context) ([]*library_models.Class, error) {
	const op = "service.Library.Classes.List"
	logger := s.logger.With(slog.String("op", op))
	logger.Debug("processing request")

	classes, err := s.provider.ListClasses(ctx)
	if err != nil {
		logger.Error("failed to list classes", slog.Any("error", err))
		return nil, err
	}

	logger.Debug("classes retrieved", slog.Int("count", len(classes)))
	return classes, nil
}
