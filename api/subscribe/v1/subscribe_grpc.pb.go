// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// SubscribeClient is the client API for Subscribe service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubscribeClient interface {
	SubscribeEntitiesByIDs(ctx context.Context, in *SubscribeEntitiesByIDsRequest, opts ...grpc.CallOption) (*SubscribeEntitiesByIDsResponse, error)
	SubscribeEntitiesByGroups(ctx context.Context, in *SubscribeEntitiesByGroupsRequest, opts ...grpc.CallOption) (*SubscribeEntitiesByGroupsResponse, error)
	SubscribeEntitiesByModels(ctx context.Context, in *SubscribeEntitiesByModelsRequest, opts ...grpc.CallOption) (*SubscribeEntitiesByModelsResponse, error)
	UnsubscribeEntitiesByIDs(ctx context.Context, in *UnsubscribeEntitiesByIDsRequest, opts ...grpc.CallOption) (*UnsubscribeEntitiesByIDsResponse, error)
	DeleteEntitiesByID(ctx context.Context, in *DeleteEntitiesByIDRequest, opts ...grpc.CallOption) (*DeleteEntitiesByIDResponse, error)
	ListSubscribeEntities(ctx context.Context, in *ListSubscribeEntitiesRequest, opts ...grpc.CallOption) (*ListSubscribeEntitiesResponse, error)
	CreateSubscribe(ctx context.Context, in *CreateSubscribeRequest, opts ...grpc.CallOption) (*CreateSubscribeResponse, error)
	UpdateSubscribe(ctx context.Context, in *UpdateSubscribeRequest, opts ...grpc.CallOption) (*UpdateSubscribeResponse, error)
	DeleteSubscribe(ctx context.Context, in *DeleteSubscribeRequest, opts ...grpc.CallOption) (*DeleteSubscribeResponse, error)
	GetSubscribe(ctx context.Context, in *GetSubscribeRequest, opts ...grpc.CallOption) (*GetSubscribeResponse, error)
	ListSubscribe(ctx context.Context, in *ListSubscribeRequest, opts ...grpc.CallOption) (*ListSubscribeResponse, error)
	ChangeSubscribed(ctx context.Context, in *ChangeSubscribedRequest, opts ...grpc.CallOption) (*ChangeSubscribedResponse, error)
	ValidateSubscribed(ctx context.Context, in *ValidateSubscribedRequest, opts ...grpc.CallOption) (*ValidateSubscribedResponse, error)
	SubscribeByDevice(ctx context.Context, in *SubscribeByDeviceRequest, opts ...grpc.CallOption) (*SubscribeByDeviceResponse, error)
}

type subscribeClient struct {
	cc grpc.ClientConnInterface
}

func NewSubscribeClient(cc grpc.ClientConnInterface) SubscribeClient {
	return &subscribeClient{cc}
}

func (c *subscribeClient) SubscribeEntitiesByIDs(ctx context.Context, in *SubscribeEntitiesByIDsRequest, opts ...grpc.CallOption) (*SubscribeEntitiesByIDsResponse, error) {
	out := new(SubscribeEntitiesByIDsResponse)
	err := c.cc.Invoke(ctx, "/api.subscribe.v1.Subscribe/SubscribeEntitiesByIDs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscribeClient) SubscribeEntitiesByGroups(ctx context.Context, in *SubscribeEntitiesByGroupsRequest, opts ...grpc.CallOption) (*SubscribeEntitiesByGroupsResponse, error) {
	out := new(SubscribeEntitiesByGroupsResponse)
	err := c.cc.Invoke(ctx, "/api.subscribe.v1.Subscribe/SubscribeEntitiesByGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscribeClient) SubscribeEntitiesByModels(ctx context.Context, in *SubscribeEntitiesByModelsRequest, opts ...grpc.CallOption) (*SubscribeEntitiesByModelsResponse, error) {
	out := new(SubscribeEntitiesByModelsResponse)
	err := c.cc.Invoke(ctx, "/api.subscribe.v1.Subscribe/SubscribeEntitiesByModels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscribeClient) UnsubscribeEntitiesByIDs(ctx context.Context, in *UnsubscribeEntitiesByIDsRequest, opts ...grpc.CallOption) (*UnsubscribeEntitiesByIDsResponse, error) {
	out := new(UnsubscribeEntitiesByIDsResponse)
	err := c.cc.Invoke(ctx, "/api.subscribe.v1.Subscribe/UnsubscribeEntitiesByIDs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscribeClient) DeleteEntitiesByID(ctx context.Context, in *DeleteEntitiesByIDRequest, opts ...grpc.CallOption) (*DeleteEntitiesByIDResponse, error) {
	out := new(DeleteEntitiesByIDResponse)
	err := c.cc.Invoke(ctx, "/api.subscribe.v1.Subscribe/DeleteEntitiesByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscribeClient) ListSubscribeEntities(ctx context.Context, in *ListSubscribeEntitiesRequest, opts ...grpc.CallOption) (*ListSubscribeEntitiesResponse, error) {
	out := new(ListSubscribeEntitiesResponse)
	err := c.cc.Invoke(ctx, "/api.subscribe.v1.Subscribe/ListSubscribeEntities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscribeClient) CreateSubscribe(ctx context.Context, in *CreateSubscribeRequest, opts ...grpc.CallOption) (*CreateSubscribeResponse, error) {
	out := new(CreateSubscribeResponse)
	err := c.cc.Invoke(ctx, "/api.subscribe.v1.Subscribe/CreateSubscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscribeClient) UpdateSubscribe(ctx context.Context, in *UpdateSubscribeRequest, opts ...grpc.CallOption) (*UpdateSubscribeResponse, error) {
	out := new(UpdateSubscribeResponse)
	err := c.cc.Invoke(ctx, "/api.subscribe.v1.Subscribe/UpdateSubscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscribeClient) DeleteSubscribe(ctx context.Context, in *DeleteSubscribeRequest, opts ...grpc.CallOption) (*DeleteSubscribeResponse, error) {
	out := new(DeleteSubscribeResponse)
	err := c.cc.Invoke(ctx, "/api.subscribe.v1.Subscribe/DeleteSubscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscribeClient) GetSubscribe(ctx context.Context, in *GetSubscribeRequest, opts ...grpc.CallOption) (*GetSubscribeResponse, error) {
	out := new(GetSubscribeResponse)
	err := c.cc.Invoke(ctx, "/api.subscribe.v1.Subscribe/GetSubscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscribeClient) ListSubscribe(ctx context.Context, in *ListSubscribeRequest, opts ...grpc.CallOption) (*ListSubscribeResponse, error) {
	out := new(ListSubscribeResponse)
	err := c.cc.Invoke(ctx, "/api.subscribe.v1.Subscribe/ListSubscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscribeClient) ChangeSubscribed(ctx context.Context, in *ChangeSubscribedRequest, opts ...grpc.CallOption) (*ChangeSubscribedResponse, error) {
	out := new(ChangeSubscribedResponse)
	err := c.cc.Invoke(ctx, "/api.subscribe.v1.Subscribe/ChangeSubscribed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscribeClient) ValidateSubscribed(ctx context.Context, in *ValidateSubscribedRequest, opts ...grpc.CallOption) (*ValidateSubscribedResponse, error) {
	out := new(ValidateSubscribedResponse)
	err := c.cc.Invoke(ctx, "/api.subscribe.v1.Subscribe/ValidateSubscribed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscribeClient) SubscribeByDevice(ctx context.Context, in *SubscribeByDeviceRequest, opts ...grpc.CallOption) (*SubscribeByDeviceResponse, error) {
	out := new(SubscribeByDeviceResponse)
	err := c.cc.Invoke(ctx, "/api.subscribe.v1.Subscribe/SubscribeByDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubscribeServer is the server API for Subscribe service.
// All implementations must embed UnimplementedSubscribeServer
// for forward compatibility
type SubscribeServer interface {
	SubscribeEntitiesByIDs(context.Context, *SubscribeEntitiesByIDsRequest) (*SubscribeEntitiesByIDsResponse, error)
	SubscribeEntitiesByGroups(context.Context, *SubscribeEntitiesByGroupsRequest) (*SubscribeEntitiesByGroupsResponse, error)
	SubscribeEntitiesByModels(context.Context, *SubscribeEntitiesByModelsRequest) (*SubscribeEntitiesByModelsResponse, error)
	UnsubscribeEntitiesByIDs(context.Context, *UnsubscribeEntitiesByIDsRequest) (*UnsubscribeEntitiesByIDsResponse, error)
	DeleteEntitiesByID(context.Context, *DeleteEntitiesByIDRequest) (*DeleteEntitiesByIDResponse, error)
	ListSubscribeEntities(context.Context, *ListSubscribeEntitiesRequest) (*ListSubscribeEntitiesResponse, error)
	CreateSubscribe(context.Context, *CreateSubscribeRequest) (*CreateSubscribeResponse, error)
	UpdateSubscribe(context.Context, *UpdateSubscribeRequest) (*UpdateSubscribeResponse, error)
	DeleteSubscribe(context.Context, *DeleteSubscribeRequest) (*DeleteSubscribeResponse, error)
	GetSubscribe(context.Context, *GetSubscribeRequest) (*GetSubscribeResponse, error)
	ListSubscribe(context.Context, *ListSubscribeRequest) (*ListSubscribeResponse, error)
	ChangeSubscribed(context.Context, *ChangeSubscribedRequest) (*ChangeSubscribedResponse, error)
	ValidateSubscribed(context.Context, *ValidateSubscribedRequest) (*ValidateSubscribedResponse, error)
	SubscribeByDevice(context.Context, *SubscribeByDeviceRequest) (*SubscribeByDeviceResponse, error)
	mustEmbedUnimplementedSubscribeServer()
}

// UnimplementedSubscribeServer must be embedded to have forward compatible implementations.
type UnimplementedSubscribeServer struct {
}

func (UnimplementedSubscribeServer) SubscribeEntitiesByIDs(context.Context, *SubscribeEntitiesByIDsRequest) (*SubscribeEntitiesByIDsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubscribeEntitiesByIDs not implemented")
}
func (UnimplementedSubscribeServer) SubscribeEntitiesByGroups(context.Context, *SubscribeEntitiesByGroupsRequest) (*SubscribeEntitiesByGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubscribeEntitiesByGroups not implemented")
}
func (UnimplementedSubscribeServer) SubscribeEntitiesByModels(context.Context, *SubscribeEntitiesByModelsRequest) (*SubscribeEntitiesByModelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubscribeEntitiesByModels not implemented")
}
func (UnimplementedSubscribeServer) UnsubscribeEntitiesByIDs(context.Context, *UnsubscribeEntitiesByIDsRequest) (*UnsubscribeEntitiesByIDsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnsubscribeEntitiesByIDs not implemented")
}
func (UnimplementedSubscribeServer) DeleteEntitiesByID(context.Context, *DeleteEntitiesByIDRequest) (*DeleteEntitiesByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEntitiesByID not implemented")
}
func (UnimplementedSubscribeServer) ListSubscribeEntities(context.Context, *ListSubscribeEntitiesRequest) (*ListSubscribeEntitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSubscribeEntities not implemented")
}
func (UnimplementedSubscribeServer) CreateSubscribe(context.Context, *CreateSubscribeRequest) (*CreateSubscribeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSubscribe not implemented")
}
func (UnimplementedSubscribeServer) UpdateSubscribe(context.Context, *UpdateSubscribeRequest) (*UpdateSubscribeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSubscribe not implemented")
}
func (UnimplementedSubscribeServer) DeleteSubscribe(context.Context, *DeleteSubscribeRequest) (*DeleteSubscribeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSubscribe not implemented")
}
func (UnimplementedSubscribeServer) GetSubscribe(context.Context, *GetSubscribeRequest) (*GetSubscribeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubscribe not implemented")
}
func (UnimplementedSubscribeServer) ListSubscribe(context.Context, *ListSubscribeRequest) (*ListSubscribeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSubscribe not implemented")
}
func (UnimplementedSubscribeServer) ChangeSubscribed(context.Context, *ChangeSubscribedRequest) (*ChangeSubscribedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeSubscribed not implemented")
}
func (UnimplementedSubscribeServer) ValidateSubscribed(context.Context, *ValidateSubscribedRequest) (*ValidateSubscribedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateSubscribed not implemented")
}
func (UnimplementedSubscribeServer) SubscribeByDevice(context.Context, *SubscribeByDeviceRequest) (*SubscribeByDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubscribeByDevice not implemented")
}
func (UnimplementedSubscribeServer) mustEmbedUnimplementedSubscribeServer() {}

// UnsafeSubscribeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubscribeServer will
// result in compilation errors.
type UnsafeSubscribeServer interface {
	mustEmbedUnimplementedSubscribeServer()
}

func RegisterSubscribeServer(s grpc.ServiceRegistrar, srv SubscribeServer) {
	s.RegisterService(&Subscribe_ServiceDesc, srv)
}

func _Subscribe_SubscribeEntitiesByIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeEntitiesByIDsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscribeServer).SubscribeEntitiesByIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.subscribe.v1.Subscribe/SubscribeEntitiesByIDs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscribeServer).SubscribeEntitiesByIDs(ctx, req.(*SubscribeEntitiesByIDsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subscribe_SubscribeEntitiesByGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeEntitiesByGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscribeServer).SubscribeEntitiesByGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.subscribe.v1.Subscribe/SubscribeEntitiesByGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscribeServer).SubscribeEntitiesByGroups(ctx, req.(*SubscribeEntitiesByGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subscribe_SubscribeEntitiesByModels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeEntitiesByModelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscribeServer).SubscribeEntitiesByModels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.subscribe.v1.Subscribe/SubscribeEntitiesByModels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscribeServer).SubscribeEntitiesByModels(ctx, req.(*SubscribeEntitiesByModelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subscribe_UnsubscribeEntitiesByIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnsubscribeEntitiesByIDsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscribeServer).UnsubscribeEntitiesByIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.subscribe.v1.Subscribe/UnsubscribeEntitiesByIDs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscribeServer).UnsubscribeEntitiesByIDs(ctx, req.(*UnsubscribeEntitiesByIDsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subscribe_DeleteEntitiesByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEntitiesByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscribeServer).DeleteEntitiesByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.subscribe.v1.Subscribe/DeleteEntitiesByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscribeServer).DeleteEntitiesByID(ctx, req.(*DeleteEntitiesByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subscribe_ListSubscribeEntities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSubscribeEntitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscribeServer).ListSubscribeEntities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.subscribe.v1.Subscribe/ListSubscribeEntities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscribeServer).ListSubscribeEntities(ctx, req.(*ListSubscribeEntitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subscribe_CreateSubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscribeServer).CreateSubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.subscribe.v1.Subscribe/CreateSubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscribeServer).CreateSubscribe(ctx, req.(*CreateSubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subscribe_UpdateSubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscribeServer).UpdateSubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.subscribe.v1.Subscribe/UpdateSubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscribeServer).UpdateSubscribe(ctx, req.(*UpdateSubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subscribe_DeleteSubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscribeServer).DeleteSubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.subscribe.v1.Subscribe/DeleteSubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscribeServer).DeleteSubscribe(ctx, req.(*DeleteSubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subscribe_GetSubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscribeServer).GetSubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.subscribe.v1.Subscribe/GetSubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscribeServer).GetSubscribe(ctx, req.(*GetSubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subscribe_ListSubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscribeServer).ListSubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.subscribe.v1.Subscribe/ListSubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscribeServer).ListSubscribe(ctx, req.(*ListSubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subscribe_ChangeSubscribed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeSubscribedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscribeServer).ChangeSubscribed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.subscribe.v1.Subscribe/ChangeSubscribed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscribeServer).ChangeSubscribed(ctx, req.(*ChangeSubscribedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subscribe_ValidateSubscribed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateSubscribedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscribeServer).ValidateSubscribed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.subscribe.v1.Subscribe/ValidateSubscribed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscribeServer).ValidateSubscribed(ctx, req.(*ValidateSubscribedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subscribe_SubscribeByDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeByDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscribeServer).SubscribeByDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.subscribe.v1.Subscribe/SubscribeByDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscribeServer).SubscribeByDevice(ctx, req.(*SubscribeByDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Subscribe_ServiceDesc is the grpc.ServiceDesc for Subscribe service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Subscribe_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.subscribe.v1.Subscribe",
	HandlerType: (*SubscribeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubscribeEntitiesByIDs",
			Handler:    _Subscribe_SubscribeEntitiesByIDs_Handler,
		},
		{
			MethodName: "SubscribeEntitiesByGroups",
			Handler:    _Subscribe_SubscribeEntitiesByGroups_Handler,
		},
		{
			MethodName: "SubscribeEntitiesByModels",
			Handler:    _Subscribe_SubscribeEntitiesByModels_Handler,
		},
		{
			MethodName: "UnsubscribeEntitiesByIDs",
			Handler:    _Subscribe_UnsubscribeEntitiesByIDs_Handler,
		},
		{
			MethodName: "DeleteEntitiesByID",
			Handler:    _Subscribe_DeleteEntitiesByID_Handler,
		},
		{
			MethodName: "ListSubscribeEntities",
			Handler:    _Subscribe_ListSubscribeEntities_Handler,
		},
		{
			MethodName: "CreateSubscribe",
			Handler:    _Subscribe_CreateSubscribe_Handler,
		},
		{
			MethodName: "UpdateSubscribe",
			Handler:    _Subscribe_UpdateSubscribe_Handler,
		},
		{
			MethodName: "DeleteSubscribe",
			Handler:    _Subscribe_DeleteSubscribe_Handler,
		},
		{
			MethodName: "GetSubscribe",
			Handler:    _Subscribe_GetSubscribe_Handler,
		},
		{
			MethodName: "ListSubscribe",
			Handler:    _Subscribe_ListSubscribe_Handler,
		},
		{
			MethodName: "ChangeSubscribed",
			Handler:    _Subscribe_ChangeSubscribed_Handler,
		},
		{
			MethodName: "ValidateSubscribed",
			Handler:    _Subscribe_ValidateSubscribed_Handler,
		},
		{
			MethodName: "SubscribeByDevice",
			Handler:    _Subscribe_SubscribeByDevice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/subscribe/v1/subscribe.proto",
}
