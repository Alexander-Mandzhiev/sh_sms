package service

import "log/slog"

type RolesProvider interface {
}

type Service struct {
	logger   *slog.Logger
	provider RolesProvider
}

func New(provider RolesProvider, logger *slog.Logger) *Service {
	const op = "service.New"

	if logger == nil {
		logger = slog.Default()
	}

	logger.Info("initializing sso handle - service roles", slog.String("op", op))
	return &Service{
		provider: provider,
		logger:   logger,
	}
}
