// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: ordererpb/ordererpb.proto

package ordererpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	_ "github.com/filecoin-project/mir/pkg/pb/mir"
	_ "github.com/filecoin-project/mir/pkg/pb/net"
	pbftpb "github.com/filecoin-project/mir/pkg/pb/pbftpb"
	trantorpb "github.com/filecoin-project/mir/pkg/pb/trantorpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//	*Event_Pbft
	Type isEvent_Type `protobuf_oneof:"type"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ordererpb_ordererpb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_ordererpb_ordererpb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_ordererpb_ordererpb_proto_rawDescGZIP(), []int{0}
}

func (m *Event) GetType() isEvent_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *Event) GetPbft() *pbftpb.Event {
	if x, ok := x.GetType().(*Event_Pbft); ok {
		return x.Pbft
	}
	return nil
}

type isEvent_Type interface {
	isEvent_Type()
}

type Event_Pbft struct {
	Pbft *pbftpb.Event `protobuf:"bytes,1,opt,name=pbft,proto3,oneof"`
}

func (*Event_Pbft) isEvent_Type() {}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//	*Message_Pbft
	Type isMessage_Type `protobuf_oneof:"type"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ordererpb_ordererpb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_ordererpb_ordererpb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_ordererpb_ordererpb_proto_rawDescGZIP(), []int{1}
}

func (m *Message) GetType() isMessage_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *Message) GetPbft() *pbftpb.Message {
	if x, ok := x.GetType().(*Message_Pbft); ok {
		return x.Pbft
	}
	return nil
}

type isMessage_Type interface {
	isMessage_Type()
}

type Message_Pbft struct {
	Pbft *pbftpb.Message `protobuf:"bytes,1,opt,name=pbft,proto3,oneof"`
}

func (*Message_Pbft) isMessage_Type() {}

type PBFTSegment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The leader node of the orderer.
	Leader string `protobuf:"bytes,1,opt,name=leader,proto3" json:"leader,omitempty"`
	// All nodes executing the orderer implementation.
	Membership *trantorpb.Membership `protobuf:"bytes,2,opt,name=membership,proto3" json:"membership,omitempty"`
	// Sequence numbers for which the orderer is responsible, along with corresponding (optional) pre-defined proposals.
	// The keys of this map are the actual "segment" of the commit log.
	// A nil value means that no proposal is specified (and the protocol implementation will decide what to propose).
	// A non-nil value will be proposed (by this node) for that sequence number whenever possible.
	// Currently, such a "free" proposal is a new availability certificate in view 0,
	// and a special empty one in other views.
	Proposals map[uint64][]byte `protobuf:"bytes,3,rep,name=proposals,proto3" json:"proposals,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PBFTSegment) Reset() {
	*x = PBFTSegment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ordererpb_ordererpb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PBFTSegment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PBFTSegment) ProtoMessage() {}

func (x *PBFTSegment) ProtoReflect() protoreflect.Message {
	mi := &file_ordererpb_ordererpb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PBFTSegment.ProtoReflect.Descriptor instead.
func (*PBFTSegment) Descriptor() ([]byte, []int) {
	return file_ordererpb_ordererpb_proto_rawDescGZIP(), []int{2}
}

func (x *PBFTSegment) GetLeader() string {
	if x != nil {
		return x.Leader
	}
	return ""
}

func (x *PBFTSegment) GetMembership() *trantorpb.Membership {
	if x != nil {
		return x.Membership
	}
	return nil
}

func (x *PBFTSegment) GetProposals() map[uint64][]byte {
	if x != nil {
		return x.Proposals
	}
	return nil
}

type PBFTModule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Segment         *PBFTSegment `protobuf:"bytes,1,opt,name=segment,proto3" json:"segment,omitempty"`
	AvailabilityId  string       `protobuf:"bytes,2,opt,name=availability_id,json=availabilityId,proto3" json:"availability_id,omitempty"`
	Epoch           uint64       `protobuf:"varint,3,opt,name=epoch,proto3" json:"epoch,omitempty"`
	ValidityChecker uint64       `protobuf:"varint,4,opt,name=validity_checker,json=validityChecker,proto3" json:"validity_checker,omitempty"`
}

func (x *PBFTModule) Reset() {
	*x = PBFTModule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ordererpb_ordererpb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PBFTModule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PBFTModule) ProtoMessage() {}

func (x *PBFTModule) ProtoReflect() protoreflect.Message {
	mi := &file_ordererpb_ordererpb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PBFTModule.ProtoReflect.Descriptor instead.
func (*PBFTModule) Descriptor() ([]byte, []int) {
	return file_ordererpb_ordererpb_proto_rawDescGZIP(), []int{3}
}

func (x *PBFTModule) GetSegment() *PBFTSegment {
	if x != nil {
		return x.Segment
	}
	return nil
}

func (x *PBFTModule) GetAvailabilityId() string {
	if x != nil {
		return x.AvailabilityId
	}
	return ""
}

func (x *PBFTModule) GetEpoch() uint64 {
	if x != nil {
		return x.Epoch
	}
	return 0
}

func (x *PBFTModule) GetValidityChecker() uint64 {
	if x != nil {
		return x.ValidityChecker
	}
	return 0
}

var File_ordererpb_ordererpb_proto protoreflect.FileDescriptor

var file_ordererpb_ordererpb_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x70, 0x62, 0x2f, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x65, 0x72, 0x70, 0x62, 0x1a, 0x19, 0x74, 0x72, 0x61, 0x6e, 0x74, 0x6f, 0x72, 0x70,
	0x62, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x13, 0x70, 0x62, 0x66, 0x74, 0x70, 0x62, 0x2f, 0x70, 0x62, 0x66, 0x74, 0x70, 0x62,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x6d, 0x69, 0x72, 0x2f, 0x63, 0x6f, 0x64, 0x65,
	0x67, 0x65, 0x6e, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x6e, 0x65, 0x74, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x67, 0x65,
	0x6e, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x40, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x04, 0x70,
	0x62, 0x66, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x66, 0x74,
	0x70, 0x62, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x04, 0x70, 0x62, 0x66, 0x74,
	0x3a, 0x04, 0x90, 0xa6, 0x1d, 0x01, 0x42, 0x0c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x04,
	0x80, 0xa6, 0x1d, 0x01, 0x22, 0x44, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x25, 0x0a, 0x04, 0x70, 0x62, 0x66, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x70, 0x62, 0x66, 0x74, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00,
	0x52, 0x04, 0x70, 0x62, 0x66, 0x74, 0x3a, 0x04, 0xc8, 0xe4, 0x1d, 0x01, 0x42, 0x0c, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x04, 0xc8, 0xe4, 0x1d, 0x01, 0x22, 0xd9, 0x02, 0x0a, 0x0b, 0x50,
	0x42, 0x46, 0x54, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x4c, 0x0a, 0x06, 0x6c, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x34, 0x82, 0xa6, 0x1d, 0x30,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x63,
	0x6f, 0x69, 0x6e, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x6d, 0x69, 0x72, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x44,
	0x52, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x0a, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74,
	0x72, 0x61, 0x6e, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73,
	0x68, 0x69, 0x70, 0x52, 0x0a, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x12,
	0x80, 0x01, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x70, 0x62, 0x2e,
	0x50, 0x42, 0x46, 0x54, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x70,
	0x6f, 0x73, 0x61, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x3b, 0xaa, 0xa6, 0x1d, 0x37,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x63,
	0x6f, 0x69, 0x6e, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x6d, 0x69, 0x72, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x74, 0x6f, 0x72, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x53, 0x65, 0x71, 0x4e, 0x72, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61,
	0x6c, 0x73, 0x1a, 0x3c, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x3a, 0x04, 0x80, 0xa6, 0x1d, 0x01, 0x22, 0xae, 0x01, 0x0a, 0x0a, 0x50, 0x42, 0x46, 0x54, 0x4d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x30, 0x0a, 0x07, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72,
	0x70, 0x62, 0x2e, 0x50, 0x42, 0x46, 0x54, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07,
	0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x12, 0x29, 0x0a, 0x10, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x69,
	0x74, 0x79, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x69, 0x74, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x65,
	0x72, 0x3a, 0x04, 0x80, 0xa6, 0x1d, 0x01, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x63, 0x6f, 0x69, 0x6e, 0x2d, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x6d, 0x69, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70,
	0x62, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_ordererpb_ordererpb_proto_rawDescOnce sync.Once
	file_ordererpb_ordererpb_proto_rawDescData = file_ordererpb_ordererpb_proto_rawDesc
)

func file_ordererpb_ordererpb_proto_rawDescGZIP() []byte {
	file_ordererpb_ordererpb_proto_rawDescOnce.Do(func() {
		file_ordererpb_ordererpb_proto_rawDescData = protoimpl.X.CompressGZIP(file_ordererpb_ordererpb_proto_rawDescData)
	})
	return file_ordererpb_ordererpb_proto_rawDescData
}

var file_ordererpb_ordererpb_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_ordererpb_ordererpb_proto_goTypes = []interface{}{
	(*Event)(nil),                // 0: ordererpb.Event
	(*Message)(nil),              // 1: ordererpb.Message
	(*PBFTSegment)(nil),          // 2: ordererpb.PBFTSegment
	(*PBFTModule)(nil),           // 3: ordererpb.PBFTModule
	nil,                          // 4: ordererpb.PBFTSegment.ProposalsEntry
	(*pbftpb.Event)(nil),         // 5: pbftpb.Event
	(*pbftpb.Message)(nil),       // 6: pbftpb.Message
	(*trantorpb.Membership)(nil), // 7: trantorpb.Membership
}
var file_ordererpb_ordererpb_proto_depIdxs = []int32{
	5, // 0: ordererpb.Event.pbft:type_name -> pbftpb.Event
	6, // 1: ordererpb.Message.pbft:type_name -> pbftpb.Message
	7, // 2: ordererpb.PBFTSegment.membership:type_name -> trantorpb.Membership
	4, // 3: ordererpb.PBFTSegment.proposals:type_name -> ordererpb.PBFTSegment.ProposalsEntry
	2, // 4: ordererpb.PBFTModule.segment:type_name -> ordererpb.PBFTSegment
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_ordererpb_ordererpb_proto_init() }
func file_ordererpb_ordererpb_proto_init() {
	if File_ordererpb_ordererpb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ordererpb_ordererpb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_ordererpb_ordererpb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_ordererpb_ordererpb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PBFTSegment); i {
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
		file_ordererpb_ordererpb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PBFTModule); i {
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
	file_ordererpb_ordererpb_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Event_Pbft)(nil),
	}
	file_ordererpb_ordererpb_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Message_Pbft)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ordererpb_ordererpb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ordererpb_ordererpb_proto_goTypes,
		DependencyIndexes: file_ordererpb_ordererpb_proto_depIdxs,
		MessageInfos:      file_ordererpb_ordererpb_proto_msgTypes,
	}.Build()
	File_ordererpb_ordererpb_proto = out.File
	file_ordererpb_ordererpb_proto_rawDesc = nil
	file_ordererpb_ordererpb_proto_goTypes = nil
	file_ordererpb_ordererpb_proto_depIdxs = nil
}
