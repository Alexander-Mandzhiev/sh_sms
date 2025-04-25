package main

import (
	config "backend/pkg/config/service"
	"backend/pkg/dbManager"
	sl "backend/pkg/logger"
	"backend/pkg/server/grpc_server"
	userHandler "backend/service/sso/users/handle"
	userRepository "backend/service/sso/users/repository"
	userService "backend/service/sso/users/service"

	roleHandler "backend/service/sso/roles/handle"
	roleRepository "backend/service/sso/roles/repository"
	roleService "backend/service/sso/roles/service"

	permissionsHandler "backend/service/sso/permissions/handle"
	permissionsRepository "backend/service/sso/permissions/repository"
	permissionsService "backend/service/sso/permissions/service"
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
	reposU, err := userRepository.New(dbPool, logger)
	if err != nil {
		logger.Error("Failed to create user repository", sl.Err(err, true))
		return
	}
	reposR, err := roleRepository.New(dbPool, logger)
	if err != nil {
		logger.Error("Failed to create user repository", sl.Err(err, true))
		return
	}

	reposP, err := permissionsRepository.New(dbPool, logger)
	if err != nil {
		logger.Error("Failed to create permissions repository", sl.Err(err, true))
		return
	}

	// 5. Инициализация сервиса
	srvU := userService.New(reposU, logger)
	srvR := roleService.New(reposR, logger)
	srvP := permissionsService.New(reposP, logger)

	// 6. Инициализация gRPC сервера
	app := grpc_server.New()

	// 7. Регистрация сервиса в gRPC сервере
	userHandler.Register(app.GRPCServer, srvU, logger)
	roleHandler.Register(app.GRPCServer, srvR, logger)
	permissionsHandler.Register(app.GRPCServer, srvP, logger)

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
