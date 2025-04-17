package repository

import (
	"backend/service/apps/constants"
	"backend/service/apps/models"
	"context"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func (r *Repository) ListRotations(ctx context.Context, filter models.ListFilter) ([]*models.RotationHistory, int, error) {
	const op = "repository.Secret.ListRotations"
	logger := r.logger.With(slog.String("op", op), slog.Int("page", filter.Page), slog.Int("count", filter.Count))
	baseQuery := `SELECT client_id, app_id, secret_type, old_secret, new_secret, rotated_by, rotated_at FROM secret_rotation_history WHERE 1=1`

	var filters []string
	var args []interface{}
	argNum := 1

	if filter.ClientID != nil {
		filters = append(filters, fmt.Sprintf("AND client_id = $%d", argNum))
		args = append(args, *filter.ClientID)
		argNum++
	}

	if filter.AppID != nil {
		filters = append(filters, fmt.Sprintf("AND app_id = $%d", argNum))
		args = append(args, *filter.AppID)
		argNum++
	}

	if filter.SecretType != nil {
		filters = append(filters, fmt.Sprintf("AND secret_type = $%d", argNum))
		args = append(args, *filter.SecretType)
		argNum++
	}

	if filter.RotatedBy != nil {
		filters = append(filters, fmt.Sprintf("AND rotated_by = $%d", argNum))
		args = append(args, *filter.RotatedBy)
		argNum++
	}

	if filter.RotatedAfter != nil {
		filters = append(filters, fmt.Sprintf("AND rotated_at >= $%d", argNum))
		args = append(args, *filter.RotatedAfter)
		argNum++
	}

	if filter.RotatedBefore != nil {
		filters = append(filters, fmt.Sprintf("AND rotated_at <= $%d", argNum))
		args = append(args, *filter.RotatedBefore)
		argNum++
	}

	fullQuery := baseQuery + " " + strings.Join(filters, " ") + " ORDER BY rotated_at DESC LIMIT $%d OFFSET $%d"
	limit := filter.Count
	offset := (filter.Page - 1) * filter.Count
	args = append(args, limit, offset)
	fullQuery = fmt.Sprintf(fullQuery, argNum, argNum+1)
	logger.Debug("executing query", slog.String("query", fullQuery), slog.Any("args", args))
	rows, err := r.db.Query(ctx, fullQuery, args...)
	if err != nil {
		logger.Error("query failed", slog.Any("error", err))
		return nil, 0, fmt.Errorf("%w: query failed", constants.ErrInternal)
	}
	defer rows.Close()

	var rotations []*models.RotationHistory
	for rows.Next() {
		var rh models.RotationHistory
		var rotatedBy pgtype.Text
		var rotatedAt time.Time

		if err = rows.Scan(&rh.ClientID, &rh.AppID, &rh.SecretType, &rh.OldSecret, &rh.NewSecret, &rotatedBy, &rotatedAt); err != nil {
			logger.Error("scan failed", slog.Any("error", err))
			return nil, 0, fmt.Errorf("%w: scan failed", constants.ErrInternal)
		}

		rh.RotatedAt = rotatedAt
		if rotatedBy.Valid {
			rh.RotatedBy = rotatedBy.String
		}

		rotations = append(rotations, &rh)
	}

	countQuery := "SELECT COUNT(*) FROM secret_rotation_history WHERE 1=1 " + strings.Join(filters, " ")
	var total int
	err = r.db.QueryRow(ctx, countQuery, args[:len(args)-2]...).Scan(&total)
	if err != nil {
		logger.Error("count query failed", slog.Any("error", err))
		return nil, 0, fmt.Errorf("%w: count query failed", constants.ErrInternal)
	}

	logger.Debug("rotations listed successfully", slog.Int("total", total), slog.Int("returned", len(rotations)))
	return rotations, total, nil
}
