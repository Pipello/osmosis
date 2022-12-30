// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/concentrated-liquidity/position.proto

// this is a legacy package that requires additional migration logic
// in order to use the correct packge. Decision made to use legacy package path
// until clear steps for migration logic and the unknowns for state breaking are
// investigated for changing proto package.

package model

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type Position struct {
	Liquidity github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=liquidity,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"liquidity" yaml:"liquidity"`
	JoinTime  time.Time                              `protobuf:"bytes,2,opt,name=join_time,json=joinTime,proto3,stdtime" json:"join_time" yaml:"join_time"`
}

func (m *Position) Reset()         { *m = Position{} }
func (m *Position) String() string { return proto.CompactTextString(m) }
func (*Position) ProtoMessage()    {}
func (*Position) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffdfd7b30d37d326, []int{0}
}
func (m *Position) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Position) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Position.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Position) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Position.Merge(m, src)
}
func (m *Position) XXX_Size() int {
	return m.Size()
}
func (m *Position) XXX_DiscardUnknown() {
	xxx_messageInfo_Position.DiscardUnknown(m)
}

var xxx_messageInfo_Position proto.InternalMessageInfo

func (m *Position) GetJoinTime() time.Time {
	if m != nil {
		return m.JoinTime
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*Position)(nil), "osmosis.concentratedliquidity.v1beta1.Position")
}

func init() {
	proto.RegisterFile("osmosis/concentrated-liquidity/position.proto", fileDescriptor_ffdfd7b30d37d326)
}

var fileDescriptor_ffdfd7b30d37d326 = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x3f, 0x4e, 0xc3, 0x30,
	0x18, 0xc5, 0x63, 0x06, 0xd4, 0x86, 0x05, 0x55, 0x0c, 0xa5, 0x42, 0x4e, 0x15, 0x09, 0xd4, 0x25,
	0xb6, 0x4a, 0x37, 0xc6, 0x88, 0x03, 0xa0, 0x0a, 0x16, 0x84, 0x54, 0xf2, 0xc7, 0x04, 0x43, 0x9c,
	0x2f, 0xd4, 0x6e, 0x45, 0x6f, 0xd1, 0x3b, 0xb1, 0x74, 0xec, 0x88, 0x18, 0x02, 0x4a, 0x6e, 0xc0,
	0x09, 0x50, 0x12, 0x27, 0xed, 0xc2, 0x64, 0xbf, 0x4f, 0xef, 0x7b, 0xef, 0x27, 0xdb, 0x74, 0x40,
	0x0a, 0x90, 0x5c, 0xd2, 0x00, 0x92, 0x80, 0x25, 0x6a, 0xee, 0x29, 0x16, 0x3a, 0x31, 0x7f, 0x5b,
	0xf0, 0x90, 0xab, 0x15, 0x4d, 0x41, 0x72, 0xc5, 0x21, 0x21, 0xe9, 0x1c, 0x14, 0xf4, 0xce, 0xb5,
	0x9d, 0xec, 0xdb, 0x5b, 0x37, 0x59, 0x8e, 0x7d, 0xa6, 0xbc, 0xf1, 0xe0, 0x34, 0xa8, 0x7c, 0xb3,
	0x6a, 0x89, 0xd6, 0xa2, 0x4e, 0x18, 0x9c, 0x44, 0x10, 0x41, 0x3d, 0x2f, 0x6f, 0x7a, 0x6a, 0x45,
	0x00, 0x51, 0xcc, 0x68, 0xa5, 0xfc, 0xc5, 0x13, 0x55, 0x5c, 0x30, 0xa9, 0x3c, 0x91, 0xd6, 0x06,
	0xfb, 0x03, 0x99, 0x9d, 0x1b, 0xcd, 0xd2, 0x7b, 0x34, 0xbb, 0x6d, 0x67, 0x1f, 0x0d, 0xd1, 0xa8,
	0xeb, 0xba, 0x9b, 0xcc, 0x32, 0xbe, 0x32, 0xeb, 0x22, 0xe2, 0xea, 0x79, 0xe1, 0x93, 0x00, 0x84,
	0xee, 0xd5, 0x87, 0x23, 0xc3, 0x57, 0xaa, 0x56, 0x29, 0x93, 0xe4, 0x9a, 0x05, 0xbf, 0x99, 0x75,
	0xbc, 0xf2, 0x44, 0x7c, 0x65, 0xb7, 0x41, 0xf6, 0x74, 0x17, 0xda, 0xbb, 0x33, 0xbb, 0x2f, 0xc0,
	0x93, 0x59, 0x89, 0xd1, 0x3f, 0x18, 0xa2, 0xd1, 0xd1, 0xe5, 0x80, 0xd4, 0x8c, 0xa4, 0x61, 0x24,
	0xb7, 0x0d, 0xa3, 0x7b, 0x56, 0xb6, 0xef, 0x32, 0xdb, 0x55, 0x7b, 0xfd, 0x6d, 0xa1, 0x69, 0xa7,
	0xd4, 0xa5, 0xd9, 0x7d, 0xd8, 0xe4, 0x18, 0x6d, 0x73, 0x8c, 0x7e, 0x72, 0x8c, 0xd6, 0x05, 0x36,
	0xb6, 0x05, 0x36, 0x3e, 0x0b, 0x6c, 0xdc, 0xbb, 0x7b, 0xdc, 0xfa, 0x8d, 0x9d, 0xd8, 0xf3, 0x65,
	0x23, 0xe8, 0x72, 0x3c, 0xa1, 0xef, 0xff, 0xfd, 0x92, 0x80, 0x90, 0xc5, 0xfe, 0x61, 0x45, 0x36,
	0xf9, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x4b, 0xa0, 0xe9, 0xd4, 0x01, 0x00, 0x00,
}

func (m *Position) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Position) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Position) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.JoinTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.JoinTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintPosition(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	{
		size := m.Liquidity.Size()
		i -= size
		if _, err := m.Liquidity.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPosition(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintPosition(dAtA []byte, offset int, v uint64) int {
	offset -= sovPosition(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Position) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Liquidity.Size()
	n += 1 + l + sovPosition(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.JoinTime)
	n += 1 + l + sovPosition(uint64(l))
	return n
}

func sovPosition(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPosition(x uint64) (n int) {
	return sovPosition(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Position) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPosition
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
			return fmt.Errorf("proto: Position: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Position: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Liquidity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPosition
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
				return ErrInvalidLengthPosition
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPosition
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Liquidity.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JoinTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPosition
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
				return ErrInvalidLengthPosition
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPosition
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.JoinTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPosition(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPosition
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
func skipPosition(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPosition
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
					return 0, ErrIntOverflowPosition
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
					return 0, ErrIntOverflowPosition
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
				return 0, ErrInvalidLengthPosition
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPosition
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPosition
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPosition        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPosition          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPosition = fmt.Errorf("proto: unexpected end of group")
)