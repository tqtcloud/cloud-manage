// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: apps/disk/pb/disk.proto

package disk

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
	SyncDisk(ctx context.Context, in *Disk, opts ...grpc.CallOption) (*Disk, error)
	QueryDisk(ctx context.Context, in *QueryDiskRequest, opts ...grpc.CallOption) (*DiskSet, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) SyncDisk(ctx context.Context, in *Disk, opts ...grpc.CallOption) (*Disk, error) {
	out := new(Disk)
	err := c.cc.Invoke(ctx, "/infraboard.cmdb.disk.Service/SyncDisk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) QueryDisk(ctx context.Context, in *QueryDiskRequest, opts ...grpc.CallOption) (*DiskSet, error) {
	out := new(DiskSet)
	err := c.cc.Invoke(ctx, "/infraboard.cmdb.disk.Service/QueryDisk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	SyncDisk(context.Context, *Disk) (*Disk, error)
	QueryDisk(context.Context, *QueryDiskRequest) (*DiskSet, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) SyncDisk(context.Context, *Disk) (*Disk, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncDisk not implemented")
}
func (UnimplementedServiceServer) QueryDisk(context.Context, *QueryDiskRequest) (*DiskSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryDisk not implemented")
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

func _Service_SyncDisk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Disk)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).SyncDisk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.cmdb.disk.Service/SyncDisk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).SyncDisk(ctx, req.(*Disk))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_QueryDisk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDiskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).QueryDisk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.cmdb.disk.Service/QueryDisk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).QueryDisk(ctx, req.(*QueryDiskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "infraboard.cmdb.disk.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SyncDisk",
			Handler:    _Service_SyncDisk_Handler,
		},
		{
			MethodName: "QueryDisk",
			Handler:    _Service_QueryDisk_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apps/disk/pb/disk.proto",
}
