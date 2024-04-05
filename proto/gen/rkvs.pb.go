// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: proto/rkvs.proto

package rkvs

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

type TransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operations []*Operation `protobuf:"bytes,1,rep,name=operations,proto3" json:"operations,omitempty"`
}

func (x *TransactionRequest) Reset() {
	*x = TransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rkvs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionRequest) ProtoMessage() {}

func (x *TransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rkvs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionRequest.ProtoReflect.Descriptor instead.
func (*TransactionRequest) Descriptor() ([]byte, []int) {
	return file_proto_rkvs_proto_rawDescGZIP(), []int{0}
}

func (x *TransactionRequest) GetOperations() []*Operation {
	if x != nil {
		return x.Operations
	}
	return nil
}

type Operation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to OperationType:
	//
	//	*Operation_Set
	//	*Operation_Delete
	OperationType isOperation_OperationType `protobuf_oneof:"operation_type"`
}

func (x *Operation) Reset() {
	*x = Operation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rkvs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Operation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operation) ProtoMessage() {}

func (x *Operation) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rkvs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operation.ProtoReflect.Descriptor instead.
func (*Operation) Descriptor() ([]byte, []int) {
	return file_proto_rkvs_proto_rawDescGZIP(), []int{1}
}

func (m *Operation) GetOperationType() isOperation_OperationType {
	if m != nil {
		return m.OperationType
	}
	return nil
}

func (x *Operation) GetSet() *SetOperation {
	if x, ok := x.GetOperationType().(*Operation_Set); ok {
		return x.Set
	}
	return nil
}

func (x *Operation) GetDelete() *DeleteOperation {
	if x, ok := x.GetOperationType().(*Operation_Delete); ok {
		return x.Delete
	}
	return nil
}

type isOperation_OperationType interface {
	isOperation_OperationType()
}

type Operation_Set struct {
	Set *SetOperation `protobuf:"bytes,1,opt,name=set,proto3,oneof"`
}

type Operation_Delete struct {
	Delete *DeleteOperation `protobuf:"bytes,2,opt,name=delete,proto3,oneof"`
}

func (*Operation_Set) isOperation_OperationType() {}

func (*Operation_Delete) isOperation_OperationType() {}

type SetOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SetOperation) Reset() {
	*x = SetOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rkvs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetOperation) ProtoMessage() {}

func (x *SetOperation) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rkvs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetOperation.ProtoReflect.Descriptor instead.
func (*SetOperation) Descriptor() ([]byte, []int) {
	return file_proto_rkvs_proto_rawDescGZIP(), []int{2}
}

func (x *SetOperation) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SetOperation) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type DeleteOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *DeleteOperation) Reset() {
	*x = DeleteOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rkvs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOperation) ProtoMessage() {}

func (x *DeleteOperation) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rkvs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteOperation.ProtoReflect.Descriptor instead.
func (*DeleteOperation) Descriptor() ([]byte, []int) {
	return file_proto_rkvs_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteOperation) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type Ack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *Ack) Reset() {
	*x = Ack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rkvs_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ack) ProtoMessage() {}

func (x *Ack) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rkvs_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ack.ProtoReflect.Descriptor instead.
func (*Ack) Descriptor() ([]byte, []int) {
	return file_proto_rkvs_proto_rawDescGZIP(), []int{4}
}

func (x *Ack) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type Key struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *Key) Reset() {
	*x = Key{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rkvs_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Key) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Key) ProtoMessage() {}

func (x *Key) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rkvs_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Key.ProtoReflect.Descriptor instead.
func (*Key) Descriptor() ([]byte, []int) {
	return file_proto_rkvs_proto_rawDescGZIP(), []int{5}
}

func (x *Key) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Value) Reset() {
	*x = Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rkvs_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value) ProtoMessage() {}

func (x *Value) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rkvs_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value.ProtoReflect.Descriptor instead.
func (*Value) Descriptor() ([]byte, []int) {
	return file_proto_rkvs_proto_rawDescGZIP(), []int{6}
}

func (x *Value) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_proto_rkvs_proto protoreflect.FileDescriptor

var file_proto_rkvs_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x6b, 0x76, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x72, 0x6b, 0x76, 0x73, 0x22, 0x45, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f,
	0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x6b, 0x76, 0x73, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x76, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x03,
	0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x72, 0x6b, 0x76, 0x73,
	0x2e, 0x53, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52,
	0x03, 0x73, 0x65, 0x74, 0x12, 0x2f, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x72, 0x6b, 0x76, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x06, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x10, 0x0a, 0x0e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x36, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x23, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x22, 0x1f, 0x0a, 0x03, 0x41, 0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x17, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x1d,
	0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x64, 0x0a,
	0x04, 0x52, 0x6b, 0x76, 0x73, 0x12, 0x1f, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x09, 0x2e, 0x72,
	0x6b, 0x76, 0x73, 0x2e, 0x4b, 0x65, 0x79, 0x1a, 0x0b, 0x2e, 0x72, 0x6b, 0x76, 0x73, 0x2e, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x12, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x2e, 0x72,
	0x6b, 0x76, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x72, 0x6b, 0x76, 0x73, 0x2e, 0x41, 0x63,
	0x6b, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x65, 0x6e, 0x3b, 0x72, 0x6b, 0x76, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_rkvs_proto_rawDescOnce sync.Once
	file_proto_rkvs_proto_rawDescData = file_proto_rkvs_proto_rawDesc
)

func file_proto_rkvs_proto_rawDescGZIP() []byte {
	file_proto_rkvs_proto_rawDescOnce.Do(func() {
		file_proto_rkvs_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_rkvs_proto_rawDescData)
	})
	return file_proto_rkvs_proto_rawDescData
}

var file_proto_rkvs_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_rkvs_proto_goTypes = []interface{}{
	(*TransactionRequest)(nil), // 0: rkvs.TransactionRequest
	(*Operation)(nil),          // 1: rkvs.Operation
	(*SetOperation)(nil),       // 2: rkvs.SetOperation
	(*DeleteOperation)(nil),    // 3: rkvs.DeleteOperation
	(*Ack)(nil),                // 4: rkvs.Ack
	(*Key)(nil),                // 5: rkvs.Key
	(*Value)(nil),              // 6: rkvs.Value
}
var file_proto_rkvs_proto_depIdxs = []int32{
	1, // 0: rkvs.TransactionRequest.operations:type_name -> rkvs.Operation
	2, // 1: rkvs.Operation.set:type_name -> rkvs.SetOperation
	3, // 2: rkvs.Operation.delete:type_name -> rkvs.DeleteOperation
	5, // 3: rkvs.Rkvs.Get:input_type -> rkvs.Key
	0, // 4: rkvs.Rkvs.ExecuteTransaction:input_type -> rkvs.TransactionRequest
	6, // 5: rkvs.Rkvs.Get:output_type -> rkvs.Value
	4, // 6: rkvs.Rkvs.ExecuteTransaction:output_type -> rkvs.Ack
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_rkvs_proto_init() }
func file_proto_rkvs_proto_init() {
	if File_proto_rkvs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_rkvs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionRequest); i {
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
		file_proto_rkvs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Operation); i {
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
		file_proto_rkvs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetOperation); i {
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
		file_proto_rkvs_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteOperation); i {
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
		file_proto_rkvs_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ack); i {
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
		file_proto_rkvs_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Key); i {
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
		file_proto_rkvs_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Value); i {
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
	file_proto_rkvs_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Operation_Set)(nil),
		(*Operation_Delete)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_rkvs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_rkvs_proto_goTypes,
		DependencyIndexes: file_proto_rkvs_proto_depIdxs,
		MessageInfos:      file_proto_rkvs_proto_msgTypes,
	}.Build()
	File_proto_rkvs_proto = out.File
	file_proto_rkvs_proto_rawDesc = nil
	file_proto_rkvs_proto_goTypes = nil
	file_proto_rkvs_proto_depIdxs = nil
}