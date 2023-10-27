// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: pb/agent_service.proto

package pb

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

type Status int32

const (
	Status_OK  Status = 0
	Status_ERR Status = 1
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "OK",
		1: "ERR",
	}
	Status_value = map[string]int32{
		"OK":  0,
		"ERR": 1,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_agent_service_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_pb_agent_service_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_pb_agent_service_proto_rawDescGZIP(), []int{0}
}

// Update target service params
type AgentServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Component string            `protobuf:"bytes,1,opt,name=component,proto3" json:"component,omitempty"`
	HopstIP   string            `protobuf:"bytes,2,opt,name=hopstIP,proto3" json:"hopstIP,omitempty"`
	ParamMap  map[string]string `protobuf:"bytes,3,rep,name=paramMap,proto3" json:"paramMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Timestamp string            `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *AgentServiceRequest) Reset() {
	*x = AgentServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_agent_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentServiceRequest) ProtoMessage() {}

func (x *AgentServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_agent_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentServiceRequest.ProtoReflect.Descriptor instead.
func (*AgentServiceRequest) Descriptor() ([]byte, []int) {
	return file_pb_agent_service_proto_rawDescGZIP(), []int{0}
}

func (x *AgentServiceRequest) GetComponent() string {
	if x != nil {
		return x.Component
	}
	return ""
}

func (x *AgentServiceRequest) GetHopstIP() string {
	if x != nil {
		return x.HopstIP
	}
	return ""
}

func (x *AgentServiceRequest) GetParamMap() map[string]string {
	if x != nil {
		return x.ParamMap
	}
	return nil
}

func (x *AgentServiceRequest) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

// Upload files with big size
type AgentFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Component string `protobuf:"bytes,1,opt,name=component,proto3" json:"component,omitempty"`
	HopstIP   string `protobuf:"bytes,2,opt,name=hopstIP,proto3" json:"hopstIP,omitempty"`
	Content   []byte `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Timestamp string `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *AgentFileRequest) Reset() {
	*x = AgentFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_agent_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentFileRequest) ProtoMessage() {}

func (x *AgentFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_agent_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentFileRequest.ProtoReflect.Descriptor instead.
func (*AgentFileRequest) Descriptor() ([]byte, []int) {
	return file_pb_agent_service_proto_rawDescGZIP(), []int{1}
}

func (x *AgentFileRequest) GetComponent() string {
	if x != nil {
		return x.Component
	}
	return ""
}

func (x *AgentFileRequest) GetHopstIP() string {
	if x != nil {
		return x.HopstIP
	}
	return ""
}

func (x *AgentFileRequest) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *AgentFileRequest) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

// response from server
type AgentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    Status `protobuf:"varint,1,opt,name=status,proto3,enum=pb.Status" json:"status,omitempty"`
	Data      string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Timestamp string `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *AgentResponse) Reset() {
	*x = AgentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_agent_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentResponse) ProtoMessage() {}

func (x *AgentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_agent_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentResponse.ProtoReflect.Descriptor instead.
func (*AgentResponse) Descriptor() ([]byte, []int) {
	return file_pb_agent_service_proto_rawDescGZIP(), []int{2}
}

func (x *AgentResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_OK
}

func (x *AgentResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *AgentResponse) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

var File_pb_agent_service_proto protoreflect.FileDescriptor

var file_pb_agent_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x62, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xeb, 0x01, 0x0a,
	0x13, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x6f, 0x70, 0x73, 0x74, 0x49, 0x50, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x6f, 0x70, 0x73, 0x74, 0x49, 0x50, 0x12, 0x41, 0x0a, 0x08,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x4d, 0x61, 0x70, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x70, 0x62, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x4d, 0x61, 0x70,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x4d, 0x61, 0x70, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x1a, 0x3b, 0x0a,
	0x0d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x82, 0x01, 0x0a, 0x10, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x68, 0x6f, 0x70, 0x73, 0x74, 0x49, 0x50, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x68, 0x6f, 0x70, 0x73, 0x74, 0x49, 0x50, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22,
	0x65, 0x0a, 0x0d, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x22, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2a, 0x19, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x45, 0x52, 0x52, 0x10,
	0x01, 0x32, 0x8a, 0x01, 0x0a, 0x0b, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x41, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x75, 0x65, 0x6e,
	0x74, 0x62, 0x69, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x69,
	0x67, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x62,
	0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06,
	0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_agent_service_proto_rawDescOnce sync.Once
	file_pb_agent_service_proto_rawDescData = file_pb_agent_service_proto_rawDesc
)

func file_pb_agent_service_proto_rawDescGZIP() []byte {
	file_pb_agent_service_proto_rawDescOnce.Do(func() {
		file_pb_agent_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_agent_service_proto_rawDescData)
	})
	return file_pb_agent_service_proto_rawDescData
}

var file_pb_agent_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pb_agent_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pb_agent_service_proto_goTypes = []interface{}{
	(Status)(0),                 // 0: pb.Status
	(*AgentServiceRequest)(nil), // 1: pb.AgentServiceRequest
	(*AgentFileRequest)(nil),    // 2: pb.AgentFileRequest
	(*AgentResponse)(nil),       // 3: pb.AgentResponse
	nil,                         // 4: pb.AgentServiceRequest.ParamMapEntry
}
var file_pb_agent_service_proto_depIdxs = []int32{
	4, // 0: pb.AgentServiceRequest.paramMap:type_name -> pb.AgentServiceRequest.ParamMapEntry
	0, // 1: pb.AgentResponse.status:type_name -> pb.Status
	1, // 2: pb.AgentAction.UpdateFluentbitHost:input_type -> pb.AgentServiceRequest
	2, // 3: pb.AgentAction.UploadBigFile:input_type -> pb.AgentFileRequest
	3, // 4: pb.AgentAction.UpdateFluentbitHost:output_type -> pb.AgentResponse
	3, // 5: pb.AgentAction.UploadBigFile:output_type -> pb.AgentResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pb_agent_service_proto_init() }
func file_pb_agent_service_proto_init() {
	if File_pb_agent_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_agent_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentServiceRequest); i {
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
		file_pb_agent_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentFileRequest); i {
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
		file_pb_agent_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentResponse); i {
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
			RawDescriptor: file_pb_agent_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_agent_service_proto_goTypes,
		DependencyIndexes: file_pb_agent_service_proto_depIdxs,
		EnumInfos:         file_pb_agent_service_proto_enumTypes,
		MessageInfos:      file_pb_agent_service_proto_msgTypes,
	}.Build()
	File_pb_agent_service_proto = out.File
	file_pb_agent_service_proto_rawDesc = nil
	file_pb_agent_service_proto_goTypes = nil
	file_pb_agent_service_proto_depIdxs = nil
}
