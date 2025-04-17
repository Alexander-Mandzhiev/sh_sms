package repository

import (
	"backend/service/apps/models"
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"strings"
)

func (r *Repository) List(ctx context.Context, filter models.ListFilter) ([]*models.Secret, int, error) {
	const op = "repository.Secret.List"
	logger := r.logger.With(slog.String("op", op))
	baseQuery := `SELECT client_id, app_id, secret_type, current_secret, algorithm, secret_version, generated_at, revoked_at FROM secrets WHERE 1=1`

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

	if filter.IsActive != nil {
		if *filter.IsActive {
			filters = append(filters, "AND revoked_at IS NULL")
		} else {
			filters = append(filters, "AND revoked_at IS NOT NULL")
		}
	}

	fullQuery := baseQuery + " " + strings.Join(filters, " ") + " ORDER BY generated_at DESC LIMIT $%d OFFSET $%d"
	limit := filter.Count
	offset := (filter.Page - 1) * filter.Count
	args = append(args, limit, offset)
	fullQuery = fmt.Sprintf(fullQuery, argNum, argNum+1)

	rows, err := r.db.Query(ctx, fullQuery, args...)
	if err != nil {
		logger.Error("failed to query secrets", slog.Any("error", err))
		return nil, 0, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var secrets []*models.Secret
	for rows.Next() {
		var secret models.Secret
		var revokedAt sql.NullTime

		err = rows.Scan(&secret.ClientID, &secret.AppID, &secret.SecretType, &secret.CurrentSecret, &secret.Algorithm, &secret.SecretVersion, &secret.GeneratedAt, &revokedAt)
		if err != nil {
			logger.Error("failed to scan secret", slog.Any("error", err))
			return nil, 0, fmt.Errorf("scan failed: %w", err)
		}

		if revokedAt.Valid {
			secret.RevokedAt = &revokedAt.Time
		}

		secrets = append(secrets, &secret)
	}

	countQuery := "SELECT COUNT(*) FROM secrets WHERE 1=1 " + strings.Join(filters, " ")
	var total int
	err = r.db.QueryRow(ctx, countQuery, args[:len(args)-2]...).Scan(&total)
	if err != nil {
		logger.Error("failed to get total count", slog.Any("error", err))
		return nil, 0, fmt.Errorf("count query failed: %w", err)
	}

	return secrets, total, nil
}
