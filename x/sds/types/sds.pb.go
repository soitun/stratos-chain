// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stratos/sds/v1/sds.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
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

// Params defines the parameters for the sds module.
type Params struct {
	BondDenom string `protobuf:"bytes,1,opt,name=bond_denom,json=bondDenom,proto3" json:"bond_denom" yaml:"bond_denom"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_a89f3959b8649eb2, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetBondDenom() string {
	if m != nil {
		return m.BondDenom
	}
	return ""
}

type FileInfo struct {
	Height    github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=height,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"height" yaml:"height"`
	Reporters []byte                                 `protobuf:"bytes,2,opt,name=reporters,proto3" json:"reporters" yaml:"reporters"`
	Uploader  string                                 `protobuf:"bytes,3,opt,name=uploader,proto3" json:"uploader" yaml:"uploader"`
}

func (m *FileInfo) Reset()         { *m = FileInfo{} }
func (m *FileInfo) String() string { return proto.CompactTextString(m) }
func (*FileInfo) ProtoMessage()    {}
func (*FileInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a89f3959b8649eb2, []int{1}
}
func (m *FileInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FileInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FileInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FileInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileInfo.Merge(m, src)
}
func (m *FileInfo) XXX_Size() int {
	return m.Size()
}
func (m *FileInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FileInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FileInfo proto.InternalMessageInfo

func (m *FileInfo) GetReporters() []byte {
	if m != nil {
		return m.Reporters
	}
	return nil
}

func (m *FileInfo) GetUploader() string {
	if m != nil {
		return m.Uploader
	}
	return ""
}

func init() {
	proto.RegisterType((*Params)(nil), "stratos.sds.v1.Params")
	proto.RegisterType((*FileInfo)(nil), "stratos.sds.v1.FileInfo")
}

func init() { proto.RegisterFile("stratos/sds/v1/sds.proto", fileDescriptor_a89f3959b8649eb2) }

var fileDescriptor_a89f3959b8649eb2 = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0x3f, 0xeb, 0xd3, 0x40,
	0x18, 0xc7, 0x73, 0x3f, 0xa1, 0xb4, 0x87, 0xff, 0x1a, 0x3a, 0xc4, 0x82, 0xb9, 0x1a, 0x41, 0x8a,
	0xd0, 0x84, 0xe2, 0x20, 0x74, 0x11, 0x83, 0x08, 0x1d, 0x04, 0x49, 0x07, 0xc1, 0xa5, 0xa4, 0xcd,
	0x99, 0x84, 0x36, 0x77, 0xe1, 0xee, 0x5a, 0xec, 0x5b, 0x70, 0xd2, 0xcd, 0xd1, 0xd1, 0xd1, 0xc1,
	0x17, 0xd1, 0xb1, 0x38, 0x89, 0xc3, 0x21, 0xe9, 0x20, 0x64, 0xcc, 0x2b, 0x90, 0xe4, 0xae, 0x7f,
	0x96, 0xdc, 0xf3, 0x7c, 0xbe, 0xf7, 0xe4, 0xf9, 0x3e, 0xcf, 0x41, 0x8b, 0x0b, 0x16, 0x0a, 0xca,
	0x3d, 0x1e, 0x71, 0x6f, 0x3b, 0xae, 0x0f, 0x37, 0x67, 0x54, 0x50, 0xf3, 0xae, 0x56, 0xdc, 0x1a,
	0x6d, 0xc7, 0xfd, 0x5e, 0x4c, 0x63, 0xda, 0x48, 0x5e, 0x1d, 0xa9, 0x5b, 0xfd, 0x07, 0x4b, 0xca,
	0x33, 0xca, 0xe7, 0x4a, 0x50, 0x89, 0x96, 0xba, 0x61, 0x96, 0x12, 0xea, 0x35, 0x5f, 0x85, 0x9c,
	0x15, 0x6c, 0xbd, 0x0d, 0x59, 0x98, 0x71, 0xd3, 0x87, 0x70, 0x41, 0x49, 0x34, 0x8f, 0x30, 0xa1,
	0x99, 0x05, 0x06, 0x60, 0xd8, 0xf1, 0x1f, 0x97, 0x12, 0x5d, 0xd1, 0x4a, 0xa2, 0xee, 0x2e, 0xcc,
	0xd6, 0x13, 0xe7, 0xc2, 0x9c, 0xa0, 0x53, 0x27, 0xaf, 0xea, 0x78, 0xf2, 0xf0, 0xeb, 0x37, 0x04,
	0x3e, 0xfd, 0xfb, 0xf1, 0xb4, 0x77, 0x1a, 0xe2, 0x63, 0x33, 0x86, 0x6a, 0xe1, 0x7c, 0xb9, 0x81,
	0xed, 0xd7, 0xe9, 0x1a, 0x4f, 0xc9, 0x07, 0x6a, 0x52, 0xd8, 0x4a, 0x70, 0x1a, 0x27, 0x42, 0xf7,
	0x7a, 0xb7, 0x97, 0xc8, 0xf8, 0x23, 0xd1, 0x93, 0x38, 0x15, 0xc9, 0x66, 0xe1, 0x2e, 0x69, 0xa6,
	0xdd, 0xeb, 0x63, 0xc4, 0xa3, 0x95, 0x27, 0x76, 0x39, 0xe6, 0xee, 0x94, 0x88, 0x52, 0x22, 0x5d,
	0x5f, 0x49, 0x74, 0x47, 0xb9, 0x52, 0xb9, 0xf3, 0xeb, 0xe7, 0x08, 0xea, 0xb9, 0xa7, 0x44, 0x04,
	0xfa, 0x9a, 0xf9, 0x02, 0x76, 0x18, 0xce, 0x29, 0x13, 0x98, 0x71, 0xeb, 0x66, 0x00, 0x86, 0xb7,
	0xfd, 0x47, 0xa5, 0x44, 0x17, 0x58, 0x49, 0x74, 0x5f, 0xfd, 0xe8, 0x8c, 0x9c, 0xe0, 0x22, 0x9b,
	0x33, 0xd8, 0xde, 0xe4, 0x6b, 0x1a, 0x46, 0x98, 0x59, 0xb7, 0x1a, 0xcf, 0xcf, 0x4b, 0x89, 0xce,
	0xac, 0x92, 0xe8, 0x9e, 0x2a, 0x3f, 0x91, 0xda, 0x49, 0x4f, 0x3b, 0x79, 0x19, 0x45, 0x0c, 0x73,
	0x3e, 0x13, 0x2c, 0x25, 0x71, 0x70, 0x2e, 0xf2, 0xdf, 0x7c, 0x2f, 0x6c, 0xb0, 0x2f, 0x6c, 0x70,
	0x28, 0x6c, 0xf0, 0xb7, 0xb0, 0xc1, 0xe7, 0xa3, 0x6d, 0x1c, 0x8e, 0xb6, 0xf1, 0xfb, 0x68, 0x1b,
	0xef, 0xbd, 0xab, 0x65, 0xe8, 0x95, 0x12, 0x2c, 0x4e, 0xe1, 0x68, 0x99, 0x84, 0x29, 0xd1, 0x3b,
	0x6e, 0x36, 0xb3, 0x68, 0x35, 0xcf, 0xfa, 0xec, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x07, 0x87,
	0x09, 0x64, 0x46, 0x02, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
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
	if this.BondDenom != that1.BondDenom {
		return false
	}
	return true
}
func (this *FileInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FileInfo)
	if !ok {
		that2, ok := that.(FileInfo)
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
	if !this.Height.Equal(that1.Height) {
		return false
	}
	if !bytes.Equal(this.Reporters, that1.Reporters) {
		return false
	}
	if this.Uploader != that1.Uploader {
		return false
	}
	return true
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BondDenom) > 0 {
		i -= len(m.BondDenom)
		copy(dAtA[i:], m.BondDenom)
		i = encodeVarintSds(dAtA, i, uint64(len(m.BondDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *FileInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FileInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FileInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Uploader) > 0 {
		i -= len(m.Uploader)
		copy(dAtA[i:], m.Uploader)
		i = encodeVarintSds(dAtA, i, uint64(len(m.Uploader)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Reporters) > 0 {
		i -= len(m.Reporters)
		copy(dAtA[i:], m.Reporters)
		i = encodeVarintSds(dAtA, i, uint64(len(m.Reporters)))
		i--
		dAtA[i] = 0x12
	}
	{
		size := m.Height.Size()
		i -= size
		if _, err := m.Height.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSds(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintSds(dAtA []byte, offset int, v uint64) int {
	offset -= sovSds(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BondDenom)
	if l > 0 {
		n += 1 + l + sovSds(uint64(l))
	}
	return n
}

func (m *FileInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Height.Size()
	n += 1 + l + sovSds(uint64(l))
	l = len(m.Reporters)
	if l > 0 {
		n += 1 + l + sovSds(uint64(l))
	}
	l = len(m.Uploader)
	if l > 0 {
		n += 1 + l + sovSds(uint64(l))
	}
	return n
}

func sovSds(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSds(x uint64) (n int) {
	return sovSds(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSds
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BondDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSds
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
				return ErrInvalidLengthSds
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSds
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BondDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSds
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
func (m *FileInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSds
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
			return fmt.Errorf("proto: FileInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FileInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSds
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
				return ErrInvalidLengthSds
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSds
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Height.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reporters", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSds
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSds
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSds
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reporters = append(m.Reporters[:0], dAtA[iNdEx:postIndex]...)
			if m.Reporters == nil {
				m.Reporters = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uploader", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSds
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
				return ErrInvalidLengthSds
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSds
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uploader = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSds
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
func skipSds(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSds
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
					return 0, ErrIntOverflowSds
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
					return 0, ErrIntOverflowSds
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
				return 0, ErrInvalidLengthSds
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSds
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSds
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSds        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSds          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSds = fmt.Errorf("proto: unexpected end of group")
)
