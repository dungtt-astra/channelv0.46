// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: channel/channel/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgOpenchannel struct {
	Creator      string      `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	PartA        string      `protobuf:"bytes,2,opt,name=partA,proto3" json:"partA,omitempty"`
	PartB        string      `protobuf:"bytes,3,opt,name=partB,proto3" json:"partB,omitempty"`
	CoinA        *types.Coin `protobuf:"bytes,4,opt,name=coinA,proto3" json:"coinA,omitempty"`
	CoinB        *types.Coin `protobuf:"bytes,5,opt,name=coinB,proto3" json:"coinB,omitempty"`
	MultisigAddr string      `protobuf:"bytes,6,opt,name=multisigAddr,proto3" json:"multisigAddr,omitempty"`
	Sequence     string      `protobuf:"bytes,7,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (m *MsgOpenchannel) Reset()         { *m = MsgOpenchannel{} }
func (m *MsgOpenchannel) String() string { return proto.CompactTextString(m) }
func (*MsgOpenchannel) ProtoMessage()    {}
func (*MsgOpenchannel) Descriptor() ([]byte, []int) {
	return fileDescriptor_89a552858311affe, []int{0}
}
func (m *MsgOpenchannel) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgOpenchannel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgOpenchannel.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgOpenchannel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgOpenchannel.Merge(m, src)
}
func (m *MsgOpenchannel) XXX_Size() int {
	return m.Size()
}
func (m *MsgOpenchannel) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgOpenchannel.DiscardUnknown(m)
}

var xxx_messageInfo_MsgOpenchannel proto.InternalMessageInfo

func (m *MsgOpenchannel) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgOpenchannel) GetPartA() string {
	if m != nil {
		return m.PartA
	}
	return ""
}

func (m *MsgOpenchannel) GetPartB() string {
	if m != nil {
		return m.PartB
	}
	return ""
}

func (m *MsgOpenchannel) GetCoinA() *types.Coin {
	if m != nil {
		return m.CoinA
	}
	return nil
}

func (m *MsgOpenchannel) GetCoinB() *types.Coin {
	if m != nil {
		return m.CoinB
	}
	return nil
}

func (m *MsgOpenchannel) GetMultisigAddr() string {
	if m != nil {
		return m.MultisigAddr
	}
	return ""
}

func (m *MsgOpenchannel) GetSequence() string {
	if m != nil {
		return m.Sequence
	}
	return ""
}

type MsgOpenchannelResponse struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *MsgOpenchannelResponse) Reset()         { *m = MsgOpenchannelResponse{} }
func (m *MsgOpenchannelResponse) String() string { return proto.CompactTextString(m) }
func (*MsgOpenchannelResponse) ProtoMessage()    {}
func (*MsgOpenchannelResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_89a552858311affe, []int{1}
}
func (m *MsgOpenchannelResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgOpenchannelResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgOpenchannelResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgOpenchannelResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgOpenchannelResponse.Merge(m, src)
}
func (m *MsgOpenchannelResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgOpenchannelResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgOpenchannelResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgOpenchannelResponse proto.InternalMessageInfo

func (m *MsgOpenchannelResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type MsgClosechannel struct {
	Creator      string      `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	MultisigAddr string      `protobuf:"bytes,2,opt,name=multisigAddr,proto3" json:"multisigAddr,omitempty"`
	PartA        string      `protobuf:"bytes,3,opt,name=partA,proto3" json:"partA,omitempty"`
	CoinA        *types.Coin `protobuf:"bytes,4,opt,name=coinA,proto3" json:"coinA,omitempty"`
	PartB        string      `protobuf:"bytes,5,opt,name=partB,proto3" json:"partB,omitempty"`
	CoinB        *types.Coin `protobuf:"bytes,6,opt,name=coinB,proto3" json:"coinB,omitempty"`
	Channelid    string      `protobuf:"bytes,7,opt,name=channelid,proto3" json:"channelid,omitempty"`
}

func (m *MsgClosechannel) Reset()         { *m = MsgClosechannel{} }
func (m *MsgClosechannel) String() string { return proto.CompactTextString(m) }
func (*MsgClosechannel) ProtoMessage()    {}
func (*MsgClosechannel) Descriptor() ([]byte, []int) {
	return fileDescriptor_89a552858311affe, []int{2}
}
func (m *MsgClosechannel) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgClosechannel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgClosechannel.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgClosechannel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgClosechannel.Merge(m, src)
}
func (m *MsgClosechannel) XXX_Size() int {
	return m.Size()
}
func (m *MsgClosechannel) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgClosechannel.DiscardUnknown(m)
}

var xxx_messageInfo_MsgClosechannel proto.InternalMessageInfo

func (m *MsgClosechannel) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgClosechannel) GetMultisigAddr() string {
	if m != nil {
		return m.MultisigAddr
	}
	return ""
}

func (m *MsgClosechannel) GetPartA() string {
	if m != nil {
		return m.PartA
	}
	return ""
}

func (m *MsgClosechannel) GetCoinA() *types.Coin {
	if m != nil {
		return m.CoinA
	}
	return nil
}

func (m *MsgClosechannel) GetPartB() string {
	if m != nil {
		return m.PartB
	}
	return ""
}

func (m *MsgClosechannel) GetCoinB() *types.Coin {
	if m != nil {
		return m.CoinB
	}
	return nil
}

func (m *MsgClosechannel) GetChannelid() string {
	if m != nil {
		return m.Channelid
	}
	return ""
}

type MsgClosechannelResponse struct {
}

func (m *MsgClosechannelResponse) Reset()         { *m = MsgClosechannelResponse{} }
func (m *MsgClosechannelResponse) String() string { return proto.CompactTextString(m) }
func (*MsgClosechannelResponse) ProtoMessage()    {}
func (*MsgClosechannelResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_89a552858311affe, []int{3}
}
func (m *MsgClosechannelResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgClosechannelResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgClosechannelResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgClosechannelResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgClosechannelResponse.Merge(m, src)
}
func (m *MsgClosechannelResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgClosechannelResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgClosechannelResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgClosechannelResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgOpenchannel)(nil), "channel.channel.MsgOpenchannel")
	proto.RegisterType((*MsgOpenchannelResponse)(nil), "channel.channel.MsgOpenchannelResponse")
	proto.RegisterType((*MsgClosechannel)(nil), "channel.channel.MsgClosechannel")
	proto.RegisterType((*MsgClosechannelResponse)(nil), "channel.channel.MsgClosechannelResponse")
}

func init() { proto.RegisterFile("channel/channel/tx.proto", fileDescriptor_89a552858311affe) }

var fileDescriptor_89a552858311affe = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xc1, 0x6a, 0xea, 0x40,
	0x14, 0x86, 0x9d, 0x78, 0xa3, 0xd7, 0xa3, 0x28, 0x0c, 0x97, 0xeb, 0x18, 0x2e, 0xb9, 0x92, 0x4d,
	0xb3, 0x9a, 0xa0, 0x7d, 0x02, 0xe3, 0x5a, 0x0a, 0x6e, 0x0a, 0xee, 0x62, 0x32, 0xa4, 0x01, 0xcd,
	0xa4, 0x99, 0x58, 0xec, 0x5b, 0xf4, 0x79, 0xfa, 0x04, 0x5d, 0xba, 0xec, 0xb2, 0xe8, 0x4b, 0x14,
	0xba, 0x29, 0x26, 0x99, 0x68, 0x44, 0xd0, 0xae, 0x86, 0xf3, 0x7b, 0x7e, 0xe7, 0x3f, 0x5f, 0xe6,
	0x00, 0x71, 0x1f, 0x9c, 0x30, 0x64, 0x0b, 0x4b, 0x9e, 0xc9, 0x9a, 0x46, 0x31, 0x4f, 0x38, 0xee,
	0xe4, 0x0a, 0xcd, 0x4f, 0x4d, 0x77, 0xb9, 0x58, 0x72, 0x61, 0xcd, 0x1d, 0xc1, 0xac, 0xa7, 0xc1,
	0x9c, 0x25, 0xce, 0xc0, 0x72, 0x79, 0x10, 0x66, 0x06, 0xe3, 0x13, 0x41, 0x7b, 0x22, 0xfc, 0xbb,
	0x88, 0x85, 0xb9, 0x05, 0x13, 0xa8, 0xbb, 0x31, 0x73, 0x12, 0x1e, 0x13, 0xd4, 0x47, 0x66, 0x63,
	0x2a, 0x4b, 0xfc, 0x07, 0xd4, 0xc8, 0x89, 0x93, 0x11, 0x51, 0x52, 0x3d, 0x2b, 0xa4, 0x6a, 0x93,
	0xea, 0x41, 0xb5, 0xb1, 0x05, 0xea, 0xfe, 0x9a, 0x11, 0xf9, 0xd5, 0x47, 0x66, 0x73, 0xd8, 0xa3,
	0x59, 0x10, 0xba, 0x0f, 0x42, 0xf3, 0x20, 0x74, 0xcc, 0x83, 0x70, 0x9a, 0xf5, 0x49, 0x83, 0x4d,
	0xd4, 0xab, 0x0c, 0x36, 0x36, 0xa0, 0xb5, 0x5c, 0x2d, 0x92, 0x40, 0x04, 0xfe, 0xc8, 0xf3, 0x62,
	0x52, 0x4b, 0xaf, 0x2f, 0x69, 0x58, 0x83, 0xdf, 0x82, 0x3d, 0xae, 0x58, 0xe8, 0x32, 0x52, 0x4f,
	0x7f, 0x2f, 0x6a, 0xc3, 0x84, 0xbf, 0xe5, 0xc9, 0xa7, 0x4c, 0x44, 0x3c, 0x14, 0x0c, 0xb7, 0x41,
	0x09, 0xbc, 0x7c, 0x78, 0x25, 0xf0, 0x8c, 0x2f, 0x04, 0x9d, 0x89, 0xf0, 0xc7, 0x0b, 0x2e, 0xd8,
	0x65, 0x4a, 0xa7, 0xb9, 0x94, 0x33, 0xb9, 0x0a, 0x92, 0xd5, 0x63, 0x92, 0x3f, 0x66, 0x56, 0xa0,
	0x57, 0xcf, 0xa0, 0xb7, 0x53, 0x22, 0xd7, 0x90, 0xfc, 0x07, 0x8d, 0x7c, 0xac, 0xc0, 0xcb, 0x31,
	0x1d, 0x04, 0xa3, 0x07, 0xdd, 0x93, 0xe1, 0x25, 0xa8, 0xe1, 0x2b, 0x82, 0xea, 0x44, 0xf8, 0xf8,
	0x1e, 0x9a, 0xc7, 0x2f, 0xe8, 0x3f, 0x3d, 0x79, 0x86, 0xb4, 0x0c, 0x5a, 0xbb, 0xb9, 0xd0, 0x50,
	0x7c, 0x89, 0x19, 0xb4, 0x4a, 0xd4, 0xfb, 0xe7, 0x8c, 0xc7, 0x1d, 0x9a, 0x79, 0xa9, 0x43, 0xfe,
	0xb7, 0x3d, 0x78, 0xdb, 0xea, 0x68, 0xb3, 0xd5, 0xd1, 0xc7, 0x56, 0x47, 0x2f, 0x3b, 0xbd, 0xb2,
	0xd9, 0xe9, 0x95, 0xf7, 0x9d, 0x5e, 0x99, 0x75, 0xe5, 0x5e, 0xad, 0x0f, 0x1b, 0xf6, 0x1c, 0x31,
	0x31, 0xaf, 0xa5, 0x4b, 0x73, 0xfb, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x22, 0x7e, 0x17, 0x47, 0x81,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	Openchannel(ctx context.Context, in *MsgOpenchannel, opts ...grpc.CallOption) (*MsgOpenchannelResponse, error)
	Closechannel(ctx context.Context, in *MsgClosechannel, opts ...grpc.CallOption) (*MsgClosechannelResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Openchannel(ctx context.Context, in *MsgOpenchannel, opts ...grpc.CallOption) (*MsgOpenchannelResponse, error) {
	out := new(MsgOpenchannelResponse)
	err := c.cc.Invoke(ctx, "/channel.channel.Msg/Openchannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Closechannel(ctx context.Context, in *MsgClosechannel, opts ...grpc.CallOption) (*MsgClosechannelResponse, error) {
	out := new(MsgClosechannelResponse)
	err := c.cc.Invoke(ctx, "/channel.channel.Msg/Closechannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	Openchannel(context.Context, *MsgOpenchannel) (*MsgOpenchannelResponse, error)
	Closechannel(context.Context, *MsgClosechannel) (*MsgClosechannelResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Openchannel(ctx context.Context, req *MsgOpenchannel) (*MsgOpenchannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Openchannel not implemented")
}
func (*UnimplementedMsgServer) Closechannel(ctx context.Context, req *MsgClosechannel) (*MsgClosechannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Closechannel not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Openchannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgOpenchannel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Openchannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/channel.channel.Msg/Openchannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Openchannel(ctx, req.(*MsgOpenchannel))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Closechannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgClosechannel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Closechannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/channel.channel.Msg/Closechannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Closechannel(ctx, req.(*MsgClosechannel))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "channel.channel.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Openchannel",
			Handler:    _Msg_Openchannel_Handler,
		},
		{
			MethodName: "Closechannel",
			Handler:    _Msg_Closechannel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "channel/channel/tx.proto",
}

func (m *MsgOpenchannel) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgOpenchannel) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgOpenchannel) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sequence) > 0 {
		i -= len(m.Sequence)
		copy(dAtA[i:], m.Sequence)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sequence)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.MultisigAddr) > 0 {
		i -= len(m.MultisigAddr)
		copy(dAtA[i:], m.MultisigAddr)
		i = encodeVarintTx(dAtA, i, uint64(len(m.MultisigAddr)))
		i--
		dAtA[i] = 0x32
	}
	if m.CoinB != nil {
		{
			size, err := m.CoinB.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.CoinA != nil {
		{
			size, err := m.CoinA.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.PartB) > 0 {
		i -= len(m.PartB)
		copy(dAtA[i:], m.PartB)
		i = encodeVarintTx(dAtA, i, uint64(len(m.PartB)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.PartA) > 0 {
		i -= len(m.PartA)
		copy(dAtA[i:], m.PartA)
		i = encodeVarintTx(dAtA, i, uint64(len(m.PartA)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgOpenchannelResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgOpenchannelResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgOpenchannelResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgClosechannel) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgClosechannel) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgClosechannel) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Channelid) > 0 {
		i -= len(m.Channelid)
		copy(dAtA[i:], m.Channelid)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Channelid)))
		i--
		dAtA[i] = 0x3a
	}
	if m.CoinB != nil {
		{
			size, err := m.CoinB.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if len(m.PartB) > 0 {
		i -= len(m.PartB)
		copy(dAtA[i:], m.PartB)
		i = encodeVarintTx(dAtA, i, uint64(len(m.PartB)))
		i--
		dAtA[i] = 0x2a
	}
	if m.CoinA != nil {
		{
			size, err := m.CoinA.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.PartA) > 0 {
		i -= len(m.PartA)
		copy(dAtA[i:], m.PartA)
		i = encodeVarintTx(dAtA, i, uint64(len(m.PartA)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.MultisigAddr) > 0 {
		i -= len(m.MultisigAddr)
		copy(dAtA[i:], m.MultisigAddr)
		i = encodeVarintTx(dAtA, i, uint64(len(m.MultisigAddr)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgClosechannelResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgClosechannelResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgClosechannelResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgOpenchannel) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.PartA)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.PartB)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.CoinA != nil {
		l = m.CoinA.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if m.CoinB != nil {
		l = m.CoinB.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.MultisigAddr)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Sequence)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgOpenchannelResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgClosechannel) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.MultisigAddr)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.PartA)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.CoinA != nil {
		l = m.CoinA.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.PartB)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.CoinB != nil {
		l = m.CoinB.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Channelid)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgClosechannelResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgOpenchannel) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgOpenchannel: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgOpenchannel: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PartA", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PartA = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PartB", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PartB = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoinA", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CoinA == nil {
				m.CoinA = &types.Coin{}
			}
			if err := m.CoinA.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoinB", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CoinB == nil {
				m.CoinB = &types.Coin{}
			}
			if err := m.CoinB.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MultisigAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MultisigAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sequence = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgOpenchannelResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgOpenchannelResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgOpenchannelResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgClosechannel) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgClosechannel: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgClosechannel: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MultisigAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MultisigAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PartA", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PartA = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoinA", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CoinA == nil {
				m.CoinA = &types.Coin{}
			}
			if err := m.CoinA.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PartB", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PartB = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoinB", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CoinB == nil {
				m.CoinB = &types.Coin{}
			}
			if err := m.CoinB.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Channelid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Channelid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgClosechannelResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgClosechannelResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgClosechannelResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)