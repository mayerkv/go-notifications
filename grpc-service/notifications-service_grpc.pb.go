// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc_service

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

// NotificationsServiceClient is the client API for NotificationsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotificationsServiceClient interface {
	SendEmail(ctx context.Context, in *SendEmailRequest, opts ...grpc.CallOption) (*Empty, error)
	CreateTemplate(ctx context.Context, in *CreateTemplateRequest, opts ...grpc.CallOption) (*CreateTemplateResponse, error)
	SearchTemplates(ctx context.Context, in *SearchTemplatesRequest, opts ...grpc.CallOption) (*SearchTemplatesResponse, error)
	SearchNotifications(ctx context.Context, in *SearchNotificationsRequest, opts ...grpc.CallOption) (*SearchNotificationsResponse, error)
}

type notificationsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationsServiceClient(cc grpc.ClientConnInterface) NotificationsServiceClient {
	return &notificationsServiceClient{cc}
}

func (c *notificationsServiceClient) SendEmail(ctx context.Context, in *SendEmailRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/NotificationsService/SendEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsServiceClient) CreateTemplate(ctx context.Context, in *CreateTemplateRequest, opts ...grpc.CallOption) (*CreateTemplateResponse, error) {
	out := new(CreateTemplateResponse)
	err := c.cc.Invoke(ctx, "/NotificationsService/CreateTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsServiceClient) SearchTemplates(ctx context.Context, in *SearchTemplatesRequest, opts ...grpc.CallOption) (*SearchTemplatesResponse, error) {
	out := new(SearchTemplatesResponse)
	err := c.cc.Invoke(ctx, "/NotificationsService/SearchTemplates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsServiceClient) SearchNotifications(ctx context.Context, in *SearchNotificationsRequest, opts ...grpc.CallOption) (*SearchNotificationsResponse, error) {
	out := new(SearchNotificationsResponse)
	err := c.cc.Invoke(ctx, "/NotificationsService/SearchNotifications", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationsServiceServer is the server API for NotificationsService service.
// All implementations must embed UnimplementedNotificationsServiceServer
// for forward compatibility
type NotificationsServiceServer interface {
	SendEmail(context.Context, *SendEmailRequest) (*Empty, error)
	CreateTemplate(context.Context, *CreateTemplateRequest) (*CreateTemplateResponse, error)
	SearchTemplates(context.Context, *SearchTemplatesRequest) (*SearchTemplatesResponse, error)
	SearchNotifications(context.Context, *SearchNotificationsRequest) (*SearchNotificationsResponse, error)
	mustEmbedUnimplementedNotificationsServiceServer()
}

// UnimplementedNotificationsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNotificationsServiceServer struct {
}

func (UnimplementedNotificationsServiceServer) SendEmail(context.Context, *SendEmailRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEmail not implemented")
}
func (UnimplementedNotificationsServiceServer) CreateTemplate(context.Context, *CreateTemplateRequest) (*CreateTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTemplate not implemented")
}
func (UnimplementedNotificationsServiceServer) SearchTemplates(context.Context, *SearchTemplatesRequest) (*SearchTemplatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchTemplates not implemented")
}
func (UnimplementedNotificationsServiceServer) SearchNotifications(context.Context, *SearchNotificationsRequest) (*SearchNotificationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchNotifications not implemented")
}
func (UnimplementedNotificationsServiceServer) mustEmbedUnimplementedNotificationsServiceServer() {}

// UnsafeNotificationsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotificationsServiceServer will
// result in compilation errors.
type UnsafeNotificationsServiceServer interface {
	mustEmbedUnimplementedNotificationsServiceServer()
}

func RegisterNotificationsServiceServer(s grpc.ServiceRegistrar, srv NotificationsServiceServer) {
	s.RegisterService(&NotificationsService_ServiceDesc, srv)
}

func _NotificationsService_SendEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServiceServer).SendEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NotificationsService/SendEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServiceServer).SendEmail(ctx, req.(*SendEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationsService_CreateTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServiceServer).CreateTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NotificationsService/CreateTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServiceServer).CreateTemplate(ctx, req.(*CreateTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationsService_SearchTemplates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchTemplatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServiceServer).SearchTemplates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NotificationsService/SearchTemplates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServiceServer).SearchTemplates(ctx, req.(*SearchTemplatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationsService_SearchNotifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchNotificationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServiceServer).SearchNotifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NotificationsService/SearchNotifications",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServiceServer).SearchNotifications(ctx, req.(*SearchNotificationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NotificationsService_ServiceDesc is the grpc.ServiceDesc for NotificationsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NotificationsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "NotificationsService",
	HandlerType: (*NotificationsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendEmail",
			Handler:    _NotificationsService_SendEmail_Handler,
		},
		{
			MethodName: "CreateTemplate",
			Handler:    _NotificationsService_CreateTemplate_Handler,
		},
		{
			MethodName: "SearchTemplates",
			Handler:    _NotificationsService_SearchTemplates_Handler,
		},
		{
			MethodName: "SearchNotifications",
			Handler:    _NotificationsService_SearchNotifications_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc-service/notifications-service.proto",
}