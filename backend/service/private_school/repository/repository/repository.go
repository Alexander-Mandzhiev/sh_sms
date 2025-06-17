package students_repository

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
)

type Repository struct {
	db     *pgxpool.Pool
	logger *slog.Logger
}

func New(db *pgxpool.Pool, logger *slog.Logger) (*Repository, error) {
	op := "repository.New.PrivateSchool.StudentsRepository"
	if db == nil {
		logger.Error("Database connection is nil", slog.String("op", op))
		return nil, fmt.Errorf("database connection is nil")
	}
	logger.Info("Repository initialized", slog.String("op", op))
	return &Repository{db: db, logger: logger}, nil
}
