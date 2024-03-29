// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: query_agent.proto

package query_agent

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

// PromptServiceClient is the client API for PromptService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PromptServiceClient interface {
	GetQuery(ctx context.Context, in *Prompt, opts ...grpc.CallOption) (*Response, error)
}

type promptServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPromptServiceClient(cc grpc.ClientConnInterface) PromptServiceClient {
	return &promptServiceClient{cc}
}

func (c *promptServiceClient) GetQuery(ctx context.Context, in *Prompt, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/query_agent.PromptService/GetQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PromptServiceServer is the server API for PromptService service.
// All implementations must embed UnimplementedPromptServiceServer
// for forward compatibility
type PromptServiceServer interface {
	GetQuery(context.Context, *Prompt) (*Response, error)
	mustEmbedUnimplementedPromptServiceServer()
}

// UnimplementedPromptServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPromptServiceServer struct {
}

func (UnimplementedPromptServiceServer) GetQuery(context.Context, *Prompt) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuery not implemented")
}
func (UnimplementedPromptServiceServer) mustEmbedUnimplementedPromptServiceServer() {}

// UnsafePromptServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PromptServiceServer will
// result in compilation errors.
type UnsafePromptServiceServer interface {
	mustEmbedUnimplementedPromptServiceServer()
}

func RegisterPromptServiceServer(s grpc.ServiceRegistrar, srv PromptServiceServer) {
	s.RegisterService(&PromptService_ServiceDesc, srv)
}

func _PromptService_GetQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Prompt)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromptServiceServer).GetQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/query_agent.PromptService/GetQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromptServiceServer).GetQuery(ctx, req.(*Prompt))
	}
	return interceptor(ctx, in, info, handler)
}

// PromptService_ServiceDesc is the grpc.ServiceDesc for PromptService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PromptService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "query_agent.PromptService",
	HandlerType: (*PromptServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetQuery",
			Handler:    _PromptService_GetQuery_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "query_agent.proto",
}
