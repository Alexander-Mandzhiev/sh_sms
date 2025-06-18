package groups_repository

import (
	"backend/pkg/models/groups"
	"backend/pkg/utils"
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *Repository) CreateGroup(ctx context.Context, s *groups_models.CreateGroup) (*groups_models.Group, error) {
	const op = "repository.PrivateSchool.GroupsRepository.CreateGroup"
	logger := r.logger.With(slog.String("op", op), slog.String("client_id", s.ClientID.String()), slog.String("name", s.Name), slog.String("trace_id", utils.TraceIDFromContext(ctx)))
	logger.Debug("creating group in repository")

	newPublicID := uuid.New()

	query := `INSERT INTO groups (public_id, client_id, name, curator_id) VALUES ($1, $2, $3, $4)
		RETURNING internal_id, public_id, client_id, name, curator_id, created_at, updated_at`

	if ctx.Err() != nil {
		logger.Warn("context canceled before query execution")
		return nil, ctx.Err()
	}

	var group groups_models.Group
	var curatorID *uuid.UUID

	err := r.db.QueryRow(ctx, query, newPublicID, s.ClientID, s.Name, s.CuratorID).Scan(&group.InternalID, &group.PublicID, &group.ClientID, &group.Name, &curatorID, &group.CreatedAt, &group.UpdatedAt)
	group.CuratorID = curatorID
	if err != nil {
		return nil, r.handleCreateError(err, query, logger)
	}

	logger.Info("group created", "public_id", group.PublicID, "internal_id", group.InternalID)
	return &group, nil
}

func (r *Repository) handleCreateError(err error, query string, logger *slog.Logger) error {
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		logger.Warn("context error", "error", err)
		return err
	}
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == "23505" && pgErr.ConstraintName == "groups_client_id_name_key" {
			logger.Warn("duplicate group name detected", "detail", pgErr.Detail)
			return groups_models.ErrDuplicateGroupName
		}
		if pgErr.Code == "23503" && pgErr.ConstraintName == "groups_curator_id_fkey" {
			logger.Warn("invalid curator reference", "detail", pgErr.Detail)
			return groups_models.ErrInvalidCuratorID
		}
		logger.Warn("database constraint violation", "code", pgErr.Code, "constraint", pgErr.ConstraintName, "detail", pgErr.Detail)
		return fmt.Errorf("%w: %s", groups_models.ErrCreateFailed, pgErr.Message)
	}

	if errors.Is(err, pgx.ErrNoRows) {
		logger.Warn("no group created after insert")
		return groups_models.ErrCreateFailed
	}

	logger.Error("database operation failed", "error", err, "query", query)
	return fmt.Errorf("%w: %v", groups_models.ErrCreateFailed, err)
}
