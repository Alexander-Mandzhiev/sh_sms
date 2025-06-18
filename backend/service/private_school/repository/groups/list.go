package groups_repository

import (
	"backend/pkg/models/groups"
	"backend/pkg/utils"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"log/slog"
	"strings"
)

func (r *Repository) ListGroups(ctx context.Context, params *groups_models.ListGroupsRequest) (*groups_models.GroupListResponse, error) {
	const op = "repository.PrivateSchool.GroupsRepository.ListGroups"
	logger := r.logger.With(slog.String("op", op), slog.String("client_id", params.ClientID.String()), slog.Int("page_size", int(params.PageSize)), slog.String("trace_id", utils.TraceIDFromContext(ctx)))

	if params.Cursor != nil {
		logger = logger.With(slog.Int64("cursor", *params.Cursor))
	} else {
		logger = logger.With(slog.String("cursor", "nil"))
	}

	if params.NameFilter != nil {
		logger = logger.With(slog.String("name_filter", *params.NameFilter))
	}

	logger.Debug("listing groups from repository")
	baseQuery := `SELECT internal_id, public_id, client_id, name, curator_id, created_at, updated_at FROM groups WHERE client_id = $1 `
	conditions := []string{}
	queryParams := []interface{}{params.ClientID}
	paramIndex := 2

	if params.NameFilter != nil && *params.NameFilter != "" {
		conditions = append(conditions, fmt.Sprintf("name ILIKE $%d", paramIndex))
		queryParams = append(queryParams, "%"+*params.NameFilter+"%")
		paramIndex++
	}

	if params.Cursor != nil && *params.Cursor > 0 {
		conditions = append(conditions, fmt.Sprintf("internal_id < $%d", paramIndex))
		queryParams = append(queryParams, *params.Cursor)
		paramIndex++
	}

	fullQuery := baseQuery
	if len(conditions) > 0 {
		fullQuery += " AND " + strings.Join(conditions, " AND ")
	}

	fullQuery += fmt.Sprintf(" ORDER BY internal_id DESC LIMIT $%d", paramIndex)
	queryParams = append(queryParams, params.PageSize+1)

	logger.Debug("executing query", "query", fullQuery, "params", queryParams)

	if ctx.Err() != nil {
		logger.Warn("context canceled before query execution")
		return nil, ctx.Err()
	}

	rows, err := r.db.Query(ctx, fullQuery, queryParams...)
	if err != nil {
		return nil, r.handleListError(err, fullQuery, logger)
	}
	defer rows.Close()

	var groups []*groups_models.Group
	for rows.Next() {
		var group groups_models.Group
		var curatorID *uuid.UUID
		err = rows.Scan(&group.InternalID, &group.PublicID, &group.ClientID, &group.Name, &curatorID, &group.CreatedAt, &group.UpdatedAt)
		if err != nil {
			logger.Error("failed to scan group row", "error", err)
			return nil, fmt.Errorf("%w: failed to scan row", groups_models.ErrListFailed)
		}
		group.CuratorID = curatorID
		groups = append(groups, &group)
	}

	if err = rows.Err(); err != nil {
		logger.Error("rows iteration error", "error", err)
		return nil, fmt.Errorf("%w: %v", groups_models.ErrListFailed, err)
	}

	logger.Debug("groups fetched", "count", len(groups))
	nextCursor := int64(0)
	if len(groups) > int(params.PageSize) && len(groups) > 0 {
		groups = groups[:params.PageSize]
		nextCursor = groups[len(groups)-1].InternalID
		logger.Debug("next page available", "next_cursor", nextCursor)
	} else {
		logger.Debug("last page reached")
	}

	logger.Info("groups listed", "returned_count", len(groups), "next_cursor", nextCursor)
	return &groups_models.GroupListResponse{Groups: groups, NextCursor: nextCursor}, nil
}

func (r *Repository) handleListError(err error, query string, logger *slog.Logger) error {
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		logger.Warn("context error", "error", err)
		return err
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		logger.Warn("database error during list operation", "code", pgErr.Code, "message", pgErr.Message, "detail", pgErr.Detail)
		return fmt.Errorf("%w: %s", groups_models.ErrListFailed, pgErr.Message)
	}

	logger.Error("database operation failed", "error", err, "query", query)
	return fmt.Errorf("%w: %v", groups_models.ErrListFailed, err)
}
