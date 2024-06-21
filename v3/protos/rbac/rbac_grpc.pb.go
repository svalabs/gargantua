// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: rbac/rbac.proto

package rbacpb

import (
	context "context"
	authr "github.com/hobbyfarm/gargantua/v3/protos/authr"
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
	RbacSvc_Grants_FullMethodName                   = "/rbac.RbacSvc/Grants"
	RbacSvc_GetAccessSet_FullMethodName             = "/rbac.RbacSvc/GetAccessSet"
	RbacSvc_GetHobbyfarmRoleBindings_FullMethodName = "/rbac.RbacSvc/GetHobbyfarmRoleBindings"
	RbacSvc_CreateRole_FullMethodName               = "/rbac.RbacSvc/CreateRole"
	RbacSvc_GetRole_FullMethodName                  = "/rbac.RbacSvc/GetRole"
	RbacSvc_UpdateRole_FullMethodName               = "/rbac.RbacSvc/UpdateRole"
	RbacSvc_DeleteRole_FullMethodName               = "/rbac.RbacSvc/DeleteRole"
	RbacSvc_ListRole_FullMethodName                 = "/rbac.RbacSvc/ListRole"
	RbacSvc_CreateRolebinding_FullMethodName        = "/rbac.RbacSvc/CreateRolebinding"
	RbacSvc_GetRolebinding_FullMethodName           = "/rbac.RbacSvc/GetRolebinding"
	RbacSvc_UpdateRolebinding_FullMethodName        = "/rbac.RbacSvc/UpdateRolebinding"
	RbacSvc_DeleteRolebinding_FullMethodName        = "/rbac.RbacSvc/DeleteRolebinding"
	RbacSvc_ListRolebinding_FullMethodName          = "/rbac.RbacSvc/ListRolebinding"
)

// RbacSvcClient is the client API for RbacSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RbacSvcClient interface {
	Grants(ctx context.Context, in *GrantRequest, opts ...grpc.CallOption) (*authr.AuthRResponse, error)
	GetAccessSet(ctx context.Context, in *general.ResourceId, opts ...grpc.CallOption) (*AccessSet, error)
	GetHobbyfarmRoleBindings(ctx context.Context, in *general.ResourceId, opts ...grpc.CallOption) (*RoleBindings, error)
	// Resource oriented RPCs for roles:
	CreateRole(ctx context.Context, in *Role, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetRole(ctx context.Context, in *general.GetRequest, opts ...grpc.CallOption) (*Role, error)
	UpdateRole(ctx context.Context, in *Role, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteRole(ctx context.Context, in *general.ResourceId, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListRole(ctx context.Context, in *general.ListOptions, opts ...grpc.CallOption) (*Roles, error)
	// Resource oriented RPCs for rolebindings:
	CreateRolebinding(ctx context.Context, in *RoleBinding, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetRolebinding(ctx context.Context, in *general.GetRequest, opts ...grpc.CallOption) (*RoleBinding, error)
	UpdateRolebinding(ctx context.Context, in *RoleBinding, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteRolebinding(ctx context.Context, in *general.ResourceId, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListRolebinding(ctx context.Context, in *general.ListOptions, opts ...grpc.CallOption) (*RoleBindings, error)
}

type rbacSvcClient struct {
	cc grpc.ClientConnInterface
}

func NewRbacSvcClient(cc grpc.ClientConnInterface) RbacSvcClient {
	return &rbacSvcClient{cc}
}

func (c *rbacSvcClient) Grants(ctx context.Context, in *GrantRequest, opts ...grpc.CallOption) (*authr.AuthRResponse, error) {
	out := new(authr.AuthRResponse)
	err := c.cc.Invoke(ctx, RbacSvc_Grants_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacSvcClient) GetAccessSet(ctx context.Context, in *general.ResourceId, opts ...grpc.CallOption) (*AccessSet, error) {
	out := new(AccessSet)
	err := c.cc.Invoke(ctx, RbacSvc_GetAccessSet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacSvcClient) GetHobbyfarmRoleBindings(ctx context.Context, in *general.ResourceId, opts ...grpc.CallOption) (*RoleBindings, error) {
	out := new(RoleBindings)
	err := c.cc.Invoke(ctx, RbacSvc_GetHobbyfarmRoleBindings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacSvcClient) CreateRole(ctx context.Context, in *Role, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, RbacSvc_CreateRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacSvcClient) GetRole(ctx context.Context, in *general.GetRequest, opts ...grpc.CallOption) (*Role, error) {
	out := new(Role)
	err := c.cc.Invoke(ctx, RbacSvc_GetRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacSvcClient) UpdateRole(ctx context.Context, in *Role, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, RbacSvc_UpdateRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacSvcClient) DeleteRole(ctx context.Context, in *general.ResourceId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, RbacSvc_DeleteRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacSvcClient) ListRole(ctx context.Context, in *general.ListOptions, opts ...grpc.CallOption) (*Roles, error) {
	out := new(Roles)
	err := c.cc.Invoke(ctx, RbacSvc_ListRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacSvcClient) CreateRolebinding(ctx context.Context, in *RoleBinding, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, RbacSvc_CreateRolebinding_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacSvcClient) GetRolebinding(ctx context.Context, in *general.GetRequest, opts ...grpc.CallOption) (*RoleBinding, error) {
	out := new(RoleBinding)
	err := c.cc.Invoke(ctx, RbacSvc_GetRolebinding_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacSvcClient) UpdateRolebinding(ctx context.Context, in *RoleBinding, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, RbacSvc_UpdateRolebinding_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacSvcClient) DeleteRolebinding(ctx context.Context, in *general.ResourceId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, RbacSvc_DeleteRolebinding_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacSvcClient) ListRolebinding(ctx context.Context, in *general.ListOptions, opts ...grpc.CallOption) (*RoleBindings, error) {
	out := new(RoleBindings)
	err := c.cc.Invoke(ctx, RbacSvc_ListRolebinding_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RbacSvcServer is the server API for RbacSvc service.
// All implementations must embed UnimplementedRbacSvcServer
// for forward compatibility
type RbacSvcServer interface {
	Grants(context.Context, *GrantRequest) (*authr.AuthRResponse, error)
	GetAccessSet(context.Context, *general.ResourceId) (*AccessSet, error)
	GetHobbyfarmRoleBindings(context.Context, *general.ResourceId) (*RoleBindings, error)
	// Resource oriented RPCs for roles:
	CreateRole(context.Context, *Role) (*emptypb.Empty, error)
	GetRole(context.Context, *general.GetRequest) (*Role, error)
	UpdateRole(context.Context, *Role) (*emptypb.Empty, error)
	DeleteRole(context.Context, *general.ResourceId) (*emptypb.Empty, error)
	ListRole(context.Context, *general.ListOptions) (*Roles, error)
	// Resource oriented RPCs for rolebindings:
	CreateRolebinding(context.Context, *RoleBinding) (*emptypb.Empty, error)
	GetRolebinding(context.Context, *general.GetRequest) (*RoleBinding, error)
	UpdateRolebinding(context.Context, *RoleBinding) (*emptypb.Empty, error)
	DeleteRolebinding(context.Context, *general.ResourceId) (*emptypb.Empty, error)
	ListRolebinding(context.Context, *general.ListOptions) (*RoleBindings, error)
	mustEmbedUnimplementedRbacSvcServer()
}

// UnimplementedRbacSvcServer must be embedded to have forward compatible implementations.
type UnimplementedRbacSvcServer struct {
}

func (UnimplementedRbacSvcServer) Grants(context.Context, *GrantRequest) (*authr.AuthRResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Grants not implemented")
}
func (UnimplementedRbacSvcServer) GetAccessSet(context.Context, *general.ResourceId) (*AccessSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccessSet not implemented")
}
func (UnimplementedRbacSvcServer) GetHobbyfarmRoleBindings(context.Context, *general.ResourceId) (*RoleBindings, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHobbyfarmRoleBindings not implemented")
}
func (UnimplementedRbacSvcServer) CreateRole(context.Context, *Role) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRole not implemented")
}
func (UnimplementedRbacSvcServer) GetRole(context.Context, *general.GetRequest) (*Role, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRole not implemented")
}
func (UnimplementedRbacSvcServer) UpdateRole(context.Context, *Role) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRole not implemented")
}
func (UnimplementedRbacSvcServer) DeleteRole(context.Context, *general.ResourceId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRole not implemented")
}
func (UnimplementedRbacSvcServer) ListRole(context.Context, *general.ListOptions) (*Roles, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRole not implemented")
}
func (UnimplementedRbacSvcServer) CreateRolebinding(context.Context, *RoleBinding) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRolebinding not implemented")
}
func (UnimplementedRbacSvcServer) GetRolebinding(context.Context, *general.GetRequest) (*RoleBinding, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRolebinding not implemented")
}
func (UnimplementedRbacSvcServer) UpdateRolebinding(context.Context, *RoleBinding) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRolebinding not implemented")
}
func (UnimplementedRbacSvcServer) DeleteRolebinding(context.Context, *general.ResourceId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRolebinding not implemented")
}
func (UnimplementedRbacSvcServer) ListRolebinding(context.Context, *general.ListOptions) (*RoleBindings, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRolebinding not implemented")
}
func (UnimplementedRbacSvcServer) mustEmbedUnimplementedRbacSvcServer() {}

// UnsafeRbacSvcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RbacSvcServer will
// result in compilation errors.
type UnsafeRbacSvcServer interface {
	mustEmbedUnimplementedRbacSvcServer()
}

func RegisterRbacSvcServer(s grpc.ServiceRegistrar, srv RbacSvcServer) {
	s.RegisterService(&RbacSvc_ServiceDesc, srv)
}

func _RbacSvc_Grants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GrantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacSvcServer).Grants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RbacSvc_Grants_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacSvcServer).Grants(ctx, req.(*GrantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacSvc_GetAccessSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(general.ResourceId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacSvcServer).GetAccessSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RbacSvc_GetAccessSet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacSvcServer).GetAccessSet(ctx, req.(*general.ResourceId))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacSvc_GetHobbyfarmRoleBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(general.ResourceId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacSvcServer).GetHobbyfarmRoleBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RbacSvc_GetHobbyfarmRoleBindings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacSvcServer).GetHobbyfarmRoleBindings(ctx, req.(*general.ResourceId))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacSvc_CreateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Role)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacSvcServer).CreateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RbacSvc_CreateRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacSvcServer).CreateRole(ctx, req.(*Role))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacSvc_GetRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(general.GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacSvcServer).GetRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RbacSvc_GetRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacSvcServer).GetRole(ctx, req.(*general.GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacSvc_UpdateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Role)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacSvcServer).UpdateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RbacSvc_UpdateRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacSvcServer).UpdateRole(ctx, req.(*Role))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacSvc_DeleteRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(general.ResourceId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacSvcServer).DeleteRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RbacSvc_DeleteRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacSvcServer).DeleteRole(ctx, req.(*general.ResourceId))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacSvc_ListRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(general.ListOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacSvcServer).ListRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RbacSvc_ListRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacSvcServer).ListRole(ctx, req.(*general.ListOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacSvc_CreateRolebinding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleBinding)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacSvcServer).CreateRolebinding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RbacSvc_CreateRolebinding_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacSvcServer).CreateRolebinding(ctx, req.(*RoleBinding))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacSvc_GetRolebinding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(general.GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacSvcServer).GetRolebinding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RbacSvc_GetRolebinding_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacSvcServer).GetRolebinding(ctx, req.(*general.GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacSvc_UpdateRolebinding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleBinding)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacSvcServer).UpdateRolebinding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RbacSvc_UpdateRolebinding_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacSvcServer).UpdateRolebinding(ctx, req.(*RoleBinding))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacSvc_DeleteRolebinding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(general.ResourceId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacSvcServer).DeleteRolebinding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RbacSvc_DeleteRolebinding_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacSvcServer).DeleteRolebinding(ctx, req.(*general.ResourceId))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacSvc_ListRolebinding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(general.ListOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacSvcServer).ListRolebinding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RbacSvc_ListRolebinding_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacSvcServer).ListRolebinding(ctx, req.(*general.ListOptions))
	}
	return interceptor(ctx, in, info, handler)
}

// RbacSvc_ServiceDesc is the grpc.ServiceDesc for RbacSvc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RbacSvc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rbac.RbacSvc",
	HandlerType: (*RbacSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Grants",
			Handler:    _RbacSvc_Grants_Handler,
		},
		{
			MethodName: "GetAccessSet",
			Handler:    _RbacSvc_GetAccessSet_Handler,
		},
		{
			MethodName: "GetHobbyfarmRoleBindings",
			Handler:    _RbacSvc_GetHobbyfarmRoleBindings_Handler,
		},
		{
			MethodName: "CreateRole",
			Handler:    _RbacSvc_CreateRole_Handler,
		},
		{
			MethodName: "GetRole",
			Handler:    _RbacSvc_GetRole_Handler,
		},
		{
			MethodName: "UpdateRole",
			Handler:    _RbacSvc_UpdateRole_Handler,
		},
		{
			MethodName: "DeleteRole",
			Handler:    _RbacSvc_DeleteRole_Handler,
		},
		{
			MethodName: "ListRole",
			Handler:    _RbacSvc_ListRole_Handler,
		},
		{
			MethodName: "CreateRolebinding",
			Handler:    _RbacSvc_CreateRolebinding_Handler,
		},
		{
			MethodName: "GetRolebinding",
			Handler:    _RbacSvc_GetRolebinding_Handler,
		},
		{
			MethodName: "UpdateRolebinding",
			Handler:    _RbacSvc_UpdateRolebinding_Handler,
		},
		{
			MethodName: "DeleteRolebinding",
			Handler:    _RbacSvc_DeleteRolebinding_Handler,
		},
		{
			MethodName: "ListRolebinding",
			Handler:    _RbacSvc_ListRolebinding_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rbac/rbac.proto",
}
