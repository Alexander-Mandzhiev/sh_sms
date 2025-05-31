package classes_service

import (
	"backend/pkg/models/library"
	"context"
	"errors"
	"log/slog"
)

func (s *Service) Get(ctx context.Context, id int) (*library_models.Class, error) {
	const op = "service.Library.Classes.Get"
	logger := s.logger.With(slog.String("op", op), slog.Int("id", id))
	logger.Debug("processing request")

	if id <= 0 {
		logger.Warn("invalid ID received")
		return nil, errors.New("invalid class ID")
	}

	class, err := s.provider.GetClassByID(ctx, id)
	if err != nil {
		logger.Error("failed to get class", slog.Any("error", err))
		return nil, err
	}

	if class.Grade < 1 || class.Grade > 11 {
		logger.Warn("invalid grade in class", slog.Int("grade", int(class.Grade)))
		return nil, errors.New("invalid class grade")
	}

	logger.Debug("class retrieved")
	return class, nil
}
