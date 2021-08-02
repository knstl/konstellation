// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: konstellation/issue/allowance.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

type Allowance struct {
	Amount  github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount"`
	Spender string                                 `protobuf:"bytes,2,opt,name=spender,proto3" json:"spender,omitempty"`
}

func (m *Allowance) Reset()         { *m = Allowance{} }
func (m *Allowance) String() string { return proto.CompactTextString(m) }
func (*Allowance) ProtoMessage()    {}
func (*Allowance) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f998091bc225886, []int{0}
}
func (m *Allowance) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Allowance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Allowance.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Allowance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Allowance.Merge(m, src)
}
func (m *Allowance) XXX_Size() int {
	return m.Size()
}
func (m *Allowance) XXX_DiscardUnknown() {
	xxx_messageInfo_Allowance.DiscardUnknown(m)
}

var xxx_messageInfo_Allowance proto.InternalMessageInfo

func (m *Allowance) GetSpender() string {
	if m != nil {
		return m.Spender
	}
	return ""
}

type AllowanceList struct {
	Allowances []*Allowance `protobuf:"bytes,1,rep,name=allowances,proto3" json:"allowances,omitempty"`
}

func (m *AllowanceList) Reset()         { *m = AllowanceList{} }
func (m *AllowanceList) String() string { return proto.CompactTextString(m) }
func (*AllowanceList) ProtoMessage()    {}
func (*AllowanceList) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f998091bc225886, []int{1}
}
func (m *AllowanceList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AllowanceList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AllowanceList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AllowanceList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllowanceList.Merge(m, src)
}
func (m *AllowanceList) XXX_Size() int {
	return m.Size()
}
func (m *AllowanceList) XXX_DiscardUnknown() {
	xxx_messageInfo_AllowanceList.DiscardUnknown(m)
}

var xxx_messageInfo_AllowanceList proto.InternalMessageInfo

func (m *AllowanceList) GetAllowances() []*Allowance {
	if m != nil {
		return m.Allowances
	}
	return nil
}

func init() {
	proto.RegisterType((*Allowance)(nil), "konstellation.issue.Allowance")
	proto.RegisterType((*AllowanceList)(nil), "konstellation.issue.AllowanceList")
}

func init() {
	proto.RegisterFile("konstellation/issue/allowance.proto", fileDescriptor_0f998091bc225886)
}

var fileDescriptor_0f998091bc225886 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xce, 0xce, 0xcf, 0x2b,
	0x2e, 0x49, 0xcd, 0xc9, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0xcf, 0x2c, 0x2e, 0x2e, 0x4d, 0xd5,
	0x4f, 0xcc, 0xc9, 0xc9, 0x2f, 0x4f, 0xcc, 0x4b, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x12, 0x46, 0x51, 0xa4, 0x07, 0x56, 0x24, 0x25, 0x92, 0x9e, 0x9f, 0x9e, 0x0f, 0x96, 0xd7, 0x07,
	0xb1, 0x20, 0x4a, 0x95, 0x72, 0xb9, 0x38, 0x1d, 0x61, 0xba, 0x85, 0xdc, 0xb8, 0xd8, 0x12, 0x73,
	0xf3, 0x4b, 0xf3, 0x4a, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x9d, 0xf4, 0x4e, 0xdc, 0x93, 0x67,
	0xb8, 0x75, 0x4f, 0x5e, 0x2d, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x3f,
	0x39, 0xbf, 0x38, 0x37, 0xbf, 0x18, 0x4a, 0xe9, 0x16, 0xa7, 0x64, 0xeb, 0x97, 0x54, 0x16, 0xa4,
	0x16, 0xeb, 0x79, 0xe6, 0x95, 0x04, 0x41, 0x75, 0x0b, 0x49, 0x70, 0xb1, 0x17, 0x17, 0xa4, 0xe6,
	0xa5, 0xa4, 0x16, 0x49, 0x30, 0x81, 0x0c, 0x0a, 0x82, 0x71, 0x95, 0xfc, 0xb9, 0x78, 0xe1, 0xd6,
	0xf9, 0x64, 0x16, 0x97, 0x08, 0xd9, 0x71, 0x71, 0xc1, 0x5d, 0x5f, 0x2c, 0xc1, 0xa8, 0xc0, 0xac,
	0xc1, 0x6d, 0x24, 0xa7, 0x87, 0xc5, 0xfd, 0x7a, 0x70, 0x7d, 0x41, 0x48, 0x3a, 0x9c, 0xd4, 0x4f,
	0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18,
	0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x8a, 0xb7, 0x02, 0x1a, 0x3a, 0x60, 0xf7,
	0x25, 0xb1, 0x81, 0xfd, 0x6b, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x29, 0x88, 0x05, 0x2e, 0x41,
	0x01, 0x00, 0x00,
}

func (m *Allowance) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Allowance) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Allowance) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Spender) > 0 {
		i -= len(m.Spender)
		copy(dAtA[i:], m.Spender)
		i = encodeVarintAllowance(dAtA, i, uint64(len(m.Spender)))
		i--
		dAtA[i] = 0x12
	}
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAllowance(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *AllowanceList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AllowanceList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AllowanceList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Allowances) > 0 {
		for iNdEx := len(m.Allowances) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Allowances[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAllowance(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintAllowance(dAtA []byte, offset int, v uint64) int {
	offset -= sovAllowance(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Allowance) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Amount.Size()
	n += 1 + l + sovAllowance(uint64(l))
	l = len(m.Spender)
	if l > 0 {
		n += 1 + l + sovAllowance(uint64(l))
	}
	return n
}

func (m *AllowanceList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Allowances) > 0 {
		for _, e := range m.Allowances {
			l = e.Size()
			n += 1 + l + sovAllowance(uint64(l))
		}
	}
	return n
}

func sovAllowance(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAllowance(x uint64) (n int) {
	return sovAllowance(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Allowance) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAllowance
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
			return fmt.Errorf("proto: Allowance: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Allowance: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAllowance
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
				return ErrInvalidLengthAllowance
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAllowance
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAllowance
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
				return ErrInvalidLengthAllowance
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAllowance
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Spender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAllowance(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAllowance
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
func (m *AllowanceList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAllowance
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
			return fmt.Errorf("proto: AllowanceList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AllowanceList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Allowances", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAllowance
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
				return ErrInvalidLengthAllowance
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAllowance
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Allowances = append(m.Allowances, &Allowance{})
			if err := m.Allowances[len(m.Allowances)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAllowance(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAllowance
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
func skipAllowance(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAllowance
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
					return 0, ErrIntOverflowAllowance
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
					return 0, ErrIntOverflowAllowance
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
				return 0, ErrInvalidLengthAllowance
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAllowance
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAllowance
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAllowance        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAllowance          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAllowance = fmt.Errorf("proto: unexpected end of group")
)
