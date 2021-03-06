// Code generated by protoc-gen-go. DO NOT EDIT.
// source: errors.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Error_Code_Type int32

const (
	Error_Code_Undefined Error_Code_Type = 0
	Error_Code_Service   Error_Code_Type = 1
	Error_Code_Network   Error_Code_Type = 2
	Error_Code_DateTime  Error_Code_Type = 3
	Error_Code_Micro     Error_Code_Type = 4
	Error_Code_Library   Error_Code_Type = 5
	Error_Code_Wrapper   Error_Code_Type = 6
)

var Error_Code_Type_name = map[int32]string{
	0: "Undefined",
	1: "Service",
	2: "Network",
	3: "DateTime",
	4: "Micro",
	5: "Library",
	6: "Wrapper",
}

var Error_Code_Type_value = map[string]int32{
	"Undefined": 0,
	"Service":   1,
	"Network":   2,
	"DateTime":  3,
	"Micro":     4,
	"Library":   5,
	"Wrapper":   6,
}

func (x Error_Code_Type) String() string {
	return proto.EnumName(Error_Code_Type_name, int32(x))
}

func (Error_Code_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_24fe73c7f0ddb19c, []int{0, 0, 0}
}

type Error struct {
	Code                 *Error_Code `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string      `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Error                *Error      `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_24fe73c7f0ddb19c, []int{0}
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

func (m *Error) GetCode() *Error_Code {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Error) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type Error_Code struct {
	Service              uint32          `protobuf:"varint,1,opt,name=service,proto3" json:"service,omitempty"`
	Type                 Error_Code_Type `protobuf:"varint,2,opt,name=type,proto3,enum=proto.Error_Code_Type" json:"type,omitempty"`
	Number               uint32          `protobuf:"varint,3,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Error_Code) Reset()         { *m = Error_Code{} }
func (m *Error_Code) String() string { return proto.CompactTextString(m) }
func (*Error_Code) ProtoMessage()    {}
func (*Error_Code) Descriptor() ([]byte, []int) {
	return fileDescriptor_24fe73c7f0ddb19c, []int{0, 0}
}

func (m *Error_Code) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error_Code.Unmarshal(m, b)
}
func (m *Error_Code) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error_Code.Marshal(b, m, deterministic)
}
func (m *Error_Code) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error_Code.Merge(m, src)
}
func (m *Error_Code) XXX_Size() int {
	return xxx_messageInfo_Error_Code.Size(m)
}
func (m *Error_Code) XXX_DiscardUnknown() {
	xxx_messageInfo_Error_Code.DiscardUnknown(m)
}

var xxx_messageInfo_Error_Code proto.InternalMessageInfo

func (m *Error_Code) GetService() uint32 {
	if m != nil {
		return m.Service
	}
	return 0
}

func (m *Error_Code) GetType() Error_Code_Type {
	if m != nil {
		return m.Type
	}
	return Error_Code_Undefined
}

func (m *Error_Code) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func init() {
	proto.RegisterEnum("proto.Error_Code_Type", Error_Code_Type_name, Error_Code_Type_value)
	proto.RegisterType((*Error)(nil), "proto.Error")
	proto.RegisterType((*Error_Code)(nil), "proto.Error.Code")
}

func init() { proto.RegisterFile("errors.proto", fileDescriptor_24fe73c7f0ddb19c) }

var fileDescriptor_24fe73c7f0ddb19c = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8e, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0xbf, 0xa4, 0x49, 0xfa, 0xe5, 0x36, 0x91, 0xf1, 0x2e, 0x4a, 0x70, 0x55, 0x02, 0x42,
	0x71, 0x91, 0x45, 0x7d, 0x84, 0xea, 0x4e, 0x5d, 0x8c, 0x15, 0xd7, 0xf9, 0x73, 0x95, 0x41, 0x92,
	0x19, 0x6e, 0xa2, 0x92, 0x27, 0xf3, 0x15, 0x7c, 0x2c, 0x99, 0x99, 0x16, 0x04, 0x57, 0xc3, 0x6f,
	0xce, 0xe1, 0xfc, 0x2e, 0x64, 0xc4, 0xac, 0x79, 0xac, 0x0c, 0xeb, 0x49, 0x63, 0xec, 0x9e, 0xf2,
	0x2b, 0x84, 0xf8, 0xd6, 0xfe, 0xe3, 0x25, 0x44, 0xad, 0xee, 0xa8, 0x08, 0x36, 0xc1, 0x76, 0xb5,
	0x3b, 0xf7, 0xb5, 0xca, 0x65, 0xd5, 0x5e, 0x77, 0x24, 0x5d, 0x8c, 0x05, 0x2c, 0x7b, 0x1a, 0xc7,
	0xfa, 0x95, 0x8a, 0x70, 0x13, 0x6c, 0x53, 0x79, 0x42, 0x2c, 0x21, 0x76, 0x86, 0x62, 0xe1, 0x16,
	0xb2, 0xdf, 0x0b, 0xd2, 0x47, 0x17, 0xdf, 0x01, 0x44, 0xfb, 0xe3, 0xcc, 0x48, 0xfc, 0xa1, 0x5a,
	0x2f, 0xcc, 0xe5, 0x09, 0xf1, 0x0a, 0xa2, 0x69, 0x36, 0x7e, 0xfd, 0x6c, 0xb7, 0xfe, 0x73, 0x47,
	0x75, 0x98, 0x0d, 0x49, 0xd7, 0xc1, 0x35, 0x24, 0xc3, 0x7b, 0xdf, 0x90, 0x77, 0xe6, 0xf2, 0x48,
	0x65, 0x03, 0x91, 0x6d, 0x61, 0x0e, 0xe9, 0xd3, 0xd0, 0xd1, 0x8b, 0x1a, 0xa8, 0x13, 0xff, 0x70,
	0x05, 0xcb, 0x47, 0x6f, 0x11, 0x81, 0x85, 0x07, 0x9a, 0x3e, 0x35, 0xbf, 0x89, 0x10, 0x33, 0xf8,
	0x7f, 0x53, 0x4f, 0x74, 0x50, 0x3d, 0x89, 0x05, 0xa6, 0x10, 0xdf, 0xab, 0x96, 0xb5, 0x88, 0x6c,
	0xeb, 0x4e, 0x35, 0x5c, 0xf3, 0x2c, 0x62, 0x0b, 0xcf, 0x5c, 0x1b, 0x43, 0x2c, 0x92, 0x26, 0x71,
	0x87, 0x5d, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x19, 0xd2, 0xb2, 0x4b, 0x57, 0x01, 0x00, 0x00,
}
