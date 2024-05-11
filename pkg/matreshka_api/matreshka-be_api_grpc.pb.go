// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: grpc/matreshka-be_api.proto

package matreshka_api

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
	MatreshkaBeAPI_ApiVersion_FullMethodName     = "/matreshka_be_api.MatreshkaBeAPI/ApiVersion"
	MatreshkaBeAPI_UpsertConfig_FullMethodName   = "/matreshka_be_api.MatreshkaBeAPI/UpsertConfig"
	MatreshkaBeAPI_GetConfig_FullMethodName      = "/matreshka_be_api.MatreshkaBeAPI/GetConfig"
	MatreshkaBeAPI_GetConfigRaw_FullMethodName   = "/matreshka_be_api.MatreshkaBeAPI/GetConfigRaw"
	MatreshkaBeAPI_PatchConfigRaw_FullMethodName = "/matreshka_be_api.MatreshkaBeAPI/PatchConfigRaw"
	MatreshkaBeAPI_ListConfigs_FullMethodName    = "/matreshka_be_api.MatreshkaBeAPI/ListConfigs"
)

// MatreshkaBeAPIClient is the client API for MatreshkaBeAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MatreshkaBeAPIClient interface {
	ApiVersion(ctx context.Context, in *ApiVersion_Request, opts ...grpc.CallOption) (*ApiVersion_Response, error)
	UpsertConfig(ctx context.Context, in *PatchConfig_Request, opts ...grpc.CallOption) (*PatchConfig_Response, error)
	GetConfig(ctx context.Context, in *GetConfig_Request, opts ...grpc.CallOption) (*GetConfig_Response, error)
	GetConfigRaw(ctx context.Context, in *GetConfigRaw_Request, opts ...grpc.CallOption) (*GetConfigRaw_Response, error)
	PatchConfigRaw(ctx context.Context, in *PatchConfigRaw_Request, opts ...grpc.CallOption) (*PatchConfigRaw_Response, error)
	ListConfigs(ctx context.Context, in *ListConfigs_Request, opts ...grpc.CallOption) (*ListConfigs_Response, error)
}

type matreshkaBeAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewMatreshkaBeAPIClient(cc grpc.ClientConnInterface) MatreshkaBeAPIClient {
	return &matreshkaBeAPIClient{cc}
}

func (c *matreshkaBeAPIClient) ApiVersion(ctx context.Context, in *ApiVersion_Request, opts ...grpc.CallOption) (*ApiVersion_Response, error) {
	out := new(ApiVersion_Response)
	err := c.cc.Invoke(ctx, MatreshkaBeAPI_ApiVersion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matreshkaBeAPIClient) UpsertConfig(ctx context.Context, in *PatchConfig_Request, opts ...grpc.CallOption) (*PatchConfig_Response, error) {
	out := new(PatchConfig_Response)
	err := c.cc.Invoke(ctx, MatreshkaBeAPI_UpsertConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matreshkaBeAPIClient) GetConfig(ctx context.Context, in *GetConfig_Request, opts ...grpc.CallOption) (*GetConfig_Response, error) {
	out := new(GetConfig_Response)
	err := c.cc.Invoke(ctx, MatreshkaBeAPI_GetConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matreshkaBeAPIClient) GetConfigRaw(ctx context.Context, in *GetConfigRaw_Request, opts ...grpc.CallOption) (*GetConfigRaw_Response, error) {
	out := new(GetConfigRaw_Response)
	err := c.cc.Invoke(ctx, MatreshkaBeAPI_GetConfigRaw_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matreshkaBeAPIClient) PatchConfigRaw(ctx context.Context, in *PatchConfigRaw_Request, opts ...grpc.CallOption) (*PatchConfigRaw_Response, error) {
	out := new(PatchConfigRaw_Response)
	err := c.cc.Invoke(ctx, MatreshkaBeAPI_PatchConfigRaw_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matreshkaBeAPIClient) ListConfigs(ctx context.Context, in *ListConfigs_Request, opts ...grpc.CallOption) (*ListConfigs_Response, error) {
	out := new(ListConfigs_Response)
	err := c.cc.Invoke(ctx, MatreshkaBeAPI_ListConfigs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MatreshkaBeAPIServer is the server API for MatreshkaBeAPI service.
// All implementations must embed UnimplementedMatreshkaBeAPIServer
// for forward compatibility
type MatreshkaBeAPIServer interface {
	ApiVersion(context.Context, *ApiVersion_Request) (*ApiVersion_Response, error)
	UpsertConfig(context.Context, *PatchConfig_Request) (*PatchConfig_Response, error)
	GetConfig(context.Context, *GetConfig_Request) (*GetConfig_Response, error)
	GetConfigRaw(context.Context, *GetConfigRaw_Request) (*GetConfigRaw_Response, error)
	PatchConfigRaw(context.Context, *PatchConfigRaw_Request) (*PatchConfigRaw_Response, error)
	ListConfigs(context.Context, *ListConfigs_Request) (*ListConfigs_Response, error)
	mustEmbedUnimplementedMatreshkaBeAPIServer()
}

// UnimplementedMatreshkaBeAPIServer must be embedded to have forward compatible implementations.
type UnimplementedMatreshkaBeAPIServer struct {
}

func (UnimplementedMatreshkaBeAPIServer) ApiVersion(context.Context, *ApiVersion_Request) (*ApiVersion_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApiVersion not implemented")
}
func (UnimplementedMatreshkaBeAPIServer) UpsertConfig(context.Context, *PatchConfig_Request) (*PatchConfig_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertConfig not implemented")
}
func (UnimplementedMatreshkaBeAPIServer) GetConfig(context.Context, *GetConfig_Request) (*GetConfig_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfig not implemented")
}
func (UnimplementedMatreshkaBeAPIServer) GetConfigRaw(context.Context, *GetConfigRaw_Request) (*GetConfigRaw_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfigRaw not implemented")
}
func (UnimplementedMatreshkaBeAPIServer) PatchConfigRaw(context.Context, *PatchConfigRaw_Request) (*PatchConfigRaw_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchConfigRaw not implemented")
}
func (UnimplementedMatreshkaBeAPIServer) ListConfigs(context.Context, *ListConfigs_Request) (*ListConfigs_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListConfigs not implemented")
}
func (UnimplementedMatreshkaBeAPIServer) mustEmbedUnimplementedMatreshkaBeAPIServer() {}

// UnsafeMatreshkaBeAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MatreshkaBeAPIServer will
// result in compilation errors.
type UnsafeMatreshkaBeAPIServer interface {
	mustEmbedUnimplementedMatreshkaBeAPIServer()
}

func RegisterMatreshkaBeAPIServer(s grpc.ServiceRegistrar, srv MatreshkaBeAPIServer) {
	s.RegisterService(&MatreshkaBeAPI_ServiceDesc, srv)
}

func _MatreshkaBeAPI_ApiVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApiVersion_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatreshkaBeAPIServer).ApiVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MatreshkaBeAPI_ApiVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatreshkaBeAPIServer).ApiVersion(ctx, req.(*ApiVersion_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatreshkaBeAPI_UpsertConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchConfig_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatreshkaBeAPIServer).UpsertConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MatreshkaBeAPI_UpsertConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatreshkaBeAPIServer).UpsertConfig(ctx, req.(*PatchConfig_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatreshkaBeAPI_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfig_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatreshkaBeAPIServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MatreshkaBeAPI_GetConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatreshkaBeAPIServer).GetConfig(ctx, req.(*GetConfig_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatreshkaBeAPI_GetConfigRaw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigRaw_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatreshkaBeAPIServer).GetConfigRaw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MatreshkaBeAPI_GetConfigRaw_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatreshkaBeAPIServer).GetConfigRaw(ctx, req.(*GetConfigRaw_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatreshkaBeAPI_PatchConfigRaw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchConfigRaw_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatreshkaBeAPIServer).PatchConfigRaw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MatreshkaBeAPI_PatchConfigRaw_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatreshkaBeAPIServer).PatchConfigRaw(ctx, req.(*PatchConfigRaw_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatreshkaBeAPI_ListConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListConfigs_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatreshkaBeAPIServer).ListConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MatreshkaBeAPI_ListConfigs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatreshkaBeAPIServer).ListConfigs(ctx, req.(*ListConfigs_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// MatreshkaBeAPI_ServiceDesc is the grpc.ServiceDesc for MatreshkaBeAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MatreshkaBeAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "matreshka_be_api.MatreshkaBeAPI",
	HandlerType: (*MatreshkaBeAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ApiVersion",
			Handler:    _MatreshkaBeAPI_ApiVersion_Handler,
		},
		{
			MethodName: "UpsertConfig",
			Handler:    _MatreshkaBeAPI_UpsertConfig_Handler,
		},
		{
			MethodName: "GetConfig",
			Handler:    _MatreshkaBeAPI_GetConfig_Handler,
		},
		{
			MethodName: "GetConfigRaw",
			Handler:    _MatreshkaBeAPI_GetConfigRaw_Handler,
		},
		{
			MethodName: "PatchConfigRaw",
			Handler:    _MatreshkaBeAPI_PatchConfigRaw_Handler,
		},
		{
			MethodName: "ListConfigs",
			Handler:    _MatreshkaBeAPI_ListConfigs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/matreshka-be_api.proto",
}
