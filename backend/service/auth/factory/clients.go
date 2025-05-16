package factory

import (
	"backend/protos/gen/go/clients/clients"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

type ClientsType interface {
	clients.ClientServiceClient
	Close() error
}

type ClientsClient struct {
	clients.ClientServiceClient
	conn *grpc.ClientConn
}

func (c *ClientsClient) Close() error {
	return c.conn.Close()
}

func (p *ClientProvider) GetClientsClient(ctx context.Context) (ClientsType, error) {
	client, err := p.getClient(ctx, ServiceClients)
	if err != nil {
		return nil, err
	}
	clientsClient, ok := client.(ClientsType)
	if !ok {
		return nil, fmt.Errorf("type assertion failed for SSO client")
	}
	return clientsClient, nil
}
