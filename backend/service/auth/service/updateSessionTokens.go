package service

import (
	"backend/pkg/jwt_manager"
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"
)

func (s *AuthService) UpdateSessionTokens(ctx context.Context, sessionID uuid.UUID, accessToken string, refreshToken string) error {
	const op = "auth.service.session.UpdateSessionTokens"
	accessHash := jwt_manager.HashToken(accessToken)
	refreshHash := jwt_manager.HashToken(refreshToken)
	expiresAt := time.Now().Add(s.cfg.JWT.AccessDuration)

	if err := s.session.UpdateTokens(ctx, sessionID, accessHash, refreshHash, expiresAt); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
