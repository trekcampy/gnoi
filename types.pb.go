// Code generated by protoc-gen-go. DO NOT EDIT.
// source: types.proto

/*
Package gnoi is a generated protocol buffer package.

It is generated from these files:
	types.proto

It has these top-level messages:
	HashType
	Path
	PathElem
*/
package gnoi

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

// Generic Layer 3 Protocol enumeration.
type L3Protocol int32

const (
	L3Protocol_UNSPECIFIED L3Protocol = 0
	L3Protocol_IPV4        L3Protocol = 1
	L3Protocol_IPV6        L3Protocol = 2
)

var L3Protocol_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "IPV4",
	2: "IPV6",
}
var L3Protocol_value = map[string]int32{
	"UNSPECIFIED": 0,
	"IPV4":        1,
	"IPV6":        2,
}

func (x L3Protocol) String() string {
	return proto.EnumName(L3Protocol_name, int32(x))
}
func (L3Protocol) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type HashType_HashMethod int32

const (
	HashType_UNSPECIFIED HashType_HashMethod = 0
	HashType_SHA256      HashType_HashMethod = 1
	HashType_SHA512      HashType_HashMethod = 2
	HashType_MD5         HashType_HashMethod = 3
)

var HashType_HashMethod_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "SHA256",
	2: "SHA512",
	3: "MD5",
}
var HashType_HashMethod_value = map[string]int32{
	"UNSPECIFIED": 0,
	"SHA256":      1,
	"SHA512":      2,
	"MD5":         3,
}

func (x HashType_HashMethod) String() string {
	return proto.EnumName(HashType_HashMethod_name, int32(x))
}
func (HashType_HashMethod) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// HashType defines the valid hash methods for data verification.  UNSPECIFIED
// should be treated an error.
type HashType struct {
	Method HashType_HashMethod `protobuf:"varint,1,opt,name=method,enum=gnoi.HashType_HashMethod" json:"method,omitempty"`
	Hash   []byte              `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *HashType) Reset()                    { *m = HashType{} }
func (m *HashType) String() string            { return proto.CompactTextString(m) }
func (*HashType) ProtoMessage()               {}
func (*HashType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HashType) GetMethod() HashType_HashMethod {
	if m != nil {
		return m.Method
	}
	return HashType_UNSPECIFIED
}

func (m *HashType) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

// Path encodes a data tree path as a series of repeated strings, with
// each element of the path representing a data tree node name and the
// associated attributes.
// Reference: gNMI Specification Section 2.2.2.
type Path struct {
	Origin string      `protobuf:"bytes,2,opt,name=origin" json:"origin,omitempty"`
	Elem   []*PathElem `protobuf:"bytes,3,rep,name=elem" json:"elem,omitempty"`
}

func (m *Path) Reset()                    { *m = Path{} }
func (m *Path) String() string            { return proto.CompactTextString(m) }
func (*Path) ProtoMessage()               {}
func (*Path) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Path) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Path) GetElem() []*PathElem {
	if m != nil {
		return m.Elem
	}
	return nil
}

// PathElem encodes an element of a gNMI path, along with any attributes (keys)
// that may be associated with it.
// Reference: gNMI Specification Section 2.2.2.
type PathElem struct {
	Name string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Key  map[string]string `protobuf:"bytes,2,rep,name=key" json:"key,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *PathElem) Reset()                    { *m = PathElem{} }
func (m *PathElem) String() string            { return proto.CompactTextString(m) }
func (*PathElem) ProtoMessage()               {}
func (*PathElem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PathElem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PathElem) GetKey() map[string]string {
	if m != nil {
		return m.Key
	}
	return nil
}

func init() {
	proto.RegisterType((*HashType)(nil), "gnoi.HashType")
	proto.RegisterType((*Path)(nil), "gnoi.Path")
	proto.RegisterType((*PathElem)(nil), "gnoi.PathElem")
	proto.RegisterEnum("gnoi.L3Protocol", L3Protocol_name, L3Protocol_value)
	proto.RegisterEnum("gnoi.HashType_HashMethod", HashType_HashMethod_name, HashType_HashMethod_value)
}

func init() { proto.RegisterFile("types.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x41, 0x4f, 0x83, 0x40,
	0x10, 0x85, 0x5d, 0x40, 0xa4, 0x53, 0x53, 0x37, 0x13, 0xa3, 0xe8, 0xa9, 0xe1, 0x54, 0x3d, 0x90,
	0x94, 0xda, 0xc6, 0x78, 0x30, 0x51, 0x8b, 0x69, 0xa3, 0x35, 0x64, 0xab, 0xde, 0x51, 0x37, 0x85,
	0x08, 0x2c, 0x01, 0x34, 0xe1, 0xe8, 0x4f, 0xf0, 0x1f, 0x9b, 0x5d, 0x20, 0xc6, 0x78, 0xfb, 0x66,
	0xe6, 0xcd, 0xcb, 0xdb, 0x59, 0xe8, 0x57, 0x75, 0xce, 0x4b, 0x37, 0x2f, 0x44, 0x25, 0xd0, 0xd8,
	0x64, 0x22, 0x76, 0xbe, 0x09, 0x58, 0x8b, 0xb0, 0x8c, 0x1e, 0xeb, 0x9c, 0xe3, 0x18, 0xcc, 0x94,
	0x57, 0x91, 0x78, 0xb3, 0xc9, 0x90, 0x8c, 0x06, 0xde, 0x91, 0x2b, 0x35, 0x6e, 0x37, 0x57, 0xb0,
	0x52, 0x02, 0xd6, 0x0a, 0x11, 0xc1, 0x88, 0xc2, 0x32, 0xb2, 0xb5, 0x21, 0x19, 0xed, 0x32, 0xc5,
	0xce, 0x25, 0xc0, 0xaf, 0x12, 0xf7, 0xa0, 0xff, 0xf4, 0xb0, 0x0e, 0xfc, 0x9b, 0xe5, 0xed, 0xd2,
	0x9f, 0xd3, 0x2d, 0x04, 0x30, 0xd7, 0x8b, 0x2b, 0x6f, 0x3a, 0xa3, 0xa4, 0xe5, 0xe9, 0xd8, 0xa3,
	0x1a, 0xee, 0x80, 0xbe, 0x9a, 0x4f, 0xa9, 0xee, 0x5c, 0x83, 0x11, 0x84, 0x55, 0x84, 0x07, 0x60,
	0x8a, 0x22, 0xde, 0xc4, 0x99, 0x72, 0xef, 0xb1, 0xb6, 0x42, 0x07, 0x0c, 0x9e, 0xf0, 0xd4, 0xd6,
	0x87, 0xfa, 0xa8, 0xef, 0x0d, 0x9a, 0x90, 0x72, 0xc3, 0x4f, 0x78, 0xca, 0xd4, 0xcc, 0xf9, 0x22,
	0x60, 0x75, 0x2d, 0x19, 0x32, 0x0b, 0x53, 0xae, 0x5e, 0xd5, 0x63, 0x8a, 0xf1, 0x04, 0xf4, 0x77,
	0x5e, 0xdb, 0x9a, 0xf2, 0x38, 0xfc, 0xeb, 0xe1, 0xde, 0xf1, 0xda, 0xcf, 0xaa, 0xa2, 0x66, 0x52,
	0x73, 0x3c, 0x03, 0xab, 0x6b, 0x20, 0x6d, 0xd6, 0x1a, 0x27, 0x89, 0xb8, 0x0f, 0xdb, 0x9f, 0x61,
	0xf2, 0xc1, 0xdb, 0x90, 0x4d, 0x71, 0xa1, 0x9d, 0x93, 0xd3, 0x31, 0xc0, 0xfd, 0x24, 0x90, 0xc7,
	0x7e, 0x15, 0xc9, 0xff, 0x3b, 0x58, 0x60, 0x2c, 0x83, 0xe7, 0x33, 0x4a, 0x5a, 0x9a, 0x51, 0xed,
	0xc5, 0x54, 0x7f, 0x33, 0xf9, 0x09, 0x00, 0x00, 0xff, 0xff, 0xea, 0xe4, 0xaf, 0x79, 0xaa, 0x01,
	0x00, 0x00,
}
