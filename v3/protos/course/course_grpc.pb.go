// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: course/course.proto

package coursepb

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
	CourseSvc_CreateCourse_FullMethodName           = "/course.CourseSvc/CreateCourse"
	CourseSvc_GetCourse_FullMethodName              = "/course.CourseSvc/GetCourse"
	CourseSvc_UpdateCourse_FullMethodName           = "/course.CourseSvc/UpdateCourse"
	CourseSvc_DeleteCourse_FullMethodName           = "/course.CourseSvc/DeleteCourse"
	CourseSvc_DeleteCollectionCourse_FullMethodName = "/course.CourseSvc/DeleteCollectionCourse"
	CourseSvc_ListCourse_FullMethodName             = "/course.CourseSvc/ListCourse"
)

// CourseSvcClient is the client API for CourseSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CourseSvcClient interface {
	CreateCourse(ctx context.Context, in *CreateCourseRequest, opts ...grpc.CallOption) (*general.ResourceId, error)
	GetCourse(ctx context.Context, in *general.GetRequest, opts ...grpc.CallOption) (*Course, error)
	UpdateCourse(ctx context.Context, in *UpdateCourseRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteCourse(ctx context.Context, in *general.ResourceId, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteCollectionCourse(ctx context.Context, in *general.ListOptions, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListCourse(ctx context.Context, in *general.ListOptions, opts ...grpc.CallOption) (*ListCoursesResponse, error)
}

type courseSvcClient struct {
	cc grpc.ClientConnInterface
}

func NewCourseSvcClient(cc grpc.ClientConnInterface) CourseSvcClient {
	return &courseSvcClient{cc}
}

func (c *courseSvcClient) CreateCourse(ctx context.Context, in *CreateCourseRequest, opts ...grpc.CallOption) (*general.ResourceId, error) {
	out := new(general.ResourceId)
	err := c.cc.Invoke(ctx, CourseSvc_CreateCourse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseSvcClient) GetCourse(ctx context.Context, in *general.GetRequest, opts ...grpc.CallOption) (*Course, error) {
	out := new(Course)
	err := c.cc.Invoke(ctx, CourseSvc_GetCourse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseSvcClient) UpdateCourse(ctx context.Context, in *UpdateCourseRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CourseSvc_UpdateCourse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseSvcClient) DeleteCourse(ctx context.Context, in *general.ResourceId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CourseSvc_DeleteCourse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseSvcClient) DeleteCollectionCourse(ctx context.Context, in *general.ListOptions, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CourseSvc_DeleteCollectionCourse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseSvcClient) ListCourse(ctx context.Context, in *general.ListOptions, opts ...grpc.CallOption) (*ListCoursesResponse, error) {
	out := new(ListCoursesResponse)
	err := c.cc.Invoke(ctx, CourseSvc_ListCourse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CourseSvcServer is the server API for CourseSvc service.
// All implementations must embed UnimplementedCourseSvcServer
// for forward compatibility
type CourseSvcServer interface {
	CreateCourse(context.Context, *CreateCourseRequest) (*general.ResourceId, error)
	GetCourse(context.Context, *general.GetRequest) (*Course, error)
	UpdateCourse(context.Context, *UpdateCourseRequest) (*emptypb.Empty, error)
	DeleteCourse(context.Context, *general.ResourceId) (*emptypb.Empty, error)
	DeleteCollectionCourse(context.Context, *general.ListOptions) (*emptypb.Empty, error)
	ListCourse(context.Context, *general.ListOptions) (*ListCoursesResponse, error)
	mustEmbedUnimplementedCourseSvcServer()
}

// UnimplementedCourseSvcServer must be embedded to have forward compatible implementations.
type UnimplementedCourseSvcServer struct {
}

func (UnimplementedCourseSvcServer) CreateCourse(context.Context, *CreateCourseRequest) (*general.ResourceId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCourse not implemented")
}
func (UnimplementedCourseSvcServer) GetCourse(context.Context, *general.GetRequest) (*Course, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourse not implemented")
}
func (UnimplementedCourseSvcServer) UpdateCourse(context.Context, *UpdateCourseRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCourse not implemented")
}
func (UnimplementedCourseSvcServer) DeleteCourse(context.Context, *general.ResourceId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCourse not implemented")
}
func (UnimplementedCourseSvcServer) DeleteCollectionCourse(context.Context, *general.ListOptions) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCollectionCourse not implemented")
}
func (UnimplementedCourseSvcServer) ListCourse(context.Context, *general.ListOptions) (*ListCoursesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCourse not implemented")
}
func (UnimplementedCourseSvcServer) mustEmbedUnimplementedCourseSvcServer() {}

// UnsafeCourseSvcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CourseSvcServer will
// result in compilation errors.
type UnsafeCourseSvcServer interface {
	mustEmbedUnimplementedCourseSvcServer()
}

func RegisterCourseSvcServer(s grpc.ServiceRegistrar, srv CourseSvcServer) {
	s.RegisterService(&CourseSvc_ServiceDesc, srv)
}

func _CourseSvc_CreateCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseSvcServer).CreateCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CourseSvc_CreateCourse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseSvcServer).CreateCourse(ctx, req.(*CreateCourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourseSvc_GetCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(general.GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseSvcServer).GetCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CourseSvc_GetCourse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseSvcServer).GetCourse(ctx, req.(*general.GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourseSvc_UpdateCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseSvcServer).UpdateCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CourseSvc_UpdateCourse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseSvcServer).UpdateCourse(ctx, req.(*UpdateCourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourseSvc_DeleteCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(general.ResourceId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseSvcServer).DeleteCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CourseSvc_DeleteCourse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseSvcServer).DeleteCourse(ctx, req.(*general.ResourceId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourseSvc_DeleteCollectionCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(general.ListOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseSvcServer).DeleteCollectionCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CourseSvc_DeleteCollectionCourse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseSvcServer).DeleteCollectionCourse(ctx, req.(*general.ListOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourseSvc_ListCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(general.ListOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseSvcServer).ListCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CourseSvc_ListCourse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseSvcServer).ListCourse(ctx, req.(*general.ListOptions))
	}
	return interceptor(ctx, in, info, handler)
}

// CourseSvc_ServiceDesc is the grpc.ServiceDesc for CourseSvc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CourseSvc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "course.CourseSvc",
	HandlerType: (*CourseSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCourse",
			Handler:    _CourseSvc_CreateCourse_Handler,
		},
		{
			MethodName: "GetCourse",
			Handler:    _CourseSvc_GetCourse_Handler,
		},
		{
			MethodName: "UpdateCourse",
			Handler:    _CourseSvc_UpdateCourse_Handler,
		},
		{
			MethodName: "DeleteCourse",
			Handler:    _CourseSvc_DeleteCourse_Handler,
		},
		{
			MethodName: "DeleteCollectionCourse",
			Handler:    _CourseSvc_DeleteCollectionCourse_Handler,
		},
		{
			MethodName: "ListCourse",
			Handler:    _CourseSvc_ListCourse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "course/course.proto",
}