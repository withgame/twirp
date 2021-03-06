// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/commo.proto

/*
Package common is a generated protocol buffer package.

It is generated from these files:
	common/commo.proto

It has these top-level messages:
	Empty
*/
package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Interface common.
type Empty struct {
	Latitude  int32 `protobuf:"varint,1,opt,name=latitude" json:"latitude,omitempty"`
	Longitude int32 `protobuf:"varint,2,opt,name=longitude" json:"longitude,omitempty"`
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Empty) GetLatitude() int32 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *Empty) GetLongitude() int32 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func init() {
	proto.RegisterType((*Empty)(nil), "common.Empty")
}

func init() { proto.RegisterFile("common/commo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 94 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0xd3, 0x07, 0x53, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x6c, 0x10, 0x31, 0x25,
	0x47, 0x2e, 0x56, 0xd7, 0xdc, 0x82, 0x92, 0x4a, 0x21, 0x29, 0x2e, 0x8e, 0x9c, 0xc4, 0x92, 0xcc,
	0x92, 0xd2, 0x94, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x38, 0x5f, 0x48, 0x86, 0x8b,
	0x33, 0x27, 0x3f, 0x2f, 0x1d, 0x22, 0xc9, 0x04, 0x96, 0x44, 0x08, 0x24, 0xb1, 0x81, 0x4d, 0x34,
	0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x7d, 0x65, 0x1f, 0x67, 0x00, 0x00, 0x00,
}
