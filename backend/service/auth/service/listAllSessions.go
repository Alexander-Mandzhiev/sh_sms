package service

import (
	"backend/pkg/utils"
	"backend/protos/gen/go/sso/users"
	"backend/service/auth/handle"
	"backend/service/auth/models"
	"context"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
)

func (s *AuthService) ListAllSessions(ctx context.Context, filter models.AllSessionsFilter) ([]models.Session, error) {
	const op = "service.ListAllSessions"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", filter.ClientID.String()), slog.Int("app_id", filter.AppID), slog.Int("page", filter.Page), slog.Int("count", filter.Count))
	logger.Debug("starting global sessions fetch")

	if err := utils.ValidatePagination(filter.Page, filter.Count); err != nil {
		logger.Warn("invalid page", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err := utils.ValidateUUID(filter.ClientID); err != nil {
		logger.Warn("invalid client_id", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err := utils.ValidateAppID(filter.AppID); err != nil {
		logger.Warn("invalid app_id", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if filter.Email != nil {
		if err := utils.ValidateEmail(*filter.Email); err != nil {
			logger.Warn("invalid email filter", slog.Any("error", err))
			return nil, fmt.Errorf("%s: %w", op, handle.ErrInvalidEmail)
		}
	}

	sessions, err := s.session.ListAllSessions(ctx, filter)
	if err != nil {
		logger.Error("failed to list all sessions", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	userIDs := make([]string, 0, len(sessions))
	seen := make(map[uuid.UUID]bool)
	for _, session := range sessions {
		if !seen[session.UserID] {
			userIDs = append(userIDs, session.UserID.String())
			seen[session.UserID] = true
		}
	}

	usersResp, err := s.users.BatchGetUsers(ctx, &users.BatchGetRequest{UserIds: userIDs, ClientId: filter.ClientID.String()})
	if err != nil {
		logger.Error("failed to get users data", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	userMap := make(map[string]*users.User)
	for _, u := range usersResp.Users {
		userMap[u.Id] = u
	}

	enrichedSessions := make([]models.Session, 0, len(sessions))
	for _, session := range sessions {
		if user, exists := userMap[session.UserID.String()]; exists {
			session.FullName = user.FullName
			session.Phone = user.Phone
			session.Email = user.Email
		} else {
			logger.Warn("user not found for session", slog.String("user_id", session.UserID.String()), slog.String("session_id", session.SessionID.String()))
		}
		enrichedSessions = append(enrichedSessions, session)
	}

	logger.Info("global sessions fetch completed", slog.Int("total_found", len(enrichedSessions)), slog.Int("users_found", len(userMap)))

	return enrichedSessions, nil

}
