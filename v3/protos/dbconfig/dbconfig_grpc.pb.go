// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: dbconfig/dbconfig.proto

package dbconfigpb

import (
	context "context"
	general "github.com/hobbyfarm/gargantua/v3/protos/general"
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
	DynamicBindConfigSvc_CreateDynamicBindConfig_FullMethodName           = "/dbconfig.DynamicBindConfigSvc/CreateDynamicBindConfig"
	DynamicBindConfigSvc_GetDynamicBindConfig_FullMethodName              = "/dbconfig.DynamicBindConfigSvc/GetDynamicBindConfig"
	DynamicBindConfigSvc_UpdateDynamicBindConfig_FullMethodName           = "/dbconfig.DynamicBindConfigSvc/UpdateDynamicBindConfig"
	DynamicBindConfigSvc_DeleteDynamicBindConfig_FullMethodName           = "/dbconfig.DynamicBindConfigSvc/DeleteDynamicBindConfig"
	DynamicBindConfigSvc_DeleteCollectionDynamicBindConfig_FullMethodName = "/dbconfig.DynamicBindConfigSvc/DeleteCollectionDynamicBindConfig"
	DynamicBindConfigSvc_ListDynamicBindConfig_FullMethodName             = "/dbconfig.DynamicBindConfigSvc/ListDynamicBindConfig"
)

// DynamicBindConfigSvcClient is the client API for DynamicBindConfigSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DynamicBindConfigSvcClient interface {
	CreateDynamicBindConfig(ctx context.Context, in *CreateDynamicBindConfigRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetDynamicBindConfig(ctx context.Context, in *general.GetRequest, opts ...grpc.CallOption) (*DynamicBindConfig, error)
	UpdateDynamicBindConfig(ctx context.Context, in *UpdateDynamicBindConfigRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteDynamicBindConfig(ctx context.Context, in *general.ResourceId, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteCollectionDynamicBindConfig(ctx context.Context, in *general.ListOptions, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListDynamicBindConfig(ctx context.Context, in *general.ListOptions, opts ...grpc.CallOption) (*ListDynamicBindConfigsResponse, error)
}

type dynamicBindConfigSvcClient struct {
	cc grpc.ClientConnInterface
}

func NewDynamicBindConfigSvcClient(cc grpc.ClientConnInterface) DynamicBindConfigSvcClient {
	return &dynamicBindConfigSvcClient{cc}
}

func (c *dynamicBindConfigSvcClient) CreateDynamicBindConfig(ctx context.Context, in *CreateDynamicBindConfigRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DynamicBindConfigSvc_CreateDynamicBindConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dynamicBindConfigSvcClient) GetDynamicBindConfig(ctx context.Context, in *general.GetRequest, opts ...grpc.CallOption) (*DynamicBindConfig, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DynamicBindConfig)
	err := c.cc.Invoke(ctx, DynamicBindConfigSvc_GetDynamicBindConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dynamicBindConfigSvcClient) UpdateDynamicBindConfig(ctx context.Context, in *UpdateDynamicBindConfigRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DynamicBindConfigSvc_UpdateDynamicBindConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dynamicBindConfigSvcClient) DeleteDynamicBindConfig(ctx context.Context, in *general.ResourceId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DynamicBindConfigSvc_DeleteDynamicBindConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dynamicBindConfigSvcClient) DeleteCollectionDynamicBindConfig(ctx context.Context, in *general.ListOptions, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DynamicBindConfigSvc_DeleteCollectionDynamicBindConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dynamicBindConfigSvcClient) ListDynamicBindConfig(ctx context.Context, in *general.ListOptions, opts ...grpc.CallOption) (*ListDynamicBindConfigsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListDynamicBindConfigsResponse)
	err := c.cc.Invoke(ctx, DynamicBindConfigSvc_ListDynamicBindConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DynamicBindConfigSvcServer is the server API for DynamicBindConfigSvc service.
// All implementations must embed UnimplementedDynamicBindConfigSvcServer
// for forward compatibility.
type DynamicBindConfigSvcServer interface {
	CreateDynamicBindConfig(context.Context, *CreateDynamicBindConfigRequest) (*emptypb.Empty, error)
	GetDynamicBindConfig(context.Context, *general.GetRequest) (*DynamicBindConfig, error)
	UpdateDynamicBindConfig(context.Context, *UpdateDynamicBindConfigRequest) (*emptypb.Empty, error)
	DeleteDynamicBindConfig(context.Context, *general.ResourceId) (*emptypb.Empty, error)
	DeleteCollectionDynamicBindConfig(context.Context, *general.ListOptions) (*emptypb.Empty, error)
	ListDynamicBindConfig(context.Context, *general.ListOptions) (*ListDynamicBindConfigsResponse, error)
	mustEmbedUnimplementedDynamicBindConfigSvcServer()
}

// UnimplementedDynamicBindConfigSvcServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDynamicBindConfigSvcServer struct{}

func (UnimplementedDynamicBindConfigSvcServer) CreateDynamicBindConfig(context.Context, *CreateDynamicBindConfigRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDynamicBindConfig not implemented")
}
func (UnimplementedDynamicBindConfigSvcServer) GetDynamicBindConfig(context.Context, *general.GetRequest) (*DynamicBindConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDynamicBindConfig not implemented")
}
func (UnimplementedDynamicBindConfigSvcServer) UpdateDynamicBindConfig(context.Context, *UpdateDynamicBindConfigRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDynamicBindConfig not implemented")
}
func (UnimplementedDynamicBindConfigSvcServer) DeleteDynamicBindConfig(context.Context, *general.ResourceId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDynamicBindConfig not implemented")
}
func (UnimplementedDynamicBindConfigSvcServer) DeleteCollectionDynamicBindConfig(context.Context, *general.ListOptions) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCollectionDynamicBindConfig not implemented")
}
func (UnimplementedDynamicBindConfigSvcServer) ListDynamicBindConfig(context.Context, *general.ListOptions) (*ListDynamicBindConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDynamicBindConfig not implemented")
}
func (UnimplementedDynamicBindConfigSvcServer) mustEmbedUnimplementedDynamicBindConfigSvcServer() {}
func (UnimplementedDynamicBindConfigSvcServer) testEmbeddedByValue()                              {}

// UnsafeDynamicBindConfigSvcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DynamicBindConfigSvcServer will
// result in compilation errors.
type UnsafeDynamicBindConfigSvcServer interface {
	mustEmbedUnimplementedDynamicBindConfigSvcServer()
}

func RegisterDynamicBindConfigSvcServer(s grpc.ServiceRegistrar, srv DynamicBindConfigSvcServer) {
	// If the following call pancis, it indicates UnimplementedDynamicBindConfigSvcServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DynamicBindConfigSvc_ServiceDesc, srv)
}

func _DynamicBindConfigSvc_CreateDynamicBindConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDynamicBindConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DynamicBindConfigSvcServer).CreateDynamicBindConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DynamicBindConfigSvc_CreateDynamicBindConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DynamicBindConfigSvcServer).CreateDynamicBindConfig(ctx, req.(*CreateDynamicBindConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DynamicBindConfigSvc_GetDynamicBindConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(general.GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DynamicBindConfigSvcServer).GetDynamicBindConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DynamicBindConfigSvc_GetDynamicBindConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DynamicBindConfigSvcServer).GetDynamicBindConfig(ctx, req.(*general.GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DynamicBindConfigSvc_UpdateDynamicBindConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDynamicBindConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DynamicBindConfigSvcServer).UpdateDynamicBindConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DynamicBindConfigSvc_UpdateDynamicBindConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DynamicBindConfigSvcServer).UpdateDynamicBindConfig(ctx, req.(*UpdateDynamicBindConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DynamicBindConfigSvc_DeleteDynamicBindConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(general.ResourceId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DynamicBindConfigSvcServer).DeleteDynamicBindConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DynamicBindConfigSvc_DeleteDynamicBindConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DynamicBindConfigSvcServer).DeleteDynamicBindConfig(ctx, req.(*general.ResourceId))
	}
	return interceptor(ctx, in, info, handler)
}

func _DynamicBindConfigSvc_DeleteCollectionDynamicBindConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(general.ListOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DynamicBindConfigSvcServer).DeleteCollectionDynamicBindConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DynamicBindConfigSvc_DeleteCollectionDynamicBindConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DynamicBindConfigSvcServer).DeleteCollectionDynamicBindConfig(ctx, req.(*general.ListOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _DynamicBindConfigSvc_ListDynamicBindConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(general.ListOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DynamicBindConfigSvcServer).ListDynamicBindConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DynamicBindConfigSvc_ListDynamicBindConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DynamicBindConfigSvcServer).ListDynamicBindConfig(ctx, req.(*general.ListOptions))
	}
	return interceptor(ctx, in, info, handler)
}

// DynamicBindConfigSvc_ServiceDesc is the grpc.ServiceDesc for DynamicBindConfigSvc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DynamicBindConfigSvc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dbconfig.DynamicBindConfigSvc",
	HandlerType: (*DynamicBindConfigSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDynamicBindConfig",
			Handler:    _DynamicBindConfigSvc_CreateDynamicBindConfig_Handler,
		},
		{
			MethodName: "GetDynamicBindConfig",
			Handler:    _DynamicBindConfigSvc_GetDynamicBindConfig_Handler,
		},
		{
			MethodName: "UpdateDynamicBindConfig",
			Handler:    _DynamicBindConfigSvc_UpdateDynamicBindConfig_Handler,
		},
		{
			MethodName: "DeleteDynamicBindConfig",
			Handler:    _DynamicBindConfigSvc_DeleteDynamicBindConfig_Handler,
		},
		{
			MethodName: "DeleteCollectionDynamicBindConfig",
			Handler:    _DynamicBindConfigSvc_DeleteCollectionDynamicBindConfig_Handler,
		},
		{
			MethodName: "ListDynamicBindConfig",
			Handler:    _DynamicBindConfigSvc_ListDynamicBindConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dbconfig/dbconfig.proto",
}
