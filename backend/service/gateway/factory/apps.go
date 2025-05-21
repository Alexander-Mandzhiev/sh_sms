package factory

import (
	"backend/protos/gen/go/apps/app_manager"
	"backend/protos/gen/go/apps/clients_apps"
	"backend/protos/gen/go/apps/secrets"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

type AppsClientType interface {
	app_manager.AppServiceClient
	client_apps.ClientsAppServiceClient
	secrets.SecretServiceClient
	Close() error
}

type AppsClient struct {
	app_manager.AppServiceClient
	client_apps.ClientsAppServiceClient
	secrets.SecretServiceClient
	conn *grpc.ClientConn
}

func (c *AppsClient) Close() error {
	return c.conn.Close()
}

func (p *ClientProvider) GetAppsClient(ctx context.Context) (AppsClientType, error) {
	client, err := p.getClient(ctx, ServiceApps)
	if err != nil {
		return nil, err
	}
	appsClient, ok := client.(AppsClientType)
	if !ok {
		return nil, fmt.Errorf("type assertion failed for SSO client")
	}
	return appsClient, nil
}
