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
	SetPassword(ctx context.Context, clientID, userID uuid.UUID, password string) error
	Restore(ctx context.Context, clientID, userID uuid.UUID) (*models.User, error)
	GetUserByLogin(ctx context.Context, login, password string, clientID uuid.UUID) (*models.UserInfo, error)
	BatchGetUsers(ctx context.Context, clientID uuid.UUID, userIDs []uuid.UUID, includeInactive bool) ([]*models.User, []uuid.UUID, error)
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
