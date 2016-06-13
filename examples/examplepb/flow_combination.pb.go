// Code generated by protoc-gen-go.
// source: examples/examplepb/flow_combination.proto
// DO NOT EDIT!

package examplepb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gengo/grpc-gateway/third_party/googleapis/google/api"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type EmptyProto struct {
}

func (m *EmptyProto) Reset()                    { *m = EmptyProto{} }
func (m *EmptyProto) String() string            { return proto.CompactTextString(m) }
func (*EmptyProto) ProtoMessage()               {}
func (*EmptyProto) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type NonEmptyProto struct {
	A string `protobuf:"bytes,1,opt,name=a" json:"a,omitempty"`
	B string `protobuf:"bytes,2,opt,name=b" json:"b,omitempty"`
	C string `protobuf:"bytes,3,opt,name=c" json:"c,omitempty"`
}

func (m *NonEmptyProto) Reset()                    { *m = NonEmptyProto{} }
func (m *NonEmptyProto) String() string            { return proto.CompactTextString(m) }
func (*NonEmptyProto) ProtoMessage()               {}
func (*NonEmptyProto) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

type UnaryProto struct {
	Str string `protobuf:"bytes,1,opt,name=str" json:"str,omitempty"`
}

func (m *UnaryProto) Reset()                    { *m = UnaryProto{} }
func (m *UnaryProto) String() string            { return proto.CompactTextString(m) }
func (*UnaryProto) ProtoMessage()               {}
func (*UnaryProto) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

type NestedProto struct {
	A *UnaryProto `protobuf:"bytes,1,opt,name=a" json:"a,omitempty"`
	B string      `protobuf:"bytes,2,opt,name=b" json:"b,omitempty"`
	C string      `protobuf:"bytes,3,opt,name=c" json:"c,omitempty"`
}

func (m *NestedProto) Reset()                    { *m = NestedProto{} }
func (m *NestedProto) String() string            { return proto.CompactTextString(m) }
func (*NestedProto) ProtoMessage()               {}
func (*NestedProto) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *NestedProto) GetA() *UnaryProto {
	if m != nil {
		return m.A
	}
	return nil
}

type SingleNestedProto struct {
	A *UnaryProto `protobuf:"bytes,1,opt,name=a" json:"a,omitempty"`
}

func (m *SingleNestedProto) Reset()                    { *m = SingleNestedProto{} }
func (m *SingleNestedProto) String() string            { return proto.CompactTextString(m) }
func (*SingleNestedProto) ProtoMessage()               {}
func (*SingleNestedProto) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *SingleNestedProto) GetA() *UnaryProto {
	if m != nil {
		return m.A
	}
	return nil
}

func init() {
	proto.RegisterType((*EmptyProto)(nil), "gengo.grpc.gateway.examples.examplepb.EmptyProto")
	proto.RegisterType((*NonEmptyProto)(nil), "gengo.grpc.gateway.examples.examplepb.NonEmptyProto")
	proto.RegisterType((*UnaryProto)(nil), "gengo.grpc.gateway.examples.examplepb.UnaryProto")
	proto.RegisterType((*NestedProto)(nil), "gengo.grpc.gateway.examples.examplepb.NestedProto")
	proto.RegisterType((*SingleNestedProto)(nil), "gengo.grpc.gateway.examples.examplepb.SingleNestedProto")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for FlowCombination service

type FlowCombinationClient interface {
	RpcEmptyRpc(ctx context.Context, in *EmptyProto, opts ...grpc.CallOption) (*EmptyProto, error)
	RpcEmptyStream(ctx context.Context, in *EmptyProto, opts ...grpc.CallOption) (FlowCombination_RpcEmptyStreamClient, error)
	StreamEmptyRpc(ctx context.Context, opts ...grpc.CallOption) (FlowCombination_StreamEmptyRpcClient, error)
	StreamEmptyStream(ctx context.Context, opts ...grpc.CallOption) (FlowCombination_StreamEmptyStreamClient, error)
	RpcBodyRpc(ctx context.Context, in *NonEmptyProto, opts ...grpc.CallOption) (*EmptyProto, error)
	RpcPathSingleNestedRpc(ctx context.Context, in *SingleNestedProto, opts ...grpc.CallOption) (*EmptyProto, error)
	RpcPathNestedRpc(ctx context.Context, in *NestedProto, opts ...grpc.CallOption) (*EmptyProto, error)
	RpcBodyStream(ctx context.Context, in *NonEmptyProto, opts ...grpc.CallOption) (FlowCombination_RpcBodyStreamClient, error)
	RpcPathSingleNestedStream(ctx context.Context, in *SingleNestedProto, opts ...grpc.CallOption) (FlowCombination_RpcPathSingleNestedStreamClient, error)
	RpcPathNestedStream(ctx context.Context, in *NestedProto, opts ...grpc.CallOption) (FlowCombination_RpcPathNestedStreamClient, error)
}

type flowCombinationClient struct {
	cc *grpc.ClientConn
}

func NewFlowCombinationClient(cc *grpc.ClientConn) FlowCombinationClient {
	return &flowCombinationClient{cc}
}

func (c *flowCombinationClient) RpcEmptyRpc(ctx context.Context, in *EmptyProto, opts ...grpc.CallOption) (*EmptyProto, error) {
	out := new(EmptyProto)
	err := grpc.Invoke(ctx, "/gengo.grpc.gateway.examples.examplepb.FlowCombination/RpcEmptyRpc", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowCombinationClient) RpcEmptyStream(ctx context.Context, in *EmptyProto, opts ...grpc.CallOption) (FlowCombination_RpcEmptyStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_FlowCombination_serviceDesc.Streams[0], c.cc, "/gengo.grpc.gateway.examples.examplepb.FlowCombination/RpcEmptyStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &flowCombinationRpcEmptyStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FlowCombination_RpcEmptyStreamClient interface {
	Recv() (*EmptyProto, error)
	grpc.ClientStream
}

type flowCombinationRpcEmptyStreamClient struct {
	grpc.ClientStream
}

func (x *flowCombinationRpcEmptyStreamClient) Recv() (*EmptyProto, error) {
	m := new(EmptyProto)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *flowCombinationClient) StreamEmptyRpc(ctx context.Context, opts ...grpc.CallOption) (FlowCombination_StreamEmptyRpcClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_FlowCombination_serviceDesc.Streams[1], c.cc, "/gengo.grpc.gateway.examples.examplepb.FlowCombination/StreamEmptyRpc", opts...)
	if err != nil {
		return nil, err
	}
	x := &flowCombinationStreamEmptyRpcClient{stream}
	return x, nil
}

type FlowCombination_StreamEmptyRpcClient interface {
	Send(*EmptyProto) error
	CloseAndRecv() (*EmptyProto, error)
	grpc.ClientStream
}

type flowCombinationStreamEmptyRpcClient struct {
	grpc.ClientStream
}

func (x *flowCombinationStreamEmptyRpcClient) Send(m *EmptyProto) error {
	return x.ClientStream.SendMsg(m)
}

func (x *flowCombinationStreamEmptyRpcClient) CloseAndRecv() (*EmptyProto, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(EmptyProto)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *flowCombinationClient) StreamEmptyStream(ctx context.Context, opts ...grpc.CallOption) (FlowCombination_StreamEmptyStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_FlowCombination_serviceDesc.Streams[2], c.cc, "/gengo.grpc.gateway.examples.examplepb.FlowCombination/StreamEmptyStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &flowCombinationStreamEmptyStreamClient{stream}
	return x, nil
}

type FlowCombination_StreamEmptyStreamClient interface {
	Send(*EmptyProto) error
	Recv() (*EmptyProto, error)
	grpc.ClientStream
}

type flowCombinationStreamEmptyStreamClient struct {
	grpc.ClientStream
}

func (x *flowCombinationStreamEmptyStreamClient) Send(m *EmptyProto) error {
	return x.ClientStream.SendMsg(m)
}

func (x *flowCombinationStreamEmptyStreamClient) Recv() (*EmptyProto, error) {
	m := new(EmptyProto)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *flowCombinationClient) RpcBodyRpc(ctx context.Context, in *NonEmptyProto, opts ...grpc.CallOption) (*EmptyProto, error) {
	out := new(EmptyProto)
	err := grpc.Invoke(ctx, "/gengo.grpc.gateway.examples.examplepb.FlowCombination/RpcBodyRpc", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowCombinationClient) RpcPathSingleNestedRpc(ctx context.Context, in *SingleNestedProto, opts ...grpc.CallOption) (*EmptyProto, error) {
	out := new(EmptyProto)
	err := grpc.Invoke(ctx, "/gengo.grpc.gateway.examples.examplepb.FlowCombination/RpcPathSingleNestedRpc", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowCombinationClient) RpcPathNestedRpc(ctx context.Context, in *NestedProto, opts ...grpc.CallOption) (*EmptyProto, error) {
	out := new(EmptyProto)
	err := grpc.Invoke(ctx, "/gengo.grpc.gateway.examples.examplepb.FlowCombination/RpcPathNestedRpc", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowCombinationClient) RpcBodyStream(ctx context.Context, in *NonEmptyProto, opts ...grpc.CallOption) (FlowCombination_RpcBodyStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_FlowCombination_serviceDesc.Streams[3], c.cc, "/gengo.grpc.gateway.examples.examplepb.FlowCombination/RpcBodyStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &flowCombinationRpcBodyStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FlowCombination_RpcBodyStreamClient interface {
	Recv() (*EmptyProto, error)
	grpc.ClientStream
}

type flowCombinationRpcBodyStreamClient struct {
	grpc.ClientStream
}

func (x *flowCombinationRpcBodyStreamClient) Recv() (*EmptyProto, error) {
	m := new(EmptyProto)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *flowCombinationClient) RpcPathSingleNestedStream(ctx context.Context, in *SingleNestedProto, opts ...grpc.CallOption) (FlowCombination_RpcPathSingleNestedStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_FlowCombination_serviceDesc.Streams[4], c.cc, "/gengo.grpc.gateway.examples.examplepb.FlowCombination/RpcPathSingleNestedStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &flowCombinationRpcPathSingleNestedStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FlowCombination_RpcPathSingleNestedStreamClient interface {
	Recv() (*EmptyProto, error)
	grpc.ClientStream
}

type flowCombinationRpcPathSingleNestedStreamClient struct {
	grpc.ClientStream
}

func (x *flowCombinationRpcPathSingleNestedStreamClient) Recv() (*EmptyProto, error) {
	m := new(EmptyProto)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *flowCombinationClient) RpcPathNestedStream(ctx context.Context, in *NestedProto, opts ...grpc.CallOption) (FlowCombination_RpcPathNestedStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_FlowCombination_serviceDesc.Streams[5], c.cc, "/gengo.grpc.gateway.examples.examplepb.FlowCombination/RpcPathNestedStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &flowCombinationRpcPathNestedStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FlowCombination_RpcPathNestedStreamClient interface {
	Recv() (*EmptyProto, error)
	grpc.ClientStream
}

type flowCombinationRpcPathNestedStreamClient struct {
	grpc.ClientStream
}

func (x *flowCombinationRpcPathNestedStreamClient) Recv() (*EmptyProto, error) {
	m := new(EmptyProto)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for FlowCombination service

type FlowCombinationServer interface {
	RpcEmptyRpc(context.Context, *EmptyProto) (*EmptyProto, error)
	RpcEmptyStream(*EmptyProto, FlowCombination_RpcEmptyStreamServer) error
	StreamEmptyRpc(FlowCombination_StreamEmptyRpcServer) error
	StreamEmptyStream(FlowCombination_StreamEmptyStreamServer) error
	RpcBodyRpc(context.Context, *NonEmptyProto) (*EmptyProto, error)
	RpcPathSingleNestedRpc(context.Context, *SingleNestedProto) (*EmptyProto, error)
	RpcPathNestedRpc(context.Context, *NestedProto) (*EmptyProto, error)
	RpcBodyStream(*NonEmptyProto, FlowCombination_RpcBodyStreamServer) error
	RpcPathSingleNestedStream(*SingleNestedProto, FlowCombination_RpcPathSingleNestedStreamServer) error
	RpcPathNestedStream(*NestedProto, FlowCombination_RpcPathNestedStreamServer) error
}

func RegisterFlowCombinationServer(s *grpc.Server, srv FlowCombinationServer) {
	s.RegisterService(&_FlowCombination_serviceDesc, srv)
}

func _FlowCombination_RpcEmptyRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlowCombinationServer).RpcEmptyRpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gengo.grpc.gateway.examples.examplepb.FlowCombination/RpcEmptyRpc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlowCombinationServer).RpcEmptyRpc(ctx, req.(*EmptyProto))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlowCombination_RpcEmptyStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EmptyProto)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FlowCombinationServer).RpcEmptyStream(m, &flowCombinationRpcEmptyStreamServer{stream})
}

type FlowCombination_RpcEmptyStreamServer interface {
	Send(*EmptyProto) error
	grpc.ServerStream
}

type flowCombinationRpcEmptyStreamServer struct {
	grpc.ServerStream
}

func (x *flowCombinationRpcEmptyStreamServer) Send(m *EmptyProto) error {
	return x.ServerStream.SendMsg(m)
}

func _FlowCombination_StreamEmptyRpc_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FlowCombinationServer).StreamEmptyRpc(&flowCombinationStreamEmptyRpcServer{stream})
}

type FlowCombination_StreamEmptyRpcServer interface {
	SendAndClose(*EmptyProto) error
	Recv() (*EmptyProto, error)
	grpc.ServerStream
}

type flowCombinationStreamEmptyRpcServer struct {
	grpc.ServerStream
}

func (x *flowCombinationStreamEmptyRpcServer) SendAndClose(m *EmptyProto) error {
	return x.ServerStream.SendMsg(m)
}

func (x *flowCombinationStreamEmptyRpcServer) Recv() (*EmptyProto, error) {
	m := new(EmptyProto)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _FlowCombination_StreamEmptyStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FlowCombinationServer).StreamEmptyStream(&flowCombinationStreamEmptyStreamServer{stream})
}

type FlowCombination_StreamEmptyStreamServer interface {
	Send(*EmptyProto) error
	Recv() (*EmptyProto, error)
	grpc.ServerStream
}

type flowCombinationStreamEmptyStreamServer struct {
	grpc.ServerStream
}

func (x *flowCombinationStreamEmptyStreamServer) Send(m *EmptyProto) error {
	return x.ServerStream.SendMsg(m)
}

func (x *flowCombinationStreamEmptyStreamServer) Recv() (*EmptyProto, error) {
	m := new(EmptyProto)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _FlowCombination_RpcBodyRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NonEmptyProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlowCombinationServer).RpcBodyRpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gengo.grpc.gateway.examples.examplepb.FlowCombination/RpcBodyRpc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlowCombinationServer).RpcBodyRpc(ctx, req.(*NonEmptyProto))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlowCombination_RpcPathSingleNestedRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleNestedProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlowCombinationServer).RpcPathSingleNestedRpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gengo.grpc.gateway.examples.examplepb.FlowCombination/RpcPathSingleNestedRpc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlowCombinationServer).RpcPathSingleNestedRpc(ctx, req.(*SingleNestedProto))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlowCombination_RpcPathNestedRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NestedProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlowCombinationServer).RpcPathNestedRpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gengo.grpc.gateway.examples.examplepb.FlowCombination/RpcPathNestedRpc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlowCombinationServer).RpcPathNestedRpc(ctx, req.(*NestedProto))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlowCombination_RpcBodyStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NonEmptyProto)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FlowCombinationServer).RpcBodyStream(m, &flowCombinationRpcBodyStreamServer{stream})
}

type FlowCombination_RpcBodyStreamServer interface {
	Send(*EmptyProto) error
	grpc.ServerStream
}

type flowCombinationRpcBodyStreamServer struct {
	grpc.ServerStream
}

func (x *flowCombinationRpcBodyStreamServer) Send(m *EmptyProto) error {
	return x.ServerStream.SendMsg(m)
}

func _FlowCombination_RpcPathSingleNestedStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SingleNestedProto)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FlowCombinationServer).RpcPathSingleNestedStream(m, &flowCombinationRpcPathSingleNestedStreamServer{stream})
}

type FlowCombination_RpcPathSingleNestedStreamServer interface {
	Send(*EmptyProto) error
	grpc.ServerStream
}

type flowCombinationRpcPathSingleNestedStreamServer struct {
	grpc.ServerStream
}

func (x *flowCombinationRpcPathSingleNestedStreamServer) Send(m *EmptyProto) error {
	return x.ServerStream.SendMsg(m)
}

func _FlowCombination_RpcPathNestedStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NestedProto)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FlowCombinationServer).RpcPathNestedStream(m, &flowCombinationRpcPathNestedStreamServer{stream})
}

type FlowCombination_RpcPathNestedStreamServer interface {
	Send(*EmptyProto) error
	grpc.ServerStream
}

type flowCombinationRpcPathNestedStreamServer struct {
	grpc.ServerStream
}

func (x *flowCombinationRpcPathNestedStreamServer) Send(m *EmptyProto) error {
	return x.ServerStream.SendMsg(m)
}

var _FlowCombination_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gengo.grpc.gateway.examples.examplepb.FlowCombination",
	HandlerType: (*FlowCombinationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RpcEmptyRpc",
			Handler:    _FlowCombination_RpcEmptyRpc_Handler,
		},
		{
			MethodName: "RpcBodyRpc",
			Handler:    _FlowCombination_RpcBodyRpc_Handler,
		},
		{
			MethodName: "RpcPathSingleNestedRpc",
			Handler:    _FlowCombination_RpcPathSingleNestedRpc_Handler,
		},
		{
			MethodName: "RpcPathNestedRpc",
			Handler:    _FlowCombination_RpcPathNestedRpc_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RpcEmptyStream",
			Handler:       _FlowCombination_RpcEmptyStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamEmptyRpc",
			Handler:       _FlowCombination_StreamEmptyRpc_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamEmptyStream",
			Handler:       _FlowCombination_StreamEmptyStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "RpcBodyStream",
			Handler:       _FlowCombination_RpcBodyStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RpcPathSingleNestedStream",
			Handler:       _FlowCombination_RpcPathSingleNestedStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RpcPathNestedStream",
			Handler:       _FlowCombination_RpcPathNestedStream_Handler,
			ServerStreams: true,
		},
	},
}

func init() { proto.RegisterFile("examples/examplepb/flow_combination.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 658 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x96, 0xbf, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x75, 0xa9, 0x84, 0xd4, 0x4b, 0x53, 0x52, 0x07, 0x85, 0x24, 0x84, 0xfe, 0x38, 0x8a,
	0x08, 0x48, 0xd8, 0x25, 0x30, 0x00, 0x0b, 0x52, 0x11, 0x8c, 0x55, 0xd5, 0xc2, 0xe2, 0x05, 0x9d,
	0x5d, 0xe3, 0x06, 0x25, 0x3e, 0xd7, 0x31, 0x0a, 0x55, 0xd5, 0x85, 0x95, 0xb1, 0x23, 0x7f, 0x02,
	0x23, 0x13, 0x6c, 0x0c, 0xcc, 0x0c, 0x95, 0xf8, 0x0b, 0x58, 0xf8, 0x0b, 0x58, 0xb9, 0x7b, 0x77,
	0xf6, 0xc5, 0x90, 0x10, 0x53, 0x90, 0x87, 0xaa, 0xf7, 0xeb, 0x3d, 0x7f, 0xde, 0xf7, 0x7d, 0xef,
	0x14, 0x7c, 0xdd, 0x7b, 0x45, 0x07, 0x61, 0xdf, 0x1b, 0x5a, 0x6a, 0x10, 0x3a, 0xd6, 0xf3, 0x3e,
	0x1b, 0x3d, 0x73, 0xd9, 0xc0, 0xe9, 0x05, 0x34, 0xee, 0xb1, 0xc0, 0x0c, 0x23, 0x16, 0x33, 0xe3,
	0xaa, 0xef, 0x05, 0x3e, 0x33, 0xfd, 0x28, 0x74, 0x4d, 0x9f, 0xc6, 0xde, 0x88, 0x1e, 0x9a, 0x49,
	0xb4, 0x99, 0x46, 0xb7, 0xda, 0x3e, 0x63, 0x7e, 0xdf, 0xb3, 0x68, 0xd8, 0xb3, 0x68, 0x10, 0xb0,
	0x18, 0x72, 0x0c, 0x65, 0x12, 0xb2, 0x80, 0xf1, 0xa3, 0x41, 0x18, 0x1f, 0x6e, 0xc3, 0xec, 0x1e,
	0xae, 0x6c, 0xb1, 0x40, 0x2f, 0x18, 0x0b, 0x18, 0xd1, 0x06, 0x5a, 0x45, 0x9d, 0xf9, 0x1d, 0x44,
	0xc5, 0xcc, 0x69, 0x94, 0xe4, 0xcc, 0x11, 0x33, 0xb7, 0x31, 0x27, 0x67, 0x2e, 0x59, 0xc6, 0xf8,
	0x69, 0x40, 0x23, 0x15, 0x57, 0xc5, 0x73, 0xc3, 0x38, 0x52, 0x91, 0x62, 0x48, 0x5e, 0xe0, 0xf2,
	0x96, 0x37, 0x8c, 0xbd, 0x3d, 0x79, 0xe0, 0x41, 0x92, 0xb8, 0xdc, 0xbd, 0x65, 0xe6, 0x2a, 0xc4,
	0xd4, 0xe9, 0x67, 0xb1, 0x3c, 0xc1, 0x4b, 0xbb, 0xbd, 0x80, 0x17, 0xfd, 0x3f, 0xbf, 0xd8, 0xfd,
	0x5c, 0xc5, 0xe7, 0x1f, 0xf3, 0x56, 0x3c, 0xd4, 0x9d, 0x30, 0xde, 0x20, 0x5c, 0xde, 0x09, 0x5d,
	0x50, 0x8c, 0xff, 0x37, 0xf2, 0x66, 0xd6, 0x12, 0xb7, 0xfe, 0x3e, 0x84, 0xd4, 0x5f, 0x9f, 0x7e,
	0x3b, 0x29, 0x55, 0xc9, 0xa2, 0xc5, 0x63, 0x2c, 0x4f, 0x6c, 0x88, 0x91, 0x71, 0x82, 0xf0, 0x62,
	0x42, 0xb3, 0x1b, 0x47, 0x1e, 0x1d, 0x14, 0x04, 0xd4, 0x04, 0xa0, 0x1a, 0x59, 0x1a, 0x03, 0x1a,
	0x02, 0xc0, 0x06, 0x02, 0x2a, 0x49, 0x53, 0xb0, 0x4c, 0x9a, 0x4a, 0xb2, 0x68, 0xa5, 0x3a, 0xc8,
	0x78, 0x8b, 0xb8, 0x49, 0x34, 0x55, 0xa1, 0x72, 0xb5, 0x01, 0xac, 0x4e, 0x2e, 0x64, 0xc1, 0xe4,
	0xa4, 0x83, 0xb8, 0x66, 0x1f, 0x4a, 0x18, 0x73, 0xa1, 0x36, 0xd9, 0x1e, 0xe8, 0x75, 0x27, 0xe7,
	0x37, 0x32, 0x97, 0xf7, 0x2c, 0x64, 0x9f, 0x10, 0xa0, 0x7d, 0x44, 0xa4, 0x02, 0xad, 0x74, 0x38,
	0x80, 0x18, 0xdc, 0x47, 0x37, 0xec, 0x4b, 0xa4, 0x09, 0x6b, 0x21, 0x8d, 0xf7, 0xad, 0x23, 0x7a,
	0x6c, 0x1d, 0x39, 0xfc, 0xcf, 0x3d, 0x16, 0x8b, 0x76, 0x62, 0xc6, 0x83, 0x97, 0x5e, 0x04, 0x11,
	0xf6, 0x0a, 0x69, 0xe9, 0x14, 0x99, 0x18, 0xc8, 0xe7, 0xda, 0x0d, 0x52, 0xd3, 0x07, 0xd2, 0x38,
	0xb1, 0xb3, 0x46, 0xda, 0x13, 0x42, 0x33, 0x47, 0x9a, 0xe4, 0x62, 0x16, 0x26, 0xdd, 0x35, 0xde,
	0x21, 0x5c, 0xe7, 0xa2, 0x6d, 0xf3, 0x9d, 0xf1, 0x57, 0x40, 0xe8, 0x78, 0x37, 0xa7, 0x22, 0xbf,
	0xbd, 0x1e, 0x67, 0xd1, 0x72, 0x1d, 0xa4, 0x5c, 0x56, 0xb5, 0x08, 0xd0, 0x9b, 0x01, 0x64, 0xe4,
	0xbc, 0x26, 0xef, 0x35, 0x08, 0x61, 0x7c, 0x47, 0xb8, 0xaa, 0x68, 0x35, 0x67, 0x37, 0x6f, 0xbf,
	0xff, 0x8d, 0x30, 0x00, 0xc2, 0x7d, 0xb2, 0x3a, 0x95, 0x70, 0xac, 0x5d, 0x33, 0x0a, 0x49, 0x9b,
	0x36, 0x65, 0x9f, 0xa7, 0x30, 0x4e, 0x4b, 0xb8, 0xa2, 0x5c, 0xad, 0xee, 0x5b, 0x61, 0xc6, 0xfe,
	0x2a, 0x8d, 0xfd, 0x05, 0x91, 0xaa, 0xb6, 0x96, 0xbc, 0x70, 0xc2, 0xdb, 0xe3, 0xc5, 0x65, 0xbc,
	0x2d, 0x8f, 0xd8, 0xc9, 0xd3, 0x26, 0x5d, 0xa6, 0x16, 0x09, 0xb9, 0x3c, 0xc5, 0xe1, 0x49, 0x62,
	0x97, 0x5f, 0x9a, 0xfa, 0xaf, 0x26, 0xd7, 0x9b, 0xeb, 0x64, 0x65, 0xaa, 0xcf, 0xf5, 0xa9, 0xb6,
	0xba, 0x48, 0x13, 0x0f, 0xf0, 0xb7, 0xe2, 0x3d, 0xc2, 0xcd, 0x09, 0x7e, 0x57, 0x0a, 0x17, 0x6a,
	0xf9, 0x6b, 0x20, 0xf2, 0x9a, 0x2a, 0x6b, 0x92, 0x13, 0x52, 0xea, 0x1f, 0x08, 0xd7, 0x32, 0xbe,
	0x57, 0xbc, 0x05, 0x59, 0x7f, 0x04, 0xa4, 0x07, 0xe4, 0xca, 0x1f, 0xad, 0xaf, 0x9b, 0x30, 0xbb,
	0xa6, 0xb4, 0x9b, 0xd3, 0x8f, 0xf0, 0x44, 0x1b, 0x68, 0xb3, 0x6c, 0xcf, 0xa7, 0x48, 0xce, 0x39,
	0xf8, 0x19, 0x76, 0xfb, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x90, 0xd0, 0xd7, 0xf8, 0x09,
	0x00, 0x00,
}
