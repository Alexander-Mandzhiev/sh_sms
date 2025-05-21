package factory

import (
	"backend/pkg/server/grpc_client"
	"log/slog"
	"sync"
)

type ServiceType int

const (
	ServiceSSO ServiceType = iota
	ServiceApps
	ServiceClients
	ServiceAuth
)

func (s ServiceType) String() string {
	names := [...]string{"sso", "apps", "clients"}
	if int(s) < len(names) {
		return names[s]
	}
	return "unknown"
}

type ClientProvider struct {
	manager      *grpc_client.GRPCClientManager
	clients      map[ServiceType]interface{}
	clientsMutex sync.RWMutex
	logger       *slog.Logger
	serviceMap   map[ServiceType]string
}

func New(manager *grpc_client.GRPCClientManager, serviceMap map[ServiceType]string, logger *slog.Logger) *ClientProvider {
	return &ClientProvider{
		manager:    manager,
		clients:    make(map[ServiceType]interface{}),
		logger:     logger,
		serviceMap: serviceMap,
	}
}
