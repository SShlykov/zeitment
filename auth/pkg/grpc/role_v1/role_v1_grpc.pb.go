// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.0
// source: role_v1.proto

package role_v1

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
	RolesService_List_FullMethodName      = "/role_v1.RolesService/List"
	RolesService_GetByName_FullMethodName = "/role_v1.RolesService/GetByName"
	RolesService_Get_FullMethodName       = "/role_v1.RolesService/Get"
	RolesService_Create_FullMethodName    = "/role_v1.RolesService/Create"
	RolesService_Update_FullMethodName    = "/role_v1.RolesService/Update"
	RolesService_Delete_FullMethodName    = "/role_v1.RolesService/Delete"
)

// RolesServiceClient is the client API for RolesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RolesServiceClient interface {
	// Список ролей
	List(ctx context.Context, in *ListRolesRequest, opts ...grpc.CallOption) (*ListRolesResponse, error)
	// Получить роль по имени
	GetByName(ctx context.Context, in *GetByNameReq, opts ...grpc.CallOption) (*RoleResponse, error)
	// Получить пользователя
	Get(ctx context.Context, in *GetByIdReq, opts ...grpc.CallOption) (*RoleResponse, error)
	// Создать пользователя
	Create(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*RoleResponse, error)
	// Обновить пользователя
	Update(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*RoleResponse, error)
	// Удалить пользователя
	Delete(ctx context.Context, in *DeleteRoleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type rolesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRolesServiceClient(cc grpc.ClientConnInterface) RolesServiceClient {
	return &rolesServiceClient{cc}
}

func (c *rolesServiceClient) List(ctx context.Context, in *ListRolesRequest, opts ...grpc.CallOption) (*ListRolesResponse, error) {
	out := new(ListRolesResponse)
	err := c.cc.Invoke(ctx, RolesService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesServiceClient) GetByName(ctx context.Context, in *GetByNameReq, opts ...grpc.CallOption) (*RoleResponse, error) {
	out := new(RoleResponse)
	err := c.cc.Invoke(ctx, RolesService_GetByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesServiceClient) Get(ctx context.Context, in *GetByIdReq, opts ...grpc.CallOption) (*RoleResponse, error) {
	out := new(RoleResponse)
	err := c.cc.Invoke(ctx, RolesService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesServiceClient) Create(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*RoleResponse, error) {
	out := new(RoleResponse)
	err := c.cc.Invoke(ctx, RolesService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesServiceClient) Update(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*RoleResponse, error) {
	out := new(RoleResponse)
	err := c.cc.Invoke(ctx, RolesService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesServiceClient) Delete(ctx context.Context, in *DeleteRoleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, RolesService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RolesServiceServer is the server API for RolesService service.
// All implementations must embed UnimplementedRolesServiceServer
// for forward compatibility
type RolesServiceServer interface {
	// Список ролей
	List(context.Context, *ListRolesRequest) (*ListRolesResponse, error)
	// Получить роль по имени
	GetByName(context.Context, *GetByNameReq) (*RoleResponse, error)
	// Получить пользователя
	Get(context.Context, *GetByIdReq) (*RoleResponse, error)
	// Создать пользователя
	Create(context.Context, *CreateRoleRequest) (*RoleResponse, error)
	// Обновить пользователя
	Update(context.Context, *UpdateRoleRequest) (*RoleResponse, error)
	// Удалить пользователя
	Delete(context.Context, *DeleteRoleRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedRolesServiceServer()
}

// UnimplementedRolesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRolesServiceServer struct {
}

func (UnimplementedRolesServiceServer) List(context.Context, *ListRolesRequest) (*ListRolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedRolesServiceServer) GetByName(context.Context, *GetByNameReq) (*RoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByName not implemented")
}
func (UnimplementedRolesServiceServer) Get(context.Context, *GetByIdReq) (*RoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedRolesServiceServer) Create(context.Context, *CreateRoleRequest) (*RoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedRolesServiceServer) Update(context.Context, *UpdateRoleRequest) (*RoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedRolesServiceServer) Delete(context.Context, *DeleteRoleRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedRolesServiceServer) mustEmbedUnimplementedRolesServiceServer() {}

// UnsafeRolesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RolesServiceServer will
// result in compilation errors.
type UnsafeRolesServiceServer interface {
	mustEmbedUnimplementedRolesServiceServer()
}

func RegisterRolesServiceServer(s grpc.ServiceRegistrar, srv RolesServiceServer) {
	s.RegisterService(&RolesService_ServiceDesc, srv)
}

func _RolesService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RolesService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServiceServer).List(ctx, req.(*ListRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RolesService_GetByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByNameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServiceServer).GetByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RolesService_GetByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServiceServer).GetByName(ctx, req.(*GetByNameReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RolesService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RolesService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServiceServer).Get(ctx, req.(*GetByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RolesService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RolesService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServiceServer).Create(ctx, req.(*CreateRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RolesService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RolesService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServiceServer).Update(ctx, req.(*UpdateRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RolesService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RolesService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServiceServer).Delete(ctx, req.(*DeleteRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RolesService_ServiceDesc is the grpc.ServiceDesc for RolesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RolesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "role_v1.RolesService",
	HandlerType: (*RolesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _RolesService_List_Handler,
		},
		{
			MethodName: "GetByName",
			Handler:    _RolesService_GetByName_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _RolesService_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _RolesService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _RolesService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _RolesService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "role_v1.proto",
}