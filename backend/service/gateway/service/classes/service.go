package classes_service

import (
	library "backend/protos/gen/go/library"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"log/slog"
)

type ClassesService interface {
	GetClass(ctx context.Context, req *library.GetClassRequest) (*library.Class, error)
	ListClasses(ctx context.Context) (*library.ListClassesResponse, error)
}

type classesService struct {
	client library.ClassServiceClient
	logger *slog.Logger
}

func NewClassesService(client library.ClassServiceClient, logger *slog.Logger) ClassesService {
	return &classesService{
		client: client,
		logger: logger.With("service", "classes"),
	}
}

func (s *classesService) GetClass(ctx context.Context, req *library.GetClassRequest) (*library.Class, error) {
	s.logger.Debug("getting class", "class_id", req.Id)
	return s.client.GetClass(ctx, req)
}

func (s *classesService) ListClasses(ctx context.Context) (*library.ListClassesResponse, error) {
	s.logger.Debug("listing classes")
	return s.client.ListClasses(ctx, &emptypb.Empty{})
}
