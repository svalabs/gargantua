// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: accesscode/accesscode.proto

package accesscodepb

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
	AccessCodeSvc_CreateAc_FullMethodName                = "/accesscode.AccessCodeSvc/CreateAc"
	AccessCodeSvc_GetAc_FullMethodName                   = "/accesscode.AccessCodeSvc/GetAc"
	AccessCodeSvc_UpdateAc_FullMethodName                = "/accesscode.AccessCodeSvc/UpdateAc"
	AccessCodeSvc_DeleteAc_FullMethodName                = "/accesscode.AccessCodeSvc/DeleteAc"
	AccessCodeSvc_DeleteCollectionAc_FullMethodName      = "/accesscode.AccessCodeSvc/DeleteCollectionAc"
	AccessCodeSvc_ListAc_FullMethodName                  = "/accesscode.AccessCodeSvc/ListAc"
	AccessCodeSvc_CreateOtac_FullMethodName              = "/accesscode.AccessCodeSvc/CreateOtac"
	AccessCodeSvc_GetOtac_FullMethodName                 = "/accesscode.AccessCodeSvc/GetOtac"
	AccessCodeSvc_UpdateOtac_FullMethodName              = "/accesscode.AccessCodeSvc/UpdateOtac"
	AccessCodeSvc_DeleteOtac_FullMethodName              = "/accesscode.AccessCodeSvc/DeleteOtac"
	AccessCodeSvc_DeleteCollectionOtac_FullMethodName    = "/accesscode.AccessCodeSvc/DeleteCollectionOtac"
	AccessCodeSvc_ListOtac_FullMethodName                = "/accesscode.AccessCodeSvc/ListOtac"
	AccessCodeSvc_ValidateExistence_FullMethodName       = "/accesscode.AccessCodeSvc/ValidateExistence"
	AccessCodeSvc_GetAccessCodesWithOTACs_FullMethodName = "/accesscode.AccessCodeSvc/GetAccessCodesWithOTACs"
	AccessCodeSvc_GetAccessCodeWithOTACs_FullMethodName  = "/accesscode.AccessCodeSvc/GetAccessCodeWithOTACs"
	AccessCodeSvc_GetAcOwnerReferences_FullMethodName    = "/accesscode.AccessCodeSvc/GetAcOwnerReferences"
)

// AccessCodeSvcClient is the client API for AccessCodeSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service definition
type AccessCodeSvcClient interface {
	// Resource oriented RPCs for AccessCodes:
	CreateAc(ctx context.Context, in *CreateAcRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetAc(ctx context.Context, in *general.GetRequest, opts ...grpc.CallOption) (*AccessCode, error)
	UpdateAc(ctx context.Context, in *UpdateAccessCodeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteAc(ctx context.Context, in *general.ResourceId, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteCollectionAc(ctx context.Context, in *general.ListOptions, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListAc(ctx context.Context, in *general.ListOptions, opts ...grpc.CallOption) (*ListAcsResponse, error)
	// Resource oriented RPCs for OneTimeAccessCodes:
	CreateOtac(ctx context.Context, in *CreateOtacRequest, opts ...grpc.CallOption) (*OneTimeAccessCode, error)
	GetOtac(ctx context.Context, in *general.GetRequest, opts ...grpc.CallOption) (*OneTimeAccessCode, error)
	UpdateOtac(ctx context.Context, in *OneTimeAccessCode, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteOtac(ctx context.Context, in *general.ResourceId, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteCollectionOtac(ctx context.Context, in *general.ListOptions, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListOtac(ctx context.Context, in *general.ListOptions, opts ...grpc.CallOption) (*ListOtacsResponse, error)
	// Helper functions
	ValidateExistence(ctx context.Context, in *general.ResourceId, opts ...grpc.CallOption) (*ResourceValidation, error)
	GetAccessCodesWithOTACs(ctx context.Context, in *ResourceIds, opts ...grpc.CallOption) (*ListAcsResponse, error)
	GetAccessCodeWithOTACs(ctx context.Context, in *general.ResourceId, opts ...grpc.CallOption) (*AccessCode, error)
	GetAcOwnerReferences(ctx context.Context, in *general.GetRequest, opts ...grpc.CallOption) (*general.OwnerReferences, error)
}

type accessCodeSvcClient struct {
	cc grpc.ClientConnInterface
}

func NewAccessCodeSvcClient(cc grpc.ClientConnInterface) AccessCodeSvcClient {
	return &accessCodeSvcClient{cc}
}

func (c *accessCodeSvcClient) CreateAc(ctx context.Context, in *CreateAcRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AccessCodeSvc_CreateAc_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessCodeSvcClient) GetAc(ctx context.Context, in *general.GetRequest, opts ...grpc.CallOption) (*AccessCode, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AccessCode)
	err := c.cc.Invoke(ctx, AccessCodeSvc_GetAc_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessCodeSvcClient) UpdateAc(ctx context.Context, in *UpdateAccessCodeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AccessCodeSvc_UpdateAc_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessCodeSvcClient) DeleteAc(ctx context.Context, in *general.ResourceId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AccessCodeSvc_DeleteAc_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessCodeSvcClient) DeleteCollectionAc(ctx context.Context, in *general.ListOptions, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AccessCodeSvc_DeleteCollectionAc_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessCodeSvcClient) ListAc(ctx context.Context, in *general.ListOptions, opts ...grpc.CallOption) (*ListAcsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListAcsResponse)
	err := c.cc.Invoke(ctx, AccessCodeSvc_ListAc_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessCodeSvcClient) CreateOtac(ctx context.Context, in *CreateOtacRequest, opts ...grpc.CallOption) (*OneTimeAccessCode, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OneTimeAccessCode)
	err := c.cc.Invoke(ctx, AccessCodeSvc_CreateOtac_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessCodeSvcClient) GetOtac(ctx context.Context, in *general.GetRequest, opts ...grpc.CallOption) (*OneTimeAccessCode, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OneTimeAccessCode)
	err := c.cc.Invoke(ctx, AccessCodeSvc_GetOtac_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessCodeSvcClient) UpdateOtac(ctx context.Context, in *OneTimeAccessCode, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AccessCodeSvc_UpdateOtac_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessCodeSvcClient) DeleteOtac(ctx context.Context, in *general.ResourceId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AccessCodeSvc_DeleteOtac_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessCodeSvcClient) DeleteCollectionOtac(ctx context.Context, in *general.ListOptions, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AccessCodeSvc_DeleteCollectionOtac_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessCodeSvcClient) ListOtac(ctx context.Context, in *general.ListOptions, opts ...grpc.CallOption) (*ListOtacsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListOtacsResponse)
	err := c.cc.Invoke(ctx, AccessCodeSvc_ListOtac_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessCodeSvcClient) ValidateExistence(ctx context.Context, in *general.ResourceId, opts ...grpc.CallOption) (*ResourceValidation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResourceValidation)
	err := c.cc.Invoke(ctx, AccessCodeSvc_ValidateExistence_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessCodeSvcClient) GetAccessCodesWithOTACs(ctx context.Context, in *ResourceIds, opts ...grpc.CallOption) (*ListAcsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListAcsResponse)
	err := c.cc.Invoke(ctx, AccessCodeSvc_GetAccessCodesWithOTACs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessCodeSvcClient) GetAccessCodeWithOTACs(ctx context.Context, in *general.ResourceId, opts ...grpc.CallOption) (*AccessCode, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AccessCode)
	err := c.cc.Invoke(ctx, AccessCodeSvc_GetAccessCodeWithOTACs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessCodeSvcClient) GetAcOwnerReferences(ctx context.Context, in *general.GetRequest, opts ...grpc.CallOption) (*general.OwnerReferences, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(general.OwnerReferences)
	err := c.cc.Invoke(ctx, AccessCodeSvc_GetAcOwnerReferences_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccessCodeSvcServer is the server API for AccessCodeSvc service.
// All implementations must embed UnimplementedAccessCodeSvcServer
// for forward compatibility.
//
// Service definition
type AccessCodeSvcServer interface {
	// Resource oriented RPCs for AccessCodes:
	CreateAc(context.Context, *CreateAcRequest) (*emptypb.Empty, error)
	GetAc(context.Context, *general.GetRequest) (*AccessCode, error)
	UpdateAc(context.Context, *UpdateAccessCodeRequest) (*emptypb.Empty, error)
	DeleteAc(context.Context, *general.ResourceId) (*emptypb.Empty, error)
	DeleteCollectionAc(context.Context, *general.ListOptions) (*emptypb.Empty, error)
	ListAc(context.Context, *general.ListOptions) (*ListAcsResponse, error)
	// Resource oriented RPCs for OneTimeAccessCodes:
	CreateOtac(context.Context, *CreateOtacRequest) (*OneTimeAccessCode, error)
	GetOtac(context.Context, *general.GetRequest) (*OneTimeAccessCode, error)
	UpdateOtac(context.Context, *OneTimeAccessCode) (*emptypb.Empty, error)
	DeleteOtac(context.Context, *general.ResourceId) (*emptypb.Empty, error)
	DeleteCollectionOtac(context.Context, *general.ListOptions) (*emptypb.Empty, error)
	ListOtac(context.Context, *general.ListOptions) (*ListOtacsResponse, error)
	// Helper functions
	ValidateExistence(context.Context, *general.ResourceId) (*ResourceValidation, error)
	GetAccessCodesWithOTACs(context.Context, *ResourceIds) (*ListAcsResponse, error)
	GetAccessCodeWithOTACs(context.Context, *general.ResourceId) (*AccessCode, error)
	GetAcOwnerReferences(context.Context, *general.GetRequest) (*general.OwnerReferences, error)
	mustEmbedUnimplementedAccessCodeSvcServer()
}

// UnimplementedAccessCodeSvcServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAccessCodeSvcServer struct{}

func (UnimplementedAccessCodeSvcServer) CreateAc(context.Context, *CreateAcRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAc not implemented")
}
func (UnimplementedAccessCodeSvcServer) GetAc(context.Context, *general.GetRequest) (*AccessCode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAc not implemented")
}
func (UnimplementedAccessCodeSvcServer) UpdateAc(context.Context, *UpdateAccessCodeRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAc not implemented")
}
func (UnimplementedAccessCodeSvcServer) DeleteAc(context.Context, *general.ResourceId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAc not implemented")
}
func (UnimplementedAccessCodeSvcServer) DeleteCollectionAc(context.Context, *general.ListOptions) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCollectionAc not implemented")
}
func (UnimplementedAccessCodeSvcServer) ListAc(context.Context, *general.ListOptions) (*ListAcsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAc not implemented")
}
func (UnimplementedAccessCodeSvcServer) CreateOtac(context.Context, *CreateOtacRequest) (*OneTimeAccessCode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOtac not implemented")
}
func (UnimplementedAccessCodeSvcServer) GetOtac(context.Context, *general.GetRequest) (*OneTimeAccessCode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOtac not implemented")
}
func (UnimplementedAccessCodeSvcServer) UpdateOtac(context.Context, *OneTimeAccessCode) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOtac not implemented")
}
func (UnimplementedAccessCodeSvcServer) DeleteOtac(context.Context, *general.ResourceId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOtac not implemented")
}
func (UnimplementedAccessCodeSvcServer) DeleteCollectionOtac(context.Context, *general.ListOptions) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCollectionOtac not implemented")
}
func (UnimplementedAccessCodeSvcServer) ListOtac(context.Context, *general.ListOptions) (*ListOtacsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOtac not implemented")
}
func (UnimplementedAccessCodeSvcServer) ValidateExistence(context.Context, *general.ResourceId) (*ResourceValidation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateExistence not implemented")
}
func (UnimplementedAccessCodeSvcServer) GetAccessCodesWithOTACs(context.Context, *ResourceIds) (*ListAcsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccessCodesWithOTACs not implemented")
}
func (UnimplementedAccessCodeSvcServer) GetAccessCodeWithOTACs(context.Context, *general.ResourceId) (*AccessCode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccessCodeWithOTACs not implemented")
}
func (UnimplementedAccessCodeSvcServer) GetAcOwnerReferences(context.Context, *general.GetRequest) (*general.OwnerReferences, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAcOwnerReferences not implemented")
}
func (UnimplementedAccessCodeSvcServer) mustEmbedUnimplementedAccessCodeSvcServer() {}
func (UnimplementedAccessCodeSvcServer) testEmbeddedByValue()                       {}

// UnsafeAccessCodeSvcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccessCodeSvcServer will
// result in compilation errors.
type UnsafeAccessCodeSvcServer interface {
	mustEmbedUnimplementedAccessCodeSvcServer()
}

func RegisterAccessCodeSvcServer(s grpc.ServiceRegistrar, srv AccessCodeSvcServer) {
	// If the following call pancis, it indicates UnimplementedAccessCodeSvcServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AccessCodeSvc_ServiceDesc, srv)
}

func _AccessCodeSvc_CreateAc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessCodeSvcServer).CreateAc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessCodeSvc_CreateAc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessCodeSvcServer).CreateAc(ctx, req.(*CreateAcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessCodeSvc_GetAc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(general.GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessCodeSvcServer).GetAc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessCodeSvc_GetAc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessCodeSvcServer).GetAc(ctx, req.(*general.GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessCodeSvc_UpdateAc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAccessCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessCodeSvcServer).UpdateAc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessCodeSvc_UpdateAc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessCodeSvcServer).UpdateAc(ctx, req.(*UpdateAccessCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessCodeSvc_DeleteAc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(general.ResourceId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessCodeSvcServer).DeleteAc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessCodeSvc_DeleteAc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessCodeSvcServer).DeleteAc(ctx, req.(*general.ResourceId))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessCodeSvc_DeleteCollectionAc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(general.ListOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessCodeSvcServer).DeleteCollectionAc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessCodeSvc_DeleteCollectionAc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessCodeSvcServer).DeleteCollectionAc(ctx, req.(*general.ListOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessCodeSvc_ListAc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(general.ListOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessCodeSvcServer).ListAc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessCodeSvc_ListAc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessCodeSvcServer).ListAc(ctx, req.(*general.ListOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessCodeSvc_CreateOtac_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOtacRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessCodeSvcServer).CreateOtac(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessCodeSvc_CreateOtac_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessCodeSvcServer).CreateOtac(ctx, req.(*CreateOtacRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessCodeSvc_GetOtac_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(general.GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessCodeSvcServer).GetOtac(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessCodeSvc_GetOtac_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessCodeSvcServer).GetOtac(ctx, req.(*general.GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessCodeSvc_UpdateOtac_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OneTimeAccessCode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessCodeSvcServer).UpdateOtac(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessCodeSvc_UpdateOtac_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessCodeSvcServer).UpdateOtac(ctx, req.(*OneTimeAccessCode))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessCodeSvc_DeleteOtac_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(general.ResourceId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessCodeSvcServer).DeleteOtac(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessCodeSvc_DeleteOtac_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessCodeSvcServer).DeleteOtac(ctx, req.(*general.ResourceId))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessCodeSvc_DeleteCollectionOtac_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(general.ListOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessCodeSvcServer).DeleteCollectionOtac(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessCodeSvc_DeleteCollectionOtac_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessCodeSvcServer).DeleteCollectionOtac(ctx, req.(*general.ListOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessCodeSvc_ListOtac_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(general.ListOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessCodeSvcServer).ListOtac(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessCodeSvc_ListOtac_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessCodeSvcServer).ListOtac(ctx, req.(*general.ListOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessCodeSvc_ValidateExistence_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(general.ResourceId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessCodeSvcServer).ValidateExistence(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessCodeSvc_ValidateExistence_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessCodeSvcServer).ValidateExistence(ctx, req.(*general.ResourceId))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessCodeSvc_GetAccessCodesWithOTACs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceIds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessCodeSvcServer).GetAccessCodesWithOTACs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessCodeSvc_GetAccessCodesWithOTACs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessCodeSvcServer).GetAccessCodesWithOTACs(ctx, req.(*ResourceIds))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessCodeSvc_GetAccessCodeWithOTACs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(general.ResourceId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessCodeSvcServer).GetAccessCodeWithOTACs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessCodeSvc_GetAccessCodeWithOTACs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessCodeSvcServer).GetAccessCodeWithOTACs(ctx, req.(*general.ResourceId))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessCodeSvc_GetAcOwnerReferences_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(general.GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessCodeSvcServer).GetAcOwnerReferences(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessCodeSvc_GetAcOwnerReferences_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessCodeSvcServer).GetAcOwnerReferences(ctx, req.(*general.GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccessCodeSvc_ServiceDesc is the grpc.ServiceDesc for AccessCodeSvc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccessCodeSvc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "accesscode.AccessCodeSvc",
	HandlerType: (*AccessCodeSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAc",
			Handler:    _AccessCodeSvc_CreateAc_Handler,
		},
		{
			MethodName: "GetAc",
			Handler:    _AccessCodeSvc_GetAc_Handler,
		},
		{
			MethodName: "UpdateAc",
			Handler:    _AccessCodeSvc_UpdateAc_Handler,
		},
		{
			MethodName: "DeleteAc",
			Handler:    _AccessCodeSvc_DeleteAc_Handler,
		},
		{
			MethodName: "DeleteCollectionAc",
			Handler:    _AccessCodeSvc_DeleteCollectionAc_Handler,
		},
		{
			MethodName: "ListAc",
			Handler:    _AccessCodeSvc_ListAc_Handler,
		},
		{
			MethodName: "CreateOtac",
			Handler:    _AccessCodeSvc_CreateOtac_Handler,
		},
		{
			MethodName: "GetOtac",
			Handler:    _AccessCodeSvc_GetOtac_Handler,
		},
		{
			MethodName: "UpdateOtac",
			Handler:    _AccessCodeSvc_UpdateOtac_Handler,
		},
		{
			MethodName: "DeleteOtac",
			Handler:    _AccessCodeSvc_DeleteOtac_Handler,
		},
		{
			MethodName: "DeleteCollectionOtac",
			Handler:    _AccessCodeSvc_DeleteCollectionOtac_Handler,
		},
		{
			MethodName: "ListOtac",
			Handler:    _AccessCodeSvc_ListOtac_Handler,
		},
		{
			MethodName: "ValidateExistence",
			Handler:    _AccessCodeSvc_ValidateExistence_Handler,
		},
		{
			MethodName: "GetAccessCodesWithOTACs",
			Handler:    _AccessCodeSvc_GetAccessCodesWithOTACs_Handler,
		},
		{
			MethodName: "GetAccessCodeWithOTACs",
			Handler:    _AccessCodeSvc_GetAccessCodeWithOTACs_Handler,
		},
		{
			MethodName: "GetAcOwnerReferences",
			Handler:    _AccessCodeSvc_GetAcOwnerReferences_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accesscode/accesscode.proto",
}
