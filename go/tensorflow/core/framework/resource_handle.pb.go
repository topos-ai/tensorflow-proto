// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/resource_handle.proto

package framework

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

// Protocol buffer representing a handle to a tensorflow resource. Handles are
// not valid across executions, but can be serialized back and forth from within
// a single run.
type ResourceHandleProto struct {
	// Unique name for the device containing the resource.
	Device string `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	// Container in which this resource is placed.
	Container string `protobuf:"bytes,2,opt,name=container,proto3" json:"container,omitempty"`
	// Unique name of this resource.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Hash code for the type of the resource. Is only valid in the same device
	// and in the same execution.
	HashCode uint64 `protobuf:"varint,4,opt,name=hash_code,json=hashCode,proto3" json:"hash_code,omitempty"`
	// For debug-only, the name of the type pointed to by this handle, if
	// available.
	MaybeTypeName string `protobuf:"bytes,5,opt,name=maybe_type_name,json=maybeTypeName,proto3" json:"maybe_type_name,omitempty"`
	// Data types and shapes for the underlying resource.
	DtypesAndShapes      []*ResourceHandleProto_DtypeAndShape `protobuf:"bytes,6,rep,name=dtypes_and_shapes,json=dtypesAndShapes,proto3" json:"dtypes_and_shapes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *ResourceHandleProto) Reset()         { *m = ResourceHandleProto{} }
func (m *ResourceHandleProto) String() string { return proto.CompactTextString(m) }
func (*ResourceHandleProto) ProtoMessage()    {}
func (*ResourceHandleProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_a36024d2bd9a2afd, []int{0}
}

func (m *ResourceHandleProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceHandleProto.Unmarshal(m, b)
}
func (m *ResourceHandleProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceHandleProto.Marshal(b, m, deterministic)
}
func (m *ResourceHandleProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceHandleProto.Merge(m, src)
}
func (m *ResourceHandleProto) XXX_Size() int {
	return xxx_messageInfo_ResourceHandleProto.Size(m)
}
func (m *ResourceHandleProto) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceHandleProto.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceHandleProto proto.InternalMessageInfo

func (m *ResourceHandleProto) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *ResourceHandleProto) GetContainer() string {
	if m != nil {
		return m.Container
	}
	return ""
}

func (m *ResourceHandleProto) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ResourceHandleProto) GetHashCode() uint64 {
	if m != nil {
		return m.HashCode
	}
	return 0
}

func (m *ResourceHandleProto) GetMaybeTypeName() string {
	if m != nil {
		return m.MaybeTypeName
	}
	return ""
}

func (m *ResourceHandleProto) GetDtypesAndShapes() []*ResourceHandleProto_DtypeAndShape {
	if m != nil {
		return m.DtypesAndShapes
	}
	return nil
}

// Protocol buffer representing a pair of (data type, tensor shape).
type ResourceHandleProto_DtypeAndShape struct {
	Dtype                DataType          `protobuf:"varint,1,opt,name=dtype,proto3,enum=tensorflow.DataType" json:"dtype,omitempty"`
	Shape                *TensorShapeProto `protobuf:"bytes,2,opt,name=shape,proto3" json:"shape,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ResourceHandleProto_DtypeAndShape) Reset()         { *m = ResourceHandleProto_DtypeAndShape{} }
func (m *ResourceHandleProto_DtypeAndShape) String() string { return proto.CompactTextString(m) }
func (*ResourceHandleProto_DtypeAndShape) ProtoMessage()    {}
func (*ResourceHandleProto_DtypeAndShape) Descriptor() ([]byte, []int) {
	return fileDescriptor_a36024d2bd9a2afd, []int{0, 0}
}

func (m *ResourceHandleProto_DtypeAndShape) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceHandleProto_DtypeAndShape.Unmarshal(m, b)
}
func (m *ResourceHandleProto_DtypeAndShape) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceHandleProto_DtypeAndShape.Marshal(b, m, deterministic)
}
func (m *ResourceHandleProto_DtypeAndShape) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceHandleProto_DtypeAndShape.Merge(m, src)
}
func (m *ResourceHandleProto_DtypeAndShape) XXX_Size() int {
	return xxx_messageInfo_ResourceHandleProto_DtypeAndShape.Size(m)
}
func (m *ResourceHandleProto_DtypeAndShape) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceHandleProto_DtypeAndShape.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceHandleProto_DtypeAndShape proto.InternalMessageInfo

func (m *ResourceHandleProto_DtypeAndShape) GetDtype() DataType {
	if m != nil {
		return m.Dtype
	}
	return DataType_DT_INVALID
}

func (m *ResourceHandleProto_DtypeAndShape) GetShape() *TensorShapeProto {
	if m != nil {
		return m.Shape
	}
	return nil
}

func init() {
	proto.RegisterType((*ResourceHandleProto)(nil), "tensorflow.ResourceHandleProto")
	proto.RegisterType((*ResourceHandleProto_DtypeAndShape)(nil), "tensorflow.ResourceHandleProto.DtypeAndShape")
}

func init() {
	proto.RegisterFile("tensorflow/core/framework/resource_handle.proto", fileDescriptor_a36024d2bd9a2afd)
}

var fileDescriptor_a36024d2bd9a2afd = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x4f, 0x4f, 0xc2, 0x40,
	0x10, 0xc5, 0x53, 0xfe, 0x45, 0x96, 0x00, 0x71, 0x35, 0xa6, 0x41, 0x0e, 0xc4, 0x44, 0x43, 0x8c,
	0xb4, 0x49, 0xfd, 0x04, 0x20, 0x07, 0x4f, 0xc6, 0x54, 0x2e, 0x7a, 0x69, 0x96, 0xee, 0x40, 0x1b,
	0xe9, 0x4e, 0xb3, 0x5b, 0x24, 0x7c, 0x6b, 0x8f, 0x1e, 0x4d, 0xa7, 0x28, 0xd5, 0x88, 0xb7, 0xee,
	0xcc, 0xef, 0xcd, 0xbc, 0xd7, 0x5d, 0xe6, 0x66, 0xa0, 0x0c, 0xea, 0xc5, 0x0a, 0x37, 0x6e, 0x88,
	0x1a, 0xdc, 0x85, 0x16, 0x09, 0x6c, 0x50, 0xbf, 0xba, 0x1a, 0x0c, 0xae, 0x75, 0x08, 0x41, 0x24,
	0x94, 0x5c, 0x81, 0x93, 0x6a, 0xcc, 0x90, 0xb3, 0xbd, 0xa0, 0x77, 0x73, 0x58, 0x5c, 0x74, 0x02,
	0x13, 0x89, 0x74, 0xa7, 0xec, 0x5d, 0xfe, 0x43, 0x6f, 0x53, 0x30, 0x05, 0x76, 0xf1, 0x5e, 0x61,
	0x27, 0xfe, 0x6e, 0xf5, 0x3d, 0x6d, 0x7e, 0xa4, 0xc5, 0x67, 0xac, 0x21, 0xe1, 0x2d, 0x0e, 0xc1,
	0xb6, 0x06, 0xd6, 0xb0, 0xe9, 0xef, 0x4e, 0xbc, 0xcf, 0x9a, 0x21, 0xaa, 0x4c, 0xc4, 0x0a, 0xb4,
	0x5d, 0xa1, 0xd6, 0xbe, 0xc0, 0x39, 0xab, 0x29, 0x91, 0x80, 0x5d, 0xa5, 0x06, 0x7d, 0xf3, 0x73,
	0xd6, 0x8c, 0x84, 0x89, 0x82, 0x10, 0x25, 0xd8, 0xb5, 0x81, 0x35, 0xac, 0xf9, 0x47, 0x79, 0xe1,
	0x0e, 0x25, 0xf0, 0x2b, 0xd6, 0x4d, 0xc4, 0x76, 0x0e, 0x41, 0xee, 0x29, 0x20, 0x6d, 0x9d, 0xb4,
	0x6d, 0x2a, 0xcf, 0xb6, 0x29, 0x3c, 0xe4, 0x43, 0x9e, 0xd9, 0xb1, 0x24, 0xdb, 0x81, 0x50, 0xb2,
	0xc8, 0x69, 0xec, 0xc6, 0xa0, 0x3a, 0x6c, 0x79, 0x23, 0x67, 0x9f, 0xd4, 0xf9, 0x23, 0x8a, 0x33,
	0xcd, 0x85, 0x63, 0x25, 0x9f, 0x72, 0x95, 0xdf, 0x2d, 0xe6, 0x7c, 0x9d, 0x4d, 0x0f, 0x59, 0xfb,
	0x07, 0xc1, 0xaf, 0x59, 0x9d, 0x18, 0x4a, 0xde, 0xf1, 0x4e, 0xcb, 0xf3, 0xa7, 0x22, 0x13, 0xb9,
	0x29, 0xbf, 0x40, 0xb8, 0xc7, 0xea, 0x64, 0x86, 0x7e, 0x45, 0xcb, 0xeb, 0x97, 0xd9, 0x19, 0x7d,
	0xd2, 0x4c, 0x32, 0xe2, 0x17, 0xe8, 0x44, 0x33, 0x1b, 0xf5, 0xb2, 0x4c, 0x7e, 0x5f, 0xcd, 0xa4,
	0xf3, 0x2b, 0x80, 0xf5, 0x32, 0x5e, 0xc6, 0x59, 0xb4, 0x9e, 0x3b, 0x21, 0x26, 0x6e, 0x86, 0x29,
	0x9a, 0x91, 0x88, 0x4b, 0xcf, 0x68, 0x44, 0xd7, 0xe8, 0x2e, 0xf1, 0xf0, 0xd3, 0xfa, 0xb0, 0xac,
	0x79, 0x83, 0xb0, 0xdb, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x89, 0x00, 0xb4, 0x2e, 0x81, 0x02,
	0x00, 0x00,
}