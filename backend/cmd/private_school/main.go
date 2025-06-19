package main

import (
	config "backend/pkg/config/library"
	"backend/pkg/server/grpc_server"
	"os/signal"
	"syscall"

	handleG "backend/service/private_school/handle/groups"
	repositoryG "backend/service/private_school/repository/groups"
	serviceG "backend/service/private_school/service/groups"

	"backend/pkg/dbManager"
	sl "backend/pkg/logger"
	"os"
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
	repoG, err := repositoryG.New(dbPool, logger)

	// 5. Инициализация сервиса
	srvG := serviceG.New(repoG, logger)
	// 6. Инициализация gRPC сервера
	app := grpc_server.New()

	// 7. Регистрация сервиса в gRPC сервере
	handleG.Register(app.GRPCServer, srvG, logger)

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
