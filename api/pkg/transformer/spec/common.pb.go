// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: transformer/spec/common.proto

package spec

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

type FromTable_JsonFormat int32

const (
	FromTable_INVALID FromTable_JsonFormat = 0
	FromTable_RECORD  FromTable_JsonFormat = 1
	FromTable_VALUES  FromTable_JsonFormat = 2
	FromTable_SPLIT   FromTable_JsonFormat = 3
)

// Enum value maps for FromTable_JsonFormat.
var (
	FromTable_JsonFormat_name = map[int32]string{
		0: "INVALID",
		1: "RECORD",
		2: "VALUES",
		3: "SPLIT",
	}
	FromTable_JsonFormat_value = map[string]int32{
		"INVALID": 0,
		"RECORD":  1,
		"VALUES":  2,
		"SPLIT":   3,
	}
)

func (x FromTable_JsonFormat) Enum() *FromTable_JsonFormat {
	p := new(FromTable_JsonFormat)
	*p = x
	return p
}

func (x FromTable_JsonFormat) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FromTable_JsonFormat) Descriptor() protoreflect.EnumDescriptor {
	return file_transformer_spec_common_proto_enumTypes[0].Descriptor()
}

func (FromTable_JsonFormat) Type() protoreflect.EnumType {
	return &file_transformer_spec_common_proto_enumTypes[0]
}

func (x FromTable_JsonFormat) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FromTable_JsonFormat.Descriptor instead.
func (FromTable_JsonFormat) EnumDescriptor() ([]byte, []int) {
	return file_transformer_spec_common_proto_rawDescGZIP(), []int{0, 0}
}

type FromTable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TableName string               `protobuf:"bytes,1,opt,name=tableName,proto3" json:"tableName,omitempty"`
	Format    FromTable_JsonFormat `protobuf:"varint,2,opt,name=format,proto3,enum=merlin.transformer.FromTable_JsonFormat" json:"format,omitempty"`
}

func (x *FromTable) Reset() {
	*x = FromTable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transformer_spec_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FromTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FromTable) ProtoMessage() {}

func (x *FromTable) ProtoReflect() protoreflect.Message {
	mi := &file_transformer_spec_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FromTable.ProtoReflect.Descriptor instead.
func (*FromTable) Descriptor() ([]byte, []int) {
	return file_transformer_spec_common_proto_rawDescGZIP(), []int{0}
}

func (x *FromTable) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

func (x *FromTable) GetFormat() FromTable_JsonFormat {
	if x != nil {
		return x.Format
	}
	return FromTable_INVALID
}

type FromJson struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JsonPath     string `protobuf:"bytes,1,opt,name=jsonPath,proto3" json:"jsonPath,omitempty"`
	AddRowNumber bool   `protobuf:"varint,2,opt,name=addRowNumber,proto3" json:"addRowNumber,omitempty"`
}

func (x *FromJson) Reset() {
	*x = FromJson{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transformer_spec_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FromJson) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FromJson) ProtoMessage() {}

func (x *FromJson) ProtoReflect() protoreflect.Message {
	mi := &file_transformer_spec_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FromJson.ProtoReflect.Descriptor instead.
func (*FromJson) Descriptor() ([]byte, []int) {
	return file_transformer_spec_common_proto_rawDescGZIP(), []int{1}
}

func (x *FromJson) GetJsonPath() string {
	if x != nil {
		return x.JsonPath
	}
	return ""
}

func (x *FromJson) GetAddRowNumber() bool {
	if x != nil {
		return x.AddRowNumber
	}
	return false
}

var File_transformer_spec_common_proto protoreflect.FileDescriptor

var file_transformer_spec_common_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x2f, 0x73, 0x70,
	0x65, 0x63, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x12, 0x6d, 0x65, 0x72, 0x6c, 0x69, 0x6e, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72,
	0x6d, 0x65, 0x72, 0x22, 0xa9, 0x01, 0x0a, 0x09, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x40, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x28, 0x2e, 0x6d, 0x65, 0x72, 0x6c, 0x69, 0x6e, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f,
	0x72, 0x6d, 0x65, 0x72, 0x2e, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x4a,
	0x73, 0x6f, 0x6e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x22, 0x3c, 0x0a, 0x0a, 0x4a, 0x73, 0x6f, 0x6e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12,
	0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x56, 0x41, 0x4c, 0x55,
	0x45, 0x53, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x50, 0x4c, 0x49, 0x54, 0x10, 0x03, 0x22,
	0x4a, 0x0a, 0x08, 0x46, 0x72, 0x6f, 0x6d, 0x4a, 0x73, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6a,
	0x73, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6a,
	0x73, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x64, 0x64, 0x52, 0x6f,
	0x77, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x61,
	0x64, 0x64, 0x52, 0x6f, 0x77, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x2e, 0x5a, 0x2c, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6a, 0x65, 0x6b, 0x2f,
	0x6d, 0x65, 0x72, 0x6c, 0x69, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_transformer_spec_common_proto_rawDescOnce sync.Once
	file_transformer_spec_common_proto_rawDescData = file_transformer_spec_common_proto_rawDesc
)

func file_transformer_spec_common_proto_rawDescGZIP() []byte {
	file_transformer_spec_common_proto_rawDescOnce.Do(func() {
		file_transformer_spec_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_transformer_spec_common_proto_rawDescData)
	})
	return file_transformer_spec_common_proto_rawDescData
}

var file_transformer_spec_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_transformer_spec_common_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_transformer_spec_common_proto_goTypes = []interface{}{
	(FromTable_JsonFormat)(0), // 0: merlin.transformer.FromTable.JsonFormat
	(*FromTable)(nil),         // 1: merlin.transformer.FromTable
	(*FromJson)(nil),          // 2: merlin.transformer.FromJson
}
var file_transformer_spec_common_proto_depIdxs = []int32{
	0, // 0: merlin.transformer.FromTable.format:type_name -> merlin.transformer.FromTable.JsonFormat
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_transformer_spec_common_proto_init() }
func file_transformer_spec_common_proto_init() {
	if File_transformer_spec_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transformer_spec_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FromTable); i {
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
		file_transformer_spec_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FromJson); i {
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
			RawDescriptor: file_transformer_spec_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transformer_spec_common_proto_goTypes,
		DependencyIndexes: file_transformer_spec_common_proto_depIdxs,
		EnumInfos:         file_transformer_spec_common_proto_enumTypes,
		MessageInfos:      file_transformer_spec_common_proto_msgTypes,
	}.Build()
	File_transformer_spec_common_proto = out.File
	file_transformer_spec_common_proto_rawDesc = nil
	file_transformer_spec_common_proto_goTypes = nil
	file_transformer_spec_common_proto_depIdxs = nil
}
