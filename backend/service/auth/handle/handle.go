package handle

import (
	config "backend/pkg/config/auth"
	"backend/protos/gen/go/auth"
	"backend/service/auth/models"
	"context"
	"errors"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log/slog"
)

var (
	ErrInvalidArgument  = errors.New("invalid argument")
	ErrPermissionDenied = errors.New("permission denied")
	ErrUnauthenticated  = errors.New("unauthenticated")
)

type AuthService interface {
	Login(ctx context.Context, req models.AuthRequest) (*models.UserInfo, string, string, error)
	Logout(ctx context.Context, accessToken, refreshToken string) error
	RefreshToken(ctx context.Context, refreshToken string, clientID uuid.UUID, appID int) (*models.UserInfo, string, string, error)

	ValidateToken(ctx context.Context, token string, tokenTypeHint string) (*models.TokenValidationResult, error)
	IntrospectToken(ctx context.Context, token string, tokenTypeHint string) (*models.TokenIntrospection, error)
	CheckPermission(ctx context.Context, check *models.PermissionCheck) (bool, []string, []string, error)

	ListActiveSessions(ctx context.Context, filter models.SessionFilter) ([]models.Session, error)
	TerminateSession(ctx context.Context, sessionID uuid.UUID) error
}

type serverAPI struct {
	auth.UnimplementedAuthServiceServer
	service AuthService
	logger  *slog.Logger
	cfg     *config.Config
}

func Register(gRPCServer *grpc.Server, service AuthService, logger *slog.Logger, cfg *config.Config) {
	auth.RegisterAuthServiceServer(gRPCServer, &serverAPI{
		service: service,
		logger:  logger,
		cfg:     cfg,
	})
}
