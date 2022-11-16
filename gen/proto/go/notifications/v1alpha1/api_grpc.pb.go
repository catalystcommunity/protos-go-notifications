// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package notificationsv1alpha1

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
	UpsertUsers(ctx context.Context, in *NotificationsServiceUpsertUsersRequest, opts ...grpc.CallOption) (*NotificationsServiceUpsertUsersResponse, error)
	GetUsers(ctx context.Context, in *NotificationsServiceGetUsersRequest, opts ...grpc.CallOption) (*NotificationsServiceGetUsersResponse, error)
	ListUsers(ctx context.Context, in *NotificationsServiceListUsersRequest, opts ...grpc.CallOption) (*NotificationsServiceListUsersResponse, error)
	DeleteUsers(ctx context.Context, in *NotificationsServiceDeleteUsersRequest, opts ...grpc.CallOption) (*NotificationsServiceDeleteUsersResponse, error)
	GetNotifications(ctx context.Context, in *NotificationsServiceGetNotificationsRequest, opts ...grpc.CallOption) (*NotificationsServiceGetNotificationsResponse, error)
	SendNotifications(ctx context.Context, in *NotificationsServiceSendNotificationsRequest, opts ...grpc.CallOption) (*NotificationsServiceSendNotificationsResponse, error)
	UpdateSubscriptions(ctx context.Context, in *NotificationsServiceUpdateSubscriptionsRequest, opts ...grpc.CallOption) (*NotificationsServiceUpdateSubscriptionsResponse, error)
	UpsertScheduledNotifications(ctx context.Context, in *NotificationsServiceUpsertScheduledNotificationsRequest, opts ...grpc.CallOption) (*NotificationsServiceUpsertScheduledNotificationsResponse, error)
	GetScheduledNotifications(ctx context.Context, in *NotificationsServiceGetScheduledNotificationsRequest, opts ...grpc.CallOption) (*NotificationsServiceGetScheduledNotificationsResponse, error)
	DeleteScheduledNotifications(ctx context.Context, in *NotificationsServiceDeleteScheduledNotificationsRequest, opts ...grpc.CallOption) (*NotificationsServiceDeleteScheduledNotificationsResponse, error)
}

type notificationsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationsServiceClient(cc grpc.ClientConnInterface) NotificationsServiceClient {
	return &notificationsServiceClient{cc}
}

func (c *notificationsServiceClient) UpsertUsers(ctx context.Context, in *NotificationsServiceUpsertUsersRequest, opts ...grpc.CallOption) (*NotificationsServiceUpsertUsersResponse, error) {
	out := new(NotificationsServiceUpsertUsersResponse)
	err := c.cc.Invoke(ctx, "/notifications.v1alpha1.NotificationsService/UpsertUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsServiceClient) GetUsers(ctx context.Context, in *NotificationsServiceGetUsersRequest, opts ...grpc.CallOption) (*NotificationsServiceGetUsersResponse, error) {
	out := new(NotificationsServiceGetUsersResponse)
	err := c.cc.Invoke(ctx, "/notifications.v1alpha1.NotificationsService/GetUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsServiceClient) ListUsers(ctx context.Context, in *NotificationsServiceListUsersRequest, opts ...grpc.CallOption) (*NotificationsServiceListUsersResponse, error) {
	out := new(NotificationsServiceListUsersResponse)
	err := c.cc.Invoke(ctx, "/notifications.v1alpha1.NotificationsService/ListUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsServiceClient) DeleteUsers(ctx context.Context, in *NotificationsServiceDeleteUsersRequest, opts ...grpc.CallOption) (*NotificationsServiceDeleteUsersResponse, error) {
	out := new(NotificationsServiceDeleteUsersResponse)
	err := c.cc.Invoke(ctx, "/notifications.v1alpha1.NotificationsService/DeleteUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsServiceClient) GetNotifications(ctx context.Context, in *NotificationsServiceGetNotificationsRequest, opts ...grpc.CallOption) (*NotificationsServiceGetNotificationsResponse, error) {
	out := new(NotificationsServiceGetNotificationsResponse)
	err := c.cc.Invoke(ctx, "/notifications.v1alpha1.NotificationsService/GetNotifications", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsServiceClient) SendNotifications(ctx context.Context, in *NotificationsServiceSendNotificationsRequest, opts ...grpc.CallOption) (*NotificationsServiceSendNotificationsResponse, error) {
	out := new(NotificationsServiceSendNotificationsResponse)
	err := c.cc.Invoke(ctx, "/notifications.v1alpha1.NotificationsService/SendNotifications", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsServiceClient) UpdateSubscriptions(ctx context.Context, in *NotificationsServiceUpdateSubscriptionsRequest, opts ...grpc.CallOption) (*NotificationsServiceUpdateSubscriptionsResponse, error) {
	out := new(NotificationsServiceUpdateSubscriptionsResponse)
	err := c.cc.Invoke(ctx, "/notifications.v1alpha1.NotificationsService/UpdateSubscriptions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsServiceClient) UpsertScheduledNotifications(ctx context.Context, in *NotificationsServiceUpsertScheduledNotificationsRequest, opts ...grpc.CallOption) (*NotificationsServiceUpsertScheduledNotificationsResponse, error) {
	out := new(NotificationsServiceUpsertScheduledNotificationsResponse)
	err := c.cc.Invoke(ctx, "/notifications.v1alpha1.NotificationsService/UpsertScheduledNotifications", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsServiceClient) GetScheduledNotifications(ctx context.Context, in *NotificationsServiceGetScheduledNotificationsRequest, opts ...grpc.CallOption) (*NotificationsServiceGetScheduledNotificationsResponse, error) {
	out := new(NotificationsServiceGetScheduledNotificationsResponse)
	err := c.cc.Invoke(ctx, "/notifications.v1alpha1.NotificationsService/GetScheduledNotifications", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsServiceClient) DeleteScheduledNotifications(ctx context.Context, in *NotificationsServiceDeleteScheduledNotificationsRequest, opts ...grpc.CallOption) (*NotificationsServiceDeleteScheduledNotificationsResponse, error) {
	out := new(NotificationsServiceDeleteScheduledNotificationsResponse)
	err := c.cc.Invoke(ctx, "/notifications.v1alpha1.NotificationsService/DeleteScheduledNotifications", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationsServiceServer is the server API for NotificationsService service.
// All implementations should embed UnimplementedNotificationsServiceServer
// for forward compatibility
type NotificationsServiceServer interface {
	UpsertUsers(context.Context, *NotificationsServiceUpsertUsersRequest) (*NotificationsServiceUpsertUsersResponse, error)
	GetUsers(context.Context, *NotificationsServiceGetUsersRequest) (*NotificationsServiceGetUsersResponse, error)
	ListUsers(context.Context, *NotificationsServiceListUsersRequest) (*NotificationsServiceListUsersResponse, error)
	DeleteUsers(context.Context, *NotificationsServiceDeleteUsersRequest) (*NotificationsServiceDeleteUsersResponse, error)
	GetNotifications(context.Context, *NotificationsServiceGetNotificationsRequest) (*NotificationsServiceGetNotificationsResponse, error)
	SendNotifications(context.Context, *NotificationsServiceSendNotificationsRequest) (*NotificationsServiceSendNotificationsResponse, error)
	UpdateSubscriptions(context.Context, *NotificationsServiceUpdateSubscriptionsRequest) (*NotificationsServiceUpdateSubscriptionsResponse, error)
	UpsertScheduledNotifications(context.Context, *NotificationsServiceUpsertScheduledNotificationsRequest) (*NotificationsServiceUpsertScheduledNotificationsResponse, error)
	GetScheduledNotifications(context.Context, *NotificationsServiceGetScheduledNotificationsRequest) (*NotificationsServiceGetScheduledNotificationsResponse, error)
	DeleteScheduledNotifications(context.Context, *NotificationsServiceDeleteScheduledNotificationsRequest) (*NotificationsServiceDeleteScheduledNotificationsResponse, error)
}

// UnimplementedNotificationsServiceServer should be embedded to have forward compatible implementations.
type UnimplementedNotificationsServiceServer struct {
}

func (UnimplementedNotificationsServiceServer) UpsertUsers(context.Context, *NotificationsServiceUpsertUsersRequest) (*NotificationsServiceUpsertUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertUsers not implemented")
}
func (UnimplementedNotificationsServiceServer) GetUsers(context.Context, *NotificationsServiceGetUsersRequest) (*NotificationsServiceGetUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedNotificationsServiceServer) ListUsers(context.Context, *NotificationsServiceListUsersRequest) (*NotificationsServiceListUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (UnimplementedNotificationsServiceServer) DeleteUsers(context.Context, *NotificationsServiceDeleteUsersRequest) (*NotificationsServiceDeleteUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUsers not implemented")
}
func (UnimplementedNotificationsServiceServer) GetNotifications(context.Context, *NotificationsServiceGetNotificationsRequest) (*NotificationsServiceGetNotificationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotifications not implemented")
}
func (UnimplementedNotificationsServiceServer) SendNotifications(context.Context, *NotificationsServiceSendNotificationsRequest) (*NotificationsServiceSendNotificationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendNotifications not implemented")
}
func (UnimplementedNotificationsServiceServer) UpdateSubscriptions(context.Context, *NotificationsServiceUpdateSubscriptionsRequest) (*NotificationsServiceUpdateSubscriptionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSubscriptions not implemented")
}
func (UnimplementedNotificationsServiceServer) UpsertScheduledNotifications(context.Context, *NotificationsServiceUpsertScheduledNotificationsRequest) (*NotificationsServiceUpsertScheduledNotificationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertScheduledNotifications not implemented")
}
func (UnimplementedNotificationsServiceServer) GetScheduledNotifications(context.Context, *NotificationsServiceGetScheduledNotificationsRequest) (*NotificationsServiceGetScheduledNotificationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScheduledNotifications not implemented")
}
func (UnimplementedNotificationsServiceServer) DeleteScheduledNotifications(context.Context, *NotificationsServiceDeleteScheduledNotificationsRequest) (*NotificationsServiceDeleteScheduledNotificationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteScheduledNotifications not implemented")
}

// UnsafeNotificationsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotificationsServiceServer will
// result in compilation errors.
type UnsafeNotificationsServiceServer interface {
	mustEmbedUnimplementedNotificationsServiceServer()
}

func RegisterNotificationsServiceServer(s grpc.ServiceRegistrar, srv NotificationsServiceServer) {
	s.RegisterService(&NotificationsService_ServiceDesc, srv)
}

func _NotificationsService_UpsertUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationsServiceUpsertUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServiceServer).UpsertUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notifications.v1alpha1.NotificationsService/UpsertUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServiceServer).UpsertUsers(ctx, req.(*NotificationsServiceUpsertUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationsService_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationsServiceGetUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServiceServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notifications.v1alpha1.NotificationsService/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServiceServer).GetUsers(ctx, req.(*NotificationsServiceGetUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationsService_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationsServiceListUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServiceServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notifications.v1alpha1.NotificationsService/ListUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServiceServer).ListUsers(ctx, req.(*NotificationsServiceListUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationsService_DeleteUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationsServiceDeleteUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServiceServer).DeleteUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notifications.v1alpha1.NotificationsService/DeleteUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServiceServer).DeleteUsers(ctx, req.(*NotificationsServiceDeleteUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationsService_GetNotifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationsServiceGetNotificationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServiceServer).GetNotifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notifications.v1alpha1.NotificationsService/GetNotifications",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServiceServer).GetNotifications(ctx, req.(*NotificationsServiceGetNotificationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationsService_SendNotifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationsServiceSendNotificationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServiceServer).SendNotifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notifications.v1alpha1.NotificationsService/SendNotifications",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServiceServer).SendNotifications(ctx, req.(*NotificationsServiceSendNotificationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationsService_UpdateSubscriptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationsServiceUpdateSubscriptionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServiceServer).UpdateSubscriptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notifications.v1alpha1.NotificationsService/UpdateSubscriptions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServiceServer).UpdateSubscriptions(ctx, req.(*NotificationsServiceUpdateSubscriptionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationsService_UpsertScheduledNotifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationsServiceUpsertScheduledNotificationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServiceServer).UpsertScheduledNotifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notifications.v1alpha1.NotificationsService/UpsertScheduledNotifications",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServiceServer).UpsertScheduledNotifications(ctx, req.(*NotificationsServiceUpsertScheduledNotificationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationsService_GetScheduledNotifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationsServiceGetScheduledNotificationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServiceServer).GetScheduledNotifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notifications.v1alpha1.NotificationsService/GetScheduledNotifications",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServiceServer).GetScheduledNotifications(ctx, req.(*NotificationsServiceGetScheduledNotificationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationsService_DeleteScheduledNotifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationsServiceDeleteScheduledNotificationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServiceServer).DeleteScheduledNotifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notifications.v1alpha1.NotificationsService/DeleteScheduledNotifications",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServiceServer).DeleteScheduledNotifications(ctx, req.(*NotificationsServiceDeleteScheduledNotificationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NotificationsService_ServiceDesc is the grpc.ServiceDesc for NotificationsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NotificationsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notifications.v1alpha1.NotificationsService",
	HandlerType: (*NotificationsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpsertUsers",
			Handler:    _NotificationsService_UpsertUsers_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _NotificationsService_GetUsers_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _NotificationsService_ListUsers_Handler,
		},
		{
			MethodName: "DeleteUsers",
			Handler:    _NotificationsService_DeleteUsers_Handler,
		},
		{
			MethodName: "GetNotifications",
			Handler:    _NotificationsService_GetNotifications_Handler,
		},
		{
			MethodName: "SendNotifications",
			Handler:    _NotificationsService_SendNotifications_Handler,
		},
		{
			MethodName: "UpdateSubscriptions",
			Handler:    _NotificationsService_UpdateSubscriptions_Handler,
		},
		{
			MethodName: "UpsertScheduledNotifications",
			Handler:    _NotificationsService_UpsertScheduledNotifications_Handler,
		},
		{
			MethodName: "GetScheduledNotifications",
			Handler:    _NotificationsService_GetScheduledNotifications_Handler,
		},
		{
			MethodName: "DeleteScheduledNotifications",
			Handler:    _NotificationsService_DeleteScheduledNotifications_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notifications/v1alpha1/api.proto",
}
