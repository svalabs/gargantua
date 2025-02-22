// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: environment/environment.proto

package environmentpb

import (
	general "github.com/hobbyfarm/gargantua/v3/protos/general"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Environment struct {
	state                protoimpl.MessageState        `protogen:"open.v1"`
	Id                   string                        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid                  string                        `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	DisplayName          string                        `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Dnssuffix            string                        `protobuf:"bytes,4,opt,name=dnssuffix,proto3" json:"dnssuffix,omitempty"`
	Provider             string                        `protobuf:"bytes,5,opt,name=provider,proto3" json:"provider,omitempty"`
	TemplateMapping      map[string]*general.StringMap `protobuf:"bytes,6,rep,name=template_mapping,json=templateMapping,proto3" json:"template_mapping,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	EnvironmentSpecifics map[string]string             `protobuf:"bytes,7,rep,name=environment_specifics,json=environmentSpecifics,proto3" json:"environment_specifics,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	IpTranslationMap     map[string]string             `protobuf:"bytes,8,rep,name=ip_translation_map,json=ipTranslationMap,proto3" json:"ip_translation_map,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	WsEndpoint           string                        `protobuf:"bytes,9,opt,name=ws_endpoint,json=wsEndpoint,proto3" json:"ws_endpoint,omitempty"`
	CountCapacity        map[string]uint32             `protobuf:"bytes,10,rep,name=count_capacity,json=countCapacity,proto3" json:"count_capacity,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	Annotations          map[string]string             `protobuf:"bytes,11,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *Environment) Reset() {
	*x = Environment{}
	mi := &file_environment_environment_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Environment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Environment) ProtoMessage() {}

func (x *Environment) ProtoReflect() protoreflect.Message {
	mi := &file_environment_environment_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Environment.ProtoReflect.Descriptor instead.
func (*Environment) Descriptor() ([]byte, []int) {
	return file_environment_environment_proto_rawDescGZIP(), []int{0}
}

func (x *Environment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Environment) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Environment) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Environment) GetDnssuffix() string {
	if x != nil {
		return x.Dnssuffix
	}
	return ""
}

func (x *Environment) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *Environment) GetTemplateMapping() map[string]*general.StringMap {
	if x != nil {
		return x.TemplateMapping
	}
	return nil
}

func (x *Environment) GetEnvironmentSpecifics() map[string]string {
	if x != nil {
		return x.EnvironmentSpecifics
	}
	return nil
}

func (x *Environment) GetIpTranslationMap() map[string]string {
	if x != nil {
		return x.IpTranslationMap
	}
	return nil
}

func (x *Environment) GetWsEndpoint() string {
	if x != nil {
		return x.WsEndpoint
	}
	return ""
}

func (x *Environment) GetCountCapacity() map[string]uint32 {
	if x != nil {
		return x.CountCapacity
	}
	return nil
}

func (x *Environment) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

type CreateEnvironmentRequest struct {
	state                protoimpl.MessageState `protogen:"open.v1"`
	DisplayName          string                 `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Dnssuffix            string                 `protobuf:"bytes,2,opt,name=dnssuffix,proto3" json:"dnssuffix,omitempty"`
	Provider             string                 `protobuf:"bytes,3,opt,name=provider,proto3" json:"provider,omitempty"`
	TemplateMapping      string                 `protobuf:"bytes,4,opt,name=template_mapping,json=templateMapping,proto3" json:"template_mapping,omitempty"`
	EnvironmentSpecifics string                 `protobuf:"bytes,5,opt,name=environment_specifics,json=environmentSpecifics,proto3" json:"environment_specifics,omitempty"`
	IpTranslationMap     string                 `protobuf:"bytes,6,opt,name=ip_translation_map,json=ipTranslationMap,proto3" json:"ip_translation_map,omitempty"`
	WsEndpoint           string                 `protobuf:"bytes,7,opt,name=ws_endpoint,json=wsEndpoint,proto3" json:"ws_endpoint,omitempty"`
	CountCapacity        string                 `protobuf:"bytes,8,opt,name=count_capacity,json=countCapacity,proto3" json:"count_capacity,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *CreateEnvironmentRequest) Reset() {
	*x = CreateEnvironmentRequest{}
	mi := &file_environment_environment_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateEnvironmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEnvironmentRequest) ProtoMessage() {}

func (x *CreateEnvironmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_environment_environment_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEnvironmentRequest.ProtoReflect.Descriptor instead.
func (*CreateEnvironmentRequest) Descriptor() ([]byte, []int) {
	return file_environment_environment_proto_rawDescGZIP(), []int{1}
}

func (x *CreateEnvironmentRequest) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *CreateEnvironmentRequest) GetDnssuffix() string {
	if x != nil {
		return x.Dnssuffix
	}
	return ""
}

func (x *CreateEnvironmentRequest) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *CreateEnvironmentRequest) GetTemplateMapping() string {
	if x != nil {
		return x.TemplateMapping
	}
	return ""
}

func (x *CreateEnvironmentRequest) GetEnvironmentSpecifics() string {
	if x != nil {
		return x.EnvironmentSpecifics
	}
	return ""
}

func (x *CreateEnvironmentRequest) GetIpTranslationMap() string {
	if x != nil {
		return x.IpTranslationMap
	}
	return ""
}

func (x *CreateEnvironmentRequest) GetWsEndpoint() string {
	if x != nil {
		return x.WsEndpoint
	}
	return ""
}

func (x *CreateEnvironmentRequest) GetCountCapacity() string {
	if x != nil {
		return x.CountCapacity
	}
	return ""
}

type UpdateEnvironmentRequest struct {
	state                protoimpl.MessageState  `protogen:"open.v1"`
	Id                   string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DisplayName          string                  `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Dnssuffix            *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=dnssuffix,proto3" json:"dnssuffix,omitempty"` // optional
	Provider             string                  `protobuf:"bytes,4,opt,name=provider,proto3" json:"provider,omitempty"`
	TemplateMapping      string                  `protobuf:"bytes,5,opt,name=template_mapping,json=templateMapping,proto3" json:"template_mapping,omitempty"`
	EnvironmentSpecifics string                  `protobuf:"bytes,6,opt,name=environment_specifics,json=environmentSpecifics,proto3" json:"environment_specifics,omitempty"`
	IpTranslationMap     string                  `protobuf:"bytes,7,opt,name=ip_translation_map,json=ipTranslationMap,proto3" json:"ip_translation_map,omitempty"`
	WsEndpoint           string                  `protobuf:"bytes,8,opt,name=ws_endpoint,json=wsEndpoint,proto3" json:"ws_endpoint,omitempty"`
	CountCapacity        string                  `protobuf:"bytes,9,opt,name=count_capacity,json=countCapacity,proto3" json:"count_capacity,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *UpdateEnvironmentRequest) Reset() {
	*x = UpdateEnvironmentRequest{}
	mi := &file_environment_environment_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateEnvironmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEnvironmentRequest) ProtoMessage() {}

func (x *UpdateEnvironmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_environment_environment_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEnvironmentRequest.ProtoReflect.Descriptor instead.
func (*UpdateEnvironmentRequest) Descriptor() ([]byte, []int) {
	return file_environment_environment_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateEnvironmentRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateEnvironmentRequest) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *UpdateEnvironmentRequest) GetDnssuffix() *wrapperspb.StringValue {
	if x != nil {
		return x.Dnssuffix
	}
	return nil
}

func (x *UpdateEnvironmentRequest) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *UpdateEnvironmentRequest) GetTemplateMapping() string {
	if x != nil {
		return x.TemplateMapping
	}
	return ""
}

func (x *UpdateEnvironmentRequest) GetEnvironmentSpecifics() string {
	if x != nil {
		return x.EnvironmentSpecifics
	}
	return ""
}

func (x *UpdateEnvironmentRequest) GetIpTranslationMap() string {
	if x != nil {
		return x.IpTranslationMap
	}
	return ""
}

func (x *UpdateEnvironmentRequest) GetWsEndpoint() string {
	if x != nil {
		return x.WsEndpoint
	}
	return ""
}

func (x *UpdateEnvironmentRequest) GetCountCapacity() string {
	if x != nil {
		return x.CountCapacity
	}
	return ""
}

type ListEnvironmentsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Environments  []*Environment         `protobuf:"bytes,1,rep,name=environments,proto3" json:"environments,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListEnvironmentsResponse) Reset() {
	*x = ListEnvironmentsResponse{}
	mi := &file_environment_environment_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListEnvironmentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEnvironmentsResponse) ProtoMessage() {}

func (x *ListEnvironmentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_environment_environment_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEnvironmentsResponse.ProtoReflect.Descriptor instead.
func (*ListEnvironmentsResponse) Descriptor() ([]byte, []int) {
	return file_environment_environment_proto_rawDescGZIP(), []int{3}
}

func (x *ListEnvironmentsResponse) GetEnvironments() []*Environment {
	if x != nil {
		return x.Environments
	}
	return nil
}

var File_environment_environment_proto protoreflect.FileDescriptor

var file_environment_environment_proto_rawDesc = string([]byte{
	0x0a, 0x1d, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x65, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x15, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xd7, 0x07, 0x0a, 0x0b, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x6e, 0x73, 0x73, 0x75, 0x66, 0x66,
	0x69, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x6e, 0x73, 0x73, 0x75, 0x66,
	0x66, 0x69, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12,
	0x58, 0x0a, 0x10, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x65, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x67, 0x0a, 0x15, 0x65, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69,
	0x63, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x70,
	0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x14, 0x65, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69,
	0x63, 0x73, 0x12, 0x5c, 0x0a, 0x12, 0x69, 0x70, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e,
	0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x49, 0x70, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x10,
	0x69, 0x70, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x70,
	0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x73, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x73, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x12, 0x52, 0x0a, 0x0e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x63, 0x61, 0x70, 0x61, 0x63,
	0x69, 0x74, 0x79, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x65, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74,
	0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x61, 0x70,
	0x61, 0x63, 0x69, 0x74, 0x79, 0x12, 0x4b, 0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x65, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x1a, 0x56, 0x0a, 0x14, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4d, 0x61,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x70, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x47, 0x0a, 0x19, 0x45, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69,
	0x63, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x1a, 0x43, 0x0a, 0x15, 0x49, 0x70, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x40, 0x0a, 0x12, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3e, 0x0a, 0x10, 0x41, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xcd, 0x02, 0x0a, 0x18, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x6e,
	0x73, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64,
	0x6e, 0x73, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12,
	0x33, 0x0a, 0x15, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73,
	0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14,
	0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x70, 0x65, 0x63, 0x69,
	0x66, 0x69, 0x63, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x69, 0x70, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x69, 0x70, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d,
	0x61, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x73, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x73, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x63, 0x61, 0x70,
	0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x22, 0xfb, 0x02, 0x0a, 0x18, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x09, 0x64, 0x6e,
	0x73, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x64, 0x6e, 0x73,
	0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x6d,
	0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x33, 0x0a,
	0x15, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x70, 0x65,
	0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x65, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69,
	0x63, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x69, 0x70, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x69, 0x70, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x70,
	0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x73, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x73, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x63, 0x61, 0x70, 0x61, 0x63,
	0x69, 0x74, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x22, 0x58, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74,
	0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0c, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x32, 0xd5, 0x03, 0x0a, 0x0e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x53, 0x76, 0x63, 0x12, 0x4f, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45,
	0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x2e, 0x65, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45,
	0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x13, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x3f, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x13, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x52, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x2e, 0x65,
	0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x40, 0x0a, 0x11, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x13, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4b, 0x0a,
	0x1b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x2e, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4e, 0x0a, 0x0f, 0x4c, 0x69,
	0x73, 0x74, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x2e,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x1a, 0x25, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x6f, 0x62, 0x62, 0x79, 0x66, 0x61,
	0x72, 0x6d, 0x2f, 0x67, 0x61, 0x72, 0x67, 0x61, 0x6e, 0x74, 0x75, 0x61, 0x2f, 0x76, 0x33, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x3b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_environment_environment_proto_rawDescOnce sync.Once
	file_environment_environment_proto_rawDescData []byte
)

func file_environment_environment_proto_rawDescGZIP() []byte {
	file_environment_environment_proto_rawDescOnce.Do(func() {
		file_environment_environment_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_environment_environment_proto_rawDesc), len(file_environment_environment_proto_rawDesc)))
	})
	return file_environment_environment_proto_rawDescData
}

var file_environment_environment_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_environment_environment_proto_goTypes = []any{
	(*Environment)(nil),              // 0: environment.Environment
	(*CreateEnvironmentRequest)(nil), // 1: environment.CreateEnvironmentRequest
	(*UpdateEnvironmentRequest)(nil), // 2: environment.UpdateEnvironmentRequest
	(*ListEnvironmentsResponse)(nil), // 3: environment.ListEnvironmentsResponse
	nil,                              // 4: environment.Environment.TemplateMappingEntry
	nil,                              // 5: environment.Environment.EnvironmentSpecificsEntry
	nil,                              // 6: environment.Environment.IpTranslationMapEntry
	nil,                              // 7: environment.Environment.CountCapacityEntry
	nil,                              // 8: environment.Environment.AnnotationsEntry
	(*wrapperspb.StringValue)(nil),   // 9: google.protobuf.StringValue
	(*general.StringMap)(nil),        // 10: general.StringMap
	(*general.GetRequest)(nil),       // 11: general.GetRequest
	(*general.ResourceId)(nil),       // 12: general.ResourceId
	(*general.ListOptions)(nil),      // 13: general.ListOptions
	(*emptypb.Empty)(nil),            // 14: google.protobuf.Empty
}
var file_environment_environment_proto_depIdxs = []int32{
	4,  // 0: environment.Environment.template_mapping:type_name -> environment.Environment.TemplateMappingEntry
	5,  // 1: environment.Environment.environment_specifics:type_name -> environment.Environment.EnvironmentSpecificsEntry
	6,  // 2: environment.Environment.ip_translation_map:type_name -> environment.Environment.IpTranslationMapEntry
	7,  // 3: environment.Environment.count_capacity:type_name -> environment.Environment.CountCapacityEntry
	8,  // 4: environment.Environment.annotations:type_name -> environment.Environment.AnnotationsEntry
	9,  // 5: environment.UpdateEnvironmentRequest.dnssuffix:type_name -> google.protobuf.StringValue
	0,  // 6: environment.ListEnvironmentsResponse.environments:type_name -> environment.Environment
	10, // 7: environment.Environment.TemplateMappingEntry.value:type_name -> general.StringMap
	1,  // 8: environment.EnvironmentSvc.CreateEnvironment:input_type -> environment.CreateEnvironmentRequest
	11, // 9: environment.EnvironmentSvc.GetEnvironment:input_type -> general.GetRequest
	2,  // 10: environment.EnvironmentSvc.UpdateEnvironment:input_type -> environment.UpdateEnvironmentRequest
	12, // 11: environment.EnvironmentSvc.DeleteEnvironment:input_type -> general.ResourceId
	13, // 12: environment.EnvironmentSvc.DeleteCollectionEnvironment:input_type -> general.ListOptions
	13, // 13: environment.EnvironmentSvc.ListEnvironment:input_type -> general.ListOptions
	12, // 14: environment.EnvironmentSvc.CreateEnvironment:output_type -> general.ResourceId
	0,  // 15: environment.EnvironmentSvc.GetEnvironment:output_type -> environment.Environment
	14, // 16: environment.EnvironmentSvc.UpdateEnvironment:output_type -> google.protobuf.Empty
	14, // 17: environment.EnvironmentSvc.DeleteEnvironment:output_type -> google.protobuf.Empty
	14, // 18: environment.EnvironmentSvc.DeleteCollectionEnvironment:output_type -> google.protobuf.Empty
	3,  // 19: environment.EnvironmentSvc.ListEnvironment:output_type -> environment.ListEnvironmentsResponse
	14, // [14:20] is the sub-list for method output_type
	8,  // [8:14] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_environment_environment_proto_init() }
func file_environment_environment_proto_init() {
	if File_environment_environment_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_environment_environment_proto_rawDesc), len(file_environment_environment_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_environment_environment_proto_goTypes,
		DependencyIndexes: file_environment_environment_proto_depIdxs,
		MessageInfos:      file_environment_environment_proto_msgTypes,
	}.Build()
	File_environment_environment_proto = out.File
	file_environment_environment_proto_goTypes = nil
	file_environment_environment_proto_depIdxs = nil
}
