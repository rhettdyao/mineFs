// Code generated by protoc-gen-go. DO NOT EDIT.
// source: raft_kv.proto

package mineFs_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type PutRequest struct {
	Seq                  int64    `protobuf:"varint,1,opt,name=seq" json:"seq,omitempty"`
	Key                  []byte   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Del                  bool     `protobuf:"varint,4,opt,name=del" json:"del,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutRequest) Reset()         { *m = PutRequest{} }
func (m *PutRequest) String() string { return proto.CompactTextString(m) }
func (*PutRequest) ProtoMessage()    {}
func (*PutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_raft_kv_75e3b60daff0f8cd, []int{0}
}
func (m *PutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutRequest.Unmarshal(m, b)
}
func (m *PutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutRequest.Marshal(b, m, deterministic)
}
func (dst *PutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutRequest.Merge(dst, src)
}
func (m *PutRequest) XXX_Size() int {
	return xxx_messageInfo_PutRequest.Size(m)
}
func (m *PutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PutRequest proto.InternalMessageInfo

func (m *PutRequest) GetSeq() int64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *PutRequest) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *PutRequest) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *PutRequest) GetDel() bool {
	if m != nil {
		return m.Del
	}
	return false
}

type PutResponse struct {
	Seq                  int64    `protobuf:"varint,1,opt,name=seq" json:"seq,omitempty"`
	Status               int64    `protobuf:"varint,2,opt,name=status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutResponse) Reset()         { *m = PutResponse{} }
func (m *PutResponse) String() string { return proto.CompactTextString(m) }
func (*PutResponse) ProtoMessage()    {}
func (*PutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_raft_kv_75e3b60daff0f8cd, []int{1}
}
func (m *PutResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutResponse.Unmarshal(m, b)
}
func (m *PutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutResponse.Marshal(b, m, deterministic)
}
func (dst *PutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutResponse.Merge(dst, src)
}
func (m *PutResponse) XXX_Size() int {
	return xxx_messageInfo_PutResponse.Size(m)
}
func (m *PutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PutResponse proto.InternalMessageInfo

func (m *PutResponse) GetSeq() int64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *PutResponse) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

type GetRequest struct {
	Seq                  int64    `protobuf:"varint,1,opt,name=seq" json:"seq,omitempty"`
	Key                  []byte   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_raft_kv_75e3b60daff0f8cd, []int{2}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (dst *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(dst, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetSeq() int64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *GetRequest) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type GetResponse struct {
	Seq                  int64    `protobuf:"varint,1,opt,name=seq" json:"seq,omitempty"`
	Status               int64    `protobuf:"varint,2,opt,name=status" json:"status,omitempty"`
	Value                []byte   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_raft_kv_75e3b60daff0f8cd, []int{3}
}
func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (dst *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(dst, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetSeq() int64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *GetResponse) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *GetResponse) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*PutRequest)(nil), "mineFs.proto.PutRequest")
	proto.RegisterType((*PutResponse)(nil), "mineFs.proto.PutResponse")
	proto.RegisterType((*GetRequest)(nil), "mineFs.proto.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "mineFs.proto.GetResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RaftKvClient is the client API for RaftKv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RaftKvClient interface {
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type raftKvClient struct {
	cc *grpc.ClientConn
}

func NewRaftKvClient(cc *grpc.ClientConn) RaftKvClient {
	return &raftKvClient{cc}
}

func (c *raftKvClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := c.cc.Invoke(ctx, "/mineFs.proto.RaftKv/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftKvClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/mineFs.proto.RaftKv/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RaftKv service

type RaftKvServer interface {
	Put(context.Context, *PutRequest) (*PutResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
}

func RegisterRaftKvServer(s *grpc.Server, srv RaftKvServer) {
	s.RegisterService(&_RaftKv_serviceDesc, srv)
}

func _RaftKv_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftKvServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mineFs.proto.RaftKv/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftKvServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftKv_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftKvServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mineFs.proto.RaftKv/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftKvServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RaftKv_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mineFs.proto.RaftKv",
	HandlerType: (*RaftKvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Put",
			Handler:    _RaftKv_Put_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _RaftKv_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "raft_kv.proto",
}

func init() { proto.RegisterFile("raft_kv.proto", fileDescriptor_raft_kv_75e3b60daff0f8cd) }

var fileDescriptor_raft_kv_75e3b60daff0f8cd = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4a, 0x4c, 0x2b,
	0x89, 0xcf, 0x2e, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xc9, 0xcd, 0xcc, 0x4b, 0x75,
	0x2b, 0x86, 0xf0, 0x94, 0x22, 0xb8, 0xb8, 0x02, 0x4a, 0x4b, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b,
	0x4b, 0x84, 0x04, 0xb8, 0x98, 0x8b, 0x53, 0x0b, 0x25, 0x18, 0x15, 0x18, 0x35, 0x98, 0x83, 0x40,
	0x4c, 0x90, 0x48, 0x76, 0x6a, 0xa5, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x4f, 0x10, 0x88, 0x29, 0x24,
	0xc2, 0xc5, 0x5a, 0x96, 0x98, 0x53, 0x9a, 0x2a, 0xc1, 0x0c, 0x16, 0x83, 0x70, 0x40, 0xea, 0x52,
	0x52, 0x73, 0x24, 0x58, 0x14, 0x18, 0x35, 0x38, 0x82, 0x40, 0x4c, 0x25, 0x73, 0x2e, 0x6e, 0xb0,
	0xc9, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x58, 0x8c, 0x16, 0xe3, 0x62, 0x2b, 0x2e, 0x49, 0x2c,
	0x29, 0x2d, 0x06, 0x9b, 0xce, 0x1c, 0x04, 0xe5, 0x29, 0x19, 0x70, 0x71, 0xb9, 0xa7, 0x92, 0xe2,
	0x24, 0x25, 0x5f, 0x2e, 0x6e, 0xb0, 0x0e, 0x52, 0xad, 0xc2, 0xee, 0x17, 0xa3, 0x06, 0x46, 0x2e,
	0xb6, 0xa0, 0xc4, 0xb4, 0x12, 0xef, 0x32, 0x21, 0x2b, 0x2e, 0xe6, 0x80, 0xd2, 0x12, 0x21, 0x09,
	0x3d, 0xe4, 0x40, 0xd3, 0x43, 0x84, 0x98, 0x94, 0x24, 0x16, 0x19, 0xa8, 0x33, 0xac, 0xb8, 0x98,
	0xdd, 0x53, 0x31, 0xf4, 0x22, 0xbc, 0x86, 0xae, 0x17, 0xc9, 0x0b, 0x49, 0x6c, 0x60, 0x21, 0x63,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x51, 0xab, 0xa5, 0xbc, 0x01, 0x00, 0x00,
}