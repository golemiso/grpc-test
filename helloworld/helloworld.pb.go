// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helloworld.proto

package helloworld

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type HelloRequest struct {
	Name                 string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	StrVal               *wrappers.StringValue `protobuf:"bytes,2,opt,name=strVal,proto3" json:"strVal,omitempty"`
	FloatVal             *wrappers.FloatValue  `protobuf:"bytes,3,opt,name=floatVal,proto3" json:"floatVal,omitempty"`
	DoubleVal            *wrappers.DoubleValue `protobuf:"bytes,4,opt,name=doubleVal,proto3" json:"doubleVal,omitempty"`
	BoolVal              *wrappers.BoolValue   `protobuf:"bytes,5,opt,name=boolVal,proto3" json:"boolVal,omitempty"`
	BytesVal             *wrappers.BytesValue  `protobuf:"bytes,6,opt,name=bytesVal,proto3" json:"bytesVal,omitempty"`
	Int32Val             *wrappers.Int32Value  `protobuf:"bytes,7,opt,name=int32Val,proto3" json:"int32Val,omitempty"`
	Uint32Val            *wrappers.UInt32Value `protobuf:"bytes,8,opt,name=uint32Val,proto3" json:"uint32Val,omitempty"`
	Int64Val             *wrappers.Int64Value  `protobuf:"bytes,9,opt,name=int64Val,proto3" json:"int64Val,omitempty"`
	Uint64Val            *wrappers.UInt64Value `protobuf:"bytes,10,opt,name=uint64Val,proto3" json:"uint64Val,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HelloRequest) GetStrVal() *wrappers.StringValue {
	if m != nil {
		return m.StrVal
	}
	return nil
}

func (m *HelloRequest) GetFloatVal() *wrappers.FloatValue {
	if m != nil {
		return m.FloatVal
	}
	return nil
}

func (m *HelloRequest) GetDoubleVal() *wrappers.DoubleValue {
	if m != nil {
		return m.DoubleVal
	}
	return nil
}

func (m *HelloRequest) GetBoolVal() *wrappers.BoolValue {
	if m != nil {
		return m.BoolVal
	}
	return nil
}

func (m *HelloRequest) GetBytesVal() *wrappers.BytesValue {
	if m != nil {
		return m.BytesVal
	}
	return nil
}

func (m *HelloRequest) GetInt32Val() *wrappers.Int32Value {
	if m != nil {
		return m.Int32Val
	}
	return nil
}

func (m *HelloRequest) GetUint32Val() *wrappers.UInt32Value {
	if m != nil {
		return m.Uint32Val
	}
	return nil
}

func (m *HelloRequest) GetInt64Val() *wrappers.Int64Value {
	if m != nil {
		return m.Int64Val
	}
	return nil
}

func (m *HelloRequest) GetUint64Val() *wrappers.UInt64Value {
	if m != nil {
		return m.Uint64Val
	}
	return nil
}

type HelloReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{1}
}

func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (m *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(m, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type TestRequest struct {
	// Types that are valid to be assigned to Params:
	//	*TestRequest_Name
	//	*TestRequest_City
	Params               isTestRequest_Params `protobuf_oneof:"params"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TestRequest) Reset()         { *m = TestRequest{} }
func (m *TestRequest) String() string { return proto.CompactTextString(m) }
func (*TestRequest) ProtoMessage()    {}
func (*TestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{2}
}

func (m *TestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestRequest.Unmarshal(m, b)
}
func (m *TestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestRequest.Marshal(b, m, deterministic)
}
func (m *TestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestRequest.Merge(m, src)
}
func (m *TestRequest) XXX_Size() int {
	return xxx_messageInfo_TestRequest.Size(m)
}
func (m *TestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TestRequest proto.InternalMessageInfo

type isTestRequest_Params interface {
	isTestRequest_Params()
}

type TestRequest_Name struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3,oneof"`
}

type TestRequest_City struct {
	City *TestMessage `protobuf:"bytes,2,opt,name=city,proto3,oneof"`
}

func (*TestRequest_Name) isTestRequest_Params() {}

func (*TestRequest_City) isTestRequest_Params() {}

func (m *TestRequest) GetParams() isTestRequest_Params {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *TestRequest) GetName() string {
	if x, ok := m.GetParams().(*TestRequest_Name); ok {
		return x.Name
	}
	return ""
}

func (m *TestRequest) GetCity() *TestMessage {
	if x, ok := m.GetParams().(*TestRequest_City); ok {
		return x.City
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TestRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TestRequest_Name)(nil),
		(*TestRequest_City)(nil),
	}
}

type TestReply struct {
	Req                  *TestRequest `protobuf:"bytes,1,opt,name=req,proto3" json:"req,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *TestReply) Reset()         { *m = TestReply{} }
func (m *TestReply) String() string { return proto.CompactTextString(m) }
func (*TestReply) ProtoMessage()    {}
func (*TestReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{3}
}

func (m *TestReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestReply.Unmarshal(m, b)
}
func (m *TestReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestReply.Marshal(b, m, deterministic)
}
func (m *TestReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestReply.Merge(m, src)
}
func (m *TestReply) XXX_Size() int {
	return xxx_messageInfo_TestReply.Size(m)
}
func (m *TestReply) XXX_DiscardUnknown() {
	xxx_messageInfo_TestReply.DiscardUnknown(m)
}

var xxx_messageInfo_TestReply proto.InternalMessageInfo

func (m *TestReply) GetReq() *TestRequest {
	if m != nil {
		return m.Req
	}
	return nil
}

type TestMessage struct {
	Me                   string   `protobuf:"bytes,3,opt,name=me,proto3" json:"me,omitempty"`
	Meme                 []string `protobuf:"bytes,4,rep,name=meme,proto3" json:"meme,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestMessage) Reset()         { *m = TestMessage{} }
func (m *TestMessage) String() string { return proto.CompactTextString(m) }
func (*TestMessage) ProtoMessage()    {}
func (*TestMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{4}
}

func (m *TestMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestMessage.Unmarshal(m, b)
}
func (m *TestMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestMessage.Marshal(b, m, deterministic)
}
func (m *TestMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestMessage.Merge(m, src)
}
func (m *TestMessage) XXX_Size() int {
	return xxx_messageInfo_TestMessage.Size(m)
}
func (m *TestMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_TestMessage.DiscardUnknown(m)
}

var xxx_messageInfo_TestMessage proto.InternalMessageInfo

func (m *TestMessage) GetMe() string {
	if m != nil {
		return m.Me
	}
	return ""
}

func (m *TestMessage) GetMeme() []string {
	if m != nil {
		return m.Meme
	}
	return nil
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "helloworld.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "helloworld.HelloReply")
	proto.RegisterType((*TestRequest)(nil), "helloworld.TestRequest")
	proto.RegisterType((*TestReply)(nil), "helloworld.TestReply")
	proto.RegisterType((*TestMessage)(nil), "helloworld.TestMessage")
}

func init() { proto.RegisterFile("helloworld.proto", fileDescriptor_17b8c58d586b62f2) }

var fileDescriptor_17b8c58d586b62f2 = []byte{
	// 589 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xcd, 0x6e, 0xd3, 0x4e,
	0x14, 0xc5, 0xff, 0x75, 0xdc, 0x7c, 0x4c, 0xff, 0x20, 0x34, 0x2a, 0xc5, 0x32, 0x15, 0x8a, 0xbc,
	0x40, 0x65, 0x41, 0x2c, 0xd2, 0x28, 0x48, 0x2c, 0x23, 0x3e, 0xca, 0x82, 0x8d, 0x0b, 0x5d, 0x44,
	0x6c, 0x26, 0xf4, 0x36, 0x44, 0x1a, 0x7b, 0x5c, 0x7b, 0x4c, 0x65, 0x55, 0xd9, 0xb0, 0x63, 0x87,
	0xc4, 0x1b, 0xb0, 0xe0, 0x75, 0x58, 0xf0, 0x0a, 0x3c, 0x08, 0xba, 0xd7, 0x33, 0xb1, 0x51, 0x6a,
	0x76, 0x33, 0x3e, 0xe7, 0xdc, 0x7b, 0x3d, 0xf3, 0x1b, 0x76, 0xe7, 0x23, 0x48, 0xa9, 0xae, 0x54,
	0x26, 0xcf, 0x47, 0x69, 0xa6, 0xb4, 0xe2, 0xac, 0xfe, 0xe2, 0x1f, 0x2e, 0x95, 0x5a, 0x4a, 0x08,
	0x45, 0xba, 0x0a, 0x45, 0x92, 0x28, 0x2d, 0xf4, 0x4a, 0x25, 0x79, 0xe5, 0xf4, 0x1f, 0x18, 0x95,
	0x76, 0x8b, 0xe2, 0x22, 0xbc, 0xca, 0x44, 0x9a, 0x42, 0x66, 0xf4, 0xe0, 0x87, 0xcb, 0xfe, 0x3f,
	0xc1, 0x62, 0x11, 0x5c, 0x16, 0x90, 0x6b, 0xce, 0x99, 0x9b, 0x88, 0x18, 0xbc, 0x9d, 0xe1, 0xce,
	0xd1, 0x20, 0xa2, 0x35, 0x9f, 0xb0, 0x6e, 0xae, 0xb3, 0x33, 0x21, 0x3d, 0x67, 0xb8, 0x73, 0xb4,
	0x37, 0x3e, 0x1c, 0x55, 0x55, 0x47, 0xb6, 0xea, 0xe8, 0x54, 0x67, 0xab, 0x64, 0x79, 0x26, 0x64,
	0x01, 0x91, 0xf1, 0xf2, 0xa7, 0xac, 0x7f, 0x21, 0x95, 0xd0, 0x98, 0xeb, 0x50, 0xee, 0xfe, 0x56,
	0xee, 0xa5, 0x31, 0x14, 0x10, 0x6d, 0xcc, 0xfc, 0x19, 0x1b, 0x9c, 0xab, 0x62, 0x21, 0x01, 0x93,
	0x6e, 0x4b, 0xc7, 0xe7, 0xd6, 0x51, 0x40, 0x54, 0xdb, 0xf9, 0x84, 0xf5, 0x16, 0x4a, 0x49, 0x4c,
	0xee, 0x52, 0xd2, 0xdf, 0x4a, 0xce, 0x2a, 0xbd, 0x80, 0xc8, 0x5a, 0x71, 0xd4, 0x45, 0xa9, 0x21,
	0xc7, 0x58, 0xb7, 0x65, 0xd4, 0x99, 0x31, 0xe0, 0xa8, 0xd6, 0x8c, 0xc1, 0x55, 0xa2, 0x8f, 0xc7,
	0x18, 0xec, 0xb5, 0x04, 0x5f, 0x1b, 0x03, 0x06, 0xad, 0x19, 0xff, 0xb1, 0xd8, 0x24, 0xfb, 0x2d,
	0xff, 0xf8, 0xae, 0x11, 0xad, 0xed, 0xa6, 0xe9, 0x74, 0x82, 0xd1, 0x41, 0x7b, 0x53, 0x32, 0x98,
	0xa6, 0xb4, 0xb6, 0x4d, 0xab, 0x24, 0xfb, 0x47, 0x53, 0x1b, 0xad, 0xed, 0xc1, 0x43, 0xc6, 0x0c,
	0x27, 0xa9, 0x2c, 0xb9, 0xc7, 0x7a, 0x31, 0xe4, 0xb9, 0x58, 0x5a, 0x50, 0xec, 0x36, 0x78, 0xcf,
	0xf6, 0xde, 0x42, 0xae, 0x2d, 0x4e, 0xfb, 0x4d, 0x9c, 0x4e, 0xfe, 0x33, 0x40, 0x3d, 0x66, 0xee,
	0x87, 0x95, 0x2e, 0x0d, 0x4e, 0xf7, 0x46, 0x0d, 0xc0, 0x31, 0xfc, 0xa6, 0xaa, 0x85, 0x76, 0xb4,
	0xcd, 0xfa, 0xac, 0x9b, 0x8a, 0x4c, 0xc4, 0x79, 0x30, 0x65, 0x83, 0xaa, 0x3a, 0x0e, 0xf1, 0x88,
	0x75, 0x32, 0xb8, 0xa4, 0xd2, 0x37, 0x14, 0x31, 0x13, 0x44, 0xe8, 0x09, 0x9e, 0x54, 0x53, 0x99,
	0xc2, 0xfc, 0x36, 0x73, 0x62, 0x20, 0x28, 0x07, 0x91, 0x13, 0x03, 0x42, 0x1f, 0x43, 0x0c, 0x9e,
	0x3b, 0xec, 0x20, 0xf4, 0xb8, 0x1e, 0x7f, 0xef, 0xb0, 0xde, 0xab, 0x0c, 0x40, 0x43, 0xc6, 0x7f,
	0x3a, 0xac, 0x7f, 0x2a, 0x4a, 0x3a, 0x00, 0xee, 0x35, 0x3b, 0x35, 0xdf, 0x8e, 0x7f, 0x70, 0x83,
	0x92, 0xca, 0x32, 0xf8, 0xea, 0x7c, 0xfe, 0xf5, 0xfb, 0x9b, 0xf3, 0xc5, 0xe1, 0x7b, 0x61, 0x2e,
	0xca, 0xf0, 0x1a, 0x0f, 0x61, 0x3d, 0x3f, 0xe0, 0xfb, 0xb4, 0xcd, 0x75, 0xf6, 0x49, 0xc8, 0xf0,
	0xba, 0x7a, 0x37, 0xeb, 0xb9, 0xcf, 0x3d, 0xfa, 0x4e, 0x0f, 0x82, 0x14, 0xfb, 0x34, 0xd6, 0xf3,
	0x43, 0xee, 0x93, 0x56, 0x11, 0x4f, 0xe2, 0x06, 0xfe, 0xf5, 0xdc, 0xe3, 0x07, 0xa4, 0x22, 0xd7,
	0xa4, 0x19, 0xc0, 0xeb, 0x9a, 0x44, 0x6e, 0x25, 0x19, 0x86, 0x6b, 0x8d, 0x00, 0x23, 0xcd, 0xa2,
	0x56, 0xf7, 0x2b, 0x6a, 0xb1, 0xa8, 0xd5, 0x46, 0x72, 0x3a, 0xb1, 0x49, 0xe2, 0xe5, 0xef, 0xa4,
	0x11, 0x37, 0x34, 0xad, 0xf9, 0x0b, 0xe6, 0xe2, 0x85, 0xf0, 0xb6, 0x6b, 0xf3, 0xef, 0x6e, 0x0b,
	0x78, 0x94, 0xb7, 0xe8, 0x24, 0x7b, 0x7c, 0x37, 0xd4, 0x90, 0xeb, 0x45, 0x97, 0xb0, 0x3d, 0xfe,
	0x13, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x5c, 0xd3, 0x73, 0x23, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	Test(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestReply, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/helloworld.Greeter/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) Test(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestReply, error) {
	out := new(TestReply)
	err := c.cc.Invoke(ctx, "/helloworld.Greeter/Test", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	Test(context.Context, *TestRequest) (*TestReply, error)
}

// UnimplementedGreeterServer can be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (*UnimplementedGreeterServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (*UnimplementedGreeterServer) Test(ctx context.Context, req *TestRequest) (*TestReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Test not implemented")
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_Test_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Test(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/Test",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Test(ctx, req.(*TestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
		{
			MethodName: "Test",
			Handler:    _Greeter_Test_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld.proto",
}