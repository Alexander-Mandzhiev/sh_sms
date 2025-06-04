package books_handle

import (
	"backend/pkg/models/library"
	library "backend/protos/gen/go/library"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type BooksService interface {
	CreateBook(ctx context.Context, params *library_models.CreateBookParams) (*library.BookResponse, error)
	GetBook(ctx context.Context, id int64, clientID string) (*library.BookResponse, error)
	UpdateBook(ctx context.Context, params *library_models.UpdateBookRequest) (*library.BookResponse, error)
	DeleteBook(ctx context.Context, id int64, clientID string) error
	ListBooks(ctx context.Context, req *library.ListBooksRequest) (*library.ListBooksResponse, error)
}

type Handler struct {
	logger  *slog.Logger
	service BooksService
}

func New(service BooksService, logger *slog.Logger) *Handler {
	return &Handler{service: service, logger: logger}
}

func (h *Handler) InitRoutes(router *gin.RouterGroup) {
	books := router.Group("/books")
	{
		books.POST("/", h.create)
		books.GET("/:id", h.get)
		books.PUT("/:id", h.update)
		books.DELETE("/:id", h.delete)
		books.GET("/", h.list)
	}
}
