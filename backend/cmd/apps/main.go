package main

import (
	config "backend/pkg/config/service"
	"backend/pkg/dbManager"
	sl "backend/pkg/logger"
	"backend/pkg/server/grpc_server"
	handleAppManager "backend/service/apps/app_manager/handle"
	repoAppManager "backend/service/apps/app_manager/repository"
	srvAppManager "backend/service/apps/app_manager/service"
	handleclientApp "backend/service/apps/client_apps/handle"
	repoClientApps "backend/service/apps/client_apps/repository"
	srvClientApps "backend/service/apps/client_apps/service"

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
	repoAM, err := repoAppManager.New(dbPool, logger)
	if err != nil {
		logger.Error("Failed to create repository", sl.Err(err, true))
		return
	}

	reposCA, err := repoClientApps.New(dbPool, logger)
	if err != nil {
		logger.Error("Failed to create repository", sl.Err(err, true))
		return
	}

	// 5. Инициализация сервиса
	srvAM := srvAppManager.New(repoAM, logger)
	srvCA := srvClientApps.New(reposCA, logger)

	// 6. Инициализация gRPC сервера
	app := grpc_server.New()

	// 7. Регистрация сервиса в gRPC сервере
	handleAppManager.Register(app.GRPCServer, srvAM, logger)
	handleclientApp.Register(app.GRPCServer, srvCA, logger)
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
