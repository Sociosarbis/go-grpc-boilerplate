// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/cmd.proto

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

// CmdServiceClient is the client API for CmdService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CmdServiceClient interface {
	CmdCall(ctx context.Context, in *Cmd, opts ...grpc.CallOption) (CmdService_CmdCallClient, error)
	CmdListFolder(ctx context.Context, in *CmdListFolderReq, opts ...grpc.CallOption) (*CmdListFolderRes, error)
}

type cmdServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCmdServiceClient(cc grpc.ClientConnInterface) CmdServiceClient {
	return &cmdServiceClient{cc}
}

func (c *cmdServiceClient) CmdCall(ctx context.Context, in *Cmd, opts ...grpc.CallOption) (CmdService_CmdCallClient, error) {
	stream, err := c.cc.NewStream(ctx, &CmdService_ServiceDesc.Streams[0], "/proto.CmdService/CmdCall", opts...)
	if err != nil {
		return nil, err
	}
	x := &cmdServiceCmdCallClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CmdService_CmdCallClient interface {
	Recv() (*CmdCallRes, error)
	grpc.ClientStream
}

type cmdServiceCmdCallClient struct {
	grpc.ClientStream
}

func (x *cmdServiceCmdCallClient) Recv() (*CmdCallRes, error) {
	m := new(CmdCallRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *cmdServiceClient) CmdListFolder(ctx context.Context, in *CmdListFolderReq, opts ...grpc.CallOption) (*CmdListFolderRes, error) {
	out := new(CmdListFolderRes)
	err := c.cc.Invoke(ctx, "/proto.CmdService/CmdListFolder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CmdServiceServer is the server API for CmdService service.
// All implementations must embed UnimplementedCmdServiceServer
// for forward compatibility
type CmdServiceServer interface {
	CmdCall(*Cmd, CmdService_CmdCallServer) error
	CmdListFolder(context.Context, *CmdListFolderReq) (*CmdListFolderRes, error)
	mustEmbedUnimplementedCmdServiceServer()
}

// UnimplementedCmdServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCmdServiceServer struct {
}

func (UnimplementedCmdServiceServer) CmdCall(*Cmd, CmdService_CmdCallServer) error {
	return status.Errorf(codes.Unimplemented, "method CmdCall not implemented")
}
func (UnimplementedCmdServiceServer) CmdListFolder(context.Context, *CmdListFolderReq) (*CmdListFolderRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CmdListFolder not implemented")
}
func (UnimplementedCmdServiceServer) mustEmbedUnimplementedCmdServiceServer() {}

// UnsafeCmdServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CmdServiceServer will
// result in compilation errors.
type UnsafeCmdServiceServer interface {
	mustEmbedUnimplementedCmdServiceServer()
}

func RegisterCmdServiceServer(s grpc.ServiceRegistrar, srv CmdServiceServer) {
	s.RegisterService(&CmdService_ServiceDesc, srv)
}

func _CmdService_CmdCall_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Cmd)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CmdServiceServer).CmdCall(m, &cmdServiceCmdCallServer{stream})
}

type CmdService_CmdCallServer interface {
	Send(*CmdCallRes) error
	grpc.ServerStream
}

type cmdServiceCmdCallServer struct {
	grpc.ServerStream
}

func (x *cmdServiceCmdCallServer) Send(m *CmdCallRes) error {
	return x.ServerStream.SendMsg(m)
}

func _CmdService_CmdListFolder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CmdListFolderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdServiceServer).CmdListFolder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CmdService/CmdListFolder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdServiceServer).CmdListFolder(ctx, req.(*CmdListFolderReq))
	}
	return interceptor(ctx, in, info, handler)
}

// CmdService_ServiceDesc is the grpc.ServiceDesc for CmdService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CmdService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CmdService",
	HandlerType: (*CmdServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CmdListFolder",
			Handler:    _CmdService_CmdListFolder_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CmdCall",
			Handler:       _CmdService_CmdCall_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/cmd.proto",
}
