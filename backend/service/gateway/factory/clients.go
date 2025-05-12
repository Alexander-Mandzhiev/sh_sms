package factory

import (
	"backend/protos/gen/go/clients/addresses"
	"backend/protos/gen/go/clients/client_types"
	"backend/protos/gen/go/clients/clients"
	"backend/protos/gen/go/clients/contacts"
	"context"
	"google.golang.org/grpc"
)

type ClientsClientType interface {
	clients.ClientServiceClient
	client_types.ClientTypeServiceClient
	addresses.AddressServiceClient
	contacts.ContactServiceClient
	Close() error
}

type ClientsClient struct {
	clients.ClientServiceClient
	client_types.ClientTypeServiceClient
	addresses.AddressServiceClient
	contacts.ContactServiceClient
	conn *grpc.ClientConn
}

func (c *ClientsClient) Close() error {
	return c.conn.Close()
}

func (p *ClientProvider) GetClientsClient(ctx context.Context) (ClientsClientType, error) {
	client, err := p.getClient(ctx, ServiceClients)
	if err != nil {
		return nil, err
	}
	return client.(ClientsClientType), nil
}
