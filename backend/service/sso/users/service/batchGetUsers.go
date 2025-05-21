package service

import (
	"backend/service/sso/models"
	"context"
	"github.com/google/uuid"
	"log/slog"
)

func (s *Service) BatchGetUsers(ctx context.Context, clientID uuid.UUID, userIDs []uuid.UUID, includeInactive bool) ([]*models.User, []uuid.UUID, error) {
	const op = "service.BatchGetUsers"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", clientID.String()), slog.Int("user_count", len(userIDs)))

	users, err := s.provider.BatchGetUsers(ctx, clientID, userIDs)
	if err != nil {
		logger.Error("failed to get users from provider", slog.Any("error", err))
		return nil, nil, ErrInternal
	}

	foundIDs := make(map[uuid.UUID]bool)
	for _, u := range users {
		foundIDs[u.ID] = true
	}

	var missing []uuid.UUID
	for _, id := range userIDs {
		if !foundIDs[id] {
			missing = append(missing, id)
		}
	}

	if !includeInactive {
		var activeUsers []*models.User
		for _, u := range users {
			if u.IsActive {
				activeUsers = append(activeUsers, u)
			} else {
				missing = append(missing, u.ID)
			}
		}
		users = activeUsers
	}

	logger.Debug("batch get users completed", slog.Int("found", len(users)), slog.Int("missing", len(missing)))
	return users, missing, nil
}
