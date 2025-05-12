// service/client_type_service.go
package service

import (
	"context"
	"log/slog"

	"backend/protos/gen/go/clients/client_types"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ClientTypeService interface {
	CreateClientType(ctx context.Context, req *client_types.CreateRequest) (*client_types.ClientType, error)
	GetClientType(ctx context.Context, req *client_types.GetRequest) (*client_types.ClientType, error)
	UpdateClientType(ctx context.Context, req *client_types.UpdateRequest) (*client_types.ClientType, error)
	ListClientTypes(ctx context.Context, req *client_types.ListRequest) (*client_types.ListResponse, error)
	DeleteClientType(ctx context.Context, req *client_types.DeleteRequest) (*emptypb.Empty, error)
	RestoreClientType(ctx context.Context, req *client_types.RestoreRequest) (*client_types.ClientType, error)
}

type clientTypeService struct {
	client client_types.ClientTypeServiceClient
	logger *slog.Logger
}

func NewClientTypeService(client client_types.ClientTypeServiceClient, logger *slog.Logger) ClientTypeService {
	return &clientTypeService{
		client: client,
		logger: logger.With("service", "client_type"),
	}
}

func (s *clientTypeService) CreateClientType(ctx context.Context, req *client_types.CreateRequest) (*client_types.ClientType, error) {
	s.logger.Debug("creating client type", "code", req.Code, "name", req.Name, "active", req.GetIsActive())
	return s.client.CreateClientType(ctx, req)
}

func (s *clientTypeService) GetClientType(ctx context.Context, req *client_types.GetRequest) (*client_types.ClientType, error) {
	s.logger.Debug("getting client type", "type_id", req.Id)
	return s.client.GetClientType(ctx, req)
}

func (s *clientTypeService) UpdateClientType(ctx context.Context, req *client_types.UpdateRequest) (*client_types.ClientType, error) {
	s.logger.Debug("updating client type", "type_id", req.Id, "code", req.Code, "name", req.Name)
	return s.client.UpdateClientType(ctx, req)
}

func (s *clientTypeService) ListClientTypes(ctx context.Context, req *client_types.ListRequest) (*client_types.ListResponse, error) {
	s.logger.Debug("listing client types", "page", req.Page, "count", req.Count, "search", req.GetSearch(), "active_only", req.GetActiveOnly())
	return s.client.ListClientType(ctx, req)
}

func (s *clientTypeService) DeleteClientType(ctx context.Context, req *client_types.DeleteRequest) (*emptypb.Empty, error) {
	s.logger.Debug("deleting client type", "type_id", req.Id, "permanent", req.GetPermanent())
	return s.client.DeleteClientType(ctx, req)
}

func (s *clientTypeService) RestoreClientType(ctx context.Context, req *client_types.RestoreRequest) (*client_types.ClientType, error) {
	s.logger.Debug("restoring client type", "type_id", req.Id)
	return s.client.RestoreClientType(ctx, req)
}
