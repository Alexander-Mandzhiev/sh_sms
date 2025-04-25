package handle

import (
	"backend/protos/gen/go/sso/users"
	"backend/service/sso/models"
	"context"
	"errors"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log/slog"
)

var (
	ErrInvalidArgument  = errors.New("invalid argument")
	ErrPermissionDenied = errors.New("permission denied")
	ErrAlreadyExists    = errors.New("user already exists")
	ErrNotFound         = errors.New("user not found")
	ErrUnauthenticated  = errors.New("unauthenticated")
)

type UserService interface {
	Create(ctx context.Context, user *models.User, password string) error
	Get(ctx context.Context, clientID, userID uuid.UUID) (*models.User, error)
	Update(ctx context.Context, user *models.User) (*models.User, error)
	Delete(ctx context.Context, clientID, userID uuid.UUID, permanent bool) error
	List(ctx context.Context, req models.ListRequest) ([]models.User, int, error)
	SetPassword(ctx context.Context, clientID, userID uuid.UUID, password string) error
	Restore(ctx context.Context, clientID, userID uuid.UUID) (*models.User, error)
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
