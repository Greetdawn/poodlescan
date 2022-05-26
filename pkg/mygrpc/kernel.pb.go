// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: kernel.proto

package mygrpc

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

// 发送控制码的包
type SendOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TerminalParamJson string `protobuf:"bytes,1,opt,name=TerminalParamJson,proto3" json:"TerminalParamJson,omitempty"`
}

func (x *SendOrderRequest) Reset() {
	*x = SendOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kernel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendOrderRequest) ProtoMessage() {}

func (x *SendOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kernel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendOrderRequest.ProtoReflect.Descriptor instead.
func (*SendOrderRequest) Descriptor() ([]byte, []int) {
	return file_kernel_proto_rawDescGZIP(), []int{0}
}

func (x *SendOrderRequest) GetTerminalParamJson() string {
	if x != nil {
		return x.TerminalParamJson
	}
	return ""
}

type SendOrderReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentPkgID   uint32 `protobuf:"varint,1,opt,name=CurrentPkgID,proto3" json:"CurrentPkgID,omitempty"`
	TotalPkgNumber uint32 `protobuf:"varint,2,opt,name=TotalPkgNumber,proto3" json:"TotalPkgNumber,omitempty"`
	Type           uint32 `protobuf:"varint,3,opt,name=Type,proto3" json:"Type,omitempty"`
	Info           string `protobuf:"bytes,4,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *SendOrderReply) Reset() {
	*x = SendOrderReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kernel_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendOrderReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendOrderReply) ProtoMessage() {}

func (x *SendOrderReply) ProtoReflect() protoreflect.Message {
	mi := &file_kernel_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendOrderReply.ProtoReflect.Descriptor instead.
func (*SendOrderReply) Descriptor() ([]byte, []int) {
	return file_kernel_proto_rawDescGZIP(), []int{1}
}

func (x *SendOrderReply) GetCurrentPkgID() uint32 {
	if x != nil {
		return x.CurrentPkgID
	}
	return 0
}

func (x *SendOrderReply) GetTotalPkgNumber() uint32 {
	if x != nil {
		return x.TotalPkgNumber
	}
	return 0
}

func (x *SendOrderReply) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *SendOrderReply) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

var File_kernel_proto protoreflect.FileDescriptor

var file_kernel_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x6d, 0x79, 0x67, 0x72, 0x70, 0x63, 0x22, 0x40, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x54, 0x65,
	0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x4a, 0x73, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x4a, 0x73, 0x6f, 0x6e, 0x22, 0x84, 0x01, 0x0a, 0x0e, 0x53, 0x65, 0x6e,
	0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x6b, 0x67, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x6b, 0x67, 0x49, 0x44, 0x12,
	0x26, 0x0a, 0x0e, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x6b, 0x67, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x6b,
	0x67, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x32,
	0x4b, 0x0a, 0x06, 0x4b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x12, 0x41, 0x0a, 0x09, 0x53, 0x65, 0x6e,
	0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x6d, 0x79, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x53, 0x65, 0x6e, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x6d, 0x79, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x30, 0x01, 0x42, 0x0b, 0x5a, 0x09,
	0x2e, 0x2e, 0x2f, 0x6d, 0x79, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_kernel_proto_rawDescOnce sync.Once
	file_kernel_proto_rawDescData = file_kernel_proto_rawDesc
)

func file_kernel_proto_rawDescGZIP() []byte {
	file_kernel_proto_rawDescOnce.Do(func() {
		file_kernel_proto_rawDescData = protoimpl.X.CompressGZIP(file_kernel_proto_rawDescData)
	})
	return file_kernel_proto_rawDescData
}

var file_kernel_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_kernel_proto_goTypes = []interface{}{
	(*SendOrderRequest)(nil), // 0: mygrpc.SendOrderRequest
	(*SendOrderReply)(nil),   // 1: mygrpc.SendOrderReply
}
var file_kernel_proto_depIdxs = []int32{
	0, // 0: mygrpc.Kernel.SendOrder:input_type -> mygrpc.SendOrderRequest
	1, // 1: mygrpc.Kernel.SendOrder:output_type -> mygrpc.SendOrderReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_kernel_proto_init() }
func file_kernel_proto_init() {
	if File_kernel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kernel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendOrderRequest); i {
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
		file_kernel_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendOrderReply); i {
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
			RawDescriptor: file_kernel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kernel_proto_goTypes,
		DependencyIndexes: file_kernel_proto_depIdxs,
		MessageInfos:      file_kernel_proto_msgTypes,
	}.Build()
	File_kernel_proto = out.File
	file_kernel_proto_rawDesc = nil
	file_kernel_proto_goTypes = nil
	file_kernel_proto_depIdxs = nil
}
