// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: lol/lol.proto

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

type GetChampionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChampionId int32 `protobuf:"varint,1,opt,name=champion_id,json=championId,proto3" json:"champion_id,omitempty"`
}

func (x *GetChampionRequest) Reset() {
	*x = GetChampionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lol_lol_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChampionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChampionRequest) ProtoMessage() {}

func (x *GetChampionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lol_lol_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChampionRequest.ProtoReflect.Descriptor instead.
func (*GetChampionRequest) Descriptor() ([]byte, []int) {
	return file_lol_lol_proto_rawDescGZIP(), []int{0}
}

func (x *GetChampionRequest) GetChampionId() int32 {
	if x != nil {
		return x.ChampionId
	}
	return 0
}

type GetChampionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Champion *Champion `protobuf:"bytes,1,opt,name=champion,proto3" json:"champion,omitempty"`
}

func (x *GetChampionResponse) Reset() {
	*x = GetChampionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lol_lol_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChampionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChampionResponse) ProtoMessage() {}

func (x *GetChampionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lol_lol_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChampionResponse.ProtoReflect.Descriptor instead.
func (*GetChampionResponse) Descriptor() ([]byte, []int) {
	return file_lol_lol_proto_rawDescGZIP(), []int{1}
}

func (x *GetChampionResponse) GetChampion() *Champion {
	if x != nil {
		return x.Champion
	}
	return nil
}

type ListChampionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListChampionsRequest) Reset() {
	*x = ListChampionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lol_lol_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListChampionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListChampionsRequest) ProtoMessage() {}

func (x *ListChampionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lol_lol_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListChampionsRequest.ProtoReflect.Descriptor instead.
func (*ListChampionsRequest) Descriptor() ([]byte, []int) {
	return file_lol_lol_proto_rawDescGZIP(), []int{2}
}

type ListChampionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Champions []*Champion `protobuf:"bytes,1,rep,name=champions,proto3" json:"champions,omitempty"`
}

func (x *ListChampionsResponse) Reset() {
	*x = ListChampionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lol_lol_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListChampionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListChampionsResponse) ProtoMessage() {}

func (x *ListChampionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lol_lol_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListChampionsResponse.ProtoReflect.Descriptor instead.
func (*ListChampionsResponse) Descriptor() ([]byte, []int) {
	return file_lol_lol_proto_rawDescGZIP(), []int{3}
}

func (x *ListChampionsResponse) GetChampions() []*Champion {
	if x != nil {
		return x.Champions
	}
	return nil
}

type GetBattleFieldRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetBattleFieldRequest) Reset() {
	*x = GetBattleFieldRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lol_lol_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBattleFieldRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBattleFieldRequest) ProtoMessage() {}

func (x *GetBattleFieldRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lol_lol_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBattleFieldRequest.ProtoReflect.Descriptor instead.
func (*GetBattleFieldRequest) Descriptor() ([]byte, []int) {
	return file_lol_lol_proto_rawDescGZIP(), []int{4}
}

type GetBattleFieldResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BattleField *BattleField `protobuf:"bytes,1,opt,name=battle_field,json=battleField,proto3" json:"battle_field,omitempty"`
}

func (x *GetBattleFieldResponse) Reset() {
	*x = GetBattleFieldResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lol_lol_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBattleFieldResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBattleFieldResponse) ProtoMessage() {}

func (x *GetBattleFieldResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lol_lol_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBattleFieldResponse.ProtoReflect.Descriptor instead.
func (*GetBattleFieldResponse) Descriptor() ([]byte, []int) {
	return file_lol_lol_proto_rawDescGZIP(), []int{5}
}

func (x *GetBattleFieldResponse) GetBattleField() *BattleField {
	if x != nil {
		return x.BattleField
	}
	return nil
}

var File_lol_lol_proto protoreflect.FileDescriptor

var file_lol_lol_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6c, 0x6f, 0x6c, 0x2f, 0x6c, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x6c, 0x6f, 0x6c, 0x1a, 0x16, 0x6c, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65,
	0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x6c, 0x6f,
	0x6c, 0x2f, 0x63, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x35, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6d, 0x70, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x68, 0x61,
	0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x45, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x68,
	0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e,
	0x0a, 0x08, 0x63, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x63, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x68, 0x61, 0x6d,
	0x70, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x63, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x22, 0x16,
	0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x49, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68,
	0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x30, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x68,
	0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0x17, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x56, 0x0a, 0x16, 0x47, 0x65,
	0x74, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x61, 0x74,
	0x74, 0x6c, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x0b, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x32, 0xe6, 0x01, 0x0a, 0x0f, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x4f, 0x66, 0x4c,
	0x65, 0x67, 0x65, 0x6e, 0x64, 0x73, 0x12, 0x40, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61,
	0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x2e, 0x6c, 0x6f, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x6c, 0x6f, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x19, 0x2e, 0x6c, 0x6f, 0x6c, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6c, 0x6f, 0x6c, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x49, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x12, 0x1a, 0x2e, 0x6c, 0x6f, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x74, 0x74,
	0x6c, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x6c, 0x6f, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0f, 0x5a, 0x0d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x6f, 0x6c, 0x3b, 0x6c, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lol_lol_proto_rawDescOnce sync.Once
	file_lol_lol_proto_rawDescData = file_lol_lol_proto_rawDesc
)

func file_lol_lol_proto_rawDescGZIP() []byte {
	file_lol_lol_proto_rawDescOnce.Do(func() {
		file_lol_lol_proto_rawDescData = protoimpl.X.CompressGZIP(file_lol_lol_proto_rawDescData)
	})
	return file_lol_lol_proto_rawDescData
}

var file_lol_lol_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_lol_lol_proto_goTypes = []interface{}{
	(*GetChampionRequest)(nil),     // 0: lol.GetChampionRequest
	(*GetChampionResponse)(nil),    // 1: lol.GetChampionResponse
	(*ListChampionsRequest)(nil),   // 2: lol.ListChampionsRequest
	(*ListChampionsResponse)(nil),  // 3: lol.ListChampionsResponse
	(*GetBattleFieldRequest)(nil),  // 4: lol.GetBattleFieldRequest
	(*GetBattleFieldResponse)(nil), // 5: lol.GetBattleFieldResponse
	(*Champion)(nil),               // 6: champion.Champion
	(*BattleField)(nil),            // 7: battle_field.BattleField
}
var file_lol_lol_proto_depIdxs = []int32{
	6, // 0: lol.GetChampionResponse.champion:type_name -> champion.Champion
	6, // 1: lol.ListChampionsResponse.champions:type_name -> champion.Champion
	7, // 2: lol.GetBattleFieldResponse.battle_field:type_name -> battle_field.BattleField
	0, // 3: lol.LeagueOfLegends.GetChampion:input_type -> lol.GetChampionRequest
	2, // 4: lol.LeagueOfLegends.ListChampions:input_type -> lol.ListChampionsRequest
	4, // 5: lol.LeagueOfLegends.GetBattleField:input_type -> lol.GetBattleFieldRequest
	1, // 6: lol.LeagueOfLegends.GetChampion:output_type -> lol.GetChampionResponse
	3, // 7: lol.LeagueOfLegends.ListChampions:output_type -> lol.ListChampionsResponse
	5, // 8: lol.LeagueOfLegends.GetBattleField:output_type -> lol.GetBattleFieldResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_lol_lol_proto_init() }
func file_lol_lol_proto_init() {
	if File_lol_lol_proto != nil {
		return
	}
	file_lol_battle_field_proto_init()
	file_lol_champion_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_lol_lol_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChampionRequest); i {
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
		file_lol_lol_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChampionResponse); i {
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
		file_lol_lol_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListChampionsRequest); i {
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
		file_lol_lol_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListChampionsResponse); i {
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
		file_lol_lol_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBattleFieldRequest); i {
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
		file_lol_lol_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBattleFieldResponse); i {
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
			RawDescriptor: file_lol_lol_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_lol_lol_proto_goTypes,
		DependencyIndexes: file_lol_lol_proto_depIdxs,
		MessageInfos:      file_lol_lol_proto_msgTypes,
	}.Build()
	File_lol_lol_proto = out.File
	file_lol_lol_proto_rawDesc = nil
	file_lol_lol_proto_goTypes = nil
	file_lol_lol_proto_depIdxs = nil
}
