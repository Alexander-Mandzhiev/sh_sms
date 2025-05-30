// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: sso/role_permissions/role_permissions.proto

package role_permissions

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	RolePermissionService_AddPermissionsToRole_FullMethodName      = "/sso.role_permissions.RolePermissionService/AddPermissionsToRole"
	RolePermissionService_RemovePermissionsFromRole_FullMethodName = "/sso.role_permissions.RolePermissionService/RemovePermissionsFromRole"
	RolePermissionService_ListPermissionsForRole_FullMethodName    = "/sso.role_permissions.RolePermissionService/ListPermissionsForRole"
	RolePermissionService_ListRolesForPermission_FullMethodName    = "/sso.role_permissions.RolePermissionService/ListRolesForPermission"
	RolePermissionService_HasPermission_FullMethodName             = "/sso.role_permissions.RolePermissionService/HasPermission"
)

// RolePermissionServiceClient is the client API for RolePermissionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RolePermissionServiceClient interface {
	AddPermissionsToRole(ctx context.Context, in *PermissionsRequest, opts ...grpc.CallOption) (*OperationStatus, error)
	RemovePermissionsFromRole(ctx context.Context, in *PermissionsRequest, opts ...grpc.CallOption) (*OperationStatus, error)
	ListPermissionsForRole(ctx context.Context, in *ListPermissionsRequest, opts ...grpc.CallOption) (*ListPermissionsResponse, error)
	ListRolesForPermission(ctx context.Context, in *ListRolesRequest, opts ...grpc.CallOption) (*ListRolesResponse, error)
	HasPermission(ctx context.Context, in *HasPermissionRequest, opts ...grpc.CallOption) (*HasPermissionResponse, error)
}

type rolePermissionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRolePermissionServiceClient(cc grpc.ClientConnInterface) RolePermissionServiceClient {
	return &rolePermissionServiceClient{cc}
}

func (c *rolePermissionServiceClient) AddPermissionsToRole(ctx context.Context, in *PermissionsRequest, opts ...grpc.CallOption) (*OperationStatus, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OperationStatus)
	err := c.cc.Invoke(ctx, RolePermissionService_AddPermissionsToRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolePermissionServiceClient) RemovePermissionsFromRole(ctx context.Context, in *PermissionsRequest, opts ...grpc.CallOption) (*OperationStatus, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OperationStatus)
	err := c.cc.Invoke(ctx, RolePermissionService_RemovePermissionsFromRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolePermissionServiceClient) ListPermissionsForRole(ctx context.Context, in *ListPermissionsRequest, opts ...grpc.CallOption) (*ListPermissionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListPermissionsResponse)
	err := c.cc.Invoke(ctx, RolePermissionService_ListPermissionsForRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolePermissionServiceClient) ListRolesForPermission(ctx context.Context, in *ListRolesRequest, opts ...grpc.CallOption) (*ListRolesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListRolesResponse)
	err := c.cc.Invoke(ctx, RolePermissionService_ListRolesForPermission_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolePermissionServiceClient) HasPermission(ctx context.Context, in *HasPermissionRequest, opts ...grpc.CallOption) (*HasPermissionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HasPermissionResponse)
	err := c.cc.Invoke(ctx, RolePermissionService_HasPermission_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RolePermissionServiceServer is the server API for RolePermissionService service.
// All implementations must embed UnimplementedRolePermissionServiceServer
// for forward compatibility.
type RolePermissionServiceServer interface {
	AddPermissionsToRole(context.Context, *PermissionsRequest) (*OperationStatus, error)
	RemovePermissionsFromRole(context.Context, *PermissionsRequest) (*OperationStatus, error)
	ListPermissionsForRole(context.Context, *ListPermissionsRequest) (*ListPermissionsResponse, error)
	ListRolesForPermission(context.Context, *ListRolesRequest) (*ListRolesResponse, error)
	HasPermission(context.Context, *HasPermissionRequest) (*HasPermissionResponse, error)
	mustEmbedUnimplementedRolePermissionServiceServer()
}

// UnimplementedRolePermissionServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRolePermissionServiceServer struct{}

func (UnimplementedRolePermissionServiceServer) AddPermissionsToRole(context.Context, *PermissionsRequest) (*OperationStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPermissionsToRole not implemented")
}
func (UnimplementedRolePermissionServiceServer) RemovePermissionsFromRole(context.Context, *PermissionsRequest) (*OperationStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemovePermissionsFromRole not implemented")
}
func (UnimplementedRolePermissionServiceServer) ListPermissionsForRole(context.Context, *ListPermissionsRequest) (*ListPermissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPermissionsForRole not implemented")
}
func (UnimplementedRolePermissionServiceServer) ListRolesForPermission(context.Context, *ListRolesRequest) (*ListRolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRolesForPermission not implemented")
}
func (UnimplementedRolePermissionServiceServer) HasPermission(context.Context, *HasPermissionRequest) (*HasPermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HasPermission not implemented")
}
func (UnimplementedRolePermissionServiceServer) mustEmbedUnimplementedRolePermissionServiceServer() {}
func (UnimplementedRolePermissionServiceServer) testEmbeddedByValue()                               {}

// UnsafeRolePermissionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RolePermissionServiceServer will
// result in compilation errors.
type UnsafeRolePermissionServiceServer interface {
	mustEmbedUnimplementedRolePermissionServiceServer()
}

func RegisterRolePermissionServiceServer(s grpc.ServiceRegistrar, srv RolePermissionServiceServer) {
	// If the following call pancis, it indicates UnimplementedRolePermissionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RolePermissionService_ServiceDesc, srv)
}

func _RolePermissionService_AddPermissionsToRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolePermissionServiceServer).AddPermissionsToRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RolePermissionService_AddPermissionsToRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolePermissionServiceServer).AddPermissionsToRole(ctx, req.(*PermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RolePermissionService_RemovePermissionsFromRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolePermissionServiceServer).RemovePermissionsFromRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RolePermissionService_RemovePermissionsFromRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolePermissionServiceServer).RemovePermissionsFromRole(ctx, req.(*PermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RolePermissionService_ListPermissionsForRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolePermissionServiceServer).ListPermissionsForRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RolePermissionService_ListPermissionsForRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolePermissionServiceServer).ListPermissionsForRole(ctx, req.(*ListPermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RolePermissionService_ListRolesForPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolePermissionServiceServer).ListRolesForPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RolePermissionService_ListRolesForPermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolePermissionServiceServer).ListRolesForPermission(ctx, req.(*ListRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RolePermissionService_HasPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HasPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolePermissionServiceServer).HasPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RolePermissionService_HasPermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolePermissionServiceServer).HasPermission(ctx, req.(*HasPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RolePermissionService_ServiceDesc is the grpc.ServiceDesc for RolePermissionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RolePermissionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sso.role_permissions.RolePermissionService",
	HandlerType: (*RolePermissionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddPermissionsToRole",
			Handler:    _RolePermissionService_AddPermissionsToRole_Handler,
		},
		{
			MethodName: "RemovePermissionsFromRole",
			Handler:    _RolePermissionService_RemovePermissionsFromRole_Handler,
		},
		{
			MethodName: "ListPermissionsForRole",
			Handler:    _RolePermissionService_ListPermissionsForRole_Handler,
		},
		{
			MethodName: "ListRolesForPermission",
			Handler:    _RolePermissionService_ListRolesForPermission_Handler,
		},
		{
			MethodName: "HasPermission",
			Handler:    _RolePermissionService_HasPermission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sso/role_permissions/role_permissions.proto",
}
