// Code generated by protoc-gen-go. DO NOT EDIT.
// source: connection-service.proto

package connection

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

type Connection struct {
	Id                   int64    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Host                 string   `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	Port                 string   `protobuf:"bytes,3,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Connection) Reset()         { *m = Connection{} }
func (m *Connection) String() string { return proto.CompactTextString(m) }
func (*Connection) ProtoMessage()    {}
func (*Connection) Descriptor() ([]byte, []int) {
	return fileDescriptor_80f86ad761877d5c, []int{0}
}

func (m *Connection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Connection.Unmarshal(m, b)
}
func (m *Connection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Connection.Marshal(b, m, deterministic)
}
func (m *Connection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Connection.Merge(m, src)
}
func (m *Connection) XXX_Size() int {
	return xxx_messageInfo_Connection.Size(m)
}
func (m *Connection) XXX_DiscardUnknown() {
	xxx_messageInfo_Connection.DiscardUnknown(m)
}

var xxx_messageInfo_Connection proto.InternalMessageInfo

func (m *Connection) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Connection) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Connection) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

type ClickResponse struct {
	Click                *Click      `protobuf:"bytes,1,opt,name=click,proto3" json:"click,omitempty"`
	Connection           *Connection `protobuf:"bytes,2,opt,name=connection,proto3" json:"connection,omitempty"`
	Error                *Error      `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ClickResponse) Reset()         { *m = ClickResponse{} }
func (m *ClickResponse) String() string { return proto.CompactTextString(m) }
func (*ClickResponse) ProtoMessage()    {}
func (*ClickResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_80f86ad761877d5c, []int{1}
}

func (m *ClickResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClickResponse.Unmarshal(m, b)
}
func (m *ClickResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClickResponse.Marshal(b, m, deterministic)
}
func (m *ClickResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClickResponse.Merge(m, src)
}
func (m *ClickResponse) XXX_Size() int {
	return xxx_messageInfo_ClickResponse.Size(m)
}
func (m *ClickResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClickResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClickResponse proto.InternalMessageInfo

func (m *ClickResponse) GetClick() *Click {
	if m != nil {
		return m.Click
	}
	return nil
}

func (m *ClickResponse) GetConnection() *Connection {
	if m != nil {
		return m.Connection
	}
	return nil
}

func (m *ClickResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type ClickData struct {
	Id                   int64       `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Click                *Click      `protobuf:"bytes,2,opt,name=click,proto3" json:"click,omitempty"`
	Connection           *Connection `protobuf:"bytes,3,opt,name=connection,proto3" json:"connection,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ClickData) Reset()         { *m = ClickData{} }
func (m *ClickData) String() string { return proto.CompactTextString(m) }
func (*ClickData) ProtoMessage()    {}
func (*ClickData) Descriptor() ([]byte, []int) {
	return fileDescriptor_80f86ad761877d5c, []int{2}
}

func (m *ClickData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClickData.Unmarshal(m, b)
}
func (m *ClickData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClickData.Marshal(b, m, deterministic)
}
func (m *ClickData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClickData.Merge(m, src)
}
func (m *ClickData) XXX_Size() int {
	return xxx_messageInfo_ClickData.Size(m)
}
func (m *ClickData) XXX_DiscardUnknown() {
	xxx_messageInfo_ClickData.DiscardUnknown(m)
}

var xxx_messageInfo_ClickData proto.InternalMessageInfo

func (m *ClickData) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ClickData) GetClick() *Click {
	if m != nil {
		return m.Click
	}
	return nil
}

func (m *ClickData) GetConnection() *Connection {
	if m != nil {
		return m.Connection
	}
	return nil
}

type Error struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_80f86ad761877d5c, []int{3}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ClickSummary struct {
	// number of clicks received.
	ClickCount int32 `protobuf:"varint,1,opt,name=click_count,json=clickCount,proto3" json:"click_count,omitempty"`
	// The duration in seconds.
	ElapsedTime          int32    `protobuf:"varint,2,opt,name=elapsed_time,json=elapsedTime,proto3" json:"elapsed_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClickSummary) Reset()         { *m = ClickSummary{} }
func (m *ClickSummary) String() string { return proto.CompactTextString(m) }
func (*ClickSummary) ProtoMessage()    {}
func (*ClickSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_80f86ad761877d5c, []int{4}
}

func (m *ClickSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClickSummary.Unmarshal(m, b)
}
func (m *ClickSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClickSummary.Marshal(b, m, deterministic)
}
func (m *ClickSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClickSummary.Merge(m, src)
}
func (m *ClickSummary) XXX_Size() int {
	return xxx_messageInfo_ClickSummary.Size(m)
}
func (m *ClickSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_ClickSummary.DiscardUnknown(m)
}

var xxx_messageInfo_ClickSummary proto.InternalMessageInfo

func (m *ClickSummary) GetClickCount() int32 {
	if m != nil {
		return m.ClickCount
	}
	return 0
}

func (m *ClickSummary) GetElapsedTime() int32 {
	if m != nil {
		return m.ElapsedTime
	}
	return 0
}

func init() {
	proto.RegisterType((*Connection)(nil), "connection.Connection")
	proto.RegisterType((*ClickResponse)(nil), "connection.ClickResponse")
	proto.RegisterType((*ClickData)(nil), "connection.ClickData")
	proto.RegisterType((*Error)(nil), "connection.Error")
	proto.RegisterType((*ClickSummary)(nil), "connection.ClickSummary")
}

func init() { proto.RegisterFile("connection-service.proto", fileDescriptor_80f86ad761877d5c) }

var fileDescriptor_80f86ad761877d5c = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0x25, 0xe9, 0x97, 0xef, 0x23, 0x93, 0x7e, 0x4a, 0x17, 0x94, 0xd8, 0x8b, 0x35, 0x17, 0x7b,
	0xb1, 0x42, 0x45, 0x6f, 0x5e, 0x6c, 0x45, 0x7a, 0xdd, 0x78, 0xf2, 0x52, 0xe2, 0x66, 0xa8, 0xc1,
	0x26, 0x1b, 0x77, 0x53, 0x41, 0xf0, 0x8f, 0x88, 0x7f, 0x56, 0x76, 0xb6, 0x4d, 0xaa, 0x41, 0xd0,
	0xdb, 0xe4, 0xbd, 0xc9, 0x7b, 0x6f, 0x66, 0x07, 0x42, 0x21, 0x8b, 0x02, 0x45, 0x95, 0xc9, 0xe2,
	0x44, 0xa3, 0x7a, 0xce, 0x04, 0x8e, 0x4a, 0x25, 0x2b, 0xc9, 0xa0, 0x61, 0xfa, 0x81, 0x58, 0x66,
	0xe2, 0xd1, 0x12, 0xd1, 0x14, 0x60, 0x52, 0x53, 0x6c, 0x07, 0xdc, 0x59, 0x1a, 0x3a, 0x03, 0x67,
	0xd8, 0xe1, 0xee, 0x2c, 0x65, 0x0c, 0xfe, 0x3c, 0x48, 0x5d, 0x85, 0xee, 0xc0, 0x19, 0xfa, 0x9c,
	0x6a, 0x83, 0x95, 0x52, 0x55, 0x61, 0xc7, 0x62, 0xa6, 0x8e, 0xde, 0x1c, 0xf8, 0x3f, 0x31, 0xaa,
	0x1c, 0x75, 0x29, 0x0b, 0x8d, 0xec, 0x18, 0x3c, 0xb2, 0x21, 0xb1, 0x60, 0xdc, 0x1b, 0x35, 0x01,
	0x46, 0xb6, 0xd3, 0xf2, 0xec, 0x02, 0xb6, 0xb2, 0x91, 0x51, 0x30, 0xde, 0xff, 0xd4, 0x5d, 0x97,
	0x7c, 0xab, 0xd3, 0x18, 0xa0, 0x52, 0x52, 0x51, 0x8e, 0x2f, 0x06, 0xd7, 0x86, 0xe0, 0x96, 0x8f,
	0x5e, 0xc1, 0x27, 0xc3, 0x69, 0x52, 0x25, 0xad, 0x01, 0xeb, 0x98, 0xee, 0xaf, 0x62, 0x76, 0x7e,
	0x1a, 0x33, 0x3a, 0x07, 0x8f, 0xd2, 0x98, 0xb5, 0x09, 0x99, 0x22, 0x79, 0xfb, 0x9c, 0x6a, 0x16,
	0xc2, 0xbf, 0x1c, 0xb5, 0x4e, 0x16, 0xb8, 0xde, 0xf0, 0xe6, 0x33, 0xe2, 0xd0, 0x25, 0xfb, 0x78,
	0x95, 0xe7, 0x89, 0x7a, 0x61, 0x87, 0x60, 0x5f, 0x6d, 0x2e, 0xe4, 0xaa, 0xa8, 0x48, 0xc4, 0xe3,
	0x40, 0xd0, 0xc4, 0x20, 0xec, 0x08, 0xba, 0xb8, 0x4c, 0x4a, 0x8d, 0xe9, 0xbc, 0xca, 0x72, 0xab,
	0xe7, 0xf1, 0x60, 0x8d, 0xdd, 0x66, 0x39, 0x8e, 0xdf, 0x1d, 0xe8, 0x35, 0x29, 0x63, 0x7b, 0x1f,
	0xec, 0x12, 0xfc, 0x18, 0x8b, 0x94, 0xdc, 0xd8, 0x5e, 0x6b, 0x7e, 0xb3, 0xb5, 0xfe, 0x41, 0x7b,
	0x2d, 0x9b, 0x77, 0xbe, 0x01, 0x56, 0xff, 0xae, 0x63, 0x7c, 0x5a, 0x61, 0x21, 0xf0, 0x3b, 0x9d,
	0xb0, 0x05, 0xaf, 0xe7, 0x1b, 0x3a, 0x57, 0xbd, 0xbb, 0xdd, 0x85, 0x2a, 0xc5, 0x69, 0xd3, 0x71,
	0xff, 0x97, 0x4e, 0xf4, 0xec, 0x23, 0x00, 0x00, 0xff, 0xff, 0x1c, 0xf4, 0x55, 0x10, 0xd7, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ConnectionServiceClient is the client API for ConnectionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConnectionServiceClient interface {
	// A simple RPC.
	SendClick(ctx context.Context, in *ClickData, opts ...grpc.CallOption) (*ClickResponse, error)
	// A client-to-server streaming RPC.
	SendClicksSequence(ctx context.Context, opts ...grpc.CallOption) (ConnectionService_SendClicksSequenceClient, error)
}

type connectionServiceClient struct {
	cc *grpc.ClientConn
}

func NewConnectionServiceClient(cc *grpc.ClientConn) ConnectionServiceClient {
	return &connectionServiceClient{cc}
}

func (c *connectionServiceClient) SendClick(ctx context.Context, in *ClickData, opts ...grpc.CallOption) (*ClickResponse, error) {
	out := new(ClickResponse)
	err := c.cc.Invoke(ctx, "/connection.ConnectionService/SendClick", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) SendClicksSequence(ctx context.Context, opts ...grpc.CallOption) (ConnectionService_SendClicksSequenceClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ConnectionService_serviceDesc.Streams[0], "/connection.ConnectionService/SendClicksSequence", opts...)
	if err != nil {
		return nil, err
	}
	x := &connectionServiceSendClicksSequenceClient{stream}
	return x, nil
}

type ConnectionService_SendClicksSequenceClient interface {
	Send(*ClickData) error
	CloseAndRecv() (*ClickSummary, error)
	grpc.ClientStream
}

type connectionServiceSendClicksSequenceClient struct {
	grpc.ClientStream
}

func (x *connectionServiceSendClicksSequenceClient) Send(m *ClickData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *connectionServiceSendClicksSequenceClient) CloseAndRecv() (*ClickSummary, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ClickSummary)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ConnectionServiceServer is the server API for ConnectionService service.
type ConnectionServiceServer interface {
	// A simple RPC.
	SendClick(context.Context, *ClickData) (*ClickResponse, error)
	// A client-to-server streaming RPC.
	SendClicksSequence(ConnectionService_SendClicksSequenceServer) error
}

// UnimplementedConnectionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedConnectionServiceServer struct {
}

func (*UnimplementedConnectionServiceServer) SendClick(ctx context.Context, req *ClickData) (*ClickResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendClick not implemented")
}
func (*UnimplementedConnectionServiceServer) SendClicksSequence(srv ConnectionService_SendClicksSequenceServer) error {
	return status.Errorf(codes.Unimplemented, "method SendClicksSequence not implemented")
}

func RegisterConnectionServiceServer(s *grpc.Server, srv ConnectionServiceServer) {
	s.RegisterService(&_ConnectionService_serviceDesc, srv)
}

func _ConnectionService_SendClick_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClickData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).SendClick(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection.ConnectionService/SendClick",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).SendClick(ctx, req.(*ClickData))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_SendClicksSequence_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ConnectionServiceServer).SendClicksSequence(&connectionServiceSendClicksSequenceServer{stream})
}

type ConnectionService_SendClicksSequenceServer interface {
	SendAndClose(*ClickSummary) error
	Recv() (*ClickData, error)
	grpc.ServerStream
}

type connectionServiceSendClicksSequenceServer struct {
	grpc.ServerStream
}

func (x *connectionServiceSendClicksSequenceServer) SendAndClose(m *ClickSummary) error {
	return x.ServerStream.SendMsg(m)
}

func (x *connectionServiceSendClicksSequenceServer) Recv() (*ClickData, error) {
	m := new(ClickData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ConnectionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "connection.ConnectionService",
	HandlerType: (*ConnectionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendClick",
			Handler:    _ConnectionService_SendClick_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendClicksSequence",
			Handler:       _ConnectionService_SendClicksSequence_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "connection-service.proto",
}
