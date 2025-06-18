package students_repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
)

func (r *Repository) studentExists(ctx context.Context, id, clientID uuid.UUID) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM students WHERE id = $1 AND client_id = $2)`
	var exists bool
	err := r.db.QueryRow(ctx, query, id, clientID).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check student existence: %w", err)
	}
	return exists, nil
}
