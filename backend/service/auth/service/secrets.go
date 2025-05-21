package service

import (
	"backend/protos/gen/go/apps/secrets"
	"context"
	"log/slog"
)

type SecretService interface {
	GetSecret(ctx context.Context, req *secrets.GetRequest) (*secrets.Secret, error)
}

type secretService struct {
	client secrets.SecretServiceClient
	logger *slog.Logger
}

func NewSecretService(client secrets.SecretServiceClient, logger *slog.Logger) SecretService {
	return &secretService{
		client: client,
		logger: logger.With("service", "secrets"),
	}
}

func (s *secretService) GetSecret(ctx context.Context, req *secrets.GetRequest) (*secrets.Secret, error) {
	s.logger.Debug("getting secret", "client_id", req.ClientId, "app_id", req.AppId, "secret_type", req.SecretType)
	return s.client.GetSecret(ctx, req)
}
