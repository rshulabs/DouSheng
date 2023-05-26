// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: apps/comment/pb/comment.proto

package comment

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

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	// 发表评论
	CommentAction(ctx context.Context, in *CommentActionRequest, opts ...grpc.CallOption) (*CommentActionResponse, error)
	// 获取评论
	GetCommentList(ctx context.Context, in *GetCommentListRequest, opts ...grpc.CallOption) (*GetCommentListResponse, error)
	// 通过视频Id获取评论总数
	GetCommentCountById(ctx context.Context, in *GetCommentCountByIdRequest, opts ...grpc.CallOption) (*GetCommentCountByIdResponse, error)
	// 获取视频评论数量 map[videoId] = comment_count
	CommentCountMap(ctx context.Context, in *CommentMapRequest, opts ...grpc.CallOption) (*CommentMapResponse, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) CommentAction(ctx context.Context, in *CommentActionRequest, opts ...grpc.CallOption) (*CommentActionResponse, error) {
	out := new(CommentActionResponse)
	err := c.cc.Invoke(ctx, "/dousheng.interaction.comment.Service/CommentAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetCommentList(ctx context.Context, in *GetCommentListRequest, opts ...grpc.CallOption) (*GetCommentListResponse, error) {
	out := new(GetCommentListResponse)
	err := c.cc.Invoke(ctx, "/dousheng.interaction.comment.Service/GetCommentList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetCommentCountById(ctx context.Context, in *GetCommentCountByIdRequest, opts ...grpc.CallOption) (*GetCommentCountByIdResponse, error) {
	out := new(GetCommentCountByIdResponse)
	err := c.cc.Invoke(ctx, "/dousheng.interaction.comment.Service/GetCommentCountById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) CommentCountMap(ctx context.Context, in *CommentMapRequest, opts ...grpc.CallOption) (*CommentMapResponse, error) {
	out := new(CommentMapResponse)
	err := c.cc.Invoke(ctx, "/dousheng.interaction.comment.Service/CommentCountMap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	// 发表评论
	CommentAction(context.Context, *CommentActionRequest) (*CommentActionResponse, error)
	// 获取评论
	GetCommentList(context.Context, *GetCommentListRequest) (*GetCommentListResponse, error)
	// 通过视频Id获取评论总数
	GetCommentCountById(context.Context, *GetCommentCountByIdRequest) (*GetCommentCountByIdResponse, error)
	// 获取视频评论数量 map[videoId] = comment_count
	CommentCountMap(context.Context, *CommentMapRequest) (*CommentMapResponse, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) CommentAction(context.Context, *CommentActionRequest) (*CommentActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentAction not implemented")
}
func (UnimplementedServiceServer) GetCommentList(context.Context, *GetCommentListRequest) (*GetCommentListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentList not implemented")
}
func (UnimplementedServiceServer) GetCommentCountById(context.Context, *GetCommentCountByIdRequest) (*GetCommentCountByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentCountById not implemented")
}
func (UnimplementedServiceServer) CommentCountMap(context.Context, *CommentMapRequest) (*CommentMapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentCountMap not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_CommentAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).CommentAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dousheng.interaction.comment.Service/CommentAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).CommentAction(ctx, req.(*CommentActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetCommentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetCommentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dousheng.interaction.comment.Service/GetCommentList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetCommentList(ctx, req.(*GetCommentListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetCommentCountById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentCountByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetCommentCountById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dousheng.interaction.comment.Service/GetCommentCountById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetCommentCountById(ctx, req.(*GetCommentCountByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_CommentCountMap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentMapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).CommentCountMap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dousheng.interaction.comment.Service/CommentCountMap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).CommentCountMap(ctx, req.(*CommentMapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dousheng.interaction.comment.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CommentAction",
			Handler:    _Service_CommentAction_Handler,
		},
		{
			MethodName: "GetCommentList",
			Handler:    _Service_GetCommentList_Handler,
		},
		{
			MethodName: "GetCommentCountById",
			Handler:    _Service_GetCommentCountById_Handler,
		},
		{
			MethodName: "CommentCountMap",
			Handler:    _Service_CommentCountMap_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apps/comment/pb/comment.proto",
}