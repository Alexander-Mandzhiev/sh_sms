package service

import (
	"log/slog"
)

type AppsProvider interface {
}

type Service struct {
	logger   *slog.Logger
	provider AppsProvider
}

func New(appsProvider AppsProvider, logger *slog.Logger) *Service {
	const op = "service.New"
	if appsProvider == nil {
		logger.Error("apps provider is nil", slog.String("op", op))
		panic("apps provider cannot be nil")
	}
	logger.Info("service initialized", slog.String("op", op))
	return &Service{provider: appsProvider, logger: logger}
}
