package subjects_service

import (
	"backend/pkg/models/library"
	library "backend/protos/gen/go/library"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"log/slog"
)

type SubjectsService interface {
	CreateSubject(ctx context.Context, params *library_models.CreateSubjectParams) (*library.Subject, error)
	GetSubject(ctx context.Context, id int32) (*library.Subject, error)
	UpdateSubject(ctx context.Context, params *library_models.UpdateSubjectParams) (*library.Subject, error)
	DeleteSubject(ctx context.Context, id int32) error
	ListSubjects(ctx context.Context) (*library.ListSubjectsResponse, error)
}

type subjectsService struct {
	client library.SubjectServiceClient
	logger *slog.Logger
}

func NewSubjectsService(client library.SubjectServiceClient, logger *slog.Logger) SubjectsService {
	return &subjectsService{
		client: client,
		logger: logger.With("service", "subjects"),
	}
}

func (s *subjectsService) CreateSubject(ctx context.Context, params *library_models.CreateSubjectParams) (*library.Subject, error) {
	s.logger.Debug("Creating subject", "name", params.Name)
	req := &library.CreateSubjectRequest{Name: params.Name}
	return s.client.CreateSubject(ctx, req)
}

func (s *subjectsService) GetSubject(ctx context.Context, id int32) (*library.Subject, error) {
	s.logger.Debug("Getting subject", "id", id)
	return s.client.GetSubject(ctx, &library.GetSubjectRequest{Id: id})
}

func (s *subjectsService) UpdateSubject(ctx context.Context, params *library_models.UpdateSubjectParams) (*library.Subject, error) {
	s.logger.Debug("Updating subject", "id", params.ID, "name", params.Name)
	req := &library.UpdateSubjectRequest{Id: params.ID, Name: params.Name}
	return s.client.UpdateSubject(ctx, req)
}

func (s *subjectsService) DeleteSubject(ctx context.Context, id int32) error {
	s.logger.Debug("Deleting subject", "id", id)
	_, err := s.client.DeleteSubject(ctx, &library.DeleteSubjectRequest{Id: id})
	return err
}

func (s *subjectsService) ListSubjects(ctx context.Context) (*library.ListSubjectsResponse, error) {
	s.logger.Debug("Listing subjects")
	return s.client.ListSubjects(ctx, &emptypb.Empty{})
}
