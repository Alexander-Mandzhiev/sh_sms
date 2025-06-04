package main

import (
	config "backend/pkg/config/gateway"
	sl "backend/pkg/logger"
	"backend/pkg/server/grpc_client"
	"backend/pkg/server/http_server"
	"backend/service/gateway/factory"
	"backend/service/gateway/handle"
	"backend/service/gateway/handle/attachments_handle"
	"backend/service/gateway/handle/auth_handle"
	books_handle "backend/service/gateway/handle/books"
	"backend/service/gateway/handle/classes_handle"
	"backend/service/gateway/handle/file_formats_handle"
	"backend/service/gateway/handle/permissions_handle"
	"backend/service/gateway/handle/role_permissions_handle"
	"backend/service/gateway/handle/roles_handle"
	"backend/service/gateway/handle/subjects_handle"
	"backend/service/gateway/handle/user_roles_handle"
	"backend/service/gateway/handle/users_handle"
	"backend/service/gateway/service/appManager"
	attachments_service "backend/service/gateway/service/attachments"
	"backend/service/gateway/service/auth"
	books_service "backend/service/gateway/service/books"
	classes_service "backend/service/gateway/service/classes"
	"backend/service/gateway/service/clientApps"
	file_formats_service "backend/service/gateway/service/fileFormats"
	"backend/service/gateway/service/permissions"
	"backend/service/gateway/service/rolePermissions"
	"backend/service/gateway/service/roles"
	"backend/service/gateway/service/secrets"
	subjects_service "backend/service/gateway/service/subjects"
	"backend/service/gateway/service/userRoles"
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
		factory.ServiceAuth:    cfg.Services.AuthHAddr,
		factory.ServiceLibrary: cfg.Services.LibraryAddr,
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

	authClient, err := clientFactory.GetAuthClient(ctx)
	if err != nil {
		logger.Error("Failed to get SSO client", sl.Err(err, false))
		return
	}

	libraryClient, err := clientFactory.GetLibraryClient(ctx)
	if err != nil {
		logger.Error("Failed to get SSO client", sl.Err(err, false))
		return
	}

	// clients
	_ = client_app_service.NewClientAppsService(appsClient, logger)
	_ = app_manager_service.NewAppService(appsClient, logger)
	_ = secrets_service.NewSecretService(appsClient, logger)

	// sso
	usersSRV := users_service.NewUserService(ssoClient, logger)
	userHandle := users_handle.New(usersSRV, logger)

	rolesSRV := roles_service.NewRoleService(ssoClient, logger)
	roleHandle := roles_handle.New(rolesSRV, logger)

	permissionSRV := permission_service.NewPermissionService(ssoClient, logger)
	permissionHandle := permissions_handle.New(permissionSRV, logger)

	rolePermissionSRV := role_permission_service.NewRolePermissionService(ssoClient, logger)
	rolePermissionHandle := role_permissions_handle.New(rolePermissionSRV, logger)

	userRoleSRV := user_roles_service.NewUserRoleService(ssoClient, logger)
	userRoleHandle := user_roles_handle.New(userRoleSRV, logger)

	// auth
	authSRV := auth_service.NewAuthService(authClient, logger)
	authHandle := auth_handle.New(authSRV, logger)

	// library
	classesSRV := classes_service.NewClassesService(libraryClient, logger)
	classesHandle := classes_handle.New(classesSRV, logger)

	fileFormatsSRV := file_formats_service.NewFileFormatsService(libraryClient, logger)
	fileFormatsHandle := file_formats_handle.New(fileFormatsSRV, logger)

	subjectsSRV := subjects_service.NewSubjectsService(libraryClient, logger)
	subjectsHandle := subjects_handle.New(subjectsSRV, logger)

	booksSRV := books_service.NewBooksService(libraryClient, logger)
	booksHandle := books_handle.New(booksSRV, logger)

	attachmentSRV := attachments_service.NewAttachmentsService(libraryClient, logger)
	attachmentHandle := attachments_handle.New(attachmentSRV, logger)

	// server
	handler := handle.New(logger, cfg.MediaDir, cfg.Env, cfg.Frontend.Addr)
	handler.RegisterHandlers(
		userHandle, roleHandle, permissionHandle, rolePermissionHandle, userRoleHandle, authHandle,
		classesHandle, subjectsHandle, fileFormatsHandle, booksHandle, attachmentHandle)
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
