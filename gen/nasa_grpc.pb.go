// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: nasa.proto

package nasav1

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

// NasaClient is the client API for Nasa service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NasaClient interface {
	RandomSpaseImage(ctx context.Context, in *RandomSpaseImageRequest, opts ...grpc.CallOption) (Nasa_RandomSpaseImageClient, error)
}

type nasaClient struct {
	cc grpc.ClientConnInterface
}

func NewNasaClient(cc grpc.ClientConnInterface) NasaClient {
	return &nasaClient{cc}
}

func (c *nasaClient) RandomSpaseImage(ctx context.Context, in *RandomSpaseImageRequest, opts ...grpc.CallOption) (Nasa_RandomSpaseImageClient, error) {
	stream, err := c.cc.NewStream(ctx, &Nasa_ServiceDesc.Streams[0], "/nasa.Nasa/RandomSpaseImage", opts...)
	if err != nil {
		return nil, err
	}
	x := &nasaRandomSpaseImageClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Nasa_RandomSpaseImageClient interface {
	Recv() (*RandomSpaseImageResponse, error)
	grpc.ClientStream
}

type nasaRandomSpaseImageClient struct {
	grpc.ClientStream
}

func (x *nasaRandomSpaseImageClient) Recv() (*RandomSpaseImageResponse, error) {
	m := new(RandomSpaseImageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NasaServer is the server API for Nasa service.
// All implementations must embed UnimplementedNasaServer
// for forward compatibility
type NasaServer interface {
	RandomSpaseImage(*RandomSpaseImageRequest, Nasa_RandomSpaseImageServer) error
	mustEmbedUnimplementedNasaServer()
}

// UnimplementedNasaServer must be embedded to have forward compatible implementations.
type UnimplementedNasaServer struct {
}

func (UnimplementedNasaServer) RandomSpaseImage(*RandomSpaseImageRequest, Nasa_RandomSpaseImageServer) error {
	return status.Errorf(codes.Unimplemented, "method RandomSpaseImage not implemented")
}
func (UnimplementedNasaServer) mustEmbedUnimplementedNasaServer() {}

// UnsafeNasaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NasaServer will
// result in compilation errors.
type UnsafeNasaServer interface {
	mustEmbedUnimplementedNasaServer()
}

func RegisterNasaServer(s grpc.ServiceRegistrar, srv NasaServer) {
	s.RegisterService(&Nasa_ServiceDesc, srv)
}

func _Nasa_RandomSpaseImage_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RandomSpaseImageRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NasaServer).RandomSpaseImage(m, &nasaRandomSpaseImageServer{stream})
}

type Nasa_RandomSpaseImageServer interface {
	Send(*RandomSpaseImageResponse) error
	grpc.ServerStream
}

type nasaRandomSpaseImageServer struct {
	grpc.ServerStream
}

func (x *nasaRandomSpaseImageServer) Send(m *RandomSpaseImageResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Nasa_ServiceDesc is the grpc.ServiceDesc for Nasa service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Nasa_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nasa.Nasa",
	HandlerType: (*NasaServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RandomSpaseImage",
			Handler:       _Nasa_RandomSpaseImage_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "nasa.proto",
}
