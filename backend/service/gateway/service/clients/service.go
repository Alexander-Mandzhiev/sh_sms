package client_service

import (
	"backend/protos/gen/go/clients/clients"
	"context"
	"log/slog"
)

type ClientService interface {
	CreateClient(ctx context.Context, req *clients.CreateRequest) (*clients.Client, error)
	GetClient(ctx context.Context, req *clients.GetRequest) (*clients.Client, error)
	UpdateClient(ctx context.Context, req *clients.UpdateRequest) (*clients.Client, error)
	DeleteClient(ctx context.Context, req *clients.DeleteRequest) (*clients.DeleteResponse, error)
	ListClients(ctx context.Context, req *clients.ListRequest) (*clients.ListResponse, error)
	RestoreClient(ctx context.Context, req *clients.RestoreRequest) (*clients.Client, error)
}

type clientService struct {
	client clients.ClientServiceClient
	logger *slog.Logger
}

func NewClientService(client clients.ClientServiceClient, logger *slog.Logger) ClientService {
	return &clientService{
		client: client,
		logger: logger.With("service", "client"),
	}
}

func (s *clientService) CreateClient(ctx context.Context, req *clients.CreateRequest) (*clients.Client, error) {
	s.logger.Debug("creating client", "name", req.Name, "type_id", req.TypeId)
	return s.client.CreateClient(ctx, req)
}

func (s *clientService) GetClient(ctx context.Context, req *clients.GetRequest) (*clients.Client, error) {
	s.logger.Debug("getting client", "client_id", req.Id)
	return s.client.GetClient(ctx, req)
}

func (s *clientService) UpdateClient(ctx context.Context, req *clients.UpdateRequest) (*clients.Client, error) {
	s.logger.Debug("updating client", "client_id", req.Id, "fields_updated", getUpdatedFields(req))
	return s.client.UpdateClient(ctx, req)
}

func (s *clientService) DeleteClient(ctx context.Context, req *clients.DeleteRequest) (*clients.DeleteResponse, error) {
	s.logger.Debug("deleting client", "client_id", req.Id, "permanent", req.GetPermanent())
	return s.client.DeleteClient(ctx, req)
}

func (s *clientService) ListClients(ctx context.Context, req *clients.ListRequest) (*clients.ListResponse, error) {
	s.logger.Debug("listing clients", "page", req.Page, "count", req.Count, "search", req.GetSearch(), "type_id", req.GetTypeId())
	return s.client.ListClients(ctx, req)
}

func (s *clientService) RestoreClient(ctx context.Context, req *clients.RestoreRequest) (*clients.Client, error) {
	s.logger.Debug("restoring client", "client_id", req.Id)
	return s.client.RestoreClient(ctx, req)
}

func getUpdatedFields(req *clients.UpdateRequest) []string {
	fields := make([]string, 0)
	if req.Name != nil {
		fields = append(fields, "name")
	}
	if req.Description != nil {
		fields = append(fields, "description")
	}
	if req.TypeId != nil {
		fields = append(fields, "type_id")
	}
	if req.Website != nil {
		fields = append(fields, "website")
	}
	return fields
}
