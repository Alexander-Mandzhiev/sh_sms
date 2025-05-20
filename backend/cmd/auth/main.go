package main

import (
	config "backend/pkg/config/auth"
	"backend/pkg/dbManager"
	sl "backend/pkg/logger"
	"backend/pkg/server/grpc_client"
	"backend/pkg/server/grpc_server"
	"backend/service/auth/factory"
	"backend/service/auth/handle"
	"backend/service/auth/repository"
	"backend/service/auth/service"
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

	sessionRepo, err := repository.New(dbPool, logger)
	if err != nil {
		logger.Error("Failed to create session repository", sl.Err(err, true))
		return
	}

	grpcManager := grpc_client.NewGRPCClientManager(logger)
	defer func() {
		if err = grpcManager.CloseAll(); err != nil {
			logger.Error("Failed to close GRPC connections", sl.Err(err, false))
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	serviceMap := map[factory.ServiceType]string{
		factory.ServiceSSO:     cfg.Services.SSOAddr,
		factory.ServiceApps:    cfg.Services.AppsAddr,
		factory.ServiceClients: cfg.Services.ClientsAddr,
	}

	clientFactory := factory.New(grpcManager, serviceMap, logger)
	defer clientFactory.Close()

	appsClient, err := clientFactory.GetAppsClient(ctx)
	if err != nil {
		logger.Error("Failed to get app client", sl.Err(err, false))
		return
	}

	ssoClient, err := clientFactory.GetSSOClient(ctx)
	if err != nil {
		logger.Error("Failed to get SSO client", sl.Err(err, false))
		return
	}

	clientsClient, err := clientFactory.GetClientsClient(ctx)
	if err != nil {
		logger.Error("Failed to get SSO client", sl.Err(err, false))
		return
	}

	usersSRV := service.NewUserService(ssoClient, logger)
	rolesSRV := service.NewRoleService(ssoClient, logger)
	permissionSRV := service.NewPermissionService(ssoClient, logger)
	rolePermissionSRV := service.NewRolePermissionService(ssoClient, logger)
	userRoleSRV := service.NewUserRoleService(ssoClient, logger)

	appManagerSRV := service.NewAppService(appsClient, logger)
	secretSRV := service.NewSecretService(appsClient, logger)
	clientAppsSRV := service.NewClientAppsService(appsClient, logger)
	clientSRV := service.NewClientService(clientsClient, logger)

	srv := service.NewAuthService(usersSRV, rolesSRV, permissionSRV, rolePermissionSRV, userRoleSRV, appManagerSRV, secretSRV, clientAppsSRV, clientSRV, sessionRepo, logger, cfg)
	app := grpc_server.New()
	handle.Register(app.GRPCServer, srv, logger, cfg)
	go func() {
		app.MustRun(logger, cfg.GRPCServer.Port)
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	logger.Info("Shutting down gracefully...")
	app.Shutdown()
	logger.Info("Service stopped")
}
