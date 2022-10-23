// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: apps/task/pb/task.proto

package task

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
	CreatTask(ctx context.Context, in *CreateTaskRequst, opts ...grpc.CallOption) (*Task, error)
	QueryTask(ctx context.Context, in *QueryTaskRequest, opts ...grpc.CallOption) (*TaskSet, error)
	DescribeTask(ctx context.Context, in *DescribeTaskRequest, opts ...grpc.CallOption) (*Task, error)
	QueryTaskRecord(ctx context.Context, in *QueryTaskRecordRequest, opts ...grpc.CallOption) (*RecordSet, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) CreatTask(ctx context.Context, in *CreateTaskRequst, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := c.cc.Invoke(ctx, "/infraboard.cmdb.task.Service/CreatTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) QueryTask(ctx context.Context, in *QueryTaskRequest, opts ...grpc.CallOption) (*TaskSet, error) {
	out := new(TaskSet)
	err := c.cc.Invoke(ctx, "/infraboard.cmdb.task.Service/QueryTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DescribeTask(ctx context.Context, in *DescribeTaskRequest, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := c.cc.Invoke(ctx, "/infraboard.cmdb.task.Service/DescribeTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) QueryTaskRecord(ctx context.Context, in *QueryTaskRecordRequest, opts ...grpc.CallOption) (*RecordSet, error) {
	out := new(RecordSet)
	err := c.cc.Invoke(ctx, "/infraboard.cmdb.task.Service/QueryTaskRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	CreatTask(context.Context, *CreateTaskRequst) (*Task, error)
	QueryTask(context.Context, *QueryTaskRequest) (*TaskSet, error)
	DescribeTask(context.Context, *DescribeTaskRequest) (*Task, error)
	QueryTaskRecord(context.Context, *QueryTaskRecordRequest) (*RecordSet, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) CreatTask(context.Context, *CreateTaskRequst) (*Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatTask not implemented")
}
func (UnimplementedServiceServer) QueryTask(context.Context, *QueryTaskRequest) (*TaskSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryTask not implemented")
}
func (UnimplementedServiceServer) DescribeTask(context.Context, *DescribeTaskRequest) (*Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeTask not implemented")
}
func (UnimplementedServiceServer) QueryTaskRecord(context.Context, *QueryTaskRecordRequest) (*RecordSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryTaskRecord not implemented")
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

func _Service_CreatTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskRequst)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).CreatTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.cmdb.task.Service/CreatTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).CreatTask(ctx, req.(*CreateTaskRequst))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_QueryTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).QueryTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.cmdb.task.Service/QueryTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).QueryTask(ctx, req.(*QueryTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DescribeTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DescribeTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.cmdb.task.Service/DescribeTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DescribeTask(ctx, req.(*DescribeTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_QueryTaskRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTaskRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).QueryTaskRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.cmdb.task.Service/QueryTaskRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).QueryTaskRecord(ctx, req.(*QueryTaskRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "infraboard.cmdb.task.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatTask",
			Handler:    _Service_CreatTask_Handler,
		},
		{
			MethodName: "QueryTask",
			Handler:    _Service_QueryTask_Handler,
		},
		{
			MethodName: "DescribeTask",
			Handler:    _Service_DescribeTask_Handler,
		},
		{
			MethodName: "QueryTaskRecord",
			Handler:    _Service_QueryTaskRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apps/task/pb/task.proto",
}
