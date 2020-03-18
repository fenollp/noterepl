// Code generated by protoc-gen-go. DO NOT EDIT.
// source: noterepl.proto

package pkg

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Value_Kind int32

const (
	Value_EMPTY      Value_Kind = 0
	Value_is_null    Value_Kind = 1
	Value_is_str     Value_Kind = 2
	Value_is_number  Value_Kind = 3
	Value_is_list    Value_Kind = 4
	Value_is_hashmap Value_Kind = 5
	Value_is_url     Value_Kind = 6
	Value_is_bool    Value_Kind = 7
)

var Value_Kind_name = map[int32]string{
	0: "EMPTY",
	1: "is_null",
	2: "is_str",
	3: "is_number",
	4: "is_list",
	5: "is_hashmap",
	6: "is_url",
	7: "is_bool",
}

var Value_Kind_value = map[string]int32{
	"EMPTY":      0,
	"is_null":    1,
	"is_str":     2,
	"is_number":  3,
	"is_list":    4,
	"is_hashmap": 5,
	"is_url":     6,
	"is_bool":    7,
}

func (x Value_Kind) String() string {
	return proto.EnumName(Value_Kind_name, int32(x))
}

func (Value_Kind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c16ddaca6537879b, []int{1, 0}
}

type Object struct {
	Value                *Value   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Ptr                  uint64   `protobuf:"varint,2,opt,name=ptr,proto3" json:"ptr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Object) Reset()         { *m = Object{} }
func (m *Object) String() string { return proto.CompactTextString(m) }
func (*Object) ProtoMessage()    {}
func (*Object) Descriptor() ([]byte, []int) {
	return fileDescriptor_c16ddaca6537879b, []int{0}
}

func (m *Object) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Object.Unmarshal(m, b)
}
func (m *Object) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Object.Marshal(b, m, deterministic)
}
func (m *Object) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Object.Merge(m, src)
}
func (m *Object) XXX_Size() int {
	return xxx_messageInfo_Object.Size(m)
}
func (m *Object) XXX_DiscardUnknown() {
	xxx_messageInfo_Object.DiscardUnknown(m)
}

var xxx_messageInfo_Object proto.InternalMessageInfo

func (m *Object) GetValue() *Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Object) GetPtr() uint64 {
	if m != nil {
		return m.Ptr
	}
	return 0
}

// Only one of these fields are ever set at once
// Not using `oneof` here as types gets less weird this way.
type Value struct {
	Kind                 Value_Kind         `protobuf:"varint,1,opt,name=kind,proto3,enum=Value_Kind" json:"kind,omitempty"`
	Str                  string             `protobuf:"bytes,2,opt,name=str,proto3" json:"str,omitempty"`
	Number               float64            `protobuf:"fixed64,3,opt,name=number,proto3" json:"number,omitempty"`
	List                 []*Object          `protobuf:"bytes,4,rep,name=list,proto3" json:"list,omitempty"`
	Hashmap              map[string]*Object `protobuf:"bytes,5,rep,name=hashmap,proto3" json:"hashmap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Url                  string             `protobuf:"bytes,6,opt,name=url,proto3" json:"url,omitempty"`
	Boolean              bool               `protobuf:"varint,7,opt,name=boolean,proto3" json:"boolean,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Value) Reset()         { *m = Value{} }
func (m *Value) String() string { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()    {}
func (*Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_c16ddaca6537879b, []int{1}
}

func (m *Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Value.Unmarshal(m, b)
}
func (m *Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Value.Marshal(b, m, deterministic)
}
func (m *Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Value.Merge(m, src)
}
func (m *Value) XXX_Size() int {
	return xxx_messageInfo_Value.Size(m)
}
func (m *Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Value.DiscardUnknown(m)
}

var xxx_messageInfo_Value proto.InternalMessageInfo

func (m *Value) GetKind() Value_Kind {
	if m != nil {
		return m.Kind
	}
	return Value_EMPTY
}

func (m *Value) GetStr() string {
	if m != nil {
		return m.Str
	}
	return ""
}

func (m *Value) GetNumber() float64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *Value) GetList() []*Object {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *Value) GetHashmap() map[string]*Object {
	if m != nil {
		return m.Hashmap
	}
	return nil
}

func (m *Value) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Value) GetBoolean() bool {
	if m != nil {
		return m.Boolean
	}
	return false
}

type EvalReq struct {
	Language             string   `protobuf:"bytes,1,opt,name=language,proto3" json:"language,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EvalReq) Reset()         { *m = EvalReq{} }
func (m *EvalReq) String() string { return proto.CompactTextString(m) }
func (*EvalReq) ProtoMessage()    {}
func (*EvalReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c16ddaca6537879b, []int{2}
}

func (m *EvalReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvalReq.Unmarshal(m, b)
}
func (m *EvalReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvalReq.Marshal(b, m, deterministic)
}
func (m *EvalReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvalReq.Merge(m, src)
}
func (m *EvalReq) XXX_Size() int {
	return xxx_messageInfo_EvalReq.Size(m)
}
func (m *EvalReq) XXX_DiscardUnknown() {
	xxx_messageInfo_EvalReq.DiscardUnknown(m)
}

var xxx_messageInfo_EvalReq proto.InternalMessageInfo

func (m *EvalReq) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *EvalReq) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type EvalRep struct {
	Stderr               string   `protobuf:"bytes,1,opt,name=stderr,proto3" json:"stderr,omitempty"`
	Result               *Object  `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EvalRep) Reset()         { *m = EvalRep{} }
func (m *EvalRep) String() string { return proto.CompactTextString(m) }
func (*EvalRep) ProtoMessage()    {}
func (*EvalRep) Descriptor() ([]byte, []int) {
	return fileDescriptor_c16ddaca6537879b, []int{3}
}

func (m *EvalRep) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvalRep.Unmarshal(m, b)
}
func (m *EvalRep) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvalRep.Marshal(b, m, deterministic)
}
func (m *EvalRep) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvalRep.Merge(m, src)
}
func (m *EvalRep) XXX_Size() int {
	return xxx_messageInfo_EvalRep.Size(m)
}
func (m *EvalRep) XXX_DiscardUnknown() {
	xxx_messageInfo_EvalRep.DiscardUnknown(m)
}

var xxx_messageInfo_EvalRep proto.InternalMessageInfo

func (m *EvalRep) GetStderr() string {
	if m != nil {
		return m.Stderr
	}
	return ""
}

func (m *EvalRep) GetResult() *Object {
	if m != nil {
		return m.Result
	}
	return nil
}

type PrintReq struct {
	// uint64 is ptr
	Path                 []uint64 `protobuf:"varint,1,rep,packed,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrintReq) Reset()         { *m = PrintReq{} }
func (m *PrintReq) String() string { return proto.CompactTextString(m) }
func (*PrintReq) ProtoMessage()    {}
func (*PrintReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c16ddaca6537879b, []int{4}
}

func (m *PrintReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrintReq.Unmarshal(m, b)
}
func (m *PrintReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrintReq.Marshal(b, m, deterministic)
}
func (m *PrintReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrintReq.Merge(m, src)
}
func (m *PrintReq) XXX_Size() int {
	return xxx_messageInfo_PrintReq.Size(m)
}
func (m *PrintReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PrintReq.DiscardUnknown(m)
}

var xxx_messageInfo_PrintReq proto.InternalMessageInfo

func (m *PrintReq) GetPath() []uint64 {
	if m != nil {
		return m.Path
	}
	return nil
}

type PrintRep struct {
	Value                *Value   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrintRep) Reset()         { *m = PrintRep{} }
func (m *PrintRep) String() string { return proto.CompactTextString(m) }
func (*PrintRep) ProtoMessage()    {}
func (*PrintRep) Descriptor() ([]byte, []int) {
	return fileDescriptor_c16ddaca6537879b, []int{5}
}

func (m *PrintRep) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrintRep.Unmarshal(m, b)
}
func (m *PrintRep) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrintRep.Marshal(b, m, deterministic)
}
func (m *PrintRep) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrintRep.Merge(m, src)
}
func (m *PrintRep) XXX_Size() int {
	return xxx_messageInfo_PrintRep.Size(m)
}
func (m *PrintRep) XXX_DiscardUnknown() {
	xxx_messageInfo_PrintRep.DiscardUnknown(m)
}

var xxx_messageInfo_PrintRep proto.InternalMessageInfo

func (m *PrintRep) GetValue() *Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterEnum("Value_Kind", Value_Kind_name, Value_Kind_value)
	proto.RegisterType((*Object)(nil), "Object")
	proto.RegisterType((*Value)(nil), "Value")
	proto.RegisterMapType((map[string]*Object)(nil), "Value.HashmapEntry")
	proto.RegisterType((*EvalReq)(nil), "EvalReq")
	proto.RegisterType((*EvalRep)(nil), "EvalRep")
	proto.RegisterType((*PrintReq)(nil), "PrintReq")
	proto.RegisterType((*PrintRep)(nil), "PrintRep")
}

func init() { proto.RegisterFile("noterepl.proto", fileDescriptor_c16ddaca6537879b) }

var fileDescriptor_c16ddaca6537879b = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x4d, 0x8f, 0xd3, 0x30,
	0x14, 0x6c, 0x1a, 0x27, 0x69, 0x5e, 0xa1, 0x8a, 0x1e, 0x12, 0x8a, 0xca, 0x42, 0xa3, 0x9c, 0x72,
	0x21, 0x87, 0x72, 0x59, 0x38, 0x2e, 0xaa, 0x04, 0xe2, 0xab, 0xb2, 0x10, 0x12, 0x5c, 0xaa, 0x74,
	0x6b, 0xb5, 0xa1, 0x5e, 0xc7, 0xd8, 0xce, 0x4a, 0x7b, 0xe3, 0xa7, 0x23, 0x3b, 0x4e, 0x85, 0x38,
	0x70, 0x9b, 0xf7, 0x31, 0x99, 0x99, 0x17, 0xc3, 0x42, 0x74, 0x86, 0x29, 0x26, 0x79, 0x2d, 0x55,
	0x67, 0xba, 0xf2, 0x1a, 0xe2, 0x2f, 0xfb, 0x9f, 0xec, 0xd6, 0xe0, 0x15, 0x44, 0xf7, 0x0d, 0xef,
	0x59, 0x1e, 0x14, 0x41, 0x35, 0x5f, 0xc7, 0xf5, 0x37, 0x5b, 0xd1, 0xa1, 0x89, 0x19, 0x84, 0xd2,
	0xa8, 0x7c, 0x5a, 0x04, 0x15, 0xa1, 0x16, 0x96, 0xbf, 0x43, 0x88, 0xdc, 0x0a, 0xae, 0x80, 0x9c,
	0x5b, 0x71, 0x70, 0xc4, 0xc5, 0x7a, 0x3e, 0x10, 0xeb, 0x0f, 0xad, 0x38, 0x50, 0x37, 0xb0, 0x64,
	0xed, 0xc9, 0x29, 0xb5, 0x10, 0x9f, 0x42, 0x2c, 0xfa, 0xbb, 0x3d, 0x53, 0x79, 0x58, 0x04, 0x55,
	0x40, 0x7d, 0x85, 0xcf, 0x80, 0xf0, 0x56, 0x9b, 0x9c, 0x14, 0x61, 0x35, 0x5f, 0x27, 0xf5, 0xe0,
	0x8d, 0xba, 0x26, 0xbe, 0x84, 0xe4, 0xd4, 0xe8, 0xd3, 0x5d, 0x23, 0xf3, 0xc8, 0xcd, 0x9f, 0x78,
	0xa9, 0x77, 0x43, 0x77, 0x23, 0x8c, 0x7a, 0xa0, 0xe3, 0x8e, 0x55, 0xed, 0x15, 0xcf, 0xe3, 0x41,
	0xb5, 0x57, 0x1c, 0x73, 0x48, 0xf6, 0x5d, 0xc7, 0x59, 0x23, 0xf2, 0xa4, 0x08, 0xaa, 0x19, 0x1d,
	0xcb, 0xe5, 0x5b, 0x78, 0xf4, 0xf7, 0x47, 0x2c, 0xf7, 0xcc, 0x1e, 0x5c, 0xa2, 0x94, 0x5a, 0x88,
	0xcf, 0xc7, 0xf3, 0x4c, 0xdd, 0x79, 0x2e, 0xd6, 0x86, 0xee, 0x9b, 0xe9, 0x75, 0x50, 0x76, 0x40,
	0x6c, 0x68, 0x4c, 0x21, 0xda, 0x7c, 0xda, 0x7e, 0xfd, 0x9e, 0x4d, 0x70, 0x0e, 0x49, 0xab, 0x77,
	0xa2, 0xe7, 0x3c, 0x0b, 0x10, 0x20, 0x6e, 0xf5, 0x4e, 0x1b, 0x95, 0x4d, 0xf1, 0x31, 0xa4, 0x6e,
	0x60, 0x53, 0x67, 0xa1, 0xdf, 0xb3, 0x29, 0x33, 0x82, 0x0b, 0x80, 0x56, 0xef, 0x7c, 0x8c, 0x2c,
	0xf2, 0xbc, 0x5e, 0xf1, 0x2c, 0xf6, 0x8b, 0xd6, 0x76, 0x96, 0x94, 0xaf, 0x21, 0xd9, 0xdc, 0x37,
	0x9c, 0xb2, 0x5f, 0xb8, 0x84, 0x19, 0x6f, 0xc4, 0xb1, 0x6f, 0x8e, 0xcc, 0xbb, 0xbe, 0xd4, 0x88,
	0x40, 0x6e, 0xbb, 0x03, 0xf3, 0xf7, 0x77, 0xb8, 0xbc, 0x19, 0xa9, 0xd2, 0xfe, 0x0b, 0x6d, 0x0e,
	0x4c, 0x29, 0x4f, 0xf4, 0x15, 0xae, 0x20, 0x56, 0x4c, 0xf7, 0xdc, 0xfc, 0x1b, 0xd9, 0xb7, 0xcb,
	0x17, 0x30, 0xdb, 0xaa, 0x56, 0x18, 0xab, 0x8f, 0x40, 0x64, 0x63, 0x4e, 0x79, 0x50, 0x84, 0x15,
	0xa1, 0x0e, 0x97, 0xd5, 0x65, 0x2e, 0xff, 0xff, 0xba, 0xd6, 0xef, 0x61, 0xf6, 0xb9, 0x33, 0x8c,
	0x6e, 0xb6, 0x1f, 0xf1, 0x0a, 0x88, 0x75, 0x86, 0xb3, 0xda, 0x67, 0x5b, 0x8e, 0x48, 0x96, 0x13,
	0x5c, 0x41, 0xe4, 0xbe, 0x89, 0x69, 0x3d, 0x6a, 0x2f, 0x2f, 0x50, 0x96, 0x93, 0x9b, 0xe8, 0x47,
	0x28, 0xcf, 0xc7, 0x7d, 0xec, 0x9e, 0xf7, 0xab, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x39, 0x02,
	0x52, 0xad, 0xf0, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NoteREPLClient is the client API for NoteREPL service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NoteREPLClient interface {
	Eval(ctx context.Context, in *EvalReq, opts ...grpc.CallOption) (*EvalRep, error)
	Print(ctx context.Context, in *PrintReq, opts ...grpc.CallOption) (*PrintRep, error)
}

type noteREPLClient struct {
	cc *grpc.ClientConn
}

func NewNoteREPLClient(cc *grpc.ClientConn) NoteREPLClient {
	return &noteREPLClient{cc}
}

func (c *noteREPLClient) Eval(ctx context.Context, in *EvalReq, opts ...grpc.CallOption) (*EvalRep, error) {
	out := new(EvalRep)
	err := c.cc.Invoke(ctx, "/NoteREPL/Eval", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noteREPLClient) Print(ctx context.Context, in *PrintReq, opts ...grpc.CallOption) (*PrintRep, error) {
	out := new(PrintRep)
	err := c.cc.Invoke(ctx, "/NoteREPL/Print", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NoteREPLServer is the server API for NoteREPL service.
type NoteREPLServer interface {
	Eval(context.Context, *EvalReq) (*EvalRep, error)
	Print(context.Context, *PrintReq) (*PrintRep, error)
}

func RegisterNoteREPLServer(s *grpc.Server, srv NoteREPLServer) {
	s.RegisterService(&_NoteREPL_serviceDesc, srv)
}

func _NoteREPL_Eval_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EvalReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteREPLServer).Eval(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NoteREPL/Eval",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteREPLServer).Eval(ctx, req.(*EvalReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NoteREPL_Print_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrintReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteREPLServer).Print(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NoteREPL/Print",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteREPLServer).Print(ctx, req.(*PrintReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _NoteREPL_serviceDesc = grpc.ServiceDesc{
	ServiceName: "NoteREPL",
	HandlerType: (*NoteREPLServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Eval",
			Handler:    _NoteREPL_Eval_Handler,
		},
		{
			MethodName: "Print",
			Handler:    _NoteREPL_Print_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "noterepl.proto",
}
