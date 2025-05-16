package service

import (
	"backend/pkg/jwt_manager"
	"backend/protos/gen/go/apps/secrets"
	"context"
	"fmt"
	"github.com/google/uuid"
)

func (s *AuthService) getJWTSecret(ctx context.Context, clientID uuid.UUID, appID int, secretType jwt_manager.TokenType) (string, error) {
	secretResp, err := s.secret.GetSecret(ctx, &secrets.GetRequest{
		ClientId:   clientID.String(),
		AppId:      int32(appID),
		SecretType: string(secretType),
	})
	if err != nil {
		return "", fmt.Errorf("failed to get JWT secret: %w", err)
	}
	return secretResp.CurrentSecret, nil
}
