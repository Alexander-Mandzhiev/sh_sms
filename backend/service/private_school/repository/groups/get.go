package groups_repository

import (
	"backend/pkg/models/groups"
	"backend/pkg/utils"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"log/slog"
)

func (r *Repository) GetGroup(ctx context.Context, publicID, clientID uuid.UUID) (*groups_models.Group, error) {
	const op = "repository.PrivateSchool.GroupsRepository.GetGroup"
	logger := r.logger.With(slog.String("op", op), slog.String("public_id", publicID.String()), slog.String("client_id", clientID.String()), slog.String("trace_id", utils.TraceIDFromContext(ctx)))
	logger.Debug("getting group from repository")

	if ctx.Err() != nil {
		logger.Warn("context canceled before query execution")
		return nil, ctx.Err()
	}

	query := `SELECT internal_id, public_id, client_id, name, curator_id, created_at, updated_at FROM groups WHERE public_id = $1 AND client_id = $2`
	var group groups_models.Group
	var curatorID *uuid.UUID

	err := r.db.QueryRow(ctx, query, publicID, clientID).Scan(&group.InternalID, &group.PublicID, &group.ClientID, &group.Name, &curatorID, &group.CreatedAt, &group.UpdatedAt)
	group.CuratorID = curatorID

	if err != nil {
		return nil, r.handleGetError(err, query, logger)
	}

	logger.Debug("group retrieved successfully")
	return &group, nil
}

func (r *Repository) handleGetError(err error, query string, logger *slog.Logger) error {
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		logger.Warn("context error", "error", err)
		return err
	}

	if errors.Is(err, pgx.ErrNoRows) {
		logger.Warn("group not found")
		return groups_models.ErrGroupNotFound
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		logger.Warn("database error", "code", pgErr.Code, "message", pgErr.Message, "detail", pgErr.Detail)
		return fmt.Errorf("%w: database error", groups_models.ErrGetFailed)
	}

	logger.Error("database operation failed", "error", err, "query", query)
	return fmt.Errorf("%w: %v", groups_models.ErrGetFailed, err)
}
