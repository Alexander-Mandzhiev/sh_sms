package service

import (
	"context"
	"log/slog"
	"time"

	"backend/protos/gen/go/apps/secrets"
)

type SecretService interface {
	GenerateSecret(ctx context.Context, req *secrets.CreateRequest) (*secrets.Secret, error)
	GetSecret(ctx context.Context, req *secrets.GetRequest) (*secrets.Secret, error)
	RotateSecret(ctx context.Context, req *secrets.RotateRequest) (*secrets.Secret, error)
	RevokeSecret(ctx context.Context, req *secrets.RevokeRequest) (*secrets.Secret, error)
	DeleteSecret(ctx context.Context, req *secrets.DeleteRequest) (*secrets.DeleteResponse, error)
	ListSecrets(ctx context.Context, req *secrets.ListRequest) (*secrets.ListResponse, error)
	GetRotationHistory(ctx context.Context, req *secrets.GetRotationHistoryRequest) (*secrets.RotationHistory, error)
	ListRotationHistory(ctx context.Context, req *secrets.ListRequest) (*secrets.ListRotationHistoryResponse, error)
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

func (s *secretService) GenerateSecret(ctx context.Context, req *secrets.CreateRequest) (*secrets.Secret, error) {
	s.logger.Debug("generating new secret", "client_id", req.ClientId, "app_id", req.AppId, "secret_type", req.SecretType, "algorithm", req.GetAlgorithm())
	return s.client.Generate(ctx, req)
}

func (s *secretService) GetSecret(ctx context.Context, req *secrets.GetRequest) (*secrets.Secret, error) {
	s.logger.Debug("getting secret", "client_id", req.ClientId, "app_id", req.AppId, "secret_type", req.SecretType)
	return s.client.GetSecret(ctx, req)
}

func (s *secretService) RotateSecret(ctx context.Context, req *secrets.RotateRequest) (*secrets.Secret, error) {
	start := time.Now()
	defer func() {
		s.logger.Info("secret rotation completed", "duration", time.Since(start), "client_id", req.ClientId, "app_id", req.AppId)
	}()

	s.logger.Debug("rotating secret", "client_id", req.ClientId, "app_id", req.AppId, "secret_type", req.SecretType, "rotated_by", req.RotatedBy)
	return s.client.RotateSecret(ctx, req)
}

func (s *secretService) RevokeSecret(ctx context.Context, req *secrets.RevokeRequest) (*secrets.Secret, error) {
	s.logger.Debug("revoking secret", "client_id", req.ClientId, "app_id", req.AppId, "secret_type", req.SecretType)
	return s.client.RevokeSecret(ctx, req)
}

func (s *secretService) DeleteSecret(ctx context.Context, req *secrets.DeleteRequest) (*secrets.DeleteResponse, error) {
	s.logger.Warn("deleting secret", "client_id", req.ClientId, "app_id", req.AppId, "secret_type", req.SecretType)
	return s.client.DeleteSecret(ctx, req)
}

func (s *secretService) ListSecrets(ctx context.Context, req *secrets.ListRequest) (*secrets.ListResponse, error) {
	s.logger.Debug("listing secrets", "filter_client_id", req.Filter.GetClientId(), "filter_app_id", req.Filter.GetAppId(), "filter_type", req.Filter.GetSecretType(), "page", req.Page, "count", req.Count)
	return s.client.ListSecrets(ctx, req)
}

func (s *secretService) GetRotationHistory(ctx context.Context, req *secrets.GetRotationHistoryRequest) (*secrets.RotationHistory, error) {
	s.logger.Debug("getting rotation history", "history_id", req.Id)
	return s.client.GetRotation(ctx, req)
}

func (s *secretService) ListRotationHistory(ctx context.Context, req *secrets.ListRequest) (*secrets.ListRotationHistoryResponse, error) {
	s.logger.Debug("listing rotation history", "filter_rotated_by", req.Filter.GetRotatedBy(), "filter_after", req.Filter.GetRotatedAfter(), "filter_before", req.Filter.GetRotatedBefore(), "page", req.Page, "count", req.Count)
	return s.client.ListRotations(ctx, req)
}
