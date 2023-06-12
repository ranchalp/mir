// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: trantorpb/trantorpb.proto

package trantorpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	_ "github.com/filecoin-project/mir/pkg/pb/mir"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	TxNo     uint64 `protobuf:"varint,2,opt,name=tx_no,json=txNo,proto3" json:"tx_no,omitempty"`
	Type     uint64 `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"` // TODO not being used at the moment? (only some boilerplate code in pkg/serializing)
	Data     []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trantorpb_trantorpb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_trantorpb_trantorpb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_trantorpb_trantorpb_proto_rawDescGZIP(), []int{0}
}

func (x *Transaction) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *Transaction) GetTxNo() uint64 {
	if x != nil {
		return x.TxNo
	}
	return 0
}

func (x *Transaction) GetType() uint64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Transaction) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type StateSnapshot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppData   []byte     `protobuf:"bytes,1,opt,name=app_data,json=appData,proto3" json:"app_data,omitempty"`
	EpochData *EpochData `protobuf:"bytes,2,opt,name=epoch_data,json=epochData,proto3" json:"epoch_data,omitempty"`
}

func (x *StateSnapshot) Reset() {
	*x = StateSnapshot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trantorpb_trantorpb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateSnapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateSnapshot) ProtoMessage() {}

func (x *StateSnapshot) ProtoReflect() protoreflect.Message {
	mi := &file_trantorpb_trantorpb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateSnapshot.ProtoReflect.Descriptor instead.
func (*StateSnapshot) Descriptor() ([]byte, []int) {
	return file_trantorpb_trantorpb_proto_rawDescGZIP(), []int{1}
}

func (x *StateSnapshot) GetAppData() []byte {
	if x != nil {
		return x.AppData
	}
	return nil
}

func (x *StateSnapshot) GetEpochData() *EpochData {
	if x != nil {
		return x.EpochData
	}
	return nil
}

type EpochData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EpochConfig        *EpochConfig    `protobuf:"bytes,1,opt,name=epoch_config,json=epochConfig,proto3" json:"epoch_config,omitempty"`
	ClientProgress     *ClientProgress `protobuf:"bytes,2,opt,name=client_progress,json=clientProgress,proto3" json:"client_progress,omitempty"`
	LeaderPolicy       []byte          `protobuf:"bytes,3,opt,name=leader_policy,json=leaderPolicy,proto3" json:"leader_policy,omitempty"`
	PreviousMembership *Membership     `protobuf:"bytes,4,opt,name=previous_membership,json=previousMembership,proto3" json:"previous_membership,omitempty"`
}

func (x *EpochData) Reset() {
	*x = EpochData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trantorpb_trantorpb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EpochData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EpochData) ProtoMessage() {}

func (x *EpochData) ProtoReflect() protoreflect.Message {
	mi := &file_trantorpb_trantorpb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EpochData.ProtoReflect.Descriptor instead.
func (*EpochData) Descriptor() ([]byte, []int) {
	return file_trantorpb_trantorpb_proto_rawDescGZIP(), []int{2}
}

func (x *EpochData) GetEpochConfig() *EpochConfig {
	if x != nil {
		return x.EpochConfig
	}
	return nil
}

func (x *EpochData) GetClientProgress() *ClientProgress {
	if x != nil {
		return x.ClientProgress
	}
	return nil
}

func (x *EpochData) GetLeaderPolicy() []byte {
	if x != nil {
		return x.LeaderPolicy
	}
	return nil
}

func (x *EpochData) GetPreviousMembership() *Membership {
	if x != nil {
		return x.PreviousMembership
	}
	return nil
}

type EpochConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EpochNr     uint64        `protobuf:"varint,1,opt,name=epoch_nr,json=epochNr,proto3" json:"epoch_nr,omitempty"`
	FirstSn     uint64        `protobuf:"varint,2,opt,name=first_sn,json=firstSn,proto3" json:"first_sn,omitempty"`
	Length      uint64        `protobuf:"varint,3,opt,name=length,proto3" json:"length,omitempty"`
	Memberships []*Membership `protobuf:"bytes,4,rep,name=memberships,proto3" json:"memberships,omitempty"`
}

func (x *EpochConfig) Reset() {
	*x = EpochConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trantorpb_trantorpb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EpochConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EpochConfig) ProtoMessage() {}

func (x *EpochConfig) ProtoReflect() protoreflect.Message {
	mi := &file_trantorpb_trantorpb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EpochConfig.ProtoReflect.Descriptor instead.
func (*EpochConfig) Descriptor() ([]byte, []int) {
	return file_trantorpb_trantorpb_proto_rawDescGZIP(), []int{3}
}

func (x *EpochConfig) GetEpochNr() uint64 {
	if x != nil {
		return x.EpochNr
	}
	return 0
}

func (x *EpochConfig) GetFirstSn() uint64 {
	if x != nil {
		return x.FirstSn
	}
	return 0
}

func (x *EpochConfig) GetLength() uint64 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *EpochConfig) GetMemberships() []*Membership {
	if x != nil {
		return x.Memberships
	}
	return nil
}

type Membership struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes map[string]*NodeIdentity `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Membership) Reset() {
	*x = Membership{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trantorpb_trantorpb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Membership) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Membership) ProtoMessage() {}

func (x *Membership) ProtoReflect() protoreflect.Message {
	mi := &file_trantorpb_trantorpb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Membership.ProtoReflect.Descriptor instead.
func (*Membership) Descriptor() ([]byte, []int) {
	return file_trantorpb_trantorpb_proto_rawDescGZIP(), []int{4}
}

func (x *Membership) GetNodes() map[string]*NodeIdentity {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type NodeIdentity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Addr   string `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"` // Value type is Multiaddr, convert in code directly.
	Key    []byte `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Weight uint64 `protobuf:"varint,4,opt,name=weight,proto3" json:"weight,omitempty"`
}

func (x *NodeIdentity) Reset() {
	*x = NodeIdentity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trantorpb_trantorpb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeIdentity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeIdentity) ProtoMessage() {}

func (x *NodeIdentity) ProtoReflect() protoreflect.Message {
	mi := &file_trantorpb_trantorpb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeIdentity.ProtoReflect.Descriptor instead.
func (*NodeIdentity) Descriptor() ([]byte, []int) {
	return file_trantorpb_trantorpb_proto_rawDescGZIP(), []int{5}
}

func (x *NodeIdentity) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NodeIdentity) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *NodeIdentity) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *NodeIdentity) GetWeight() uint64 {
	if x != nil {
		return x.Weight
	}
	return 0
}

type ClientProgress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Progress map[string]*DeliveredTXs `protobuf:"bytes,1,rep,name=progress,proto3" json:"progress,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ClientProgress) Reset() {
	*x = ClientProgress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trantorpb_trantorpb_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientProgress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientProgress) ProtoMessage() {}

func (x *ClientProgress) ProtoReflect() protoreflect.Message {
	mi := &file_trantorpb_trantorpb_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientProgress.ProtoReflect.Descriptor instead.
func (*ClientProgress) Descriptor() ([]byte, []int) {
	return file_trantorpb_trantorpb_proto_rawDescGZIP(), []int{6}
}

func (x *ClientProgress) GetProgress() map[string]*DeliveredTXs {
	if x != nil {
		return x.Progress
	}
	return nil
}

type DeliveredTXs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LowWm     uint64   `protobuf:"varint,1,opt,name=low_wm,json=lowWm,proto3" json:"low_wm,omitempty"`
	Delivered []uint64 `protobuf:"varint,2,rep,packed,name=delivered,proto3" json:"delivered,omitempty"`
}

func (x *DeliveredTXs) Reset() {
	*x = DeliveredTXs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trantorpb_trantorpb_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliveredTXs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveredTXs) ProtoMessage() {}

func (x *DeliveredTXs) ProtoReflect() protoreflect.Message {
	mi := &file_trantorpb_trantorpb_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveredTXs.ProtoReflect.Descriptor instead.
func (*DeliveredTXs) Descriptor() ([]byte, []int) {
	return file_trantorpb_trantorpb_proto_rawDescGZIP(), []int{7}
}

func (x *DeliveredTXs) GetLowWm() uint64 {
	if x != nil {
		return x.LowWm
	}
	return 0
}

func (x *DeliveredTXs) GetDelivered() []uint64 {
	if x != nil {
		return x.Delivered
	}
	return nil
}

var File_trantorpb_trantorpb_proto protoreflect.FileDescriptor

var file_trantorpb_trantorpb_proto_rawDesc = []byte{
	0x0a, 0x19, 0x74, 0x72, 0x61, 0x6e, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x2f, 0x74, 0x72, 0x61, 0x6e,
	0x74, 0x6f, 0x72, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x74, 0x72, 0x61,
	0x6e, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x1a, 0x1c, 0x6d, 0x69, 0x72, 0x2f, 0x63, 0x6f, 0x64, 0x65,
	0x67, 0x65, 0x6e, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe9, 0x01, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3e, 0x82, 0xa6, 0x1d, 0x3a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x63, 0x6f, 0x69, 0x6e,
	0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x6d, 0x69, 0x72, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x74, 0x72, 0x61, 0x6e, 0x74, 0x6f, 0x72, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x4f, 0x0a, 0x05, 0x74, 0x78, 0x5f, 0x6e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x42, 0x3a, 0x82, 0xa6, 0x1d, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x66, 0x69, 0x6c, 0x65, 0x63, 0x6f, 0x69, 0x6e, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2f, 0x6d, 0x69, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x74, 0x6f,
	0x72, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x54, 0x78, 0x4e, 0x6f, 0x52, 0x04, 0x74, 0x78,
	0x4e, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x3a, 0x04, 0x80, 0xa6, 0x1d, 0x01,
	0x22, 0x65, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f,
	0x74, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x70, 0x70, 0x44, 0x61, 0x74, 0x61, 0x12, 0x33, 0x0a, 0x0a,
	0x65, 0x70, 0x6f, 0x63, 0x68, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x2e, 0x45, 0x70, 0x6f,
	0x63, 0x68, 0x44, 0x61, 0x74, 0x61, 0x52, 0x09, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x44, 0x61, 0x74,
	0x61, 0x3a, 0x04, 0x80, 0xa6, 0x1d, 0x01, 0x22, 0xfd, 0x01, 0x0a, 0x09, 0x45, 0x70, 0x6f, 0x63,
	0x68, 0x44, 0x61, 0x74, 0x61, 0x12, 0x39, 0x0a, 0x0c, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x2e, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x0b, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x42, 0x0a, 0x0f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x72, 0x61, 0x6e,
	0x74, 0x6f, 0x72, 0x70, 0x62, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x52, 0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x6c, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x46, 0x0a, 0x13, 0x70, 0x72, 0x65,
	0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x74, 0x6f, 0x72,
	0x70, 0x62, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x52, 0x12, 0x70,
	0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69,
	0x70, 0x3a, 0x04, 0x80, 0xa6, 0x1d, 0x01, 0x22, 0x96, 0x02, 0x0a, 0x0b, 0x45, 0x70, 0x6f, 0x63,
	0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x58, 0x0a, 0x08, 0x65, 0x70, 0x6f, 0x63, 0x68,
	0x5f, 0x6e, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x3d, 0x82, 0xa6, 0x1d, 0x39, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x63, 0x6f,
	0x69, 0x6e, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x6d, 0x69, 0x72, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x74, 0x6f, 0x72, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x4e, 0x72, 0x52, 0x07, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x4e,
	0x72, 0x12, 0x56, 0x0a, 0x08, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x73, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x42, 0x3b, 0x82, 0xa6, 0x1d, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x63, 0x6f, 0x69, 0x6e, 0x2d, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2f, 0x6d, 0x69, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x72, 0x61,
	0x6e, 0x74, 0x6f, 0x72, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x71, 0x4e, 0x72,
	0x52, 0x07, 0x66, 0x69, 0x72, 0x73, 0x74, 0x53, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74,
	0x68, 0x12, 0x37, 0x0a, 0x0b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x74, 0x6f, 0x72,
	0x70, 0x62, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x52, 0x0b, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x3a, 0x04, 0x98, 0xa6, 0x1d, 0x01,
	0x22, 0xd3, 0x01, 0x0a, 0x0a, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x12,
	0x6c, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x74, 0x72, 0x61, 0x6e, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x42, 0x34, 0xaa, 0xa6, 0x1d, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x66, 0x69, 0x6c, 0x65, 0x63, 0x6f, 0x69, 0x6e, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2f, 0x6d, 0x69, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x4e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x1a, 0x51, 0x0a,
	0x0a, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74,
	0x72, 0x61, 0x6e, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x3a, 0x04, 0x80, 0xa6, 0x1d, 0x01, 0x22, 0xda, 0x01, 0x0a, 0x0c, 0x4e, 0x6f, 0x64, 0x65, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x44, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x34, 0x82, 0xa6, 0x1d, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x63, 0x6f, 0x69, 0x6e, 0x2d, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2f, 0x6d, 0x69, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64,
	0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x58, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x04, 0x42, 0x40, 0x82, 0xa6, 0x1d, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x63, 0x6f, 0x69, 0x6e, 0x2d, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2f, 0x6d, 0x69, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x72, 0x61,
	0x6e, 0x74, 0x6f, 0x72, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x57,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x04, 0x80,
	0xa6, 0x1d, 0x01, 0x22, 0xf2, 0x01, 0x0a, 0x0e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72,
	0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x83, 0x01, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x74, 0x72, 0x61, 0x6e,
	0x74, 0x6f, 0x72, 0x70, 0x62, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x42, 0x3e, 0xaa, 0xa6, 0x1d, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x63, 0x6f, 0x69, 0x6e, 0x2d, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2f, 0x6d, 0x69, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x72, 0x61, 0x6e,
	0x74, 0x6f, 0x72, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x44, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x1a, 0x54, 0x0a, 0x0d,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x2d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x74, 0x72, 0x61, 0x6e, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x65, 0x64, 0x54, 0x58, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x3a, 0x04, 0x98, 0xa6, 0x1d, 0x01, 0x22, 0x49, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x65, 0x64, 0x54, 0x58, 0x73, 0x12, 0x15, 0x0a, 0x06, 0x6c, 0x6f, 0x77, 0x5f,
	0x77, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x6f, 0x77, 0x57, 0x6d, 0x12,
	0x1c, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x64, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x04, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x64, 0x3a, 0x04, 0x80,
	0xa6, 0x1d, 0x01, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x63, 0x6f, 0x69, 0x6e, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2f, 0x6d, 0x69, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_trantorpb_trantorpb_proto_rawDescOnce sync.Once
	file_trantorpb_trantorpb_proto_rawDescData = file_trantorpb_trantorpb_proto_rawDesc
)

func file_trantorpb_trantorpb_proto_rawDescGZIP() []byte {
	file_trantorpb_trantorpb_proto_rawDescOnce.Do(func() {
		file_trantorpb_trantorpb_proto_rawDescData = protoimpl.X.CompressGZIP(file_trantorpb_trantorpb_proto_rawDescData)
	})
	return file_trantorpb_trantorpb_proto_rawDescData
}

var file_trantorpb_trantorpb_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_trantorpb_trantorpb_proto_goTypes = []interface{}{
	(*Transaction)(nil),    // 0: trantorpb.Transaction
	(*StateSnapshot)(nil),  // 1: trantorpb.StateSnapshot
	(*EpochData)(nil),      // 2: trantorpb.EpochData
	(*EpochConfig)(nil),    // 3: trantorpb.EpochConfig
	(*Membership)(nil),     // 4: trantorpb.Membership
	(*NodeIdentity)(nil),   // 5: trantorpb.NodeIdentity
	(*ClientProgress)(nil), // 6: trantorpb.ClientProgress
	(*DeliveredTXs)(nil),   // 7: trantorpb.DeliveredTXs
	nil,                    // 8: trantorpb.Membership.NodesEntry
	nil,                    // 9: trantorpb.ClientProgress.ProgressEntry
}
var file_trantorpb_trantorpb_proto_depIdxs = []int32{
	2, // 0: trantorpb.StateSnapshot.epoch_data:type_name -> trantorpb.EpochData
	3, // 1: trantorpb.EpochData.epoch_config:type_name -> trantorpb.EpochConfig
	6, // 2: trantorpb.EpochData.client_progress:type_name -> trantorpb.ClientProgress
	4, // 3: trantorpb.EpochData.previous_membership:type_name -> trantorpb.Membership
	4, // 4: trantorpb.EpochConfig.memberships:type_name -> trantorpb.Membership
	8, // 5: trantorpb.Membership.nodes:type_name -> trantorpb.Membership.NodesEntry
	9, // 6: trantorpb.ClientProgress.progress:type_name -> trantorpb.ClientProgress.ProgressEntry
	5, // 7: trantorpb.Membership.NodesEntry.value:type_name -> trantorpb.NodeIdentity
	7, // 8: trantorpb.ClientProgress.ProgressEntry.value:type_name -> trantorpb.DeliveredTXs
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_trantorpb_trantorpb_proto_init() }
func file_trantorpb_trantorpb_proto_init() {
	if File_trantorpb_trantorpb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_trantorpb_trantorpb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_trantorpb_trantorpb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateSnapshot); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_trantorpb_trantorpb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EpochData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_trantorpb_trantorpb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EpochConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_trantorpb_trantorpb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Membership); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_trantorpb_trantorpb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeIdentity); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_trantorpb_trantorpb_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientProgress); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_trantorpb_trantorpb_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliveredTXs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_trantorpb_trantorpb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_trantorpb_trantorpb_proto_goTypes,
		DependencyIndexes: file_trantorpb_trantorpb_proto_depIdxs,
		MessageInfos:      file_trantorpb_trantorpb_proto_msgTypes,
	}.Build()
	File_trantorpb_trantorpb_proto = out.File
	file_trantorpb_trantorpb_proto_rawDesc = nil
	file_trantorpb_trantorpb_proto_goTypes = nil
	file_trantorpb_trantorpb_proto_depIdxs = nil
}
