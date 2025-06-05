package main

import (
	config "backend/pkg/config/library"
	"backend/pkg/server/grpc_server"
	attachmentRepository "backend/service/library/repository/attachment"
	booksRepository "backend/service/library/repository/books"
	classesRepository "backend/service/library/repository/classes"
	fileFormatRepository "backend/service/library/repository/fileFormats"
	subjectsRepository "backend/service/library/repository/subjects"
	serviceA "backend/service/library/service/attachment"
	serviceB "backend/service/library/service/books"
	serviceC "backend/service/library/service/classes"
	serviceFF "backend/service/library/service/fileFormats"
	serviceS "backend/service/library/service/subjects"

	handlerA "backend/service/library/handle/attachment"
	handlerB "backend/service/library/handle/books"
	handlerC "backend/service/library/handle/classes"
	handlerFF "backend/service/library/handle/fileFormats"
	handlerS "backend/service/library/handle/subjects"

	"os/signal"
	"syscall"

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
	reposC, err := classesRepository.New(dbPool, logger)
	if err != nil {
		logger.Error("Failed to create user repository", sl.Err(err, true))
		return
	}

	reposFF, err := fileFormatRepository.New(dbPool, logger)
	if err != nil {
		logger.Error("Failed to create file format repository", sl.Err(err, true))
		return
	}

	reposS, err := subjectsRepository.New(dbPool, logger)
	if err != nil {
		logger.Error("Failed to create subjects repository", sl.Err(err, true))
		return
	}

	reposB, err := booksRepository.New(dbPool, logger)
	if err != nil {
		logger.Error("Failed to create books repository", sl.Err(err, true))
		return
	}

	reposA, err := attachmentRepository.New(dbPool, logger)

	// 5. Инициализация сервиса
	srvS := serviceS.New(reposS, logger)
	srvC := serviceC.New(reposC, logger)
	srvFF := serviceFF.New(reposFF, logger)
	srvA := serviceA.New(reposA, logger)
	srvB := serviceB.New(reposB, logger, reposC, reposS)
	// 6. Инициализация gRPC сервера
	app := grpc_server.New()

	// 7. Регистрация сервиса в gRPC сервере
	handlerS.Register(app.GRPCServer, srvS, logger)
	handlerFF.Register(app.GRPCServer, srvFF, logger)
	handlerC.Register(app.GRPCServer, srvC, logger)
	handlerA.Register(app.GRPCServer, srvA, logger)
	handlerB.Register(app.GRPCServer, srvB, logger)
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
