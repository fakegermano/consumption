// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: proto/meterusage.proto

package proto

import (
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

type MeterUsageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *MeterUsageRequest) Reset() {
	*x = MeterUsageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_meterusage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeterUsageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeterUsageRequest) ProtoMessage() {}

func (x *MeterUsageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_meterusage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeterUsageRequest.ProtoReflect.Descriptor instead.
func (*MeterUsageRequest) Descriptor() ([]byte, []int) {
	return file_proto_meterusage_proto_rawDescGZIP(), []int{0}
}

func (x *MeterUsageRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *MeterUsageRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type MeterUsageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Usages []*MeterUsage `protobuf:"bytes,1,rep,name=usages,proto3" json:"usages,omitempty"`
}

func (x *MeterUsageResponse) Reset() {
	*x = MeterUsageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_meterusage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeterUsageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeterUsageResponse) ProtoMessage() {}

func (x *MeterUsageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_meterusage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeterUsageResponse.ProtoReflect.Descriptor instead.
func (*MeterUsageResponse) Descriptor() ([]byte, []int) {
	return file_proto_meterusage_proto_rawDescGZIP(), []int{1}
}

func (x *MeterUsageResponse) GetUsages() []*MeterUsage {
	if x != nil {
		return x.Usages
	}
	return nil
}

type MeterUsage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time  int64   `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	Usage float64 `protobuf:"fixed64,2,opt,name=usage,proto3" json:"usage,omitempty"`
}

func (x *MeterUsage) Reset() {
	*x = MeterUsage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_meterusage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeterUsage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeterUsage) ProtoMessage() {}

func (x *MeterUsage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_meterusage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeterUsage.ProtoReflect.Descriptor instead.
func (*MeterUsage) Descriptor() ([]byte, []int) {
	return file_proto_meterusage_proto_rawDescGZIP(), []int{2}
}

func (x *MeterUsage) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *MeterUsage) GetUsage() float64 {
	if x != nil {
		return x.Usage
	}
	return 0
}

var File_proto_meterusage_proto protoreflect.FileDescriptor

var file_proto_meterusage_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x75, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x41, 0x0a, 0x11, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x22, 0x3f, 0x0a, 0x12, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x55, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x75, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x06, 0x75, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x22, 0x36, 0x0a, 0x0a, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x55, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x32, 0x51, 0x0a, 0x11, 0x4d,
	0x65, 0x74, 0x65, 0x72, 0x55, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3c, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4d, 0x65, 0x74, 0x65, 0x72, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x55,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2a,
	0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x61, 0x6b,
	0x65, 0x67, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_meterusage_proto_rawDescOnce sync.Once
	file_proto_meterusage_proto_rawDescData = file_proto_meterusage_proto_rawDesc
)

func file_proto_meterusage_proto_rawDescGZIP() []byte {
	file_proto_meterusage_proto_rawDescOnce.Do(func() {
		file_proto_meterusage_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_meterusage_proto_rawDescData)
	})
	return file_proto_meterusage_proto_rawDescData
}

var file_proto_meterusage_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_meterusage_proto_goTypes = []interface{}{
	(*MeterUsageRequest)(nil),  // 0: proto.MeterUsageRequest
	(*MeterUsageResponse)(nil), // 1: proto.MeterUsageResponse
	(*MeterUsage)(nil),         // 2: proto.MeterUsage
}
var file_proto_meterusage_proto_depIdxs = []int32{
	2, // 0: proto.MeterUsageResponse.usages:type_name -> proto.MeterUsage
	0, // 1: proto.MeterUsageService.Get:input_type -> proto.MeterUsageRequest
	1, // 2: proto.MeterUsageService.Get:output_type -> proto.MeterUsageResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_meterusage_proto_init() }
func file_proto_meterusage_proto_init() {
	if File_proto_meterusage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_meterusage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeterUsageRequest); i {
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
		file_proto_meterusage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeterUsageResponse); i {
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
		file_proto_meterusage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeterUsage); i {
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
			RawDescriptor: file_proto_meterusage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_meterusage_proto_goTypes,
		DependencyIndexes: file_proto_meterusage_proto_depIdxs,
		MessageInfos:      file_proto_meterusage_proto_msgTypes,
	}.Build()
	File_proto_meterusage_proto = out.File
	file_proto_meterusage_proto_rawDesc = nil
	file_proto_meterusage_proto_goTypes = nil
	file_proto_meterusage_proto_depIdxs = nil
}
