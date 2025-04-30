package service

import (
	"backend/service/clients/client_types/models"
	"context"
	"errors"
	"log/slog"
)

var (
	ErrInternal        = errors.New("internal server error")
	ErrInvalidArgument = errors.New("invalid argument")
	ErrCodeExists      = errors.New("code exists")
	ErrNotFound        = errors.New("not found")
	ErrAlreadyActive   = errors.New("already active")
	ErrCodeConflict    = errors.New("code conflict")
	ErrConflict        = errors.New("conflict")
)

type ClientTypesProvider interface {
	Create(ctx context.Context, ct *models.CreateParams) (*models.ClientType, error)
	Get(ctx context.Context, id int) (*models.ClientType, error)
	Update(ctx context.Context, ct *models.UpdateParams) (*models.ClientType, error)
	Delete(ctx context.Context, id int, permanent bool) error
	List(ctx context.Context, filter models.Filter, pagination models.Pagination) ([]*models.ClientType, int, error)
	Restore(ctx context.Context, id int) (*models.ClientType, error)
	CodeExists(ctx context.Context, code string) (bool, error)
	HasDependentClients(ctx context.Context, typeID int) (bool, error)
	Exist(ctx context.Context, id int) (bool, error)
}

type Service struct {
	logger   *slog.Logger
	provider ClientTypesProvider
}

func New(provider ClientTypesProvider, logger *slog.Logger) *Service {
	const op = "service.ClientTypes.New"
	if logger == nil {
		logger = slog.Default()
	}

	logger.Info("initializing clients service - handle client_type", slog.String("op", op))
	return &Service{
		provider: provider,
		logger:   logger,
	}
}
