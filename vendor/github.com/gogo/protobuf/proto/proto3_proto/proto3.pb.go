// Code generated by protoc-gen-gogo.
// source: proto3_proto/proto3.proto
// DO NOT EDIT!

/*
Package proto3_proto is a generated protocol buffer package.

It is generated from these files:
	proto3_proto/proto3.proto

It has these top-level messages:
	Message
	Nested
	MessageWithMap
*/
package proto3_proto

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import testdata "github.com/gogo/protobuf/proto/testdata"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

type Message_Humour int32

const (
	Message_UNKNOWN     Message_Humour = 0
	Message_PUNS        Message_Humour = 1
	Message_SLAPSTICK   Message_Humour = 2
	Message_BILL_BAILEY Message_Humour = 3
)

var Message_Humour_name = map[int32]string{
	0: "UNKNOWN",
	1: "PUNS",
	2: "SLAPSTICK",
	3: "BILL_BAILEY",
}
var Message_Humour_value = map[string]int32{
	"UNKNOWN":     0,
	"PUNS":        1,
	"SLAPSTICK":   2,
	"BILL_BAILEY": 3,
}

func (x Message_Humour) String() string {
	return proto.EnumName(Message_Humour_name, int32(x))
}
func (Message_Humour) EnumDescriptor() ([]byte, []int) { return fileDescriptorProto3, []int{0, 0} }

type Message struct {
	Name         string                           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Hilarity     Message_Humour                   `protobuf:"varint,2,opt,name=hilarity,proto3,enum=proto3_proto.Message_Humour" json:"hilarity,omitempty"`
	HeightInCm   uint32                           `protobuf:"varint,3,opt,name=height_in_cm,json=heightInCm,proto3" json:"height_in_cm,omitempty"`
	Data         []byte                           `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	ResultCount  int64                            `protobuf:"varint,7,opt,name=result_count,json=resultCount,proto3" json:"result_count,omitempty"`
	TrueScotsman bool                             `protobuf:"varint,8,opt,name=true_scotsman,json=trueScotsman,proto3" json:"true_scotsman,omitempty"`
	Score        float32                          `protobuf:"fixed32,9,opt,name=score,proto3" json:"score,omitempty"`
	Key          []uint64                         `protobuf:"varint,5,rep,name=key" json:"key,omitempty"`
	Nested       *Nested                          `protobuf:"bytes,6,opt,name=nested" json:"nested,omitempty"`
	RFunny       []Message_Humour                 `protobuf:"varint,16,rep,name=r_funny,json=rFunny,enum=proto3_proto.Message_Humour" json:"r_funny,omitempty"`
	Terrain      map[string]*Nested               `protobuf:"bytes,10,rep,name=terrain" json:"terrain,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
	Proto2Field  *testdata.SubDefaults            `protobuf:"bytes,11,opt,name=proto2_field,json=proto2Field" json:"proto2_field,omitempty"`
	Proto2Value  map[string]*testdata.SubDefaults `protobuf:"bytes,13,rep,name=proto2_value,json=proto2Value" json:"proto2_value,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptorProto3, []int{0} }

func (m *Message) GetNested() *Nested {
	if m != nil {
		return m.Nested
	}
	return nil
}

func (m *Message) GetTerrain() map[string]*Nested {
	if m != nil {
		return m.Terrain
	}
	return nil
}

func (m *Message) GetProto2Field() *testdata.SubDefaults {
	if m != nil {
		return m.Proto2Field
	}
	return nil
}

func (m *Message) GetProto2Value() map[string]*testdata.SubDefaults {
	if m != nil {
		return m.Proto2Value
	}
	return nil
}

type Nested struct {
	Bunny string `protobuf:"bytes,1,opt,name=bunny,proto3" json:"bunny,omitempty"`
	Cute  bool   `protobuf:"varint,2,opt,name=cute,proto3" json:"cute,omitempty"`
}

func (m *Nested) Reset()                    { *m = Nested{} }
func (m *Nested) String() string            { return proto.CompactTextString(m) }
func (*Nested) ProtoMessage()               {}
func (*Nested) Descriptor() ([]byte, []int) { return fileDescriptorProto3, []int{1} }

type MessageWithMap struct {
	ByteMapping map[bool][]byte `protobuf:"bytes,1,rep,name=byte_mapping,json=byteMapping" json:"byte_mapping,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *MessageWithMap) Reset()                    { *m = MessageWithMap{} }
func (m *MessageWithMap) String() string            { return proto.CompactTextString(m) }
func (*MessageWithMap) ProtoMessage()               {}
func (*MessageWithMap) Descriptor() ([]byte, []int) { return fileDescriptorProto3, []int{2} }

func (m *MessageWithMap) GetByteMapping() map[bool][]byte {
	if m != nil {
		return m.ByteMapping
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "proto3_proto.Message")
	proto.RegisterType((*Nested)(nil), "proto3_proto.Nested")
	proto.RegisterType((*MessageWithMap)(nil), "proto3_proto.MessageWithMap")
	proto.RegisterEnum("proto3_proto.Message_Humour", Message_Humour_name, Message_Humour_value)
}

func init() { proto.RegisterFile("proto3_proto/proto3.proto", fileDescriptorProto3) }

var fileDescriptorProto3 = []byte{
	// 580 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x53, 0xdd, 0x6a, 0xdb, 0x4c,
	0x10, 0xfd, 0xfc, 0x13, 0x59, 0x19, 0xc9, 0xf9, 0xc4, 0xe2, 0xc2, 0xd6, 0xf4, 0xc2, 0x75, 0xa1,
	0x84, 0xfe, 0xc8, 0xa0, 0x52, 0x08, 0xa5, 0xb4, 0xc4, 0x6e, 0x42, 0x4d, 0x6c, 0xd7, 0xac, 0xe3,
	0x86, 0x5e, 0x09, 0xc9, 0x5e, 0xcb, 0xa2, 0x96, 0x64, 0xa4, 0x55, 0xc1, 0x4f, 0xd1, 0x77, 0xe8,
	0x93, 0x76, 0x7f, 0xe4, 0x44, 0x09, 0x6e, 0x7b, 0xa5, 0xd9, 0x33, 0xe7, 0xcc, 0xcc, 0x9e, 0x59,
	0xc1, 0xe3, 0x6d, 0x9a, 0xb0, 0xe4, 0x8d, 0x2b, 0x3f, 0x3d, 0x75, 0xb0, 0xe5, 0x07, 0x99, 0xe5,
	0x54, 0xdb, 0x09, 0x42, 0xb6, 0xce, 0x7d, 0x7b, 0x91, 0x44, 0xbd, 0x20, 0x09, 0x0a, 0xae, 0x9f,
	0xaf, 0x54, 0xd0, 0x63, 0x34, 0x63, 0x4b, 0x8f, 0x79, 0x32, 0x50, 0x15, 0xba, 0x3f, 0x35, 0x68,
	0x8c, 0x69, 0x96, 0x79, 0x01, 0x45, 0x08, 0xea, 0xb1, 0x17, 0x51, 0x5c, 0xe9, 0x54, 0x4e, 0x8f,
	0x89, 0x8c, 0xd1, 0x19, 0xe8, 0xeb, 0x70, 0xe3, 0xa5, 0x21, 0xdb, 0xe1, 0x2a, 0xc7, 0x4f, 0x9c,
	0x27, 0x76, 0xb9, 0xa9, 0x5d, 0x88, 0xed, 0xcf, 0x79, 0x94, 0xe4, 0x29, 0xb9, 0x65, 0xa3, 0x0e,
	0x98, 0x6b, 0x1a, 0x06, 0x6b, 0xe6, 0x86, 0xb1, 0xbb, 0x88, 0x70, 0x8d, 0xab, 0x9b, 0x04, 0x14,
	0x36, 0x8c, 0x07, 0x91, 0xe8, 0x27, 0xc6, 0xc1, 0x75, 0x9e, 0x31, 0x89, 0x8c, 0xd1, 0x53, 0x30,
	0x53, 0x9a, 0xe5, 0x1b, 0xe6, 0x2e, 0x92, 0x3c, 0x66, 0xb8, 0xc1, 0x73, 0x35, 0x62, 0x28, 0x6c,
	0x20, 0x20, 0xf4, 0x0c, 0x9a, 0x2c, 0xcd, 0xa9, 0x9b, 0x2d, 0x12, 0x96, 0x45, 0x5e, 0x8c, 0x75,
	0xce, 0xd1, 0x89, 0x29, 0xc0, 0x59, 0x81, 0xa1, 0x16, 0x1c, 0xf1, 0x7c, 0x4a, 0xf1, 0x31, 0x4f,
	0x56, 0x89, 0x3a, 0x20, 0x0b, 0x6a, 0xdf, 0xe9, 0x0e, 0x1f, 0x75, 0x6a, 0xa7, 0x75, 0x22, 0x42,
	0xf4, 0x0a, 0xb4, 0x98, 0xbb, 0x41, 0x97, 0x58, 0xe3, 0x44, 0xc3, 0x69, 0xdd, 0xbf, 0xdd, 0x44,
	0xe6, 0x48, 0xc1, 0x41, 0x6f, 0xa1, 0x91, 0xba, 0xab, 0x3c, 0x8e, 0x77, 0xd8, 0xe2, 0x35, 0xfe,
	0x65, 0x86, 0x96, 0x5e, 0x0a, 0x2e, 0x7a, 0x0f, 0x0d, 0x46, 0xd3, 0xd4, 0x0b, 0x63, 0x0c, 0x5c,
	0x66, 0x38, 0xdd, 0xc3, 0xb2, 0x6b, 0x45, 0xba, 0x88, 0x59, 0xba, 0x23, 0x7b, 0x09, 0x5f, 0x81,
	0x5a, 0xb3, 0xe3, 0xae, 0x42, 0xba, 0x59, 0x62, 0x43, 0x0e, 0xfa, 0xc8, 0xde, 0xaf, 0xd3, 0x9e,
	0xe5, 0xfe, 0x27, 0xba, 0xf2, 0xb8, 0x41, 0x19, 0x31, 0x14, 0xf5, 0x52, 0x30, 0xd1, 0xf0, 0x56,
	0xf9, 0xc3, 0xdb, 0xe4, 0x14, 0x37, 0x65, 0xf3, 0xe7, 0x87, 0x9b, 0x4f, 0x25, 0xf3, 0xab, 0x20,
	0xaa, 0x01, 0x8a, 0x52, 0x12, 0x69, 0x4f, 0xc1, 0x2c, 0x4f, 0xb7, 0x77, 0x52, 0x3d, 0x15, 0xe9,
	0xe4, 0x0b, 0x38, 0x52, 0x5d, 0xaa, 0x7f, 0x31, 0x52, 0x51, 0xde, 0x55, 0xcf, 0x2a, 0xed, 0x39,
	0x58, 0x0f, 0x5b, 0x1e, 0xa8, 0xfa, 0xf2, 0x7e, 0xd5, 0x3f, 0xdc, 0xfa, 0xae, 0x6c, 0xf7, 0x23,
	0x68, 0xca, 0x7d, 0x64, 0x40, 0x63, 0x3e, 0xb9, 0x9a, 0x7c, 0xb9, 0x99, 0x58, 0xff, 0x21, 0x1d,
	0xea, 0xd3, 0xf9, 0x64, 0x66, 0x55, 0x50, 0x13, 0x8e, 0x67, 0xa3, 0xf3, 0xe9, 0xec, 0x7a, 0x38,
	0xb8, 0xb2, 0xaa, 0xe8, 0x7f, 0x30, 0xfa, 0xc3, 0xd1, 0xc8, 0xed, 0x9f, 0x0f, 0x47, 0x17, 0xdf,
	0xac, 0x5a, 0xd7, 0x01, 0x4d, 0x0d, 0x2b, 0xde, 0x90, 0x2f, 0x77, 0xad, 0xe6, 0x51, 0x07, 0xf1,
	0x6a, 0x17, 0x39, 0x53, 0x03, 0xe9, 0x44, 0xc6, 0xdd, 0x5f, 0x15, 0x38, 0x29, 0x7c, 0xbc, 0xe1,
	0xff, 0xe0, 0xd8, 0xdb, 0x22, 0x6e, 0x98, 0xbf, 0x63, 0xd4, 0x8d, 0xbc, 0xed, 0x36, 0x8c, 0x03,
	0x5e, 0x43, 0x78, 0xff, 0xfa, 0xa0, 0xf7, 0x85, 0xc6, 0xee, 0x73, 0xc1, 0x58, 0xf1, 0x8b, 0x15,
	0xf8, 0x77, 0x48, 0xfb, 0x03, 0x58, 0x0f, 0x09, 0x65, 0xc3, 0x74, 0x65, 0x58, 0xab, 0x6c, 0x98,
	0x59, 0x72, 0xc6, 0xd7, 0x54, 0xeb, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x75, 0x74, 0x87, 0xb4,
	0x50, 0x04, 0x00, 0x00,
}
