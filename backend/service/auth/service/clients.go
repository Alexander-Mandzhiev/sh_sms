package service

import (
	"backend/protos/gen/go/clients/clients"
	"context"
	"log/slog"
)

type ClientService interface {
	GetClient(ctx context.Context, req *clients.GetRequest) (*clients.Client, error)
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

func (s *clientService) GetClient(ctx context.Context, req *clients.GetRequest) (*clients.Client, error) {
	s.logger.Debug("getting client", "client_id", req.Id)
	return s.client.GetClient(ctx, req)
}
