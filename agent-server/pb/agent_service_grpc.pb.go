// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: pb/agent_service.proto

package pb

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

// AgentActionClient is the client API for AgentAction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentActionClient interface {
	UpdateFluentbitHost(ctx context.Context, in *AgentServiceRequest, opts ...grpc.CallOption) (*AgentResponse, error)
}

type agentActionClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentActionClient(cc grpc.ClientConnInterface) AgentActionClient {
	return &agentActionClient{cc}
}

func (c *agentActionClient) UpdateFluentbitHost(ctx context.Context, in *AgentServiceRequest, opts ...grpc.CallOption) (*AgentResponse, error) {
	out := new(AgentResponse)
	err := c.cc.Invoke(ctx, "/pb.AgentAction/UpdateFluentbitHost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentActionServer is the server API for AgentAction service.
// All implementations must embed UnimplementedAgentActionServer
// for forward compatibility
type AgentActionServer interface {
	UpdateFluentbitHost(context.Context, *AgentServiceRequest) (*AgentResponse, error)
	mustEmbedUnimplementedAgentActionServer()
}

// UnimplementedAgentActionServer must be embedded to have forward compatible implementations.
type UnimplementedAgentActionServer struct {
}

func (UnimplementedAgentActionServer) UpdateFluentbitHost(context.Context, *AgentServiceRequest) (*AgentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFluentbitHost not implemented")
}
func (UnimplementedAgentActionServer) mustEmbedUnimplementedAgentActionServer() {}

// UnsafeAgentActionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentActionServer will
// result in compilation errors.
type UnsafeAgentActionServer interface {
	mustEmbedUnimplementedAgentActionServer()
}

func RegisterAgentActionServer(s grpc.ServiceRegistrar, srv AgentActionServer) {
	s.RegisterService(&AgentAction_ServiceDesc, srv)
}

func _AgentAction_UpdateFluentbitHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentActionServer).UpdateFluentbitHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AgentAction/UpdateFluentbitHost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentActionServer).UpdateFluentbitHost(ctx, req.(*AgentServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AgentAction_ServiceDesc is the grpc.ServiceDesc for AgentAction service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AgentAction_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.AgentAction",
	HandlerType: (*AgentActionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateFluentbitHost",
			Handler:    _AgentAction_UpdateFluentbitHost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/agent_service.proto",
}

// AgentFileServiceClient is the client API for AgentFileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentFileServiceClient interface {
	UploadBigFile(ctx context.Context, opts ...grpc.CallOption) (AgentFileService_UploadBigFileClient, error)
}

type agentFileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentFileServiceClient(cc grpc.ClientConnInterface) AgentFileServiceClient {
	return &agentFileServiceClient{cc}
}

func (c *agentFileServiceClient) UploadBigFile(ctx context.Context, opts ...grpc.CallOption) (AgentFileService_UploadBigFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &AgentFileService_ServiceDesc.Streams[0], "/pb.AgentFileService/UploadBigFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &agentFileServiceUploadBigFileClient{stream}
	return x, nil
}

type AgentFileService_UploadBigFileClient interface {
	Send(*AgentFileRequest) error
	CloseAndRecv() (*AgentResponse, error)
	grpc.ClientStream
}

type agentFileServiceUploadBigFileClient struct {
	grpc.ClientStream
}

func (x *agentFileServiceUploadBigFileClient) Send(m *AgentFileRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *agentFileServiceUploadBigFileClient) CloseAndRecv() (*AgentResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AgentResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AgentFileServiceServer is the server API for AgentFileService service.
// All implementations must embed UnimplementedAgentFileServiceServer
// for forward compatibility
type AgentFileServiceServer interface {
	UploadBigFile(AgentFileService_UploadBigFileServer) error
	mustEmbedUnimplementedAgentFileServiceServer()
}

// UnimplementedAgentFileServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAgentFileServiceServer struct {
}

func (UnimplementedAgentFileServiceServer) UploadBigFile(AgentFileService_UploadBigFileServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadBigFile not implemented")
}
func (UnimplementedAgentFileServiceServer) mustEmbedUnimplementedAgentFileServiceServer() {}

// UnsafeAgentFileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentFileServiceServer will
// result in compilation errors.
type UnsafeAgentFileServiceServer interface {
	mustEmbedUnimplementedAgentFileServiceServer()
}

func RegisterAgentFileServiceServer(s grpc.ServiceRegistrar, srv AgentFileServiceServer) {
	s.RegisterService(&AgentFileService_ServiceDesc, srv)
}

func _AgentFileService_UploadBigFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AgentFileServiceServer).UploadBigFile(&agentFileServiceUploadBigFileServer{stream})
}

type AgentFileService_UploadBigFileServer interface {
	SendAndClose(*AgentResponse) error
	Recv() (*AgentFileRequest, error)
	grpc.ServerStream
}

type agentFileServiceUploadBigFileServer struct {
	grpc.ServerStream
}

func (x *agentFileServiceUploadBigFileServer) SendAndClose(m *AgentResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *agentFileServiceUploadBigFileServer) Recv() (*AgentFileRequest, error) {
	m := new(AgentFileRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AgentFileService_ServiceDesc is the grpc.ServiceDesc for AgentFileService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AgentFileService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.AgentFileService",
	HandlerType: (*AgentFileServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadBigFile",
			Handler:       _AgentFileService_UploadBigFile_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "pb/agent_service.proto",
}
