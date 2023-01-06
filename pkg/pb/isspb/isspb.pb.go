//
//Copyright IBM Corp. All Rights Reserved.
//
//SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: isspb/isspb.proto

package isspb

import (
	checkpointpb "github.com/filecoin-project/mir/pkg/pb/checkpointpb"
	commonpb "github.com/filecoin-project/mir/pkg/pb/commonpb"
	requestpb "github.com/filecoin-project/mir/pkg/pb/requestpb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ISSMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//	*ISSMessage_StableCheckpoint
	//	*ISSMessage_RetransmitRequests
	Type isISSMessage_Type `protobuf_oneof:"type"`
}

func (x *ISSMessage) Reset() {
	*x = ISSMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isspb_isspb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ISSMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ISSMessage) ProtoMessage() {}

func (x *ISSMessage) ProtoReflect() protoreflect.Message {
	mi := &file_isspb_isspb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ISSMessage.ProtoReflect.Descriptor instead.
func (*ISSMessage) Descriptor() ([]byte, []int) {
	return file_isspb_isspb_proto_rawDescGZIP(), []int{0}
}

func (m *ISSMessage) GetType() isISSMessage_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *ISSMessage) GetStableCheckpoint() *checkpointpb.StableCheckpoint {
	if x, ok := x.GetType().(*ISSMessage_StableCheckpoint); ok {
		return x.StableCheckpoint
	}
	return nil
}

func (x *ISSMessage) GetRetransmitRequests() *RetransmitRequests {
	if x, ok := x.GetType().(*ISSMessage_RetransmitRequests); ok {
		return x.RetransmitRequests
	}
	return nil
}

type isISSMessage_Type interface {
	isISSMessage_Type()
}

type ISSMessage_StableCheckpoint struct {
	// SBMessage                       sb                  = 1;
	StableCheckpoint *checkpointpb.StableCheckpoint `protobuf:"bytes,3,opt,name=stable_checkpoint,json=stableCheckpoint,proto3,oneof"`
}

type ISSMessage_RetransmitRequests struct {
	RetransmitRequests *RetransmitRequests `protobuf:"bytes,4,opt,name=retransmit_requests,json=retransmitRequests,proto3,oneof"`
}

func (*ISSMessage_StableCheckpoint) isISSMessage_Type() {}

func (*ISSMessage_RetransmitRequests) isISSMessage_Type() {}

type RetransmitRequests struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Requests []*requestpb.Request `protobuf:"bytes,1,rep,name=requests,proto3" json:"requests,omitempty"`
}

func (x *RetransmitRequests) Reset() {
	*x = RetransmitRequests{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isspb_isspb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetransmitRequests) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetransmitRequests) ProtoMessage() {}

func (x *RetransmitRequests) ProtoReflect() protoreflect.Message {
	mi := &file_isspb_isspb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetransmitRequests.ProtoReflect.Descriptor instead.
func (*RetransmitRequests) Descriptor() ([]byte, []int) {
	return file_isspb_isspb_proto_rawDescGZIP(), []int{1}
}

func (x *RetransmitRequests) GetRequests() []*requestpb.Request {
	if x != nil {
		return x.Requests
	}
	return nil
}

type ISSEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//	*ISSEvent_PersistCheckpoint
	//	*ISSEvent_PersistStableCheckpoint
	//	*ISSEvent_PushCheckpoint
	//	*ISSEvent_SbDeliver
	Type isISSEvent_Type `protobuf_oneof:"type"`
}

func (x *ISSEvent) Reset() {
	*x = ISSEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isspb_isspb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ISSEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ISSEvent) ProtoMessage() {}

func (x *ISSEvent) ProtoReflect() protoreflect.Message {
	mi := &file_isspb_isspb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ISSEvent.ProtoReflect.Descriptor instead.
func (*ISSEvent) Descriptor() ([]byte, []int) {
	return file_isspb_isspb_proto_rawDescGZIP(), []int{2}
}

func (m *ISSEvent) GetType() isISSEvent_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *ISSEvent) GetPersistCheckpoint() *PersistCheckpoint {
	if x, ok := x.GetType().(*ISSEvent_PersistCheckpoint); ok {
		return x.PersistCheckpoint
	}
	return nil
}

func (x *ISSEvent) GetPersistStableCheckpoint() *PersistStableCheckpoint {
	if x, ok := x.GetType().(*ISSEvent_PersistStableCheckpoint); ok {
		return x.PersistStableCheckpoint
	}
	return nil
}

func (x *ISSEvent) GetPushCheckpoint() *PushCheckpoint {
	if x, ok := x.GetType().(*ISSEvent_PushCheckpoint); ok {
		return x.PushCheckpoint
	}
	return nil
}

func (x *ISSEvent) GetSbDeliver() *SBDeliver {
	if x, ok := x.GetType().(*ISSEvent_SbDeliver); ok {
		return x.SbDeliver
	}
	return nil
}

type isISSEvent_Type interface {
	isISSEvent_Type()
}

type ISSEvent_PersistCheckpoint struct {
	PersistCheckpoint *PersistCheckpoint `protobuf:"bytes,1,opt,name=persist_checkpoint,json=persistCheckpoint,proto3,oneof"`
}

type ISSEvent_PersistStableCheckpoint struct {
	PersistStableCheckpoint *PersistStableCheckpoint `protobuf:"bytes,3,opt,name=persist_stable_checkpoint,json=persistStableCheckpoint,proto3,oneof"`
}

type ISSEvent_PushCheckpoint struct {
	PushCheckpoint *PushCheckpoint `protobuf:"bytes,5,opt,name=push_checkpoint,json=pushCheckpoint,proto3,oneof"`
}

type ISSEvent_SbDeliver struct {
	SbDeliver *SBDeliver `protobuf:"bytes,6,opt,name=sb_deliver,json=sbDeliver,proto3,oneof"`
}

func (*ISSEvent_PersistCheckpoint) isISSEvent_Type() {}

func (*ISSEvent_PersistStableCheckpoint) isISSEvent_Type() {}

func (*ISSEvent_PushCheckpoint) isISSEvent_Type() {}

func (*ISSEvent_SbDeliver) isISSEvent_Type() {}

type ISSHashOrigin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//	*ISSHashOrigin_LogEntrySn
	//	*ISSHashOrigin_StateSnapshotEpoch
	//	*ISSHashOrigin_Requests
	Type isISSHashOrigin_Type `protobuf_oneof:"type"`
}

func (x *ISSHashOrigin) Reset() {
	*x = ISSHashOrigin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isspb_isspb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ISSHashOrigin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ISSHashOrigin) ProtoMessage() {}

func (x *ISSHashOrigin) ProtoReflect() protoreflect.Message {
	mi := &file_isspb_isspb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ISSHashOrigin.ProtoReflect.Descriptor instead.
func (*ISSHashOrigin) Descriptor() ([]byte, []int) {
	return file_isspb_isspb_proto_rawDescGZIP(), []int{3}
}

func (m *ISSHashOrigin) GetType() isISSHashOrigin_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *ISSHashOrigin) GetLogEntrySn() uint64 {
	if x, ok := x.GetType().(*ISSHashOrigin_LogEntrySn); ok {
		return x.LogEntrySn
	}
	return 0
}

func (x *ISSHashOrigin) GetStateSnapshotEpoch() uint64 {
	if x, ok := x.GetType().(*ISSHashOrigin_StateSnapshotEpoch); ok {
		return x.StateSnapshotEpoch
	}
	return 0
}

func (x *ISSHashOrigin) GetRequests() *RequestHashOrigin {
	if x, ok := x.GetType().(*ISSHashOrigin_Requests); ok {
		return x.Requests
	}
	return nil
}

type isISSHashOrigin_Type interface {
	isISSHashOrigin_Type()
}

type ISSHashOrigin_LogEntrySn struct {
	LogEntrySn uint64 `protobuf:"varint,2,opt,name=log_entry_sn,json=logEntrySn,proto3,oneof"`
}

type ISSHashOrigin_StateSnapshotEpoch struct {
	StateSnapshotEpoch uint64 `protobuf:"varint,3,opt,name=state_snapshot_epoch,json=stateSnapshotEpoch,proto3,oneof"`
}

type ISSHashOrigin_Requests struct {
	Requests *RequestHashOrigin `protobuf:"bytes,4,opt,name=requests,proto3,oneof"`
}

func (*ISSHashOrigin_LogEntrySn) isISSHashOrigin_Type() {}

func (*ISSHashOrigin_StateSnapshotEpoch) isISSHashOrigin_Type() {}

func (*ISSHashOrigin_Requests) isISSHashOrigin_Type() {}

type RequestHashOrigin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Requests []*requestpb.Request `protobuf:"bytes,1,rep,name=requests,proto3" json:"requests,omitempty"`
}

func (x *RequestHashOrigin) Reset() {
	*x = RequestHashOrigin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isspb_isspb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestHashOrigin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestHashOrigin) ProtoMessage() {}

func (x *RequestHashOrigin) ProtoReflect() protoreflect.Message {
	mi := &file_isspb_isspb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestHashOrigin.ProtoReflect.Descriptor instead.
func (*RequestHashOrigin) Descriptor() ([]byte, []int) {
	return file_isspb_isspb_proto_rawDescGZIP(), []int{4}
}

func (x *RequestHashOrigin) GetRequests() []*requestpb.Request {
	if x != nil {
		return x.Requests
	}
	return nil
}

type ISSSigVerOrigin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//	*ISSSigVerOrigin_CheckpointEpoch
	//	*ISSSigVerOrigin_StableCheckpoint
	Type isISSSigVerOrigin_Type `protobuf_oneof:"type"`
}

func (x *ISSSigVerOrigin) Reset() {
	*x = ISSSigVerOrigin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isspb_isspb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ISSSigVerOrigin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ISSSigVerOrigin) ProtoMessage() {}

func (x *ISSSigVerOrigin) ProtoReflect() protoreflect.Message {
	mi := &file_isspb_isspb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ISSSigVerOrigin.ProtoReflect.Descriptor instead.
func (*ISSSigVerOrigin) Descriptor() ([]byte, []int) {
	return file_isspb_isspb_proto_rawDescGZIP(), []int{5}
}

func (m *ISSSigVerOrigin) GetType() isISSSigVerOrigin_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *ISSSigVerOrigin) GetCheckpointEpoch() uint64 {
	if x, ok := x.GetType().(*ISSSigVerOrigin_CheckpointEpoch); ok {
		return x.CheckpointEpoch
	}
	return 0
}

func (x *ISSSigVerOrigin) GetStableCheckpoint() *checkpointpb.StableCheckpoint {
	if x, ok := x.GetType().(*ISSSigVerOrigin_StableCheckpoint); ok {
		return x.StableCheckpoint
	}
	return nil
}

type isISSSigVerOrigin_Type interface {
	isISSSigVerOrigin_Type()
}

type ISSSigVerOrigin_CheckpointEpoch struct {
	CheckpointEpoch uint64 `protobuf:"varint,2,opt,name=checkpoint_epoch,json=checkpointEpoch,proto3,oneof"`
}

type ISSSigVerOrigin_StableCheckpoint struct {
	StableCheckpoint *checkpointpb.StableCheckpoint `protobuf:"bytes,3,opt,name=stable_checkpoint,json=stableCheckpoint,proto3,oneof"`
}

func (*ISSSigVerOrigin_CheckpointEpoch) isISSSigVerOrigin_Type() {}

func (*ISSSigVerOrigin_StableCheckpoint) isISSSigVerOrigin_Type() {}

type PersistCheckpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sn                uint64                  `protobuf:"varint,1,opt,name=sn,proto3" json:"sn,omitempty"`
	StateSnapshot     *commonpb.StateSnapshot `protobuf:"bytes,2,opt,name=state_snapshot,json=stateSnapshot,proto3" json:"state_snapshot,omitempty"`
	StateSnapshotHash []byte                  `protobuf:"bytes,3,opt,name=state_snapshot_hash,json=stateSnapshotHash,proto3" json:"state_snapshot_hash,omitempty"`
	Signature         []byte                  `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *PersistCheckpoint) Reset() {
	*x = PersistCheckpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isspb_isspb_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersistCheckpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersistCheckpoint) ProtoMessage() {}

func (x *PersistCheckpoint) ProtoReflect() protoreflect.Message {
	mi := &file_isspb_isspb_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersistCheckpoint.ProtoReflect.Descriptor instead.
func (*PersistCheckpoint) Descriptor() ([]byte, []int) {
	return file_isspb_isspb_proto_rawDescGZIP(), []int{6}
}

func (x *PersistCheckpoint) GetSn() uint64 {
	if x != nil {
		return x.Sn
	}
	return 0
}

func (x *PersistCheckpoint) GetStateSnapshot() *commonpb.StateSnapshot {
	if x != nil {
		return x.StateSnapshot
	}
	return nil
}

func (x *PersistCheckpoint) GetStateSnapshotHash() []byte {
	if x != nil {
		return x.StateSnapshotHash
	}
	return nil
}

func (x *PersistCheckpoint) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

// PersistStableCheckpoint needs to be a separate Event from StableCheckpoint, since both are ISSEvents,
// but, the protocol must differentiate between them. While the former will be applied on recovery from the WAL,
// the latter serves as a notification to the ISS protocol when a stable checkpoint has been persisted.
type PersistStableCheckpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StableCheckpoint *checkpointpb.StableCheckpoint `protobuf:"bytes,1,opt,name=stable_checkpoint,json=stableCheckpoint,proto3" json:"stable_checkpoint,omitempty"`
}

func (x *PersistStableCheckpoint) Reset() {
	*x = PersistStableCheckpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isspb_isspb_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersistStableCheckpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersistStableCheckpoint) ProtoMessage() {}

func (x *PersistStableCheckpoint) ProtoReflect() protoreflect.Message {
	mi := &file_isspb_isspb_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersistStableCheckpoint.ProtoReflect.Descriptor instead.
func (*PersistStableCheckpoint) Descriptor() ([]byte, []int) {
	return file_isspb_isspb_proto_rawDescGZIP(), []int{7}
}

func (x *PersistStableCheckpoint) GetStableCheckpoint() *checkpointpb.StableCheckpoint {
	if x != nil {
		return x.StableCheckpoint
	}
	return nil
}

type PushCheckpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PushCheckpoint) Reset() {
	*x = PushCheckpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isspb_isspb_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushCheckpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushCheckpoint) ProtoMessage() {}

func (x *PushCheckpoint) ProtoReflect() protoreflect.Message {
	mi := &file_isspb_isspb_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushCheckpoint.ProtoReflect.Descriptor instead.
func (*PushCheckpoint) Descriptor() ([]byte, []int) {
	return file_isspb_isspb_proto_rawDescGZIP(), []int{8}
}

type SBDeliver struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sn       uint64 `protobuf:"varint,1,opt,name=sn,proto3" json:"sn,omitempty"`
	CertData []byte `protobuf:"bytes,2,opt,name=cert_data,json=certData,proto3" json:"cert_data,omitempty"`
	Aborted  bool   `protobuf:"varint,3,opt,name=aborted,proto3" json:"aborted,omitempty"`
	Leader   string `protobuf:"bytes,4,opt,name=leader,proto3" json:"leader,omitempty"`
}

func (x *SBDeliver) Reset() {
	*x = SBDeliver{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isspb_isspb_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SBDeliver) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SBDeliver) ProtoMessage() {}

func (x *SBDeliver) ProtoReflect() protoreflect.Message {
	mi := &file_isspb_isspb_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SBDeliver.ProtoReflect.Descriptor instead.
func (*SBDeliver) Descriptor() ([]byte, []int) {
	return file_isspb_isspb_proto_rawDescGZIP(), []int{9}
}

func (x *SBDeliver) GetSn() uint64 {
	if x != nil {
		return x.Sn
	}
	return 0
}

func (x *SBDeliver) GetCertData() []byte {
	if x != nil {
		return x.CertData
	}
	return nil
}

func (x *SBDeliver) GetAborted() bool {
	if x != nil {
		return x.Aborted
	}
	return false
}

func (x *SBDeliver) GetLeader() string {
	if x != nil {
		return x.Leader
	}
	return ""
}

var File_isspb_isspb_proto protoreflect.FileDescriptor

var file_isspb_isspb_proto_rawDesc = []byte{
	0x0a, 0x11, 0x69, 0x73, 0x73, 0x70, 0x62, 0x2f, 0x69, 0x73, 0x73, 0x70, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x69, 0x73, 0x73, 0x70, 0x62, 0x1a, 0x17, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x70, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x70, 0x62, 0x2f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xb1, 0x01, 0x0a, 0x0a, 0x49, 0x53, 0x53, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x4d,
	0x0a, 0x11, 0x73, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x10, 0x73, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x4c, 0x0a,
	0x13, 0x72, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x69, 0x73, 0x73,
	0x70, 0x62, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x73, 0x48, 0x00, 0x52, 0x12, 0x72, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x42, 0x06, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x22, 0x44, 0x0a, 0x12, 0x52, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x2e, 0x0a, 0x08, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x22, 0xb0, 0x02, 0x0a, 0x08, 0x49, 0x53,
	0x53, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x49, 0x0a, 0x12, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73,
	0x74, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x69, 0x73, 0x73, 0x70, 0x62, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x69,
	0x73, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x11,
	0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x12, 0x5c, 0x0a, 0x19, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x5f, 0x73, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x69, 0x73, 0x73, 0x70, 0x62, 0x2e, 0x50, 0x65, 0x72,
	0x73, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x17, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x53,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12,
	0x40, 0x0a, 0x0f, 0x70, 0x75, 0x73, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x69, 0x73, 0x73, 0x70, 0x62,
	0x2e, 0x50, 0x75, 0x73, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x48,
	0x00, 0x52, 0x0e, 0x70, 0x75, 0x73, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x12, 0x31, 0x0a, 0x0a, 0x73, 0x62, 0x5f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x69, 0x73, 0x73, 0x70, 0x62, 0x2e, 0x53, 0x42,
	0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x48, 0x00, 0x52, 0x09, 0x73, 0x62, 0x44, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0xa7, 0x01, 0x0a,
	0x0d, 0x49, 0x53, 0x53, 0x48, 0x61, 0x73, 0x68, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x22,
	0x0a, 0x0c, 0x6c, 0x6f, 0x67, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x73, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x0a, 0x6c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x53, 0x6e, 0x12, 0x32, 0x0a, 0x14, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x48, 0x00, 0x52, 0x12, 0x73, 0x74, 0x61, 0x74, 0x65, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f,
	0x74, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x12, 0x36, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x69, 0x73, 0x73, 0x70, 0x62,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x61, 0x73, 0x68, 0x4f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x42, 0x06,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x43, 0x0a, 0x11, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x48, 0x61, 0x73, 0x68, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x2e, 0x0a, 0x08, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x22, 0x95, 0x01, 0x0a, 0x0f,
	0x49, 0x53, 0x53, 0x53, 0x69, 0x67, 0x56, 0x65, 0x72, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12,
	0x2b, 0x0a, 0x10, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x65, 0x70,
	0x6f, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x0f, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x12, 0x4d, 0x0a, 0x11,
	0x73, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x10, 0x73, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x42, 0x06, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x22, 0xb1, 0x01, 0x0a, 0x11, 0x50, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x73, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x73, 0x6e, 0x12, 0x3e, 0x0a, 0x0e, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x5f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x5f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x11, 0x73, 0x74, 0x61, 0x74, 0x65, 0x53, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x66, 0x0a, 0x17, 0x50, 0x65, 0x72, 0x73, 0x69,
	0x73, 0x74, 0x53, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x12, 0x4b, 0x0a, 0x11, 0x73, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x10, 0x73,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22,
	0x10, 0x0a, 0x0e, 0x50, 0x75, 0x73, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x22, 0x6a, 0x0a, 0x09, 0x53, 0x42, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x73, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x73, 0x6e, 0x12, 0x1b,
	0x0a, 0x09, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x08, 0x63, 0x65, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x62, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x61, 0x62,
	0x6f, 0x72, 0x74, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42, 0x2e, 0x5a,
	0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x6c, 0x65,
	0x63, 0x6f, 0x69, 0x6e, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x6d, 0x69, 0x72,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x69, 0x73, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_isspb_isspb_proto_rawDescOnce sync.Once
	file_isspb_isspb_proto_rawDescData = file_isspb_isspb_proto_rawDesc
)

func file_isspb_isspb_proto_rawDescGZIP() []byte {
	file_isspb_isspb_proto_rawDescOnce.Do(func() {
		file_isspb_isspb_proto_rawDescData = protoimpl.X.CompressGZIP(file_isspb_isspb_proto_rawDescData)
	})
	return file_isspb_isspb_proto_rawDescData
}

var file_isspb_isspb_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_isspb_isspb_proto_goTypes = []interface{}{
	(*ISSMessage)(nil),                    // 0: isspb.ISSMessage
	(*RetransmitRequests)(nil),            // 1: isspb.RetransmitRequests
	(*ISSEvent)(nil),                      // 2: isspb.ISSEvent
	(*ISSHashOrigin)(nil),                 // 3: isspb.ISSHashOrigin
	(*RequestHashOrigin)(nil),             // 4: isspb.RequestHashOrigin
	(*ISSSigVerOrigin)(nil),               // 5: isspb.ISSSigVerOrigin
	(*PersistCheckpoint)(nil),             // 6: isspb.PersistCheckpoint
	(*PersistStableCheckpoint)(nil),       // 7: isspb.PersistStableCheckpoint
	(*PushCheckpoint)(nil),                // 8: isspb.PushCheckpoint
	(*SBDeliver)(nil),                     // 9: isspb.SBDeliver
	(*checkpointpb.StableCheckpoint)(nil), // 10: checkpointpb.StableCheckpoint
	(*requestpb.Request)(nil),             // 11: requestpb.Request
	(*commonpb.StateSnapshot)(nil),        // 12: commonpb.StateSnapshot
}
var file_isspb_isspb_proto_depIdxs = []int32{
	10, // 0: isspb.ISSMessage.stable_checkpoint:type_name -> checkpointpb.StableCheckpoint
	1,  // 1: isspb.ISSMessage.retransmit_requests:type_name -> isspb.RetransmitRequests
	11, // 2: isspb.RetransmitRequests.requests:type_name -> requestpb.Request
	6,  // 3: isspb.ISSEvent.persist_checkpoint:type_name -> isspb.PersistCheckpoint
	7,  // 4: isspb.ISSEvent.persist_stable_checkpoint:type_name -> isspb.PersistStableCheckpoint
	8,  // 5: isspb.ISSEvent.push_checkpoint:type_name -> isspb.PushCheckpoint
	9,  // 6: isspb.ISSEvent.sb_deliver:type_name -> isspb.SBDeliver
	4,  // 7: isspb.ISSHashOrigin.requests:type_name -> isspb.RequestHashOrigin
	11, // 8: isspb.RequestHashOrigin.requests:type_name -> requestpb.Request
	10, // 9: isspb.ISSSigVerOrigin.stable_checkpoint:type_name -> checkpointpb.StableCheckpoint
	12, // 10: isspb.PersistCheckpoint.state_snapshot:type_name -> commonpb.StateSnapshot
	10, // 11: isspb.PersistStableCheckpoint.stable_checkpoint:type_name -> checkpointpb.StableCheckpoint
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_isspb_isspb_proto_init() }
func file_isspb_isspb_proto_init() {
	if File_isspb_isspb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_isspb_isspb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ISSMessage); i {
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
		file_isspb_isspb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetransmitRequests); i {
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
		file_isspb_isspb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ISSEvent); i {
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
		file_isspb_isspb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ISSHashOrigin); i {
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
		file_isspb_isspb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestHashOrigin); i {
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
		file_isspb_isspb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ISSSigVerOrigin); i {
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
		file_isspb_isspb_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersistCheckpoint); i {
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
		file_isspb_isspb_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersistStableCheckpoint); i {
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
		file_isspb_isspb_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushCheckpoint); i {
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
		file_isspb_isspb_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SBDeliver); i {
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
	file_isspb_isspb_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ISSMessage_StableCheckpoint)(nil),
		(*ISSMessage_RetransmitRequests)(nil),
	}
	file_isspb_isspb_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*ISSEvent_PersistCheckpoint)(nil),
		(*ISSEvent_PersistStableCheckpoint)(nil),
		(*ISSEvent_PushCheckpoint)(nil),
		(*ISSEvent_SbDeliver)(nil),
	}
	file_isspb_isspb_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*ISSHashOrigin_LogEntrySn)(nil),
		(*ISSHashOrigin_StateSnapshotEpoch)(nil),
		(*ISSHashOrigin_Requests)(nil),
	}
	file_isspb_isspb_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*ISSSigVerOrigin_CheckpointEpoch)(nil),
		(*ISSSigVerOrigin_StableCheckpoint)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_isspb_isspb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_isspb_isspb_proto_goTypes,
		DependencyIndexes: file_isspb_isspb_proto_depIdxs,
		MessageInfos:      file_isspb_isspb_proto_msgTypes,
	}.Build()
	File_isspb_isspb_proto = out.File
	file_isspb_isspb_proto_rawDesc = nil
	file_isspb_isspb_proto_goTypes = nil
	file_isspb_isspb_proto_depIdxs = nil
}
