// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: notif.proto

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

const (
	Notif_StreamMeow_FullMethodName = "/Notif/StreamMeow"
)

// NotifClient is the client API for Notif service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotifClient interface {
	StreamMeow(ctx context.Context, in *MeowReq, opts ...grpc.CallOption) (Notif_StreamMeowClient, error)
}

type notifClient struct {
	cc grpc.ClientConnInterface
}

func NewNotifClient(cc grpc.ClientConnInterface) NotifClient {
	return &notifClient{cc}
}

func (c *notifClient) StreamMeow(ctx context.Context, in *MeowReq, opts ...grpc.CallOption) (Notif_StreamMeowClient, error) {
	stream, err := c.cc.NewStream(ctx, &Notif_ServiceDesc.Streams[0], Notif_StreamMeow_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &notifStreamMeowClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Notif_StreamMeowClient interface {
	Recv() (*MeowRes, error)
	grpc.ClientStream
}

type notifStreamMeowClient struct {
	grpc.ClientStream
}

func (x *notifStreamMeowClient) Recv() (*MeowRes, error) {
	m := new(MeowRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NotifServer is the server API for Notif service.
// All implementations must embed UnimplementedNotifServer
// for forward compatibility
type NotifServer interface {
	StreamMeow(*MeowReq, Notif_StreamMeowServer) error
	mustEmbedUnimplementedNotifServer()
}

// UnimplementedNotifServer must be embedded to have forward compatible implementations.
type UnimplementedNotifServer struct {
}

func (UnimplementedNotifServer) StreamMeow(*MeowReq, Notif_StreamMeowServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamMeow not implemented")
}
func (UnimplementedNotifServer) mustEmbedUnimplementedNotifServer() {}

// UnsafeNotifServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotifServer will
// result in compilation errors.
type UnsafeNotifServer interface {
	mustEmbedUnimplementedNotifServer()
}

func RegisterNotifServer(s grpc.ServiceRegistrar, srv NotifServer) {
	s.RegisterService(&Notif_ServiceDesc, srv)
}

func _Notif_StreamMeow_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MeowReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NotifServer).StreamMeow(m, &notifStreamMeowServer{stream})
}

type Notif_StreamMeowServer interface {
	Send(*MeowRes) error
	grpc.ServerStream
}

type notifStreamMeowServer struct {
	grpc.ServerStream
}

func (x *notifStreamMeowServer) Send(m *MeowRes) error {
	return x.ServerStream.SendMsg(m)
}

// Notif_ServiceDesc is the grpc.ServiceDesc for Notif service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Notif_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Notif",
	HandlerType: (*NotifServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamMeow",
			Handler:       _Notif_StreamMeow_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "notif.proto",
}
