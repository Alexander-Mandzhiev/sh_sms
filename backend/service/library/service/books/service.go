package books_service

import (
	"backend/pkg/models/library"
	"backend/service/library/service/classes"
	"backend/service/library/service/subjects"
	"context"
	"github.com/google/uuid"
	"log/slog"
)

type BooksProvider interface {
	CreateBook(ctx context.Context, book *library_models.Book) error
	GetBookByID(ctx context.Context, id int64, clientID uuid.UUID) (*library_models.Book, error)
	UpdateBook(ctx context.Context, book *library_models.Book) (*library_models.Book, error)
	DeleteBook(ctx context.Context, id int64, clientID uuid.UUID) error
	ListBooks(ctx context.Context, params *library_models.ListBooksRequest) (books []*library_models.Book, nextCursor int64, err error)
	CountBooks(ctx context.Context, clientID uuid.UUID, filter *string) (int32, error)
}

type Service struct {
	logger          *slog.Logger
	provider        BooksProvider
	classesProvider classes_service.ClassesProvider
	subjectsProvide subjects_service.SubjectsProvider
}

func New(provider BooksProvider, logger *slog.Logger, classesProvider classes_service.ClassesProvider, subjectsProvide subjects_service.SubjectsProvider) *Service {
	const op = "service.New.Library.Books"

	if logger == nil {
		logger = slog.Default()
	}

	logger.Info("initializing library handle - service books", slog.String("op", op))
	return &Service{
		provider:        provider,
		logger:          logger,
		classesProvider: classesProvider,
		subjectsProvide: subjectsProvide,
	}
}
