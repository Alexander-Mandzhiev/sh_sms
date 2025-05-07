package main

import (
	config "backend/pkg/config/service"
	"backend/pkg/dbManager"
	sl "backend/pkg/logger"
	"backend/pkg/server/grpc_server"
	handleClientTypes "backend/service/clients/client_types/handle"
	repoClientTypes "backend/service/clients/client_types/repository"
	srvClientTypes "backend/service/clients/client_types/service"

	handleClients "backend/service/clients/clients/handle"
	repoClients "backend/service/clients/clients/repository"
	srvClients "backend/service/clients/clients/service"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 1. Загрузка конфигурации
	cfg := config.Initialize()
	// 2. Инициализация логгера
	logger := sl.SetupLogger(cfg.Env)
	logger.Info("configuration loaded")

	// 3. Инициализация подключения к базе данных
	dbPool, err := dbManager.OpenPostgresqlConnection(cfg.DBConfig.Postgres, logger)
	if err != nil {
		logger.Error("failed to connect to database", sl.Err(err, false))
		os.Exit(1)
	}
	defer func() {
		if err = dbManager.ClosePostgresqlConnection(dbPool, logger); err != nil {
			logger.Error("failed to close database connection", sl.Err(err, false))
		}
	}()

	// 4. Инициализация репозитория
	repoCT, err := repoClientTypes.New(dbPool, logger)
	if err != nil {
		logger.Error("Failed to create repository", sl.Err(err, true))
		return
	}
	repoC, err := repoClients.New(dbPool, logger)
	if err != nil {
		logger.Error("Failed to create repository", sl.Err(err, true))
		return
	}

	// 5. Инициализация сервиса
	srvS := srvClientTypes.New(repoCT, logger)
	srvC := srvClients.New(repoC, logger)

	// 6. Инициализация gRPC сервера
	app := grpc_server.New()

	// 7. Регистрация сервиса в gRPC сервере
	handleClientTypes.Register(app.GRPCServer, srvS, logger)
	handleClients.Register(app.GRPCServer, srvC, logger)
	// 8. Запуск gRPC сервера
	go func() {
		app.MustRun(logger, cfg.GRPCServer.Port)
	}()

	// Ожидание сигнала для graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	// Graceful shutdown
	logger.Info("Shutting down gracefully...")
	app.Shutdown()
	logger.Info("Service stopped")
}
