package factory

import (
	"backend/protos/gen/go/apps/app_manager"
	"backend/protos/gen/go/apps/clients_apps"
	"backend/protos/gen/go/apps/secrets"
	"backend/protos/gen/go/clients/clients"
	"backend/protos/gen/go/sso/permissions"
	"backend/protos/gen/go/sso/role_permissions"
	"backend/protos/gen/go/sso/roles"
	"backend/protos/gen/go/sso/users"
	"backend/protos/gen/go/sso/users_roles"
	"fmt"
	"google.golang.org/grpc"
)

func (p *ClientProvider) createClient(serviceType ServiceType, conn *grpc.ClientConn) (interface{}, error) {
	switch serviceType {
	case ServiceSSO:
		return &SSOClient{
			UserServiceClient:           users.NewUserServiceClient(conn),
			RoleServiceClient:           roles.NewRoleServiceClient(conn),
			PermissionServiceClient:     permissions.NewPermissionServiceClient(conn),
			RolePermissionServiceClient: role_permissions.NewRolePermissionServiceClient(conn),
			UserRoleServiceClient:       user_roles.NewUserRoleServiceClient(conn),
			conn:                        conn,
		}, nil
	case ServiceApps:
		return &AppsClient{
			AppServiceClient:        app_manager.NewAppServiceClient(conn),
			ClientsAppServiceClient: client_apps.NewClientsAppServiceClient(conn),
			SecretServiceClient:     secrets.NewSecretServiceClient(conn),
			conn:                    conn,
		}, nil
	case ServiceClients:
		return &ClientsClient{
			ClientServiceClient: clients.NewClientServiceClient(conn),
			conn:                conn,
		}, nil
	default:
		return nil, fmt.Errorf("unknown service type: %d", serviceType)
	}
}
