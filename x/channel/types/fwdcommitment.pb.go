// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: channel/channel/fwdcommitment.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	proto "github.com/gogo/protobuf/proto"
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

type Fwdcommitment struct {
	Index            string      `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Channelid        string      `protobuf:"bytes,2,opt,name=channelid,proto3" json:"channelid,omitempty"`
	MultisigAddr     string      `protobuf:"bytes,3,opt,name=multisigAddr,proto3" json:"multisigAddr,omitempty"`
	SenderAddr       string      `protobuf:"bytes,4,opt,name=senderAddr,proto3" json:"senderAddr,omitempty"`
	ReceiverAddr     string      `protobuf:"bytes,5,opt,name=receiverAddr,proto3" json:"receiverAddr,omitempty"`
	Timelocksender   uint64      `protobuf:"varint,6,opt,name=timelocksender,proto3" json:"timelocksender,omitempty"`
	Timelockreceiver uint64      `protobuf:"varint,7,opt,name=timelockreceiver,proto3" json:"timelockreceiver,omitempty"`
	Hashcodehtlc     string      `protobuf:"bytes,8,opt,name=hashcodehtlc,proto3" json:"hashcodehtlc,omitempty"`
	Hashcodedest     string      `protobuf:"bytes,9,opt,name=hashcodedest,proto3" json:"hashcodedest,omitempty"`
	Cointransfer     *types.Coin `protobuf:"bytes,10,opt,name=cointransfer,proto3" json:"cointransfer,omitempty"`
	Creator          string      `protobuf:"bytes,11,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *Fwdcommitment) Reset()         { *m = Fwdcommitment{} }
func (m *Fwdcommitment) String() string { return proto.CompactTextString(m) }
func (*Fwdcommitment) ProtoMessage()    {}
func (*Fwdcommitment) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fe370908b866a42, []int{0}
}
func (m *Fwdcommitment) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Fwdcommitment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Fwdcommitment.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Fwdcommitment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fwdcommitment.Merge(m, src)
}
func (m *Fwdcommitment) XXX_Size() int {
	return m.Size()
}
func (m *Fwdcommitment) XXX_DiscardUnknown() {
	xxx_messageInfo_Fwdcommitment.DiscardUnknown(m)
}

var xxx_messageInfo_Fwdcommitment proto.InternalMessageInfo

func (m *Fwdcommitment) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *Fwdcommitment) GetChannelid() string {
	if m != nil {
		return m.Channelid
	}
	return ""
}

func (m *Fwdcommitment) GetMultisigAddr() string {
	if m != nil {
		return m.MultisigAddr
	}
	return ""
}

func (m *Fwdcommitment) GetSenderAddr() string {
	if m != nil {
		return m.SenderAddr
	}
	return ""
}

func (m *Fwdcommitment) GetReceiverAddr() string {
	if m != nil {
		return m.ReceiverAddr
	}
	return ""
}

func (m *Fwdcommitment) GetTimelocksender() uint64 {
	if m != nil {
		return m.Timelocksender
	}
	return 0
}

func (m *Fwdcommitment) GetTimelockreceiver() uint64 {
	if m != nil {
		return m.Timelockreceiver
	}
	return 0
}

func (m *Fwdcommitment) GetHashcodehtlc() string {
	if m != nil {
		return m.Hashcodehtlc
	}
	return ""
}

func (m *Fwdcommitment) GetHashcodedest() string {
	if m != nil {
		return m.Hashcodedest
	}
	return ""
}

func (m *Fwdcommitment) GetCointransfer() *types.Coin {
	if m != nil {
		return m.Cointransfer
	}
	return nil
}

func (m *Fwdcommitment) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*Fwdcommitment)(nil), "channel.channel.Fwdcommitment")
}

func init() {
	proto.RegisterFile("channel/channel/fwdcommitment.proto", fileDescriptor_2fe370908b866a42)
}

var fileDescriptor_2fe370908b866a42 = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xcb, 0x4a, 0x33, 0x31,
	0x18, 0x86, 0x9b, 0xbf, 0xa7, 0xbf, 0x69, 0x3d, 0x10, 0x04, 0xa3, 0x48, 0x28, 0x15, 0xa4, 0xb8,
	0x98, 0xa1, 0xba, 0x76, 0xa1, 0x82, 0x17, 0xd0, 0xa5, 0xbb, 0x69, 0xf2, 0xd5, 0x09, 0xce, 0x24,
	0x25, 0x89, 0xb5, 0xde, 0x85, 0x97, 0xe0, 0xe5, 0xb8, 0xec, 0xd2, 0xa5, 0xb4, 0x37, 0x22, 0x93,
	0x99, 0xd1, 0x19, 0x5d, 0x85, 0x3c, 0x79, 0xbe, 0x37, 0x90, 0x37, 0xf8, 0x94, 0xc7, 0x91, 0x52,
	0x90, 0x84, 0xe5, 0x3a, 0x7f, 0x16, 0x5c, 0xa7, 0xa9, 0x74, 0x29, 0x28, 0x17, 0x2c, 0x8c, 0x76,
	0x9a, 0xec, 0x15, 0x87, 0x41, 0xb1, 0x1e, 0x33, 0xae, 0x6d, 0xaa, 0x6d, 0x38, 0x8b, 0x2c, 0x84,
	0xcb, 0xc9, 0x0c, 0x5c, 0x34, 0x09, 0xb9, 0x96, 0x2a, 0x1f, 0x18, 0xbd, 0x35, 0xf1, 0xce, 0x5d,
	0x35, 0x88, 0x1c, 0xe0, 0xb6, 0x54, 0x02, 0x56, 0x14, 0x0d, 0xd1, 0xb8, 0x37, 0xcd, 0x37, 0xe4,
	0x04, 0xf7, 0x8a, 0x48, 0x29, 0xe8, 0x3f, 0x7f, 0xf2, 0x03, 0xc8, 0x08, 0x0f, 0xd2, 0xa7, 0xc4,
	0x49, 0x2b, 0x1f, 0xae, 0x85, 0x30, 0xb4, 0xe9, 0x85, 0x1a, 0x23, 0x0c, 0x63, 0x0b, 0x4a, 0x80,
	0xf1, 0x46, 0xcb, 0x1b, 0x15, 0x92, 0x65, 0x18, 0xe0, 0x20, 0x97, 0x85, 0xd1, 0xce, 0x33, 0xaa,
	0x8c, 0x9c, 0xe1, 0x5d, 0x27, 0x53, 0x48, 0x34, 0x7f, 0xcc, 0x27, 0x69, 0x67, 0x88, 0xc6, 0xad,
	0xe9, 0x2f, 0x4a, 0xce, 0xf1, 0x7e, 0x49, 0xca, 0x79, 0xda, 0xf5, 0xe6, 0x1f, 0x9e, 0xdd, 0x1b,
	0x47, 0x36, 0xe6, 0x5a, 0x40, 0xec, 0x12, 0x4e, 0xff, 0xe7, 0xf7, 0x56, 0x59, 0xd5, 0x11, 0x60,
	0x1d, 0xed, 0xd5, 0x9d, 0x8c, 0x91, 0x2b, 0x3c, 0xc8, 0xde, 0xd5, 0x99, 0x48, 0xd9, 0x39, 0x18,
	0x8a, 0x87, 0x68, 0xdc, 0xbf, 0x38, 0x0a, 0xf2, 0x02, 0x82, 0xac, 0x80, 0xa0, 0x28, 0x20, 0xb8,
	0xd5, 0x52, 0x4d, 0x6b, 0x3a, 0xa1, 0xb8, 0xcb, 0x0d, 0x44, 0x4e, 0x1b, 0xda, 0xf7, 0xe9, 0xe5,
	0xf6, 0x66, 0xf2, 0xbe, 0x61, 0x68, 0xbd, 0x61, 0xe8, 0x73, 0xc3, 0xd0, 0xeb, 0x96, 0x35, 0xd6,
	0x5b, 0xd6, 0xf8, 0xd8, 0xb2, 0xc6, 0xfd, 0x61, 0xf9, 0x15, 0x56, 0xdf, 0x9f, 0xc2, 0xbd, 0x2c,
	0xc0, 0xce, 0x3a, 0xbe, 0xdc, 0xcb, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa9, 0xfe, 0x43, 0x5c,
	0x34, 0x02, 0x00, 0x00,
}

func (m *Fwdcommitment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Fwdcommitment) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Fwdcommitment) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintFwdcommitment(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x5a
	}
	if m.Cointransfer != nil {
		{
			size, err := m.Cointransfer.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFwdcommitment(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x52
	}
	if len(m.Hashcodedest) > 0 {
		i -= len(m.Hashcodedest)
		copy(dAtA[i:], m.Hashcodedest)
		i = encodeVarintFwdcommitment(dAtA, i, uint64(len(m.Hashcodedest)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Hashcodehtlc) > 0 {
		i -= len(m.Hashcodehtlc)
		copy(dAtA[i:], m.Hashcodehtlc)
		i = encodeVarintFwdcommitment(dAtA, i, uint64(len(m.Hashcodehtlc)))
		i--
		dAtA[i] = 0x42
	}
	if m.Timelockreceiver != 0 {
		i = encodeVarintFwdcommitment(dAtA, i, uint64(m.Timelockreceiver))
		i--
		dAtA[i] = 0x38
	}
	if m.Timelocksender != 0 {
		i = encodeVarintFwdcommitment(dAtA, i, uint64(m.Timelocksender))
		i--
		dAtA[i] = 0x30
	}
	if len(m.ReceiverAddr) > 0 {
		i -= len(m.ReceiverAddr)
		copy(dAtA[i:], m.ReceiverAddr)
		i = encodeVarintFwdcommitment(dAtA, i, uint64(len(m.ReceiverAddr)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.SenderAddr) > 0 {
		i -= len(m.SenderAddr)
		copy(dAtA[i:], m.SenderAddr)
		i = encodeVarintFwdcommitment(dAtA, i, uint64(len(m.SenderAddr)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.MultisigAddr) > 0 {
		i -= len(m.MultisigAddr)
		copy(dAtA[i:], m.MultisigAddr)
		i = encodeVarintFwdcommitment(dAtA, i, uint64(len(m.MultisigAddr)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Channelid) > 0 {
		i -= len(m.Channelid)
		copy(dAtA[i:], m.Channelid)
		i = encodeVarintFwdcommitment(dAtA, i, uint64(len(m.Channelid)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintFwdcommitment(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintFwdcommitment(dAtA []byte, offset int, v uint64) int {
	offset -= sovFwdcommitment(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Fwdcommitment) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovFwdcommitment(uint64(l))
	}
	l = len(m.Channelid)
	if l > 0 {
		n += 1 + l + sovFwdcommitment(uint64(l))
	}
	l = len(m.MultisigAddr)
	if l > 0 {
		n += 1 + l + sovFwdcommitment(uint64(l))
	}
	l = len(m.SenderAddr)
	if l > 0 {
		n += 1 + l + sovFwdcommitment(uint64(l))
	}
	l = len(m.ReceiverAddr)
	if l > 0 {
		n += 1 + l + sovFwdcommitment(uint64(l))
	}
	if m.Timelocksender != 0 {
		n += 1 + sovFwdcommitment(uint64(m.Timelocksender))
	}
	if m.Timelockreceiver != 0 {
		n += 1 + sovFwdcommitment(uint64(m.Timelockreceiver))
	}
	l = len(m.Hashcodehtlc)
	if l > 0 {
		n += 1 + l + sovFwdcommitment(uint64(l))
	}
	l = len(m.Hashcodedest)
	if l > 0 {
		n += 1 + l + sovFwdcommitment(uint64(l))
	}
	if m.Cointransfer != nil {
		l = m.Cointransfer.Size()
		n += 1 + l + sovFwdcommitment(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovFwdcommitment(uint64(l))
	}
	return n
}

func sovFwdcommitment(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFwdcommitment(x uint64) (n int) {
	return sovFwdcommitment(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Fwdcommitment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFwdcommitment
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
			return fmt.Errorf("proto: Fwdcommitment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Fwdcommitment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFwdcommitment
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
				return ErrInvalidLengthFwdcommitment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFwdcommitment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Channelid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFwdcommitment
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
				return ErrInvalidLengthFwdcommitment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFwdcommitment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Channelid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MultisigAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFwdcommitment
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
				return ErrInvalidLengthFwdcommitment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFwdcommitment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MultisigAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SenderAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFwdcommitment
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
				return ErrInvalidLengthFwdcommitment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFwdcommitment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SenderAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReceiverAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFwdcommitment
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
				return ErrInvalidLengthFwdcommitment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFwdcommitment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReceiverAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timelocksender", wireType)
			}
			m.Timelocksender = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFwdcommitment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timelocksender |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timelockreceiver", wireType)
			}
			m.Timelockreceiver = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFwdcommitment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timelockreceiver |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hashcodehtlc", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFwdcommitment
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
				return ErrInvalidLengthFwdcommitment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFwdcommitment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hashcodehtlc = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hashcodedest", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFwdcommitment
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
				return ErrInvalidLengthFwdcommitment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFwdcommitment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hashcodedest = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cointransfer", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFwdcommitment
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
				return ErrInvalidLengthFwdcommitment
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFwdcommitment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Cointransfer == nil {
				m.Cointransfer = &types.Coin{}
			}
			if err := m.Cointransfer.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFwdcommitment
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
				return ErrInvalidLengthFwdcommitment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFwdcommitment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFwdcommitment(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFwdcommitment
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
func skipFwdcommitment(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFwdcommitment
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
					return 0, ErrIntOverflowFwdcommitment
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
					return 0, ErrIntOverflowFwdcommitment
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
				return 0, ErrInvalidLengthFwdcommitment
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFwdcommitment
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFwdcommitment
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFwdcommitment        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFwdcommitment          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFwdcommitment = fmt.Errorf("proto: unexpected end of group")
)
