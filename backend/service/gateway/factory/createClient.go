package factory

import (
	"backend/protos/gen/go/apps/app_manager"
	"backend/protos/gen/go/apps/clients_apps"
	"backend/protos/gen/go/apps/secrets"
	"backend/protos/gen/go/auth"
	"backend/protos/gen/go/clients/addresses"
	"backend/protos/gen/go/clients/client_types"
	"backend/protos/gen/go/clients/clients"
	"backend/protos/gen/go/clients/contacts"
	library "backend/protos/gen/go/library"
	"backend/protos/gen/go/sso/permissions"
	"backend/protos/gen/go/sso/role_permissions"
	"backend/protos/gen/go/sso/roles"
	"backend/protos/gen/go/sso/users"
	"backend/protos/gen/go/sso/users_roles"
	"google.golang.org/grpc"
)

func (p *ClientProvider) createClient(serviceType ServiceType, conn *grpc.ClientConn) interface{} {
	switch serviceType {
	case ServiceSSO:
		return &SSOClient{
			UserServiceClient:           users.NewUserServiceClient(conn),
			RoleServiceClient:           roles.NewRoleServiceClient(conn),
			PermissionServiceClient:     permissions.NewPermissionServiceClient(conn),
			RolePermissionServiceClient: role_permissions.NewRolePermissionServiceClient(conn),
			UserRoleServiceClient:       user_roles.NewUserRoleServiceClient(conn),
			conn:                        conn,
		}
	case ServiceApps:
		return &AppsClient{
			AppServiceClient:        app_manager.NewAppServiceClient(conn),
			ClientsAppServiceClient: client_apps.NewClientsAppServiceClient(conn),
			SecretServiceClient:     secrets.NewSecretServiceClient(conn),
			conn:                    conn,
		}
	case ServiceClients:
		return &ClientsClient{
			ClientServiceClient:     clients.NewClientServiceClient(conn),
			ClientTypeServiceClient: client_types.NewClientTypeServiceClient(conn),
			AddressServiceClient:    addresses.NewAddressServiceClient(conn),
			ContactServiceClient:    contacts.NewContactServiceClient(conn),
			conn:                    conn,
		}
	case ServiceAuth:
		return &AuthClient{
			AuthServiceClient: auth.NewAuthServiceClient(conn),
			conn:              conn,
		}
	case ServiceLibrary:
		return &LibraryClient{
			AttachmentServiceClient: library.NewAttachmentServiceClient(conn),
			SubjectServiceClient:    library.NewSubjectServiceClient(conn),
			FileFormatServiceClient: library.NewFileFormatServiceClient(conn),
			BookServiceClient:       library.NewBookServiceClient(conn),
			ClassServiceClient:      library.NewClassServiceClient(conn),
			conn:                    conn,
		}
	default:
		panic("unknown service type")
	}
	return nil
}
