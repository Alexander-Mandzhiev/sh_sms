package handle

import (
	"backend/protos/gen/go/apps"
	"context"
	"google.golang.org/grpc"
)

type AppsService interface {
	Create(ctx context.Context, request *apps.CreateRequest) (*apps.App, error)
	Get(ctx context.Context, request *apps.GetRequest) (*apps.App, error)
	Update(ctx context.Context, request *apps.UpdateRequest) (*apps.App, error)
	Delete(ctx context.Context, request *apps.DeleteRequest) (*apps.DeleteResponse, error)
	List(ctx context.Context, request *apps.ListRequest) (*apps.ListResponse, error)

	GenerateSecretKey(ctx context.Context, request *apps.GenerateSecretKeyRequest) (*apps.SecretKeyResponse, error)
	RotateSecretKey(ctx context.Context, request *apps.RotateSecretKeyRequest) (*apps.SecretKeyResponse, error)
	RevokeSecretKey(ctx context.Context, request *apps.RevokeSecretKeyRequest) (*apps.RevokeSecretKeyResponse, error)
	GetKeyRotationHistory(ctx context.Context, request *apps.GetKeyRotationHistoryRequest) (*apps.KeyRotationHistoryResponse, error)
}

type serverAPI struct {
	apps.UnimplementedAppsServiceServer
	service AppsService
}

func Register(gRPCServer *grpc.Server, service AppsService) {
	apps.RegisterAppsServiceServer(gRPCServer, &serverAPI{service: service})
}
