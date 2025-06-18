package groups_repository

import (
	groups_models "backend/pkg/models/groups"
	"backend/pkg/utils"
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *Repository) UpdateGroup(ctx context.Context, ug *groups_models.UpdateGroup) (*groups_models.Group, error) {
	const op = "repository.PrivateSchool.GroupsRepository.UpdateGroup"
	logger := r.logger.With(slog.String("op", op), slog.String("public_id", ug.PublicID.String()), slog.String("client_id", ug.ClientID.String()), slog.String("trace_id", utils.TraceIDFromContext(ctx)))
	logger.Debug("updating group in repository")

	if ctx.Err() != nil {
		logger.Warn("context canceled before query execution")
		return nil, ctx.Err()
	}

	query := `UPDATE groups SET name = $3, curator_id = $4, updated_at = NOW() WHERE public_id = $1 AND client_id = $2
		RETURNING internal_id, public_id, client_id, name, curator_id, created_at, updated_at`
	var group groups_models.Group
	var curatorID *uuid.UUID

	err := r.db.QueryRow(ctx, query, ug.PublicID, ug.ClientID, ug.Name, ug.CuratorID).Scan(&group.InternalID, &group.PublicID, &group.ClientID, &group.Name, &curatorID, &group.CreatedAt, &group.UpdatedAt)
	group.CuratorID = curatorID
	if err != nil {
		return nil, r.handleUpdateError(err, query, logger)
	}

	logger.Info("group updated", "public_id", group.PublicID, "internal_id", group.InternalID)
	return &group, nil
}

func (r *Repository) handleUpdateError(err error, query string, logger *slog.Logger) error {
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		logger.Warn("context error", "error", err)
		return err
	}

	if errors.Is(err, pgx.ErrNoRows) {
		logger.Warn("group not found for update")
		return groups_models.ErrGroupNotFound
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
		return fmt.Errorf("%w: %s", groups_models.ErrUpdateFailed, pgErr.Message)
	}

	logger.Error("database operation failed", "error", err, "query", query)
	return fmt.Errorf("%w: %v", groups_models.ErrUpdateFailed, err)
}
