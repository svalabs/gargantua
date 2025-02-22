// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: vmtemplate/vmtemplate.proto

package vmtemplatepb

import (
	general "github.com/hobbyfarm/gargantua/v3/protos/general"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type VMTemplate struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid           string                 `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Image         string                 `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	ConfigMap     map[string]string      `protobuf:"bytes,5,rep,name=config_map,json=configMap,proto3" json:"config_map,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	CostBasePrice *string                `protobuf:"bytes,6,opt,name=cost_base_price,json=costBasePrice,proto3,oneof" json:"cost_base_price,omitempty"`
	CostTimeUnit  *string                `protobuf:"bytes,7,opt,name=cost_time_unit,json=costTimeUnit,proto3,oneof" json:"cost_time_unit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VMTemplate) Reset() {
	*x = VMTemplate{}
	mi := &file_vmtemplate_vmtemplate_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VMTemplate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VMTemplate) ProtoMessage() {}

func (x *VMTemplate) ProtoReflect() protoreflect.Message {
	mi := &file_vmtemplate_vmtemplate_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VMTemplate.ProtoReflect.Descriptor instead.
func (*VMTemplate) Descriptor() ([]byte, []int) {
	return file_vmtemplate_vmtemplate_proto_rawDescGZIP(), []int{0}
}

func (x *VMTemplate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *VMTemplate) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *VMTemplate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *VMTemplate) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *VMTemplate) GetConfigMap() map[string]string {
	if x != nil {
		return x.ConfigMap
	}
	return nil
}

func (x *VMTemplate) GetCostBasePrice() string {
	if x != nil && x.CostBasePrice != nil {
		return *x.CostBasePrice
	}
	return ""
}

func (x *VMTemplate) GetCostTimeUnit() string {
	if x != nil && x.CostTimeUnit != nil {
		return *x.CostTimeUnit
	}
	return ""
}

type CreateVMTemplateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Image         string                 `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	ConfigMapRaw  string                 `protobuf:"bytes,3,opt,name=config_map_raw,json=configMapRaw,proto3" json:"config_map_raw,omitempty"`
	CostBasePrice *string                `protobuf:"bytes,4,opt,name=cost_base_price,json=costBasePrice,proto3,oneof" json:"cost_base_price,omitempty"`
	CostTimeUnit  *string                `protobuf:"bytes,5,opt,name=cost_time_unit,json=costTimeUnit,proto3,oneof" json:"cost_time_unit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateVMTemplateRequest) Reset() {
	*x = CreateVMTemplateRequest{}
	mi := &file_vmtemplate_vmtemplate_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateVMTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVMTemplateRequest) ProtoMessage() {}

func (x *CreateVMTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vmtemplate_vmtemplate_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVMTemplateRequest.ProtoReflect.Descriptor instead.
func (*CreateVMTemplateRequest) Descriptor() ([]byte, []int) {
	return file_vmtemplate_vmtemplate_proto_rawDescGZIP(), []int{1}
}

func (x *CreateVMTemplateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateVMTemplateRequest) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *CreateVMTemplateRequest) GetConfigMapRaw() string {
	if x != nil {
		return x.ConfigMapRaw
	}
	return ""
}

func (x *CreateVMTemplateRequest) GetCostBasePrice() string {
	if x != nil && x.CostBasePrice != nil {
		return *x.CostBasePrice
	}
	return ""
}

func (x *CreateVMTemplateRequest) GetCostTimeUnit() string {
	if x != nil && x.CostTimeUnit != nil {
		return *x.CostTimeUnit
	}
	return ""
}

type UpdateVMTemplateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Image         string                 `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	ConfigMapRaw  string                 `protobuf:"bytes,4,opt,name=config_map_raw,json=configMapRaw,proto3" json:"config_map_raw,omitempty"`
	CostBasePrice *string                `protobuf:"bytes,5,opt,name=cost_base_price,json=costBasePrice,proto3,oneof" json:"cost_base_price,omitempty"`
	CostTimeUnit  *string                `protobuf:"bytes,6,opt,name=cost_time_unit,json=costTimeUnit,proto3,oneof" json:"cost_time_unit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateVMTemplateRequest) Reset() {
	*x = UpdateVMTemplateRequest{}
	mi := &file_vmtemplate_vmtemplate_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateVMTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateVMTemplateRequest) ProtoMessage() {}

func (x *UpdateVMTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vmtemplate_vmtemplate_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateVMTemplateRequest.ProtoReflect.Descriptor instead.
func (*UpdateVMTemplateRequest) Descriptor() ([]byte, []int) {
	return file_vmtemplate_vmtemplate_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateVMTemplateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateVMTemplateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateVMTemplateRequest) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *UpdateVMTemplateRequest) GetConfigMapRaw() string {
	if x != nil {
		return x.ConfigMapRaw
	}
	return ""
}

func (x *UpdateVMTemplateRequest) GetCostBasePrice() string {
	if x != nil && x.CostBasePrice != nil {
		return *x.CostBasePrice
	}
	return ""
}

func (x *UpdateVMTemplateRequest) GetCostTimeUnit() string {
	if x != nil && x.CostTimeUnit != nil {
		return *x.CostTimeUnit
	}
	return ""
}

type ListVMTemplatesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Vmtemplates   []*VMTemplate          `protobuf:"bytes,1,rep,name=vmtemplates,proto3" json:"vmtemplates,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListVMTemplatesResponse) Reset() {
	*x = ListVMTemplatesResponse{}
	mi := &file_vmtemplate_vmtemplate_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListVMTemplatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVMTemplatesResponse) ProtoMessage() {}

func (x *ListVMTemplatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vmtemplate_vmtemplate_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVMTemplatesResponse.ProtoReflect.Descriptor instead.
func (*ListVMTemplatesResponse) Descriptor() ([]byte, []int) {
	return file_vmtemplate_vmtemplate_proto_rawDescGZIP(), []int{3}
}

func (x *ListVMTemplatesResponse) GetVmtemplates() []*VMTemplate {
	if x != nil {
		return x.Vmtemplates
	}
	return nil
}

var File_vmtemplate_vmtemplate_proto protoreflect.FileDescriptor

var file_vmtemplate_vmtemplate_proto_rawDesc = string([]byte{
	0x0a, 0x1b, 0x76, 0x6d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x6d, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x76,
	0x6d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x1a, 0x15, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdb, 0x02,
	0x0a, 0x0a, 0x56, 0x4d, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x44, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x76,
	0x6d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x56, 0x4d, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4d, 0x61, 0x70, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4d, 0x61, 0x70, 0x12, 0x2b,
	0x0a, 0x0f, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0d, 0x63, 0x6f, 0x73, 0x74, 0x42,
	0x61, 0x73, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x0e, 0x63,
	0x6f, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0c, 0x63, 0x6f, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x55,
	0x6e, 0x69, 0x74, 0x88, 0x01, 0x01, 0x1a, 0x3c, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x62, 0x61,
	0x73, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x63, 0x6f, 0x73,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x22, 0xe8, 0x01, 0x0a, 0x17,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x4d, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x6d, 0x61, 0x70, 0x5f,
	0x72, 0x61, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x4d, 0x61, 0x70, 0x52, 0x61, 0x77, 0x12, 0x2b, 0x0a, 0x0f, 0x63, 0x6f, 0x73, 0x74, 0x5f,
	0x62, 0x61, 0x73, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x0d, 0x63, 0x6f, 0x73, 0x74, 0x42, 0x61, 0x73, 0x65, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x0e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0c,
	0x63, 0x6f, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x88, 0x01, 0x01, 0x42,
	0x12, 0x0a, 0x10, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x22, 0xf8, 0x01, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x56, 0x4d, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x0e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x6d, 0x61, 0x70, 0x5f, 0x72, 0x61, 0x77, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4d, 0x61, 0x70, 0x52,
	0x61, 0x77, 0x12, 0x2b, 0x0a, 0x0f, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0d, 0x63,
	0x6f, 0x73, 0x74, 0x42, 0x61, 0x73, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x29, 0x0a, 0x0e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x6e, 0x69,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0c, 0x63, 0x6f, 0x73, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x88, 0x01, 0x01, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x63,
	0x6f, 0x73, 0x74, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x42, 0x11,
	0x0a, 0x0f, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x6e, 0x69,
	0x74, 0x22, 0x53, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x4d, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0b,
	0x76, 0x6d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x76, 0x6d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x56,
	0x4d, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x0b, 0x76, 0x6d, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x32, 0xc6, 0x03, 0x0a, 0x0d, 0x56, 0x4d, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x76, 0x63, 0x12, 0x4c, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x56, 0x4d, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x76,
	0x6d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x56, 0x4d, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x13, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x3c, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x56, 0x4d, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x6c, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x76,
	0x6d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x56, 0x4d, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x12, 0x4f, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x4d,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x76, 0x6d, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x4d, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3f, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56,
	0x4d, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4a, 0x0a, 0x1a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x4d, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x4b, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x4d, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x23, 0x2e, 0x76, 0x6d, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x4d, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x6f,
	0x62, 0x62, 0x79, 0x66, 0x61, 0x72, 0x6d, 0x2f, 0x67, 0x61, 0x72, 0x67, 0x61, 0x6e, 0x74, 0x75,
	0x61, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x6d, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x3b, 0x76, 0x6d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_vmtemplate_vmtemplate_proto_rawDescOnce sync.Once
	file_vmtemplate_vmtemplate_proto_rawDescData []byte
)

func file_vmtemplate_vmtemplate_proto_rawDescGZIP() []byte {
	file_vmtemplate_vmtemplate_proto_rawDescOnce.Do(func() {
		file_vmtemplate_vmtemplate_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_vmtemplate_vmtemplate_proto_rawDesc), len(file_vmtemplate_vmtemplate_proto_rawDesc)))
	})
	return file_vmtemplate_vmtemplate_proto_rawDescData
}

var file_vmtemplate_vmtemplate_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_vmtemplate_vmtemplate_proto_goTypes = []any{
	(*VMTemplate)(nil),              // 0: vmtemplate.VMTemplate
	(*CreateVMTemplateRequest)(nil), // 1: vmtemplate.CreateVMTemplateRequest
	(*UpdateVMTemplateRequest)(nil), // 2: vmtemplate.UpdateVMTemplateRequest
	(*ListVMTemplatesResponse)(nil), // 3: vmtemplate.ListVMTemplatesResponse
	nil,                             // 4: vmtemplate.VMTemplate.ConfigMapEntry
	(*general.GetRequest)(nil),      // 5: general.GetRequest
	(*general.ResourceId)(nil),      // 6: general.ResourceId
	(*general.ListOptions)(nil),     // 7: general.ListOptions
	(*emptypb.Empty)(nil),           // 8: google.protobuf.Empty
}
var file_vmtemplate_vmtemplate_proto_depIdxs = []int32{
	4, // 0: vmtemplate.VMTemplate.config_map:type_name -> vmtemplate.VMTemplate.ConfigMapEntry
	0, // 1: vmtemplate.ListVMTemplatesResponse.vmtemplates:type_name -> vmtemplate.VMTemplate
	1, // 2: vmtemplate.VMTemplateSvc.CreateVMTemplate:input_type -> vmtemplate.CreateVMTemplateRequest
	5, // 3: vmtemplate.VMTemplateSvc.GetVMTemplate:input_type -> general.GetRequest
	2, // 4: vmtemplate.VMTemplateSvc.UpdateVMTemplate:input_type -> vmtemplate.UpdateVMTemplateRequest
	6, // 5: vmtemplate.VMTemplateSvc.DeleteVMTemplate:input_type -> general.ResourceId
	7, // 6: vmtemplate.VMTemplateSvc.DeleteCollectionVMTemplate:input_type -> general.ListOptions
	7, // 7: vmtemplate.VMTemplateSvc.ListVMTemplate:input_type -> general.ListOptions
	6, // 8: vmtemplate.VMTemplateSvc.CreateVMTemplate:output_type -> general.ResourceId
	0, // 9: vmtemplate.VMTemplateSvc.GetVMTemplate:output_type -> vmtemplate.VMTemplate
	8, // 10: vmtemplate.VMTemplateSvc.UpdateVMTemplate:output_type -> google.protobuf.Empty
	8, // 11: vmtemplate.VMTemplateSvc.DeleteVMTemplate:output_type -> google.protobuf.Empty
	8, // 12: vmtemplate.VMTemplateSvc.DeleteCollectionVMTemplate:output_type -> google.protobuf.Empty
	3, // 13: vmtemplate.VMTemplateSvc.ListVMTemplate:output_type -> vmtemplate.ListVMTemplatesResponse
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_vmtemplate_vmtemplate_proto_init() }
func file_vmtemplate_vmtemplate_proto_init() {
	if File_vmtemplate_vmtemplate_proto != nil {
		return
	}
	file_vmtemplate_vmtemplate_proto_msgTypes[0].OneofWrappers = []any{}
	file_vmtemplate_vmtemplate_proto_msgTypes[1].OneofWrappers = []any{}
	file_vmtemplate_vmtemplate_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_vmtemplate_vmtemplate_proto_rawDesc), len(file_vmtemplate_vmtemplate_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vmtemplate_vmtemplate_proto_goTypes,
		DependencyIndexes: file_vmtemplate_vmtemplate_proto_depIdxs,
		MessageInfos:      file_vmtemplate_vmtemplate_proto_msgTypes,
	}.Build()
	File_vmtemplate_vmtemplate_proto = out.File
	file_vmtemplate_vmtemplate_proto_goTypes = nil
	file_vmtemplate_vmtemplate_proto_depIdxs = nil
}
