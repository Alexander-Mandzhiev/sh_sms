package factory

import (
	"backend/protos/gen/go/sso/permissions"
	"backend/protos/gen/go/sso/role_permissions"
	"backend/protos/gen/go/sso/roles"
	"backend/protos/gen/go/sso/users"
	"backend/protos/gen/go/sso/users_roles"
	"context"
	"google.golang.org/grpc"
)

type SSOClientType interface {
	users.UserServiceClient
	roles.RoleServiceClient
	permissions.PermissionServiceClient
	role_permissions.RolePermissionServiceClient
	user_roles.UserRoleServiceClient
	Close() error
}

type SSOClient struct {
	users.UserServiceClient
	roles.RoleServiceClient
	permissions.PermissionServiceClient
	role_permissions.RolePermissionServiceClient
	user_roles.UserRoleServiceClient
	conn *grpc.ClientConn
}

func (c *SSOClient) Close() error {
	return c.conn.Close()
}

func (p *ClientProvider) GetSSOClient(ctx context.Context) (SSOClientType, error) {
	client, err := p.getClient(ctx, ServiceSSO)
	if err != nil {
		return nil, err
	}
	return client.(SSOClientType), nil
}
