package repository

import (
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
)

var (
	ErrNotFound      = errors.New("client app not found")
	ErrAlreadyExists = errors.New("client app already exists")
	ErrInternal      = errors.New("internal error")
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
