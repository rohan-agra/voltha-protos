// Code generated by protoc-gen-go. DO NOT EDIT.
// source: voltha_protos/afrouter.proto

package afrouter

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/opencord/voltha-protos/v3/go/common"
	grpc "google.golang.org/grpc"
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

type Result struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Info                 string   `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_be7e2f565b9e1c96, []int{0}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Result) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *Result) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_be7e2f565b9e1c96, []int{1}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Count struct {
	Count                uint32   `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Count) Reset()         { *m = Count{} }
func (m *Count) String() string { return proto.CompactTextString(m) }
func (*Count) ProtoMessage()    {}
func (*Count) Descriptor() ([]byte, []int) {
	return fileDescriptor_be7e2f565b9e1c96, []int{2}
}

func (m *Count) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Count.Unmarshal(m, b)
}
func (m *Count) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Count.Marshal(b, m, deterministic)
}
func (m *Count) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Count.Merge(m, src)
}
func (m *Count) XXX_Size() int {
	return xxx_messageInfo_Count.Size(m)
}
func (m *Count) XXX_DiscardUnknown() {
	xxx_messageInfo_Count.DiscardUnknown(m)
}

var xxx_messageInfo_Count proto.InternalMessageInfo

func (m *Count) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Conn struct {
	Server               string   `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Pkg                  string   `protobuf:"bytes,2,opt,name=pkg,proto3" json:"pkg,omitempty"`
	Svc                  string   `protobuf:"bytes,3,opt,name=svc,proto3" json:"svc,omitempty"`
	Cluster              string   `protobuf:"bytes,4,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Backend              string   `protobuf:"bytes,5,opt,name=backend,proto3" json:"backend,omitempty"`
	Connection           string   `protobuf:"bytes,6,opt,name=connection,proto3" json:"connection,omitempty"`
	Addr                 string   `protobuf:"bytes,7,opt,name=addr,proto3" json:"addr,omitempty"`
	Port                 uint64   `protobuf:"varint,8,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Conn) Reset()         { *m = Conn{} }
func (m *Conn) String() string { return proto.CompactTextString(m) }
func (*Conn) ProtoMessage()    {}
func (*Conn) Descriptor() ([]byte, []int) {
	return fileDescriptor_be7e2f565b9e1c96, []int{3}
}

func (m *Conn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Conn.Unmarshal(m, b)
}
func (m *Conn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Conn.Marshal(b, m, deterministic)
}
func (m *Conn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Conn.Merge(m, src)
}
func (m *Conn) XXX_Size() int {
	return xxx_messageInfo_Conn.Size(m)
}
func (m *Conn) XXX_DiscardUnknown() {
	xxx_messageInfo_Conn.DiscardUnknown(m)
}

var xxx_messageInfo_Conn proto.InternalMessageInfo

func (m *Conn) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

func (m *Conn) GetPkg() string {
	if m != nil {
		return m.Pkg
	}
	return ""
}

func (m *Conn) GetSvc() string {
	if m != nil {
		return m.Svc
	}
	return ""
}

func (m *Conn) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *Conn) GetBackend() string {
	if m != nil {
		return m.Backend
	}
	return ""
}

func (m *Conn) GetConnection() string {
	if m != nil {
		return m.Connection
	}
	return ""
}

func (m *Conn) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *Conn) GetPort() uint64 {
	if m != nil {
		return m.Port
	}
	return 0
}

type Affinity struct {
	Router               string   `protobuf:"bytes,1,opt,name=router,proto3" json:"router,omitempty"`
	Route                string   `protobuf:"bytes,2,opt,name=route,proto3" json:"route,omitempty"`
	Cluster              string   `protobuf:"bytes,3,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Backend              string   `protobuf:"bytes,4,opt,name=backend,proto3" json:"backend,omitempty"`
	Id                   string   `protobuf:"bytes,5,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Affinity) Reset()         { *m = Affinity{} }
func (m *Affinity) String() string { return proto.CompactTextString(m) }
func (*Affinity) ProtoMessage()    {}
func (*Affinity) Descriptor() ([]byte, []int) {
	return fileDescriptor_be7e2f565b9e1c96, []int{4}
}

func (m *Affinity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Affinity.Unmarshal(m, b)
}
func (m *Affinity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Affinity.Marshal(b, m, deterministic)
}
func (m *Affinity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Affinity.Merge(m, src)
}
func (m *Affinity) XXX_Size() int {
	return xxx_messageInfo_Affinity.Size(m)
}
func (m *Affinity) XXX_DiscardUnknown() {
	xxx_messageInfo_Affinity.DiscardUnknown(m)
}

var xxx_messageInfo_Affinity proto.InternalMessageInfo

func (m *Affinity) GetRouter() string {
	if m != nil {
		return m.Router
	}
	return ""
}

func (m *Affinity) GetRoute() string {
	if m != nil {
		return m.Route
	}
	return ""
}

func (m *Affinity) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *Affinity) GetBackend() string {
	if m != nil {
		return m.Backend
	}
	return ""
}

func (m *Affinity) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Result)(nil), "afrouter.Result")
	proto.RegisterType((*Empty)(nil), "afrouter.Empty")
	proto.RegisterType((*Count)(nil), "afrouter.Count")
	proto.RegisterType((*Conn)(nil), "afrouter.Conn")
	proto.RegisterType((*Affinity)(nil), "afrouter.Affinity")
}

func init() { proto.RegisterFile("voltha_protos/afrouter.proto", fileDescriptor_be7e2f565b9e1c96) }

var fileDescriptor_be7e2f565b9e1c96 = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0xdd, 0xb4, 0xe9, 0xc7, 0x0e, 0xea, 0xee, 0x62, 0x10, 0xb2, 0x2a, 0x40, 0x55, 0x4e, 0xbd,
	0x90, 0x22, 0x2a, 0xc4, 0x19, 0x2a, 0xd4, 0x0b, 0xa7, 0xec, 0x8d, 0x0b, 0x4a, 0x1d, 0x27, 0x6b,
	0x6d, 0xe3, 0x89, 0x6c, 0x27, 0xd2, 0x4a, 0xfc, 0x2c, 0xce, 0xfc, 0x36, 0x64, 0x3b, 0x6e, 0x5a,
	0x89, 0xbd, 0xbd, 0x8f, 0x49, 0xfc, 0x9e, 0x3d, 0xf0, 0xb6, 0xc3, 0xa3, 0x79, 0xc8, 0x7f, 0x35,
	0x0a, 0x0d, 0xea, 0x4d, 0x5e, 0x2a, 0x6c, 0x0d, 0x57, 0xa9, 0xe3, 0x64, 0x1e, 0xf8, 0x72, 0x79,
	0x39, 0xc7, 0xb0, 0xae, 0x51, 0xfa, 0xa9, 0xe4, 0x07, 0x4c, 0x33, 0xae, 0xdb, 0xa3, 0x21, 0x14,
	0x66, 0xba, 0x65, 0x8c, 0x6b, 0x4d, 0xa3, 0x55, 0xb4, 0x9e, 0x67, 0x81, 0x92, 0xd7, 0x30, 0xe1,
	0x4a, 0xa1, 0xa2, 0xa3, 0x55, 0xb4, 0xbe, 0xce, 0x3c, 0x21, 0x04, 0x62, 0x21, 0x4b, 0xa4, 0x63,
	0x27, 0x3a, 0x9c, 0xcc, 0x60, 0xf2, 0xbd, 0x6e, 0xcc, 0x53, 0xf2, 0x0e, 0x26, 0x3b, 0x6c, 0xa5,
	0xb1, 0xdf, 0x32, 0x0b, 0xdc, 0x3f, 0x17, 0x99, 0x27, 0xc9, 0xdf, 0x08, 0xe2, 0x1d, 0x4a, 0x49,
	0xde, 0xc0, 0x54, 0x73, 0xd5, 0x71, 0xe5, 0xfc, 0xeb, 0xac, 0x67, 0xe4, 0x0e, 0xc6, 0xcd, 0x63,
	0xd5, 0x1f, 0x68, 0xa1, 0x55, 0x74, 0xc7, 0xfa, 0xd3, 0x2c, 0xb4, 0x81, 0xd9, 0xb1, 0xd5, 0x86,
	0x2b, 0x1a, 0x3b, 0x35, 0x50, 0xeb, 0x1c, 0x72, 0xf6, 0xc8, 0x65, 0x41, 0x27, 0xde, 0xe9, 0x29,
	0x79, 0x0f, 0xc0, 0x50, 0x4a, 0xce, 0x8c, 0x40, 0x49, 0xa7, 0xce, 0x3c, 0x53, 0x6c, 0xa9, 0xbc,
	0x28, 0x14, 0x9d, 0xf9, 0x52, 0x16, 0x5b, 0xad, 0x41, 0x65, 0xe8, 0x7c, 0x15, 0xad, 0xe3, 0xcc,
	0xe1, 0xe4, 0x37, 0xcc, 0xbf, 0x96, 0xa5, 0x90, 0xc2, 0x3c, 0xd9, 0x0e, 0xfe, 0xa2, 0x43, 0x07,
	0xcf, 0x6c, 0x75, 0x87, 0xc2, 0xb5, 0x39, 0x72, 0x9e, 0x7a, 0xfc, 0x6c, 0xea, 0xf8, 0x32, 0xf5,
	0x0d, 0x8c, 0x44, 0xa8, 0x32, 0x12, 0xc5, 0xa7, 0x3f, 0x11, 0x2c, 0x76, 0x28, 0x4b, 0x51, 0xb5,
	0x2a, 0x77, 0xb9, 0xb7, 0xb0, 0xb8, 0xe7, 0x66, 0x37, 0x14, 0xb9, 0x49, 0x4f, 0xeb, 0x60, 0xd5,
	0xe5, 0xdd, 0xc0, 0xfd, 0x7b, 0x27, 0x57, 0xe4, 0x33, 0xbc, 0xb8, 0xe7, 0xe6, 0xd4, 0x83, 0x0c,
	0x23, 0x41, 0xfb, 0xef, 0x67, 0x5f, 0xe0, 0xe5, 0x9e, 0x9b, 0x3d, 0x5a, 0x5d, 0x48, 0xee, 0xdf,
	0xf9, 0x76, 0x18, 0x74, 0x1b, 0xb0, 0xbc, 0x3d, 0x0f, 0x60, 0xdf, 0xfc, 0xea, 0xdb, 0x1e, 0x5e,
	0xa1, 0xaa, 0x52, 0x6c, 0xb8, 0x64, 0xa8, 0x8a, 0xd4, 0xaf, 0xe5, 0xcf, 0x8f, 0x95, 0x30, 0x0f,
	0xed, 0x21, 0x65, 0x58, 0x6f, 0x82, 0xb7, 0xf1, 0xde, 0x87, 0x7e, 0x65, 0xbb, 0xed, 0xa6, 0xc2,
	0xd3, 0x82, 0x1f, 0xa6, 0x4e, 0xde, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x5d, 0x28, 0xe9, 0x22,
	0x01, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ConfigurationClient is the client API for Configuration service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConfigurationClient interface {
	SetConnection(ctx context.Context, in *Conn, opts ...grpc.CallOption) (*Result, error)
	SetAffinity(ctx context.Context, in *Affinity, opts ...grpc.CallOption) (*Result, error)
	GetGoroutineCount(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Count, error)
}

type configurationClient struct {
	cc *grpc.ClientConn
}

func NewConfigurationClient(cc *grpc.ClientConn) ConfigurationClient {
	return &configurationClient{cc}
}

func (c *configurationClient) SetConnection(ctx context.Context, in *Conn, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/afrouter.Configuration/SetConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configurationClient) SetAffinity(ctx context.Context, in *Affinity, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/afrouter.Configuration/SetAffinity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configurationClient) GetGoroutineCount(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Count, error) {
	out := new(Count)
	err := c.cc.Invoke(ctx, "/afrouter.Configuration/GetGoroutineCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigurationServer is the server API for Configuration service.
type ConfigurationServer interface {
	SetConnection(context.Context, *Conn) (*Result, error)
	SetAffinity(context.Context, *Affinity) (*Result, error)
	GetGoroutineCount(context.Context, *Empty) (*Count, error)
}

func RegisterConfigurationServer(s *grpc.Server, srv ConfigurationServer) {
	s.RegisterService(&_Configuration_serviceDesc, srv)
}

func _Configuration_SetConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Conn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationServer).SetConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/afrouter.Configuration/SetConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationServer).SetConnection(ctx, req.(*Conn))
	}
	return interceptor(ctx, in, info, handler)
}

func _Configuration_SetAffinity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Affinity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationServer).SetAffinity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/afrouter.Configuration/SetAffinity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationServer).SetAffinity(ctx, req.(*Affinity))
	}
	return interceptor(ctx, in, info, handler)
}

func _Configuration_GetGoroutineCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationServer).GetGoroutineCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/afrouter.Configuration/GetGoroutineCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationServer).GetGoroutineCount(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Configuration_serviceDesc = grpc.ServiceDesc{
	ServiceName: "afrouter.Configuration",
	HandlerType: (*ConfigurationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetConnection",
			Handler:    _Configuration_SetConnection_Handler,
		},
		{
			MethodName: "SetAffinity",
			Handler:    _Configuration_SetAffinity_Handler,
		},
		{
			MethodName: "GetGoroutineCount",
			Handler:    _Configuration_GetGoroutineCount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "voltha_protos/afrouter.proto",
}
