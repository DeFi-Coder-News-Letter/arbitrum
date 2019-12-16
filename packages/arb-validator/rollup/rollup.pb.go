// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rollup.proto

package rollup

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type ChainBuf struct {
	ContractAddress      string          `protobuf:"bytes,1,opt,name=contractAddress,proto3" json:"contractAddress,omitempty"`
	VmParams             *VMParamsBuf    `protobuf:"bytes,2,opt,name=vmParams,proto3" json:"vmParams,omitempty"`
	Nodes                []*NodeBuf      `protobuf:"bytes,3,rep,name=nodes,proto3" json:"nodes,omitempty"`
	LatestConfirmedHash  string          `protobuf:"bytes,4,opt,name=latestConfirmedHash,proto3" json:"latestConfirmedHash,omitempty"`
	Stakers              []*StakerBuf    `protobuf:"bytes,5,rep,name=stakers,proto3" json:"stakers,omitempty"`
	Challenges           []*ChallengeBuf `protobuf:"bytes,6,rep,name=challenges,proto3" json:"challenges,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ChainBuf) Reset()         { *m = ChainBuf{} }
func (m *ChainBuf) String() string { return proto.CompactTextString(m) }
func (*ChainBuf) ProtoMessage()    {}
func (*ChainBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_037f188b50610c79, []int{0}
}

func (m *ChainBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChainBuf.Unmarshal(m, b)
}
func (m *ChainBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChainBuf.Marshal(b, m, deterministic)
}
func (m *ChainBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainBuf.Merge(m, src)
}
func (m *ChainBuf) XXX_Size() int {
	return xxx_messageInfo_ChainBuf.Size(m)
}
func (m *ChainBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainBuf.DiscardUnknown(m)
}

var xxx_messageInfo_ChainBuf proto.InternalMessageInfo

func (m *ChainBuf) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *ChainBuf) GetVmParams() *VMParamsBuf {
	if m != nil {
		return m.VmParams
	}
	return nil
}

func (m *ChainBuf) GetNodes() []*NodeBuf {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func (m *ChainBuf) GetLatestConfirmedHash() string {
	if m != nil {
		return m.LatestConfirmedHash
	}
	return ""
}

func (m *ChainBuf) GetStakers() []*StakerBuf {
	if m != nil {
		return m.Stakers
	}
	return nil
}

func (m *ChainBuf) GetChallenges() []*ChallengeBuf {
	if m != nil {
		return m.Challenges
	}
	return nil
}

type VMParamsBuf struct {
	StakeRequirement     string           `protobuf:"bytes,1,opt,name=stakeRequirement,proto3" json:"stakeRequirement,omitempty"`
	GracePeriod          uint32           `protobuf:"varint,2,opt,name=gracePeriod,proto3" json:"gracePeriod,omitempty"`
	MaxExecutionSteps    uint32           `protobuf:"varint,3,opt,name=maxExecutionSteps,proto3" json:"maxExecutionSteps,omitempty"`
	PendingInbox         *PendingInboxBuf `protobuf:"bytes,4,opt,name=pendingInbox,proto3" json:"pendingInbox,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *VMParamsBuf) Reset()         { *m = VMParamsBuf{} }
func (m *VMParamsBuf) String() string { return proto.CompactTextString(m) }
func (*VMParamsBuf) ProtoMessage()    {}
func (*VMParamsBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_037f188b50610c79, []int{1}
}

func (m *VMParamsBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VMParamsBuf.Unmarshal(m, b)
}
func (m *VMParamsBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VMParamsBuf.Marshal(b, m, deterministic)
}
func (m *VMParamsBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VMParamsBuf.Merge(m, src)
}
func (m *VMParamsBuf) XXX_Size() int {
	return xxx_messageInfo_VMParamsBuf.Size(m)
}
func (m *VMParamsBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_VMParamsBuf.DiscardUnknown(m)
}

var xxx_messageInfo_VMParamsBuf proto.InternalMessageInfo

func (m *VMParamsBuf) GetStakeRequirement() string {
	if m != nil {
		return m.StakeRequirement
	}
	return ""
}

func (m *VMParamsBuf) GetGracePeriod() uint32 {
	if m != nil {
		return m.GracePeriod
	}
	return 0
}

func (m *VMParamsBuf) GetMaxExecutionSteps() uint32 {
	if m != nil {
		return m.MaxExecutionSteps
	}
	return 0
}

func (m *VMParamsBuf) GetPendingInbox() *PendingInboxBuf {
	if m != nil {
		return m.PendingInbox
	}
	return nil
}

type NodeBuf struct {
	DisputableNode       *DisputableNodeBuf `protobuf:"bytes,1,opt,name=disputableNode,proto3" json:"disputableNode,omitempty"`
	MachineHash          string             `protobuf:"bytes,2,opt,name=machineHash,proto3" json:"machineHash,omitempty"`
	InboxHash            string             `protobuf:"bytes,3,opt,name=inboxHash,proto3" json:"inboxHash,omitempty"`
	LinkType             uint32             `protobuf:"varint,4,opt,name=linkType,proto3" json:"linkType,omitempty"`
	PrevHash             string             `protobuf:"bytes,5,opt,name=prevHash,proto3" json:"prevHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *NodeBuf) Reset()         { *m = NodeBuf{} }
func (m *NodeBuf) String() string { return proto.CompactTextString(m) }
func (*NodeBuf) ProtoMessage()    {}
func (*NodeBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_037f188b50610c79, []int{2}
}

func (m *NodeBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeBuf.Unmarshal(m, b)
}
func (m *NodeBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeBuf.Marshal(b, m, deterministic)
}
func (m *NodeBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeBuf.Merge(m, src)
}
func (m *NodeBuf) XXX_Size() int {
	return xxx_messageInfo_NodeBuf.Size(m)
}
func (m *NodeBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeBuf.DiscardUnknown(m)
}

var xxx_messageInfo_NodeBuf proto.InternalMessageInfo

func (m *NodeBuf) GetDisputableNode() *DisputableNodeBuf {
	if m != nil {
		return m.DisputableNode
	}
	return nil
}

func (m *NodeBuf) GetMachineHash() string {
	if m != nil {
		return m.MachineHash
	}
	return ""
}

func (m *NodeBuf) GetInboxHash() string {
	if m != nil {
		return m.InboxHash
	}
	return ""
}

func (m *NodeBuf) GetLinkType() uint32 {
	if m != nil {
		return m.LinkType
	}
	return 0
}

func (m *NodeBuf) GetPrevHash() string {
	if m != nil {
		return m.PrevHash
	}
	return ""
}

type StakerBuf struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Location             string   `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	CreationTime         string   `protobuf:"bytes,3,opt,name=creationTime,proto3" json:"creationTime,omitempty"`
	ChallengeAddr        string   `protobuf:"bytes,4,opt,name=challengeAddr,proto3" json:"challengeAddr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StakerBuf) Reset()         { *m = StakerBuf{} }
func (m *StakerBuf) String() string { return proto.CompactTextString(m) }
func (*StakerBuf) ProtoMessage()    {}
func (*StakerBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_037f188b50610c79, []int{3}
}

func (m *StakerBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StakerBuf.Unmarshal(m, b)
}
func (m *StakerBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StakerBuf.Marshal(b, m, deterministic)
}
func (m *StakerBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StakerBuf.Merge(m, src)
}
func (m *StakerBuf) XXX_Size() int {
	return xxx_messageInfo_StakerBuf.Size(m)
}
func (m *StakerBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_StakerBuf.DiscardUnknown(m)
}

var xxx_messageInfo_StakerBuf proto.InternalMessageInfo

func (m *StakerBuf) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *StakerBuf) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *StakerBuf) GetCreationTime() string {
	if m != nil {
		return m.CreationTime
	}
	return ""
}

func (m *StakerBuf) GetChallengeAddr() string {
	if m != nil {
		return m.ChallengeAddr
	}
	return ""
}

type ChallengeBuf struct {
	Contract             string   `protobuf:"bytes,1,opt,name=contract,proto3" json:"contract,omitempty"`
	Asserter             string   `protobuf:"bytes,2,opt,name=asserter,proto3" json:"asserter,omitempty"`
	Challenger           string   `protobuf:"bytes,3,opt,name=challenger,proto3" json:"challenger,omitempty"`
	Kind                 uint32   `protobuf:"varint,4,opt,name=kind,proto3" json:"kind,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChallengeBuf) Reset()         { *m = ChallengeBuf{} }
func (m *ChallengeBuf) String() string { return proto.CompactTextString(m) }
func (*ChallengeBuf) ProtoMessage()    {}
func (*ChallengeBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_037f188b50610c79, []int{4}
}

func (m *ChallengeBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChallengeBuf.Unmarshal(m, b)
}
func (m *ChallengeBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChallengeBuf.Marshal(b, m, deterministic)
}
func (m *ChallengeBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChallengeBuf.Merge(m, src)
}
func (m *ChallengeBuf) XXX_Size() int {
	return xxx_messageInfo_ChallengeBuf.Size(m)
}
func (m *ChallengeBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_ChallengeBuf.DiscardUnknown(m)
}

var xxx_messageInfo_ChallengeBuf proto.InternalMessageInfo

func (m *ChallengeBuf) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *ChallengeBuf) GetAsserter() string {
	if m != nil {
		return m.Asserter
	}
	return ""
}

func (m *ChallengeBuf) GetChallenger() string {
	if m != nil {
		return m.Challenger
	}
	return ""
}

func (m *ChallengeBuf) GetKind() uint32 {
	if m != nil {
		return m.Kind
	}
	return 0
}

type DisputableNodeBuf struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DisputableNodeBuf) Reset()         { *m = DisputableNodeBuf{} }
func (m *DisputableNodeBuf) String() string { return proto.CompactTextString(m) }
func (*DisputableNodeBuf) ProtoMessage()    {}
func (*DisputableNodeBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_037f188b50610c79, []int{5}
}

func (m *DisputableNodeBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DisputableNodeBuf.Unmarshal(m, b)
}
func (m *DisputableNodeBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DisputableNodeBuf.Marshal(b, m, deterministic)
}
func (m *DisputableNodeBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisputableNodeBuf.Merge(m, src)
}
func (m *DisputableNodeBuf) XXX_Size() int {
	return xxx_messageInfo_DisputableNodeBuf.Size(m)
}
func (m *DisputableNodeBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_DisputableNodeBuf.DiscardUnknown(m)
}

var xxx_messageInfo_DisputableNodeBuf proto.InternalMessageInfo

type PendingInboxBuf struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PendingInboxBuf) Reset()         { *m = PendingInboxBuf{} }
func (m *PendingInboxBuf) String() string { return proto.CompactTextString(m) }
func (*PendingInboxBuf) ProtoMessage()    {}
func (*PendingInboxBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_037f188b50610c79, []int{6}
}

func (m *PendingInboxBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PendingInboxBuf.Unmarshal(m, b)
}
func (m *PendingInboxBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PendingInboxBuf.Marshal(b, m, deterministic)
}
func (m *PendingInboxBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PendingInboxBuf.Merge(m, src)
}
func (m *PendingInboxBuf) XXX_Size() int {
	return xxx_messageInfo_PendingInboxBuf.Size(m)
}
func (m *PendingInboxBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_PendingInboxBuf.DiscardUnknown(m)
}

var xxx_messageInfo_PendingInboxBuf proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ChainBuf)(nil), "rollup.ChainBuf")
	proto.RegisterType((*VMParamsBuf)(nil), "rollup.VMParamsBuf")
	proto.RegisterType((*NodeBuf)(nil), "rollup.NodeBuf")
	proto.RegisterType((*StakerBuf)(nil), "rollup.StakerBuf")
	proto.RegisterType((*ChallengeBuf)(nil), "rollup.ChallengeBuf")
	proto.RegisterType((*DisputableNodeBuf)(nil), "rollup.DisputableNodeBuf")
	proto.RegisterType((*PendingInboxBuf)(nil), "rollup.PendingInboxBuf")
}

func init() { proto.RegisterFile("rollup.proto", fileDescriptor_037f188b50610c79) }

var fileDescriptor_037f188b50610c79 = []byte{
	// 556 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x54, 0x5d, 0xaf, 0xd2, 0x40,
	0x10, 0x4d, 0x2f, 0x97, 0xaf, 0x29, 0x88, 0x2c, 0x26, 0x56, 0x63, 0x0c, 0x69, 0x34, 0x21, 0x7e,
	0x80, 0x41, 0xdf, 0x4c, 0x8c, 0x17, 0x34, 0xd1, 0x07, 0x0d, 0xe1, 0xde, 0xf8, 0xe0, 0xdb, 0xb6,
	0x1d, 0x60, 0x43, 0xbb, 0x5b, 0x77, 0xb7, 0x04, 0xfd, 0x0b, 0xfe, 0x0a, 0xff, 0x8b, 0xf1, 0x77,
	0x99, 0x5d, 0xda, 0x5a, 0xc0, 0xb7, 0xce, 0x39, 0x67, 0xa7, 0x33, 0x67, 0x4f, 0x0b, 0x1d, 0x29,
	0xe2, 0x38, 0x4b, 0xc7, 0xa9, 0x14, 0x5a, 0x90, 0xc6, 0xa1, 0xf2, 0x7f, 0x5d, 0x40, 0x6b, 0xbe,
	0xa1, 0x8c, 0xcf, 0xb2, 0x15, 0x19, 0x41, 0x2f, 0x14, 0x5c, 0x4b, 0x1a, 0xea, 0xab, 0x28, 0x92,
	0xa8, 0x94, 0xe7, 0x0c, 0x9d, 0x51, 0x7b, 0x79, 0x0a, 0x93, 0x09, 0xb4, 0x76, 0xc9, 0x82, 0x4a,
	0x9a, 0x28, 0xef, 0x62, 0xe8, 0x8c, 0xdc, 0xe9, 0x60, 0x9c, 0xf7, 0xff, 0xf2, 0xe9, 0x80, 0xcf,
	0xb2, 0xd5, 0xb2, 0x14, 0x91, 0xc7, 0x50, 0xe7, 0x22, 0x42, 0xe5, 0xd5, 0x86, 0xb5, 0x91, 0x3b,
	0xed, 0x15, 0xea, 0xcf, 0x22, 0x42, 0xa3, 0x3c, 0xb0, 0xe4, 0x05, 0x0c, 0x62, 0xaa, 0x51, 0xe9,
	0xb9, 0xe0, 0x2b, 0x26, 0x13, 0x8c, 0x3e, 0x50, 0xb5, 0xf1, 0x2e, 0xed, 0x14, 0xff, 0xa3, 0xc8,
	0x53, 0x68, 0x2a, 0x4d, 0xb7, 0x28, 0x95, 0x57, 0xb7, 0xad, 0xfb, 0x45, 0xeb, 0x6b, 0x0b, 0x9b,
	0xe6, 0x85, 0x82, 0xbc, 0x02, 0x08, 0x37, 0x34, 0x8e, 0x91, 0xaf, 0x51, 0x79, 0x0d, 0xab, 0xbf,
	0x53, 0xe8, 0xe7, 0x05, 0x63, 0x8e, 0x54, 0x74, 0xfe, 0x1f, 0x07, 0xdc, 0xca, 0x56, 0xe4, 0x09,
	0xdc, 0xb6, 0x0d, 0x97, 0xf8, 0x2d, 0x63, 0x12, 0x13, 0xe4, 0x3a, 0xf7, 0xe9, 0x0c, 0x27, 0x43,
	0x70, 0xd7, 0x92, 0x86, 0xb8, 0x40, 0xc9, 0x44, 0x64, 0xbd, 0xea, 0x2e, 0xab, 0x10, 0x79, 0x06,
	0xfd, 0x84, 0xee, 0xdf, 0xef, 0x31, 0xcc, 0x34, 0x13, 0xfc, 0x5a, 0x63, 0x6a, 0x5c, 0x32, 0xba,
	0x73, 0x82, 0xbc, 0x86, 0x4e, 0x8a, 0x3c, 0x62, 0x7c, 0xfd, 0x91, 0x07, 0x62, 0x6f, 0x9d, 0x71,
	0xa7, 0x77, 0x8b, 0x1d, 0x16, 0x15, 0xce, 0xac, 0x71, 0x24, 0xf6, 0x7f, 0x3b, 0xd0, 0xcc, 0x0d,
	0x27, 0x57, 0x70, 0x2b, 0x62, 0x2a, 0xcd, 0x34, 0x0d, 0x62, 0x34, 0xa0, 0x5d, 0xc1, 0x9d, 0xde,
	0x2b, 0x5a, 0xbd, 0x3b, 0x62, 0x4d, 0xb3, 0x93, 0x03, 0x66, 0xb7, 0x84, 0x86, 0x1b, 0xc6, 0xd1,
	0x5e, 0xd2, 0x85, 0xb5, 0xa0, 0x0a, 0x91, 0x07, 0xd0, 0x66, 0xe6, 0xcd, 0x96, 0xaf, 0x59, 0xfe,
	0x1f, 0x40, 0xee, 0x43, 0x2b, 0x66, 0x7c, 0x7b, 0xf3, 0x3d, 0x45, 0xbb, 0x47, 0x77, 0x59, 0xd6,
	0x86, 0x4b, 0x25, 0xee, 0xec, 0xc1, 0xba, 0x3d, 0x58, 0xd6, 0xfe, 0x4f, 0x07, 0xda, 0xe5, 0xe5,
	0x12, 0x0f, 0x9a, 0xf4, 0x28, 0xac, 0x45, 0x69, 0xfb, 0x8b, 0x90, 0x1a, 0xf3, 0xf2, 0xe1, 0xca,
	0x9a, 0xf8, 0xd0, 0x09, 0x25, 0xda, 0xe7, 0x1b, 0x96, 0x60, 0x3e, 0xdc, 0x11, 0x46, 0x1e, 0x41,
	0xb7, 0x4c, 0x81, 0x09, 0x7e, 0x1e, 0xc3, 0x63, 0xd0, 0xff, 0x01, 0x9d, 0x6a, 0x72, 0xcc, 0x5b,
	0x8b, 0xaf, 0x25, 0x1f, 0xa8, 0xac, 0x0d, 0x47, 0x95, 0x42, 0xa9, 0x51, 0x16, 0x13, 0x15, 0x35,
	0x79, 0x58, 0xc9, 0xa6, 0xcc, 0xe7, 0xa9, 0x20, 0x84, 0xc0, 0xe5, 0x96, 0xf1, 0x28, 0x77, 0xca,
	0x3e, 0xfb, 0x03, 0xe8, 0x9f, 0x5d, 0x93, 0xdf, 0x87, 0xde, 0x49, 0x0c, 0x66, 0x6f, 0xbf, 0xbe,
	0x59, 0x33, 0xbd, 0xc9, 0x82, 0x71, 0x28, 0x92, 0x89, 0x58, 0xad, 0x42, 0xf3, 0xc9, 0xc7, 0x34,
	0x50, 0x13, 0x2a, 0x03, 0xa6, 0x65, 0x96, 0x4c, 0x52, 0x1a, 0x6e, 0xe9, 0x1a, 0x2d, 0xf2, 0x7c,
	0x47, 0x63, 0x16, 0x51, 0x2d, 0xe4, 0xe4, 0x90, 0x86, 0xa0, 0x61, 0x7f, 0x1b, 0x2f, 0xff, 0x06,
	0x00, 0x00, 0xff, 0xff, 0xfc, 0xc5, 0x2b, 0x8c, 0x46, 0x04, 0x00, 0x00,
}