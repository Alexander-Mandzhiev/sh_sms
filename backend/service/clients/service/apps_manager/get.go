package apps_manager_service

import (
	"backend/protos/gen/go/apps/app_manager"
	"context"
	"log/slog"
)

func (s *Service) Get(ctx context.Context, req *app_manager.GetRequest) (*app_manager.App, error) {
	const op = "service.Get"
	logger := s.logger.With(slog.String("op", op))

	if err := validateGetRequest(req); err != nil {
		logger.Error("Invalid request", slog.Any("error", err))
		return nil, err
	}

	app, err := s.provider.Get(ctx, req)
	if err != nil {
		logger.Error("Failed to get app",
			slog.Any("request", req),
			slog.Any("error", err))
		return nil, err
	}

	return app, nil
}

// Валидация запроса
func validateGetRequest(req *app_manager.GetRequest) error {
	switch {
	case req.Id == nil && req.Code == nil:
		return ErrIdentifierRequired
	case req.Id != nil && *req.Id <= 0:
		return ErrInvalidID
	case req.Code != nil && *req.Code == "":
		return ErrEmptyCode
	case req.Id != nil && req.Code != nil:
		return ErrConflictParams
	}
	return nil
}
