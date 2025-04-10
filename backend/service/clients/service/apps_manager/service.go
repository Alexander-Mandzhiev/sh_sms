package apps_manager_service

import (
	"backend/protos/gen/go/apps/app_manager"
	"context"
	"errors"
	"log/slog"
)

var (
	// Валидация
	ErrInvalidID          = errors.New("invalid ID format")
	ErrEmptyCode          = errors.New("empty code")
	ErrInvalidPagination  = errors.New("invalid pagination")
	ErrConflictParams     = errors.New("conflicting parameters")
	ErrNoUpdateFields     = errors.New("no update fields")
	ErrIdentifierRequired = errors.New("identifier required")
	ErrInvalidName        = errors.New("invalid name")
	ErrMaxCountExceeded   = errors.New("max count exceeded")

	// Бизнес-логика
	ErrAlreadyExists = errors.New("already exists")
	ErrNotFound      = errors.New("not found")
	ErrInactiveApp   = errors.New("inactive application")

	// Системные
	ErrDatabaseConnection = errors.New("database connection failed")
	ErrTimeout            = errors.New("request timeout")
)

type AppsProvider interface {
	Create(ctx context.Context, app *app_manager.App) error
	Get(ctx context.Context, req *app_manager.GetRequest) (*app_manager.App, error)
	Update(ctx context.Context, app *app_manager.App) error
	Delete(ctx context.Context, id int32) error
	List(ctx context.Context, req *app_manager.ListRequest) ([]*app_manager.App, int32, error)
}

type Service struct {
	logger   *slog.Logger
	provider AppsProvider
}

func New(provider AppsProvider, logger *slog.Logger) *Service {
	const op = "service.New"

	if logger == nil {
		logger = slog.Default()
	}

	logger.Info("initializing apps service", slog.String("op", op))
	return &Service{
		provider: provider,
		logger:   logger,
	}
}
