package main

import (
	config "backend/pkg/config/gateway"
	sl "backend/pkg/logger"
	"backend/pkg/server/grpc_client"
	"backend/pkg/server/http_server"
	"backend/service/gateway/factory"
	"backend/service/gateway/handle"
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := config.MustLoad()
	logger := sl.SetupLogger(cfg.Env)
	logger.Info("configuration loaded")

	grpcManager := grpc_client.NewGRPCClientManager(logger)
	defer func() {
		if err := grpcManager.CloseAll(); err != nil {
			logger.Error("Failed to close GRPC connections", sl.Err(err, false))
		}
	}()

	serviceMap := map[factory.ServiceType]string{
		factory.ServiceSSO:     cfg.Services.SSOAddr,
		factory.ServiceApps:    cfg.Services.AppsAddr,
		factory.ServiceClients: cfg.Services.ClientsAddr,
	}

	clientFactory := factory.New(grpcManager, serviceMap, logger)
	defer clientFactory.Close()

	handler := handle.New(logger, cfg.MediaDir, cfg.Env, cfg.Frontend.Addr)

	server := http_server.New(handler, logger)

	go func() {
		if err := server.Start(cfg.HTTPServer); err != nil {
			logger.Error("HTTP server error", sl.Err(err, false))
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer shutdownCancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		logger.Error("Server shutdown error", sl.Err(err, false))
	}

	logger.Info("Server exited gracefully")
}
