// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/remote_fused_graph_execute_info.proto

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
type RemoteFusedGraphExecuteInfo struct {
	// Definition of remote graph
	RemoteGraph *GraphDef `protobuf:"bytes,1,opt,name=remote_graph,json=remoteGraph,proto3" json:"remote_graph,omitempty"`
	// Remote fused graph input node name
	GraphInputNodeName []string `protobuf:"bytes,2,rep,name=graph_input_node_name,json=graphInputNodeName,proto3" json:"graph_input_node_name,omitempty"`
	// Remote fused graph output node name
	GraphOutputNodeName []string `protobuf:"bytes,3,rep,name=graph_output_node_name,json=graphOutputNodeName,proto3" json:"graph_output_node_name,omitempty"`
	// Executor's name
	ExecutorName string `protobuf:"bytes,4,opt,name=executor_name,json=executorName,proto3" json:"executor_name,omitempty"`
	// Optional: Parameters given to the executor
	SerializedExecutorParameters []byte `protobuf:"bytes,5,opt,name=serialized_executor_parameters,json=serializedExecutorParameters,proto3" json:"serialized_executor_parameters,omitempty"`
	// Optional: Default graph input tensor shape used to allocate memory
	// before executing op
	DefaultGraphInputTensorShape []*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto `protobuf:"bytes,6,rep,name=default_graph_input_tensor_shape,json=defaultGraphInputTensorShape,proto3" json:"default_graph_input_tensor_shape,omitempty"`
	// Optional: Default graph input tensor shape used to allocate memory
	// before executing op
	// TODO(satok): Remote output tensor shape once shape information is stored
	// in NodeDef
	DefaultGraphOutputTensorShape []*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto `protobuf:"bytes,7,rep,name=default_graph_output_tensor_shape,json=defaultGraphOutputTensorShape,proto3" json:"default_graph_output_tensor_shape,omitempty"`
	XXX_NoUnkeyedLiteral          struct{}                                            `json:"-"`
	XXX_unrecognized              []byte                                              `json:"-"`
	XXX_sizecache                 int32                                               `json:"-"`
}

func (m *RemoteFusedGraphExecuteInfo) Reset()         { *m = RemoteFusedGraphExecuteInfo{} }
func (m *RemoteFusedGraphExecuteInfo) String() string { return proto.CompactTextString(m) }
func (*RemoteFusedGraphExecuteInfo) ProtoMessage()    {}
func (*RemoteFusedGraphExecuteInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_c15f13da5b37f691, []int{0}
}

func (m *RemoteFusedGraphExecuteInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoteFusedGraphExecuteInfo.Unmarshal(m, b)
}
func (m *RemoteFusedGraphExecuteInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoteFusedGraphExecuteInfo.Marshal(b, m, deterministic)
}
func (m *RemoteFusedGraphExecuteInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoteFusedGraphExecuteInfo.Merge(m, src)
}
func (m *RemoteFusedGraphExecuteInfo) XXX_Size() int {
	return xxx_messageInfo_RemoteFusedGraphExecuteInfo.Size(m)
}
func (m *RemoteFusedGraphExecuteInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoteFusedGraphExecuteInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RemoteFusedGraphExecuteInfo proto.InternalMessageInfo

func (m *RemoteFusedGraphExecuteInfo) GetRemoteGraph() *GraphDef {
	if m != nil {
		return m.RemoteGraph
	}
	return nil
}

func (m *RemoteFusedGraphExecuteInfo) GetGraphInputNodeName() []string {
	if m != nil {
		return m.GraphInputNodeName
	}
	return nil
}

func (m *RemoteFusedGraphExecuteInfo) GetGraphOutputNodeName() []string {
	if m != nil {
		return m.GraphOutputNodeName
	}
	return nil
}

func (m *RemoteFusedGraphExecuteInfo) GetExecutorName() string {
	if m != nil {
		return m.ExecutorName
	}
	return ""
}

func (m *RemoteFusedGraphExecuteInfo) GetSerializedExecutorParameters() []byte {
	if m != nil {
		return m.SerializedExecutorParameters
	}
	return nil
}

func (m *RemoteFusedGraphExecuteInfo) GetDefaultGraphInputTensorShape() []*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto {
	if m != nil {
		return m.DefaultGraphInputTensorShape
	}
	return nil
}

func (m *RemoteFusedGraphExecuteInfo) GetDefaultGraphOutputTensorShape() []*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto {
	if m != nil {
		return m.DefaultGraphOutputTensorShape
	}
	return nil
}

type RemoteFusedGraphExecuteInfo_TensorShapeTypeProto struct {
	Dtype                DataType          `protobuf:"varint,1,opt,name=dtype,proto3,enum=tensorflow.DataType" json:"dtype,omitempty"`
	Shape                *TensorShapeProto `protobuf:"bytes,2,opt,name=shape,proto3" json:"shape,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) Reset() {
	*m = RemoteFusedGraphExecuteInfo_TensorShapeTypeProto{}
}
func (m *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) String() string {
	return proto.CompactTextString(m)
}
func (*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) ProtoMessage() {}
func (*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_c15f13da5b37f691, []int{0, 0}
}

func (m *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoteFusedGraphExecuteInfo_TensorShapeTypeProto.Unmarshal(m, b)
}
func (m *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoteFusedGraphExecuteInfo_TensorShapeTypeProto.Marshal(b, m, deterministic)
}
func (m *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoteFusedGraphExecuteInfo_TensorShapeTypeProto.Merge(m, src)
}
func (m *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) XXX_Size() int {
	return xxx_messageInfo_RemoteFusedGraphExecuteInfo_TensorShapeTypeProto.Size(m)
}
func (m *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoteFusedGraphExecuteInfo_TensorShapeTypeProto.DiscardUnknown(m)
}

var xxx_messageInfo_RemoteFusedGraphExecuteInfo_TensorShapeTypeProto proto.InternalMessageInfo

func (m *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) GetDtype() DataType {
	if m != nil {
		return m.Dtype
	}
	return DataType_DT_INVALID
}

func (m *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) GetShape() *TensorShapeProto {
	if m != nil {
		return m.Shape
	}
	return nil
}

func init() {
	proto.RegisterType((*RemoteFusedGraphExecuteInfo)(nil), "tensorflow.RemoteFusedGraphExecuteInfo")
	proto.RegisterType((*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto)(nil), "tensorflow.RemoteFusedGraphExecuteInfo.TensorShapeTypeProto")
}

func init() {
	proto.RegisterFile("tensorflow/core/framework/remote_fused_graph_execute_info.proto", fileDescriptor_c15f13da5b37f691)
}

var fileDescriptor_c15f13da5b37f691 = []byte{
	// 462 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x41, 0x6f, 0xd3, 0x40,
	0x10, 0x85, 0xb5, 0x0d, 0x29, 0xea, 0x26, 0x70, 0x58, 0x0a, 0xb2, 0x42, 0x40, 0x06, 0x84, 0x64,
	0x21, 0x6a, 0x8b, 0xf4, 0xc0, 0x05, 0x09, 0x51, 0xa5, 0x54, 0xbd, 0x94, 0xc8, 0xf4, 0xc4, 0x65,
	0xb5, 0x8d, 0xc7, 0x89, 0x45, 0xec, 0xb1, 0xd6, 0x6b, 0x4a, 0x39, 0x57, 0xfc, 0x1f, 0xfe, 0x1d,
	0x47, 0xe4, 0x59, 0xe3, 0x6c, 0x24, 0x92, 0x13, 0xc7, 0x64, 0xbe, 0xf7, 0xfc, 0xe6, 0x79, 0xcc,
	0xdf, 0x1b, 0x28, 0x2a, 0xd4, 0xe9, 0x0a, 0xaf, 0xa3, 0x39, 0x6a, 0x88, 0x52, 0xad, 0x72, 0xb8,
	0x46, 0xfd, 0x35, 0xd2, 0x90, 0xa3, 0x01, 0x99, 0xd6, 0x15, 0x24, 0x72, 0xa1, 0x55, 0xb9, 0x94,
	0xf0, 0x1d, 0xe6, 0xb5, 0x01, 0x99, 0x15, 0x29, 0x86, 0xa5, 0x46, 0x83, 0x82, 0xaf, 0x0d, 0x46,
	0x2f, 0xb7, 0x9b, 0x91, 0xde, 0x4a, 0x46, 0xaf, 0xb7, 0x63, 0x76, 0x22, 0xab, 0xa5, 0x2a, 0xa1,
	0xa5, 0x77, 0x98, 0x9a, 0x9b, 0x12, 0x2a, 0x8b, 0x3d, 0xff, 0xd5, 0xe7, 0x8f, 0x63, 0x4a, 0xfc,
	0xb1, 0x09, 0x7c, 0xd6, 0x3c, 0xef, 0xd4, 0xc6, 0x3d, 0x2f, 0x52, 0x14, 0x6f, 0xf9, 0xb0, 0x5d,
	0x88, 0xa2, 0x78, 0xcc, 0x67, 0xc1, 0x60, 0x72, 0x18, 0xae, 0xdd, 0x43, 0xd2, 0x4c, 0x21, 0x8d,
	0x07, 0x96, 0xa4, 0xdf, 0xe2, 0x0d, 0x7f, 0x68, 0x97, 0xcf, 0x8a, 0xb2, 0x36, 0xb2, 0xc0, 0x04,
	0x64, 0xa1, 0x72, 0xf0, 0xf6, 0xfc, 0x5e, 0x70, 0x10, 0x0b, 0x1a, 0x9e, 0x37, 0xb3, 0x0b, 0x4c,
	0xe0, 0x42, 0xe5, 0x20, 0x8e, 0xf9, 0x23, 0x2b, 0xc1, 0xda, 0x6c, 0x6a, 0x7a, 0xa4, 0x79, 0x40,
	0xd3, 0x4f, 0x34, 0xec, 0x44, 0x2f, 0xf8, 0x3d, 0x5b, 0x2f, 0x6a, 0xcb, 0xde, 0xf1, 0x59, 0x70,
	0x10, 0x0f, 0xff, 0xfe, 0x49, 0xd0, 0x94, 0x3f, 0xad, 0x40, 0x67, 0x6a, 0x95, 0xfd, 0x80, 0x44,
	0x76, 0x7c, 0xa9, 0x9a, 0x4e, 0x0c, 0xe8, 0xca, 0xeb, 0xfb, 0x2c, 0x18, 0xc6, 0xe3, 0x35, 0x75,
	0xda, 0x42, 0xb3, 0x8e, 0x11, 0xb7, 0x8c, 0xfb, 0x09, 0xa4, 0xaa, 0x5e, 0x19, 0xe9, 0xee, 0xe6,
	0xb6, 0xef, 0xed, 0xfb, 0xbd, 0x60, 0x30, 0x79, 0xe7, 0x16, 0xb4, 0xa3, 0xdf, 0xf0, 0x92, 0xb0,
	0xcf, 0x8d, 0xf4, 0xf2, 0xa6, 0x84, 0x59, 0xf3, 0x52, 0xe2, 0x71, 0xfb, 0x94, 0xb3, 0xae, 0x23,
	0x07, 0x13, 0x3f, 0x19, 0x7f, 0xb6, 0x19, 0xa3, 0xed, 0x6b, 0x23, 0xc7, 0xdd, 0xff, 0x90, 0xe3,
	0x89, 0x9b, 0xc3, 0xf6, 0xee, 0x70, 0xa3, 0x6f, 0xfc, 0xf0, 0x5f, 0x32, 0xf1, 0x8a, 0xf7, 0x93,
	0xe6, 0xc6, 0xe8, 0x58, 0xee, 0x6f, 0x1e, 0xcb, 0x54, 0x19, 0xd5, 0x90, 0xb1, 0x45, 0xc4, 0x84,
	0xf7, 0x6d, 0xde, 0x3d, 0x3a, 0xac, 0xb1, 0xcb, 0x3a, 0xe6, 0x36, 0x8f, 0x45, 0x4f, 0x6e, 0x19,
	0xf7, 0x50, 0x2f, 0x5c, 0xb4, 0x3b, 0xee, 0x13, 0x7f, 0xc7, 0x96, 0xe4, 0x32, 0x63, 0x5f, 0x3e,
	0x2c, 0x32, 0xb3, 0xac, 0xaf, 0xc2, 0x39, 0xe6, 0x91, 0xc1, 0x12, 0xab, 0x23, 0x95, 0x45, 0x6b,
	0xb7, 0x23, 0xfa, 0x34, 0xa2, 0x05, 0x46, 0x5b, 0xbf, 0xa1, 0xdf, 0x8c, 0x5d, 0xed, 0x13, 0x76,
	0xfc, 0x27, 0x00, 0x00, 0xff, 0xff, 0xdd, 0xcd, 0x5c, 0x66, 0x0c, 0x04, 0x00, 0x00,
}
