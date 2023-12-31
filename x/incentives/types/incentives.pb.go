// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: blockx/incentives/v1/incentives.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

// Incentive defines an instance that organizes distribution conditions for a
// given smart contract
type Incentive struct {
	// contract address
	Contract string `protobuf:"bytes,1,opt,name=contract,proto3" json:"contract,omitempty"`
	// denoms and percentage of rewards to be allocated
	Allocations github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,2,rep,name=allocations,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"allocations"`
	// number of remaining epochs
	Epochs uint32 `protobuf:"varint,3,opt,name=epochs,proto3" json:"epochs,omitempty"`
	// distribution start time
	StartTime time.Time `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time"`
	// cumulative gas spent by all gasmeters of the incentive during the epoch
	TotalGas uint64 `protobuf:"varint,5,opt,name=total_gas,json=totalGas,proto3" json:"total_gas,omitempty"`
}

func (m *Incentive) Reset()         { *m = Incentive{} }
func (m *Incentive) String() string { return proto.CompactTextString(m) }
func (*Incentive) ProtoMessage()    {}
func (*Incentive) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e41a5285327289b, []int{0}
}
func (m *Incentive) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Incentive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Incentive.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Incentive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Incentive.Merge(m, src)
}
func (m *Incentive) XXX_Size() int {
	return m.Size()
}
func (m *Incentive) XXX_DiscardUnknown() {
	xxx_messageInfo_Incentive.DiscardUnknown(m)
}

var xxx_messageInfo_Incentive proto.InternalMessageInfo

func (m *Incentive) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *Incentive) GetAllocations() github_com_cosmos_cosmos_sdk_types.DecCoins {
	if m != nil {
		return m.Allocations
	}
	return nil
}

func (m *Incentive) GetEpochs() uint32 {
	if m != nil {
		return m.Epochs
	}
	return 0
}

func (m *Incentive) GetStartTime() time.Time {
	if m != nil {
		return m.StartTime
	}
	return time.Time{}
}

func (m *Incentive) GetTotalGas() uint64 {
	if m != nil {
		return m.TotalGas
	}
	return 0
}

// GasMeter tracks the cumulative gas spent per participant in one epoch
type GasMeter struct {
	// hex address of the incentivized contract
	Contract string `protobuf:"bytes,1,opt,name=contract,proto3" json:"contract,omitempty"`
	// participant address that interacts with the incentive
	Participant string `protobuf:"bytes,2,opt,name=participant,proto3" json:"participant,omitempty"`
	// cumulative gas spent during the epoch
	CumulativeGas uint64 `protobuf:"varint,3,opt,name=cumulative_gas,json=cumulativeGas,proto3" json:"cumulative_gas,omitempty"`
}

func (m *GasMeter) Reset()         { *m = GasMeter{} }
func (m *GasMeter) String() string { return proto.CompactTextString(m) }
func (*GasMeter) ProtoMessage()    {}
func (*GasMeter) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e41a5285327289b, []int{1}
}
func (m *GasMeter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GasMeter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GasMeter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GasMeter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GasMeter.Merge(m, src)
}
func (m *GasMeter) XXX_Size() int {
	return m.Size()
}
func (m *GasMeter) XXX_DiscardUnknown() {
	xxx_messageInfo_GasMeter.DiscardUnknown(m)
}

var xxx_messageInfo_GasMeter proto.InternalMessageInfo

func (m *GasMeter) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *GasMeter) GetParticipant() string {
	if m != nil {
		return m.Participant
	}
	return ""
}

func (m *GasMeter) GetCumulativeGas() uint64 {
	if m != nil {
		return m.CumulativeGas
	}
	return 0
}

// RegisterIncentiveProposal is a gov Content type to register an incentive
type RegisterIncentiveProposal struct {
	// title of the proposal
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// proposal description
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// contract address
	Contract string `protobuf:"bytes,3,opt,name=contract,proto3" json:"contract,omitempty"`
	// denoms and percentage of rewards to be allocated
	Allocations github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,4,rep,name=allocations,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"allocations"`
	// number of remaining epochs
	Epochs uint32 `protobuf:"varint,5,opt,name=epochs,proto3" json:"epochs,omitempty"`
}

func (m *RegisterIncentiveProposal) Reset()         { *m = RegisterIncentiveProposal{} }
func (m *RegisterIncentiveProposal) String() string { return proto.CompactTextString(m) }
func (*RegisterIncentiveProposal) ProtoMessage()    {}
func (*RegisterIncentiveProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e41a5285327289b, []int{2}
}
func (m *RegisterIncentiveProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisterIncentiveProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisterIncentiveProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisterIncentiveProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterIncentiveProposal.Merge(m, src)
}
func (m *RegisterIncentiveProposal) XXX_Size() int {
	return m.Size()
}
func (m *RegisterIncentiveProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterIncentiveProposal.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterIncentiveProposal proto.InternalMessageInfo

func (m *RegisterIncentiveProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *RegisterIncentiveProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *RegisterIncentiveProposal) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *RegisterIncentiveProposal) GetAllocations() github_com_cosmos_cosmos_sdk_types.DecCoins {
	if m != nil {
		return m.Allocations
	}
	return nil
}

func (m *RegisterIncentiveProposal) GetEpochs() uint32 {
	if m != nil {
		return m.Epochs
	}
	return 0
}

// CancelIncentiveProposal is a gov Content type to cancel an incentive
type CancelIncentiveProposal struct {
	// title of the proposal
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// proposal description
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// contract address
	Contract string `protobuf:"bytes,3,opt,name=contract,proto3" json:"contract,omitempty"`
}

func (m *CancelIncentiveProposal) Reset()         { *m = CancelIncentiveProposal{} }
func (m *CancelIncentiveProposal) String() string { return proto.CompactTextString(m) }
func (*CancelIncentiveProposal) ProtoMessage()    {}
func (*CancelIncentiveProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e41a5285327289b, []int{3}
}
func (m *CancelIncentiveProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CancelIncentiveProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CancelIncentiveProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CancelIncentiveProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CancelIncentiveProposal.Merge(m, src)
}
func (m *CancelIncentiveProposal) XXX_Size() int {
	return m.Size()
}
func (m *CancelIncentiveProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_CancelIncentiveProposal.DiscardUnknown(m)
}

var xxx_messageInfo_CancelIncentiveProposal proto.InternalMessageInfo

func (m *CancelIncentiveProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CancelIncentiveProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CancelIncentiveProposal) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func init() {
	proto.RegisterType((*Incentive)(nil), "blockx.incentives.v1.Incentive")
	proto.RegisterType((*GasMeter)(nil), "blockx.incentives.v1.GasMeter")
	proto.RegisterType((*RegisterIncentiveProposal)(nil), "blockx.incentives.v1.RegisterIncentiveProposal")
	proto.RegisterType((*CancelIncentiveProposal)(nil), "blockx.incentives.v1.CancelIncentiveProposal")
}

func init() {
	proto.RegisterFile("blockx/incentives/v1/incentives.proto", fileDescriptor_5e41a5285327289b)
}

var fileDescriptor_5e41a5285327289b = []byte{
	// 505 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0x31, 0x6f, 0xd3, 0x50,
	0x10, 0xc7, 0xe3, 0x24, 0xad, 0x92, 0x17, 0x95, 0xc1, 0x8a, 0xc0, 0x04, 0xe4, 0x58, 0x91, 0x2a,
	0x45, 0x42, 0xf1, 0x53, 0xda, 0x01, 0xc1, 0x98, 0x20, 0x55, 0x0c, 0x48, 0xc8, 0x30, 0xb1, 0x54,
	0xcf, 0x2f, 0x57, 0xf7, 0xa9, 0x8e, 0xcf, 0xf2, 0x5d, 0xa2, 0xb2, 0xf2, 0x09, 0x3a, 0x31, 0x33,
	0xf3, 0x49, 0x3a, 0x76, 0x64, 0xa2, 0x28, 0x59, 0xf8, 0x18, 0xc8, 0x76, 0x5c, 0xdc, 0xa5, 0x63,
	0x27, 0xfb, 0x7f, 0x77, 0xef, 0xfd, 0xee, 0xfe, 0x67, 0x8b, 0xc3, 0x30, 0x46, 0x7d, 0x71, 0x29,
	0x4d, 0xa2, 0x21, 0x61, 0xb3, 0x06, 0x92, 0xeb, 0x69, 0x4d, 0xf9, 0x69, 0x86, 0x8c, 0x76, 0xbf,
	0x2c, 0xf3, 0x6b, 0x89, 0xf5, 0x74, 0xd0, 0x8f, 0x30, 0xc2, 0xa2, 0x40, 0xe6, 0x6f, 0x65, 0xed,
	0x60, 0x18, 0x21, 0x46, 0x31, 0xc8, 0x42, 0x85, 0xab, 0x33, 0xc9, 0x66, 0x09, 0xc4, 0x6a, 0x99,
	0xee, 0x0a, 0x5c, 0x8d, 0xb4, 0x44, 0x92, 0xa1, 0x22, 0x90, 0xeb, 0x69, 0x08, 0xac, 0xa6, 0x52,
	0xa3, 0x49, 0xca, 0xfc, 0xe8, 0x7b, 0x53, 0x74, 0xdf, 0x57, 0x20, 0x7b, 0x20, 0x3a, 0x1a, 0x13,
	0xce, 0x94, 0x66, 0xc7, 0xf2, 0xac, 0x71, 0x37, 0xb8, 0xd3, 0x36, 0x89, 0x9e, 0x8a, 0x63, 0xd4,
	0x8a, 0x0d, 0x26, 0xe4, 0x34, 0xbd, 0xd6, 0xb8, 0x77, 0xf4, 0xd2, 0x2f, 0xef, 0xf7, 0xf3, 0xfb,
	0xfd, 0xdd, 0xfd, 0xfe, 0x3b, 0xd0, 0x73, 0x34, 0xc9, 0xec, 0xf8, 0xfa, 0xf7, 0xb0, 0xf1, 0xf3,
	0x76, 0xf8, 0x2a, 0x32, 0x7c, 0xbe, 0x0a, 0x7d, 0x8d, 0x4b, 0xb9, 0xeb, 0xa7, 0x7c, 0x4c, 0x68,
	0x71, 0x21, 0xf9, 0x6b, 0x0a, 0x54, 0x9d, 0xa1, 0xa0, 0x4e, 0xb1, 0x9f, 0x8a, 0x7d, 0x48, 0x51,
	0x9f, 0x93, 0xd3, 0xf2, 0xac, 0xf1, 0x41, 0xb0, 0x53, 0xf6, 0x5c, 0x08, 0x62, 0x95, 0xf1, 0x69,
	0x3e, 0xaf, 0xd3, 0xf6, 0xac, 0x71, 0xef, 0x68, 0xe0, 0x97, 0x66, 0xf8, 0x95, 0x19, 0xfe, 0xe7,
	0xca, 0x8c, 0x59, 0x27, 0xef, 0xe4, 0xea, 0x76, 0x68, 0x05, 0xdd, 0xe2, 0x5c, 0x9e, 0xb1, 0x5f,
	0x88, 0x2e, 0x23, 0xab, 0xf8, 0x34, 0x52, 0xe4, 0xec, 0x79, 0xd6, 0xb8, 0x1d, 0x74, 0x8a, 0xc0,
	0x89, 0xa2, 0x11, 0x8a, 0xce, 0x89, 0xa2, 0x0f, 0xc0, 0x90, 0x3d, 0x68, 0x8b, 0x27, 0x7a, 0xa9,
	0xca, 0xd8, 0x68, 0x93, 0xaa, 0x84, 0x9d, 0x66, 0x91, 0xae, 0x87, 0xec, 0x43, 0xf1, 0x44, 0xaf,
	0x96, 0xab, 0x58, 0xe5, 0x16, 0x17, 0xac, 0x56, 0xc1, 0x3a, 0xf8, 0x1f, 0xcd, 0x81, 0xdf, 0x9a,
	0xe2, 0x79, 0x00, 0x91, 0x21, 0x86, 0xec, 0x6e, 0x23, 0x1f, 0x33, 0x4c, 0x91, 0x54, 0x6c, 0xf7,
	0xc5, 0x1e, 0x1b, 0x8e, 0x61, 0xc7, 0x2f, 0x45, 0x0e, 0x5f, 0x00, 0xe9, 0xcc, 0xa4, 0xb9, 0x5d,
	0x15, 0xbc, 0x16, 0xba, 0xd7, 0x7a, 0xeb, 0xe1, 0x8d, 0xb6, 0x1f, 0x79, 0xa3, 0x7b, 0xf5, 0x8d,
	0xbe, 0x6d, 0xff, 0xfd, 0x31, 0x6c, 0x8c, 0x48, 0x3c, 0x9b, 0xab, 0x44, 0x43, 0xfc, 0x28, 0x0e,
	0x94, 0xd0, 0xd9, 0xa7, 0xeb, 0x8d, 0x6b, 0xdd, 0x6c, 0x5c, 0xeb, 0xcf, 0xc6, 0xb5, 0xae, 0xb6,
	0x6e, 0xe3, 0x66, 0xeb, 0x36, 0x7e, 0x6d, 0xdd, 0xc6, 0x97, 0x37, 0xb5, 0x31, 0x17, 0x70, 0x66,
	0x26, 0x6b, 0x48, 0x78, 0x95, 0x01, 0xc9, 0x50, 0x5f, 0x4e, 0x18, 0x88, 0x13, 0xe0, 0xc9, 0x6b,
	0x79, 0xef, 0x9f, 0x2e, 0xa6, 0x0f, 0xf7, 0x8b, 0xaf, 0xf0, 0xf8, 0x5f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xbe, 0x5f, 0x86, 0x44, 0xf5, 0x03, 0x00, 0x00,
}

func (m *Incentive) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Incentive) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Incentive) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TotalGas != 0 {
		i = encodeVarintIncentives(dAtA, i, uint64(m.TotalGas))
		i--
		dAtA[i] = 0x28
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintIncentives(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	if m.Epochs != 0 {
		i = encodeVarintIncentives(dAtA, i, uint64(m.Epochs))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Allocations) > 0 {
		for iNdEx := len(m.Allocations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Allocations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintIncentives(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintIncentives(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GasMeter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GasMeter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GasMeter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CumulativeGas != 0 {
		i = encodeVarintIncentives(dAtA, i, uint64(m.CumulativeGas))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Participant) > 0 {
		i -= len(m.Participant)
		copy(dAtA[i:], m.Participant)
		i = encodeVarintIncentives(dAtA, i, uint64(len(m.Participant)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintIncentives(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RegisterIncentiveProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterIncentiveProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegisterIncentiveProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Epochs != 0 {
		i = encodeVarintIncentives(dAtA, i, uint64(m.Epochs))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Allocations) > 0 {
		for iNdEx := len(m.Allocations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Allocations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintIncentives(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintIncentives(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintIncentives(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintIncentives(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CancelIncentiveProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CancelIncentiveProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CancelIncentiveProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintIncentives(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintIncentives(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintIncentives(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintIncentives(dAtA []byte, offset int, v uint64) int {
	offset -= sovIncentives(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Incentive) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovIncentives(uint64(l))
	}
	if len(m.Allocations) > 0 {
		for _, e := range m.Allocations {
			l = e.Size()
			n += 1 + l + sovIncentives(uint64(l))
		}
	}
	if m.Epochs != 0 {
		n += 1 + sovIncentives(uint64(m.Epochs))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovIncentives(uint64(l))
	if m.TotalGas != 0 {
		n += 1 + sovIncentives(uint64(m.TotalGas))
	}
	return n
}

func (m *GasMeter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovIncentives(uint64(l))
	}
	l = len(m.Participant)
	if l > 0 {
		n += 1 + l + sovIncentives(uint64(l))
	}
	if m.CumulativeGas != 0 {
		n += 1 + sovIncentives(uint64(m.CumulativeGas))
	}
	return n
}

func (m *RegisterIncentiveProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovIncentives(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovIncentives(uint64(l))
	}
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovIncentives(uint64(l))
	}
	if len(m.Allocations) > 0 {
		for _, e := range m.Allocations {
			l = e.Size()
			n += 1 + l + sovIncentives(uint64(l))
		}
	}
	if m.Epochs != 0 {
		n += 1 + sovIncentives(uint64(m.Epochs))
	}
	return n
}

func (m *CancelIncentiveProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovIncentives(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovIncentives(uint64(l))
	}
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovIncentives(uint64(l))
	}
	return n
}

func sovIncentives(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIncentives(x uint64) (n int) {
	return sovIncentives(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Incentive) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIncentives
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
			return fmt.Errorf("proto: Incentive: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Incentive: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentives
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
				return ErrInvalidLengthIncentives
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentives
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Allocations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentives
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
				return ErrInvalidLengthIncentives
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIncentives
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Allocations = append(m.Allocations, types.DecCoin{})
			if err := m.Allocations[len(m.Allocations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Epochs", wireType)
			}
			m.Epochs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentives
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Epochs |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentives
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
				return ErrInvalidLengthIncentives
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIncentives
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalGas", wireType)
			}
			m.TotalGas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentives
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalGas |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipIncentives(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIncentives
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
func (m *GasMeter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIncentives
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
			return fmt.Errorf("proto: GasMeter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GasMeter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentives
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
				return ErrInvalidLengthIncentives
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentives
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Participant", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentives
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
				return ErrInvalidLengthIncentives
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentives
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Participant = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CumulativeGas", wireType)
			}
			m.CumulativeGas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentives
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CumulativeGas |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipIncentives(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIncentives
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
func (m *RegisterIncentiveProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIncentives
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
			return fmt.Errorf("proto: RegisterIncentiveProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterIncentiveProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentives
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
				return ErrInvalidLengthIncentives
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentives
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentives
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
				return ErrInvalidLengthIncentives
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentives
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentives
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
				return ErrInvalidLengthIncentives
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentives
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Allocations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentives
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
				return ErrInvalidLengthIncentives
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIncentives
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Allocations = append(m.Allocations, types.DecCoin{})
			if err := m.Allocations[len(m.Allocations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Epochs", wireType)
			}
			m.Epochs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentives
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Epochs |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipIncentives(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIncentives
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
func (m *CancelIncentiveProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIncentives
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
			return fmt.Errorf("proto: CancelIncentiveProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CancelIncentiveProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentives
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
				return ErrInvalidLengthIncentives
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentives
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentives
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
				return ErrInvalidLengthIncentives
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentives
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentives
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
				return ErrInvalidLengthIncentives
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentives
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIncentives(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIncentives
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
func skipIncentives(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIncentives
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
					return 0, ErrIntOverflowIncentives
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
					return 0, ErrIntOverflowIncentives
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
				return 0, ErrInvalidLengthIncentives
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIncentives
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIncentives
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIncentives        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIncentives          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIncentives = fmt.Errorf("proto: unexpected end of group")
)
