// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.9
// source: roulette.proto

package proto

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

type NewBet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bet int32 `protobuf:"varint,1,opt,name=bet,proto3" json:"bet,omitempty"`
}

func (x *NewBet) Reset() {
	*x = NewBet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_roulette_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewBet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewBet) ProtoMessage() {}

func (x *NewBet) ProtoReflect() protoreflect.Message {
	mi := &file_roulette_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewBet.ProtoReflect.Descriptor instead.
func (*NewBet) Descriptor() ([]byte, []int) {
	return file_roulette_proto_rawDescGZIP(), []int{0}
}

func (x *NewBet) GetBet() int32 {
	if x != nil {
		return x.Bet
	}
	return 0
}

type BetResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *BetResult) Reset() {
	*x = BetResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_roulette_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BetResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BetResult) ProtoMessage() {}

func (x *BetResult) ProtoReflect() protoreflect.Message {
	mi := &file_roulette_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BetResult.ProtoReflect.Descriptor instead.
func (*BetResult) Descriptor() ([]byte, []int) {
	return file_roulette_proto_rawDescGZIP(), []int{1}
}

func (x *BetResult) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_roulette_proto protoreflect.FileDescriptor

var file_roulette_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x6f, 0x75, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x1a, 0x0a, 0x06, 0x4e, 0x65, 0x77, 0x42, 0x65, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x62, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x62, 0x65,
	0x74, 0x22, 0x23, 0x0a, 0x09, 0x42, 0x65, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x31, 0x0a, 0x08, 0x52, 0x6f, 0x75, 0x6c, 0x65, 0x74,
	0x74, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x53, 0x70, 0x69, 0x6e, 0x12, 0x0b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x4e, 0x65, 0x77, 0x42, 0x65, 0x74, 0x1a, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x70, 0x6b, 0x67,
	0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_roulette_proto_rawDescOnce sync.Once
	file_roulette_proto_rawDescData = file_roulette_proto_rawDesc
)

func file_roulette_proto_rawDescGZIP() []byte {
	file_roulette_proto_rawDescOnce.Do(func() {
		file_roulette_proto_rawDescData = protoimpl.X.CompressGZIP(file_roulette_proto_rawDescData)
	})
	return file_roulette_proto_rawDescData
}

var file_roulette_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_roulette_proto_goTypes = []interface{}{
	(*NewBet)(nil),    // 0: api.NewBet
	(*BetResult)(nil), // 1: api.BetResult
}
var file_roulette_proto_depIdxs = []int32{
	0, // 0: api.Roulette.Spin:input_type -> api.NewBet
	1, // 1: api.Roulette.Spin:output_type -> api.BetResult
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_roulette_proto_init() }
func file_roulette_proto_init() {
	if File_roulette_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_roulette_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewBet); i {
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
		file_roulette_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BetResult); i {
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
			RawDescriptor: file_roulette_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_roulette_proto_goTypes,
		DependencyIndexes: file_roulette_proto_depIdxs,
		MessageInfos:      file_roulette_proto_msgTypes,
	}.Build()
	File_roulette_proto = out.File
	file_roulette_proto_rawDesc = nil
	file_roulette_proto_goTypes = nil
	file_roulette_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RouletteClient is the client API for Roulette service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RouletteClient interface {
	Spin(ctx context.Context, in *NewBet, opts ...grpc.CallOption) (*BetResult, error)
}

type rouletteClient struct {
	cc grpc.ClientConnInterface
}

func NewRouletteClient(cc grpc.ClientConnInterface) RouletteClient {
	return &rouletteClient{cc}
}

func (c *rouletteClient) Spin(ctx context.Context, in *NewBet, opts ...grpc.CallOption) (*BetResult, error) {
	out := new(BetResult)
	err := c.cc.Invoke(ctx, "/api.Roulette/Spin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouletteServer is the server API for Roulette service.
type RouletteServer interface {
	Spin(context.Context, *NewBet) (*BetResult, error)
}

// UnimplementedRouletteServer can be embedded to have forward compatible implementations.
type UnimplementedRouletteServer struct {
}

func (*UnimplementedRouletteServer) Spin(context.Context, *NewBet) (*BetResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Spin not implemented")
}

func RegisterRouletteServer(s *grpc.Server, srv RouletteServer) {
	s.RegisterService(&_Roulette_serviceDesc, srv)
}

func _Roulette_Spin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewBet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouletteServer).Spin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Roulette/Spin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouletteServer).Spin(ctx, req.(*NewBet))
	}
	return interceptor(ctx, in, info, handler)
}

var _Roulette_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Roulette",
	HandlerType: (*RouletteServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Spin",
			Handler:    _Roulette_Spin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "roulette.proto",
}