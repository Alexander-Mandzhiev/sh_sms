package http_server

import (
	cfg "backend/pkg/config/gateway"
	"backend/service/gateway/handle"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
)

type APIServer struct {
	httpserver *http.Server
	gateway    *handle.ServerAPI
	logger     *slog.Logger
}

func New(handler *handle.ServerAPI, logger *slog.Logger) *APIServer {
	return &APIServer{
		gateway: handler,
		logger:  logger,
	}
}

func (s *APIServer) Start(httpServerCfg cfg.HTTPServer) error {
	s.httpserver = &http.Server{
		Addr:           fmt.Sprintf("%s:%d", httpServerCfg.Address, httpServerCfg.Port),
		Handler:        s.gateway.GetHTTPHandler(),
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    httpServerCfg.Timeout,
		WriteTimeout:   httpServerCfg.Timeout,
		IdleTimeout:    httpServerCfg.IdleTimeout,
	}

	s.logger.Info("Starting HTTP server", slog.String("address", s.httpserver.Addr), slog.String("env", s.gateway.GetEnv()))
	if err := s.httpserver.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		s.logger.Error("Failed to start HTTP server", slog.Any("error", err), slog.String("env", s.gateway.GetEnv()))
		return fmt.Errorf("server failed: %w", err)
	}
	return nil
}

func (s *APIServer) Shutdown(ctx context.Context) error {
	s.logger.Info("Shutting down HTTP server...",
		slog.String("env", s.gateway.GetEnv()),
	)

	if err := s.httpserver.Shutdown(ctx); err != nil {
		if !errors.Is(err, context.DeadlineExceeded) {
			s.logger.Error("HTTP server shutdown error", slog.Any("error", err), slog.String("env", s.gateway.GetEnv()))
			return fmt.Errorf("shutdown error: %w", err)
		}
		return nil
	}
	return nil
}
