// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: maximum/maximum.proto

package maximumpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Number struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *Number) Reset() {
	*x = Number{}
	if protoimpl.UnsafeEnabled {
		mi := &file_maximum_maximum_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Number) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Number) ProtoMessage() {}

func (x *Number) ProtoReflect() protoreflect.Message {
	mi := &file_maximum_maximum_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Number.ProtoReflect.Descriptor instead.
func (*Number) Descriptor() ([]byte, []int) {
	return file_maximum_maximum_proto_rawDescGZIP(), []int{0}
}

func (x *Number) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type Maximum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Maximum int32 `protobuf:"varint,1,opt,name=maximum,proto3" json:"maximum,omitempty"`
}

func (x *Maximum) Reset() {
	*x = Maximum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_maximum_maximum_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Maximum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Maximum) ProtoMessage() {}

func (x *Maximum) ProtoReflect() protoreflect.Message {
	mi := &file_maximum_maximum_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Maximum.ProtoReflect.Descriptor instead.
func (*Maximum) Descriptor() ([]byte, []int) {
	return file_maximum_maximum_proto_rawDescGZIP(), []int{1}
}

func (x *Maximum) GetMaximum() int32 {
	if x != nil {
		return x.Maximum
	}
	return 0
}

var File_maximum_maximum_proto protoreflect.FileDescriptor

var file_maximum_maximum_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x2f, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x06, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x23, 0x0a, 0x07, 0x4d, 0x61, 0x78,
	0x69, 0x6d, 0x75, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x32, 0x37,
	0x0a, 0x0e, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x25, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x12, 0x07,
	0x2e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x1a, 0x08, 0x2e, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75,
	0x6d, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0c, 0x5a, 0x0a, 0x2f, 0x6d, 0x61, 0x78, 0x69,
	0x6d, 0x75, 0x6d, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_maximum_maximum_proto_rawDescOnce sync.Once
	file_maximum_maximum_proto_rawDescData = file_maximum_maximum_proto_rawDesc
)

func file_maximum_maximum_proto_rawDescGZIP() []byte {
	file_maximum_maximum_proto_rawDescOnce.Do(func() {
		file_maximum_maximum_proto_rawDescData = protoimpl.X.CompressGZIP(file_maximum_maximum_proto_rawDescData)
	})
	return file_maximum_maximum_proto_rawDescData
}

var file_maximum_maximum_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_maximum_maximum_proto_goTypes = []interface{}{
	(*Number)(nil),  // 0: Number
	(*Maximum)(nil), // 1: Maximum
}
var file_maximum_maximum_proto_depIdxs = []int32{
	0, // 0: MaximumService.GetMaximum:input_type -> Number
	1, // 1: MaximumService.GetMaximum:output_type -> Maximum
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_maximum_maximum_proto_init() }
func file_maximum_maximum_proto_init() {
	if File_maximum_maximum_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_maximum_maximum_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Number); i {
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
		file_maximum_maximum_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Maximum); i {
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
			RawDescriptor: file_maximum_maximum_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_maximum_maximum_proto_goTypes,
		DependencyIndexes: file_maximum_maximum_proto_depIdxs,
		MessageInfos:      file_maximum_maximum_proto_msgTypes,
	}.Build()
	File_maximum_maximum_proto = out.File
	file_maximum_maximum_proto_rawDesc = nil
	file_maximum_maximum_proto_goTypes = nil
	file_maximum_maximum_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MaximumServiceClient is the client API for MaximumService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MaximumServiceClient interface {
	GetMaximum(ctx context.Context, opts ...grpc.CallOption) (MaximumService_GetMaximumClient, error)
}

type maximumServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMaximumServiceClient(cc grpc.ClientConnInterface) MaximumServiceClient {
	return &maximumServiceClient{cc}
}

func (c *maximumServiceClient) GetMaximum(ctx context.Context, opts ...grpc.CallOption) (MaximumService_GetMaximumClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MaximumService_serviceDesc.Streams[0], "/MaximumService/GetMaximum", opts...)
	if err != nil {
		return nil, err
	}
	x := &maximumServiceGetMaximumClient{stream}
	return x, nil
}

type MaximumService_GetMaximumClient interface {
	Send(*Number) error
	Recv() (*Maximum, error)
	grpc.ClientStream
}

type maximumServiceGetMaximumClient struct {
	grpc.ClientStream
}

func (x *maximumServiceGetMaximumClient) Send(m *Number) error {
	return x.ClientStream.SendMsg(m)
}

func (x *maximumServiceGetMaximumClient) Recv() (*Maximum, error) {
	m := new(Maximum)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MaximumServiceServer is the server API for MaximumService service.
type MaximumServiceServer interface {
	GetMaximum(MaximumService_GetMaximumServer) error
}

// UnimplementedMaximumServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMaximumServiceServer struct {
}

func (*UnimplementedMaximumServiceServer) GetMaximum(MaximumService_GetMaximumServer) error {
	return status.Errorf(codes.Unimplemented, "method GetMaximum not implemented")
}

func RegisterMaximumServiceServer(s *grpc.Server, srv MaximumServiceServer) {
	s.RegisterService(&_MaximumService_serviceDesc, srv)
}

func _MaximumService_GetMaximum_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MaximumServiceServer).GetMaximum(&maximumServiceGetMaximumServer{stream})
}

type MaximumService_GetMaximumServer interface {
	Send(*Maximum) error
	Recv() (*Number, error)
	grpc.ServerStream
}

type maximumServiceGetMaximumServer struct {
	grpc.ServerStream
}

func (x *maximumServiceGetMaximumServer) Send(m *Maximum) error {
	return x.ServerStream.SendMsg(m)
}

func (x *maximumServiceGetMaximumServer) Recv() (*Number, error) {
	m := new(Number)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _MaximumService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "MaximumService",
	HandlerType: (*MaximumServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetMaximum",
			Handler:       _MaximumService_GetMaximum_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "maximum/maximum.proto",
}
