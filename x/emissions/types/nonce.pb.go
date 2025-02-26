// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: emissions/v1/nonce.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type Nonce struct {
	BlockHeight int64 `protobuf:"varint,1,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
}

func (m *Nonce) Reset()         { *m = Nonce{} }
func (m *Nonce) String() string { return proto.CompactTextString(m) }
func (*Nonce) ProtoMessage()    {}
func (*Nonce) Descriptor() ([]byte, []int) {
	return fileDescriptor_529b41e523578c99, []int{0}
}
func (m *Nonce) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Nonce) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Nonce.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Nonce) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nonce.Merge(m, src)
}
func (m *Nonce) XXX_Size() int {
	return m.Size()
}
func (m *Nonce) XXX_DiscardUnknown() {
	xxx_messageInfo_Nonce.DiscardUnknown(m)
}

var xxx_messageInfo_Nonce proto.InternalMessageInfo

func (m *Nonce) GetBlockHeight() int64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

type Nonces struct {
	Nonces []*Nonce `protobuf:"bytes,1,rep,name=nonces,proto3" json:"nonces,omitempty"`
}

func (m *Nonces) Reset()         { *m = Nonces{} }
func (m *Nonces) String() string { return proto.CompactTextString(m) }
func (*Nonces) ProtoMessage()    {}
func (*Nonces) Descriptor() ([]byte, []int) {
	return fileDescriptor_529b41e523578c99, []int{1}
}
func (m *Nonces) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Nonces) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Nonces.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Nonces) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nonces.Merge(m, src)
}
func (m *Nonces) XXX_Size() int {
	return m.Size()
}
func (m *Nonces) XXX_DiscardUnknown() {
	xxx_messageInfo_Nonces.DiscardUnknown(m)
}

var xxx_messageInfo_Nonces proto.InternalMessageInfo

func (m *Nonces) GetNonces() []*Nonce {
	if m != nil {
		return m.Nonces
	}
	return nil
}

type ReputerRequestNonce struct {
	// associated with the reputer request
	ReputerNonce *Nonce `protobuf:"bytes,1,opt,name=reputer_nonce,json=reputerNonce,proto3" json:"reputer_nonce,omitempty"`
	// the reputers should respond to the reputer request with losses for work
	// found at this worker nonce
	WorkerNonce *Nonce `protobuf:"bytes,2,opt,name=worker_nonce,json=workerNonce,proto3" json:"worker_nonce,omitempty"`
}

func (m *ReputerRequestNonce) Reset()         { *m = ReputerRequestNonce{} }
func (m *ReputerRequestNonce) String() string { return proto.CompactTextString(m) }
func (*ReputerRequestNonce) ProtoMessage()    {}
func (*ReputerRequestNonce) Descriptor() ([]byte, []int) {
	return fileDescriptor_529b41e523578c99, []int{2}
}
func (m *ReputerRequestNonce) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReputerRequestNonce) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReputerRequestNonce.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReputerRequestNonce) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReputerRequestNonce.Merge(m, src)
}
func (m *ReputerRequestNonce) XXX_Size() int {
	return m.Size()
}
func (m *ReputerRequestNonce) XXX_DiscardUnknown() {
	xxx_messageInfo_ReputerRequestNonce.DiscardUnknown(m)
}

var xxx_messageInfo_ReputerRequestNonce proto.InternalMessageInfo

func (m *ReputerRequestNonce) GetReputerNonce() *Nonce {
	if m != nil {
		return m.ReputerNonce
	}
	return nil
}

func (m *ReputerRequestNonce) GetWorkerNonce() *Nonce {
	if m != nil {
		return m.WorkerNonce
	}
	return nil
}

type ReputerRequestNonces struct {
	Nonces []*ReputerRequestNonce `protobuf:"bytes,1,rep,name=nonces,proto3" json:"nonces,omitempty"`
}

func (m *ReputerRequestNonces) Reset()         { *m = ReputerRequestNonces{} }
func (m *ReputerRequestNonces) String() string { return proto.CompactTextString(m) }
func (*ReputerRequestNonces) ProtoMessage()    {}
func (*ReputerRequestNonces) Descriptor() ([]byte, []int) {
	return fileDescriptor_529b41e523578c99, []int{3}
}
func (m *ReputerRequestNonces) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReputerRequestNonces) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReputerRequestNonces.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReputerRequestNonces) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReputerRequestNonces.Merge(m, src)
}
func (m *ReputerRequestNonces) XXX_Size() int {
	return m.Size()
}
func (m *ReputerRequestNonces) XXX_DiscardUnknown() {
	xxx_messageInfo_ReputerRequestNonces.DiscardUnknown(m)
}

var xxx_messageInfo_ReputerRequestNonces proto.InternalMessageInfo

func (m *ReputerRequestNonces) GetNonces() []*ReputerRequestNonce {
	if m != nil {
		return m.Nonces
	}
	return nil
}

func init() {
	proto.RegisterType((*Nonce)(nil), "emissions.v1.Nonce")
	proto.RegisterType((*Nonces)(nil), "emissions.v1.Nonces")
	proto.RegisterType((*ReputerRequestNonce)(nil), "emissions.v1.ReputerRequestNonce")
	proto.RegisterType((*ReputerRequestNonces)(nil), "emissions.v1.ReputerRequestNonces")
}

func init() { proto.RegisterFile("emissions/v1/nonce.proto", fileDescriptor_529b41e523578c99) }

var fileDescriptor_529b41e523578c99 = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x48, 0xcd, 0xcd, 0x2c,
	0x2e, 0xce, 0xcc, 0xcf, 0x2b, 0xd6, 0x2f, 0x33, 0xd4, 0xcf, 0xcb, 0xcf, 0x4b, 0x4e, 0xd5, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x81, 0xcb, 0xe8, 0x95, 0x19, 0x4a, 0x89, 0xa4, 0xe7, 0xa7,
	0xe7, 0x83, 0x25, 0xf4, 0x41, 0x2c, 0x88, 0x1a, 0x25, 0x03, 0x2e, 0x56, 0x3f, 0x90, 0x16, 0x21,
	0x45, 0x2e, 0x9e, 0xa4, 0x9c, 0xfc, 0xe4, 0xec, 0xf8, 0x8c, 0xd4, 0xcc, 0xf4, 0x8c, 0x12, 0x09,
	0x46, 0x05, 0x46, 0x0d, 0xe6, 0x20, 0x6e, 0xb0, 0x98, 0x07, 0x58, 0xc8, 0x8a, 0xe5, 0xc5, 0x02,
	0x79, 0x46, 0x25, 0x53, 0x2e, 0x36, 0xb0, 0x8e, 0x62, 0x21, 0x6d, 0x2e, 0x36, 0xb0, 0x75, 0xc5,
	0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0xc2, 0x7a, 0xc8, 0x16, 0xea, 0x81, 0x55, 0x05, 0x41,
	0x95, 0x28, 0xf5, 0x32, 0x72, 0x09, 0x07, 0xa5, 0x16, 0x94, 0x96, 0xa4, 0x16, 0x05, 0xa5, 0x16,
	0x96, 0xa6, 0x16, 0x97, 0x40, 0xec, 0xb5, 0xe0, 0xe2, 0x2d, 0x82, 0x08, 0xc7, 0x83, 0x55, 0x82,
	0x2d, 0xc6, 0x61, 0x16, 0x0f, 0x54, 0x25, 0x44, 0xa7, 0x19, 0x17, 0x4f, 0x79, 0x7e, 0x51, 0x36,
	0x5c, 0x23, 0x13, 0x6e, 0x8d, 0xdc, 0x10, 0x85, 0x60, 0x0e, 0xd4, 0x1b, 0x81, 0x5c, 0x22, 0x58,
	0x9c, 0x53, 0x2c, 0x64, 0x89, 0xe6, 0x29, 0x45, 0x54, 0xf3, 0xb0, 0xe8, 0x81, 0x79, 0xd1, 0x29,
	0xe8, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58,
	0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x2c, 0xd2, 0x33, 0x4b, 0x32,
	0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x13, 0x73, 0x72, 0xf2, 0x8b, 0x12, 0x75, 0xf3, 0x52,
	0x4b, 0x40, 0x8e, 0x82, 0x71, 0x93, 0x33, 0x12, 0x33, 0xf3, 0xf4, 0x2b, 0xf4, 0x11, 0x91, 0x59,
	0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0x8e, 0x26, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xe5, 0xad, 0xbc, 0xfe, 0xe6, 0x01, 0x00, 0x00,
}

func (this *Nonce) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Nonce)
	if !ok {
		that2, ok := that.(Nonce)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.BlockHeight != that1.BlockHeight {
		return false
	}
	return true
}
func (this *ReputerRequestNonce) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ReputerRequestNonce)
	if !ok {
		that2, ok := that.(ReputerRequestNonce)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.ReputerNonce.Equal(that1.ReputerNonce) {
		return false
	}
	if !this.WorkerNonce.Equal(that1.WorkerNonce) {
		return false
	}
	return true
}
func (m *Nonce) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Nonce) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Nonce) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlockHeight != 0 {
		i = encodeVarintNonce(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Nonces) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Nonces) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Nonces) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Nonces) > 0 {
		for iNdEx := len(m.Nonces) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Nonces[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintNonce(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ReputerRequestNonce) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReputerRequestNonce) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReputerRequestNonce) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.WorkerNonce != nil {
		{
			size, err := m.WorkerNonce.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintNonce(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.ReputerNonce != nil {
		{
			size, err := m.ReputerNonce.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintNonce(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ReputerRequestNonces) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReputerRequestNonces) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReputerRequestNonces) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Nonces) > 0 {
		for iNdEx := len(m.Nonces) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Nonces[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintNonce(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintNonce(dAtA []byte, offset int, v uint64) int {
	offset -= sovNonce(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Nonce) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlockHeight != 0 {
		n += 1 + sovNonce(uint64(m.BlockHeight))
	}
	return n
}

func (m *Nonces) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Nonces) > 0 {
		for _, e := range m.Nonces {
			l = e.Size()
			n += 1 + l + sovNonce(uint64(l))
		}
	}
	return n
}

func (m *ReputerRequestNonce) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ReputerNonce != nil {
		l = m.ReputerNonce.Size()
		n += 1 + l + sovNonce(uint64(l))
	}
	if m.WorkerNonce != nil {
		l = m.WorkerNonce.Size()
		n += 1 + l + sovNonce(uint64(l))
	}
	return n
}

func (m *ReputerRequestNonces) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Nonces) > 0 {
		for _, e := range m.Nonces {
			l = e.Size()
			n += 1 + l + sovNonce(uint64(l))
		}
	}
	return n
}

func sovNonce(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNonce(x uint64) (n int) {
	return sovNonce(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Nonce) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNonce
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
			return fmt.Errorf("proto: Nonce: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Nonce: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNonce
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNonce(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNonce
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
func (m *Nonces) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNonce
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
			return fmt.Errorf("proto: Nonces: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Nonces: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonces", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNonce
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
				return ErrInvalidLengthNonce
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNonce
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nonces = append(m.Nonces, &Nonce{})
			if err := m.Nonces[len(m.Nonces)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNonce(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNonce
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
func (m *ReputerRequestNonce) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNonce
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
			return fmt.Errorf("proto: ReputerRequestNonce: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReputerRequestNonce: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReputerNonce", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNonce
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
				return ErrInvalidLengthNonce
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNonce
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ReputerNonce == nil {
				m.ReputerNonce = &Nonce{}
			}
			if err := m.ReputerNonce.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WorkerNonce", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNonce
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
				return ErrInvalidLengthNonce
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNonce
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.WorkerNonce == nil {
				m.WorkerNonce = &Nonce{}
			}
			if err := m.WorkerNonce.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNonce(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNonce
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
func (m *ReputerRequestNonces) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNonce
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
			return fmt.Errorf("proto: ReputerRequestNonces: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReputerRequestNonces: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonces", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNonce
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
				return ErrInvalidLengthNonce
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNonce
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nonces = append(m.Nonces, &ReputerRequestNonce{})
			if err := m.Nonces[len(m.Nonces)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNonce(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNonce
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
func skipNonce(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNonce
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
					return 0, ErrIntOverflowNonce
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
					return 0, ErrIntOverflowNonce
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
				return 0, ErrInvalidLengthNonce
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNonce
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNonce
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNonce        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNonce          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNonce = fmt.Errorf("proto: unexpected end of group")
)
