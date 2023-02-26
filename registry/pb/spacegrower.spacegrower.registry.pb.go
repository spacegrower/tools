// protoc --go_out=. --go-grpc_out=. spacegrower.spacegrower.registry.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.3
// source: etc/proto/spacegrower.spacegrower.registry.proto

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

type MoultiService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Services []*ServiceInfo `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
}

func (x *MoultiService) Reset() {
	*x = MoultiService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_etc_proto_spacegrower_spacegrower_registry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MoultiService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoultiService) ProtoMessage() {}

func (x *MoultiService) ProtoReflect() protoreflect.Message {
	mi := &file_etc_proto_spacegrower_spacegrower_registry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoultiService.ProtoReflect.Descriptor instead.
func (*MoultiService) Descriptor() ([]byte, []int) {
	return file_etc_proto_spacegrower_spacegrower_registry_proto_rawDescGZIP(), []int{0}
}

func (x *MoultiService) GetServices() []*ServiceInfo {
	if x != nil {
		return x.Services
	}
	return nil
}

type ServiceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version     string            `protobuf:"bytes,1,opt,name=Version,proto3" json:"Version,omitempty"`
	Namespace   string            `protobuf:"bytes,2,opt,name=Namespace,proto3" json:"Namespace,omitempty"`
	ServiceName string            `protobuf:"bytes,3,opt,name=ServiceName,proto3" json:"ServiceName,omitempty"`
	Methods     []*MethodInfo     `protobuf:"bytes,4,rep,name=Methods,proto3" json:"Methods,omitempty"`
	Region      string            `protobuf:"bytes,5,opt,name=Region,proto3" json:"Region,omitempty"`
	Host        string            `protobuf:"bytes,6,opt,name=Host,proto3" json:"Host,omitempty"`
	Port        int32             `protobuf:"varint,7,opt,name=Port,proto3" json:"Port,omitempty"`
	Weight      int32             `protobuf:"varint,8,opt,name=Weight,proto3" json:"Weight,omitempty"`
	Runtime     string            `protobuf:"bytes,9,opt,name=Runtime,proto3" json:"Runtime,omitempty"`
	Tags        map[string]string `protobuf:"bytes,10,rep,name=Tags,proto3" json:"Tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	OrgID       string            `protobuf:"bytes,11,opt,name=OrgID,proto3" json:"OrgID,omitempty"`
}

func (x *ServiceInfo) Reset() {
	*x = ServiceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_etc_proto_spacegrower_spacegrower_registry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceInfo) ProtoMessage() {}

func (x *ServiceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_etc_proto_spacegrower_spacegrower_registry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceInfo.ProtoReflect.Descriptor instead.
func (*ServiceInfo) Descriptor() ([]byte, []int) {
	return file_etc_proto_spacegrower_spacegrower_registry_proto_rawDescGZIP(), []int{1}
}

func (x *ServiceInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ServiceInfo) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ServiceInfo) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *ServiceInfo) GetMethods() []*MethodInfo {
	if x != nil {
		return x.Methods
	}
	return nil
}

func (x *ServiceInfo) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *ServiceInfo) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *ServiceInfo) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *ServiceInfo) GetWeight() int32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *ServiceInfo) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

func (x *ServiceInfo) GetTags() map[string]string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *ServiceInfo) GetOrgID() string {
	if x != nil {
		return x.OrgID
	}
	return ""
}

type MethodInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	IsClientStream bool   `protobuf:"varint,2,opt,name=IsClientStream,proto3" json:"IsClientStream,omitempty"`
	IsServerStream bool   `protobuf:"varint,3,opt,name=IsServerStream,proto3" json:"IsServerStream,omitempty"`
}

func (x *MethodInfo) Reset() {
	*x = MethodInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_etc_proto_spacegrower_spacegrower_registry_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MethodInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MethodInfo) ProtoMessage() {}

func (x *MethodInfo) ProtoReflect() protoreflect.Message {
	mi := &file_etc_proto_spacegrower_spacegrower_registry_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MethodInfo.ProtoReflect.Descriptor instead.
func (*MethodInfo) Descriptor() ([]byte, []int) {
	return file_etc_proto_spacegrower_spacegrower_registry_proto_rawDescGZIP(), []int{2}
}

func (x *MethodInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MethodInfo) GetIsClientStream() bool {
	if x != nil {
		return x.IsClientStream
	}
	return false
}

func (x *MethodInfo) GetIsServerStream() bool {
	if x != nil {
		return x.IsServerStream
	}
	return false
}

type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command string            `protobuf:"bytes,1,opt,name=Command,proto3" json:"Command,omitempty"`
	Args    map[string]string `protobuf:"bytes,2,rep,name=Args,proto3" json:"Args,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_etc_proto_spacegrower_spacegrower_registry_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_etc_proto_spacegrower_spacegrower_registry_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_etc_proto_spacegrower_spacegrower_registry_proto_rawDescGZIP(), []int{3}
}

func (x *Command) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *Command) GetArgs() map[string]string {
	if x != nil {
		return x.Args
	}
	return nil
}

type ResolveInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address map[string][]byte `protobuf:"bytes,1,rep,name=Address,proto3" json:"Address,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Config  []byte            `protobuf:"bytes,2,opt,name=Config,proto3" json:"Config,omitempty"`
}

func (x *ResolveInfo) Reset() {
	*x = ResolveInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_etc_proto_spacegrower_spacegrower_registry_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolveInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveInfo) ProtoMessage() {}

func (x *ResolveInfo) ProtoReflect() protoreflect.Message {
	mi := &file_etc_proto_spacegrower_spacegrower_registry_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveInfo.ProtoReflect.Descriptor instead.
func (*ResolveInfo) Descriptor() ([]byte, []int) {
	return file_etc_proto_spacegrower_spacegrower_registry_proto_rawDescGZIP(), []int{4}
}

func (x *ResolveInfo) GetAddress() map[string][]byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *ResolveInfo) GetConfig() []byte {
	if x != nil {
		return x.Config
	}
	return nil
}

type TargetInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service string `protobuf:"bytes,1,opt,name=Service,proto3" json:"Service,omitempty"`
}

func (x *TargetInfo) Reset() {
	*x = TargetInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_etc_proto_spacegrower_spacegrower_registry_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TargetInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TargetInfo) ProtoMessage() {}

func (x *TargetInfo) ProtoReflect() protoreflect.Message {
	mi := &file_etc_proto_spacegrower_spacegrower_registry_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TargetInfo.ProtoReflect.Descriptor instead.
func (*TargetInfo) Descriptor() ([]byte, []int) {
	return file_etc_proto_spacegrower_spacegrower_registry_proto_rawDescGZIP(), []int{5}
}

func (x *TargetInfo) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

var File_etc_proto_spacegrower_spacegrower_registry_proto protoreflect.FileDescriptor

var file_etc_proto_spacegrower_spacegrower_registry_proto_rawDesc = []byte{
	0x0a, 0x30, 0x65, 0x74, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x67, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x67, 0x72, 0x6f,
	0x77, 0x65, 0x72, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x3f, 0x0a, 0x0d, 0x4d, 0x6f, 0x75,
	0x6c, 0x74, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x87, 0x03, 0x0a, 0x0b, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x48, 0x6f, 0x73,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x50, 0x6f, 0x72,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x75, 0x6e,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x52, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x54, 0x61, 0x67, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x04, 0x54, 0x61, 0x67, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x4f, 0x72, 0x67, 0x49, 0x44, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4f, 0x72, 0x67, 0x49, 0x44, 0x1a, 0x37, 0x0a, 0x09, 0x54,
	0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x70, 0x0a, 0x0a, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x49, 0x73, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e,
	0x49, 0x73, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x26,
	0x0a, 0x0e, 0x49, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x49, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x22, 0x8a, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x2c, 0x0a, 0x04,
	0x41, 0x72, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x41, 0x72, 0x67, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x37, 0x0a, 0x09, 0x41, 0x72,
	0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x9c, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x39, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x6c, 0x76, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x3a, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x26, 0x0a, 0x0a, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x18, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x32, 0xb4, 0x01, 0x0a, 0x08, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x34, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0e, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x3b, 0x0a,
	0x0d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x12, 0x14,
	0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x4d, 0x6f, 0x75, 0x6c, 0x74, 0x69, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x1a, 0x0e, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x35, 0x0a, 0x08, 0x52, 0x65,
	0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x12, 0x2e, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x30,
	0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_etc_proto_spacegrower_spacegrower_registry_proto_rawDescOnce sync.Once
	file_etc_proto_spacegrower_spacegrower_registry_proto_rawDescData = file_etc_proto_spacegrower_spacegrower_registry_proto_rawDesc
)

func file_etc_proto_spacegrower_spacegrower_registry_proto_rawDescGZIP() []byte {
	file_etc_proto_spacegrower_spacegrower_registry_proto_rawDescOnce.Do(func() {
		file_etc_proto_spacegrower_spacegrower_registry_proto_rawDescData = protoimpl.X.CompressGZIP(file_etc_proto_spacegrower_spacegrower_registry_proto_rawDescData)
	})
	return file_etc_proto_spacegrower_spacegrower_registry_proto_rawDescData
}

var file_etc_proto_spacegrower_spacegrower_registry_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_etc_proto_spacegrower_spacegrower_registry_proto_goTypes = []interface{}{
	(*MoultiService)(nil), // 0: space.MoultiService
	(*ServiceInfo)(nil),   // 1: space.ServiceInfo
	(*MethodInfo)(nil),    // 2: space.MethodInfo
	(*Command)(nil),       // 3: space.Command
	(*ResolveInfo)(nil),   // 4: space.ResolveInfo
	(*TargetInfo)(nil),    // 5: space.TargetInfo
	nil,                   // 6: space.ServiceInfo.TagsEntry
	nil,                   // 7: space.Command.ArgsEntry
	nil,                   // 8: space.ResolveInfo.AddressEntry
}
var file_etc_proto_spacegrower_spacegrower_registry_proto_depIdxs = []int32{
	1, // 0: space.MoultiService.services:type_name -> space.ServiceInfo
	2, // 1: space.ServiceInfo.Methods:type_name -> space.MethodInfo
	6, // 2: space.ServiceInfo.Tags:type_name -> space.ServiceInfo.TagsEntry
	7, // 3: space.Command.Args:type_name -> space.Command.ArgsEntry
	8, // 4: space.ResolveInfo.Address:type_name -> space.ResolveInfo.AddressEntry
	1, // 5: space.Registry.Register:input_type -> space.ServiceInfo
	0, // 6: space.Registry.RegisterMulti:input_type -> space.MoultiService
	5, // 7: space.Registry.Resolver:input_type -> space.TargetInfo
	3, // 8: space.Registry.Register:output_type -> space.Command
	3, // 9: space.Registry.RegisterMulti:output_type -> space.Command
	4, // 10: space.Registry.Resolver:output_type -> space.ResolveInfo
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_etc_proto_spacegrower_spacegrower_registry_proto_init() }
func file_etc_proto_spacegrower_spacegrower_registry_proto_init() {
	if File_etc_proto_spacegrower_spacegrower_registry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_etc_proto_spacegrower_spacegrower_registry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MoultiService); i {
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
		file_etc_proto_spacegrower_spacegrower_registry_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceInfo); i {
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
		file_etc_proto_spacegrower_spacegrower_registry_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MethodInfo); i {
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
		file_etc_proto_spacegrower_spacegrower_registry_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command); i {
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
		file_etc_proto_spacegrower_spacegrower_registry_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolveInfo); i {
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
		file_etc_proto_spacegrower_spacegrower_registry_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TargetInfo); i {
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
			RawDescriptor: file_etc_proto_spacegrower_spacegrower_registry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_etc_proto_spacegrower_spacegrower_registry_proto_goTypes,
		DependencyIndexes: file_etc_proto_spacegrower_spacegrower_registry_proto_depIdxs,
		MessageInfos:      file_etc_proto_spacegrower_spacegrower_registry_proto_msgTypes,
	}.Build()
	File_etc_proto_spacegrower_spacegrower_registry_proto = out.File
	file_etc_proto_spacegrower_spacegrower_registry_proto_rawDesc = nil
	file_etc_proto_spacegrower_spacegrower_registry_proto_goTypes = nil
	file_etc_proto_spacegrower_spacegrower_registry_proto_depIdxs = nil
}