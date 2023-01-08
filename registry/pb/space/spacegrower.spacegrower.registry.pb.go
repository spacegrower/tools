// Code generated by protoc-gen-go. DO NOT EDIT.
// source: spacegrower.spacegrower.registry.proto

package space

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ServiceInfo struct {
	Version              string            `protobuf:"bytes,1,opt,name=Version,proto3" json:"Version,omitempty"`
	Namespace            string            `protobuf:"bytes,2,opt,name=Namespace,proto3" json:"Namespace,omitempty"`
	ServiceName          string            `protobuf:"bytes,3,opt,name=ServiceName,proto3" json:"ServiceName,omitempty"`
	Methods              []*MethodInfo     `protobuf:"bytes,4,rep,name=Methods,proto3" json:"Methods,omitempty"`
	Region               string            `protobuf:"bytes,5,opt,name=Region,proto3" json:"Region,omitempty"`
	Host                 string            `protobuf:"bytes,6,opt,name=Host,proto3" json:"Host,omitempty"`
	Port                 int32             `protobuf:"varint,7,opt,name=Port,proto3" json:"Port,omitempty"`
	Weight               int32             `protobuf:"varint,8,opt,name=Weight,proto3" json:"Weight,omitempty"`
	Runtime              string            `protobuf:"bytes,9,opt,name=Runtime,proto3" json:"Runtime,omitempty"`
	Tags                 map[string]string `protobuf:"bytes,10,rep,name=Tags,proto3" json:"Tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	OrgID                string            `protobuf:"bytes,11,opt,name=OrgID,proto3" json:"OrgID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ServiceInfo) Reset()         { *m = ServiceInfo{} }
func (m *ServiceInfo) String() string { return proto.CompactTextString(m) }
func (*ServiceInfo) ProtoMessage()    {}
func (*ServiceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a27cb2a1d792ff, []int{0}
}

func (m *ServiceInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceInfo.Unmarshal(m, b)
}
func (m *ServiceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceInfo.Marshal(b, m, deterministic)
}
func (m *ServiceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceInfo.Merge(m, src)
}
func (m *ServiceInfo) XXX_Size() int {
	return xxx_messageInfo_ServiceInfo.Size(m)
}
func (m *ServiceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceInfo proto.InternalMessageInfo

func (m *ServiceInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ServiceInfo) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ServiceInfo) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *ServiceInfo) GetMethods() []*MethodInfo {
	if m != nil {
		return m.Methods
	}
	return nil
}

func (m *ServiceInfo) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *ServiceInfo) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *ServiceInfo) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *ServiceInfo) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *ServiceInfo) GetRuntime() string {
	if m != nil {
		return m.Runtime
	}
	return ""
}

func (m *ServiceInfo) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *ServiceInfo) GetOrgID() string {
	if m != nil {
		return m.OrgID
	}
	return ""
}

type MethodInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	IsClientStream       bool     `protobuf:"varint,2,opt,name=IsClientStream,proto3" json:"IsClientStream,omitempty"`
	IsServerStream       bool     `protobuf:"varint,3,opt,name=IsServerStream,proto3" json:"IsServerStream,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MethodInfo) Reset()         { *m = MethodInfo{} }
func (m *MethodInfo) String() string { return proto.CompactTextString(m) }
func (*MethodInfo) ProtoMessage()    {}
func (*MethodInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a27cb2a1d792ff, []int{1}
}

func (m *MethodInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MethodInfo.Unmarshal(m, b)
}
func (m *MethodInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MethodInfo.Marshal(b, m, deterministic)
}
func (m *MethodInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MethodInfo.Merge(m, src)
}
func (m *MethodInfo) XXX_Size() int {
	return xxx_messageInfo_MethodInfo.Size(m)
}
func (m *MethodInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MethodInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MethodInfo proto.InternalMessageInfo

func (m *MethodInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MethodInfo) GetIsClientStream() bool {
	if m != nil {
		return m.IsClientStream
	}
	return false
}

func (m *MethodInfo) GetIsServerStream() bool {
	if m != nil {
		return m.IsServerStream
	}
	return false
}

type Command struct {
	Command              string            `protobuf:"bytes,1,opt,name=Command,proto3" json:"Command,omitempty"`
	Args                 map[string]string `protobuf:"bytes,2,rep,name=Args,proto3" json:"Args,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Command) Reset()         { *m = Command{} }
func (m *Command) String() string { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()    {}
func (*Command) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a27cb2a1d792ff, []int{2}
}

func (m *Command) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Command.Unmarshal(m, b)
}
func (m *Command) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Command.Marshal(b, m, deterministic)
}
func (m *Command) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Command.Merge(m, src)
}
func (m *Command) XXX_Size() int {
	return xxx_messageInfo_Command.Size(m)
}
func (m *Command) XXX_DiscardUnknown() {
	xxx_messageInfo_Command.DiscardUnknown(m)
}

var xxx_messageInfo_Command proto.InternalMessageInfo

func (m *Command) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *Command) GetArgs() map[string]string {
	if m != nil {
		return m.Args
	}
	return nil
}

type ResolveInfo struct {
	Address              map[string][]byte `protobuf:"bytes,1,rep,name=Address,proto3" json:"Address,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Config               []byte            `protobuf:"bytes,2,opt,name=Config,proto3" json:"Config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ResolveInfo) Reset()         { *m = ResolveInfo{} }
func (m *ResolveInfo) String() string { return proto.CompactTextString(m) }
func (*ResolveInfo) ProtoMessage()    {}
func (*ResolveInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a27cb2a1d792ff, []int{3}
}

func (m *ResolveInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResolveInfo.Unmarshal(m, b)
}
func (m *ResolveInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResolveInfo.Marshal(b, m, deterministic)
}
func (m *ResolveInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResolveInfo.Merge(m, src)
}
func (m *ResolveInfo) XXX_Size() int {
	return xxx_messageInfo_ResolveInfo.Size(m)
}
func (m *ResolveInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ResolveInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ResolveInfo proto.InternalMessageInfo

func (m *ResolveInfo) GetAddress() map[string][]byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *ResolveInfo) GetConfig() []byte {
	if m != nil {
		return m.Config
	}
	return nil
}

type TargetInfo struct {
	Service              string   `protobuf:"bytes,1,opt,name=Service,proto3" json:"Service,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TargetInfo) Reset()         { *m = TargetInfo{} }
func (m *TargetInfo) String() string { return proto.CompactTextString(m) }
func (*TargetInfo) ProtoMessage()    {}
func (*TargetInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a27cb2a1d792ff, []int{4}
}

func (m *TargetInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetInfo.Unmarshal(m, b)
}
func (m *TargetInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetInfo.Marshal(b, m, deterministic)
}
func (m *TargetInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetInfo.Merge(m, src)
}
func (m *TargetInfo) XXX_Size() int {
	return xxx_messageInfo_TargetInfo.Size(m)
}
func (m *TargetInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TargetInfo proto.InternalMessageInfo

func (m *TargetInfo) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func init() {
	proto.RegisterType((*ServiceInfo)(nil), "space.ServiceInfo")
	proto.RegisterMapType((map[string]string)(nil), "space.ServiceInfo.TagsEntry")
	proto.RegisterType((*MethodInfo)(nil), "space.MethodInfo")
	proto.RegisterType((*Command)(nil), "space.Command")
	proto.RegisterMapType((map[string]string)(nil), "space.Command.ArgsEntry")
	proto.RegisterType((*ResolveInfo)(nil), "space.ResolveInfo")
	proto.RegisterMapType((map[string][]byte)(nil), "space.ResolveInfo.AddressEntry")
	proto.RegisterType((*TargetInfo)(nil), "space.TargetInfo")
}

func init() {
	proto.RegisterFile("spacegrower.spacegrower.registry.proto", fileDescriptor_15a27cb2a1d792ff)
}

var fileDescriptor_15a27cb2a1d792ff = []byte{
	// 506 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4f, 0x8b, 0x13, 0x4f,
	0x10, 0xdd, 0xce, 0x24, 0x99, 0xa4, 0x12, 0x96, 0xdf, 0xf6, 0x4f, 0xa4, 0x09, 0x0b, 0x86, 0x39,
	0x84, 0x80, 0x32, 0x86, 0x55, 0x51, 0xf7, 0xb6, 0x46, 0xc1, 0x1c, 0xfc, 0x43, 0xef, 0xa2, 0xe0,
	0x6d, 0xdc, 0xd4, 0xce, 0x0e, 0x66, 0xa6, 0x43, 0x77, 0x6f, 0x96, 0x7c, 0x02, 0xc1, 0xcf, 0xe0,
	0x87, 0x95, 0xae, 0xe9, 0x4e, 0x46, 0xf7, 0x20, 0xde, 0x5e, 0xbd, 0x7a, 0xd5, 0x55, 0xf5, 0x66,
	0x0a, 0x26, 0x66, 0x9d, 0x5d, 0x62, 0xae, 0xd5, 0x2d, 0xea, 0xb4, 0x89, 0x35, 0xe6, 0x85, 0xb1,
	0x7a, 0x9b, 0xae, 0xb5, 0xb2, 0x8a, 0x77, 0x28, 0x97, 0x7c, 0x8f, 0x60, 0x70, 0x8e, 0x7a, 0x53,
	0x5c, 0xe2, 0xa2, 0xba, 0x52, 0x5c, 0x40, 0xfc, 0x09, 0xb5, 0x29, 0x54, 0x25, 0xd8, 0x98, 0x4d,
	0xfb, 0x32, 0x84, 0xfc, 0x18, 0xfa, 0xef, 0xb3, 0x12, 0xa9, 0x4c, 0xb4, 0x28, 0xb7, 0x27, 0xf8,
	0x78, 0xf7, 0x8c, 0xe3, 0x44, 0x44, 0xf9, 0x26, 0xc5, 0x1f, 0x42, 0xfc, 0x0e, 0xed, 0xb5, 0x5a,
	0x1a, 0xd1, 0x1e, 0x47, 0xd3, 0xc1, 0xc9, 0x51, 0x3d, 0x5e, 0x5a, 0xb3, 0xae, 0xbb, 0x0c, 0x0a,
	0x7e, 0x1f, 0xba, 0x12, 0x73, 0x37, 0x45, 0x87, 0x5e, 0xf2, 0x11, 0xe7, 0xd0, 0x7e, 0xab, 0x8c,
	0x15, 0x5d, 0x62, 0x09, 0x3b, 0xee, 0xa3, 0xd2, 0x56, 0xc4, 0x63, 0x36, 0xed, 0x48, 0xc2, 0xae,
	0xfe, 0x33, 0x16, 0xf9, 0xb5, 0x15, 0x3d, 0x62, 0x7d, 0xe4, 0xd6, 0x93, 0x37, 0x95, 0x2d, 0x4a,
	0x14, 0xfd, 0x7a, 0x3d, 0x1f, 0xf2, 0x19, 0xb4, 0x2f, 0xb2, 0xdc, 0x08, 0xa0, 0xd9, 0x8e, 0xfd,
	0x6c, 0x0d, 0x6b, 0x52, 0x97, 0x7e, 0x53, 0x59, 0xbd, 0x95, 0xa4, 0xe4, 0xf7, 0xa0, 0xf3, 0x41,
	0xe7, 0x8b, 0xd7, 0x62, 0x40, 0x2f, 0xd5, 0xc1, 0xe8, 0x39, 0xf4, 0x77, 0x42, 0xfe, 0x1f, 0x44,
	0xdf, 0x70, 0xeb, 0x9d, 0x74, 0xd0, 0x15, 0x6d, 0xb2, 0xd5, 0x4d, 0x70, 0xb0, 0x0e, 0x4e, 0x5b,
	0x2f, 0x58, 0xb2, 0x06, 0xd8, 0x3b, 0xe1, 0x96, 0x22, 0x23, 0xeb, 0x52, 0xc2, 0x7c, 0x02, 0x87,
	0x0b, 0x33, 0x5f, 0x15, 0x58, 0xd9, 0x73, 0xab, 0x31, 0x2b, 0xe9, 0x91, 0x9e, 0xfc, 0x83, 0xad,
	0x75, 0x6e, 0x72, 0xd4, 0x5e, 0x17, 0x05, 0x5d, 0x93, 0x4d, 0x7e, 0x30, 0x88, 0xe7, 0xaa, 0x2c,
	0xb3, 0x6a, 0xe9, 0x8c, 0xf1, 0x30, 0x7c, 0xf7, 0x90, 0x79, 0x04, 0xed, 0x33, 0x9d, 0x1b, 0xd1,
	0x22, 0x63, 0x84, 0x37, 0xc6, 0x67, 0x53, 0x97, 0xf2, 0xa6, 0x38, 0xe8, 0xd6, 0xdf, 0x51, 0xff,
	0xb4, 0xfe, 0x4f, 0x06, 0x03, 0x89, 0x46, 0xad, 0x36, 0xf5, 0x8f, 0xf8, 0x12, 0xe2, 0xb3, 0xe5,
	0x52, 0xa3, 0x31, 0x82, 0x51, 0xe7, 0x07, 0xbe, 0x73, 0x43, 0x94, 0x7a, 0x45, 0x3d, 0x40, 0xd0,
	0xbb, 0x8f, 0x3f, 0x57, 0xd5, 0x55, 0x91, 0x53, 0x97, 0xa1, 0xf4, 0xd1, 0xe8, 0x14, 0x86, 0xcd,
	0x82, 0xbf, 0x8d, 0x37, 0x6c, 0x8e, 0x37, 0x01, 0xb8, 0xc8, 0x74, 0x8e, 0x36, 0x5c, 0x89, 0xff,
	0x33, 0x82, 0x5b, 0x3e, 0x3c, 0xb9, 0x85, 0x9e, 0xf4, 0x87, 0xc6, 0x9f, 0x06, 0x8c, 0x9a, 0xf3,
	0xbb, 0x3f, 0xd4, 0xe8, 0xf0, 0x77, 0x2f, 0x93, 0x83, 0x29, 0x9b, 0x31, 0xfe, 0xcc, 0x55, 0xd1,
	0x8a, 0x9a, 0x87, 0x13, 0xd9, 0xb7, 0x1e, 0xf1, 0xbb, 0x36, 0x24, 0x07, 0x33, 0xf6, 0xea, 0xff,
	0x2f, 0x47, 0x8d, 0x6b, 0x7f, 0x4c, 0xf8, 0x6b, 0x97, 0x6e, 0xfd, 0xc9, 0xaf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x4b, 0x7b, 0xf9, 0x7d, 0x15, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RegistryClient is the client API for Registry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RegistryClient interface {
	Register(ctx context.Context, opts ...grpc.CallOption) (Registry_RegisterClient, error)
	Resolver(ctx context.Context, in *TargetInfo, opts ...grpc.CallOption) (Registry_ResolverClient, error)
}

type registryClient struct {
	cc grpc.ClientConnInterface
}

func NewRegistryClient(cc grpc.ClientConnInterface) RegistryClient {
	return &registryClient{cc}
}

func (c *registryClient) Register(ctx context.Context, opts ...grpc.CallOption) (Registry_RegisterClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Registry_serviceDesc.Streams[0], "/space.Registry/Register", opts...)
	if err != nil {
		return nil, err
	}
	x := &registryRegisterClient{stream}
	return x, nil
}

type Registry_RegisterClient interface {
	Send(*ServiceInfo) error
	Recv() (*Command, error)
	grpc.ClientStream
}

type registryRegisterClient struct {
	grpc.ClientStream
}

func (x *registryRegisterClient) Send(m *ServiceInfo) error {
	return x.ClientStream.SendMsg(m)
}

func (x *registryRegisterClient) Recv() (*Command, error) {
	m := new(Command)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *registryClient) Resolver(ctx context.Context, in *TargetInfo, opts ...grpc.CallOption) (Registry_ResolverClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Registry_serviceDesc.Streams[1], "/space.Registry/Resolver", opts...)
	if err != nil {
		return nil, err
	}
	x := &registryResolverClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Registry_ResolverClient interface {
	Recv() (*ResolveInfo, error)
	grpc.ClientStream
}

type registryResolverClient struct {
	grpc.ClientStream
}

func (x *registryResolverClient) Recv() (*ResolveInfo, error) {
	m := new(ResolveInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RegistryServer is the server API for Registry service.
type RegistryServer interface {
	Register(Registry_RegisterServer) error
	Resolver(*TargetInfo, Registry_ResolverServer) error
}

// UnimplementedRegistryServer can be embedded to have forward compatible implementations.
type UnimplementedRegistryServer struct {
}

func (*UnimplementedRegistryServer) Register(srv Registry_RegisterServer) error {
	return status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedRegistryServer) Resolver(req *TargetInfo, srv Registry_ResolverServer) error {
	return status.Errorf(codes.Unimplemented, "method Resolver not implemented")
}

func RegisterRegistryServer(s *grpc.Server, srv RegistryServer) {
	s.RegisterService(&_Registry_serviceDesc, srv)
}

func _Registry_Register_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RegistryServer).Register(&registryRegisterServer{stream})
}

type Registry_RegisterServer interface {
	Send(*Command) error
	Recv() (*ServiceInfo, error)
	grpc.ServerStream
}

type registryRegisterServer struct {
	grpc.ServerStream
}

func (x *registryRegisterServer) Send(m *Command) error {
	return x.ServerStream.SendMsg(m)
}

func (x *registryRegisterServer) Recv() (*ServiceInfo, error) {
	m := new(ServiceInfo)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Registry_Resolver_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TargetInfo)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RegistryServer).Resolver(m, &registryResolverServer{stream})
}

type Registry_ResolverServer interface {
	Send(*ResolveInfo) error
	grpc.ServerStream
}

type registryResolverServer struct {
	grpc.ServerStream
}

func (x *registryResolverServer) Send(m *ResolveInfo) error {
	return x.ServerStream.SendMsg(m)
}

var _Registry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "space.Registry",
	HandlerType: (*RegistryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Register",
			Handler:       _Registry_Register_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Resolver",
			Handler:       _Registry_Resolver_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "spacegrower.spacegrower.registry.proto",
}
