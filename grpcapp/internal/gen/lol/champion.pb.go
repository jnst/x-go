// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: lol/champion.proto

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

type Champion_ChampionType int32

const (
	Champion_UNKNOWN  Champion_ChampionType = 0
	Champion_MARKSMAN Champion_ChampionType = 1
	Champion_MAGE     Champion_ChampionType = 2
	Champion_ASSASSIN Champion_ChampionType = 3
	Champion_TANK     Champion_ChampionType = 4
	Champion_FIGHTER  Champion_ChampionType = 5
	Champion_SUPPORT  Champion_ChampionType = 6
)

// Enum value maps for Champion_ChampionType.
var (
	Champion_ChampionType_name = map[int32]string{
		0: "UNKNOWN",
		1: "MARKSMAN",
		2: "MAGE",
		3: "ASSASSIN",
		4: "TANK",
		5: "FIGHTER",
		6: "SUPPORT",
	}
	Champion_ChampionType_value = map[string]int32{
		"UNKNOWN":  0,
		"MARKSMAN": 1,
		"MAGE":     2,
		"ASSASSIN": 3,
		"TANK":     4,
		"FIGHTER":  5,
		"SUPPORT":  6,
	}
)

func (x Champion_ChampionType) Enum() *Champion_ChampionType {
	p := new(Champion_ChampionType)
	*p = x
	return p
}

func (x Champion_ChampionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Champion_ChampionType) Descriptor() protoreflect.EnumDescriptor {
	return file_lol_champion_proto_enumTypes[0].Descriptor()
}

func (Champion_ChampionType) Type() protoreflect.EnumType {
	return &file_lol_champion_proto_enumTypes[0]
}

func (x Champion_ChampionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Champion_ChampionType.Descriptor instead.
func (Champion_ChampionType) EnumDescriptor() ([]byte, []int) {
	return file_lol_champion_proto_rawDescGZIP(), []int{0, 0}
}

type Champion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChampionId int32                 `protobuf:"varint,1,opt,name=champion_id,json=championId,proto3" json:"champion_id,omitempty"`
	Type       Champion_ChampionType `protobuf:"varint,2,opt,name=type,proto3,enum=champion.Champion_ChampionType" json:"type,omitempty"`
	Name       string                `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Message    string                `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Champion) Reset() {
	*x = Champion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lol_champion_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Champion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Champion) ProtoMessage() {}

func (x *Champion) ProtoReflect() protoreflect.Message {
	mi := &file_lol_champion_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Champion.ProtoReflect.Descriptor instead.
func (*Champion) Descriptor() ([]byte, []int) {
	return file_lol_champion_proto_rawDescGZIP(), []int{0}
}

func (x *Champion) GetChampionId() int32 {
	if x != nil {
		return x.ChampionId
	}
	return 0
}

func (x *Champion) GetType() Champion_ChampionType {
	if x != nil {
		return x.Type
	}
	return Champion_UNKNOWN
}

func (x *Champion) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Champion) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_lol_champion_proto protoreflect.FileDescriptor

var file_lol_champion_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6c, 0x6f, 0x6c, 0x2f, 0x63, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x22, 0xf5,
	0x01, 0x0a, 0x08, 0x43, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x63, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x63, 0x68, 0x61,
	0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x2e, 0x43,
	0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x65, 0x0a, 0x0c, 0x43, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08,
	0x4d, 0x41, 0x52, 0x4b, 0x53, 0x4d, 0x41, 0x4e, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x41,
	0x47, 0x45, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x53, 0x53, 0x41, 0x53, 0x53, 0x49, 0x4e,
	0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x41, 0x4e, 0x4b, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07,
	0x46, 0x49, 0x47, 0x48, 0x54, 0x45, 0x52, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x50,
	0x50, 0x4f, 0x52, 0x54, 0x10, 0x06, 0x42, 0x0f, 0x5a, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x6c, 0x6f, 0x6c, 0x3b, 0x6c, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lol_champion_proto_rawDescOnce sync.Once
	file_lol_champion_proto_rawDescData = file_lol_champion_proto_rawDesc
)

func file_lol_champion_proto_rawDescGZIP() []byte {
	file_lol_champion_proto_rawDescOnce.Do(func() {
		file_lol_champion_proto_rawDescData = protoimpl.X.CompressGZIP(file_lol_champion_proto_rawDescData)
	})
	return file_lol_champion_proto_rawDescData
}

var file_lol_champion_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_lol_champion_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_lol_champion_proto_goTypes = []interface{}{
	(Champion_ChampionType)(0), // 0: champion.Champion.ChampionType
	(*Champion)(nil),           // 1: champion.Champion
}
var file_lol_champion_proto_depIdxs = []int32{
	0, // 0: champion.Champion.type:type_name -> champion.Champion.ChampionType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_lol_champion_proto_init() }
func file_lol_champion_proto_init() {
	if File_lol_champion_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lol_champion_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Champion); i {
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
			RawDescriptor: file_lol_champion_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_lol_champion_proto_goTypes,
		DependencyIndexes: file_lol_champion_proto_depIdxs,
		EnumInfos:         file_lol_champion_proto_enumTypes,
		MessageInfos:      file_lol_champion_proto_msgTypes,
	}.Build()
	File_lol_champion_proto = out.File
	file_lol_champion_proto_rawDesc = nil
	file_lol_champion_proto_goTypes = nil
	file_lol_champion_proto_depIdxs = nil
}