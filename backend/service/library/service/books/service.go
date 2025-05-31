package books_service

import (
	"backend/pkg/models/library"
	"backend/service/library/service/classes"
	"backend/service/library/service/subjects"
	"context"
	"log/slog"
)

type BooksProvider interface {
	CreateBook(ctx context.Context, book *library_models.Book) (*library_models.Book, error)
	GetBookByID(ctx context.Context, id int64) (*library_models.Book, error)
	UpdateBook(ctx context.Context, book *library_models.Book) (*library_models.Book, error)
	DeleteBook(ctx context.Context, id int64) error
	ListBooks(ctx context.Context, params *library_models.ListBooksParams) ([]*library_models.Book, error)
	CountBooks(ctx context.Context, clientID string, filter string) (int32, error)
}

type Service struct {
	logger          *slog.Logger
	provider        BooksProvider
	classesProvider classes_service.ClassesProvider
	subjectsProvide subjects_service.SubjectsProvider
}

func New(provider BooksProvider, logger *slog.Logger) *Service {
	const op = "service.New.Library.Books"

	if logger == nil {
		logger = slog.Default()
	}

	logger.Info("initializing library handle - service books", slog.String("op", op))
	return &Service{
		provider: provider,
		logger:   logger,
	}
}
