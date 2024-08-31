// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.3
// source: pkg/proto/tmp_installer/tmp_installer.proto

package tmp_installer

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	TemporaryDirectoryInstaller_CheckReadiness_FullMethodName            = "/buildbarn.tmp_installer.TemporaryDirectoryInstaller/CheckReadiness"
	TemporaryDirectoryInstaller_InstallTemporaryDirectory_FullMethodName = "/buildbarn.tmp_installer.TemporaryDirectoryInstaller/InstallTemporaryDirectory"
)

// TemporaryDirectoryInstallerClient is the client API for TemporaryDirectoryInstaller service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TemporaryDirectoryInstallerClient interface {
	CheckReadiness(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	InstallTemporaryDirectory(ctx context.Context, in *InstallTemporaryDirectoryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type temporaryDirectoryInstallerClient struct {
	cc grpc.ClientConnInterface
}

func NewTemporaryDirectoryInstallerClient(cc grpc.ClientConnInterface) TemporaryDirectoryInstallerClient {
	return &temporaryDirectoryInstallerClient{cc}
}

func (c *temporaryDirectoryInstallerClient) CheckReadiness(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TemporaryDirectoryInstaller_CheckReadiness_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *temporaryDirectoryInstallerClient) InstallTemporaryDirectory(ctx context.Context, in *InstallTemporaryDirectoryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TemporaryDirectoryInstaller_InstallTemporaryDirectory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TemporaryDirectoryInstallerServer is the server API for TemporaryDirectoryInstaller service.
// All implementations should embed UnimplementedTemporaryDirectoryInstallerServer
// for forward compatibility
type TemporaryDirectoryInstallerServer interface {
	CheckReadiness(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	InstallTemporaryDirectory(context.Context, *InstallTemporaryDirectoryRequest) (*emptypb.Empty, error)
}

// UnimplementedTemporaryDirectoryInstallerServer should be embedded to have forward compatible implementations.
type UnimplementedTemporaryDirectoryInstallerServer struct {
}

func (UnimplementedTemporaryDirectoryInstallerServer) CheckReadiness(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckReadiness not implemented")
}
func (UnimplementedTemporaryDirectoryInstallerServer) InstallTemporaryDirectory(context.Context, *InstallTemporaryDirectoryRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InstallTemporaryDirectory not implemented")
}

// UnsafeTemporaryDirectoryInstallerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TemporaryDirectoryInstallerServer will
// result in compilation errors.
type UnsafeTemporaryDirectoryInstallerServer interface {
	mustEmbedUnimplementedTemporaryDirectoryInstallerServer()
}

func RegisterTemporaryDirectoryInstallerServer(s grpc.ServiceRegistrar, srv TemporaryDirectoryInstallerServer) {
	s.RegisterService(&TemporaryDirectoryInstaller_ServiceDesc, srv)
}

func _TemporaryDirectoryInstaller_CheckReadiness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemporaryDirectoryInstallerServer).CheckReadiness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TemporaryDirectoryInstaller_CheckReadiness_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemporaryDirectoryInstallerServer).CheckReadiness(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TemporaryDirectoryInstaller_InstallTemporaryDirectory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstallTemporaryDirectoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemporaryDirectoryInstallerServer).InstallTemporaryDirectory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TemporaryDirectoryInstaller_InstallTemporaryDirectory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemporaryDirectoryInstallerServer).InstallTemporaryDirectory(ctx, req.(*InstallTemporaryDirectoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TemporaryDirectoryInstaller_ServiceDesc is the grpc.ServiceDesc for TemporaryDirectoryInstaller service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TemporaryDirectoryInstaller_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "buildbarn.tmp_installer.TemporaryDirectoryInstaller",
	HandlerType: (*TemporaryDirectoryInstallerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckReadiness",
			Handler:    _TemporaryDirectoryInstaller_CheckReadiness_Handler,
		},
		{
			MethodName: "InstallTemporaryDirectory",
			Handler:    _TemporaryDirectoryInstaller_InstallTemporaryDirectory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/tmp_installer/tmp_installer.proto",
}
