package repository

import (
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
)

var (
	ErrInternal = errors.New("internal server error")
	ErrNotFound = errors.New("role not found")
	ErrConflict = errors.New("conflict fields")
)

type Repository struct {
	db     *pgxpool.Pool
	logger *slog.Logger
}

func New(db *pgxpool.Pool, logger *slog.Logger) (*Repository, error) {
	const op = "repository.New.Roles"
	if db == nil {
		logger.Error("Database connection is nil", slog.String("op", op))
		return nil, fmt.Errorf("database connection is nil")
	}
	logger.Info("Repository initialized", slog.String("op", op))
	return &Repository{db: db, logger: logger}, nil
}
