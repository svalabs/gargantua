// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: general/general.proto

package generalpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type ResourceId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ResourceId) Reset() {
	*x = ResourceId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_general_general_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceId) ProtoMessage() {}

func (x *ResourceId) ProtoReflect() protoreflect.Message {
	mi := &file_general_general_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceId.ProtoReflect.Descriptor instead.
func (*ResourceId) Descriptor() ([]byte, []int) {
	return file_general_general_proto_rawDescGZIP(), []int{0}
}

func (x *ResourceId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	LoadFromCache bool   `protobuf:"varint,2,opt,name=load_from_cache,json=loadFromCache,proto3" json:"load_from_cache,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_general_general_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_general_general_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_general_general_proto_rawDescGZIP(), []int{1}
}

func (x *GetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetRequest) GetLoadFromCache() bool {
	if x != nil {
		return x.LoadFromCache
	}
	return false
}

type ListOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LabelSelector string `protobuf:"bytes,1,opt,name=label_selector,json=labelSelector,proto3" json:"label_selector,omitempty"`
	LoadFromCache bool   `protobuf:"varint,2,opt,name=load_from_cache,json=loadFromCache,proto3" json:"load_from_cache,omitempty"`
}

func (x *ListOptions) Reset() {
	*x = ListOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_general_general_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOptions) ProtoMessage() {}

func (x *ListOptions) ProtoReflect() protoreflect.Message {
	mi := &file_general_general_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOptions.ProtoReflect.Descriptor instead.
func (*ListOptions) Descriptor() ([]byte, []int) {
	return file_general_general_proto_rawDescGZIP(), []int{2}
}

func (x *ListOptions) GetLabelSelector() string {
	if x != nil {
		return x.LabelSelector
	}
	return ""
}

func (x *ListOptions) GetLoadFromCache() bool {
	if x != nil {
		return x.LoadFromCache
	}
	return false
}

type OwnerReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiVersion         string                `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	Kind               string                `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	Id                 string                `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	Uid                string                `protobuf:"bytes,4,opt,name=uid,proto3" json:"uid,omitempty"`
	Controller         *wrapperspb.BoolValue `protobuf:"bytes,5,opt,name=controller,proto3" json:"controller,omitempty"`
	BlockOwnerDeletion *wrapperspb.BoolValue `protobuf:"bytes,6,opt,name=block_owner_deletion,json=blockOwnerDeletion,proto3" json:"block_owner_deletion,omitempty"`
}

func (x *OwnerReference) Reset() {
	*x = OwnerReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_general_general_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OwnerReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OwnerReference) ProtoMessage() {}

func (x *OwnerReference) ProtoReflect() protoreflect.Message {
	mi := &file_general_general_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OwnerReference.ProtoReflect.Descriptor instead.
func (*OwnerReference) Descriptor() ([]byte, []int) {
	return file_general_general_proto_rawDescGZIP(), []int{3}
}

func (x *OwnerReference) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *OwnerReference) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *OwnerReference) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OwnerReference) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *OwnerReference) GetController() *wrapperspb.BoolValue {
	if x != nil {
		return x.Controller
	}
	return nil
}

func (x *OwnerReference) GetBlockOwnerDeletion() *wrapperspb.BoolValue {
	if x != nil {
		return x.BlockOwnerDeletion
	}
	return nil
}

type OwnerReferences struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerReferences []*OwnerReference `protobuf:"bytes,1,rep,name=owner_references,json=ownerReferences,proto3" json:"owner_references,omitempty"`
}

func (x *OwnerReferences) Reset() {
	*x = OwnerReferences{}
	if protoimpl.UnsafeEnabled {
		mi := &file_general_general_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OwnerReferences) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OwnerReferences) ProtoMessage() {}

func (x *OwnerReferences) ProtoReflect() protoreflect.Message {
	mi := &file_general_general_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OwnerReferences.ProtoReflect.Descriptor instead.
func (*OwnerReferences) Descriptor() ([]byte, []int) {
	return file_general_general_proto_rawDescGZIP(), []int{4}
}

func (x *OwnerReferences) GetOwnerReferences() []*OwnerReference {
	if x != nil {
		return x.OwnerReferences
	}
	return nil
}

// Since it is not possible to define an array of maps or a map of maps in protobuf,
// let's define a wrapper proto message type for string maps
type StringMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value map[string]string `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *StringMap) Reset() {
	*x = StringMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_general_general_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringMap) ProtoMessage() {}

func (x *StringMap) ProtoReflect() protoreflect.Message {
	mi := &file_general_general_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringMap.ProtoReflect.Descriptor instead.
func (*StringMap) Descriptor() ([]byte, []int) {
	return file_general_general_proto_rawDescGZIP(), []int{5}
}

func (x *StringMap) GetValue() map[string]string {
	if x != nil {
		return x.Value
	}
	return nil
}

// A wrapper for string slices in case we need to differ between an empty and unset slice
type StringArray struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *StringArray) Reset() {
	*x = StringArray{}
	if protoimpl.UnsafeEnabled {
		mi := &file_general_general_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringArray) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringArray) ProtoMessage() {}

func (x *StringArray) ProtoReflect() protoreflect.Message {
	mi := &file_general_general_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringArray.ProtoReflect.Descriptor instead.
func (*StringArray) Descriptor() ([]byte, []int) {
	return file_general_general_proto_rawDescGZIP(), []int{6}
}

func (x *StringArray) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_general_general_proto protoreflect.FileDescriptor

var file_general_general_proto_rawDesc = []byte{
	0x0a, 0x15, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x1c, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x44,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0f,
	0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x43,
	0x61, 0x63, 0x68, 0x65, 0x22, 0x5c, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x73, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x26, 0x0a, 0x0f, 0x6c, 0x6f,
	0x61, 0x64, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0d, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x63,
	0x68, 0x65, 0x22, 0xf1, 0x01, 0x0a, 0x0e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x3a, 0x0a, 0x0a,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x4c, 0x0a, 0x14, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x12, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x55, 0x0a, 0x0f, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x42, 0x0a, 0x10, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x4f, 0x77,
	0x6e, 0x65, 0x72, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x0f, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x22, 0x7a, 0x0a,
	0x09, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x70, 0x12, 0x33, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x6c, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x70, 0x2e, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x1a,
	0x38, 0x0a, 0x0a, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x25, 0x0a, 0x0b, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68,
	0x6f, 0x62, 0x62, 0x79, 0x66, 0x61, 0x72, 0x6d, 0x2f, 0x67, 0x61, 0x72, 0x67, 0x61, 0x6e, 0x74,
	0x75, 0x61, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x6c, 0x3b, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_general_general_proto_rawDescOnce sync.Once
	file_general_general_proto_rawDescData = file_general_general_proto_rawDesc
)

func file_general_general_proto_rawDescGZIP() []byte {
	file_general_general_proto_rawDescOnce.Do(func() {
		file_general_general_proto_rawDescData = protoimpl.X.CompressGZIP(file_general_general_proto_rawDescData)
	})
	return file_general_general_proto_rawDescData
}

var file_general_general_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_general_general_proto_goTypes = []interface{}{
	(*ResourceId)(nil),           // 0: general.ResourceId
	(*GetRequest)(nil),           // 1: general.GetRequest
	(*ListOptions)(nil),          // 2: general.ListOptions
	(*OwnerReference)(nil),       // 3: general.OwnerReference
	(*OwnerReferences)(nil),      // 4: general.OwnerReferences
	(*StringMap)(nil),            // 5: general.StringMap
	(*StringArray)(nil),          // 6: general.StringArray
	nil,                          // 7: general.StringMap.ValueEntry
	(*wrapperspb.BoolValue)(nil), // 8: google.protobuf.BoolValue
}
var file_general_general_proto_depIdxs = []int32{
	8, // 0: general.OwnerReference.controller:type_name -> google.protobuf.BoolValue
	8, // 1: general.OwnerReference.block_owner_deletion:type_name -> google.protobuf.BoolValue
	3, // 2: general.OwnerReferences.owner_references:type_name -> general.OwnerReference
	7, // 3: general.StringMap.value:type_name -> general.StringMap.ValueEntry
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_general_general_proto_init() }
func file_general_general_proto_init() {
	if File_general_general_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_general_general_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceId); i {
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
		file_general_general_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_general_general_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOptions); i {
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
		file_general_general_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OwnerReference); i {
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
		file_general_general_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OwnerReferences); i {
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
		file_general_general_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringMap); i {
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
		file_general_general_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringArray); i {
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
			RawDescriptor: file_general_general_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_general_general_proto_goTypes,
		DependencyIndexes: file_general_general_proto_depIdxs,
		MessageInfos:      file_general_general_proto_msgTypes,
	}.Build()
	File_general_general_proto = out.File
	file_general_general_proto_rawDesc = nil
	file_general_general_proto_goTypes = nil
	file_general_general_proto_depIdxs = nil
}
