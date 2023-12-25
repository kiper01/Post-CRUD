// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: proto/post_crud.proto

package Post_CRUD

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
	PostInfoService_AddPostValue_FullMethodName    = "/postcrud.PostInfoService/AddPostValue"
	PostInfoService_RemovePostValue_FullMethodName = "/postcrud.PostInfoService/RemovePostValue"
	PostInfoService_UpdatePostValue_FullMethodName = "/postcrud.PostInfoService/UpdatePostValue"
)

// PostInfoServiceClient is the client API for PostInfoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostInfoServiceClient interface {
	AddPostValue(ctx context.Context, in *AddPostValueRequest, opts ...grpc.CallOption) (*AddPostValueResponse, error)
	RemovePostValue(ctx context.Context, in *RemovePostValueRequest, opts ...grpc.CallOption) (*RemovePostValueResponse, error)
	UpdatePostValue(ctx context.Context, in *UpdatePostValueRequest, opts ...grpc.CallOption) (*UpdatePostValueResponse, error)
}

type postInfoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPostInfoServiceClient(cc grpc.ClientConnInterface) PostInfoServiceClient {
	return &postInfoServiceClient{cc}
}

func (c *postInfoServiceClient) AddPostValue(ctx context.Context, in *AddPostValueRequest, opts ...grpc.CallOption) (*AddPostValueResponse, error) {
	out := new(AddPostValueResponse)
	err := c.cc.Invoke(ctx, PostInfoService_AddPostValue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postInfoServiceClient) RemovePostValue(ctx context.Context, in *RemovePostValueRequest, opts ...grpc.CallOption) (*RemovePostValueResponse, error) {
	out := new(RemovePostValueResponse)
	err := c.cc.Invoke(ctx, PostInfoService_RemovePostValue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postInfoServiceClient) UpdatePostValue(ctx context.Context, in *UpdatePostValueRequest, opts ...grpc.CallOption) (*UpdatePostValueResponse, error) {
	out := new(UpdatePostValueResponse)
	err := c.cc.Invoke(ctx, PostInfoService_UpdatePostValue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostInfoServiceServer is the server API for PostInfoService service.
// All implementations must embed UnimplementedPostInfoServiceServer
// for forward compatibility
type PostInfoServiceServer interface {
	AddPostValue(context.Context, *AddPostValueRequest) (*AddPostValueResponse, error)
	RemovePostValue(context.Context, *RemovePostValueRequest) (*RemovePostValueResponse, error)
	UpdatePostValue(context.Context, *UpdatePostValueRequest) (*UpdatePostValueResponse, error)
	mustEmbedUnimplementedPostInfoServiceServer()
}

// UnimplementedPostInfoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPostInfoServiceServer struct {
}

func (UnimplementedPostInfoServiceServer) AddPostValue(context.Context, *AddPostValueRequest) (*AddPostValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPostValue not implemented")
}
func (UnimplementedPostInfoServiceServer) RemovePostValue(context.Context, *RemovePostValueRequest) (*RemovePostValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemovePostValue not implemented")
}
func (UnimplementedPostInfoServiceServer) UpdatePostValue(context.Context, *UpdatePostValueRequest) (*UpdatePostValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePostValue not implemented")
}
func (UnimplementedPostInfoServiceServer) mustEmbedUnimplementedPostInfoServiceServer() {}

// UnsafePostInfoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostInfoServiceServer will
// result in compilation errors.
type UnsafePostInfoServiceServer interface {
	mustEmbedUnimplementedPostInfoServiceServer()
}

func RegisterPostInfoServiceServer(s grpc.ServiceRegistrar, srv PostInfoServiceServer) {
	s.RegisterService(&PostInfoService_ServiceDesc, srv)
}

func _PostInfoService_AddPostValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPostValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostInfoServiceServer).AddPostValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostInfoService_AddPostValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostInfoServiceServer).AddPostValue(ctx, req.(*AddPostValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostInfoService_RemovePostValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemovePostValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostInfoServiceServer).RemovePostValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostInfoService_RemovePostValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostInfoServiceServer).RemovePostValue(ctx, req.(*RemovePostValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostInfoService_UpdatePostValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePostValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostInfoServiceServer).UpdatePostValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostInfoService_UpdatePostValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostInfoServiceServer).UpdatePostValue(ctx, req.(*UpdatePostValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PostInfoService_ServiceDesc is the grpc.ServiceDesc for PostInfoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PostInfoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "postcrud.PostInfoService",
	HandlerType: (*PostInfoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddPostValue",
			Handler:    _PostInfoService_AddPostValue_Handler,
		},
		{
			MethodName: "RemovePostValue",
			Handler:    _PostInfoService_RemovePostValue_Handler,
		},
		{
			MethodName: "UpdatePostValue",
			Handler:    _PostInfoService_UpdatePostValue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/post_crud.proto",
}
