package repository

import (
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
)

var (
	ErrAlreadyExists = errors.New("entity already exists")
	ErrInternal      = errors.New("internal error")
	ErrNotFound      = errors.New("entity not found")
)

type Repository struct {
	db     *pgxpool.Pool
	logger *slog.Logger
}

func New(db *pgxpool.Pool, logger *slog.Logger) (*Repository, error) {
	if db == nil {
		logger.Error("Database connection is nil", slog.String("op", "repository.New"))
		return nil, fmt.Errorf("database connection is nil")
	}
	logger.Info("Repository initialized", slog.String("op", "repository.New"))
	return &Repository{db: db, logger: logger}, nil
}
