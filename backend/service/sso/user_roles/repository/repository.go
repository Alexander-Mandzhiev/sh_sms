package repository

import (
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
)

var (
	ErrInternal           = errors.New("internal server error")
	ErrInvalidArgument    = errors.New("invalid argument")
	ErrConflict           = errors.New("error conflict")
	ErrAssignmentNotFound = errors.New("role assignment not found")
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
