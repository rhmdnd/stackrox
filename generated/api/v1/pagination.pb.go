// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/v1/pagination.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SortOption struct {
	Field                string   `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Reversed             bool     `protobuf:"varint,2,opt,name=reversed,proto3" json:"reversed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SortOption) Reset()         { *m = SortOption{} }
func (m *SortOption) String() string { return proto.CompactTextString(m) }
func (*SortOption) ProtoMessage()    {}
func (*SortOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_5993a9cae165a01a, []int{0}
}
func (m *SortOption) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SortOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SortOption.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SortOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SortOption.Merge(m, src)
}
func (m *SortOption) XXX_Size() int {
	return m.Size()
}
func (m *SortOption) XXX_DiscardUnknown() {
	xxx_messageInfo_SortOption.DiscardUnknown(m)
}

var xxx_messageInfo_SortOption proto.InternalMessageInfo

func (m *SortOption) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *SortOption) GetReversed() bool {
	if m != nil {
		return m.Reversed
	}
	return false
}

func (m *SortOption) MessageClone() proto.Message {
	return m.Clone()
}
func (m *SortOption) Clone() *SortOption {
	if m == nil {
		return nil
	}
	cloned := new(SortOption)
	*cloned = *m

	return cloned
}

type Pagination struct {
	Limit                int32       `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               int32       `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	SortOption           *SortOption `protobuf:"bytes,3,opt,name=sort_option,json=sortOption,proto3" json:"sort_option,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Pagination) Reset()         { *m = Pagination{} }
func (m *Pagination) String() string { return proto.CompactTextString(m) }
func (*Pagination) ProtoMessage()    {}
func (*Pagination) Descriptor() ([]byte, []int) {
	return fileDescriptor_5993a9cae165a01a, []int{1}
}
func (m *Pagination) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pagination) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pagination.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pagination) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pagination.Merge(m, src)
}
func (m *Pagination) XXX_Size() int {
	return m.Size()
}
func (m *Pagination) XXX_DiscardUnknown() {
	xxx_messageInfo_Pagination.DiscardUnknown(m)
}

var xxx_messageInfo_Pagination proto.InternalMessageInfo

func (m *Pagination) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *Pagination) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *Pagination) GetSortOption() *SortOption {
	if m != nil {
		return m.SortOption
	}
	return nil
}

func (m *Pagination) MessageClone() proto.Message {
	return m.Clone()
}
func (m *Pagination) Clone() *Pagination {
	if m == nil {
		return nil
	}
	cloned := new(Pagination)
	*cloned = *m

	cloned.SortOption = m.SortOption.Clone()
	return cloned
}

func init() {
	proto.RegisterType((*SortOption)(nil), "v1.SortOption")
	proto.RegisterType((*Pagination)(nil), "v1.Pagination")
}

func init() { proto.RegisterFile("api/v1/pagination.proto", fileDescriptor_5993a9cae165a01a) }

var fileDescriptor_5993a9cae165a01a = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0x2c, 0xc8, 0xd4,
	0x2f, 0x33, 0xd4, 0x2f, 0x48, 0x4c, 0xcf, 0xcc, 0x4b, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x33, 0x54, 0xb2, 0xe3, 0xe2, 0x0a, 0xce, 0x2f, 0x2a, 0xf1,
	0x2f, 0x00, 0x89, 0x0b, 0x89, 0x70, 0xb1, 0xa6, 0x65, 0xa6, 0xe6, 0xa4, 0x48, 0x30, 0x2a, 0x30,
	0x6a, 0x70, 0x06, 0x41, 0x38, 0x42, 0x52, 0x5c, 0x1c, 0x45, 0xa9, 0x65, 0xa9, 0x45, 0xc5, 0xa9,
	0x29, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x1c, 0x41, 0x70, 0xbe, 0x52, 0x36, 0x17, 0x57, 0x00, 0xdc,
	0x5c, 0x90, 0xfe, 0x9c, 0xcc, 0xdc, 0xcc, 0x12, 0xb0, 0x7e, 0xd6, 0x20, 0x08, 0x47, 0x48, 0x8c,
	0x8b, 0x2d, 0x3f, 0x2d, 0xad, 0x38, 0xb5, 0x04, 0xac, 0x9b, 0x35, 0x08, 0xca, 0x13, 0xd2, 0xe7,
	0xe2, 0x2e, 0xce, 0x2f, 0x2a, 0x89, 0xcf, 0x07, 0x5b, 0x2e, 0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0x6d,
	0xc4, 0xa7, 0x57, 0x66, 0xa8, 0x87, 0x70, 0x52, 0x10, 0x57, 0x31, 0x9c, 0xed, 0xa4, 0x77, 0xe2,
	0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0xce, 0x78, 0x2c, 0xc7, 0xc0,
	0x25, 0x91, 0x99, 0xaf, 0x57, 0x5c, 0x92, 0x98, 0x9c, 0x5d, 0x94, 0x5f, 0x01, 0xf1, 0x95, 0x5e,
	0x62, 0x41, 0xa6, 0x5e, 0x99, 0x61, 0x14, 0x53, 0x99, 0x61, 0x12, 0x1b, 0x58, 0xc4, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0x99, 0x29, 0xff, 0x9d, 0x02, 0x01, 0x00, 0x00,
}

func (m *SortOption) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SortOption) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SortOption) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Reversed {
		i--
		if m.Reversed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.Field) > 0 {
		i -= len(m.Field)
		copy(dAtA[i:], m.Field)
		i = encodeVarintPagination(dAtA, i, uint64(len(m.Field)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Pagination) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pagination) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Pagination) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.SortOption != nil {
		{
			size, err := m.SortOption.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPagination(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Offset != 0 {
		i = encodeVarintPagination(dAtA, i, uint64(m.Offset))
		i--
		dAtA[i] = 0x10
	}
	if m.Limit != 0 {
		i = encodeVarintPagination(dAtA, i, uint64(m.Limit))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPagination(dAtA []byte, offset int, v uint64) int {
	offset -= sovPagination(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SortOption) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Field)
	if l > 0 {
		n += 1 + l + sovPagination(uint64(l))
	}
	if m.Reversed {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Pagination) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Limit != 0 {
		n += 1 + sovPagination(uint64(m.Limit))
	}
	if m.Offset != 0 {
		n += 1 + sovPagination(uint64(m.Offset))
	}
	if m.SortOption != nil {
		l = m.SortOption.Size()
		n += 1 + l + sovPagination(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovPagination(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPagination(x uint64) (n int) {
	return sovPagination(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SortOption) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPagination
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
			return fmt.Errorf("proto: SortOption: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SortOption: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPagination
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
				return ErrInvalidLengthPagination
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPagination
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Field = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reversed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPagination
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
			m.Reversed = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipPagination(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPagination
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Pagination) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPagination
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
			return fmt.Errorf("proto: Pagination: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pagination: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPagination
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Limit |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPagination
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Offset |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SortOption", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPagination
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
				return ErrInvalidLengthPagination
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPagination
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SortOption == nil {
				m.SortOption = &SortOption{}
			}
			if err := m.SortOption.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPagination(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPagination
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPagination(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPagination
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
					return 0, ErrIntOverflowPagination
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
					return 0, ErrIntOverflowPagination
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
				return 0, ErrInvalidLengthPagination
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPagination
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPagination
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPagination        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPagination          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPagination = fmt.Errorf("proto: unexpected end of group")
)
