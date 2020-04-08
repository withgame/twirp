// Code generated by protoc-gen-go. DO NOT EDIT.
// source: routeguide/hello.proto

/*
Package routeguide is a generated protocol buffer package.

It is generated from these files:
	routeguide/hello.proto

It has these top-level messages:
	Empty
*/
package routeguide

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "twirp/protoc-gen-twirp/pb/common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Empty struct {
	Latitude2  int32 `protobuf:"varint,1,opt,name=latitude2" json:"latitude2,omitempty"`
	Longitude2 int32 `protobuf:"varint,2,opt,name=longitude2" json:"longitude2,omitempty"`
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Empty) GetLatitude2() int32 {
	if m != nil {
		return m.Latitude2
	}
	return 0
}

func (m *Empty) GetLongitude2() int32 {
	if m != nil {
		return m.Longitude2
	}
	return 0
}

func init() {
	proto.RegisterType((*Empty)(nil), "routeguide.Empty")
}

func init() { proto.RegisterFile("routeguide/hello.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 143 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0xca, 0x2f, 0x2d,
	0x49, 0x4d, 0x2f, 0xcd, 0x4c, 0x49, 0xd5, 0xcf, 0x48, 0xcd, 0xc9, 0xc9, 0xd7, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xe2, 0x42, 0x88, 0x4b, 0x09, 0x25, 0xe7, 0xe7, 0xe6, 0xe6, 0xe7, 0xe9, 0x83,
	0x29, 0x88, 0xbc, 0x92, 0x2b, 0x17, 0xab, 0x6b, 0x6e, 0x41, 0x49, 0xa5, 0x90, 0x0c, 0x17, 0x67,
	0x4e, 0x62, 0x49, 0x66, 0x49, 0x69, 0x4a, 0xaa, 0x91, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x6b, 0x10,
	0x42, 0x40, 0x48, 0x8e, 0x8b, 0x2b, 0x27, 0x3f, 0x2f, 0x1d, 0x2a, 0xcd, 0x04, 0x96, 0x46, 0x12,
	0x31, 0xb2, 0xe2, 0xe2, 0x0a, 0x02, 0x59, 0xe4, 0x0e, 0xb2, 0x48, 0x48, 0x07, 0xca, 0x83, 0x98,
	0xcc, 0xab, 0x07, 0xb1, 0x57, 0x0f, 0xcc, 0x95, 0x42, 0xe5, 0x2a, 0x31, 0x24, 0xb1, 0x81, 0x5d,
	0x62, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x66, 0x47, 0x03, 0x88, 0xc3, 0x00, 0x00, 0x00,
}