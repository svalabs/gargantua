// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: vm/vm.proto

package vmpb

import (
	general "github.com/hobbyfarm/gargantua/v3/protos/general"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type VM struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	Id                string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid               string                 `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	VmTemplateId      string                 `protobuf:"bytes,3,opt,name=vm_template_id,json=vmTemplateId,proto3" json:"vm_template_id,omitempty"`
	SshUsername       string                 `protobuf:"bytes,4,opt,name=ssh_username,json=sshUsername,proto3" json:"ssh_username,omitempty"`
	Protocol          string                 `protobuf:"bytes,5,opt,name=protocol,proto3" json:"protocol,omitempty"`
	SecretName        string                 `protobuf:"bytes,6,opt,name=secret_name,json=secretName,proto3" json:"secret_name,omitempty"`
	VmClaimId         string                 `protobuf:"bytes,7,opt,name=vm_claim_id,json=vmClaimId,proto3" json:"vm_claim_id,omitempty"`
	User              string                 `protobuf:"bytes,8,opt,name=user,proto3" json:"user,omitempty"`
	Provision         bool                   `protobuf:"varint,9,opt,name=provision,proto3" json:"provision,omitempty"`
	VmSetId           string                 `protobuf:"bytes,10,opt,name=vm_set_id,json=vmSetId,proto3" json:"vm_set_id,omitempty"`
	Labels            map[string]string      `protobuf:"bytes,11,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Finalizers        []string               `protobuf:"bytes,12,rep,name=finalizers,proto3" json:"finalizers,omitempty"`
	Status            *VMStatus              `protobuf:"bytes,13,opt,name=status,proto3" json:"status,omitempty"`
	Annotations       map[string]string      `protobuf:"bytes,14,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	DeletionTimestamp *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=deletion_timestamp,json=deletionTimestamp,proto3" json:"deletion_timestamp,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *VM) Reset() {
	*x = VM{}
	mi := &file_vm_vm_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VM) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VM) ProtoMessage() {}

func (x *VM) ProtoReflect() protoreflect.Message {
	mi := &file_vm_vm_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VM.ProtoReflect.Descriptor instead.
func (*VM) Descriptor() ([]byte, []int) {
	return file_vm_vm_proto_rawDescGZIP(), []int{0}
}

func (x *VM) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *VM) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *VM) GetVmTemplateId() string {
	if x != nil {
		return x.VmTemplateId
	}
	return ""
}

func (x *VM) GetSshUsername() string {
	if x != nil {
		return x.SshUsername
	}
	return ""
}

func (x *VM) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *VM) GetSecretName() string {
	if x != nil {
		return x.SecretName
	}
	return ""
}

func (x *VM) GetVmClaimId() string {
	if x != nil {
		return x.VmClaimId
	}
	return ""
}

func (x *VM) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *VM) GetProvision() bool {
	if x != nil {
		return x.Provision
	}
	return false
}

func (x *VM) GetVmSetId() string {
	if x != nil {
		return x.VmSetId
	}
	return ""
}

func (x *VM) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *VM) GetFinalizers() []string {
	if x != nil {
		return x.Finalizers
	}
	return nil
}

func (x *VM) GetStatus() *VMStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *VM) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

func (x *VM) GetDeletionTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletionTimestamp
	}
	return nil
}

type CreateVMRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	VmTemplateId  string                 `protobuf:"bytes,2,opt,name=vm_template_id,json=vmTemplateId,proto3" json:"vm_template_id,omitempty"`
	SshUsername   string                 `protobuf:"bytes,3,opt,name=ssh_username,json=sshUsername,proto3" json:"ssh_username,omitempty"`
	Protocol      string                 `protobuf:"bytes,4,opt,name=protocol,proto3" json:"protocol,omitempty"`
	SecretName    string                 `protobuf:"bytes,5,opt,name=secret_name,json=secretName,proto3" json:"secret_name,omitempty"`
	VmClaimId     string                 `protobuf:"bytes,6,opt,name=vm_claim_id,json=vmClaimId,proto3" json:"vm_claim_id,omitempty"`
	VmClaimUid    string                 `protobuf:"bytes,7,opt,name=vm_claim_uid,json=vmClaimUid,proto3" json:"vm_claim_uid,omitempty"`
	User          string                 `protobuf:"bytes,8,opt,name=user,proto3" json:"user,omitempty"`
	Provision     bool                   `protobuf:"varint,9,opt,name=provision,proto3" json:"provision,omitempty"`
	VmSetId       string                 `protobuf:"bytes,10,opt,name=vm_set_id,json=vmSetId,proto3" json:"vm_set_id,omitempty"`
	VmSetUid      string                 `protobuf:"bytes,11,opt,name=vm_set_uid,json=vmSetUid,proto3" json:"vm_set_uid,omitempty"`
	Labels        map[string]string      `protobuf:"bytes,12,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Finalizers    []string               `protobuf:"bytes,13,rep,name=finalizers,proto3" json:"finalizers,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateVMRequest) Reset() {
	*x = CreateVMRequest{}
	mi := &file_vm_vm_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateVMRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVMRequest) ProtoMessage() {}

func (x *CreateVMRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vm_vm_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVMRequest.ProtoReflect.Descriptor instead.
func (*CreateVMRequest) Descriptor() ([]byte, []int) {
	return file_vm_vm_proto_rawDescGZIP(), []int{1}
}

func (x *CreateVMRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateVMRequest) GetVmTemplateId() string {
	if x != nil {
		return x.VmTemplateId
	}
	return ""
}

func (x *CreateVMRequest) GetSshUsername() string {
	if x != nil {
		return x.SshUsername
	}
	return ""
}

func (x *CreateVMRequest) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *CreateVMRequest) GetSecretName() string {
	if x != nil {
		return x.SecretName
	}
	return ""
}

func (x *CreateVMRequest) GetVmClaimId() string {
	if x != nil {
		return x.VmClaimId
	}
	return ""
}

func (x *CreateVMRequest) GetVmClaimUid() string {
	if x != nil {
		return x.VmClaimUid
	}
	return ""
}

func (x *CreateVMRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *CreateVMRequest) GetProvision() bool {
	if x != nil {
		return x.Provision
	}
	return false
}

func (x *CreateVMRequest) GetVmSetId() string {
	if x != nil {
		return x.VmSetId
	}
	return ""
}

func (x *CreateVMRequest) GetVmSetUid() string {
	if x != nil {
		return x.VmSetUid
	}
	return ""
}

func (x *CreateVMRequest) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *CreateVMRequest) GetFinalizers() []string {
	if x != nil {
		return x.Finalizers
	}
	return nil
}

type UpdateVMRequest struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Id            string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Bound         string                  `protobuf:"bytes,2,opt,name=bound,proto3" json:"bound,omitempty"`
	VmClaimId     *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=vm_claim_id,json=vmClaimId,proto3" json:"vm_claim_id,omitempty"`
	User          *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	SecretName    string                  `protobuf:"bytes,5,opt,name=secret_name,json=secretName,proto3" json:"secret_name,omitempty"`
	Finalizers    *general.StringArray    `protobuf:"bytes,6,opt,name=finalizers,proto3" json:"finalizers,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateVMRequest) Reset() {
	*x = UpdateVMRequest{}
	mi := &file_vm_vm_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateVMRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateVMRequest) ProtoMessage() {}

func (x *UpdateVMRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vm_vm_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateVMRequest.ProtoReflect.Descriptor instead.
func (*UpdateVMRequest) Descriptor() ([]byte, []int) {
	return file_vm_vm_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateVMRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateVMRequest) GetBound() string {
	if x != nil {
		return x.Bound
	}
	return ""
}

func (x *UpdateVMRequest) GetVmClaimId() *wrapperspb.StringValue {
	if x != nil {
		return x.VmClaimId
	}
	return nil
}

func (x *UpdateVMRequest) GetUser() *wrapperspb.StringValue {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UpdateVMRequest) GetSecretName() string {
	if x != nil {
		return x.SecretName
	}
	return ""
}

func (x *UpdateVMRequest) GetFinalizers() *general.StringArray {
	if x != nil {
		return x.Finalizers
	}
	return nil
}

type UpdateVMStatusRequest struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Id            string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status        string                  `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Allocated     *wrapperspb.BoolValue   `protobuf:"bytes,3,opt,name=allocated,proto3" json:"allocated,omitempty"`
	Tainted       *wrapperspb.BoolValue   `protobuf:"bytes,4,opt,name=tainted,proto3" json:"tainted,omitempty"`
	PublicIp      *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=public_ip,json=publicIp,proto3" json:"public_ip,omitempty"`
	PrivateIp     *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=private_ip,json=privateIp,proto3" json:"private_ip,omitempty"`
	Hostname      *wrapperspb.StringValue `protobuf:"bytes,7,opt,name=hostname,proto3" json:"hostname,omitempty"`
	EnvironmentId string                  `protobuf:"bytes,8,opt,name=environment_id,json=environmentId,proto3" json:"environment_id,omitempty"`
	Tfstate       string                  `protobuf:"bytes,9,opt,name=tfstate,proto3" json:"tfstate,omitempty"`
	WsEndpoint    string                  `protobuf:"bytes,10,opt,name=ws_endpoint,json=wsEndpoint,proto3" json:"ws_endpoint,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateVMStatusRequest) Reset() {
	*x = UpdateVMStatusRequest{}
	mi := &file_vm_vm_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateVMStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateVMStatusRequest) ProtoMessage() {}

func (x *UpdateVMStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vm_vm_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateVMStatusRequest.ProtoReflect.Descriptor instead.
func (*UpdateVMStatusRequest) Descriptor() ([]byte, []int) {
	return file_vm_vm_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateVMStatusRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateVMStatusRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *UpdateVMStatusRequest) GetAllocated() *wrapperspb.BoolValue {
	if x != nil {
		return x.Allocated
	}
	return nil
}

func (x *UpdateVMStatusRequest) GetTainted() *wrapperspb.BoolValue {
	if x != nil {
		return x.Tainted
	}
	return nil
}

func (x *UpdateVMStatusRequest) GetPublicIp() *wrapperspb.StringValue {
	if x != nil {
		return x.PublicIp
	}
	return nil
}

func (x *UpdateVMStatusRequest) GetPrivateIp() *wrapperspb.StringValue {
	if x != nil {
		return x.PrivateIp
	}
	return nil
}

func (x *UpdateVMStatusRequest) GetHostname() *wrapperspb.StringValue {
	if x != nil {
		return x.Hostname
	}
	return nil
}

func (x *UpdateVMStatusRequest) GetEnvironmentId() string {
	if x != nil {
		return x.EnvironmentId
	}
	return ""
}

func (x *UpdateVMStatusRequest) GetTfstate() string {
	if x != nil {
		return x.Tfstate
	}
	return ""
}

func (x *UpdateVMStatusRequest) GetWsEndpoint() string {
	if x != nil {
		return x.WsEndpoint
	}
	return ""
}

type VMStatus struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Allocated     bool                   `protobuf:"varint,2,opt,name=allocated,proto3" json:"allocated,omitempty"`
	Tainted       bool                   `protobuf:"varint,3,opt,name=tainted,proto3" json:"tainted,omitempty"`
	PublicIp      string                 `protobuf:"bytes,4,opt,name=public_ip,json=publicIp,proto3" json:"public_ip,omitempty"`
	PrivateIp     string                 `protobuf:"bytes,5,opt,name=private_ip,json=privateIp,proto3" json:"private_ip,omitempty"`
	Hostname      string                 `protobuf:"bytes,6,opt,name=hostname,proto3" json:"hostname,omitempty"`
	EnvironmentId string                 `protobuf:"bytes,7,opt,name=environment_id,json=environmentId,proto3" json:"environment_id,omitempty"`
	Tfstate       string                 `protobuf:"bytes,8,opt,name=tfstate,proto3" json:"tfstate,omitempty"`
	WsEndpoint    string                 `protobuf:"bytes,9,opt,name=ws_endpoint,json=wsEndpoint,proto3" json:"ws_endpoint,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VMStatus) Reset() {
	*x = VMStatus{}
	mi := &file_vm_vm_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VMStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VMStatus) ProtoMessage() {}

func (x *VMStatus) ProtoReflect() protoreflect.Message {
	mi := &file_vm_vm_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VMStatus.ProtoReflect.Descriptor instead.
func (*VMStatus) Descriptor() ([]byte, []int) {
	return file_vm_vm_proto_rawDescGZIP(), []int{4}
}

func (x *VMStatus) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *VMStatus) GetAllocated() bool {
	if x != nil {
		return x.Allocated
	}
	return false
}

func (x *VMStatus) GetTainted() bool {
	if x != nil {
		return x.Tainted
	}
	return false
}

func (x *VMStatus) GetPublicIp() string {
	if x != nil {
		return x.PublicIp
	}
	return ""
}

func (x *VMStatus) GetPrivateIp() string {
	if x != nil {
		return x.PrivateIp
	}
	return ""
}

func (x *VMStatus) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *VMStatus) GetEnvironmentId() string {
	if x != nil {
		return x.EnvironmentId
	}
	return ""
}

func (x *VMStatus) GetTfstate() string {
	if x != nil {
		return x.Tfstate
	}
	return ""
}

func (x *VMStatus) GetWsEndpoint() string {
	if x != nil {
		return x.WsEndpoint
	}
	return ""
}

type ListVMsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Vms           []*VM                  `protobuf:"bytes,1,rep,name=vms,proto3" json:"vms,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListVMsResponse) Reset() {
	*x = ListVMsResponse{}
	mi := &file_vm_vm_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListVMsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVMsResponse) ProtoMessage() {}

func (x *ListVMsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vm_vm_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVMsResponse.ProtoReflect.Descriptor instead.
func (*ListVMsResponse) Descriptor() ([]byte, []int) {
	return file_vm_vm_proto_rawDescGZIP(), []int{5}
}

func (x *ListVMsResponse) GetVms() []*VM {
	if x != nil {
		return x.Vms
	}
	return nil
}

var File_vm_vm_proto protoreflect.FileDescriptor

var file_vm_vm_proto_rawDesc = string([]byte{
	0x0a, 0x0b, 0x76, 0x6d, 0x2f, 0x76, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76,
	0x6d, 0x1a, 0x15, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x05, 0x0a, 0x02, 0x56, 0x4d, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12,
	0x24, 0x0a, 0x0e, 0x76, 0x6d, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x76, 0x6d, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x73, 0x68, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x73, 0x68,
	0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0b, 0x76, 0x6d, 0x5f, 0x63, 0x6c, 0x61, 0x69,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x6d, 0x43, 0x6c,
	0x61, 0x69, 0x6d, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x09, 0x76, 0x6d, 0x5f, 0x73, 0x65,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x6d, 0x53, 0x65,
	0x74, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x0b, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x76, 0x6d, 0x2e, 0x56, 0x4d, 0x2e, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12,
	0x1e, 0x0a, 0x0a, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x73, 0x18, 0x0c, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x73, 0x12,
	0x24, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x76, 0x6d, 0x2e, 0x56, 0x4d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x76, 0x6d, 0x2e,
	0x56, 0x4d, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x49, 0x0a, 0x12, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x11, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x1a, 0x39, 0x0a, 0x0b, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3e, 0x0a, 0x10, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xe9, 0x03, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x56, 0x4d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x76, 0x6d,
	0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x76, 0x6d, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x73, 0x73, 0x68, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x73, 0x68, 0x55, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1e, 0x0a, 0x0b, 0x76, 0x6d, 0x5f, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x5f, 0x69, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x6d, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x49, 0x64,
	0x12, 0x20, 0x0a, 0x0c, 0x76, 0x6d, 0x5f, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x5f, 0x75, 0x69, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x6d, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x55,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x09, 0x76, 0x6d, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x6d, 0x53, 0x65, 0x74, 0x49, 0x64,
	0x12, 0x1c, 0x0a, 0x0a, 0x76, 0x6d, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76, 0x6d, 0x53, 0x65, 0x74, 0x55, 0x69, 0x64, 0x12, 0x37,
	0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x76, 0x6d, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x4d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x69, 0x6e, 0x61, 0x6c,
	0x69, 0x7a, 0x65, 0x72, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x69, 0x6e,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0xfe, 0x01, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x4d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x3c, 0x0a, 0x0b,
	0x76, 0x6d, 0x5f, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x09, 0x76, 0x6d, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a,
	0x0a, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x41, 0x72, 0x72, 0x61, 0x79, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a,
	0x65, 0x72, 0x73, 0x22, 0xc3, 0x03, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x4d,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x38, 0x0a, 0x09, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x34, 0x0a, 0x07, 0x74, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x74, 0x61,
	0x69, 0x6e, 0x74, 0x65, 0x64, 0x12, 0x39, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f,
	0x69, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x70,
	0x12, 0x3b, 0x0a, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x70, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x09, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x49, 0x70, 0x12, 0x38, 0x0a,
	0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x68,
	0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x74, 0x66, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x74, 0x66, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x73, 0x5f, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77,
	0x73, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x94, 0x02, 0x0a, 0x08, 0x56, 0x4d,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x74, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x74,
	0x61, 0x69, 0x6e, 0x74, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x5f, 0x69, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x49, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x69,
	0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x49, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25,
	0x0a, 0x0e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x66, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x66, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x77, 0x73, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x73, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x22, 0x2b, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x4d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x03, 0x76, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x06, 0x2e, 0x76, 0x6d, 0x2e, 0x56, 0x4d, 0x52, 0x03, 0x76, 0x6d, 0x73, 0x32, 0x96, 0x03,
	0x0a, 0x05, 0x56, 0x4d, 0x53, 0x76, 0x63, 0x12, 0x37, 0x0a, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x56, 0x4d, 0x12, 0x13, 0x2e, 0x76, 0x6d, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56,
	0x4d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x24, 0x0a, 0x05, 0x47, 0x65, 0x74, 0x56, 0x4d, 0x12, 0x13, 0x2e, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x06,
	0x2e, 0x76, 0x6d, 0x2e, 0x56, 0x4d, 0x12, 0x37, 0x0a, 0x08, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x56, 0x4d, 0x12, 0x13, 0x2e, 0x76, 0x6d, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x4d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x43, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x4d, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x19, 0x2e, 0x76, 0x6d, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x4d, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x37, 0x0a, 0x08, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x4d,
	0x12, 0x13, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x42, 0x0a,
	0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x56, 0x4d, 0x12, 0x14, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x33, 0x0a, 0x06, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x4d, 0x12, 0x14, 0x2e, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x1a, 0x13, 0x2e, 0x76, 0x6d, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x4d, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x6f, 0x62, 0x62, 0x79, 0x66, 0x61, 0x72, 0x6d, 0x2f, 0x67,
	0x61, 0x72, 0x67, 0x61, 0x6e, 0x74, 0x75, 0x61, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x76, 0x6d, 0x3b, 0x76, 0x6d, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_vm_vm_proto_rawDescOnce sync.Once
	file_vm_vm_proto_rawDescData []byte
)

func file_vm_vm_proto_rawDescGZIP() []byte {
	file_vm_vm_proto_rawDescOnce.Do(func() {
		file_vm_vm_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_vm_vm_proto_rawDesc), len(file_vm_vm_proto_rawDesc)))
	})
	return file_vm_vm_proto_rawDescData
}

var file_vm_vm_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_vm_vm_proto_goTypes = []any{
	(*VM)(nil),                     // 0: vm.VM
	(*CreateVMRequest)(nil),        // 1: vm.CreateVMRequest
	(*UpdateVMRequest)(nil),        // 2: vm.UpdateVMRequest
	(*UpdateVMStatusRequest)(nil),  // 3: vm.UpdateVMStatusRequest
	(*VMStatus)(nil),               // 4: vm.VMStatus
	(*ListVMsResponse)(nil),        // 5: vm.ListVMsResponse
	nil,                            // 6: vm.VM.LabelsEntry
	nil,                            // 7: vm.VM.AnnotationsEntry
	nil,                            // 8: vm.CreateVMRequest.LabelsEntry
	(*timestamppb.Timestamp)(nil),  // 9: google.protobuf.Timestamp
	(*wrapperspb.StringValue)(nil), // 10: google.protobuf.StringValue
	(*general.StringArray)(nil),    // 11: general.StringArray
	(*wrapperspb.BoolValue)(nil),   // 12: google.protobuf.BoolValue
	(*general.GetRequest)(nil),     // 13: general.GetRequest
	(*general.ResourceId)(nil),     // 14: general.ResourceId
	(*general.ListOptions)(nil),    // 15: general.ListOptions
	(*emptypb.Empty)(nil),          // 16: google.protobuf.Empty
}
var file_vm_vm_proto_depIdxs = []int32{
	6,  // 0: vm.VM.labels:type_name -> vm.VM.LabelsEntry
	4,  // 1: vm.VM.status:type_name -> vm.VMStatus
	7,  // 2: vm.VM.annotations:type_name -> vm.VM.AnnotationsEntry
	9,  // 3: vm.VM.deletion_timestamp:type_name -> google.protobuf.Timestamp
	8,  // 4: vm.CreateVMRequest.labels:type_name -> vm.CreateVMRequest.LabelsEntry
	10, // 5: vm.UpdateVMRequest.vm_claim_id:type_name -> google.protobuf.StringValue
	10, // 6: vm.UpdateVMRequest.user:type_name -> google.protobuf.StringValue
	11, // 7: vm.UpdateVMRequest.finalizers:type_name -> general.StringArray
	12, // 8: vm.UpdateVMStatusRequest.allocated:type_name -> google.protobuf.BoolValue
	12, // 9: vm.UpdateVMStatusRequest.tainted:type_name -> google.protobuf.BoolValue
	10, // 10: vm.UpdateVMStatusRequest.public_ip:type_name -> google.protobuf.StringValue
	10, // 11: vm.UpdateVMStatusRequest.private_ip:type_name -> google.protobuf.StringValue
	10, // 12: vm.UpdateVMStatusRequest.hostname:type_name -> google.protobuf.StringValue
	0,  // 13: vm.ListVMsResponse.vms:type_name -> vm.VM
	1,  // 14: vm.VMSvc.CreateVM:input_type -> vm.CreateVMRequest
	13, // 15: vm.VMSvc.GetVM:input_type -> general.GetRequest
	2,  // 16: vm.VMSvc.UpdateVM:input_type -> vm.UpdateVMRequest
	3,  // 17: vm.VMSvc.UpdateVMStatus:input_type -> vm.UpdateVMStatusRequest
	14, // 18: vm.VMSvc.DeleteVM:input_type -> general.ResourceId
	15, // 19: vm.VMSvc.DeleteCollectionVM:input_type -> general.ListOptions
	15, // 20: vm.VMSvc.ListVM:input_type -> general.ListOptions
	16, // 21: vm.VMSvc.CreateVM:output_type -> google.protobuf.Empty
	0,  // 22: vm.VMSvc.GetVM:output_type -> vm.VM
	16, // 23: vm.VMSvc.UpdateVM:output_type -> google.protobuf.Empty
	16, // 24: vm.VMSvc.UpdateVMStatus:output_type -> google.protobuf.Empty
	16, // 25: vm.VMSvc.DeleteVM:output_type -> google.protobuf.Empty
	16, // 26: vm.VMSvc.DeleteCollectionVM:output_type -> google.protobuf.Empty
	5,  // 27: vm.VMSvc.ListVM:output_type -> vm.ListVMsResponse
	21, // [21:28] is the sub-list for method output_type
	14, // [14:21] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_vm_vm_proto_init() }
func file_vm_vm_proto_init() {
	if File_vm_vm_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_vm_vm_proto_rawDesc), len(file_vm_vm_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vm_vm_proto_goTypes,
		DependencyIndexes: file_vm_vm_proto_depIdxs,
		MessageInfos:      file_vm_vm_proto_msgTypes,
	}.Build()
	File_vm_vm_proto = out.File
	file_vm_vm_proto_goTypes = nil
	file_vm_vm_proto_depIdxs = nil
}
