package repository

import (
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
)

var (
	ErrInternal      = errors.New("internal server error")
	ErrCodeExists    = errors.New("code exists")
	ErrNotFound      = errors.New("not found")
	ErrConflict      = errors.New("conflict")
	ErrCodeConflict  = errors.New("code conflict")
	ErrAlreadyActive = errors.New("already active")
)

type Repository struct {
	db     *pgxpool.Pool
	logger *slog.Logger
}

func New(db *pgxpool.Pool, logger *slog.Logger) (*Repository, error) {
	op := "repository.New.ClientTypes"
	if db == nil {
		logger.Error("Database connection is nil", slog.String("op", op))
		return nil, fmt.Errorf("database connection is nil")
	}
	logger.Info("Repository initialized", slog.String("op", op))
	return &Repository{db: db, logger: logger}, nil
}
