package books_handle

import (
	"backend/pkg/models/library"
	library "backend/protos/gen/go/library"
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log/slog"
)

type BookService interface {
	CreateBook(ctx context.Context, params *library_models.CreateBookParams) (*library_models.Book, error)
	GetBook(ctx context.Context, id int64, clientID uuid.UUID) (*library_models.Book, error)
	UpdateBook(ctx context.Context, params *library_models.UpdateBookParams) (*library_models.Book, error)
	DeleteBook(ctx context.Context, id int64, clientID uuid.UUID) error
	ListBooks(ctx context.Context, params *library_models.ListBooksParams) (*library_models.ListBooksResult, error)
}

type serverAPI struct {
	library.UnimplementedBookServiceServer
	service BookService
	logger  *slog.Logger
}

func Register(gRPCServer *grpc.Server, service BookService, logger *slog.Logger) {
	library.RegisterBookServiceServer(gRPCServer, &serverAPI{
		service: service,
		logger:  logger.With("module", "books"),
	})
}
