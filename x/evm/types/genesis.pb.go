// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stratos/evm/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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

// GenesisState defines the evm module's genesis state.
type GenesisState struct {
	// params defines all the parameters of the module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params" yaml:"params"`
	// accounts is an array containing the ethereum genesis accounts.
	Accounts []GenesisAccount `protobuf:"bytes,2,rep,name=accounts,proto3" json:"accounts" yaml:"accounts"`
	// block gas is the amount of gas used on the last block before the upgrade.
	// Zero by default.
	BlockGas uint64 `protobuf:"varint,3,opt,name=block_gas,json=blockGas,proto3" json:"block_gas" yaml:"block_gas"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bfc0e5ccf7131d4, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetAccounts() []GenesisAccount {
	if m != nil {
		return m.Accounts
	}
	return nil
}

func (m *GenesisState) GetBlockGas() uint64 {
	if m != nil {
		return m.BlockGas
	}
	return 0
}

// GenesisAccount defines an account to be initialized in the genesis state.
// Its main difference between with Geth's GenesisAccount is that it uses a
// custom storage type and that it doesn't contain the private key field.
type GenesisAccount struct {
	// address defines an ethereum hex formated address of an account
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address" yaml:"address"`
	// code defines the hex bytes of the account code.
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code" yaml:"code"`
	// storage defines the set of state key values for the account.
	Storage Storage `protobuf:"bytes,3,rep,name=storage,proto3,castrepeated=Storage" json:"storage" yaml:"storage"`
}

func (m *GenesisAccount) Reset()         { *m = GenesisAccount{} }
func (m *GenesisAccount) String() string { return proto.CompactTextString(m) }
func (*GenesisAccount) ProtoMessage()    {}
func (*GenesisAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bfc0e5ccf7131d4, []int{1}
}
func (m *GenesisAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisAccount.Merge(m, src)
}
func (m *GenesisAccount) XXX_Size() int {
	return m.Size()
}
func (m *GenesisAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisAccount.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisAccount proto.InternalMessageInfo

func (m *GenesisAccount) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *GenesisAccount) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *GenesisAccount) GetStorage() Storage {
	if m != nil {
		return m.Storage
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "stratos.evm.v1.GenesisState")
	proto.RegisterType((*GenesisAccount)(nil), "stratos.evm.v1.GenesisAccount")
}

func init() { proto.RegisterFile("stratos/evm/v1/genesis.proto", fileDescriptor_8bfc0e5ccf7131d4) }

var fileDescriptor_8bfc0e5ccf7131d4 = []byte{
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x52, 0x4d, 0x6b, 0xd4, 0x40,
	0x18, 0xce, 0xec, 0x96, 0xae, 0x9d, 0xd5, 0x55, 0x83, 0x1f, 0xb1, 0xe8, 0xcc, 0x12, 0x44, 0x82,
	0x1f, 0x19, 0x5a, 0x0f, 0x82, 0x37, 0xe3, 0xa1, 0x27, 0x41, 0x52, 0x4f, 0x82, 0x94, 0xd9, 0x64,
	0x48, 0x83, 0x4d, 0x66, 0xc9, 0x4c, 0x83, 0xfd, 0x15, 0xfa, 0x33, 0x8a, 0x27, 0x7f, 0x46, 0x8f,
	0xbd, 0x08, 0x9e, 0x46, 0xc9, 0x1e, 0x84, 0x1c, 0xf3, 0x0b, 0xca, 0xce, 0x4c, 0xb6, 0xec, 0x5e,
	0xc2, 0xfb, 0x3e, 0xcf, 0xf3, 0xbe, 0xcf, 0xc3, 0x9b, 0x81, 0x8f, 0x85, 0xac, 0xa8, 0xe4, 0x82,
	0xb0, 0xba, 0x20, 0xf5, 0x1e, 0xc9, 0x58, 0xc9, 0x44, 0x2e, 0xc2, 0x79, 0xc5, 0x25, 0x77, 0x27,
	0x96, 0x0d, 0x59, 0x5d, 0x84, 0xf5, 0xde, 0xee, 0xbd, 0x8c, 0x67, 0x5c, 0x53, 0x64, 0x59, 0x19,
	0xd5, 0xee, 0xa3, 0x84, 0x8b, 0x82, 0x8b, 0x23, 0x43, 0x98, 0xc6, 0x52, 0x77, 0x69, 0x91, 0x97,
	0x9c, 0xe8, 0xaf, 0x85, 0xbc, 0x0d, 0xc7, 0xe5, 0x6a, 0xcd, 0xf8, 0xdf, 0x07, 0xf0, 0xe6, 0x81,
	0xf1, 0x3f, 0x94, 0x54, 0x32, 0xf7, 0x13, 0xdc, 0x9e, 0xd3, 0x8a, 0x16, 0xc2, 0x03, 0x53, 0x10,
	0x8c, 0xf7, 0x1f, 0x84, 0xeb, 0x79, 0xc2, 0x8f, 0x9a, 0x8d, 0x9e, 0x5e, 0x28, 0xec, 0xb4, 0x0a,
	0x5b, 0x75, 0xa7, 0xf0, 0xad, 0x33, 0x5a, 0x9c, 0xbc, 0xf5, 0x4d, 0xef, 0x9f, 0xff, 0xff, 0xf5,
	0x1c, 0xc4, 0x96, 0x75, 0x53, 0x78, 0x83, 0x26, 0x09, 0x3f, 0x2d, 0xa5, 0xf0, 0x06, 0xd3, 0x61,
	0x30, 0xde, 0x47, 0x9b, 0x7b, 0x6d, 0x8a, 0x77, 0x46, 0x16, 0x05, 0x76, 0xff, 0x6a, 0xae, 0x53,
	0xf8, 0xb6, 0x71, 0xe8, 0x11, 0xeb, 0xb1, 0x52, 0xb8, 0xef, 0xe1, 0xce, 0xec, 0x84, 0x27, 0x5f,
	0x8f, 0x32, 0x2a, 0xbc, 0xe1, 0x14, 0x04, 0x5b, 0xd1, 0xb3, 0x56, 0xe1, 0x6b, 0xb0, 0x53, 0xf8,
	0x8e, 0xd9, 0xb1, 0x82, 0xfa, 0x25, 0x1a, 0x38, 0xa0, 0xc2, 0xff, 0x0d, 0xe0, 0x64, 0x3d, 0x8b,
	0xfb, 0x06, 0x8e, 0x68, 0x9a, 0x56, 0x4c, 0x98, 0xa3, 0xec, 0x44, 0x4f, 0x5a, 0x85, 0x7b, 0xa8,
	0x53, 0x78, 0x62, 0x73, 0x19, 0xc0, 0x8f, 0x7b, 0xca, 0x7d, 0x01, 0xb7, 0x12, 0x9e, 0x32, 0x6f,
	0xa0, 0xa7, 0x1e, 0xb6, 0x0a, 0xeb, 0xbe, 0x53, 0x78, 0x6c, 0x46, 0x96, 0x9d, 0x1f, 0x6b, 0xd0,
	0xfd, 0x02, 0x47, 0x42, 0xf2, 0x8a, 0x66, 0xcc, 0x1b, 0xea, 0x13, 0xdd, 0xdf, 0x3c, 0x91, 0xfe,
	0x43, 0xd1, 0x4b, 0x7b, 0x99, 0x5e, 0x7d, 0x1d, 0xc0, 0x02, 0xfe, 0xcf, 0xbf, 0x78, 0x74, 0x68,
	0xea, 0xb8, 0x57, 0x45, 0x1f, 0xce, 0x1b, 0x04, 0x2e, 0x1a, 0x04, 0x2e, 0x1b, 0x04, 0xfe, 0x35,
	0x08, 0xfc, 0x58, 0x20, 0xe7, 0x72, 0x81, 0x9c, 0x3f, 0x0b, 0xe4, 0x7c, 0x26, 0x59, 0x2e, 0x8f,
	0x4f, 0x67, 0x61, 0xc2, 0x0b, 0x62, 0x5d, 0x4b, 0x26, 0xfb, 0xf2, 0x55, 0x72, 0x4c, 0xf3, 0x92,
	0x7c, 0xd3, 0xef, 0x47, 0x9e, 0xcd, 0x99, 0x98, 0x6d, 0xeb, 0xf7, 0xf3, 0xfa, 0x2a, 0x00, 0x00,
	0xff, 0xff, 0x55, 0x8f, 0x0b, 0xcd, 0xcd, 0x02, 0x00, 0x00,
}

func (this *GenesisState) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GenesisState)
	if !ok {
		that2, ok := that.(GenesisState)
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
	if !this.Params.Equal(&that1.Params) {
		return false
	}
	if len(this.Accounts) != len(that1.Accounts) {
		return false
	}
	for i := range this.Accounts {
		if !this.Accounts[i].Equal(&that1.Accounts[i]) {
			return false
		}
	}
	if this.BlockGas != that1.BlockGas {
		return false
	}
	return true
}
func (this *GenesisAccount) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GenesisAccount)
	if !ok {
		that2, ok := that.(GenesisAccount)
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
	if this.Address != that1.Address {
		return false
	}
	if this.Code != that1.Code {
		return false
	}
	if len(this.Storage) != len(that1.Storage) {
		return false
	}
	for i := range this.Storage {
		if !this.Storage[i].Equal(&that1.Storage[i]) {
			return false
		}
	}
	return true
}
func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlockGas != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.BlockGas))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Accounts) > 0 {
		for iNdEx := len(m.Accounts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Accounts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *GenesisAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Storage) > 0 {
		for iNdEx := len(m.Storage) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Storage[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Code) > 0 {
		i -= len(m.Code)
		copy(dAtA[i:], m.Code)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Code)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Accounts) > 0 {
		for _, e := range m.Accounts {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.BlockGas != 0 {
		n += 1 + sovGenesis(uint64(m.BlockGas))
	}
	return n
}

func (m *GenesisAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.Code)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.Storage) > 0 {
		for _, e := range m.Storage {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Accounts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Accounts = append(m.Accounts, GenesisAccount{})
			if err := m.Accounts[len(m.Accounts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockGas", wireType)
			}
			m.BlockGas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockGas |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *GenesisAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Code = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Storage", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Storage = append(m.Storage, State{})
			if err := m.Storage[len(m.Storage)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)