// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: vmtemplate/vmtemplate.proto

package vmtemplatepb

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
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	VMTemplateSvc_CreateVMTemplate_FullMethodName           = "/vmtemplate.VMTemplateSvc/CreateVMTemplate"
	VMTemplateSvc_GetVMTemplate_FullMethodName              = "/vmtemplate.VMTemplateSvc/GetVMTemplate"
	VMTemplateSvc_UpdateVMTemplate_FullMethodName           = "/vmtemplate.VMTemplateSvc/UpdateVMTemplate"
	VMTemplateSvc_DeleteVMTemplate_FullMethodName           = "/vmtemplate.VMTemplateSvc/DeleteVMTemplate"
	VMTemplateSvc_DeleteCollectionVMTemplate_FullMethodName = "/vmtemplate.VMTemplateSvc/DeleteCollectionVMTemplate"
	VMTemplateSvc_ListVMTemplate_FullMethodName             = "/vmtemplate.VMTemplateSvc/ListVMTemplate"
)

// VMTemplateSvcClient is the client API for VMTemplateSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VMTemplateSvcClient interface {
	CreateVMTemplate(ctx context.Context, in *CreateVMTemplateRequest, opts ...grpc.CallOption) (*general.ResourceId, error)
	GetVMTemplate(ctx context.Context, in *general.GetRequest, opts ...grpc.CallOption) (*VMTemplate, error)
	UpdateVMTemplate(ctx context.Context, in *UpdateVMTemplateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteVMTemplate(ctx context.Context, in *general.ResourceId, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteCollectionVMTemplate(ctx context.Context, in *general.ListOptions, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListVMTemplate(ctx context.Context, in *general.ListOptions, opts ...grpc.CallOption) (*ListVMTemplatesResponse, error)
}

type vMTemplateSvcClient struct {
	cc grpc.ClientConnInterface
}

func NewVMTemplateSvcClient(cc grpc.ClientConnInterface) VMTemplateSvcClient {
	return &vMTemplateSvcClient{cc}
}

func (c *vMTemplateSvcClient) CreateVMTemplate(ctx context.Context, in *CreateVMTemplateRequest, opts ...grpc.CallOption) (*general.ResourceId, error) {
	out := new(general.ResourceId)
	err := c.cc.Invoke(ctx, VMTemplateSvc_CreateVMTemplate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vMTemplateSvcClient) GetVMTemplate(ctx context.Context, in *general.GetRequest, opts ...grpc.CallOption) (*VMTemplate, error) {
	out := new(VMTemplate)
	err := c.cc.Invoke(ctx, VMTemplateSvc_GetVMTemplate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vMTemplateSvcClient) UpdateVMTemplate(ctx context.Context, in *UpdateVMTemplateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, VMTemplateSvc_UpdateVMTemplate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vMTemplateSvcClient) DeleteVMTemplate(ctx context.Context, in *general.ResourceId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, VMTemplateSvc_DeleteVMTemplate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vMTemplateSvcClient) DeleteCollectionVMTemplate(ctx context.Context, in *general.ListOptions, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, VMTemplateSvc_DeleteCollectionVMTemplate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vMTemplateSvcClient) ListVMTemplate(ctx context.Context, in *general.ListOptions, opts ...grpc.CallOption) (*ListVMTemplatesResponse, error) {
	out := new(ListVMTemplatesResponse)
	err := c.cc.Invoke(ctx, VMTemplateSvc_ListVMTemplate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VMTemplateSvcServer is the server API for VMTemplateSvc service.
// All implementations must embed UnimplementedVMTemplateSvcServer
// for forward compatibility
type VMTemplateSvcServer interface {
	CreateVMTemplate(context.Context, *CreateVMTemplateRequest) (*general.ResourceId, error)
	GetVMTemplate(context.Context, *general.GetRequest) (*VMTemplate, error)
	UpdateVMTemplate(context.Context, *UpdateVMTemplateRequest) (*emptypb.Empty, error)
	DeleteVMTemplate(context.Context, *general.ResourceId) (*emptypb.Empty, error)
	DeleteCollectionVMTemplate(context.Context, *general.ListOptions) (*emptypb.Empty, error)
	ListVMTemplate(context.Context, *general.ListOptions) (*ListVMTemplatesResponse, error)
	mustEmbedUnimplementedVMTemplateSvcServer()
}

// UnimplementedVMTemplateSvcServer must be embedded to have forward compatible implementations.
type UnimplementedVMTemplateSvcServer struct {
}

func (UnimplementedVMTemplateSvcServer) CreateVMTemplate(context.Context, *CreateVMTemplateRequest) (*general.ResourceId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVMTemplate not implemented")
}
func (UnimplementedVMTemplateSvcServer) GetVMTemplate(context.Context, *general.GetRequest) (*VMTemplate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVMTemplate not implemented")
}
func (UnimplementedVMTemplateSvcServer) UpdateVMTemplate(context.Context, *UpdateVMTemplateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateVMTemplate not implemented")
}
func (UnimplementedVMTemplateSvcServer) DeleteVMTemplate(context.Context, *general.ResourceId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVMTemplate not implemented")
}
func (UnimplementedVMTemplateSvcServer) DeleteCollectionVMTemplate(context.Context, *general.ListOptions) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCollectionVMTemplate not implemented")
}
func (UnimplementedVMTemplateSvcServer) ListVMTemplate(context.Context, *general.ListOptions) (*ListVMTemplatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVMTemplate not implemented")
}
func (UnimplementedVMTemplateSvcServer) mustEmbedUnimplementedVMTemplateSvcServer() {}

// UnsafeVMTemplateSvcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VMTemplateSvcServer will
// result in compilation errors.
type UnsafeVMTemplateSvcServer interface {
	mustEmbedUnimplementedVMTemplateSvcServer()
}

func RegisterVMTemplateSvcServer(s grpc.ServiceRegistrar, srv VMTemplateSvcServer) {
	s.RegisterService(&VMTemplateSvc_ServiceDesc, srv)
}

func _VMTemplateSvc_CreateVMTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVMTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMTemplateSvcServer).CreateVMTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VMTemplateSvc_CreateVMTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMTemplateSvcServer).CreateVMTemplate(ctx, req.(*CreateVMTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VMTemplateSvc_GetVMTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(general.GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMTemplateSvcServer).GetVMTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VMTemplateSvc_GetVMTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMTemplateSvcServer).GetVMTemplate(ctx, req.(*general.GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VMTemplateSvc_UpdateVMTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateVMTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMTemplateSvcServer).UpdateVMTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VMTemplateSvc_UpdateVMTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMTemplateSvcServer).UpdateVMTemplate(ctx, req.(*UpdateVMTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VMTemplateSvc_DeleteVMTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(general.ResourceId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMTemplateSvcServer).DeleteVMTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VMTemplateSvc_DeleteVMTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMTemplateSvcServer).DeleteVMTemplate(ctx, req.(*general.ResourceId))
	}
	return interceptor(ctx, in, info, handler)
}

func _VMTemplateSvc_DeleteCollectionVMTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(general.ListOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMTemplateSvcServer).DeleteCollectionVMTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VMTemplateSvc_DeleteCollectionVMTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMTemplateSvcServer).DeleteCollectionVMTemplate(ctx, req.(*general.ListOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _VMTemplateSvc_ListVMTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(general.ListOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMTemplateSvcServer).ListVMTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VMTemplateSvc_ListVMTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMTemplateSvcServer).ListVMTemplate(ctx, req.(*general.ListOptions))
	}
	return interceptor(ctx, in, info, handler)
}

// VMTemplateSvc_ServiceDesc is the grpc.ServiceDesc for VMTemplateSvc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VMTemplateSvc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vmtemplate.VMTemplateSvc",
	HandlerType: (*VMTemplateSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateVMTemplate",
			Handler:    _VMTemplateSvc_CreateVMTemplate_Handler,
		},
		{
			MethodName: "GetVMTemplate",
			Handler:    _VMTemplateSvc_GetVMTemplate_Handler,
		},
		{
			MethodName: "UpdateVMTemplate",
			Handler:    _VMTemplateSvc_UpdateVMTemplate_Handler,
		},
		{
			MethodName: "DeleteVMTemplate",
			Handler:    _VMTemplateSvc_DeleteVMTemplate_Handler,
		},
		{
			MethodName: "DeleteCollectionVMTemplate",
			Handler:    _VMTemplateSvc_DeleteCollectionVMTemplate_Handler,
		},
		{
			MethodName: "ListVMTemplate",
			Handler:    _VMTemplateSvc_ListVMTemplate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vmtemplate/vmtemplate.proto",
}