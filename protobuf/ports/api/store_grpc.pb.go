// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: ports/api/store.proto

package api

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

// PortsClient is the client API for Ports service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PortsClient interface {
	StreamPorts(ctx context.Context, opts ...grpc.CallOption) (Ports_StreamPortsClient, error)
}

type portsClient struct {
	cc grpc.ClientConnInterface
}

func NewPortsClient(cc grpc.ClientConnInterface) PortsClient {
	return &portsClient{cc}
}

func (c *portsClient) StreamPorts(ctx context.Context, opts ...grpc.CallOption) (Ports_StreamPortsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Ports_ServiceDesc.Streams[0], "/ports.api.Ports/StreamPorts", opts...)
	if err != nil {
		return nil, err
	}
	x := &portsStreamPortsClient{stream}
	return x, nil
}

type Ports_StreamPortsClient interface {
	Send(*StreamPortsRequest) error
	CloseAndRecv() (*StreamPortsResponse, error)
	grpc.ClientStream
}

type portsStreamPortsClient struct {
	grpc.ClientStream
}

func (x *portsStreamPortsClient) Send(m *StreamPortsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *portsStreamPortsClient) CloseAndRecv() (*StreamPortsResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamPortsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PortsServer is the server API for Ports service.
// All implementations must embed UnimplementedPortsServer
// for forward compatibility
type PortsServer interface {
	StreamPorts(Ports_StreamPortsServer) error
	mustEmbedUnimplementedPortsServer()
}

// UnimplementedPortsServer must be embedded to have forward compatible implementations.
type UnimplementedPortsServer struct {
}

func (UnimplementedPortsServer) StreamPorts(Ports_StreamPortsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamPorts not implemented")
}
func (UnimplementedPortsServer) mustEmbedUnimplementedPortsServer() {}

// UnsafePortsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PortsServer will
// result in compilation errors.
type UnsafePortsServer interface {
	mustEmbedUnimplementedPortsServer()
}

func RegisterPortsServer(s grpc.ServiceRegistrar, srv PortsServer) {
	s.RegisterService(&Ports_ServiceDesc, srv)
}

func _Ports_StreamPorts_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PortsServer).StreamPorts(&portsStreamPortsServer{stream})
}

type Ports_StreamPortsServer interface {
	SendAndClose(*StreamPortsResponse) error
	Recv() (*StreamPortsRequest, error)
	grpc.ServerStream
}

type portsStreamPortsServer struct {
	grpc.ServerStream
}

func (x *portsStreamPortsServer) SendAndClose(m *StreamPortsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *portsStreamPortsServer) Recv() (*StreamPortsRequest, error) {
	m := new(StreamPortsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Ports_ServiceDesc is the grpc.ServiceDesc for Ports service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Ports_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ports.api.Ports",
	HandlerType: (*PortsServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamPorts",
			Handler:       _Ports_StreamPorts_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "ports/api/store.proto",
}
