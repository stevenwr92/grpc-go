// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.6
// source: proto/user.proto

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

// AddServiceUsersClient is the client API for AddServiceUsers service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AddServiceUsersClient interface {
	GetAllUser(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ResponseGetAllUser, error)
	CreateUser(ctx context.Context, in *CreateUserDto, opts ...grpc.CallOption) (*ResponseCreateUserDto, error)
	UpdateUser(ctx context.Context, in *UpdateUserDto, opts ...grpc.CallOption) (*ResponseUpdateUserDto, error)
	DeleteUser(ctx context.Context, in *DeleteUserDto, opts ...grpc.CallOption) (*ResponseDeleteUserDto, error)
}

type addServiceUsersClient struct {
	cc grpc.ClientConnInterface
}

func NewAddServiceUsersClient(cc grpc.ClientConnInterface) AddServiceUsersClient {
	return &addServiceUsersClient{cc}
}

func (c *addServiceUsersClient) GetAllUser(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ResponseGetAllUser, error) {
	out := new(ResponseGetAllUser)
	err := c.cc.Invoke(ctx, "/user.AddServiceUsers/GetAllUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addServiceUsersClient) CreateUser(ctx context.Context, in *CreateUserDto, opts ...grpc.CallOption) (*ResponseCreateUserDto, error) {
	out := new(ResponseCreateUserDto)
	err := c.cc.Invoke(ctx, "/user.AddServiceUsers/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addServiceUsersClient) UpdateUser(ctx context.Context, in *UpdateUserDto, opts ...grpc.CallOption) (*ResponseUpdateUserDto, error) {
	out := new(ResponseUpdateUserDto)
	err := c.cc.Invoke(ctx, "/user.AddServiceUsers/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addServiceUsersClient) DeleteUser(ctx context.Context, in *DeleteUserDto, opts ...grpc.CallOption) (*ResponseDeleteUserDto, error) {
	out := new(ResponseDeleteUserDto)
	err := c.cc.Invoke(ctx, "/user.AddServiceUsers/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddServiceUsersServer is the server API for AddServiceUsers service.
// All implementations must embed UnimplementedAddServiceUsersServer
// for forward compatibility
type AddServiceUsersServer interface {
	GetAllUser(context.Context, *Empty) (*ResponseGetAllUser, error)
	CreateUser(context.Context, *CreateUserDto) (*ResponseCreateUserDto, error)
	UpdateUser(context.Context, *UpdateUserDto) (*ResponseUpdateUserDto, error)
	DeleteUser(context.Context, *DeleteUserDto) (*ResponseDeleteUserDto, error)
	mustEmbedUnimplementedAddServiceUsersServer()
}

// UnimplementedAddServiceUsersServer must be embedded to have forward compatible implementations.
type UnimplementedAddServiceUsersServer struct {
}

func (UnimplementedAddServiceUsersServer) GetAllUser(context.Context, *Empty) (*ResponseGetAllUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUser not implemented")
}
func (UnimplementedAddServiceUsersServer) CreateUser(context.Context, *CreateUserDto) (*ResponseCreateUserDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedAddServiceUsersServer) UpdateUser(context.Context, *UpdateUserDto) (*ResponseUpdateUserDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedAddServiceUsersServer) DeleteUser(context.Context, *DeleteUserDto) (*ResponseDeleteUserDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedAddServiceUsersServer) mustEmbedUnimplementedAddServiceUsersServer() {}

// UnsafeAddServiceUsersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AddServiceUsersServer will
// result in compilation errors.
type UnsafeAddServiceUsersServer interface {
	mustEmbedUnimplementedAddServiceUsersServer()
}

func RegisterAddServiceUsersServer(s grpc.ServiceRegistrar, srv AddServiceUsersServer) {
	s.RegisterService(&AddServiceUsers_ServiceDesc, srv)
}

func _AddServiceUsers_GetAllUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServiceUsersServer).GetAllUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.AddServiceUsers/GetAllUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServiceUsersServer).GetAllUser(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddServiceUsers_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServiceUsersServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.AddServiceUsers/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServiceUsersServer).CreateUser(ctx, req.(*CreateUserDto))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddServiceUsers_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServiceUsersServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.AddServiceUsers/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServiceUsersServer).UpdateUser(ctx, req.(*UpdateUserDto))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddServiceUsers_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServiceUsersServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.AddServiceUsers/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServiceUsersServer).DeleteUser(ctx, req.(*DeleteUserDto))
	}
	return interceptor(ctx, in, info, handler)
}

// AddServiceUsers_ServiceDesc is the grpc.ServiceDesc for AddServiceUsers service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AddServiceUsers_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.AddServiceUsers",
	HandlerType: (*AddServiceUsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllUser",
			Handler:    _AddServiceUsers_GetAllUser_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _AddServiceUsers_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _AddServiceUsers_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _AddServiceUsers_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/user.proto",
}
