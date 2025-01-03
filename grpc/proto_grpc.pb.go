// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// HomeworkFourServiceClient is the client API for HomeworkFourService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HomeworkFourServiceClient interface {
	SendTokenToNextCLient(ctx context.Context, in *TokenSendRequest, opts ...grpc.CallOption) (*TokenSendResponse, error)
	SendIDToNextClient(ctx context.Context, in *IDSendRequest, opts ...grpc.CallOption) (*IDSendResponse, error)
}

type homeworkFourServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHomeworkFourServiceClient(cc grpc.ClientConnInterface) HomeworkFourServiceClient {
	return &homeworkFourServiceClient{cc}
}

func (c *homeworkFourServiceClient) SendTokenToNextCLient(ctx context.Context, in *TokenSendRequest, opts ...grpc.CallOption) (*TokenSendResponse, error) {
	out := new(TokenSendResponse)
	err := c.cc.Invoke(ctx, "/HomeworkFourService/SendTokenToNextCLient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *homeworkFourServiceClient) SendIDToNextClient(ctx context.Context, in *IDSendRequest, opts ...grpc.CallOption) (*IDSendResponse, error) {
	out := new(IDSendResponse)
	err := c.cc.Invoke(ctx, "/HomeworkFourService/SendIDToNextClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HomeworkFourServiceServer is the server API for HomeworkFourService service.
// All implementations must embed UnimplementedHomeworkFourServiceServer
// for forward compatibility
type HomeworkFourServiceServer interface {
	SendTokenToNextCLient(context.Context, *TokenSendRequest) (*TokenSendResponse, error)
	SendIDToNextClient(context.Context, *IDSendRequest) (*IDSendResponse, error)
	mustEmbedUnimplementedHomeworkFourServiceServer()
}

// UnimplementedHomeworkFourServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHomeworkFourServiceServer struct {
}

func (UnimplementedHomeworkFourServiceServer) SendTokenToNextCLient(context.Context, *TokenSendRequest) (*TokenSendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTokenToNextCLient not implemented")
}
func (UnimplementedHomeworkFourServiceServer) SendIDToNextClient(context.Context, *IDSendRequest) (*IDSendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendIDToNextClient not implemented")
}
func (UnimplementedHomeworkFourServiceServer) mustEmbedUnimplementedHomeworkFourServiceServer() {}

// UnsafeHomeworkFourServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HomeworkFourServiceServer will
// result in compilation errors.
type UnsafeHomeworkFourServiceServer interface {
	mustEmbedUnimplementedHomeworkFourServiceServer()
}

func RegisterHomeworkFourServiceServer(s grpc.ServiceRegistrar, srv HomeworkFourServiceServer) {
	s.RegisterService(&HomeworkFourService_ServiceDesc, srv)
}

func _HomeworkFourService_SendTokenToNextCLient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenSendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeworkFourServiceServer).SendTokenToNextCLient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HomeworkFourService/SendTokenToNextCLient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeworkFourServiceServer).SendTokenToNextCLient(ctx, req.(*TokenSendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HomeworkFourService_SendIDToNextClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDSendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeworkFourServiceServer).SendIDToNextClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HomeworkFourService/SendIDToNextClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeworkFourServiceServer).SendIDToNextClient(ctx, req.(*IDSendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HomeworkFourService_ServiceDesc is the grpc.ServiceDesc for HomeworkFourService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HomeworkFourService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "HomeworkFourService",
	HandlerType: (*HomeworkFourServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendTokenToNextCLient",
			Handler:    _HomeworkFourService_SendTokenToNextCLient_Handler,
		},
		{
			MethodName: "SendIDToNextClient",
			Handler:    _HomeworkFourService_SendIDToNextClient_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/proto.proto",
}
