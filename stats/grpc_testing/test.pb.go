// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc_testing/test.proto

package grpc_testing

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc/1291"
	codes "google.golang.org/grpc/1291/codes"
	status "google.golang.org/grpc/1291/status"
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

type SimpleRequest struct {
	Id                   int32    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleRequest) Reset()         { *m = SimpleRequest{} }
func (m *SimpleRequest) String() string { return proto.CompactTextString(m) }
func (*SimpleRequest) ProtoMessage()    {}
func (*SimpleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1cda82041fed8bf, []int{0}
}

func (m *SimpleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleRequest.Unmarshal(m, b)
}
func (m *SimpleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleRequest.Marshal(b, m, deterministic)
}
func (m *SimpleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleRequest.Merge(m, src)
}
func (m *SimpleRequest) XXX_Size() int {
	return xxx_messageInfo_SimpleRequest.Size(m)
}
func (m *SimpleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleRequest proto.InternalMessageInfo

func (m *SimpleRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type SimpleResponse struct {
	Id                   int32    `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleResponse) Reset()         { *m = SimpleResponse{} }
func (m *SimpleResponse) String() string { return proto.CompactTextString(m) }
func (*SimpleResponse) ProtoMessage()    {}
func (*SimpleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1cda82041fed8bf, []int{1}
}

func (m *SimpleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleResponse.Unmarshal(m, b)
}
func (m *SimpleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleResponse.Marshal(b, m, deterministic)
}
func (m *SimpleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleResponse.Merge(m, src)
}
func (m *SimpleResponse) XXX_Size() int {
	return xxx_messageInfo_SimpleResponse.Size(m)
}
func (m *SimpleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleResponse proto.InternalMessageInfo

func (m *SimpleResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*SimpleRequest)(nil), "grpc.testing.SimpleRequest")
	proto.RegisterType((*SimpleResponse)(nil), "grpc.testing.SimpleResponse")
}

func init() { proto.RegisterFile("grpc_testing/test.proto", fileDescriptor_e1cda82041fed8bf) }

var fileDescriptor_e1cda82041fed8bf = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0x2f, 0x2a, 0x48,
	0x8e, 0x2f, 0x49, 0x2d, 0x2e, 0xc9, 0xcc, 0x4b, 0xd7, 0x07, 0xd1, 0x7a, 0x05, 0x45, 0xf9, 0x25,
	0xf9, 0x42, 0x3c, 0x20, 0x09, 0x3d, 0xa8, 0x84, 0x92, 0x3c, 0x17, 0x6f, 0x70, 0x66, 0x6e, 0x41,
	0x4e, 0x6a, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x10, 0x1f, 0x17, 0x53, 0x66, 0x8a, 0x04,
	0x93, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x53, 0x66, 0x8a, 0x92, 0x02, 0x17, 0x1f, 0x4c, 0x41, 0x71,
	0x41, 0x7e, 0x5e, 0x71, 0x2a, 0x54, 0x05, 0x33, 0x4c, 0x85, 0xd1, 0x09, 0x26, 0x2e, 0xee, 0x90,
	0xd4, 0xe2, 0x92, 0xe0, 0xd4, 0xa2, 0xb2, 0xcc, 0xe4, 0x54, 0x21, 0x37, 0x2e, 0xce, 0xd0, 0xbc,
	0xc4, 0xa2, 0x4a, 0xe7, 0xc4, 0x9c, 0x1c, 0x21, 0x69, 0x3d, 0x64, 0xeb, 0xf4, 0x50, 0xec, 0x92,
	0x92, 0xc1, 0x2e, 0x09, 0xb5, 0xc7, 0x9f, 0x8b, 0xcf, 0xad, 0x34, 0x27, 0xc7, 0xa5, 0xb4, 0x20,
	0x27, 0xb5, 0x82, 0x42, 0xc3, 0x34, 0x18, 0x0d, 0x18, 0x85, 0xfc, 0xb9, 0x04, 0x9c, 0x73, 0x32,
	0x53, 0xf3, 0x4a, 0x82, 0x4b, 0x8a, 0x52, 0x13, 0x73, 0x29, 0x36, 0x12, 0x64, 0x20, 0xc8, 0xd3,
	0xa9, 0x45, 0x54, 0x31, 0xd0, 0x80, 0x31, 0x89, 0x0d, 0x1c, 0x45, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x4c, 0x43, 0x27, 0x67, 0xbd, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TestServiceClient is the client API for TestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestServiceClient interface {
	// One request followed by one response.
	// The server returns the client id as-is.
	UnaryCall(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error)
	// A sequence of requests with each request served by the server immediately.
	// As one request could lead to multiple responses, this interface
	// demonstrates the idea of full duplexing.
	FullDuplexCall(ctx context.Context, opts ...grpc.CallOption) (TestService_FullDuplexCallClient, error)
	// Client stream
	ClientStreamCall(ctx context.Context, opts ...grpc.CallOption) (TestService_ClientStreamCallClient, error)
	// Server stream
	ServerStreamCall(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (TestService_ServerStreamCallClient, error)
}

type testServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTestServiceClient(cc grpc.ClientConnInterface) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) UnaryCall(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error) {
	out := new(SimpleResponse)
	err := c.cc.Invoke(ctx, "/grpc.testing.TestService/UnaryCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) FullDuplexCall(ctx context.Context, opts ...grpc.CallOption) (TestService_FullDuplexCallClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TestService_serviceDesc.Streams[0], "/grpc.testing.TestService/FullDuplexCall", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceFullDuplexCallClient{stream}
	return x, nil
}

type TestService_FullDuplexCallClient interface {
	Send(*SimpleRequest) error
	Recv() (*SimpleResponse, error)
	grpc.ClientStream
}

type testServiceFullDuplexCallClient struct {
	grpc.ClientStream
}

func (x *testServiceFullDuplexCallClient) Send(m *SimpleRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServiceFullDuplexCallClient) Recv() (*SimpleResponse, error) {
	m := new(SimpleResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) ClientStreamCall(ctx context.Context, opts ...grpc.CallOption) (TestService_ClientStreamCallClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TestService_serviceDesc.Streams[1], "/grpc.testing.TestService/ClientStreamCall", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceClientStreamCallClient{stream}
	return x, nil
}

type TestService_ClientStreamCallClient interface {
	Send(*SimpleRequest) error
	CloseAndRecv() (*SimpleResponse, error)
	grpc.ClientStream
}

type testServiceClientStreamCallClient struct {
	grpc.ClientStream
}

func (x *testServiceClientStreamCallClient) Send(m *SimpleRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServiceClientStreamCallClient) CloseAndRecv() (*SimpleResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SimpleResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) ServerStreamCall(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (TestService_ServerStreamCallClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TestService_serviceDesc.Streams[2], "/grpc.testing.TestService/ServerStreamCall", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceServerStreamCallClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TestService_ServerStreamCallClient interface {
	Recv() (*SimpleResponse, error)
	grpc.ClientStream
}

type testServiceServerStreamCallClient struct {
	grpc.ClientStream
}

func (x *testServiceServerStreamCallClient) Recv() (*SimpleResponse, error) {
	m := new(SimpleResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TestServiceServer is the server API for TestService service.
type TestServiceServer interface {
	// One request followed by one response.
	// The server returns the client id as-is.
	UnaryCall(context.Context, *SimpleRequest) (*SimpleResponse, error)
	// A sequence of requests with each request served by the server immediately.
	// As one request could lead to multiple responses, this interface
	// demonstrates the idea of full duplexing.
	FullDuplexCall(TestService_FullDuplexCallServer) error
	// Client stream
	ClientStreamCall(TestService_ClientStreamCallServer) error
	// Server stream
	ServerStreamCall(*SimpleRequest, TestService_ServerStreamCallServer) error
}

// UnimplementedTestServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTestServiceServer struct {
}

func (*UnimplementedTestServiceServer) UnaryCall(ctx context.Context, req *SimpleRequest) (*SimpleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnaryCall not implemented")
}
func (*UnimplementedTestServiceServer) FullDuplexCall(srv TestService_FullDuplexCallServer) error {
	return status.Errorf(codes.Unimplemented, "method FullDuplexCall not implemented")
}
func (*UnimplementedTestServiceServer) ClientStreamCall(srv TestService_ClientStreamCallServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientStreamCall not implemented")
}
func (*UnimplementedTestServiceServer) ServerStreamCall(req *SimpleRequest, srv TestService_ServerStreamCallServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerStreamCall not implemented")
}

func RegisterTestServiceServer(s *grpc.Server, srv TestServiceServer) {
	s.RegisterService(&_TestService_serviceDesc, srv)
}

func _TestService_UnaryCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).UnaryCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.testing.TestService/UnaryCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).UnaryCall(ctx, req.(*SimpleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_FullDuplexCall_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).FullDuplexCall(&testServiceFullDuplexCallServer{stream})
}

type TestService_FullDuplexCallServer interface {
	Send(*SimpleResponse) error
	Recv() (*SimpleRequest, error)
	grpc.ServerStream
}

type testServiceFullDuplexCallServer struct {
	grpc.ServerStream
}

func (x *testServiceFullDuplexCallServer) Send(m *SimpleResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServiceFullDuplexCallServer) Recv() (*SimpleRequest, error) {
	m := new(SimpleRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TestService_ClientStreamCall_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).ClientStreamCall(&testServiceClientStreamCallServer{stream})
}

type TestService_ClientStreamCallServer interface {
	SendAndClose(*SimpleResponse) error
	Recv() (*SimpleRequest, error)
	grpc.ServerStream
}

type testServiceClientStreamCallServer struct {
	grpc.ServerStream
}

func (x *testServiceClientStreamCallServer) SendAndClose(m *SimpleResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServiceClientStreamCallServer) Recv() (*SimpleRequest, error) {
	m := new(SimpleRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TestService_ServerStreamCall_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SimpleRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestServiceServer).ServerStreamCall(m, &testServiceServerStreamCallServer{stream})
}

type TestService_ServerStreamCallServer interface {
	Send(*SimpleResponse) error
	grpc.ServerStream
}

type testServiceServerStreamCallServer struct {
	grpc.ServerStream
}

func (x *testServiceServerStreamCallServer) Send(m *SimpleResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _TestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.testing.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnaryCall",
			Handler:    _TestService_UnaryCall_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FullDuplexCall",
			Handler:       _TestService_FullDuplexCall_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ClientStreamCall",
			Handler:       _TestService_ClientStreamCall_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ServerStreamCall",
			Handler:       _TestService_ServerStreamCall_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "grpc_testing/test.proto",
}
