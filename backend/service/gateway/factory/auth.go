package factory

import (
	"backend/protos/gen/go/auth"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

type AuthClientType interface {
	auth.AuthServiceClient
	Close() error
}

type AuthClient struct {
	auth.AuthServiceClient
	conn *grpc.ClientConn
}

func (c *AuthClient) Close() error {
	return c.conn.Close()
}

func (p *ClientProvider) GetAuthClient(ctx context.Context) (AuthClientType, error) {
	client, err := p.getClient(ctx, ServiceSSO)
	if err != nil {
		return nil, err
	}
	authClient, ok := client.(AuthClientType)
	if !ok {
		return nil, fmt.Errorf("type assertion failed for auth client")
	}
	return authClient, nil
}
