// Code generated with goa v3.7.6, DO NOT EDIT.
//
// add protocol buffer definition
//
// Command:
// $ goa gen add/design

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: goadesign_goagen_add.proto

package addpb

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

type AddnumsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Left operand
	A int32 `protobuf:"zigzag32,1,opt,name=a,proto3" json:"a,omitempty"`
	// Right operand
	B int32 `protobuf:"zigzag32,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *AddnumsRequest) Reset() {
	*x = AddnumsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_add_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddnumsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddnumsRequest) ProtoMessage() {}

func (x *AddnumsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_add_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddnumsRequest.ProtoReflect.Descriptor instead.
func (*AddnumsRequest) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_add_proto_rawDescGZIP(), []int{0}
}

func (x *AddnumsRequest) GetA() int32 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *AddnumsRequest) GetB() int32 {
	if x != nil {
		return x.B
	}
	return 0
}

type AddnumsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field int32 `protobuf:"zigzag32,1,opt,name=field,proto3" json:"field,omitempty"`
}

func (x *AddnumsResponse) Reset() {
	*x = AddnumsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_add_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddnumsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddnumsResponse) ProtoMessage() {}

func (x *AddnumsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_add_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddnumsResponse.ProtoReflect.Descriptor instead.
func (*AddnumsResponse) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_add_proto_rawDescGZIP(), []int{1}
}

func (x *AddnumsResponse) GetField() int32 {
	if x != nil {
		return x.Field
	}
	return 0
}

var File_goadesign_goagen_add_proto protoreflect.FileDescriptor

var file_goadesign_goagen_add_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x67, 0x6f, 0x61, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x67, 0x6f, 0x61, 0x67,
	0x65, 0x6e, 0x5f, 0x61, 0x64, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x64,
	0x64, 0x22, 0x2c, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x6e, 0x75, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x11, 0x52, 0x01,
	0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x11, 0x52, 0x01, 0x62, 0x22,
	0x27, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x6e, 0x75, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x11, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x3b, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12,
	0x34, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x6e, 0x75, 0x6d, 0x73, 0x12, 0x13, 0x2e, 0x61, 0x64, 0x64,
	0x2e, 0x41, 0x64, 0x64, 0x6e, 0x75, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x61, 0x64, 0x64, 0x2e, 0x41, 0x64, 0x64, 0x6e, 0x75, 0x6d, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x61, 0x64, 0x64, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_goadesign_goagen_add_proto_rawDescOnce sync.Once
	file_goadesign_goagen_add_proto_rawDescData = file_goadesign_goagen_add_proto_rawDesc
)

func file_goadesign_goagen_add_proto_rawDescGZIP() []byte {
	file_goadesign_goagen_add_proto_rawDescOnce.Do(func() {
		file_goadesign_goagen_add_proto_rawDescData = protoimpl.X.CompressGZIP(file_goadesign_goagen_add_proto_rawDescData)
	})
	return file_goadesign_goagen_add_proto_rawDescData
}

var file_goadesign_goagen_add_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_goadesign_goagen_add_proto_goTypes = []interface{}{
	(*AddnumsRequest)(nil),  // 0: add.AddnumsRequest
	(*AddnumsResponse)(nil), // 1: add.AddnumsResponse
}
var file_goadesign_goagen_add_proto_depIdxs = []int32{
	0, // 0: add.Add.Addnums:input_type -> add.AddnumsRequest
	1, // 1: add.Add.Addnums:output_type -> add.AddnumsResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_goadesign_goagen_add_proto_init() }
func file_goadesign_goagen_add_proto_init() {
	if File_goadesign_goagen_add_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_goadesign_goagen_add_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddnumsRequest); i {
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
		file_goadesign_goagen_add_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddnumsResponse); i {
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
			RawDescriptor: file_goadesign_goagen_add_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_goadesign_goagen_add_proto_goTypes,
		DependencyIndexes: file_goadesign_goagen_add_proto_depIdxs,
		MessageInfos:      file_goadesign_goagen_add_proto_msgTypes,
	}.Build()
	File_goadesign_goagen_add_proto = out.File
	file_goadesign_goagen_add_proto_rawDesc = nil
	file_goadesign_goagen_add_proto_goTypes = nil
	file_goadesign_goagen_add_proto_depIdxs = nil
}
