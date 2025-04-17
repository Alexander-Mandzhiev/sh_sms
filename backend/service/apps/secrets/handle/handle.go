package handle

import (
	pb "backend/protos/gen/go/apps/secrets"
	"backend/service/apps/models"
	"context"
	"google.golang.org/grpc"
	"log/slog"
	"time"
)

type SecretService interface {
	Generate(ctx context.Context, params models.CreateSecretParams) (*models.Secret, error)
	Get(ctx context.Context, clientID string, appID int, secretType string) (*models.Secret, error)
	Rotate(ctx context.Context, params models.RotateSecretParams) (*models.Secret, error)
	Revoke(ctx context.Context, clientID string, appID int, secretType string) (*models.Secret, error)
	Delete(ctx context.Context, clientID string, appID int, secretType string) error
	List(ctx context.Context, filter models.ListFilter) ([]*models.Secret, int, error)
	GetRotation(ctx context.Context, clientID string, appID int, secretType string, rotatedAt time.Time) (*models.RotationHistory, error)
	ListRotations(ctx context.Context, filter models.ListFilter) ([]*models.RotationHistory, int, error)
}

type serverAPI struct {
	pb.UnimplementedSecretServiceServer
	service SecretService
	logger  *slog.Logger
}

func Register(gRPCServer *grpc.Server, service SecretService, logger *slog.Logger) {
	pb.RegisterSecretServiceServer(gRPCServer, &serverAPI{
		service: service,
		logger:  logger,
	})
}
