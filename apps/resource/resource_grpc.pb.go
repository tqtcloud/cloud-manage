// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: apps/resource/pb/resource.proto

package resource

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

// RPCClient is the client API for RPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RPCClient interface {
	PutResource(ctx context.Context, in *ResourceSet, opts ...grpc.CallOption) (*ResourceSet, error)
	DeleteResource(ctx context.Context, in *ResourceSet, opts ...grpc.CallOption) (*ResourceSet, error)
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*ResourceSet, error)
	QueryTag(ctx context.Context, in *QueryTagRequest, opts ...grpc.CallOption) (*TagSet, error)
	UpdateTag(ctx context.Context, in *UpdateTagRequest, opts ...grpc.CallOption) (*Resource, error)
}

type rPCClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCClient(cc grpc.ClientConnInterface) RPCClient {
	return &rPCClient{cc}
}

func (c *rPCClient) PutResource(ctx context.Context, in *ResourceSet, opts ...grpc.CallOption) (*ResourceSet, error) {
	out := new(ResourceSet)
	err := c.cc.Invoke(ctx, "/infraboard.cmdb.resource.RPC/PutResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) DeleteResource(ctx context.Context, in *ResourceSet, opts ...grpc.CallOption) (*ResourceSet, error) {
	out := new(ResourceSet)
	err := c.cc.Invoke(ctx, "/infraboard.cmdb.resource.RPC/DeleteResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*ResourceSet, error) {
	out := new(ResourceSet)
	err := c.cc.Invoke(ctx, "/infraboard.cmdb.resource.RPC/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) QueryTag(ctx context.Context, in *QueryTagRequest, opts ...grpc.CallOption) (*TagSet, error) {
	out := new(TagSet)
	err := c.cc.Invoke(ctx, "/infraboard.cmdb.resource.RPC/QueryTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) UpdateTag(ctx context.Context, in *UpdateTagRequest, opts ...grpc.CallOption) (*Resource, error) {
	out := new(Resource)
	err := c.cc.Invoke(ctx, "/infraboard.cmdb.resource.RPC/UpdateTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCServer is the server API for RPC service.
// All implementations must embed UnimplementedRPCServer
// for forward compatibility
type RPCServer interface {
	PutResource(context.Context, *ResourceSet) (*ResourceSet, error)
	DeleteResource(context.Context, *ResourceSet) (*ResourceSet, error)
	Search(context.Context, *SearchRequest) (*ResourceSet, error)
	QueryTag(context.Context, *QueryTagRequest) (*TagSet, error)
	UpdateTag(context.Context, *UpdateTagRequest) (*Resource, error)
	mustEmbedUnimplementedRPCServer()
}

// UnimplementedRPCServer must be embedded to have forward compatible implementations.
type UnimplementedRPCServer struct {
}

func (UnimplementedRPCServer) PutResource(context.Context, *ResourceSet) (*ResourceSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutResource not implemented")
}
func (UnimplementedRPCServer) DeleteResource(context.Context, *ResourceSet) (*ResourceSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteResource not implemented")
}
func (UnimplementedRPCServer) Search(context.Context, *SearchRequest) (*ResourceSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedRPCServer) QueryTag(context.Context, *QueryTagRequest) (*TagSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryTag not implemented")
}
func (UnimplementedRPCServer) UpdateTag(context.Context, *UpdateTagRequest) (*Resource, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTag not implemented")
}
func (UnimplementedRPCServer) mustEmbedUnimplementedRPCServer() {}

// UnsafeRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RPCServer will
// result in compilation errors.
type UnsafeRPCServer interface {
	mustEmbedUnimplementedRPCServer()
}

func RegisterRPCServer(s grpc.ServiceRegistrar, srv RPCServer) {
	s.RegisterService(&RPC_ServiceDesc, srv)
}

func _RPC_PutResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceSet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).PutResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.cmdb.resource.RPC/PutResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).PutResource(ctx, req.(*ResourceSet))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_DeleteResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceSet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).DeleteResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.cmdb.resource.RPC/DeleteResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).DeleteResource(ctx, req.(*ResourceSet))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.cmdb.resource.RPC/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_QueryTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).QueryTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.cmdb.resource.RPC/QueryTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).QueryTag(ctx, req.(*QueryTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_UpdateTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).UpdateTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.cmdb.resource.RPC/UpdateTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).UpdateTag(ctx, req.(*UpdateTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RPC_ServiceDesc is the grpc.ServiceDesc for RPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "infraboard.cmdb.resource.RPC",
	HandlerType: (*RPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PutResource",
			Handler:    _RPC_PutResource_Handler,
		},
		{
			MethodName: "DeleteResource",
			Handler:    _RPC_DeleteResource_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _RPC_Search_Handler,
		},
		{
			MethodName: "QueryTag",
			Handler:    _RPC_QueryTag_Handler,
		},
		{
			MethodName: "UpdateTag",
			Handler:    _RPC_UpdateTag_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apps/resource/pb/resource.proto",
}