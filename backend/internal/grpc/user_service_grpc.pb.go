// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	Signup(ctx context.Context, in *SignupParams, opts ...grpc.CallOption) (*AuthToken, error)
	Dropout(ctx context.Context, in *AuthToken, opts ...grpc.CallOption) (*Result, error)
	Login(ctx context.Context, in *LoginParams, opts ...grpc.CallOption) (*TokenWithUrl, error)
	Logout(ctx context.Context, in *AuthToken, opts ...grpc.CallOption) (*Result, error)
	RegisterService(ctx context.Context, in *ServiceName, opts ...grpc.CallOption) (*AuthToken, error)
	UnregisterService(ctx context.Context, in *AuthToken, opts ...grpc.CallOption) (*Result, error)
	Authenticate(ctx context.Context, in *TokenWithPermission, opts ...grpc.CallOption) (*Result, error)
	Authorize(ctx context.Context, in *AuthorizeParams, opts ...grpc.CallOption) (*Result, error)
	RegisterPermission(ctx context.Context, in *TokenWithPermission, opts ...grpc.CallOption) (*Result, error)
	UnregisterPermission(ctx context.Context, in *TokenWithPermission, opts ...grpc.CallOption) (*Result, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Signup(ctx context.Context, in *SignupParams, opts ...grpc.CallOption) (*AuthToken, error) {
	out := new(AuthToken)
	err := c.cc.Invoke(ctx, "/grpc.UserService/Signup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Dropout(ctx context.Context, in *AuthToken, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/grpc.UserService/Dropout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Login(ctx context.Context, in *LoginParams, opts ...grpc.CallOption) (*TokenWithUrl, error) {
	out := new(TokenWithUrl)
	err := c.cc.Invoke(ctx, "/grpc.UserService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Logout(ctx context.Context, in *AuthToken, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/grpc.UserService/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) RegisterService(ctx context.Context, in *ServiceName, opts ...grpc.CallOption) (*AuthToken, error) {
	out := new(AuthToken)
	err := c.cc.Invoke(ctx, "/grpc.UserService/RegisterService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UnregisterService(ctx context.Context, in *AuthToken, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/grpc.UserService/UnregisterService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Authenticate(ctx context.Context, in *TokenWithPermission, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/grpc.UserService/Authenticate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Authorize(ctx context.Context, in *AuthorizeParams, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/grpc.UserService/Authorize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) RegisterPermission(ctx context.Context, in *TokenWithPermission, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/grpc.UserService/RegisterPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UnregisterPermission(ctx context.Context, in *TokenWithPermission, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/grpc.UserService/UnregisterPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	Signup(context.Context, *SignupParams) (*AuthToken, error)
	Dropout(context.Context, *AuthToken) (*Result, error)
	Login(context.Context, *LoginParams) (*TokenWithUrl, error)
	Logout(context.Context, *AuthToken) (*Result, error)
	RegisterService(context.Context, *ServiceName) (*AuthToken, error)
	UnregisterService(context.Context, *AuthToken) (*Result, error)
	Authenticate(context.Context, *TokenWithPermission) (*Result, error)
	Authorize(context.Context, *AuthorizeParams) (*Result, error)
	RegisterPermission(context.Context, *TokenWithPermission) (*Result, error)
	UnregisterPermission(context.Context, *TokenWithPermission) (*Result, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) Signup(context.Context, *SignupParams) (*AuthToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Signup not implemented")
}
func (UnimplementedUserServiceServer) Dropout(context.Context, *AuthToken) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Dropout not implemented")
}
func (UnimplementedUserServiceServer) Login(context.Context, *LoginParams) (*TokenWithUrl, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServiceServer) Logout(context.Context, *AuthToken) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedUserServiceServer) RegisterService(context.Context, *ServiceName) (*AuthToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterService not implemented")
}
func (UnimplementedUserServiceServer) UnregisterService(context.Context, *AuthToken) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnregisterService not implemented")
}
func (UnimplementedUserServiceServer) Authenticate(context.Context, *TokenWithPermission) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}
func (UnimplementedUserServiceServer) Authorize(context.Context, *AuthorizeParams) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authorize not implemented")
}
func (UnimplementedUserServiceServer) RegisterPermission(context.Context, *TokenWithPermission) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterPermission not implemented")
}
func (UnimplementedUserServiceServer) UnregisterPermission(context.Context, *TokenWithPermission) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnregisterPermission not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_Signup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignupParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Signup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.UserService/Signup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Signup(ctx, req.(*SignupParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Dropout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Dropout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.UserService/Dropout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Dropout(ctx, req.(*AuthToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.UserService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Login(ctx, req.(*LoginParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.UserService/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Logout(ctx, req.(*AuthToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_RegisterService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).RegisterService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.UserService/RegisterService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).RegisterService(ctx, req.(*ServiceName))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UnregisterService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UnregisterService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.UserService/UnregisterService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UnregisterService(ctx, req.(*AuthToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenWithPermission)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.UserService/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Authenticate(ctx, req.(*TokenWithPermission))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Authorize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Authorize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.UserService/Authorize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Authorize(ctx, req.(*AuthorizeParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_RegisterPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenWithPermission)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).RegisterPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.UserService/RegisterPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).RegisterPermission(ctx, req.(*TokenWithPermission))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UnregisterPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenWithPermission)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UnregisterPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.UserService/UnregisterPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UnregisterPermission(ctx, req.(*TokenWithPermission))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Signup",
			Handler:    _UserService_Signup_Handler,
		},
		{
			MethodName: "Dropout",
			Handler:    _UserService_Dropout_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UserService_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _UserService_Logout_Handler,
		},
		{
			MethodName: "RegisterService",
			Handler:    _UserService_RegisterService_Handler,
		},
		{
			MethodName: "UnregisterService",
			Handler:    _UserService_UnregisterService_Handler,
		},
		{
			MethodName: "Authenticate",
			Handler:    _UserService_Authenticate_Handler,
		},
		{
			MethodName: "Authorize",
			Handler:    _UserService_Authorize_Handler,
		},
		{
			MethodName: "RegisterPermission",
			Handler:    _UserService_RegisterPermission_Handler,
		},
		{
			MethodName: "UnregisterPermission",
			Handler:    _UserService_UnregisterPermission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_service.proto",
}
