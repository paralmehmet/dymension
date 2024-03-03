// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dymension/sequencer/sequencer.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types1 "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Sequencer defines a sequencer identified by its' address (sequencerAddress).
// The sequencer could be attached to only one rollapp (rollappId).
type Sequencer struct {
	// sequencerAddress is the bech32-encoded address of the sequencer account which is the account that the message was sent from.
	SequencerAddress string `protobuf:"bytes,1,opt,name=sequencerAddress,proto3" json:"sequencerAddress,omitempty"`
	// pubkey is the public key of the sequencers' dymint client, as a Protobuf Any.
	DymintPubKey *types.Any `protobuf:"bytes,2,opt,name=dymintPubKey,proto3" json:"dymintPubKey,omitempty"`
	// rollappId defines the rollapp to which the sequencer belongs.
	RollappId string `protobuf:"bytes,3,opt,name=rollappId,proto3" json:"rollappId,omitempty"`
	// description defines the descriptive terms for the sequencer.
	Description Description `protobuf:"bytes,4,opt,name=description,proto3" json:"description"`
	// jailed defined whether the sequencer has been jailed from bonded status or not.
	Jailed bool `protobuf:"varint,5,opt,name=jailed,proto3" json:"jailed,omitempty"`
	// status is the sequencer status (bonded/unbonding/unbonded).
	Status OperatingStatus `protobuf:"varint,6,opt,name=status,proto3,enum=dymensionxyz.dymension.sequencer.OperatingStatus" json:"status,omitempty"`
	// tokens define the delegated tokens (incl. self-delegation).
	Tokens github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,7,rep,name=tokens,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"tokens"`
	// unbonding_height defines, if unbonding, the height at which this sequencer has begun unbonding.
	UnbondingHeight int64 `protobuf:"varint,8,opt,name=unbonding_height,json=unbondingHeight,proto3" json:"unbonding_height,omitempty"`
	// unbond_time defines, if unbonding, the min time for the sequencer to complete unbonding.
	UnbondTime time.Time `protobuf:"bytes,9,opt,name=unbond_time,json=unbondTime,proto3,stdtime" json:"unbond_time"`
}

func (m *Sequencer) Reset()         { *m = Sequencer{} }
func (m *Sequencer) String() string { return proto.CompactTextString(m) }
func (*Sequencer) ProtoMessage()    {}
func (*Sequencer) Descriptor() ([]byte, []int) {
	return fileDescriptor_17d99b644bf09274, []int{0}
}
func (m *Sequencer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Sequencer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Sequencer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Sequencer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sequencer.Merge(m, src)
}
func (m *Sequencer) XXX_Size() int {
	return m.Size()
}
func (m *Sequencer) XXX_DiscardUnknown() {
	xxx_messageInfo_Sequencer.DiscardUnknown(m)
}

var xxx_messageInfo_Sequencer proto.InternalMessageInfo

func (m *Sequencer) GetSequencerAddress() string {
	if m != nil {
		return m.SequencerAddress
	}
	return ""
}

func (m *Sequencer) GetDymintPubKey() *types.Any {
	if m != nil {
		return m.DymintPubKey
	}
	return nil
}

func (m *Sequencer) GetRollappId() string {
	if m != nil {
		return m.RollappId
	}
	return ""
}

func (m *Sequencer) GetDescription() Description {
	if m != nil {
		return m.Description
	}
	return Description{}
}

func (m *Sequencer) GetJailed() bool {
	if m != nil {
		return m.Jailed
	}
	return false
}

func (m *Sequencer) GetStatus() OperatingStatus {
	if m != nil {
		return m.Status
	}
	return Unspecified
}

func (m *Sequencer) GetTokens() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Tokens
	}
	return nil
}

func (m *Sequencer) GetUnbondingHeight() int64 {
	if m != nil {
		return m.UnbondingHeight
	}
	return 0
}

func (m *Sequencer) GetUnbondTime() time.Time {
	if m != nil {
		return m.UnbondTime
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*Sequencer)(nil), "dymensionxyz.dymension.sequencer.Sequencer")
}

func init() {
	proto.RegisterFile("dymension/sequencer/sequencer.proto", fileDescriptor_17d99b644bf09274)
}

var fileDescriptor_17d99b644bf09274 = []byte{
	// 533 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x4d, 0x6f, 0xd3, 0x30,
	0x18, 0x6e, 0x68, 0x29, 0xad, 0x8b, 0x60, 0x8a, 0x2a, 0xf0, 0x2a, 0x94, 0x46, 0x20, 0xa4, 0x30,
	0xa9, 0x36, 0xed, 0x24, 0xee, 0x2b, 0x20, 0x31, 0x71, 0x60, 0xca, 0xe0, 0xc2, 0xa5, 0xca, 0x87,
	0x49, 0xcd, 0x1a, 0x3b, 0xc4, 0x4e, 0xb5, 0xf0, 0x2b, 0xf6, 0x3b, 0x38, 0x73, 0xe2, 0x17, 0x4c,
	0x9c, 0x76, 0xe4, 0xc4, 0x50, 0xfb, 0x47, 0x50, 0x1c, 0xa7, 0x0d, 0x6c, 0xd2, 0x4e, 0xf1, 0xfb,
	0xf1, 0x3c, 0xef, 0xc7, 0xf3, 0x06, 0x3c, 0x09, 0xf3, 0x98, 0x30, 0x41, 0x39, 0xc3, 0x82, 0x7c,
	0xc9, 0x08, 0x0b, 0x48, 0xba, 0x7d, 0xa1, 0x24, 0xe5, 0x92, 0x9b, 0xf6, 0x26, 0xe9, 0x34, 0xff,
	0x8a, 0x36, 0x06, 0xda, 0xe4, 0x0d, 0x9e, 0x5e, 0x47, 0x13, 0x12, 0x11, 0xa4, 0x34, 0x91, 0x45,
	0xaa, 0x22, 0x1a, 0xec, 0x5d, 0x97, 0xc6, 0x13, 0x92, 0x7a, 0x92, 0xb2, 0x68, 0x26, 0xa4, 0x27,
	0x33, 0xa1, 0x73, 0x77, 0x03, 0x2e, 0x62, 0x2e, 0x66, 0xca, 0xc2, 0xa5, 0x51, 0x85, 0x22, 0xce,
	0xa3, 0x05, 0xc1, 0xca, 0xf2, 0xb3, 0x4f, 0xd8, 0x63, 0xb9, 0x0e, 0xf5, 0x23, 0x1e, 0xf1, 0x12,
	0x52, 0xbc, 0xb4, 0x77, 0xf8, 0x3f, 0x40, 0xd2, 0x98, 0x08, 0xe9, 0xc5, 0x89, 0x4e, 0xb0, 0x4a,
	0x7e, 0xec, 0x7b, 0x82, 0xe0, 0xe5, 0xd8, 0x27, 0xd2, 0x1b, 0xe3, 0x80, 0xd3, 0xaa, 0xf1, 0x87,
	0x3a, 0x1e, 0x8b, 0x08, 0x2f, 0xc7, 0xc5, 0xa7, 0x0c, 0x3c, 0xfe, 0xd1, 0x02, 0xdd, 0xe3, 0x6a,
	0x14, 0x73, 0x0f, 0xec, 0x6c, 0xe6, 0x3a, 0x08, 0xc3, 0x94, 0x08, 0x01, 0x0d, 0xdb, 0x70, 0xba,
	0xee, 0x15, 0xbf, 0xe9, 0x82, 0xbb, 0x61, 0x1e, 0x53, 0x26, 0x8f, 0x32, 0xff, 0x2d, 0xc9, 0xe1,
	0x2d, 0xdb, 0x70, 0x7a, 0x93, 0x3e, 0x2a, 0x5b, 0x45, 0x55, 0xab, 0xe8, 0x80, 0xe5, 0x53, 0xf8,
	0xf3, 0xfb, 0xa8, 0xaf, 0x57, 0x10, 0xa4, 0x79, 0x22, 0x39, 0x2a, 0x51, 0xee, 0x3f, 0x1c, 0xe6,
	0x23, 0xd0, 0x4d, 0xf9, 0x62, 0xe1, 0x25, 0xc9, 0x61, 0x08, 0x9b, 0xaa, 0xf0, 0xd6, 0x61, 0x7e,
	0x00, 0xbd, 0x9a, 0x24, 0xb0, 0xa5, 0x0a, 0x8e, 0xd0, 0x4d, 0xe2, 0xa2, 0x57, 0x5b, 0xd0, 0xb4,
	0x75, 0xfe, 0x7b, 0xd8, 0x70, 0xeb, 0x3c, 0xe6, 0x03, 0xd0, 0xfe, 0xec, 0xd1, 0x05, 0x09, 0xe1,
	0x6d, 0xdb, 0x70, 0x3a, 0xae, 0xb6, 0xcc, 0x43, 0xd0, 0x2e, 0x05, 0x85, 0x6d, 0xdb, 0x70, 0xee,
	0x4d, 0xc6, 0x37, 0x57, 0x7a, 0x57, 0x9d, 0xc2, 0xb1, 0x02, 0xba, 0x9a, 0xc0, 0x0c, 0x40, 0x5b,
	0xf2, 0x13, 0xc2, 0x04, 0xbc, 0x63, 0x37, 0x9d, 0xde, 0x64, 0x17, 0xe9, 0x65, 0x14, 0x7a, 0x21,
	0xad, 0x17, 0x7a, 0xc9, 0x29, 0x9b, 0x3e, 0x2f, 0x1a, 0xfc, 0x76, 0x39, 0x74, 0x22, 0x2a, 0xe7,
	0x99, 0x8f, 0x02, 0x1e, 0xeb, 0xe3, 0xd1, 0x9f, 0x91, 0x08, 0x4f, 0xb0, 0xcc, 0x13, 0x22, 0x14,
	0x40, 0xb8, 0x9a, 0xda, 0x7c, 0x06, 0x76, 0x32, 0xe6, 0x73, 0x16, 0x16, 0xa7, 0x38, 0x27, 0x34,
	0x9a, 0x4b, 0xd8, 0xb1, 0x0d, 0xa7, 0xe9, 0xde, 0xdf, 0xf8, 0xdf, 0x28, 0xb7, 0xf9, 0x1a, 0xf4,
	0x4a, 0xd7, 0xac, 0x38, 0x24, 0xd8, 0x55, 0x9b, 0x1c, 0x5c, 0x91, 0xee, 0x7d, 0x75, 0x65, 0xd3,
	0x4e, 0xd1, 0xd5, 0xd9, 0xe5, 0xd0, 0x70, 0x41, 0x09, 0x2c, 0x42, 0xd3, 0xa3, 0xf3, 0x95, 0x65,
	0x5c, 0xac, 0x2c, 0xe3, 0xcf, 0xca, 0x32, 0xce, 0xd6, 0x56, 0xe3, 0x62, 0x6d, 0x35, 0x7e, 0xad,
	0xad, 0xc6, 0xc7, 0x17, 0xb5, 0xee, 0xeb, 0x5b, 0xdb, 0x1a, 0x78, 0xb9, 0x8f, 0x4f, 0x6b, 0x7f,
	0x91, 0x9a, 0xc8, 0x6f, 0xab, 0xda, 0xfb, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x4a, 0x10,
	0xfd, 0xd7, 0x03, 0x00, 0x00,
}

func (m *Sequencer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Sequencer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Sequencer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.UnbondTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.UnbondTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintSequencer(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x4a
	if m.UnbondingHeight != 0 {
		i = encodeVarintSequencer(dAtA, i, uint64(m.UnbondingHeight))
		i--
		dAtA[i] = 0x40
	}
	if len(m.Tokens) > 0 {
		for iNdEx := len(m.Tokens) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Tokens[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSequencer(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.Status != 0 {
		i = encodeVarintSequencer(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x30
	}
	if m.Jailed {
		i--
		if m.Jailed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	{
		size, err := m.Description.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintSequencer(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.RollappId) > 0 {
		i -= len(m.RollappId)
		copy(dAtA[i:], m.RollappId)
		i = encodeVarintSequencer(dAtA, i, uint64(len(m.RollappId)))
		i--
		dAtA[i] = 0x1a
	}
	if m.DymintPubKey != nil {
		{
			size, err := m.DymintPubKey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSequencer(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.SequencerAddress) > 0 {
		i -= len(m.SequencerAddress)
		copy(dAtA[i:], m.SequencerAddress)
		i = encodeVarintSequencer(dAtA, i, uint64(len(m.SequencerAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSequencer(dAtA []byte, offset int, v uint64) int {
	offset -= sovSequencer(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Sequencer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SequencerAddress)
	if l > 0 {
		n += 1 + l + sovSequencer(uint64(l))
	}
	if m.DymintPubKey != nil {
		l = m.DymintPubKey.Size()
		n += 1 + l + sovSequencer(uint64(l))
	}
	l = len(m.RollappId)
	if l > 0 {
		n += 1 + l + sovSequencer(uint64(l))
	}
	l = m.Description.Size()
	n += 1 + l + sovSequencer(uint64(l))
	if m.Jailed {
		n += 2
	}
	if m.Status != 0 {
		n += 1 + sovSequencer(uint64(m.Status))
	}
	if len(m.Tokens) > 0 {
		for _, e := range m.Tokens {
			l = e.Size()
			n += 1 + l + sovSequencer(uint64(l))
		}
	}
	if m.UnbondingHeight != 0 {
		n += 1 + sovSequencer(uint64(m.UnbondingHeight))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.UnbondTime)
	n += 1 + l + sovSequencer(uint64(l))
	return n
}

func sovSequencer(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSequencer(x uint64) (n int) {
	return sovSequencer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Sequencer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSequencer
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
			return fmt.Errorf("proto: Sequencer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Sequencer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SequencerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSequencer
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
				return ErrInvalidLengthSequencer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSequencer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SequencerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DymintPubKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSequencer
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
				return ErrInvalidLengthSequencer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSequencer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DymintPubKey == nil {
				m.DymintPubKey = &types.Any{}
			}
			if err := m.DymintPubKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RollappId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSequencer
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
				return ErrInvalidLengthSequencer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSequencer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RollappId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSequencer
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
				return ErrInvalidLengthSequencer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSequencer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Description.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Jailed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSequencer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Jailed = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSequencer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= OperatingStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tokens", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSequencer
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
				return ErrInvalidLengthSequencer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSequencer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tokens = append(m.Tokens, types1.Coin{})
			if err := m.Tokens[len(m.Tokens)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnbondingHeight", wireType)
			}
			m.UnbondingHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSequencer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UnbondingHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnbondTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSequencer
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
				return ErrInvalidLengthSequencer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSequencer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.UnbondTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSequencer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSequencer
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
func skipSequencer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSequencer
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
					return 0, ErrIntOverflowSequencer
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
					return 0, ErrIntOverflowSequencer
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
				return 0, ErrInvalidLengthSequencer
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSequencer
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSequencer
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSequencer        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSequencer          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSequencer = fmt.Errorf("proto: unexpected end of group")
)
