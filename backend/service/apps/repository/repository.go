package repository

import (
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
)

var (
	ErrAppNotFound   = errors.New("app not found")
	ErrNilConnection = errors.New("database connection is nil")
)

type Repository struct {
	db     *pgxpool.Pool
	logger *slog.Logger
}

func New(db *pgxpool.Pool, logger *slog.Logger) *Repository {
	return &Repository{
		db:     db,
		logger: logger.With(slog.String("component", "repository")),
	}
}
