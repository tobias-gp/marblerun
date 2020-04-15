// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.11.4
// source: rpc/coordinator.proto

package rpc

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ActivationReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Quote   []byte `protobuf:"bytes,1,opt,name=quote,proto3" json:"quote,omitempty"`
	Csr     string `protobuf:"bytes,2,opt,name=csr,proto3" json:"csr,omitempty"`
	Package string `protobuf:"bytes,3,opt,name=package,proto3" json:"package,omitempty"`
	Variant string `protobuf:"bytes,4,opt,name=variant,proto3" json:"variant,omitempty"`
}

func (x *ActivationReq) Reset() {
	*x = ActivationReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_coordinator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivationReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivationReq) ProtoMessage() {}

func (x *ActivationReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_coordinator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivationReq.ProtoReflect.Descriptor instead.
func (*ActivationReq) Descriptor() ([]byte, []int) {
	return file_rpc_coordinator_proto_rawDescGZIP(), []int{0}
}

func (x *ActivationReq) GetQuote() []byte {
	if x != nil {
		return x.Quote
	}
	return nil
}

func (x *ActivationReq) GetCsr() string {
	if x != nil {
		return x.Csr
	}
	return ""
}

func (x *ActivationReq) GetPackage() string {
	if x != nil {
		return x.Package
	}
	return ""
}

func (x *ActivationReq) GetVariant() string {
	if x != nil {
		return x.Variant
	}
	return ""
}

type ActivationRepl struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success     bool        `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Certificate string      `protobuf:"bytes,2,opt,name=certificate,proto3" json:"certificate,omitempty"`
	Params      *Parameters `protobuf:"bytes,3,opt,name=params,proto3" json:"params,omitempty"`
}

func (x *ActivationRepl) Reset() {
	*x = ActivationRepl{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_coordinator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivationRepl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivationRepl) ProtoMessage() {}

func (x *ActivationRepl) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_coordinator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivationRepl.ProtoReflect.Descriptor instead.
func (*ActivationRepl) Descriptor() ([]byte, []int) {
	return file_rpc_coordinator_proto_rawDescGZIP(), []int{1}
}

func (x *ActivationRepl) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ActivationRepl) GetCertificate() string {
	if x != nil {
		return x.Certificate
	}
	return ""
}

func (x *ActivationRepl) GetParams() *Parameters {
	if x != nil {
		return x.Params
	}
	return nil
}

type Parameters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Env  map[string]string `protobuf:"bytes,1,rep,name=env,proto3" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Argv []string          `protobuf:"bytes,2,rep,name=argv,proto3" json:"argv,omitempty"`
}

func (x *Parameters) Reset() {
	*x = Parameters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_coordinator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Parameters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Parameters) ProtoMessage() {}

func (x *Parameters) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_coordinator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Parameters.ProtoReflect.Descriptor instead.
func (*Parameters) Descriptor() ([]byte, []int) {
	return file_rpc_coordinator_proto_rawDescGZIP(), []int{2}
}

func (x *Parameters) GetEnv() map[string]string {
	if x != nil {
		return x.Env
	}
	return nil
}

func (x *Parameters) GetArgv() []string {
	if x != nil {
		return x.Argv
	}
	return nil
}

var File_rpc_coordinator_proto protoreflect.FileDescriptor

var file_rpc_coordinator_proto_rawDesc = []byte{
	0x0a, 0x15, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x70, 0x63, 0x22, 0x6b, 0x0a, 0x0d,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a,
	0x05, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x71, 0x75,
	0x6f, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x73, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x63, 0x73, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x22, 0x75, 0x0a, 0x0e, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x22, 0x84, 0x01, 0x0a, 0x0a, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12,
	0x2a, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x45, 0x6e,
	0x76, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x12, 0x12, 0x0a, 0x04, 0x61,
	0x72, 0x67, 0x76, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x76, 0x1a,
	0x36, 0x0a, 0x08, 0x45, 0x6e, 0x76, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x46, 0x0a, 0x0b, 0x43, 0x6f, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x37, 0x0a, 0x0c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x42,
	0x27, 0x5a, 0x25, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e,
	0x61, 0x74, 0x6f, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_coordinator_proto_rawDescOnce sync.Once
	file_rpc_coordinator_proto_rawDescData = file_rpc_coordinator_proto_rawDesc
)

func file_rpc_coordinator_proto_rawDescGZIP() []byte {
	file_rpc_coordinator_proto_rawDescOnce.Do(func() {
		file_rpc_coordinator_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_coordinator_proto_rawDescData)
	})
	return file_rpc_coordinator_proto_rawDescData
}

var file_rpc_coordinator_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_rpc_coordinator_proto_goTypes = []interface{}{
	(*ActivationReq)(nil),  // 0: rpc.ActivationReq
	(*ActivationRepl)(nil), // 1: rpc.ActivationRepl
	(*Parameters)(nil),     // 2: rpc.Parameters
	nil,                    // 3: rpc.Parameters.EnvEntry
}
var file_rpc_coordinator_proto_depIdxs = []int32{
	2, // 0: rpc.ActivationRepl.params:type_name -> rpc.Parameters
	3, // 1: rpc.Parameters.env:type_name -> rpc.Parameters.EnvEntry
	0, // 2: rpc.Coordinator.ActivateNode:input_type -> rpc.ActivationReq
	1, // 3: rpc.Coordinator.ActivateNode:output_type -> rpc.ActivationRepl
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_rpc_coordinator_proto_init() }
func file_rpc_coordinator_proto_init() {
	if File_rpc_coordinator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_coordinator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivationReq); i {
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
		file_rpc_coordinator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivationRepl); i {
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
		file_rpc_coordinator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Parameters); i {
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
			RawDescriptor: file_rpc_coordinator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_coordinator_proto_goTypes,
		DependencyIndexes: file_rpc_coordinator_proto_depIdxs,
		MessageInfos:      file_rpc_coordinator_proto_msgTypes,
	}.Build()
	File_rpc_coordinator_proto = out.File
	file_rpc_coordinator_proto_rawDesc = nil
	file_rpc_coordinator_proto_goTypes = nil
	file_rpc_coordinator_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CoordinatorClient is the client API for Coordinator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CoordinatorClient interface {
	ActivateNode(ctx context.Context, in *ActivationReq, opts ...grpc.CallOption) (*ActivationRepl, error)
}

type coordinatorClient struct {
	cc grpc.ClientConnInterface
}

func NewCoordinatorClient(cc grpc.ClientConnInterface) CoordinatorClient {
	return &coordinatorClient{cc}
}

func (c *coordinatorClient) ActivateNode(ctx context.Context, in *ActivationReq, opts ...grpc.CallOption) (*ActivationRepl, error) {
	out := new(ActivationRepl)
	err := c.cc.Invoke(ctx, "/rpc.Coordinator/ActivateNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoordinatorServer is the server API for Coordinator service.
type CoordinatorServer interface {
	ActivateNode(context.Context, *ActivationReq) (*ActivationRepl, error)
}

// UnimplementedCoordinatorServer can be embedded to have forward compatible implementations.
type UnimplementedCoordinatorServer struct {
}

func (*UnimplementedCoordinatorServer) ActivateNode(context.Context, *ActivationReq) (*ActivationRepl, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivateNode not implemented")
}

func RegisterCoordinatorServer(s *grpc.Server, srv CoordinatorServer) {
	s.RegisterService(&_Coordinator_serviceDesc, srv)
}

func _Coordinator_ActivateNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).ActivateNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Coordinator/ActivateNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).ActivateNode(ctx, req.(*ActivationReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Coordinator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.Coordinator",
	HandlerType: (*CoordinatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ActivateNode",
			Handler:    _Coordinator_ActivateNode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/coordinator.proto",
}
