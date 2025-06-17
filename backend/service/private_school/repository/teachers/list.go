package teachers_repository

import (
	"backend/pkg/models/teacher"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) ListTeachers(ctx context.Context, filter *teachers_models.ListTeachersFilter) (*teachers_models.ListTeachersResponse, error) {
	const op = "repository.PrivateSchool.Teachers.ListTeachers"
	logger := r.logger.With(slog.String("op", op))
	logger.Debug("Listing teachers in repository", "client_id", filter.ClientID, "limit", filter.Limit, "include_deleted", filter.IncludeDeleted)

	query := `SELECT id, client_id, full_name, phone, email, additional_info, deleted_at, created_at, updated_at FROM teachers WHERE client_id = $1`

	args := []interface{}{filter.ClientID}
	paramCount := 2

	if filter.Cursor != nil && filter.Cursor.LastID != "" && !filter.Cursor.CreatedAt.IsZero() {
		query += fmt.Sprintf(" AND (created_at, id) > ($%d, $%d)", paramCount, paramCount+1)
		args = append(args, filter.Cursor.CreatedAt, filter.Cursor.LastID)
		paramCount += 2
	}

	if !filter.IncludeDeleted {
		query += " AND deleted_at IS NULL"
	}

	query += fmt.Sprintf(" ORDER BY created_at ASC, id ASC LIMIT $%d", paramCount)
	args = append(args, filter.Limit+1)
	logger.Debug("Executing query", "query", query, "args", args)

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		logger.Error("failed to query teachers", "error", err)
		return nil, fmt.Errorf("failed to query teachers: %w", err)
	}
	defer rows.Close()

	teachers := make([]*teachers_models.Teacher, 0, filter.Limit)
	for rows.Next() {
		var t teachers_models.Teacher
		err = rows.Scan(&t.ID, &t.ClientID, &t.FullName, &t.Phone, &t.Email, &t.AdditionalInfo, &t.DeletedAt, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			logger.Error("failed to scan teacher", "error", err)
			return nil, fmt.Errorf("failed to scan teacher: %w", err)
		}
		teachers = append(teachers, &t)
	}

	if err = rows.Err(); err != nil {
		logger.Error("rows iteration error", "error", err)
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	var nextCursor *teachers_models.Cursor
	if len(teachers) > int(filter.Limit) {
		last := teachers[len(teachers)-1]
		teachers = teachers[:len(teachers)-1]

		nextCursor = &teachers_models.Cursor{
			LastID:    last.ID.String(),
			CreatedAt: last.CreatedAt,
		}
		logger.Debug("Next cursor prepared", "last_id", nextCursor.LastID, "created_at", nextCursor.CreatedAt)
	}

	response := &teachers_models.ListTeachersResponse{
		Teachers:   teachers,
		NextCursor: nextCursor,
	}

	logger.Info("teachers listed successfully", "count", len(teachers), "has_next", nextCursor != nil)
	return response, nil
}
