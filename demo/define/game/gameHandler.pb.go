// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v5.29.1
// source: gameHandler.proto

package game

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

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Num   int64 `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
	Level int32 `protobuf:"varint,3,opt,name=level,proto3" json:"level,omitempty"`
	Star  int32 `protobuf:"varint,4,opt,name=star,proto3" json:"star,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gameHandler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_gameHandler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_gameHandler_proto_rawDescGZIP(), []int{0}
}

func (x *Item) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Item) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *Item) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *Item) GetStar() int32 {
	if x != nil {
		return x.Star
	}
	return 0
}

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemList []*Item `protobuf:"bytes,1,rep,name=itemList,proto3" json:"itemList,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gameHandler_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_gameHandler_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_gameHandler_proto_rawDescGZIP(), []int{1}
}

func (x *Role) GetItemList() []*Item {
	if x != nil {
		return x.ItemList
	}
	return nil
}

type LoginDefine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A     string `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *LoginDefine) Reset() {
	*x = LoginDefine{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gameHandler_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginDefine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginDefine) ProtoMessage() {}

func (x *LoginDefine) ProtoReflect() protoreflect.Message {
	mi := &file_gameHandler_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginDefine.ProtoReflect.Descriptor instead.
func (*LoginDefine) Descriptor() ([]byte, []int) {
	return file_gameHandler_proto_rawDescGZIP(), []int{2}
}

func (x *LoginDefine) GetA() string {
	if x != nil {
		return x.A
	}
	return ""
}

func (x *LoginDefine) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_gameHandler_proto protoreflect.FileDescriptor

var file_gameHandler_proto_rawDesc = []byte{
	0x0a, 0x11, 0x67, 0x61, 0x6d, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e,
	0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x73, 0x74, 0x61, 0x72, 0x22, 0x29, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12,
	0x21, 0x0a, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x05, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4c, 0x69,
	0x73, 0x74, 0x22, 0x31, 0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x44, 0x65, 0x66, 0x69, 0x6e,
	0x65, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x61, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x07, 0x5a, 0x05, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gameHandler_proto_rawDescOnce sync.Once
	file_gameHandler_proto_rawDescData = file_gameHandler_proto_rawDesc
)

func file_gameHandler_proto_rawDescGZIP() []byte {
	file_gameHandler_proto_rawDescOnce.Do(func() {
		file_gameHandler_proto_rawDescData = protoimpl.X.CompressGZIP(file_gameHandler_proto_rawDescData)
	})
	return file_gameHandler_proto_rawDescData
}

var file_gameHandler_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_gameHandler_proto_goTypes = []interface{}{
	(*Item)(nil),        // 0: Item
	(*Role)(nil),        // 1: Role
	(*LoginDefine)(nil), // 2: LoginDefine
}
var file_gameHandler_proto_depIdxs = []int32{
	0, // 0: Role.itemList:type_name -> Item
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_gameHandler_proto_init() }
func file_gameHandler_proto_init() {
	if File_gameHandler_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gameHandler_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_gameHandler_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
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
		file_gameHandler_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginDefine); i {
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
			RawDescriptor: file_gameHandler_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gameHandler_proto_goTypes,
		DependencyIndexes: file_gameHandler_proto_depIdxs,
		MessageInfos:      file_gameHandler_proto_msgTypes,
	}.Build()
	File_gameHandler_proto = out.File
	file_gameHandler_proto_rawDesc = nil
	file_gameHandler_proto_goTypes = nil
	file_gameHandler_proto_depIdxs = nil
}
