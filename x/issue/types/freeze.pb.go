// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: konstellation/issue/freeze.proto

package types

import (
	fmt "fmt"
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

type Freeze struct {
	In  bool `protobuf:"varint,1,opt,name=in,proto3" json:"in,omitempty"`
	Out bool `protobuf:"varint,2,opt,name=out,proto3" json:"out,omitempty"`
}

func (m *Freeze) Reset()         { *m = Freeze{} }
func (m *Freeze) String() string { return proto.CompactTextString(m) }
func (*Freeze) ProtoMessage()    {}
func (*Freeze) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6927cdb03dd0514, []int{0}
}
func (m *Freeze) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Freeze) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Freeze.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Freeze) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Freeze.Merge(m, src)
}
func (m *Freeze) XXX_Size() int {
	return m.Size()
}
func (m *Freeze) XXX_DiscardUnknown() {
	xxx_messageInfo_Freeze.DiscardUnknown(m)
}

var xxx_messageInfo_Freeze proto.InternalMessageInfo

func (m *Freeze) GetIn() bool {
	if m != nil {
		return m.In
	}
	return false
}

func (m *Freeze) GetOut() bool {
	if m != nil {
		return m.Out
	}
	return false
}

type AddressFreezeList struct {
	AddressFreezes []*AddressFreeze `protobuf:"bytes,1,rep,name=addressFreezes,proto3" json:"addressFreezes,omitempty"`
}

func (m *AddressFreezeList) Reset()         { *m = AddressFreezeList{} }
func (m *AddressFreezeList) String() string { return proto.CompactTextString(m) }
func (*AddressFreezeList) ProtoMessage()    {}
func (*AddressFreezeList) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6927cdb03dd0514, []int{1}
}
func (m *AddressFreezeList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AddressFreezeList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AddressFreezeList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AddressFreezeList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressFreezeList.Merge(m, src)
}
func (m *AddressFreezeList) XXX_Size() int {
	return m.Size()
}
func (m *AddressFreezeList) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressFreezeList.DiscardUnknown(m)
}

var xxx_messageInfo_AddressFreezeList proto.InternalMessageInfo

func (m *AddressFreezeList) GetAddressFreezes() []*AddressFreeze {
	if m != nil {
		return m.AddressFreezes
	}
	return nil
}

type AddressFreeze struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	In      bool   `protobuf:"varint,2,opt,name=in,proto3" json:"in,omitempty"`
	Out     bool   `protobuf:"varint,3,opt,name=out,proto3" json:"out,omitempty"`
}

func (m *AddressFreeze) Reset()         { *m = AddressFreeze{} }
func (m *AddressFreeze) String() string { return proto.CompactTextString(m) }
func (*AddressFreeze) ProtoMessage()    {}
func (*AddressFreeze) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6927cdb03dd0514, []int{2}
}
func (m *AddressFreeze) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AddressFreeze) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AddressFreeze.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AddressFreeze) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressFreeze.Merge(m, src)
}
func (m *AddressFreeze) XXX_Size() int {
	return m.Size()
}
func (m *AddressFreeze) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressFreeze.DiscardUnknown(m)
}

var xxx_messageInfo_AddressFreeze proto.InternalMessageInfo

func (m *AddressFreeze) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *AddressFreeze) GetIn() bool {
	if m != nil {
		return m.In
	}
	return false
}

func (m *AddressFreeze) GetOut() bool {
	if m != nil {
		return m.Out
	}
	return false
}

func init() {
	proto.RegisterType((*Freeze)(nil), "konstellation.issue.Freeze")
	proto.RegisterType((*AddressFreezeList)(nil), "konstellation.issue.AddressFreezeList")
	proto.RegisterType((*AddressFreeze)(nil), "konstellation.issue.AddressFreeze")
}

func init() { proto.RegisterFile("konstellation/issue/freeze.proto", fileDescriptor_e6927cdb03dd0514) }

var fileDescriptor_e6927cdb03dd0514 = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xc8, 0xce, 0xcf, 0x2b,
	0x2e, 0x49, 0xcd, 0xc9, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0xcf, 0x2c, 0x2e, 0x2e, 0x4d, 0xd5,
	0x4f, 0x2b, 0x4a, 0x4d, 0xad, 0x4a, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x46, 0x51,
	0xa1, 0x07, 0x56, 0xa1, 0xa4, 0xc5, 0xc5, 0xe6, 0x06, 0x56, 0x24, 0xc4, 0xc7, 0xc5, 0x94, 0x99,
	0x27, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x11, 0xc4, 0x94, 0x99, 0x27, 0x24, 0xc0, 0xc5, 0x9c, 0x5f,
	0x5a, 0x22, 0xc1, 0x04, 0x16, 0x00, 0x31, 0x95, 0xe2, 0xb9, 0x04, 0x1d, 0x53, 0x52, 0x8a, 0x52,
	0x8b, 0x8b, 0x21, 0x5a, 0x7c, 0x32, 0x8b, 0x4b, 0x84, 0xbc, 0xb8, 0xf8, 0x12, 0x91, 0x05, 0x8b,
	0x25, 0x18, 0x15, 0x98, 0x35, 0xb8, 0x8d, 0x94, 0xf4, 0xb0, 0x58, 0xa7, 0x87, 0xa2, 0x3f, 0x08,
	0x4d, 0xa7, 0x92, 0x37, 0x17, 0x2f, 0x8a, 0x02, 0x21, 0x09, 0x2e, 0x76, 0xa8, 0x12, 0xb0, 0xc3,
	0x38, 0x83, 0x60, 0x5c, 0xa8, 0x6b, 0x99, 0xd0, 0x5d, 0xcb, 0x0c, 0x77, 0xad, 0x93, 0xfa, 0x89,
	0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3,
	0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0xf1, 0x56, 0x40, 0x83, 0xa7, 0xa4, 0xb2,
	0x20, 0xb5, 0x38, 0x89, 0x0d, 0x1c, 0x3c, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x89,
	0xc4, 0x75, 0x42, 0x01, 0x00, 0x00,
}

func (m *Freeze) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Freeze) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Freeze) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Out {
		i--
		if m.Out {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.In {
		i--
		if m.In {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *AddressFreezeList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AddressFreezeList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AddressFreezeList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AddressFreezes) > 0 {
		for iNdEx := len(m.AddressFreezes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AddressFreezes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintFreeze(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *AddressFreeze) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AddressFreeze) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AddressFreeze) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Out {
		i--
		if m.Out {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.In {
		i--
		if m.In {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintFreeze(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintFreeze(dAtA []byte, offset int, v uint64) int {
	offset -= sovFreeze(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Freeze) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.In {
		n += 2
	}
	if m.Out {
		n += 2
	}
	return n
}

func (m *AddressFreezeList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.AddressFreezes) > 0 {
		for _, e := range m.AddressFreezes {
			l = e.Size()
			n += 1 + l + sovFreeze(uint64(l))
		}
	}
	return n
}

func (m *AddressFreeze) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovFreeze(uint64(l))
	}
	if m.In {
		n += 2
	}
	if m.Out {
		n += 2
	}
	return n
}

func sovFreeze(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFreeze(x uint64) (n int) {
	return sovFreeze(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Freeze) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFreeze
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
			return fmt.Errorf("proto: Freeze: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Freeze: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field In", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFreeze
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
			m.In = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Out", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFreeze
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
			m.Out = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipFreeze(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFreeze
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
func (m *AddressFreezeList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFreeze
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
			return fmt.Errorf("proto: AddressFreezeList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AddressFreezeList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddressFreezes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFreeze
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
				return ErrInvalidLengthFreeze
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFreeze
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AddressFreezes = append(m.AddressFreezes, &AddressFreeze{})
			if err := m.AddressFreezes[len(m.AddressFreezes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFreeze(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFreeze
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
func (m *AddressFreeze) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFreeze
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
			return fmt.Errorf("proto: AddressFreeze: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AddressFreeze: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFreeze
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
				return ErrInvalidLengthFreeze
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFreeze
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field In", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFreeze
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
			m.In = bool(v != 0)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Out", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFreeze
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
			m.Out = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipFreeze(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFreeze
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
func skipFreeze(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFreeze
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
					return 0, ErrIntOverflowFreeze
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
					return 0, ErrIntOverflowFreeze
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
				return 0, ErrInvalidLengthFreeze
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFreeze
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFreeze
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFreeze        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFreeze          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFreeze = fmt.Errorf("proto: unexpected end of group")
)
