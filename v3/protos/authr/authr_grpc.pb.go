// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: authr/authr.proto

package authrpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AuthR_AuthR_FullMethodName = "/authr.AuthR/AuthR"
)

// AuthRClient is the client API for AuthR service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthRClient interface {
	AuthR(ctx context.Context, in *AuthRRequest, opts ...grpc.CallOption) (*AuthRResponse, error)
}

type authRClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthRClient(cc grpc.ClientConnInterface) AuthRClient {
	return &authRClient{cc}
}

func (c *authRClient) AuthR(ctx context.Context, in *AuthRRequest, opts ...grpc.CallOption) (*AuthRResponse, error) {
	out := new(AuthRResponse)
	err := c.cc.Invoke(ctx, AuthR_AuthR_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthRServer is the server API for AuthR service.
// All implementations must embed UnimplementedAuthRServer
// for forward compatibility
type AuthRServer interface {
	AuthR(context.Context, *AuthRRequest) (*AuthRResponse, error)
	mustEmbedUnimplementedAuthRServer()
}

// UnimplementedAuthRServer must be embedded to have forward compatible implementations.
type UnimplementedAuthRServer struct {
}

func (UnimplementedAuthRServer) AuthR(context.Context, *AuthRRequest) (*AuthRResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthR not implemented")
}
func (UnimplementedAuthRServer) mustEmbedUnimplementedAuthRServer() {}

// UnsafeAuthRServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthRServer will
// result in compilation errors.
type UnsafeAuthRServer interface {
	mustEmbedUnimplementedAuthRServer()
}

func RegisterAuthRServer(s grpc.ServiceRegistrar, srv AuthRServer) {
	s.RegisterService(&AuthR_ServiceDesc, srv)
}

func _AuthR_AuthR_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthRServer).AuthR(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthR_AuthR_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthRServer).AuthR(ctx, req.(*AuthRRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthR_ServiceDesc is the grpc.ServiceDesc for AuthR service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthR_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "authr.AuthR",
	HandlerType: (*AuthRServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuthR",
			Handler:    _AuthR_AuthR_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authr/authr.proto",
}
