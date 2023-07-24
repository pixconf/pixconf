// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.4
// source: internal/protos/src/hub_service.proto

package protos

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

// HubServiceClient is the client API for HubService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HubServiceClient interface {
	Subscribe(ctx context.Context, opts ...grpc.CallOption) (HubService_SubscribeClient, error)
}

type hubServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHubServiceClient(cc grpc.ClientConnInterface) HubServiceClient {
	return &hubServiceClient{cc}
}

func (c *hubServiceClient) Subscribe(ctx context.Context, opts ...grpc.CallOption) (HubService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &HubService_ServiceDesc.Streams[0], "/HubService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &hubServiceSubscribeClient{stream}
	return x, nil
}

type HubService_SubscribeClient interface {
	Send(*HubSubscribeAgentResponse) error
	Recv() (*HubSubscribeAgentRequest, error)
	grpc.ClientStream
}

type hubServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *hubServiceSubscribeClient) Send(m *HubSubscribeAgentResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *hubServiceSubscribeClient) Recv() (*HubSubscribeAgentRequest, error) {
	m := new(HubSubscribeAgentRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HubServiceServer is the server API for HubService service.
// All implementations must embed UnimplementedHubServiceServer
// for forward compatibility
type HubServiceServer interface {
	Subscribe(HubService_SubscribeServer) error
	mustEmbedUnimplementedHubServiceServer()
}

// UnimplementedHubServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHubServiceServer struct {
}

func (UnimplementedHubServiceServer) Subscribe(HubService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedHubServiceServer) mustEmbedUnimplementedHubServiceServer() {}

// UnsafeHubServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HubServiceServer will
// result in compilation errors.
type UnsafeHubServiceServer interface {
	mustEmbedUnimplementedHubServiceServer()
}

func RegisterHubServiceServer(s grpc.ServiceRegistrar, srv HubServiceServer) {
	s.RegisterService(&HubService_ServiceDesc, srv)
}

func _HubService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HubServiceServer).Subscribe(&hubServiceSubscribeServer{stream})
}

type HubService_SubscribeServer interface {
	Send(*HubSubscribeAgentRequest) error
	Recv() (*HubSubscribeAgentResponse, error)
	grpc.ServerStream
}

type hubServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *hubServiceSubscribeServer) Send(m *HubSubscribeAgentRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *hubServiceSubscribeServer) Recv() (*HubSubscribeAgentResponse, error) {
	m := new(HubSubscribeAgentResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HubService_ServiceDesc is the grpc.ServiceDesc for HubService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HubService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "HubService",
	HandlerType: (*HubServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _HubService_Subscribe_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "internal/protos/src/hub_service.proto",
}
