package http_server

import (
	cfg "backend/pkg/config/gateway"
	"backend/service/gateway/handle"
	"context"
	"fmt"
	"github.com/rs/cors"
	"log/slog"
	"net/http"
	"time"
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

func (s *APIServer) Start(frontendAddr string, httpServerCfg cfg.HTTPServer) error {
	if frontendAddr == "" {
		s.logger.Error("Frontend address is not configured")
		return fmt.Errorf("frontend address is not configured")
	}

	allowedOrigins := []string{
		"http://localhost:*", "https://localhost:*",
		fmt.Sprintf("http://%s", frontendAddr),
		fmt.Sprintf("https://%s", frontendAddr),
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "X-Requested-With", "X-API-Key", "X-Csrf-Token"},
	})

	handlerWithCORS := c.Handler(s.gateway.GetHTTPHandler())

	s.httpserver = &http.Server{
		Addr:           fmt.Sprintf("%s:%d", httpServerCfg.Address, httpServerCfg.Port),
		Handler:        handlerWithCORS,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    httpServerCfg.Timeout,
		WriteTimeout:   httpServerCfg.Timeout,
		IdleTimeout:    httpServerCfg.IdleTimeout,
	}

	s.logger.Info("Starting HTTP server", slog.String("address", s.httpserver.Addr))

	if err := s.httpserver.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		s.logger.Error("Failed to start HTTP server", slog.Any("error", err))
		return err
	}

	return nil
}

func (s *APIServer) Shutdown(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	s.logger.Info("Shutting down HTTP server...")
	err := s.httpserver.Shutdown(ctx)
	if err != nil {
		s.logger.Error("HTTP server shutdown error", slog.Any("error", err))
	}
	return err
}
