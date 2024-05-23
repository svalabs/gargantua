// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: vmset/virtualmachineset.proto

package vmsetpb

import (
	general "github.com/hobbyfarm/gargantua/v3/protos/general"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type VMSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                  string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid                 string            `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Count               uint32            `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	Environment         string            `protobuf:"bytes,4,opt,name=environment,proto3" json:"environment,omitempty"`
	VmTemplate          string            `protobuf:"bytes,5,opt,name=vm_template,json=vmTemplate,proto3" json:"vm_template,omitempty"`
	BaseName            string            `protobuf:"bytes,6,opt,name=base_name,json=baseName,proto3" json:"base_name,omitempty"`
	RestrictedBind      bool              `protobuf:"varint,7,opt,name=restricted_bind,json=restrictedBind,proto3" json:"restricted_bind,omitempty"`
	RestrictedBindValue string            `protobuf:"bytes,8,opt,name=restricted_bind_value,json=restrictedBindValue,proto3" json:"restricted_bind_value,omitempty"`
	Status              *VMSetStatus      `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
	Labels              map[string]string `protobuf:"bytes,10,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *VMSet) Reset() {
	*x = VMSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vmset_virtualmachineset_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VMSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VMSet) ProtoMessage() {}

func (x *VMSet) ProtoReflect() protoreflect.Message {
	mi := &file_vmset_virtualmachineset_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VMSet.ProtoReflect.Descriptor instead.
func (*VMSet) Descriptor() ([]byte, []int) {
	return file_vmset_virtualmachineset_proto_rawDescGZIP(), []int{0}
}

func (x *VMSet) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *VMSet) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *VMSet) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *VMSet) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

func (x *VMSet) GetVmTemplate() string {
	if x != nil {
		return x.VmTemplate
	}
	return ""
}

func (x *VMSet) GetBaseName() string {
	if x != nil {
		return x.BaseName
	}
	return ""
}

func (x *VMSet) GetRestrictedBind() bool {
	if x != nil {
		return x.RestrictedBind
	}
	return false
}

func (x *VMSet) GetRestrictedBindValue() string {
	if x != nil {
		return x.RestrictedBindValue
	}
	return ""
}

func (x *VMSet) GetStatus() *VMSetStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *VMSet) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type CreateVMSetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                  string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Count               uint32            `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Environment         string            `protobuf:"bytes,3,opt,name=environment,proto3" json:"environment,omitempty"`
	VmTemplate          string            `protobuf:"bytes,4,opt,name=vm_template,json=vmTemplate,proto3" json:"vm_template,omitempty"`
	BaseName            string            `protobuf:"bytes,5,opt,name=base_name,json=baseName,proto3" json:"base_name,omitempty"`
	RestrictedBind      bool              `protobuf:"varint,6,opt,name=restricted_bind,json=restrictedBind,proto3" json:"restricted_bind,omitempty"`
	RestrictedBindValue string            `protobuf:"bytes,7,opt,name=restricted_bind_value,json=restrictedBindValue,proto3" json:"restricted_bind_value,omitempty"`
	SeName              string            `protobuf:"bytes,8,opt,name=se_name,json=seName,proto3" json:"se_name,omitempty"`
	SeUid               string            `protobuf:"bytes,9,opt,name=se_uid,json=seUid,proto3" json:"se_uid,omitempty"`
	Labels              map[string]string `protobuf:"bytes,10,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CreateVMSetRequest) Reset() {
	*x = CreateVMSetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vmset_virtualmachineset_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateVMSetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVMSetRequest) ProtoMessage() {}

func (x *CreateVMSetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vmset_virtualmachineset_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVMSetRequest.ProtoReflect.Descriptor instead.
func (*CreateVMSetRequest) Descriptor() ([]byte, []int) {
	return file_vmset_virtualmachineset_proto_rawDescGZIP(), []int{1}
}

func (x *CreateVMSetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateVMSetRequest) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *CreateVMSetRequest) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

func (x *CreateVMSetRequest) GetVmTemplate() string {
	if x != nil {
		return x.VmTemplate
	}
	return ""
}

func (x *CreateVMSetRequest) GetBaseName() string {
	if x != nil {
		return x.BaseName
	}
	return ""
}

func (x *CreateVMSetRequest) GetRestrictedBind() bool {
	if x != nil {
		return x.RestrictedBind
	}
	return false
}

func (x *CreateVMSetRequest) GetRestrictedBindValue() string {
	if x != nil {
		return x.RestrictedBindValue
	}
	return ""
}

func (x *CreateVMSetRequest) GetSeName() string {
	if x != nil {
		return x.SeName
	}
	return ""
}

func (x *CreateVMSetRequest) GetSeUid() string {
	if x != nil {
		return x.SeUid
	}
	return ""
}

func (x *CreateVMSetRequest) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type VMSetStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Machines    []*VMProvision `protobuf:"bytes,1,rep,name=machines,proto3" json:"machines,omitempty"`
	Available   uint32         `protobuf:"varint,2,opt,name=available,proto3" json:"available,omitempty"`
	Provisioned uint32         `protobuf:"varint,3,opt,name=provisioned,proto3" json:"provisioned,omitempty"`
}

func (x *VMSetStatus) Reset() {
	*x = VMSetStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vmset_virtualmachineset_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VMSetStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VMSetStatus) ProtoMessage() {}

func (x *VMSetStatus) ProtoReflect() protoreflect.Message {
	mi := &file_vmset_virtualmachineset_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VMSetStatus.ProtoReflect.Descriptor instead.
func (*VMSetStatus) Descriptor() ([]byte, []int) {
	return file_vmset_virtualmachineset_proto_rawDescGZIP(), []int{2}
}

func (x *VMSetStatus) GetMachines() []*VMProvision {
	if x != nil {
		return x.Machines
	}
	return nil
}

func (x *VMSetStatus) GetAvailable() uint32 {
	if x != nil {
		return x.Available
	}
	return 0
}

func (x *VMSetStatus) GetProvisioned() uint32 {
	if x != nil {
		return x.Provisioned
	}
	return 0
}

type UpdateVMSetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Count          *wrapperspb.UInt32Value `protobuf:"bytes,2,opt,name=count,proto3" json:"count,omitempty"`
	Environment    string                  `protobuf:"bytes,3,opt,name=environment,proto3" json:"environment,omitempty"`
	VmTemplate     string                  `protobuf:"bytes,4,opt,name=vm_template,json=vmTemplate,proto3" json:"vm_template,omitempty"`
	RestrictedBind *wrapperspb.BoolValue   `protobuf:"bytes,5,opt,name=restricted_bind,json=restrictedBind,proto3" json:"restricted_bind,omitempty"`
}

func (x *UpdateVMSetRequest) Reset() {
	*x = UpdateVMSetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vmset_virtualmachineset_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateVMSetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateVMSetRequest) ProtoMessage() {}

func (x *UpdateVMSetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vmset_virtualmachineset_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateVMSetRequest.ProtoReflect.Descriptor instead.
func (*UpdateVMSetRequest) Descriptor() ([]byte, []int) {
	return file_vmset_virtualmachineset_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateVMSetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateVMSetRequest) GetCount() *wrapperspb.UInt32Value {
	if x != nil {
		return x.Count
	}
	return nil
}

func (x *UpdateVMSetRequest) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

func (x *UpdateVMSetRequest) GetVmTemplate() string {
	if x != nil {
		return x.VmTemplate
	}
	return ""
}

func (x *UpdateVMSetRequest) GetRestrictedBind() *wrapperspb.BoolValue {
	if x != nil {
		return x.RestrictedBind
	}
	return nil
}

type VMProvision struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VmName   string `protobuf:"bytes,1,opt,name=vm_name,json=vmName,proto3" json:"vm_name,omitempty"`
	TfcState string `protobuf:"bytes,2,opt,name=tfc_state,json=tfcState,proto3" json:"tfc_state,omitempty"`
	TfcCm    string `protobuf:"bytes,3,opt,name=tfc_cm,json=tfcCm,proto3" json:"tfc_cm,omitempty"`
}

func (x *VMProvision) Reset() {
	*x = VMProvision{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vmset_virtualmachineset_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VMProvision) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VMProvision) ProtoMessage() {}

func (x *VMProvision) ProtoReflect() protoreflect.Message {
	mi := &file_vmset_virtualmachineset_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VMProvision.ProtoReflect.Descriptor instead.
func (*VMProvision) Descriptor() ([]byte, []int) {
	return file_vmset_virtualmachineset_proto_rawDescGZIP(), []int{4}
}

func (x *VMProvision) GetVmName() string {
	if x != nil {
		return x.VmName
	}
	return ""
}

func (x *VMProvision) GetTfcState() string {
	if x != nil {
		return x.TfcState
	}
	return ""
}

func (x *VMProvision) GetTfcCm() string {
	if x != nil {
		return x.TfcCm
	}
	return ""
}

type UpdateVMSetStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Machines    []*VMProvision          `protobuf:"bytes,2,rep,name=machines,proto3" json:"machines,omitempty"`
	Available   *wrapperspb.UInt32Value `protobuf:"bytes,3,opt,name=available,proto3" json:"available,omitempty"`
	Provisioned *wrapperspb.UInt32Value `protobuf:"bytes,4,opt,name=provisioned,proto3" json:"provisioned,omitempty"`
}

func (x *UpdateVMSetStatusRequest) Reset() {
	*x = UpdateVMSetStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vmset_virtualmachineset_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateVMSetStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateVMSetStatusRequest) ProtoMessage() {}

func (x *UpdateVMSetStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vmset_virtualmachineset_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateVMSetStatusRequest.ProtoReflect.Descriptor instead.
func (*UpdateVMSetStatusRequest) Descriptor() ([]byte, []int) {
	return file_vmset_virtualmachineset_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateVMSetStatusRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateVMSetStatusRequest) GetMachines() []*VMProvision {
	if x != nil {
		return x.Machines
	}
	return nil
}

func (x *UpdateVMSetStatusRequest) GetAvailable() *wrapperspb.UInt32Value {
	if x != nil {
		return x.Available
	}
	return nil
}

func (x *UpdateVMSetStatusRequest) GetProvisioned() *wrapperspb.UInt32Value {
	if x != nil {
		return x.Provisioned
	}
	return nil
}

type ListVMSetsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vmsets []*VMSet `protobuf:"bytes,1,rep,name=vmsets,proto3" json:"vmsets,omitempty"`
}

func (x *ListVMSetsResponse) Reset() {
	*x = ListVMSetsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vmset_virtualmachineset_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListVMSetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVMSetsResponse) ProtoMessage() {}

func (x *ListVMSetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vmset_virtualmachineset_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVMSetsResponse.ProtoReflect.Descriptor instead.
func (*ListVMSetsResponse) Descriptor() ([]byte, []int) {
	return file_vmset_virtualmachineset_proto_rawDescGZIP(), []int{6}
}

func (x *ListVMSetsResponse) GetVmsets() []*VMSet {
	if x != nil {
		return x.Vmsets
	}
	return nil
}

var File_vmset_virtualmachineset_proto protoreflect.FileDescriptor

var file_vmset_virtualmachineset_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x76, 0x6d, 0x73, 0x65, 0x74, 0x2f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x6d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x76, 0x6d, 0x73, 0x65, 0x74, 0x1a, 0x15, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2f,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x03, 0x0a, 0x05, 0x56,
	0x4d, 0x53, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b,
	0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x76, 0x6d, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x6d, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f,
	0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x65,
	0x64, 0x42, 0x69, 0x6e, 0x64, 0x12, 0x32, 0x0a, 0x15, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63,
	0x74, 0x65, 0x64, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x65, 0x64,
	0x42, 0x69, 0x6e, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x76, 0x6d, 0x73, 0x65,
	0x74, 0x2e, 0x56, 0x4d, 0x53, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x30, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18,
	0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x76, 0x6d, 0x73, 0x65, 0x74, 0x2e, 0x56, 0x4d,
	0x53, 0x65, 0x74, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0xa1, 0x03, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x4d, 0x53,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x76, 0x6d, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x6d, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x27, 0x0a, 0x0f, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x69,
	0x6e, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69,
	0x63, 0x74, 0x65, 0x64, 0x42, 0x69, 0x6e, 0x64, 0x12, 0x32, 0x0a, 0x15, 0x72, 0x65, 0x73, 0x74,
	0x72, 0x69, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63,
	0x74, 0x65, 0x64, 0x42, 0x69, 0x6e, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x17, 0x0a, 0x07,
	0x73, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x65, 0x5f, 0x75, 0x69, 0x64, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x65, 0x55, 0x69, 0x64, 0x12, 0x3d, 0x0a, 0x06,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x76,
	0x6d, 0x73, 0x65, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x4d, 0x53, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x7d, 0x0a, 0x0b, 0x56, 0x4d, 0x53, 0x65, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2e, 0x0a, 0x08, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x76, 0x6d, 0x73, 0x65, 0x74, 0x2e,
	0x56, 0x4d, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x65, 0x64, 0x22, 0xe0, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x56, 0x4d, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x32, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49,
	0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x76, 0x6d, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x6d, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x12, 0x43, 0x0a, 0x0f, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x65,
	0x64, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42,
	0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69,
	0x63, 0x74, 0x65, 0x64, 0x42, 0x69, 0x6e, 0x64, 0x22, 0x5a, 0x0a, 0x0b, 0x56, 0x4d, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x76, 0x6d, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x6d, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x74, 0x66, 0x63, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x66, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x15, 0x0a,
	0x06, 0x74, 0x66, 0x63, 0x5f, 0x63, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x66, 0x63, 0x43, 0x6d, 0x22, 0xd6, 0x01, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56,
	0x4d, 0x53, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x2e, 0x0a, 0x08, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x76, 0x6d, 0x73, 0x65, 0x74, 0x2e, 0x56, 0x4d, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x73, 0x12, 0x3a, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x3e, 0x0a,
	0x0b, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x0b, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x64, 0x22, 0x3a, 0x0a,
	0x12, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x4d, 0x53, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x76, 0x6d, 0x73, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x76, 0x6d, 0x73, 0x65, 0x74, 0x2e, 0x56, 0x4d, 0x53, 0x65,
	0x74, 0x52, 0x06, 0x76, 0x6d, 0x73, 0x65, 0x74, 0x73, 0x32, 0x8b, 0x04, 0x0a, 0x08, 0x56, 0x4d,
	0x53, 0x65, 0x74, 0x53, 0x76, 0x63, 0x12, 0x40, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x56, 0x4d, 0x53, 0x65, 0x74, 0x12, 0x19, 0x2e, 0x76, 0x6d, 0x73, 0x65, 0x74, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x56, 0x4d, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x2d, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x56,
	0x4d, 0x53, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x76, 0x6d, 0x73, 0x65,
	0x74, 0x2e, 0x56, 0x4d, 0x53, 0x65, 0x74, 0x12, 0x40, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x56, 0x4d, 0x53, 0x65, 0x74, 0x12, 0x19, 0x2e, 0x76, 0x6d, 0x73, 0x65, 0x74, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x4d, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4c, 0x0a, 0x11, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x56, 0x4d, 0x53, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f,
	0x2e, 0x76, 0x6d, 0x73, 0x65, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x4d, 0x53,
	0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3a, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x56, 0x4d, 0x53, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c,
	0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x45, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x4d, 0x53, 0x65, 0x74, 0x12, 0x14, 0x2e, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3c, 0x0a, 0x09, 0x4c, 0x69,
	0x73, 0x74, 0x56, 0x4d, 0x53, 0x65, 0x74, 0x12, 0x14, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x6c, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x19, 0x2e,
	0x76, 0x6d, 0x73, 0x65, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x4d, 0x53, 0x65, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x54,
	0x6f, 0x57, 0x6f, 0x72, 0x6b, 0x71, 0x75, 0x65, 0x75, 0x65, 0x12, 0x13, 0x2e, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x6f, 0x62, 0x62, 0x79, 0x66, 0x61, 0x72, 0x6d, 0x2f,
	0x67, 0x61, 0x72, 0x67, 0x61, 0x6e, 0x74, 0x75, 0x61, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x76, 0x6d, 0x73, 0x65, 0x74, 0x3b, 0x76, 0x6d, 0x73, 0x65, 0x74, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vmset_virtualmachineset_proto_rawDescOnce sync.Once
	file_vmset_virtualmachineset_proto_rawDescData = file_vmset_virtualmachineset_proto_rawDesc
)

func file_vmset_virtualmachineset_proto_rawDescGZIP() []byte {
	file_vmset_virtualmachineset_proto_rawDescOnce.Do(func() {
		file_vmset_virtualmachineset_proto_rawDescData = protoimpl.X.CompressGZIP(file_vmset_virtualmachineset_proto_rawDescData)
	})
	return file_vmset_virtualmachineset_proto_rawDescData
}

var file_vmset_virtualmachineset_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_vmset_virtualmachineset_proto_goTypes = []interface{}{
	(*VMSet)(nil),                    // 0: vmset.VMSet
	(*CreateVMSetRequest)(nil),       // 1: vmset.CreateVMSetRequest
	(*VMSetStatus)(nil),              // 2: vmset.VMSetStatus
	(*UpdateVMSetRequest)(nil),       // 3: vmset.UpdateVMSetRequest
	(*VMProvision)(nil),              // 4: vmset.VMProvision
	(*UpdateVMSetStatusRequest)(nil), // 5: vmset.UpdateVMSetStatusRequest
	(*ListVMSetsResponse)(nil),       // 6: vmset.ListVMSetsResponse
	nil,                              // 7: vmset.VMSet.LabelsEntry
	nil,                              // 8: vmset.CreateVMSetRequest.LabelsEntry
	(*wrapperspb.UInt32Value)(nil),   // 9: google.protobuf.UInt32Value
	(*wrapperspb.BoolValue)(nil),     // 10: google.protobuf.BoolValue
	(*general.GetRequest)(nil),       // 11: general.GetRequest
	(*general.ResourceId)(nil),       // 12: general.ResourceId
	(*general.ListOptions)(nil),      // 13: general.ListOptions
	(*emptypb.Empty)(nil),            // 14: google.protobuf.Empty
}
var file_vmset_virtualmachineset_proto_depIdxs = []int32{
	2,  // 0: vmset.VMSet.status:type_name -> vmset.VMSetStatus
	7,  // 1: vmset.VMSet.labels:type_name -> vmset.VMSet.LabelsEntry
	8,  // 2: vmset.CreateVMSetRequest.labels:type_name -> vmset.CreateVMSetRequest.LabelsEntry
	4,  // 3: vmset.VMSetStatus.machines:type_name -> vmset.VMProvision
	9,  // 4: vmset.UpdateVMSetRequest.count:type_name -> google.protobuf.UInt32Value
	10, // 5: vmset.UpdateVMSetRequest.restricted_bind:type_name -> google.protobuf.BoolValue
	4,  // 6: vmset.UpdateVMSetStatusRequest.machines:type_name -> vmset.VMProvision
	9,  // 7: vmset.UpdateVMSetStatusRequest.available:type_name -> google.protobuf.UInt32Value
	9,  // 8: vmset.UpdateVMSetStatusRequest.provisioned:type_name -> google.protobuf.UInt32Value
	0,  // 9: vmset.ListVMSetsResponse.vmsets:type_name -> vmset.VMSet
	1,  // 10: vmset.VMSetSvc.CreateVMSet:input_type -> vmset.CreateVMSetRequest
	11, // 11: vmset.VMSetSvc.GetVMSet:input_type -> general.GetRequest
	3,  // 12: vmset.VMSetSvc.UpdateVMSet:input_type -> vmset.UpdateVMSetRequest
	5,  // 13: vmset.VMSetSvc.UpdateVMSetStatus:input_type -> vmset.UpdateVMSetStatusRequest
	12, // 14: vmset.VMSetSvc.DeleteVMSet:input_type -> general.ResourceId
	13, // 15: vmset.VMSetSvc.DeleteCollectionVMSet:input_type -> general.ListOptions
	13, // 16: vmset.VMSetSvc.ListVMSet:input_type -> general.ListOptions
	12, // 17: vmset.VMSetSvc.AddToWorkqueue:input_type -> general.ResourceId
	14, // 18: vmset.VMSetSvc.CreateVMSet:output_type -> google.protobuf.Empty
	0,  // 19: vmset.VMSetSvc.GetVMSet:output_type -> vmset.VMSet
	14, // 20: vmset.VMSetSvc.UpdateVMSet:output_type -> google.protobuf.Empty
	14, // 21: vmset.VMSetSvc.UpdateVMSetStatus:output_type -> google.protobuf.Empty
	14, // 22: vmset.VMSetSvc.DeleteVMSet:output_type -> google.protobuf.Empty
	14, // 23: vmset.VMSetSvc.DeleteCollectionVMSet:output_type -> google.protobuf.Empty
	6,  // 24: vmset.VMSetSvc.ListVMSet:output_type -> vmset.ListVMSetsResponse
	14, // 25: vmset.VMSetSvc.AddToWorkqueue:output_type -> google.protobuf.Empty
	18, // [18:26] is the sub-list for method output_type
	10, // [10:18] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_vmset_virtualmachineset_proto_init() }
func file_vmset_virtualmachineset_proto_init() {
	if File_vmset_virtualmachineset_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vmset_virtualmachineset_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VMSet); i {
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
		file_vmset_virtualmachineset_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateVMSetRequest); i {
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
		file_vmset_virtualmachineset_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VMSetStatus); i {
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
		file_vmset_virtualmachineset_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateVMSetRequest); i {
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
		file_vmset_virtualmachineset_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VMProvision); i {
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
		file_vmset_virtualmachineset_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateVMSetStatusRequest); i {
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
		file_vmset_virtualmachineset_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListVMSetsResponse); i {
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
			RawDescriptor: file_vmset_virtualmachineset_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vmset_virtualmachineset_proto_goTypes,
		DependencyIndexes: file_vmset_virtualmachineset_proto_depIdxs,
		MessageInfos:      file_vmset_virtualmachineset_proto_msgTypes,
	}.Build()
	File_vmset_virtualmachineset_proto = out.File
	file_vmset_virtualmachineset_proto_rawDesc = nil
	file_vmset_virtualmachineset_proto_goTypes = nil
	file_vmset_virtualmachineset_proto_depIdxs = nil
}
