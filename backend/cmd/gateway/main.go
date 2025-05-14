package main

import (
	config "backend/pkg/config/gateway"
	sl "backend/pkg/logger"
	"backend/pkg/server/grpc_client"
	"backend/pkg/server/http_server"
	"backend/service/gateway/factory"
	"backend/service/gateway/handle"
	"backend/service/gateway/handle/permissions_handle"
	"backend/service/gateway/handle/role_permissions_handle"
	"backend/service/gateway/handle/roles_handle"
	"backend/service/gateway/handle/users_handle"
	"backend/service/gateway/service/appManager"
	"backend/service/gateway/service/clientApps"
	permission_service "backend/service/gateway/service/permissions"
	role_permission_service "backend/service/gateway/service/rolePermissions"
	"backend/service/gateway/service/roles"
	"backend/service/gateway/service/secrets"
	"backend/service/gateway/service/users"
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

	_ = client_app_service.NewClientAppsService(appsClient, logger)
	_ = app_manager_service.NewAppService(appsClient, logger)
	_ = secrets_service.NewSecretService(appsClient, logger)

	usersSRV := users_service.NewUserService(ssoClient, logger)
	rolesSRV := roles_service.NewRoleService(ssoClient, logger)
	permissionSRV := permission_service.NewPermissionService(ssoClient, logger)
	rolePermissionSRV := role_permission_service.NewRolePermissionService(ssoClient, logger)

	userHandle := users_handle.New(usersSRV, logger)
	roleHandle := roles_handle.New(rolesSRV, logger)
	permissionHandle := permissions_handle.New(permissionSRV, logger)
	rolePermissionHandle := role_permissions_handle.New(rolePermissionSRV, logger)

	handler := handle.New(logger, cfg.MediaDir, cfg.Env, cfg.Frontend.Addr)

	handler.RegisterHandlers(userHandle, roleHandle, permissionHandle, rolePermissionHandle)

	server := http_server.New(handler, logger)

	go func() {
		if err = server.Start(cfg.HTTPServer); err != nil {
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

	if err = server.Shutdown(shutdownCtx); err != nil {
		logger.Error("Server shutdown error", sl.Err(err, false))
	}

	logger.Info("Server exited gracefully")
}
