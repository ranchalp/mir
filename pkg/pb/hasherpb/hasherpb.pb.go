// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: hasherpb/hasherpb.proto

package hasherpb

import (
	checkpointpb "github.com/filecoin-project/mir/pkg/pb/checkpointpb"
	commonpb "github.com/filecoin-project/mir/pkg/pb/commonpb"
	contextstorepb "github.com/filecoin-project/mir/pkg/pb/contextstorepb"
	dslpb "github.com/filecoin-project/mir/pkg/pb/dslpb"
	_ "github.com/filecoin-project/mir/pkg/pb/mir"
	ordererpb "github.com/filecoin-project/mir/pkg/pb/ordererpb"
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

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//	*Event_Request
	//	*Event_Result
	//	*Event_RequestOne
	//	*Event_ResultOne
	Type isEvent_Type `protobuf_oneof:"Type"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hasherpb_hasherpb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_hasherpb_hasherpb_proto_msgTypes[0]
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
	return file_hasherpb_hasherpb_proto_rawDescGZIP(), []int{0}
}

func (m *Event) GetType() isEvent_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *Event) GetRequest() *Request {
	if x, ok := x.GetType().(*Event_Request); ok {
		return x.Request
	}
	return nil
}

func (x *Event) GetResult() *Result {
	if x, ok := x.GetType().(*Event_Result); ok {
		return x.Result
	}
	return nil
}

func (x *Event) GetRequestOne() *RequestOne {
	if x, ok := x.GetType().(*Event_RequestOne); ok {
		return x.RequestOne
	}
	return nil
}

func (x *Event) GetResultOne() *ResultOne {
	if x, ok := x.GetType().(*Event_ResultOne); ok {
		return x.ResultOne
	}
	return nil
}

type isEvent_Type interface {
	isEvent_Type()
}

type Event_Request struct {
	Request *Request `protobuf:"bytes,1,opt,name=request,proto3,oneof"`
}

type Event_Result struct {
	Result *Result `protobuf:"bytes,2,opt,name=result,proto3,oneof"`
}

type Event_RequestOne struct {
	RequestOne *RequestOne `protobuf:"bytes,3,opt,name=request_one,json=requestOne,proto3,oneof"`
}

type Event_ResultOne struct {
	ResultOne *ResultOne `protobuf:"bytes,4,opt,name=result_one,json=resultOne,proto3,oneof"`
}

func (*Event_Request) isEvent_Type() {}

func (*Event_Result) isEvent_Type() {}

func (*Event_RequestOne) isEvent_Type() {}

func (*Event_ResultOne) isEvent_Type() {}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data   []*commonpb.HashData `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Origin *HashOrigin          `protobuf:"bytes,2,opt,name=origin,proto3" json:"origin,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hasherpb_hasherpb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_hasherpb_hasherpb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_hasherpb_hasherpb_proto_rawDescGZIP(), []int{1}
}

func (x *Request) GetData() []*commonpb.HashData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Request) GetOrigin() *HashOrigin {
	if x != nil {
		return x.Origin
	}
	return nil
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Digests [][]byte    `protobuf:"bytes,1,rep,name=digests,proto3" json:"digests,omitempty"`
	Origin  *HashOrigin `protobuf:"bytes,2,opt,name=origin,proto3" json:"origin,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hasherpb_hasherpb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_hasherpb_hasherpb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_hasherpb_hasherpb_proto_rawDescGZIP(), []int{2}
}

func (x *Result) GetDigests() [][]byte {
	if x != nil {
		return x.Digests
	}
	return nil
}

func (x *Result) GetOrigin() *HashOrigin {
	if x != nil {
		return x.Origin
	}
	return nil
}

type RequestOne struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data   *commonpb.HashData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Origin *HashOrigin        `protobuf:"bytes,2,opt,name=origin,proto3" json:"origin,omitempty"`
}

func (x *RequestOne) Reset() {
	*x = RequestOne{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hasherpb_hasherpb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestOne) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestOne) ProtoMessage() {}

func (x *RequestOne) ProtoReflect() protoreflect.Message {
	mi := &file_hasherpb_hasherpb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestOne.ProtoReflect.Descriptor instead.
func (*RequestOne) Descriptor() ([]byte, []int) {
	return file_hasherpb_hasherpb_proto_rawDescGZIP(), []int{3}
}

func (x *RequestOne) GetData() *commonpb.HashData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *RequestOne) GetOrigin() *HashOrigin {
	if x != nil {
		return x.Origin
	}
	return nil
}

type ResultOne struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Digest []byte      `protobuf:"bytes,1,opt,name=digest,proto3" json:"digest,omitempty"`
	Origin *HashOrigin `protobuf:"bytes,2,opt,name=origin,proto3" json:"origin,omitempty"`
}

func (x *ResultOne) Reset() {
	*x = ResultOne{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hasherpb_hasherpb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultOne) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultOne) ProtoMessage() {}

func (x *ResultOne) ProtoReflect() protoreflect.Message {
	mi := &file_hasherpb_hasherpb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultOne.ProtoReflect.Descriptor instead.
func (*ResultOne) Descriptor() ([]byte, []int) {
	return file_hasherpb_hasherpb_proto_rawDescGZIP(), []int{4}
}

func (x *ResultOne) GetDigest() []byte {
	if x != nil {
		return x.Digest
	}
	return nil
}

func (x *ResultOne) GetOrigin() *HashOrigin {
	if x != nil {
		return x.Origin
	}
	return nil
}

type HashOrigin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Module string `protobuf:"bytes,1,opt,name=module,proto3" json:"module,omitempty"`
	// Types that are assignable to Type:
	//	*HashOrigin_ContextStore
	//	*HashOrigin_Request
	//	*HashOrigin_Dsl
	//	*HashOrigin_Checkpoint
	//	*HashOrigin_Sb
	Type isHashOrigin_Type `protobuf_oneof:"type"`
}

func (x *HashOrigin) Reset() {
	*x = HashOrigin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hasherpb_hasherpb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HashOrigin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HashOrigin) ProtoMessage() {}

func (x *HashOrigin) ProtoReflect() protoreflect.Message {
	mi := &file_hasherpb_hasherpb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HashOrigin.ProtoReflect.Descriptor instead.
func (*HashOrigin) Descriptor() ([]byte, []int) {
	return file_hasherpb_hasherpb_proto_rawDescGZIP(), []int{5}
}

func (x *HashOrigin) GetModule() string {
	if x != nil {
		return x.Module
	}
	return ""
}

func (m *HashOrigin) GetType() isHashOrigin_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *HashOrigin) GetContextStore() *contextstorepb.Origin {
	if x, ok := x.GetType().(*HashOrigin_ContextStore); ok {
		return x.ContextStore
	}
	return nil
}

func (x *HashOrigin) GetRequest() *requestpb.Request {
	if x, ok := x.GetType().(*HashOrigin_Request); ok {
		return x.Request
	}
	return nil
}

func (x *HashOrigin) GetDsl() *dslpb.Origin {
	if x, ok := x.GetType().(*HashOrigin_Dsl); ok {
		return x.Dsl
	}
	return nil
}

func (x *HashOrigin) GetCheckpoint() *checkpointpb.HashOrigin {
	if x, ok := x.GetType().(*HashOrigin_Checkpoint); ok {
		return x.Checkpoint
	}
	return nil
}

func (x *HashOrigin) GetSb() *ordererpb.HashOrigin {
	if x, ok := x.GetType().(*HashOrigin_Sb); ok {
		return x.Sb
	}
	return nil
}

type isHashOrigin_Type interface {
	isHashOrigin_Type()
}

type HashOrigin_ContextStore struct {
	ContextStore *contextstorepb.Origin `protobuf:"bytes,2,opt,name=context_store,json=contextStore,proto3,oneof"`
}

type HashOrigin_Request struct {
	Request *requestpb.Request `protobuf:"bytes,3,opt,name=request,proto3,oneof"`
}

type HashOrigin_Dsl struct {
	Dsl *dslpb.Origin `protobuf:"bytes,4,opt,name=dsl,proto3,oneof"`
}

type HashOrigin_Checkpoint struct {
	Checkpoint *checkpointpb.HashOrigin `protobuf:"bytes,5,opt,name=checkpoint,proto3,oneof"`
}

type HashOrigin_Sb struct {
	Sb *ordererpb.HashOrigin `protobuf:"bytes,6,opt,name=sb,proto3,oneof"`
}

func (*HashOrigin_ContextStore) isHashOrigin_Type() {}

func (*HashOrigin_Request) isHashOrigin_Type() {}

func (*HashOrigin_Dsl) isHashOrigin_Type() {}

func (*HashOrigin_Checkpoint) isHashOrigin_Type() {}

func (*HashOrigin_Sb) isHashOrigin_Type() {}

var File_hasherpb_hasherpb_proto protoreflect.FileDescriptor

var file_hasherpb_hasherpb_proto_rawDesc = []byte{
	0x0a, 0x17, 0x68, 0x61, 0x73, 0x68, 0x65, 0x72, 0x70, 0x62, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x65,
	0x72, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x68, 0x61, 0x73, 0x68, 0x65,
	0x72, 0x70, 0x62, 0x1a, 0x17, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x62, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x64, 0x73,
	0x6c, 0x70, 0x62, 0x2f, 0x64, 0x73, 0x6c, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x70, 0x62, 0x2f, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x6d, 0x69, 0x72,
	0x2f, 0x63, 0x6f, 0x64, 0x65, 0x67, 0x65, 0x6e, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe5, 0x01, 0x0a, 0x05, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x2d, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x65, 0x72, 0x70, 0x62, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2a, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x37,
	0x0a, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x6e, 0x65, 0x48, 0x00, 0x52, 0x0a, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x4f, 0x6e, 0x65, 0x12, 0x34, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x5f, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x68, 0x61,
	0x73, 0x68, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4f, 0x6e, 0x65,
	0x48, 0x00, 0x52, 0x09, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4f, 0x6e, 0x65, 0x3a, 0x04, 0x90,
	0xa6, 0x1d, 0x01, 0x42, 0x0c, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x04, 0x80, 0xa6, 0x1d,
	0x01, 0x22, 0x6b, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x70, 0x62, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x32, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x65, 0x72, 0x70, 0x62, 0x2e,
	0x48, 0x61, 0x73, 0x68, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x42, 0x04, 0x98, 0xa6, 0x1d, 0x01,
	0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x3a, 0x04, 0x98, 0xa6, 0x1d, 0x01, 0x22, 0x5c,
	0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x69, 0x67, 0x65,
	0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x07, 0x64, 0x69, 0x67, 0x65, 0x73,
	0x74, 0x73, 0x12, 0x32, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x48, 0x61,
	0x73, 0x68, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x42, 0x04, 0xa0, 0xa6, 0x1d, 0x01, 0x52, 0x06,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x3a, 0x04, 0x98, 0xa6, 0x1d, 0x01, 0x22, 0x6e, 0x0a, 0x0a,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x6e, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x70, 0x62, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x32, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x48, 0x61,
	0x73, 0x68, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x42, 0x04, 0x98, 0xa6, 0x1d, 0x01, 0x52, 0x06,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x3a, 0x04, 0x98, 0xa6, 0x1d, 0x01, 0x22, 0x5d, 0x0a, 0x09,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4f, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69, 0x67,
	0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73,
	0x74, 0x12, 0x32, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x48, 0x61, 0x73,
	0x68, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x42, 0x04, 0xa0, 0xa6, 0x1d, 0x01, 0x52, 0x06, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x3a, 0x04, 0x98, 0xa6, 0x1d, 0x01, 0x22, 0xe1, 0x02, 0x0a, 0x0a,
	0x48, 0x61, 0x73, 0x68, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x4e, 0x0a, 0x06, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x36, 0x82, 0xa6, 0x1d, 0x32,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x63,
	0x6f, 0x69, 0x6e, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x6d, 0x69, 0x72, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x49, 0x44, 0x52, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x3d, 0x0a, 0x0d, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x70, 0x62, 0x2e, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x48, 0x00, 0x52, 0x0c, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00,
	0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x03, 0x64, 0x73, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x64, 0x73, 0x6c, 0x70, 0x62, 0x2e, 0x4f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x48, 0x00, 0x52, 0x03, 0x64, 0x73, 0x6c, 0x12, 0x3a, 0x0a, 0x0a,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x70, 0x62, 0x2e,
	0x48, 0x61, 0x73, 0x68, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x48, 0x00, 0x52, 0x0a, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x02, 0x73, 0x62, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x70, 0x62,
	0x2e, 0x48, 0x61, 0x73, 0x68, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x48, 0x00, 0x52, 0x02, 0x73,
	0x62, 0x3a, 0x04, 0x80, 0xa6, 0x1d, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x42,
	0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69,
	0x6c, 0x65, 0x63, 0x6f, 0x69, 0x6e, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x6d,
	0x69, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x65, 0x72,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hasherpb_hasherpb_proto_rawDescOnce sync.Once
	file_hasherpb_hasherpb_proto_rawDescData = file_hasherpb_hasherpb_proto_rawDesc
)

func file_hasherpb_hasherpb_proto_rawDescGZIP() []byte {
	file_hasherpb_hasherpb_proto_rawDescOnce.Do(func() {
		file_hasherpb_hasherpb_proto_rawDescData = protoimpl.X.CompressGZIP(file_hasherpb_hasherpb_proto_rawDescData)
	})
	return file_hasherpb_hasherpb_proto_rawDescData
}

var file_hasherpb_hasherpb_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_hasherpb_hasherpb_proto_goTypes = []interface{}{
	(*Event)(nil),                   // 0: hasherpb.Event
	(*Request)(nil),                 // 1: hasherpb.Request
	(*Result)(nil),                  // 2: hasherpb.Result
	(*RequestOne)(nil),              // 3: hasherpb.RequestOne
	(*ResultOne)(nil),               // 4: hasherpb.ResultOne
	(*HashOrigin)(nil),              // 5: hasherpb.HashOrigin
	(*commonpb.HashData)(nil),       // 6: commonpb.HashData
	(*contextstorepb.Origin)(nil),   // 7: contextstorepb.Origin
	(*requestpb.Request)(nil),       // 8: requestpb.Request
	(*dslpb.Origin)(nil),            // 9: dslpb.Origin
	(*checkpointpb.HashOrigin)(nil), // 10: checkpointpb.HashOrigin
	(*ordererpb.HashOrigin)(nil),    // 11: ordererpb.HashOrigin
}
var file_hasherpb_hasherpb_proto_depIdxs = []int32{
	1,  // 0: hasherpb.Event.request:type_name -> hasherpb.Request
	2,  // 1: hasherpb.Event.result:type_name -> hasherpb.Result
	3,  // 2: hasherpb.Event.request_one:type_name -> hasherpb.RequestOne
	4,  // 3: hasherpb.Event.result_one:type_name -> hasherpb.ResultOne
	6,  // 4: hasherpb.Request.data:type_name -> commonpb.HashData
	5,  // 5: hasherpb.Request.origin:type_name -> hasherpb.HashOrigin
	5,  // 6: hasherpb.Result.origin:type_name -> hasherpb.HashOrigin
	6,  // 7: hasherpb.RequestOne.data:type_name -> commonpb.HashData
	5,  // 8: hasherpb.RequestOne.origin:type_name -> hasherpb.HashOrigin
	5,  // 9: hasherpb.ResultOne.origin:type_name -> hasherpb.HashOrigin
	7,  // 10: hasherpb.HashOrigin.context_store:type_name -> contextstorepb.Origin
	8,  // 11: hasherpb.HashOrigin.request:type_name -> requestpb.Request
	9,  // 12: hasherpb.HashOrigin.dsl:type_name -> dslpb.Origin
	10, // 13: hasherpb.HashOrigin.checkpoint:type_name -> checkpointpb.HashOrigin
	11, // 14: hasherpb.HashOrigin.sb:type_name -> ordererpb.HashOrigin
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_hasherpb_hasherpb_proto_init() }
func file_hasherpb_hasherpb_proto_init() {
	if File_hasherpb_hasherpb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hasherpb_hasherpb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_hasherpb_hasherpb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_hasherpb_hasherpb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_hasherpb_hasherpb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestOne); i {
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
		file_hasherpb_hasherpb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultOne); i {
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
		file_hasherpb_hasherpb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HashOrigin); i {
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
	file_hasherpb_hasherpb_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Event_Request)(nil),
		(*Event_Result)(nil),
		(*Event_RequestOne)(nil),
		(*Event_ResultOne)(nil),
	}
	file_hasherpb_hasherpb_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*HashOrigin_ContextStore)(nil),
		(*HashOrigin_Request)(nil),
		(*HashOrigin_Dsl)(nil),
		(*HashOrigin_Checkpoint)(nil),
		(*HashOrigin_Sb)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_hasherpb_hasherpb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_hasherpb_hasherpb_proto_goTypes,
		DependencyIndexes: file_hasherpb_hasherpb_proto_depIdxs,
		MessageInfos:      file_hasherpb_hasherpb_proto_msgTypes,
	}.Build()
	File_hasherpb_hasherpb_proto = out.File
	file_hasherpb_hasherpb_proto_rawDesc = nil
	file_hasherpb_hasherpb_proto_goTypes = nil
	file_hasherpb_hasherpb_proto_depIdxs = nil
}
