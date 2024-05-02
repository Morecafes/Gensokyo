// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: c2cpic.proto

package protobufc2c

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

// 定义一个Proto消息结构来解析特定的嵌套数据
type ImageData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 实例化Level1消息
	NestedData *ImageData_Level1 `protobuf:"bytes,1,opt,name=nestedData,proto3" json:"nestedData,omitempty"`
}

func (x *ImageData) Reset() {
	*x = ImageData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_c2cpic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageData) ProtoMessage() {}

func (x *ImageData) ProtoReflect() protoreflect.Message {
	mi := &file_c2cpic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageData.ProtoReflect.Descriptor instead.
func (*ImageData) Descriptor() ([]byte, []int) {
	return file_c2cpic_proto_rawDescGZIP(), []int{0}
}

func (x *ImageData) GetNestedData() *ImageData_Level1 {
	if x != nil {
		return x.NestedData
	}
	return nil
}

// 内嵌消息定义为一级字段1的结构
type ImageData_Level1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 实例化Level3消息
	Details *ImageData_Level1_Level3 `protobuf:"bytes,3,opt,name=details,proto3" json:"details,omitempty"`
}

func (x *ImageData_Level1) Reset() {
	*x = ImageData_Level1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_c2cpic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageData_Level1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageData_Level1) ProtoMessage() {}

func (x *ImageData_Level1) ProtoReflect() protoreflect.Message {
	mi := &file_c2cpic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageData_Level1.ProtoReflect.Descriptor instead.
func (*ImageData_Level1) Descriptor() ([]byte, []int) {
	return file_c2cpic_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ImageData_Level1) GetDetails() *ImageData_Level1_Level3 {
	if x != nil {
		return x.Details
	}
	return nil
}

// 内嵌消息定义为二级字段3的结构
type ImageData_Level1_Level3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 可选字段，解析图片宽度
	Width int32 `protobuf:"varint,8,opt,name=width,proto3" json:"width,omitempty"`
	// 可选字段，解析图片高度
	Height int32 `protobuf:"varint,9,opt,name=height,proto3" json:"height,omitempty"`
	// 实例化Level29消息
	Urls *ImageData_Level1_Level3_Level29 `protobuf:"bytes,29,opt,name=urls,proto3" json:"urls,omitempty"`
}

func (x *ImageData_Level1_Level3) Reset() {
	*x = ImageData_Level1_Level3{}
	if protoimpl.UnsafeEnabled {
		mi := &file_c2cpic_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageData_Level1_Level3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageData_Level1_Level3) ProtoMessage() {}

func (x *ImageData_Level1_Level3) ProtoReflect() protoreflect.Message {
	mi := &file_c2cpic_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageData_Level1_Level3.ProtoReflect.Descriptor instead.
func (*ImageData_Level1_Level3) Descriptor() ([]byte, []int) {
	return file_c2cpic_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *ImageData_Level1_Level3) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *ImageData_Level1_Level3) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *ImageData_Level1_Level3) GetUrls() *ImageData_Level1_Level3_Level29 {
	if x != nil {
		return x.Urls
	}
	return nil
}

// 内嵌消息定义为三级字段29的结构
type ImageData_Level1_Level3_Level29 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 可选字段，解析图片URL
	Url string `protobuf:"bytes,30,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *ImageData_Level1_Level3_Level29) Reset() {
	*x = ImageData_Level1_Level3_Level29{}
	if protoimpl.UnsafeEnabled {
		mi := &file_c2cpic_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageData_Level1_Level3_Level29) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageData_Level1_Level3_Level29) ProtoMessage() {}

func (x *ImageData_Level1_Level3_Level29) ProtoReflect() protoreflect.Message {
	mi := &file_c2cpic_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageData_Level1_Level3_Level29.ProtoReflect.Descriptor instead.
func (*ImageData_Level1_Level3_Level29) Descriptor() ([]byte, []int) {
	return file_c2cpic_proto_rawDescGZIP(), []int{0, 0, 0, 0}
}

func (x *ImageData_Level1_Level3_Level29) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_c2cpic_proto protoreflect.FileDescriptor

var file_c2cpic_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x32, 0x63, 0x70, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x63, 0x32, 0x63, 0x22, 0xad, 0x02, 0x0a, 0x09,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x3d, 0x0a, 0x0a, 0x6e, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x63, 0x32, 0x63, 0x2e, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x31, 0x52, 0x0a, 0x6e, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x1a, 0xe0, 0x01, 0x0a, 0x06, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x31, 0x12, 0x3e, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x63,
	0x32, 0x63, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x31, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x33, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x1a, 0x95, 0x01, 0x0a, 0x06, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x33, 0x12, 0x14,
	0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x77,
	0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x40, 0x0a, 0x04,
	0x75, 0x72, 0x6c, 0x73, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x63, 0x32, 0x63, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x31, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x33,
	0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x32, 0x39, 0x52, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x1a, 0x1b,
	0x0a, 0x07, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x32, 0x39, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x42, 0x30, 0x5a, 0x2e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x6f, 0x73, 0x68, 0x69, 0x6e,
	0x6f, 0x6e, 0x79, 0x61, 0x72, 0x75, 0x6b, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x73, 0x6f, 0x6b, 0x79,
	0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x63, 0x32, 0x63, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_c2cpic_proto_rawDescOnce sync.Once
	file_c2cpic_proto_rawDescData = file_c2cpic_proto_rawDesc
)

func file_c2cpic_proto_rawDescGZIP() []byte {
	file_c2cpic_proto_rawDescOnce.Do(func() {
		file_c2cpic_proto_rawDescData = protoimpl.X.CompressGZIP(file_c2cpic_proto_rawDescData)
	})
	return file_c2cpic_proto_rawDescData
}

var file_c2cpic_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_c2cpic_proto_goTypes = []interface{}{
	(*ImageData)(nil),                       // 0: protobufc2c.ImageData
	(*ImageData_Level1)(nil),                // 1: protobufc2c.ImageData.Level1
	(*ImageData_Level1_Level3)(nil),         // 2: protobufc2c.ImageData.Level1.Level3
	(*ImageData_Level1_Level3_Level29)(nil), // 3: protobufc2c.ImageData.Level1.Level3.Level29
}
var file_c2cpic_proto_depIdxs = []int32{
	1, // 0: protobufc2c.ImageData.nestedData:type_name -> protobufc2c.ImageData.Level1
	2, // 1: protobufc2c.ImageData.Level1.details:type_name -> protobufc2c.ImageData.Level1.Level3
	3, // 2: protobufc2c.ImageData.Level1.Level3.urls:type_name -> protobufc2c.ImageData.Level1.Level3.Level29
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_c2cpic_proto_init() }
func file_c2cpic_proto_init() {
	if File_c2cpic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_c2cpic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageData); i {
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
		file_c2cpic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageData_Level1); i {
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
		file_c2cpic_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageData_Level1_Level3); i {
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
		file_c2cpic_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageData_Level1_Level3_Level29); i {
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
			RawDescriptor: file_c2cpic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_c2cpic_proto_goTypes,
		DependencyIndexes: file_c2cpic_proto_depIdxs,
		MessageInfos:      file_c2cpic_proto_msgTypes,
	}.Build()
	File_c2cpic_proto = out.File
	file_c2cpic_proto_rawDesc = nil
	file_c2cpic_proto_goTypes = nil
	file_c2cpic_proto_depIdxs = nil
}
