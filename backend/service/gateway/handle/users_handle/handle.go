package users_handle

import (
	"backend/protos/gen/go/sso/users"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type UsersService interface {
	CreateUser(ctx context.Context, req *users.CreateRequest) (*users.User, error)
	GetUser(ctx context.Context, req *users.GetRequest) (*users.User, error)
	UpdateUser(ctx context.Context, req *users.UpdateRequest) (*users.User, error)
	DeleteUser(ctx context.Context, req *users.DeleteRequest) (*users.SuccessResponse, error)
	ListUsers(ctx context.Context, req *users.ListRequest) (*users.ListResponse, error)
	SetPassword(ctx context.Context, req *users.SetPasswordRequest) (*users.SuccessResponse, error)
	RestoreUser(ctx context.Context, req *users.RestoreRequest) (*users.User, error)
}

type Handler struct {
	logger  *slog.Logger
	service UsersService
}

func New(service UsersService, logger *slog.Logger) *Handler {
	return &Handler{service: service, logger: logger}
}

func (h *Handler) InitRoutes(router *gin.RouterGroup) {
	user := router.Group("/users")
	{
		user.POST("", h.createUser)
		user.GET("/:id", h.getUser)
		user.PUT("/:id", h.updateUser)
		user.DELETE("/:id", h.deleteUser)
		user.POST("/list", h.listUsers)
		user.POST("/:id/password", h.setUserPassword)
		user.GET("/:id/restore", h.restoreUser)
	}
}
