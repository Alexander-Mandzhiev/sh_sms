package handle

import (
	"backend/protos/gen/go/sso/users"
	"backend/service/sso/models"
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log/slog"
)

type UserService interface {
	Create(ctx context.Context, user *models.User, password string) error
	Get(ctx context.Context, clientID, userID uuid.UUID) (*models.User, error)
	Update(ctx context.Context, user *models.User) (*models.User, error)
	Delete(ctx context.Context, clientID, userID uuid.UUID, permanent bool) error
	List(ctx context.Context, req models.ListRequest) ([]models.User, int, error)
	SetPassword(ctx context.Context, userID uuid.UUID, passwordHash string) error
}

type serverAPI struct {
	users.UnimplementedUserServiceServer
	service UserService
	logger  *slog.Logger
}

func Register(gRPCServer *grpc.Server, service UserService, logger *slog.Logger) {
	users.RegisterUserServiceServer(gRPCServer, &serverAPI{
		service: service,
		logger:  logger,
	})
}
