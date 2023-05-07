// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/cosmwasmpool/v1beta1/model/pool.proto

package model

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
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

type CosmWasmPool struct {
	PoolAddress     string `protobuf:"bytes,1,opt,name=pool_address,json=poolAddress,proto3" json:"pool_address,omitempty" yaml:"pool_address"`
	ContractAddress string `protobuf:"bytes,2,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty" yaml:"contract_address"`
	PoolId          uint64 `protobuf:"varint,3,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	CodeId          uint64 `protobuf:"varint,4,opt,name=code_id,json=codeId,proto3" json:"code_id,omitempty"`
}

func (m *CosmWasmPool) Reset()      { *m = CosmWasmPool{} }
func (*CosmWasmPool) ProtoMessage() {}
func (*CosmWasmPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0cb64564a744af1, []int{0}
}
func (m *CosmWasmPool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CosmWasmPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CosmWasmPool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CosmWasmPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CosmWasmPool.Merge(m, src)
}
func (m *CosmWasmPool) XXX_Size() int {
	return m.Size()
}
func (m *CosmWasmPool) XXX_DiscardUnknown() {
	xxx_messageInfo_CosmWasmPool.DiscardUnknown(m)
}

var xxx_messageInfo_CosmWasmPool proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CosmWasmPool)(nil), "osmosis.cosmwasmpool.v1beta1.CosmWasmPool")
}

func init() {
	proto.RegisterFile("osmosis/cosmwasmpool/v1beta1/model/pool.proto", fileDescriptor_a0cb64564a744af1)
}

var fileDescriptor_a0cb64564a744af1 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x4d, 0x4a, 0xc3, 0x40,
	0x14, 0xc7, 0x33, 0x5a, 0x2b, 0xc6, 0x82, 0x5a, 0x85, 0xd6, 0x2a, 0x49, 0xc9, 0xaa, 0x9b, 0x66,
	0x28, 0x22, 0x48, 0x77, 0x56, 0x10, 0xba, 0x93, 0x6c, 0x04, 0x37, 0x65, 0x92, 0x19, 0x63, 0x60,
	0xc6, 0x57, 0x3a, 0xd3, 0xaa, 0x37, 0x70, 0xe9, 0xd2, 0x65, 0x0f, 0xe1, 0x21, 0xc4, 0x55, 0x77,
	0xba, 0x2a, 0xd2, 0xde, 0xa0, 0x27, 0x90, 0x99, 0x49, 0xa5, 0xba, 0xcb, 0xef, 0xff, 0x91, 0xc7,
	0xbc, 0xe7, 0x36, 0x41, 0x0a, 0x90, 0x99, 0xc4, 0x09, 0x48, 0xf1, 0x40, 0xa4, 0xe8, 0x03, 0x70,
	0x3c, 0x6a, 0xc5, 0x4c, 0x91, 0x16, 0x16, 0x40, 0x19, 0xc7, 0x5a, 0x0a, 0xfb, 0x03, 0x50, 0x50,
	0x3e, 0xce, 0xe3, 0xe1, 0x6a, 0x3c, 0xcc, 0xe3, 0xb5, 0xc3, 0xc4, 0xd8, 0x3d, 0x93, 0xc5, 0x16,
	0x6c, 0xb1, 0x76, 0x90, 0x42, 0x0a, 0x56, 0xd7, 0x5f, 0xb9, 0xea, 0xa7, 0x00, 0x29, 0x67, 0xd8,
	0x50, 0x3c, 0xbc, 0xc5, 0x2a, 0x13, 0x4c, 0x2a, 0x22, 0xfa, 0x36, 0x10, 0x7c, 0x22, 0xb7, 0x74,
	0x01, 0x52, 0x5c, 0x13, 0x29, 0xae, 0x00, 0x78, 0xb9, 0xed, 0x96, 0xf4, 0xc8, 0x1e, 0xa1, 0x74,
	0xc0, 0xa4, 0xac, 0xa2, 0x3a, 0x6a, 0x6c, 0x75, 0x2a, 0x8b, 0xa9, 0xbf, 0xff, 0x44, 0x04, 0x6f,
	0x07, 0xab, 0x6e, 0x10, 0x6d, 0x6b, 0x3c, 0xb7, 0x54, 0xbe, 0x74, 0x77, 0x13, 0xb8, 0x57, 0x03,
	0x92, 0xa8, 0xdf, 0xfe, 0x9a, 0xe9, 0x1f, 0x2d, 0xa6, 0x7e, 0xc5, 0xf6, 0xff, 0x27, 0x82, 0x68,
	0x67, 0x29, 0x2d, 0xff, 0x53, 0x71, 0x37, 0xcd, 0x94, 0x8c, 0x56, 0xd7, 0xeb, 0xa8, 0x51, 0x88,
	0x8a, 0x1a, 0xbb, 0x54, 0x1b, 0x09, 0x50, 0xa6, 0x8d, 0x82, 0x35, 0x34, 0x76, 0x69, 0x7b, 0xef,
	0x79, 0xec, 0x3b, 0xaf, 0x63, 0xdf, 0xf9, 0x78, 0x6b, 0x6e, 0xe8, 0x77, 0x74, 0x3b, 0xd1, 0xfb,
	0xcc, 0x43, 0x93, 0x99, 0x87, 0xbe, 0x67, 0x1e, 0x7a, 0x99, 0x7b, 0xce, 0x64, 0xee, 0x39, 0x5f,
	0x73, 0xcf, 0xb9, 0x39, 0x4b, 0x33, 0x75, 0x37, 0x8c, 0xc3, 0x04, 0x04, 0xce, 0xd7, 0xdd, 0xe4,
	0x24, 0x96, 0x4b, 0xc0, 0xa3, 0xd6, 0x29, 0x7e, 0xfc, 0x7b, 0x30, 0x73, 0xa8, 0xb8, 0x68, 0x96,
	0x76, 0xf2, 0x13, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x92, 0xde, 0xee, 0xd5, 0x01, 0x00, 0x00,
}

func (m *CosmWasmPool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CosmWasmPool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CosmWasmPool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CodeId != 0 {
		i = encodeVarintPool(dAtA, i, uint64(m.CodeId))
		i--
		dAtA[i] = 0x20
	}
	if m.PoolId != 0 {
		i = encodeVarintPool(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintPool(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PoolAddress) > 0 {
		i -= len(m.PoolAddress)
		copy(dAtA[i:], m.PoolAddress)
		i = encodeVarintPool(dAtA, i, uint64(len(m.PoolAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPool(dAtA []byte, offset int, v uint64) int {
	offset -= sovPool(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CosmWasmPool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PoolAddress)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	if m.PoolId != 0 {
		n += 1 + sovPool(uint64(m.PoolId))
	}
	if m.CodeId != 0 {
		n += 1 + sovPool(uint64(m.CodeId))
	}
	return n
}

func sovPool(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPool(x uint64) (n int) {
	return sovPool(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CosmWasmPool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPool
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
			return fmt.Errorf("proto: CosmWasmPool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CosmWasmPool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeId", wireType)
			}
			m.CodeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CodeId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPool
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
func skipPool(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPool
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
					return 0, ErrIntOverflowPool
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
					return 0, ErrIntOverflowPool
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
				return 0, ErrInvalidLengthPool
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPool
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPool
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPool        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPool          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPool = fmt.Errorf("proto: unexpected end of group")
)
