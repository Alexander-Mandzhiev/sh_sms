package teachers_service

import (
	"backend/pkg/models/teacher"
	"context"
	"fmt"
	"log/slog"
)

func (s *Service) ListTeachers(ctx context.Context, filter *teachers_models.ListTeachersFilter) (*teachers_models.ListTeachersResponse, error) {
	const op = "teachers_service.ListTeachers"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", filter.ClientID.String()), slog.Int("limit", int(filter.Limit)))
	logger.Debug("listing teachers")

	response, err := s.provider.ListTeachers(ctx, filter)
	if err != nil {
		logger.Error("failed to list teachers", "error", err)
		return nil, fmt.Errorf("failed to list teachers: %w", err)
	}

	logger.Info("teachers listed successfully", "count", len(response.Teachers))
	return response, nil
}
