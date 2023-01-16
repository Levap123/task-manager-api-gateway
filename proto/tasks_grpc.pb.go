// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: tasks.proto

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

// TaskManagerClient is the client API for TaskManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskManagerClient interface {
	Create(ctx context.Context, in *Task, opts ...grpc.CallOption) (*TaskHelperBody, error)
	Update(ctx context.Context, in *Task, opts ...grpc.CallOption) (*TaskHelperBody, error)
	GetAll(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*TaskArray, error)
	Get(ctx context.Context, in *UserAndTask, opts ...grpc.CallOption) (*Task, error)
	Delete(ctx context.Context, in *TaskHelperBody, opts ...grpc.CallOption) (*TaskHelperBody, error)
}

type taskManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskManagerClient(cc grpc.ClientConnInterface) TaskManagerClient {
	return &taskManagerClient{cc}
}

func (c *taskManagerClient) Create(ctx context.Context, in *Task, opts ...grpc.CallOption) (*TaskHelperBody, error) {
	out := new(TaskHelperBody)
	err := c.cc.Invoke(ctx, "/proto.TaskManager/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskManagerClient) Update(ctx context.Context, in *Task, opts ...grpc.CallOption) (*TaskHelperBody, error) {
	out := new(TaskHelperBody)
	err := c.cc.Invoke(ctx, "/proto.TaskManager/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskManagerClient) GetAll(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*TaskArray, error) {
	out := new(TaskArray)
	err := c.cc.Invoke(ctx, "/proto.TaskManager/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskManagerClient) Get(ctx context.Context, in *UserAndTask, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := c.cc.Invoke(ctx, "/proto.TaskManager/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskManagerClient) Delete(ctx context.Context, in *TaskHelperBody, opts ...grpc.CallOption) (*TaskHelperBody, error) {
	out := new(TaskHelperBody)
	err := c.cc.Invoke(ctx, "/proto.TaskManager/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskManagerServer is the server API for TaskManager service.
// All implementations must embed UnimplementedTaskManagerServer
// for forward compatibility
type TaskManagerServer interface {
	Create(context.Context, *Task) (*TaskHelperBody, error)
	Update(context.Context, *Task) (*TaskHelperBody, error)
	GetAll(context.Context, *UserRequest) (*TaskArray, error)
	Get(context.Context, *UserAndTask) (*Task, error)
	Delete(context.Context, *TaskHelperBody) (*TaskHelperBody, error)
	mustEmbedUnimplementedTaskManagerServer()
}

// UnimplementedTaskManagerServer must be embedded to have forward compatible implementations.
type UnimplementedTaskManagerServer struct {
}

func (UnimplementedTaskManagerServer) Create(context.Context, *Task) (*TaskHelperBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTaskManagerServer) Update(context.Context, *Task) (*TaskHelperBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedTaskManagerServer) GetAll(context.Context, *UserRequest) (*TaskArray, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedTaskManagerServer) Get(context.Context, *UserAndTask) (*Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedTaskManagerServer) Delete(context.Context, *TaskHelperBody) (*TaskHelperBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedTaskManagerServer) mustEmbedUnimplementedTaskManagerServer() {}

// UnsafeTaskManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskManagerServer will
// result in compilation errors.
type UnsafeTaskManagerServer interface {
	mustEmbedUnimplementedTaskManagerServer()
}

func RegisterTaskManagerServer(s grpc.ServiceRegistrar, srv TaskManagerServer) {
	s.RegisterService(&TaskManager_ServiceDesc, srv)
}

func _TaskManager_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Task)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TaskManager/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagerServer).Create(ctx, req.(*Task))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskManager_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Task)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TaskManager/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagerServer).Update(ctx, req.(*Task))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskManager_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagerServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TaskManager/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagerServer).GetAll(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskManager_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAndTask)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TaskManager/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagerServer).Get(ctx, req.(*UserAndTask))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskManager_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskHelperBody)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TaskManager/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagerServer).Delete(ctx, req.(*TaskHelperBody))
	}
	return interceptor(ctx, in, info, handler)
}

// TaskManager_ServiceDesc is the grpc.ServiceDesc for TaskManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaskManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.TaskManager",
	HandlerType: (*TaskManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _TaskManager_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _TaskManager_Update_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _TaskManager_GetAll_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _TaskManager_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TaskManager_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tasks.proto",
}
