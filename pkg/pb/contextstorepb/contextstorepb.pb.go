// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: contextstorepb/contextstorepb.proto

package contextstorepb

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

type Origin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemID uint64 `protobuf:"varint,1,opt,name=itemID,proto3" json:"itemID,omitempty"`
}

func (x *Origin) Reset() {
	*x = Origin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contextstorepb_contextstorepb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Origin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Origin) ProtoMessage() {}

func (x *Origin) ProtoReflect() protoreflect.Message {
	mi := &file_contextstorepb_contextstorepb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Origin.ProtoReflect.Descriptor instead.
func (*Origin) Descriptor() ([]byte, []int) {
	return file_contextstorepb_contextstorepb_proto_rawDescGZIP(), []int{0}
}

func (x *Origin) GetItemID() uint64 {
	if x != nil {
		return x.ItemID
	}
	return 0
}

var File_contextstorepb_contextstorepb_proto protoreflect.FileDescriptor

var file_contextstorepb_contextstorepb_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x70, 0x62,
	0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x70, 0x62, 0x1a, 0x1c, 0x6d, 0x69, 0x72, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x67,
	0x65, 0x6e, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x06, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x69,
	0x74, 0x65, 0x6d, 0x49, 0x44, 0x3a, 0x04, 0x80, 0xa6, 0x1d, 0x01, 0x42, 0x37, 0x5a, 0x35, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x63, 0x6f,
	0x69, 0x6e, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x6d, 0x69, 0x72, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contextstorepb_contextstorepb_proto_rawDescOnce sync.Once
	file_contextstorepb_contextstorepb_proto_rawDescData = file_contextstorepb_contextstorepb_proto_rawDesc
)

func file_contextstorepb_contextstorepb_proto_rawDescGZIP() []byte {
	file_contextstorepb_contextstorepb_proto_rawDescOnce.Do(func() {
		file_contextstorepb_contextstorepb_proto_rawDescData = protoimpl.X.CompressGZIP(file_contextstorepb_contextstorepb_proto_rawDescData)
	})
	return file_contextstorepb_contextstorepb_proto_rawDescData
}

var file_contextstorepb_contextstorepb_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_contextstorepb_contextstorepb_proto_goTypes = []interface{}{
	(*Origin)(nil), // 0: contextstorepb.Origin
}
var file_contextstorepb_contextstorepb_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_contextstorepb_contextstorepb_proto_init() }
func file_contextstorepb_contextstorepb_proto_init() {
	if File_contextstorepb_contextstorepb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contextstorepb_contextstorepb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Origin); i {
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
			RawDescriptor: file_contextstorepb_contextstorepb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_contextstorepb_contextstorepb_proto_goTypes,
		DependencyIndexes: file_contextstorepb_contextstorepb_proto_depIdxs,
		MessageInfos:      file_contextstorepb_contextstorepb_proto_msgTypes,
	}.Build()
	File_contextstorepb_contextstorepb_proto = out.File
	file_contextstorepb_contextstorepb_proto_rawDesc = nil
	file_contextstorepb_contextstorepb_proto_goTypes = nil
	file_contextstorepb_contextstorepb_proto_depIdxs = nil
}
