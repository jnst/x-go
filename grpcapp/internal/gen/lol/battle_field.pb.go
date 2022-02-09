// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: lol/battle_field.proto

package lol

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

type BattleField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BattleFieldId int32  `protobuf:"varint,1,opt,name=battle_field_id,json=battleFieldId,proto3" json:"battle_field_id,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *BattleField) Reset() {
	*x = BattleField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lol_battle_field_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BattleField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BattleField) ProtoMessage() {}

func (x *BattleField) ProtoReflect() protoreflect.Message {
	mi := &file_lol_battle_field_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BattleField.ProtoReflect.Descriptor instead.
func (*BattleField) Descriptor() ([]byte, []int) {
	return file_lol_battle_field_proto_rawDescGZIP(), []int{0}
}

func (x *BattleField) GetBattleFieldId() int32 {
	if x != nil {
		return x.BattleFieldId
	}
	return 0
}

func (x *BattleField) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BattleField) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_lol_battle_field_proto protoreflect.FileDescriptor

var file_lol_battle_field_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6c, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65,
	0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x6b, 0x0a, 0x0b, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d,
	0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x0f, 0x5a, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x6f, 0x6c,
	0x3b, 0x6c, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lol_battle_field_proto_rawDescOnce sync.Once
	file_lol_battle_field_proto_rawDescData = file_lol_battle_field_proto_rawDesc
)

func file_lol_battle_field_proto_rawDescGZIP() []byte {
	file_lol_battle_field_proto_rawDescOnce.Do(func() {
		file_lol_battle_field_proto_rawDescData = protoimpl.X.CompressGZIP(file_lol_battle_field_proto_rawDescData)
	})
	return file_lol_battle_field_proto_rawDescData
}

var file_lol_battle_field_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_lol_battle_field_proto_goTypes = []interface{}{
	(*BattleField)(nil), // 0: battle_field.BattleField
}
var file_lol_battle_field_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_lol_battle_field_proto_init() }
func file_lol_battle_field_proto_init() {
	if File_lol_battle_field_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lol_battle_field_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BattleField); i {
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
			RawDescriptor: file_lol_battle_field_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_lol_battle_field_proto_goTypes,
		DependencyIndexes: file_lol_battle_field_proto_depIdxs,
		MessageInfos:      file_lol_battle_field_proto_msgTypes,
	}.Build()
	File_lol_battle_field_proto = out.File
	file_lol_battle_field_proto_rawDesc = nil
	file_lol_battle_field_proto_goTypes = nil
	file_lol_battle_field_proto_depIdxs = nil
}
