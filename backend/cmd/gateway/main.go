package main

import (
	config "backend/pkg/config/gateway"
	sl "backend/pkg/logger"
	"backend/pkg/server/http_server"
	"net/http"
	"os"
)

func main() {
	// 1. Загрузка конфигурации
	cfg := config.MustLoad()
	// 2. Инициализация логгера
	logger := sl.SetupLogger(cfg.Env)
	logger.Info("configuration loaded")

	server := http_server.New(serverAPI, logger)

	go func() {
		if err := server.Start(cfg.Frontend.Addr, cfg.HTTPServer); err != nil && err != http.ErrServerClosed {
			logger.Error("HTTP server error", sl.Err(err, false))
			os.Exit(1)
		}
	}()

}
