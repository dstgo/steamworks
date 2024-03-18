// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: gc.proto

package message

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

type GCProtoBufMsgSrc int32

const (
	GCProtoBufMsgSrc_GCProtoBufMsgSrc_Unspecified GCProtoBufMsgSrc = 0
	GCProtoBufMsgSrc_GCProtoBufMsgSrc_FromSystem  GCProtoBufMsgSrc = 1
	GCProtoBufMsgSrc_GCProtoBufMsgSrc_FromSteamID GCProtoBufMsgSrc = 2
	GCProtoBufMsgSrc_GCProtoBufMsgSrc_FromGC      GCProtoBufMsgSrc = 3
	GCProtoBufMsgSrc_GCProtoBufMsgSrc_ReplySystem GCProtoBufMsgSrc = 4
)

// Enum value maps for GCProtoBufMsgSrc.
var (
	GCProtoBufMsgSrc_name = map[int32]string{
		0: "GCProtoBufMsgSrc_Unspecified",
		1: "GCProtoBufMsgSrc_FromSystem",
		2: "GCProtoBufMsgSrc_FromSteamID",
		3: "GCProtoBufMsgSrc_FromGC",
		4: "GCProtoBufMsgSrc_ReplySystem",
	}
	GCProtoBufMsgSrc_value = map[string]int32{
		"GCProtoBufMsgSrc_Unspecified": 0,
		"GCProtoBufMsgSrc_FromSystem":  1,
		"GCProtoBufMsgSrc_FromSteamID": 2,
		"GCProtoBufMsgSrc_FromGC":      3,
		"GCProtoBufMsgSrc_ReplySystem": 4,
	}
)

func (x GCProtoBufMsgSrc) Enum() *GCProtoBufMsgSrc {
	p := new(GCProtoBufMsgSrc)
	*p = x
	return p
}

func (x GCProtoBufMsgSrc) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GCProtoBufMsgSrc) Descriptor() protoreflect.EnumDescriptor {
	return file_gc_proto_enumTypes[0].Descriptor()
}

func (GCProtoBufMsgSrc) Type() protoreflect.EnumType {
	return &file_gc_proto_enumTypes[0]
}

func (x GCProtoBufMsgSrc) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *GCProtoBufMsgSrc) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = GCProtoBufMsgSrc(num)
	return nil
}

// Deprecated: Use GCProtoBufMsgSrc.Descriptor instead.
func (GCProtoBufMsgSrc) EnumDescriptor() ([]byte, []int) {
	return file_gc_proto_rawDescGZIP(), []int{0}
}

type CMsgProtoBufHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientSteamId    *uint64           `protobuf:"fixed64,1,opt,name=client_steam_id,json=clientSteamId" json:"client_steam_id,omitempty"`
	ClientSessionId  *int32            `protobuf:"varint,2,opt,name=client_session_id,json=clientSessionId" json:"client_session_id,omitempty"`
	SourceAppId      *uint32           `protobuf:"varint,3,opt,name=source_app_id,json=sourceAppId" json:"source_app_id,omitempty"`
	JobIdSource      *uint64           `protobuf:"fixed64,10,opt,name=job_id_source,json=jobIdSource,def=18446744073709551615" json:"job_id_source,omitempty"`
	JobIdTarget      *uint64           `protobuf:"fixed64,11,opt,name=job_id_target,json=jobIdTarget,def=18446744073709551615" json:"job_id_target,omitempty"`
	TargetJobName    *string           `protobuf:"bytes,12,opt,name=target_job_name,json=targetJobName" json:"target_job_name,omitempty"`
	Eresult          *int32            `protobuf:"varint,13,opt,name=eresult,def=2" json:"eresult,omitempty"`
	ErrorMessage     *string           `protobuf:"bytes,14,opt,name=error_message,json=errorMessage" json:"error_message,omitempty"`
	GcMsgSrc         *GCProtoBufMsgSrc `protobuf:"varint,200,opt,name=gc_msg_src,json=gcMsgSrc,enum=foobar.GCProtoBufMsgSrc,def=0" json:"gc_msg_src,omitempty"`
	GcDirIndexSource *uint32           `protobuf:"varint,201,opt,name=gc_dir_index_source,json=gcDirIndexSource" json:"gc_dir_index_source,omitempty"`
}

// Default values for CMsgProtoBufHeader fields.
const (
	Default_CMsgProtoBufHeader_JobIdSource = uint64(18446744073709551615)
	Default_CMsgProtoBufHeader_JobIdTarget = uint64(18446744073709551615)
	Default_CMsgProtoBufHeader_Eresult     = int32(2)
	Default_CMsgProtoBufHeader_GcMsgSrc    = GCProtoBufMsgSrc_GCProtoBufMsgSrc_Unspecified
)

func (x *CMsgProtoBufHeader) Reset() {
	*x = CMsgProtoBufHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgProtoBufHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgProtoBufHeader) ProtoMessage() {}

func (x *CMsgProtoBufHeader) ProtoReflect() protoreflect.Message {
	mi := &file_gc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgProtoBufHeader.ProtoReflect.Descriptor instead.
func (*CMsgProtoBufHeader) Descriptor() ([]byte, []int) {
	return file_gc_proto_rawDescGZIP(), []int{0}
}

func (x *CMsgProtoBufHeader) GetClientSteamId() uint64 {
	if x != nil && x.ClientSteamId != nil {
		return *x.ClientSteamId
	}
	return 0
}

func (x *CMsgProtoBufHeader) GetClientSessionId() int32 {
	if x != nil && x.ClientSessionId != nil {
		return *x.ClientSessionId
	}
	return 0
}

func (x *CMsgProtoBufHeader) GetSourceAppId() uint32 {
	if x != nil && x.SourceAppId != nil {
		return *x.SourceAppId
	}
	return 0
}

func (x *CMsgProtoBufHeader) GetJobIdSource() uint64 {
	if x != nil && x.JobIdSource != nil {
		return *x.JobIdSource
	}
	return Default_CMsgProtoBufHeader_JobIdSource
}

func (x *CMsgProtoBufHeader) GetJobIdTarget() uint64 {
	if x != nil && x.JobIdTarget != nil {
		return *x.JobIdTarget
	}
	return Default_CMsgProtoBufHeader_JobIdTarget
}

func (x *CMsgProtoBufHeader) GetTargetJobName() string {
	if x != nil && x.TargetJobName != nil {
		return *x.TargetJobName
	}
	return ""
}

func (x *CMsgProtoBufHeader) GetEresult() int32 {
	if x != nil && x.Eresult != nil {
		return *x.Eresult
	}
	return Default_CMsgProtoBufHeader_Eresult
}

func (x *CMsgProtoBufHeader) GetErrorMessage() string {
	if x != nil && x.ErrorMessage != nil {
		return *x.ErrorMessage
	}
	return ""
}

func (x *CMsgProtoBufHeader) GetGcMsgSrc() GCProtoBufMsgSrc {
	if x != nil && x.GcMsgSrc != nil {
		return *x.GcMsgSrc
	}
	return Default_CMsgProtoBufHeader_GcMsgSrc
}

func (x *CMsgProtoBufHeader) GetGcDirIndexSource() uint32 {
	if x != nil && x.GcDirIndexSource != nil {
		return *x.GcDirIndexSource
	}
	return 0
}

var File_gc_proto protoreflect.FileDescriptor

var file_gc_proto_rawDesc = []byte{
	0x0a, 0x08, 0x67, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x66, 0x6f, 0x6f, 0x62,
	0x61, 0x72, 0x22, 0xf1, 0x03, 0x0a, 0x12, 0x43, 0x4d, 0x73, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x42, 0x75, 0x66, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x06, 0x52, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x49,
	0x64, 0x12, 0x2a, 0x0a, 0x11, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x22, 0x0a,
	0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x70, 0x70, 0x49,
	0x64, 0x12, 0x38, 0x0a, 0x0d, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x5f, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x06, 0x3a, 0x14, 0x31, 0x38, 0x34, 0x34, 0x36, 0x37,
	0x34, 0x34, 0x30, 0x37, 0x33, 0x37, 0x30, 0x39, 0x35, 0x35, 0x31, 0x36, 0x31, 0x35, 0x52, 0x0b,
	0x6a, 0x6f, 0x62, 0x49, 0x64, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x0d, 0x6a,
	0x6f, 0x62, 0x5f, 0x69, 0x64, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x06, 0x3a, 0x14, 0x31, 0x38, 0x34, 0x34, 0x36, 0x37, 0x34, 0x34, 0x30, 0x37, 0x33, 0x37,
	0x30, 0x39, 0x35, 0x35, 0x31, 0x36, 0x31, 0x35, 0x52, 0x0b, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f,
	0x6a, 0x6f, 0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x07, 0x65, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x01,
	0x32, 0x52, 0x07, 0x65, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x55, 0x0a, 0x0a, 0x67, 0x63, 0x5f, 0x6d, 0x73, 0x67, 0x5f, 0x73, 0x72, 0x63, 0x18, 0xc8, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x66, 0x6f, 0x6f, 0x62, 0x61, 0x72, 0x2e, 0x47, 0x43,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x75, 0x66, 0x4d, 0x73, 0x67, 0x53, 0x72, 0x63, 0x3a, 0x1c,
	0x47, 0x43, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x75, 0x66, 0x4d, 0x73, 0x67, 0x53, 0x72, 0x63,
	0x5f, 0x55, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x52, 0x08, 0x67, 0x63,
	0x4d, 0x73, 0x67, 0x53, 0x72, 0x63, 0x12, 0x2e, 0x0a, 0x13, 0x67, 0x63, 0x5f, 0x64, 0x69, 0x72,
	0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0xc9, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x67, 0x63, 0x44, 0x69, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2a, 0xb6, 0x01, 0x0a, 0x10, 0x47, 0x43, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x42, 0x75, 0x66, 0x4d, 0x73, 0x67, 0x53, 0x72, 0x63, 0x12, 0x20, 0x0a, 0x1c, 0x47,
	0x43, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x75, 0x66, 0x4d, 0x73, 0x67, 0x53, 0x72, 0x63, 0x5f,
	0x55, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x10, 0x00, 0x12, 0x1f, 0x0a,
	0x1b, 0x47, 0x43, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x75, 0x66, 0x4d, 0x73, 0x67, 0x53, 0x72,
	0x63, 0x5f, 0x46, 0x72, 0x6f, 0x6d, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x10, 0x01, 0x12, 0x20,
	0x0a, 0x1c, 0x47, 0x43, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x75, 0x66, 0x4d, 0x73, 0x67, 0x53,
	0x72, 0x63, 0x5f, 0x46, 0x72, 0x6f, 0x6d, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x44, 0x10, 0x02,
	0x12, 0x1b, 0x0a, 0x17, 0x47, 0x43, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x75, 0x66, 0x4d, 0x73,
	0x67, 0x53, 0x72, 0x63, 0x5f, 0x46, 0x72, 0x6f, 0x6d, 0x47, 0x43, 0x10, 0x03, 0x12, 0x20, 0x0a,
	0x1c, 0x47, 0x43, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x75, 0x66, 0x4d, 0x73, 0x67, 0x53, 0x72,
	0x63, 0x5f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x10, 0x04, 0x42,
	0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x73,
	0x74, 0x67, 0x6f, 0x2f, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x3b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
}

var (
	file_gc_proto_rawDescOnce sync.Once
	file_gc_proto_rawDescData = file_gc_proto_rawDesc
)

func file_gc_proto_rawDescGZIP() []byte {
	file_gc_proto_rawDescOnce.Do(func() {
		file_gc_proto_rawDescData = protoimpl.X.CompressGZIP(file_gc_proto_rawDescData)
	})
	return file_gc_proto_rawDescData
}

var file_gc_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_gc_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_gc_proto_goTypes = []interface{}{
	(GCProtoBufMsgSrc)(0),      // 0: foobar.GCProtoBufMsgSrc
	(*CMsgProtoBufHeader)(nil), // 1: foobar.CMsgProtoBufHeader
}
var file_gc_proto_depIdxs = []int32{
	0, // 0: foobar.CMsgProtoBufHeader.gc_msg_src:type_name -> foobar.GCProtoBufMsgSrc
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_gc_proto_init() }
func file_gc_proto_init() {
	if File_gc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CMsgProtoBufHeader); i {
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
			RawDescriptor: file_gc_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gc_proto_goTypes,
		DependencyIndexes: file_gc_proto_depIdxs,
		EnumInfos:         file_gc_proto_enumTypes,
		MessageInfos:      file_gc_proto_msgTypes,
	}.Build()
	File_gc_proto = out.File
	file_gc_proto_rawDesc = nil
	file_gc_proto_goTypes = nil
	file_gc_proto_depIdxs = nil
}
