// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stratos/pot/v1/event.proto

package types

import (
	fmt "fmt"
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

// EventVolumeReport is emitted on Msg/MsgVolumeReport
type EventVolumeReport struct {
	ReportReference string `protobuf:"bytes,1,opt,name=report_reference,json=reportReference,proto3" json:"report_reference,omitempty"`
	Epoch           string `protobuf:"bytes,2,opt,name=epoch,proto3" json:"epoch,omitempty"`
}

func (m *EventVolumeReport) Reset()         { *m = EventVolumeReport{} }
func (m *EventVolumeReport) String() string { return proto.CompactTextString(m) }
func (*EventVolumeReport) ProtoMessage()    {}
func (*EventVolumeReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_4848852d1aafb62a, []int{0}
}
func (m *EventVolumeReport) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventVolumeReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventVolumeReport.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventVolumeReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventVolumeReport.Merge(m, src)
}
func (m *EventVolumeReport) XXX_Size() int {
	return m.Size()
}
func (m *EventVolumeReport) XXX_DiscardUnknown() {
	xxx_messageInfo_EventVolumeReport.DiscardUnknown(m)
}

var xxx_messageInfo_EventVolumeReport proto.InternalMessageInfo

func (m *EventVolumeReport) GetReportReference() string {
	if m != nil {
		return m.ReportReference
	}
	return ""
}

func (m *EventVolumeReport) GetEpoch() string {
	if m != nil {
		return m.Epoch
	}
	return ""
}

// EventWithdraw is emitted on Msg/MsgWithdraw
type EventWithdraw struct {
	Amount        string `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
	WalletAddress string `protobuf:"bytes,2,opt,name=wallet_address,json=walletAddress,proto3" json:"wallet_address,omitempty"`
	TargetAddress string `protobuf:"bytes,3,opt,name=target_address,json=targetAddress,proto3" json:"target_address,omitempty"`
}

func (m *EventWithdraw) Reset()         { *m = EventWithdraw{} }
func (m *EventWithdraw) String() string { return proto.CompactTextString(m) }
func (*EventWithdraw) ProtoMessage()    {}
func (*EventWithdraw) Descriptor() ([]byte, []int) {
	return fileDescriptor_4848852d1aafb62a, []int{1}
}
func (m *EventWithdraw) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventWithdraw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventWithdraw.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventWithdraw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventWithdraw.Merge(m, src)
}
func (m *EventWithdraw) XXX_Size() int {
	return m.Size()
}
func (m *EventWithdraw) XXX_DiscardUnknown() {
	xxx_messageInfo_EventWithdraw.DiscardUnknown(m)
}

var xxx_messageInfo_EventWithdraw proto.InternalMessageInfo

func (m *EventWithdraw) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *EventWithdraw) GetWalletAddress() string {
	if m != nil {
		return m.WalletAddress
	}
	return ""
}

func (m *EventWithdraw) GetTargetAddress() string {
	if m != nil {
		return m.TargetAddress
	}
	return ""
}

// EventFoundationDeposit is emitted on Msg/MsgFoundationDeposit
type EventFoundationDeposit struct {
	Amount string `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *EventFoundationDeposit) Reset()         { *m = EventFoundationDeposit{} }
func (m *EventFoundationDeposit) String() string { return proto.CompactTextString(m) }
func (*EventFoundationDeposit) ProtoMessage()    {}
func (*EventFoundationDeposit) Descriptor() ([]byte, []int) {
	return fileDescriptor_4848852d1aafb62a, []int{2}
}
func (m *EventFoundationDeposit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventFoundationDeposit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventFoundationDeposit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventFoundationDeposit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventFoundationDeposit.Merge(m, src)
}
func (m *EventFoundationDeposit) XXX_Size() int {
	return m.Size()
}
func (m *EventFoundationDeposit) XXX_DiscardUnknown() {
	xxx_messageInfo_EventFoundationDeposit.DiscardUnknown(m)
}

var xxx_messageInfo_EventFoundationDeposit proto.InternalMessageInfo

func (m *EventFoundationDeposit) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

// EventSlashing is emitted on Msg/MsgSlashingResourceNode
type EventSlashing struct {
	WalletAddress  string `protobuf:"bytes,1,opt,name=wallet_address,json=walletAddress,proto3" json:"wallet_address,omitempty"`
	NetworkAddress string `protobuf:"bytes,2,opt,name=network_address,json=networkAddress,proto3" json:"network_address,omitempty"`
	Amount         string `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	SlashingType   string `protobuf:"bytes,4,opt,name=slashing_type,json=slashingType,proto3" json:"slashing_type,omitempty"`
	Suspend        string `protobuf:"bytes,5,opt,name=suspend,proto3" json:"suspend,omitempty"`
}

func (m *EventSlashing) Reset()         { *m = EventSlashing{} }
func (m *EventSlashing) String() string { return proto.CompactTextString(m) }
func (*EventSlashing) ProtoMessage()    {}
func (*EventSlashing) Descriptor() ([]byte, []int) {
	return fileDescriptor_4848852d1aafb62a, []int{3}
}
func (m *EventSlashing) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventSlashing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventSlashing.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventSlashing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventSlashing.Merge(m, src)
}
func (m *EventSlashing) XXX_Size() int {
	return m.Size()
}
func (m *EventSlashing) XXX_DiscardUnknown() {
	xxx_messageInfo_EventSlashing.DiscardUnknown(m)
}

var xxx_messageInfo_EventSlashing proto.InternalMessageInfo

func (m *EventSlashing) GetWalletAddress() string {
	if m != nil {
		return m.WalletAddress
	}
	return ""
}

func (m *EventSlashing) GetNetworkAddress() string {
	if m != nil {
		return m.NetworkAddress
	}
	return ""
}

func (m *EventSlashing) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *EventSlashing) GetSlashingType() string {
	if m != nil {
		return m.SlashingType
	}
	return ""
}

func (m *EventSlashing) GetSuspend() string {
	if m != nil {
		return m.Suspend
	}
	return ""
}

func init() {
	proto.RegisterType((*EventVolumeReport)(nil), "stratos.pot.v1.EventVolumeReport")
	proto.RegisterType((*EventWithdraw)(nil), "stratos.pot.v1.EventWithdraw")
	proto.RegisterType((*EventFoundationDeposit)(nil), "stratos.pot.v1.EventFoundationDeposit")
	proto.RegisterType((*EventSlashing)(nil), "stratos.pot.v1.EventSlashing")
}

func init() { proto.RegisterFile("stratos/pot/v1/event.proto", fileDescriptor_4848852d1aafb62a) }

var fileDescriptor_4848852d1aafb62a = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcd, 0x4a, 0xc3, 0x40,
	0x14, 0x85, 0x3b, 0xd6, 0x56, 0x1c, 0x6c, 0xab, 0x41, 0x4a, 0x70, 0x11, 0x24, 0x22, 0xea, 0xc2,
	0xc4, 0xe2, 0x13, 0x28, 0x2a, 0xb8, 0xad, 0x45, 0xc1, 0x4d, 0x99, 0x26, 0xd7, 0x26, 0x98, 0xce,
	0x0c, 0x33, 0x37, 0xad, 0x7d, 0x0b, 0x9f, 0xc6, 0x67, 0x70, 0xd9, 0xa5, 0x4b, 0x69, 0x5f, 0x44,
	0x3a, 0x99, 0xe0, 0x0f, 0x75, 0x37, 0xf7, 0x3b, 0x87, 0x7b, 0x4e, 0xc8, 0xa5, 0x7b, 0x1a, 0x15,
	0x43, 0xa1, 0x43, 0x29, 0x30, 0x1c, 0x77, 0x42, 0x18, 0x03, 0xc7, 0x40, 0x2a, 0x81, 0xc2, 0x69,
	0x5a, 0x2d, 0x90, 0x02, 0x83, 0x71, 0xc7, 0xef, 0xd1, 0x9d, 0xeb, 0xa5, 0x7c, 0x2f, 0xb2, 0x7c,
	0x04, 0x5d, 0x90, 0x42, 0xa1, 0x73, 0x42, 0xb7, 0x95, 0x79, 0xf5, 0x15, 0x3c, 0x81, 0x02, 0x1e,
	0x81, 0x4b, 0xf6, 0xc9, 0xf1, 0x66, 0xb7, 0x55, 0xf0, 0x6e, 0x89, 0x9d, 0x5d, 0x5a, 0x03, 0x29,
	0xa2, 0xc4, 0x5d, 0x33, 0x7a, 0x31, 0xf8, 0x39, 0x6d, 0x98, 0xad, 0x0f, 0x29, 0x26, 0xb1, 0x62,
	0x13, 0xa7, 0x4d, 0xeb, 0x6c, 0x24, 0x72, 0x8e, 0x76, 0x8f, 0x9d, 0x9c, 0x43, 0xda, 0x9c, 0xb0,
	0x2c, 0x03, 0xec, 0xb3, 0x38, 0x56, 0xa0, 0xb5, 0xdd, 0xd3, 0x28, 0xe8, 0x45, 0x01, 0x97, 0x36,
	0x64, 0x6a, 0xf8, 0xc3, 0x56, 0x2d, 0x6c, 0x05, 0xb5, 0x36, 0xff, 0x8c, 0xb6, 0x4d, 0xec, 0x8d,
	0xc8, 0x79, 0xcc, 0x30, 0x15, 0xfc, 0x0a, 0xa4, 0xd0, 0x29, 0xfe, 0x97, 0xef, 0xbf, 0x11, 0xdb,
	0xf4, 0x2e, 0x63, 0x3a, 0x49, 0xf9, 0x70, 0x45, 0x23, 0xb2, 0xaa, 0xd1, 0x11, 0x6d, 0x71, 0xc0,
	0x89, 0x50, 0xcf, 0x7f, 0x9a, 0x37, 0x2d, 0x2e, 0x8d, 0xdf, 0xc9, 0xd5, 0x5f, 0x5f, 0x7e, 0x40,
	0x1b, 0xda, 0x66, 0xf6, 0x71, 0x2a, 0xc1, 0x5d, 0x37, 0xf2, 0x56, 0x09, 0x7b, 0x53, 0x09, 0x8e,
	0x4b, 0x37, 0x74, 0xae, 0x25, 0xf0, 0xd8, 0xad, 0x19, 0xb9, 0x1c, 0x2f, 0x6f, 0xdf, 0xe7, 0x1e,
	0x99, 0xcd, 0x3d, 0xf2, 0x39, 0xf7, 0xc8, 0xeb, 0xc2, 0xab, 0xcc, 0x16, 0x5e, 0xe5, 0x63, 0xe1,
	0x55, 0x1e, 0xc3, 0x61, 0x8a, 0x49, 0x3e, 0x08, 0x22, 0x31, 0x0a, 0xed, 0xcf, 0xe6, 0x80, 0xe5,
	0xf3, 0x34, 0x4a, 0x58, 0xca, 0xc3, 0x17, 0x73, 0x1b, 0xcb, 0x5c, 0x3d, 0xa8, 0x9b, 0xcb, 0x38,
	0xff, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x34, 0x7c, 0x3d, 0x37, 0x02, 0x00, 0x00,
}

func (m *EventVolumeReport) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventVolumeReport) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventVolumeReport) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Epoch) > 0 {
		i -= len(m.Epoch)
		copy(dAtA[i:], m.Epoch)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Epoch)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ReportReference) > 0 {
		i -= len(m.ReportReference)
		copy(dAtA[i:], m.ReportReference)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.ReportReference)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventWithdraw) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventWithdraw) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventWithdraw) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TargetAddress) > 0 {
		i -= len(m.TargetAddress)
		copy(dAtA[i:], m.TargetAddress)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.TargetAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.WalletAddress) > 0 {
		i -= len(m.WalletAddress)
		copy(dAtA[i:], m.WalletAddress)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.WalletAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventFoundationDeposit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventFoundationDeposit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventFoundationDeposit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventSlashing) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventSlashing) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventSlashing) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Suspend) > 0 {
		i -= len(m.Suspend)
		copy(dAtA[i:], m.Suspend)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Suspend)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.SlashingType) > 0 {
		i -= len(m.SlashingType)
		copy(dAtA[i:], m.SlashingType)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.SlashingType)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.NetworkAddress) > 0 {
		i -= len(m.NetworkAddress)
		copy(dAtA[i:], m.NetworkAddress)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.NetworkAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.WalletAddress) > 0 {
		i -= len(m.WalletAddress)
		copy(dAtA[i:], m.WalletAddress)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.WalletAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventVolumeReport) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ReportReference)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Epoch)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func (m *EventWithdraw) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.WalletAddress)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.TargetAddress)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func (m *EventFoundationDeposit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func (m *EventSlashing) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.WalletAddress)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.NetworkAddress)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.SlashingType)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Suspend)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func sovEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventVolumeReport) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventVolumeReport: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventVolumeReport: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReportReference", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReportReference = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Epoch", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Epoch = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *EventWithdraw) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventWithdraw: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventWithdraw: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WalletAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WalletAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TargetAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *EventFoundationDeposit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventFoundationDeposit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventFoundationDeposit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *EventSlashing) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventSlashing: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventSlashing: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WalletAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WalletAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NetworkAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NetworkAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashingType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SlashingType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Suspend", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Suspend = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func skipEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
				return 0, ErrInvalidLengthEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvent = fmt.Errorf("proto: unexpected end of group")
)