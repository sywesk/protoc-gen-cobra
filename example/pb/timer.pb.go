// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: timer.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type TickRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Interval int32 `protobuf:"varint,1,opt,name=interval,proto3" json:"interval,omitempty"`
}

func (x *TickRequest) Reset() {
	*x = TickRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_timer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TickRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TickRequest) ProtoMessage() {}

func (x *TickRequest) ProtoReflect() protoreflect.Message {
	mi := &file_timer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TickRequest.ProtoReflect.Descriptor instead.
func (*TickRequest) Descriptor() ([]byte, []int) {
	return file_timer_proto_rawDescGZIP(), []int{0}
}

func (x *TickRequest) GetInterval() int32 {
	if x != nil {
		return x.Interval
	}
	return 0
}

type TickResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time string `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *TickResponse) Reset() {
	*x = TickResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_timer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TickResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TickResponse) ProtoMessage() {}

func (x *TickResponse) ProtoReflect() protoreflect.Message {
	mi := &file_timer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TickResponse.ProtoReflect.Descriptor instead.
func (*TickResponse) Descriptor() ([]byte, []int) {
	return file_timer_proto_rawDescGZIP(), []int{1}
}

func (x *TickResponse) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

var File_timer_proto protoreflect.FileDescriptor

var file_timer_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x22, 0x29, 0x0a, 0x0b, 0x54, 0x69, 0x63, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x22, 0x22, 0x0a, 0x0c, 0x54, 0x69, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x32, 0x3e, 0x0a, 0x05, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x12, 0x35,
	0x0a, 0x04, 0x54, 0x69, 0x63, 0x6b, 0x12, 0x14, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x54, 0x69, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_timer_proto_rawDescOnce sync.Once
	file_timer_proto_rawDescData = file_timer_proto_rawDesc
)

func file_timer_proto_rawDescGZIP() []byte {
	file_timer_proto_rawDescOnce.Do(func() {
		file_timer_proto_rawDescData = protoimpl.X.CompressGZIP(file_timer_proto_rawDescData)
	})
	return file_timer_proto_rawDescData
}

var file_timer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_timer_proto_goTypes = []interface{}{
	(*TickRequest)(nil),  // 0: example.TickRequest
	(*TickResponse)(nil), // 1: example.TickResponse
}
var file_timer_proto_depIdxs = []int32{
	0, // 0: example.Timer.Tick:input_type -> example.TickRequest
	1, // 1: example.Timer.Tick:output_type -> example.TickResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_timer_proto_init() }
func file_timer_proto_init() {
	if File_timer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_timer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TickRequest); i {
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
		file_timer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TickResponse); i {
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
			RawDescriptor: file_timer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_timer_proto_goTypes,
		DependencyIndexes: file_timer_proto_depIdxs,
		MessageInfos:      file_timer_proto_msgTypes,
	}.Build()
	File_timer_proto = out.File
	file_timer_proto_rawDesc = nil
	file_timer_proto_goTypes = nil
	file_timer_proto_depIdxs = nil
}
