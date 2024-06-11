// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
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
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	MatreshkaBeAPI_ApiVersion_FullMethodName  = "/matreshka_be_api.MatreshkaBeAPI/ApiVersion"
	MatreshkaBeAPI_GetConfig_FullMethodName   = "/matreshka_be_api.MatreshkaBeAPI/GetConfig"
	MatreshkaBeAPI_ListConfigs_FullMethodName = "/matreshka_be_api.MatreshkaBeAPI/ListConfigs"
	MatreshkaBeAPI_PostConfig_FullMethodName  = "/matreshka_be_api.MatreshkaBeAPI/PostConfig"
	MatreshkaBeAPI_PatchConfig_FullMethodName = "/matreshka_be_api.MatreshkaBeAPI/PatchConfig"
)

// MatreshkaBeAPIClient is the client API for MatreshkaBeAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MatreshkaBeAPIClient interface {
	ApiVersion(ctx context.Context, in *ApiVersion_Request, opts ...grpc.CallOption) (*ApiVersion_Response, error)
	GetConfig(ctx context.Context, in *GetConfig_Request, opts ...grpc.CallOption) (*GetConfig_Response, error)
	ListConfigs(ctx context.Context, in *ListConfigs_Request, opts ...grpc.CallOption) (*ListConfigs_Response, error)
	PostConfig(ctx context.Context, in *PostConfig_Request, opts ...grpc.CallOption) (*PostConfig_Response, error)
	PatchConfig(ctx context.Context, in *PatchConfig_Request, opts ...grpc.CallOption) (*PatchConfig_Response, error)
}

type matreshkaBeAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewMatreshkaBeAPIClient(cc grpc.ClientConnInterface) MatreshkaBeAPIClient {
	return &matreshkaBeAPIClient{cc}
}

func (c *matreshkaBeAPIClient) ApiVersion(ctx context.Context, in *ApiVersion_Request, opts ...grpc.CallOption) (*ApiVersion_Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ApiVersion_Response)
	err := c.cc.Invoke(ctx, MatreshkaBeAPI_ApiVersion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matreshkaBeAPIClient) GetConfig(ctx context.Context, in *GetConfig_Request, opts ...grpc.CallOption) (*GetConfig_Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetConfig_Response)
	err := c.cc.Invoke(ctx, MatreshkaBeAPI_GetConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matreshkaBeAPIClient) ListConfigs(ctx context.Context, in *ListConfigs_Request, opts ...grpc.CallOption) (*ListConfigs_Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListConfigs_Response)
	err := c.cc.Invoke(ctx, MatreshkaBeAPI_ListConfigs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matreshkaBeAPIClient) PostConfig(ctx context.Context, in *PostConfig_Request, opts ...grpc.CallOption) (*PostConfig_Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PostConfig_Response)
	err := c.cc.Invoke(ctx, MatreshkaBeAPI_PostConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matreshkaBeAPIClient) PatchConfig(ctx context.Context, in *PatchConfig_Request, opts ...grpc.CallOption) (*PatchConfig_Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PatchConfig_Response)
	err := c.cc.Invoke(ctx, MatreshkaBeAPI_PatchConfig_FullMethodName, in, out, cOpts...)
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
	GetConfig(context.Context, *GetConfig_Request) (*GetConfig_Response, error)
	ListConfigs(context.Context, *ListConfigs_Request) (*ListConfigs_Response, error)
	PostConfig(context.Context, *PostConfig_Request) (*PostConfig_Response, error)
	PatchConfig(context.Context, *PatchConfig_Request) (*PatchConfig_Response, error)
	mustEmbedUnimplementedMatreshkaBeAPIServer()
}

// UnimplementedMatreshkaBeAPIServer must be embedded to have forward compatible implementations.
type UnimplementedMatreshkaBeAPIServer struct {
}

func (UnimplementedMatreshkaBeAPIServer) ApiVersion(context.Context, *ApiVersion_Request) (*ApiVersion_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApiVersion not implemented")
}
func (UnimplementedMatreshkaBeAPIServer) GetConfig(context.Context, *GetConfig_Request) (*GetConfig_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfig not implemented")
}
func (UnimplementedMatreshkaBeAPIServer) ListConfigs(context.Context, *ListConfigs_Request) (*ListConfigs_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListConfigs not implemented")
}
func (UnimplementedMatreshkaBeAPIServer) PostConfig(context.Context, *PostConfig_Request) (*PostConfig_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostConfig not implemented")
}
func (UnimplementedMatreshkaBeAPIServer) PatchConfig(context.Context, *PatchConfig_Request) (*PatchConfig_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchConfig not implemented")
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

func _MatreshkaBeAPI_PostConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostConfig_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatreshkaBeAPIServer).PostConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MatreshkaBeAPI_PostConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatreshkaBeAPIServer).PostConfig(ctx, req.(*PostConfig_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatreshkaBeAPI_PatchConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchConfig_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatreshkaBeAPIServer).PatchConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MatreshkaBeAPI_PatchConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatreshkaBeAPIServer).PatchConfig(ctx, req.(*PatchConfig_Request))
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
			MethodName: "GetConfig",
			Handler:    _MatreshkaBeAPI_GetConfig_Handler,
		},
		{
			MethodName: "ListConfigs",
			Handler:    _MatreshkaBeAPI_ListConfigs_Handler,
		},
		{
			MethodName: "PostConfig",
			Handler:    _MatreshkaBeAPI_PostConfig_Handler,
		},
		{
			MethodName: "PatchConfig",
			Handler:    _MatreshkaBeAPI_PatchConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/matreshka-be_api.proto",
}
