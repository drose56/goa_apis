// Code generated with goa v3.7.6, DO NOT EDIT.
//
// poke protocol buffer definition
//
// Command:
// $ goa gen poke/design

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: goadesign_goagen_poke.proto

package pokepb

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

type PokemonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// first name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *PokemonRequest) Reset() {
	*x = PokemonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_poke_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PokemonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PokemonRequest) ProtoMessage() {}

func (x *PokemonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_poke_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PokemonRequest.ProtoReflect.Descriptor instead.
func (*PokemonRequest) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_poke_proto_rawDescGZIP(), []int{0}
}

func (x *PokemonRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type PokemonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
}

func (x *PokemonResponse) Reset() {
	*x = PokemonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_poke_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PokemonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PokemonResponse) ProtoMessage() {}

func (x *PokemonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_poke_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PokemonResponse.ProtoReflect.Descriptor instead.
func (*PokemonResponse) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_poke_proto_rawDescGZIP(), []int{1}
}

func (x *PokemonResponse) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

var File_goadesign_goagen_poke_proto protoreflect.FileDescriptor

var file_goadesign_goagen_poke_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x67, 0x6f, 0x61, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x67, 0x6f, 0x61, 0x67,
	0x65, 0x6e, 0x5f, 0x70, 0x6f, 0x6b, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x70,
	0x6f, 0x6b, 0x65, 0x22, 0x24, 0x0a, 0x0e, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x27, 0x0a, 0x0f, 0x50, 0x6f, 0x6b,
	0x65, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x32, 0x3e, 0x0a, 0x04, 0x50, 0x6f, 0x6b, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x50, 0x6f,
	0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x12, 0x14, 0x2e, 0x70, 0x6f, 0x6b, 0x65, 0x2e, 0x50, 0x6f, 0x6b,
	0x65, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x6f,
	0x6b, 0x65, 0x2e, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x70, 0x6f, 0x6b, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_goadesign_goagen_poke_proto_rawDescOnce sync.Once
	file_goadesign_goagen_poke_proto_rawDescData = file_goadesign_goagen_poke_proto_rawDesc
)

func file_goadesign_goagen_poke_proto_rawDescGZIP() []byte {
	file_goadesign_goagen_poke_proto_rawDescOnce.Do(func() {
		file_goadesign_goagen_poke_proto_rawDescData = protoimpl.X.CompressGZIP(file_goadesign_goagen_poke_proto_rawDescData)
	})
	return file_goadesign_goagen_poke_proto_rawDescData
}

var file_goadesign_goagen_poke_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_goadesign_goagen_poke_proto_goTypes = []interface{}{
	(*PokemonRequest)(nil),  // 0: poke.PokemonRequest
	(*PokemonResponse)(nil), // 1: poke.PokemonResponse
}
var file_goadesign_goagen_poke_proto_depIdxs = []int32{
	0, // 0: poke.Poke.Pokemon:input_type -> poke.PokemonRequest
	1, // 1: poke.Poke.Pokemon:output_type -> poke.PokemonResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_goadesign_goagen_poke_proto_init() }
func file_goadesign_goagen_poke_proto_init() {
	if File_goadesign_goagen_poke_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_goadesign_goagen_poke_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PokemonRequest); i {
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
		file_goadesign_goagen_poke_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PokemonResponse); i {
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
			RawDescriptor: file_goadesign_goagen_poke_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_goadesign_goagen_poke_proto_goTypes,
		DependencyIndexes: file_goadesign_goagen_poke_proto_depIdxs,
		MessageInfos:      file_goadesign_goagen_poke_proto_msgTypes,
	}.Build()
	File_goadesign_goagen_poke_proto = out.File
	file_goadesign_goagen_poke_proto_rawDesc = nil
	file_goadesign_goagen_poke_proto_goTypes = nil
	file_goadesign_goagen_poke_proto_depIdxs = nil
}
