// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: clients/client_types/client_types.proto

package client_types

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ClientTypeService_CreateClientType_FullMethodName  = "/clients.client_types.v1.ClientTypeService/CreateClientType"
	ClientTypeService_GetClientType_FullMethodName     = "/clients.client_types.v1.ClientTypeService/GetClientType"
	ClientTypeService_UpdateClientType_FullMethodName  = "/clients.client_types.v1.ClientTypeService/UpdateClientType"
	ClientTypeService_ListClientType_FullMethodName    = "/clients.client_types.v1.ClientTypeService/ListClientType"
	ClientTypeService_DeleteClientType_FullMethodName  = "/clients.client_types.v1.ClientTypeService/DeleteClientType"
	ClientTypeService_RestoreClientType_FullMethodName = "/clients.client_types.v1.ClientTypeService/RestoreClientType"
)

// ClientTypeServiceClient is the client API for ClientTypeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientTypeServiceClient interface {
	CreateClientType(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*ClientType, error)
	GetClientType(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*ClientType, error)
	UpdateClientType(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*ClientType, error)
	ListClientType(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	DeleteClientType(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RestoreClientType(ctx context.Context, in *RestoreRequest, opts ...grpc.CallOption) (*ClientType, error)
}

type clientTypeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientTypeServiceClient(cc grpc.ClientConnInterface) ClientTypeServiceClient {
	return &clientTypeServiceClient{cc}
}

func (c *clientTypeServiceClient) CreateClientType(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*ClientType, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ClientType)
	err := c.cc.Invoke(ctx, ClientTypeService_CreateClientType_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientTypeServiceClient) GetClientType(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*ClientType, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ClientType)
	err := c.cc.Invoke(ctx, ClientTypeService_GetClientType_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientTypeServiceClient) UpdateClientType(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*ClientType, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ClientType)
	err := c.cc.Invoke(ctx, ClientTypeService_UpdateClientType_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientTypeServiceClient) ListClientType(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, ClientTypeService_ListClientType_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientTypeServiceClient) DeleteClientType(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ClientTypeService_DeleteClientType_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientTypeServiceClient) RestoreClientType(ctx context.Context, in *RestoreRequest, opts ...grpc.CallOption) (*ClientType, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ClientType)
	err := c.cc.Invoke(ctx, ClientTypeService_RestoreClientType_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientTypeServiceServer is the server API for ClientTypeService service.
// All implementations must embed UnimplementedClientTypeServiceServer
// for forward compatibility.
type ClientTypeServiceServer interface {
	CreateClientType(context.Context, *CreateRequest) (*ClientType, error)
	GetClientType(context.Context, *GetRequest) (*ClientType, error)
	UpdateClientType(context.Context, *UpdateRequest) (*ClientType, error)
	ListClientType(context.Context, *ListRequest) (*ListResponse, error)
	DeleteClientType(context.Context, *DeleteRequest) (*emptypb.Empty, error)
	RestoreClientType(context.Context, *RestoreRequest) (*ClientType, error)
	mustEmbedUnimplementedClientTypeServiceServer()
}

// UnimplementedClientTypeServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedClientTypeServiceServer struct{}

func (UnimplementedClientTypeServiceServer) CreateClientType(context.Context, *CreateRequest) (*ClientType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClientType not implemented")
}
func (UnimplementedClientTypeServiceServer) GetClientType(context.Context, *GetRequest) (*ClientType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClientType not implemented")
}
func (UnimplementedClientTypeServiceServer) UpdateClientType(context.Context, *UpdateRequest) (*ClientType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateClientType not implemented")
}
func (UnimplementedClientTypeServiceServer) ListClientType(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListClientType not implemented")
}
func (UnimplementedClientTypeServiceServer) DeleteClientType(context.Context, *DeleteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteClientType not implemented")
}
func (UnimplementedClientTypeServiceServer) RestoreClientType(context.Context, *RestoreRequest) (*ClientType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RestoreClientType not implemented")
}
func (UnimplementedClientTypeServiceServer) mustEmbedUnimplementedClientTypeServiceServer() {}
func (UnimplementedClientTypeServiceServer) testEmbeddedByValue()                           {}

// UnsafeClientTypeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientTypeServiceServer will
// result in compilation errors.
type UnsafeClientTypeServiceServer interface {
	mustEmbedUnimplementedClientTypeServiceServer()
}

func RegisterClientTypeServiceServer(s grpc.ServiceRegistrar, srv ClientTypeServiceServer) {
	// If the following call pancis, it indicates UnimplementedClientTypeServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ClientTypeService_ServiceDesc, srv)
}

func _ClientTypeService_CreateClientType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientTypeServiceServer).CreateClientType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientTypeService_CreateClientType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientTypeServiceServer).CreateClientType(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientTypeService_GetClientType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientTypeServiceServer).GetClientType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientTypeService_GetClientType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientTypeServiceServer).GetClientType(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientTypeService_UpdateClientType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientTypeServiceServer).UpdateClientType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientTypeService_UpdateClientType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientTypeServiceServer).UpdateClientType(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientTypeService_ListClientType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientTypeServiceServer).ListClientType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientTypeService_ListClientType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientTypeServiceServer).ListClientType(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientTypeService_DeleteClientType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientTypeServiceServer).DeleteClientType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientTypeService_DeleteClientType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientTypeServiceServer).DeleteClientType(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientTypeService_RestoreClientType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientTypeServiceServer).RestoreClientType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientTypeService_RestoreClientType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientTypeServiceServer).RestoreClientType(ctx, req.(*RestoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClientTypeService_ServiceDesc is the grpc.ServiceDesc for ClientTypeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientTypeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "clients.client_types.v1.ClientTypeService",
	HandlerType: (*ClientTypeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateClientType",
			Handler:    _ClientTypeService_CreateClientType_Handler,
		},
		{
			MethodName: "GetClientType",
			Handler:    _ClientTypeService_GetClientType_Handler,
		},
		{
			MethodName: "UpdateClientType",
			Handler:    _ClientTypeService_UpdateClientType_Handler,
		},
		{
			MethodName: "ListClientType",
			Handler:    _ClientTypeService_ListClientType_Handler,
		},
		{
			MethodName: "DeleteClientType",
			Handler:    _ClientTypeService_DeleteClientType_Handler,
		},
		{
			MethodName: "RestoreClientType",
			Handler:    _ClientTypeService_RestoreClientType_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "clients/client_types/client_types.proto",
}
