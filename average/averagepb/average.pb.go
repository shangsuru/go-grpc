// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: average/average.proto

package averagepb

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

	N int32 `protobuf:"varint,1,opt,name=n,proto3" json:"n,omitempty"`
}

func (x *Number) Reset() {
	*x = Number{}
	if protoimpl.UnsafeEnabled {
		mi := &file_average_average_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Number) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Number) ProtoMessage() {}

func (x *Number) ProtoReflect() protoreflect.Message {
	mi := &file_average_average_proto_msgTypes[0]
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
	return file_average_average_proto_rawDescGZIP(), []int{0}
}

func (x *Number) GetN() int32 {
	if x != nil {
		return x.N
	}
	return 0
}

type Average struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *Average) Reset() {
	*x = Average{}
	if protoimpl.UnsafeEnabled {
		mi := &file_average_average_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Average) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Average) ProtoMessage() {}

func (x *Average) ProtoReflect() protoreflect.Message {
	mi := &file_average_average_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Average.ProtoReflect.Descriptor instead.
func (*Average) Descriptor() ([]byte, []int) {
	return file_average_average_proto_rawDescGZIP(), []int{1}
}

func (x *Average) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_average_average_proto protoreflect.FileDescriptor

var file_average_average_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x16, 0x0a, 0x06, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x0c, 0x0a, 0x01, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x6e, 0x22,
	0x21, 0x0a, 0x07, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x32, 0x39, 0x0a, 0x0e, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x41,
	0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x07, 0x2e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x1a,
	0x08, 0x2e, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x22, 0x00, 0x28, 0x01, 0x42, 0x0c, 0x5a,
	0x0a, 0x2f, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_average_average_proto_rawDescOnce sync.Once
	file_average_average_proto_rawDescData = file_average_average_proto_rawDesc
)

func file_average_average_proto_rawDescGZIP() []byte {
	file_average_average_proto_rawDescOnce.Do(func() {
		file_average_average_proto_rawDescData = protoimpl.X.CompressGZIP(file_average_average_proto_rawDescData)
	})
	return file_average_average_proto_rawDescData
}

var file_average_average_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_average_average_proto_goTypes = []interface{}{
	(*Number)(nil),  // 0: Number
	(*Average)(nil), // 1: Average
}
var file_average_average_proto_depIdxs = []int32{
	0, // 0: AverageService.ComputeAverage:input_type -> Number
	1, // 1: AverageService.ComputeAverage:output_type -> Average
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_average_average_proto_init() }
func file_average_average_proto_init() {
	if File_average_average_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_average_average_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_average_average_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Average); i {
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
			RawDescriptor: file_average_average_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_average_average_proto_goTypes,
		DependencyIndexes: file_average_average_proto_depIdxs,
		MessageInfos:      file_average_average_proto_msgTypes,
	}.Build()
	File_average_average_proto = out.File
	file_average_average_proto_rawDesc = nil
	file_average_average_proto_goTypes = nil
	file_average_average_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AverageServiceClient is the client API for AverageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AverageServiceClient interface {
	ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (AverageService_ComputeAverageClient, error)
}

type averageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAverageServiceClient(cc grpc.ClientConnInterface) AverageServiceClient {
	return &averageServiceClient{cc}
}

func (c *averageServiceClient) ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (AverageService_ComputeAverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AverageService_serviceDesc.Streams[0], "/AverageService/ComputeAverage", opts...)
	if err != nil {
		return nil, err
	}
	x := &averageServiceComputeAverageClient{stream}
	return x, nil
}

type AverageService_ComputeAverageClient interface {
	Send(*Number) error
	CloseAndRecv() (*Average, error)
	grpc.ClientStream
}

type averageServiceComputeAverageClient struct {
	grpc.ClientStream
}

func (x *averageServiceComputeAverageClient) Send(m *Number) error {
	return x.ClientStream.SendMsg(m)
}

func (x *averageServiceComputeAverageClient) CloseAndRecv() (*Average, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Average)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AverageServiceServer is the server API for AverageService service.
type AverageServiceServer interface {
	ComputeAverage(AverageService_ComputeAverageServer) error
}

// UnimplementedAverageServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAverageServiceServer struct {
}

func (*UnimplementedAverageServiceServer) ComputeAverage(AverageService_ComputeAverageServer) error {
	return status.Errorf(codes.Unimplemented, "method ComputeAverage not implemented")
}

func RegisterAverageServiceServer(s *grpc.Server, srv AverageServiceServer) {
	s.RegisterService(&_AverageService_serviceDesc, srv)
}

func _AverageService_ComputeAverage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AverageServiceServer).ComputeAverage(&averageServiceComputeAverageServer{stream})
}

type AverageService_ComputeAverageServer interface {
	SendAndClose(*Average) error
	Recv() (*Number, error)
	grpc.ServerStream
}

type averageServiceComputeAverageServer struct {
	grpc.ServerStream
}

func (x *averageServiceComputeAverageServer) SendAndClose(m *Average) error {
	return x.ServerStream.SendMsg(m)
}

func (x *averageServiceComputeAverageServer) Recv() (*Number, error) {
	m := new(Number)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _AverageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "AverageService",
	HandlerType: (*AverageServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ComputeAverage",
			Handler:       _AverageService_ComputeAverage_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "average/average.proto",
}