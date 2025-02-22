// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: authn/authn.proto

package authnpb

import (
	context "context"
	user "github.com/hobbyfarm/gargantua/v3/protos/user"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AuthN_AuthN_FullMethodName = "/authn.AuthN/AuthN"
)

// AuthNClient is the client API for AuthN service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service definition
type AuthNClient interface {
	AuthN(ctx context.Context, in *AuthNRequest, opts ...grpc.CallOption) (*user.User, error)
}

type authNClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthNClient(cc grpc.ClientConnInterface) AuthNClient {
	return &authNClient{cc}
}

func (c *authNClient) AuthN(ctx context.Context, in *AuthNRequest, opts ...grpc.CallOption) (*user.User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(user.User)
	err := c.cc.Invoke(ctx, AuthN_AuthN_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthNServer is the server API for AuthN service.
// All implementations must embed UnimplementedAuthNServer
// for forward compatibility.
//
// Service definition
type AuthNServer interface {
	AuthN(context.Context, *AuthNRequest) (*user.User, error)
	mustEmbedUnimplementedAuthNServer()
}

// UnimplementedAuthNServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAuthNServer struct{}

func (UnimplementedAuthNServer) AuthN(context.Context, *AuthNRequest) (*user.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthN not implemented")
}
func (UnimplementedAuthNServer) mustEmbedUnimplementedAuthNServer() {}
func (UnimplementedAuthNServer) testEmbeddedByValue()               {}

// UnsafeAuthNServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthNServer will
// result in compilation errors.
type UnsafeAuthNServer interface {
	mustEmbedUnimplementedAuthNServer()
}

func RegisterAuthNServer(s grpc.ServiceRegistrar, srv AuthNServer) {
	// If the following call pancis, it indicates UnimplementedAuthNServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AuthN_ServiceDesc, srv)
}

func _AuthN_AuthN_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthNRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthNServer).AuthN(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthN_AuthN_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthNServer).AuthN(ctx, req.(*AuthNRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthN_ServiceDesc is the grpc.ServiceDesc for AuthN service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthN_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "authn.AuthN",
	HandlerType: (*AuthNServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuthN",
			Handler:    _AuthN_AuthN_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authn/authn.proto",
}
