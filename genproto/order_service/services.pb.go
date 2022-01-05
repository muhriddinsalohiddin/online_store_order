// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: order_service/services.proto

package order

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("order_service/services.proto", fileDescriptor_edc415da3680c341) }

var fileDescriptor_edc415da3680c341 = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xc9, 0x2f, 0x4a, 0x49,
	0x2d, 0x8a, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x87, 0xd2, 0xc5, 0x7a, 0x05, 0x45,
	0xf9, 0x25, 0xf9, 0x42, 0xac, 0x60, 0x59, 0x29, 0x49, 0x54, 0x45, 0x60, 0x1e, 0x44, 0x85, 0xd1,
	0x54, 0x26, 0x2e, 0x1e, 0x7f, 0x10, 0x3f, 0x18, 0x22, 0x29, 0xa4, 0xcd, 0xc5, 0xed, 0x5c, 0x94,
	0x9a, 0x58, 0x92, 0x0a, 0x16, 0x15, 0xe2, 0xd1, 0x83, 0xa8, 0x06, 0xf3, 0xa4, 0x50, 0x78, 0x4a,
	0x0c, 0x20, 0xc5, 0xa1, 0x05, 0x29, 0x44, 0x2a, 0x36, 0xe3, 0xe2, 0x71, 0x4f, 0x2d, 0x01, 0xf3,
	0x9c, 0x2a, 0x3d, 0x53, 0x84, 0xc4, 0xa0, 0xf2, 0xc8, 0x82, 0x41, 0xa9, 0x85, 0x18, 0xfa, 0x2c,
	0xb8, 0xb8, 0x5c, 0x52, 0x73, 0x52, 0x4b, 0x52, 0xf1, 0xea, 0x12, 0x80, 0x8a, 0xbb, 0xe6, 0x16,
	0x94, 0x54, 0x06, 0xa5, 0x16, 0x17, 0x28, 0x31, 0x08, 0x59, 0x72, 0x71, 0xf9, 0x64, 0x16, 0x43,
	0xd4, 0x15, 0x0b, 0x09, 0x43, 0x55, 0xc0, 0x85, 0x40, 0xda, 0x44, 0x30, 0x05, 0x41, 0x5a, 0x9d,
	0x04, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x19, 0x8f,
	0xe5, 0x18, 0x92, 0xd8, 0xc0, 0x01, 0x66, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xfd, 0xa5, 0xe7,
	0x07, 0x72, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrderServiceClient interface {
	CreateOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Order, error)
	UpdateOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Order, error)
	GetOrderById(ctx context.Context, in *GetOrderByIdReq, opts ...grpc.CallOption) (*Order, error)
	DeleteById(ctx context.Context, in *GetOrderByIdReq, opts ...grpc.CallOption) (*EmptyResp, error)
	ListOrders(ctx context.Context, in *ListOrderReq, opts ...grpc.CallOption) (*ListOrderResp, error)
}

type orderServiceClient struct {
	cc *grpc.ClientConn
}

func NewOrderServiceClient(cc *grpc.ClientConn) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) CreateOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := c.cc.Invoke(ctx, "/order.OrderService/CreateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) UpdateOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := c.cc.Invoke(ctx, "/order.OrderService/UpdateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetOrderById(ctx context.Context, in *GetOrderByIdReq, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := c.cc.Invoke(ctx, "/order.OrderService/GetOrderById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) DeleteById(ctx context.Context, in *GetOrderByIdReq, opts ...grpc.CallOption) (*EmptyResp, error) {
	out := new(EmptyResp)
	err := c.cc.Invoke(ctx, "/order.OrderService/DeleteById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) ListOrders(ctx context.Context, in *ListOrderReq, opts ...grpc.CallOption) (*ListOrderResp, error) {
	out := new(ListOrderResp)
	err := c.cc.Invoke(ctx, "/order.OrderService/ListOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
type OrderServiceServer interface {
	CreateOrder(context.Context, *Order) (*Order, error)
	UpdateOrder(context.Context, *Order) (*Order, error)
	GetOrderById(context.Context, *GetOrderByIdReq) (*Order, error)
	DeleteById(context.Context, *GetOrderByIdReq) (*EmptyResp, error)
	ListOrders(context.Context, *ListOrderReq) (*ListOrderResp, error)
}

// UnimplementedOrderServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (*UnimplementedOrderServiceServer) CreateOrder(ctx context.Context, req *Order) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (*UnimplementedOrderServiceServer) UpdateOrder(ctx context.Context, req *Order) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrder not implemented")
}
func (*UnimplementedOrderServiceServer) GetOrderById(ctx context.Context, req *GetOrderByIdReq) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderById not implemented")
}
func (*UnimplementedOrderServiceServer) DeleteById(ctx context.Context, req *GetOrderByIdReq) (*EmptyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteById not implemented")
}
func (*UnimplementedOrderServiceServer) ListOrders(ctx context.Context, req *ListOrderReq) (*ListOrderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrders not implemented")
}

func RegisterOrderServiceServer(s *grpc.Server, srv OrderServiceServer) {
	s.RegisterService(&_OrderService_serviceDesc, srv)
}

func _OrderService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.OrderService/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CreateOrder(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_UpdateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).UpdateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.OrderService/UpdateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).UpdateOrder(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetOrderById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrderById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.OrderService/GetOrderById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrderById(ctx, req.(*GetOrderByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_DeleteById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).DeleteById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.OrderService/DeleteById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).DeleteById(ctx, req.(*GetOrderByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_ListOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).ListOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.OrderService/ListOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).ListOrders(ctx, req.(*ListOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _OrderService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "order.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _OrderService_CreateOrder_Handler,
		},
		{
			MethodName: "UpdateOrder",
			Handler:    _OrderService_UpdateOrder_Handler,
		},
		{
			MethodName: "GetOrderById",
			Handler:    _OrderService_GetOrderById_Handler,
		},
		{
			MethodName: "DeleteById",
			Handler:    _OrderService_DeleteById_Handler,
		},
		{
			MethodName: "ListOrders",
			Handler:    _OrderService_ListOrders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order_service/services.proto",
}
