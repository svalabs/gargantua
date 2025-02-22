// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: scenario/scenario.proto

package scenariopb

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

type Scenario struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	Id                string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid               string                 `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Name              string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description       string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Steps             []*ScenarioStep        `protobuf:"bytes,5,rep,name=steps,proto3" json:"steps,omitempty"`
	Categories        []string               `protobuf:"bytes,6,rep,name=categories,proto3" json:"categories,omitempty"`
	Tags              []string               `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags,omitempty"`
	Vms               []*general.StringMap   `protobuf:"bytes,8,rep,name=vms,proto3" json:"vms,omitempty"`
	KeepaliveDuration string                 `protobuf:"bytes,9,opt,name=keepalive_duration,json=keepaliveDuration,proto3" json:"keepalive_duration,omitempty"`
	PauseDuration     string                 `protobuf:"bytes,10,opt,name=pause_duration,json=pauseDuration,proto3" json:"pause_duration,omitempty"`
	Pausable          bool                   `protobuf:"varint,11,opt,name=pausable,proto3" json:"pausable,omitempty"`
	VmTasks           []*VirtualMachineTasks `protobuf:"bytes,12,rep,name=vm_tasks,json=vmTasks,proto3" json:"vm_tasks,omitempty"`
	Labels            map[string]string      `protobuf:"bytes,13,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *Scenario) Reset() {
	*x = Scenario{}
	mi := &file_scenario_scenario_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Scenario) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scenario) ProtoMessage() {}

func (x *Scenario) ProtoReflect() protoreflect.Message {
	mi := &file_scenario_scenario_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scenario.ProtoReflect.Descriptor instead.
func (*Scenario) Descriptor() ([]byte, []int) {
	return file_scenario_scenario_proto_rawDescGZIP(), []int{0}
}

func (x *Scenario) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Scenario) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Scenario) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Scenario) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Scenario) GetSteps() []*ScenarioStep {
	if x != nil {
		return x.Steps
	}
	return nil
}

func (x *Scenario) GetCategories() []string {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *Scenario) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Scenario) GetVms() []*general.StringMap {
	if x != nil {
		return x.Vms
	}
	return nil
}

func (x *Scenario) GetKeepaliveDuration() string {
	if x != nil {
		return x.KeepaliveDuration
	}
	return ""
}

func (x *Scenario) GetPauseDuration() string {
	if x != nil {
		return x.PauseDuration
	}
	return ""
}

func (x *Scenario) GetPausable() bool {
	if x != nil {
		return x.Pausable
	}
	return false
}

func (x *Scenario) GetVmTasks() []*VirtualMachineTasks {
	if x != nil {
		return x.VmTasks
	}
	return nil
}

func (x *Scenario) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type CreateScenarioRequest struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	Name              string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description       string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	RawSteps          string                 `protobuf:"bytes,3,opt,name=raw_steps,json=rawSteps,proto3" json:"raw_steps,omitempty"`
	RawCategories     string                 `protobuf:"bytes,4,opt,name=raw_categories,json=rawCategories,proto3" json:"raw_categories,omitempty"`
	RawTags           string                 `protobuf:"bytes,5,opt,name=raw_tags,json=rawTags,proto3" json:"raw_tags,omitempty"`
	RawVms            string                 `protobuf:"bytes,6,opt,name=raw_vms,json=rawVms,proto3" json:"raw_vms,omitempty"`
	RawVmTasks        string                 `protobuf:"bytes,7,opt,name=raw_vm_tasks,json=rawVmTasks,proto3" json:"raw_vm_tasks,omitempty"`
	KeepaliveDuration string                 `protobuf:"bytes,8,opt,name=keepalive_duration,json=keepaliveDuration,proto3" json:"keepalive_duration,omitempty"`
	PauseDuration     string                 `protobuf:"bytes,9,opt,name=pause_duration,json=pauseDuration,proto3" json:"pause_duration,omitempty"`
	Pausable          bool                   `protobuf:"varint,10,opt,name=pausable,proto3" json:"pausable,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *CreateScenarioRequest) Reset() {
	*x = CreateScenarioRequest{}
	mi := &file_scenario_scenario_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateScenarioRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateScenarioRequest) ProtoMessage() {}

func (x *CreateScenarioRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scenario_scenario_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateScenarioRequest.ProtoReflect.Descriptor instead.
func (*CreateScenarioRequest) Descriptor() ([]byte, []int) {
	return file_scenario_scenario_proto_rawDescGZIP(), []int{1}
}

func (x *CreateScenarioRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateScenarioRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateScenarioRequest) GetRawSteps() string {
	if x != nil {
		return x.RawSteps
	}
	return ""
}

func (x *CreateScenarioRequest) GetRawCategories() string {
	if x != nil {
		return x.RawCategories
	}
	return ""
}

func (x *CreateScenarioRequest) GetRawTags() string {
	if x != nil {
		return x.RawTags
	}
	return ""
}

func (x *CreateScenarioRequest) GetRawVms() string {
	if x != nil {
		return x.RawVms
	}
	return ""
}

func (x *CreateScenarioRequest) GetRawVmTasks() string {
	if x != nil {
		return x.RawVmTasks
	}
	return ""
}

func (x *CreateScenarioRequest) GetKeepaliveDuration() string {
	if x != nil {
		return x.KeepaliveDuration
	}
	return ""
}

func (x *CreateScenarioRequest) GetPauseDuration() string {
	if x != nil {
		return x.PauseDuration
	}
	return ""
}

func (x *CreateScenarioRequest) GetPausable() bool {
	if x != nil {
		return x.Pausable
	}
	return false
}

type ScenarioStep struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content       string                 `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ScenarioStep) Reset() {
	*x = ScenarioStep{}
	mi := &file_scenario_scenario_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ScenarioStep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScenarioStep) ProtoMessage() {}

func (x *ScenarioStep) ProtoReflect() protoreflect.Message {
	mi := &file_scenario_scenario_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScenarioStep.ProtoReflect.Descriptor instead.
func (*ScenarioStep) Descriptor() ([]byte, []int) {
	return file_scenario_scenario_proto_rawDescGZIP(), []int{2}
}

func (x *ScenarioStep) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ScenarioStep) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type UpdateScenarioRequest struct {
	state             protoimpl.MessageState  `protogen:"open.v1"`
	Id                string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name              string                  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description       string                  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	RawSteps          string                  `protobuf:"bytes,4,opt,name=raw_steps,json=rawSteps,proto3" json:"raw_steps,omitempty"`
	RawCategories     string                  `protobuf:"bytes,5,opt,name=raw_categories,json=rawCategories,proto3" json:"raw_categories,omitempty"`
	RawTags           string                  `protobuf:"bytes,6,opt,name=raw_tags,json=rawTags,proto3" json:"raw_tags,omitempty"`
	RawVms            string                  `protobuf:"bytes,7,opt,name=raw_vms,json=rawVms,proto3" json:"raw_vms,omitempty"`
	RawVmTasks        string                  `protobuf:"bytes,8,opt,name=raw_vm_tasks,json=rawVmTasks,proto3" json:"raw_vm_tasks,omitempty"`
	KeepaliveDuration *wrapperspb.StringValue `protobuf:"bytes,9,opt,name=keepalive_duration,json=keepaliveDuration,proto3" json:"keepalive_duration,omitempty"`
	PauseDuration     *wrapperspb.StringValue `protobuf:"bytes,10,opt,name=pause_duration,json=pauseDuration,proto3" json:"pause_duration,omitempty"`
	Pausable          *wrapperspb.BoolValue   `protobuf:"bytes,11,opt,name=pausable,proto3" json:"pausable,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *UpdateScenarioRequest) Reset() {
	*x = UpdateScenarioRequest{}
	mi := &file_scenario_scenario_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateScenarioRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateScenarioRequest) ProtoMessage() {}

func (x *UpdateScenarioRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scenario_scenario_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateScenarioRequest.ProtoReflect.Descriptor instead.
func (*UpdateScenarioRequest) Descriptor() ([]byte, []int) {
	return file_scenario_scenario_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateScenarioRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateScenarioRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateScenarioRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateScenarioRequest) GetRawSteps() string {
	if x != nil {
		return x.RawSteps
	}
	return ""
}

func (x *UpdateScenarioRequest) GetRawCategories() string {
	if x != nil {
		return x.RawCategories
	}
	return ""
}

func (x *UpdateScenarioRequest) GetRawTags() string {
	if x != nil {
		return x.RawTags
	}
	return ""
}

func (x *UpdateScenarioRequest) GetRawVms() string {
	if x != nil {
		return x.RawVms
	}
	return ""
}

func (x *UpdateScenarioRequest) GetRawVmTasks() string {
	if x != nil {
		return x.RawVmTasks
	}
	return ""
}

func (x *UpdateScenarioRequest) GetKeepaliveDuration() *wrapperspb.StringValue {
	if x != nil {
		return x.KeepaliveDuration
	}
	return nil
}

func (x *UpdateScenarioRequest) GetPauseDuration() *wrapperspb.StringValue {
	if x != nil {
		return x.PauseDuration
	}
	return nil
}

func (x *UpdateScenarioRequest) GetPausable() *wrapperspb.BoolValue {
	if x != nil {
		return x.Pausable
	}
	return nil
}

type ListScenariosResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Scenarios     []*Scenario            `protobuf:"bytes,1,rep,name=scenarios,proto3" json:"scenarios,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListScenariosResponse) Reset() {
	*x = ListScenariosResponse{}
	mi := &file_scenario_scenario_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListScenariosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListScenariosResponse) ProtoMessage() {}

func (x *ListScenariosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_scenario_scenario_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListScenariosResponse.ProtoReflect.Descriptor instead.
func (*ListScenariosResponse) Descriptor() ([]byte, []int) {
	return file_scenario_scenario_proto_rawDescGZIP(), []int{4}
}

func (x *ListScenariosResponse) GetScenarios() []*Scenario {
	if x != nil {
		return x.Scenarios
	}
	return nil
}

type VirtualMachineTasks struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	VmName        string                 `protobuf:"bytes,1,opt,name=vm_name,json=vmName,proto3" json:"vm_name,omitempty"`
	Tasks         []*Task                `protobuf:"bytes,2,rep,name=tasks,proto3" json:"tasks,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VirtualMachineTasks) Reset() {
	*x = VirtualMachineTasks{}
	mi := &file_scenario_scenario_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VirtualMachineTasks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VirtualMachineTasks) ProtoMessage() {}

func (x *VirtualMachineTasks) ProtoReflect() protoreflect.Message {
	mi := &file_scenario_scenario_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VirtualMachineTasks.ProtoReflect.Descriptor instead.
func (*VirtualMachineTasks) Descriptor() ([]byte, []int) {
	return file_scenario_scenario_proto_rawDescGZIP(), []int{5}
}

func (x *VirtualMachineTasks) GetVmName() string {
	if x != nil {
		return x.VmName
	}
	return ""
}

func (x *VirtualMachineTasks) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type Task struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	Name                string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description         string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Command             string                 `protobuf:"bytes,3,opt,name=command,proto3" json:"command,omitempty"`
	ExpectedOutputValue string                 `protobuf:"bytes,4,opt,name=expected_output_value,json=expectedOutputValue,proto3" json:"expected_output_value,omitempty"`
	ExpectedReturnCode  int32                  `protobuf:"varint,5,opt,name=expected_return_code,json=expectedReturnCode,proto3" json:"expected_return_code,omitempty"`
	ReturnType          string                 `protobuf:"bytes,6,opt,name=return_type,json=returnType,proto3" json:"return_type,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *Task) Reset() {
	*x = Task{}
	mi := &file_scenario_scenario_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_scenario_scenario_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_scenario_scenario_proto_rawDescGZIP(), []int{6}
}

func (x *Task) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Task) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Task) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *Task) GetExpectedOutputValue() string {
	if x != nil {
		return x.ExpectedOutputValue
	}
	return ""
}

func (x *Task) GetExpectedReturnCode() int32 {
	if x != nil {
		return x.ExpectedReturnCode
	}
	return 0
}

func (x *Task) GetReturnType() string {
	if x != nil {
		return x.ReturnType
	}
	return ""
}

var File_scenario_scenario_proto protoreflect.FileDescriptor

var file_scenario_scenario_proto_rawDesc = string([]byte{
	0x0a, 0x17, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x2f, 0x73, 0x63, 0x65, 0x6e, 0x61,
	0x72, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x63, 0x65, 0x6e, 0x61,
	0x72, 0x69, 0x6f, 0x1a, 0x15, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x04, 0x0a, 0x08, 0x53, 0x63, 0x65, 0x6e,
	0x61, 0x72, 0x69, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x05,
	0x73, 0x74, 0x65, 0x70, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x63,
	0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x53,
	0x74, 0x65, 0x70, 0x52, 0x05, 0x73, 0x74, 0x65, 0x70, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x24,
	0x0a, 0x03, 0x76, 0x6d, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x70, 0x52,
	0x03, 0x76, 0x6d, 0x73, 0x12, 0x2d, 0x0a, 0x12, 0x6b, 0x65, 0x65, 0x70, 0x61, 0x6c, 0x69, 0x76,
	0x65, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x6b, 0x65, 0x65, 0x70, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x75, 0x73, 0x65, 0x5f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x75,
	0x73, 0x65, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x75, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x70, 0x61,
	0x75, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x38, 0x0a, 0x08, 0x76, 0x6d, 0x5f, 0x74, 0x61, 0x73,
	0x6b, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x63, 0x65, 0x6e, 0x61,
	0x72, 0x69, 0x6f, 0x2e, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x4d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x07, 0x76, 0x6d, 0x54, 0x61, 0x73, 0x6b, 0x73,
	0x12, 0x36, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x2e, 0x53, 0x63, 0x65, 0x6e,
	0x61, 0x72, 0x69, 0x6f, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0xd9, 0x02, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63,
	0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x61, 0x77, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x61, 0x77, 0x53, 0x74, 0x65, 0x70, 0x73,
	0x12, 0x25, 0x0a, 0x0e, 0x72, 0x61, 0x77, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x61, 0x77, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x61, 0x77, 0x5f, 0x74,
	0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x61, 0x77, 0x54, 0x61,
	0x67, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x61, 0x77, 0x5f, 0x76, 0x6d, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x61, 0x77, 0x56, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x0c, 0x72,
	0x61, 0x77, 0x5f, 0x76, 0x6d, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x72, 0x61, 0x77, 0x56, 0x6d, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x2d, 0x0a,
	0x12, 0x6b, 0x65, 0x65, 0x70, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6b, 0x65, 0x65, 0x70, 0x61,
	0x6c, 0x69, 0x76, 0x65, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e,
	0x70, 0x61, 0x75, 0x73, 0x65, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x75, 0x73, 0x65, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x75, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x70, 0x61, 0x75, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x22,
	0x3e, 0x0a, 0x0c, 0x53, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x53, 0x74, 0x65, 0x70, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22,
	0xc1, 0x03, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x65, 0x6e, 0x61, 0x72,
	0x69, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1b, 0x0a, 0x09, 0x72, 0x61, 0x77, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x72, 0x61, 0x77, 0x53, 0x74, 0x65, 0x70, 0x73, 0x12, 0x25, 0x0a, 0x0e,
	0x72, 0x61, 0x77, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x61, 0x77, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x61, 0x77, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x61, 0x77, 0x54, 0x61, 0x67, 0x73, 0x12, 0x17,
	0x0a, 0x07, 0x72, 0x61, 0x77, 0x5f, 0x76, 0x6d, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x61, 0x77, 0x56, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x0c, 0x72, 0x61, 0x77, 0x5f, 0x76,
	0x6d, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72,
	0x61, 0x77, 0x56, 0x6d, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x4b, 0x0a, 0x12, 0x6b, 0x65, 0x65,
	0x70, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x11, 0x6b, 0x65, 0x65, 0x70, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x43, 0x0a, 0x0e, 0x70, 0x61, 0x75, 0x73, 0x65, 0x5f,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0d, 0x70, 0x61,
	0x75, 0x73, 0x65, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x08, 0x70,
	0x61, 0x75, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x70, 0x61, 0x75, 0x73, 0x61,
	0x62, 0x6c, 0x65, 0x22, 0x49, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x63, 0x65, 0x6e, 0x61,
	0x72, 0x69, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x09,
	0x73, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x61,
	0x72, 0x69, 0x6f, 0x52, 0x09, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x73, 0x22, 0x54,
	0x0a, 0x13, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x76, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24,
	0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x73, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x74,
	0x61, 0x73, 0x6b, 0x73, 0x22, 0xdd, 0x01, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x32, 0x0a,
	0x15, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x65, 0x78,
	0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x30, 0x0a, 0x14, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x12, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x32, 0xe5, 0x03, 0x0a, 0x0b, 0x53, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69,
	0x6f, 0x53, 0x76, 0x63, 0x12, 0x46, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63,
	0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x12, 0x1f, 0x2e, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69,
	0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x6c, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x53, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x12, 0x13, 0x2e, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x2e, 0x53, 0x63, 0x65, 0x6e,
	0x61, 0x72, 0x69, 0x6f, 0x12, 0x49, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63,
	0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x12, 0x1f, 0x2e, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69,
	0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x3d, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69,
	0x6f, 0x12, 0x13, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x48,
	0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x12, 0x14, 0x2e, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x6c, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x45, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x12, 0x14, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x6c, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x1f,
	0x2e, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x63,
	0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3b, 0x0a, 0x0c, 0x43, 0x6f, 0x70, 0x79, 0x53, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x12,
	0x13, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x3e, 0x5a, 0x3c,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x6f, 0x62, 0x62, 0x79,
	0x66, 0x61, 0x72, 0x6d, 0x2f, 0x67, 0x61, 0x72, 0x67, 0x61, 0x6e, 0x74, 0x75, 0x61, 0x2f, 0x76,
	0x33, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69,
	0x6f, 0x3b, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_scenario_scenario_proto_rawDescOnce sync.Once
	file_scenario_scenario_proto_rawDescData []byte
)

func file_scenario_scenario_proto_rawDescGZIP() []byte {
	file_scenario_scenario_proto_rawDescOnce.Do(func() {
		file_scenario_scenario_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_scenario_scenario_proto_rawDesc), len(file_scenario_scenario_proto_rawDesc)))
	})
	return file_scenario_scenario_proto_rawDescData
}

var file_scenario_scenario_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_scenario_scenario_proto_goTypes = []any{
	(*Scenario)(nil),               // 0: scenario.Scenario
	(*CreateScenarioRequest)(nil),  // 1: scenario.CreateScenarioRequest
	(*ScenarioStep)(nil),           // 2: scenario.ScenarioStep
	(*UpdateScenarioRequest)(nil),  // 3: scenario.UpdateScenarioRequest
	(*ListScenariosResponse)(nil),  // 4: scenario.ListScenariosResponse
	(*VirtualMachineTasks)(nil),    // 5: scenario.VirtualMachineTasks
	(*Task)(nil),                   // 6: scenario.Task
	nil,                            // 7: scenario.Scenario.LabelsEntry
	(*general.StringMap)(nil),      // 8: general.StringMap
	(*wrapperspb.StringValue)(nil), // 9: google.protobuf.StringValue
	(*wrapperspb.BoolValue)(nil),   // 10: google.protobuf.BoolValue
	(*general.GetRequest)(nil),     // 11: general.GetRequest
	(*general.ResourceId)(nil),     // 12: general.ResourceId
	(*general.ListOptions)(nil),    // 13: general.ListOptions
	(*emptypb.Empty)(nil),          // 14: google.protobuf.Empty
}
var file_scenario_scenario_proto_depIdxs = []int32{
	2,  // 0: scenario.Scenario.steps:type_name -> scenario.ScenarioStep
	8,  // 1: scenario.Scenario.vms:type_name -> general.StringMap
	5,  // 2: scenario.Scenario.vm_tasks:type_name -> scenario.VirtualMachineTasks
	7,  // 3: scenario.Scenario.labels:type_name -> scenario.Scenario.LabelsEntry
	9,  // 4: scenario.UpdateScenarioRequest.keepalive_duration:type_name -> google.protobuf.StringValue
	9,  // 5: scenario.UpdateScenarioRequest.pause_duration:type_name -> google.protobuf.StringValue
	10, // 6: scenario.UpdateScenarioRequest.pausable:type_name -> google.protobuf.BoolValue
	0,  // 7: scenario.ListScenariosResponse.scenarios:type_name -> scenario.Scenario
	6,  // 8: scenario.VirtualMachineTasks.tasks:type_name -> scenario.Task
	1,  // 9: scenario.ScenarioSvc.CreateScenario:input_type -> scenario.CreateScenarioRequest
	11, // 10: scenario.ScenarioSvc.GetScenario:input_type -> general.GetRequest
	3,  // 11: scenario.ScenarioSvc.UpdateScenario:input_type -> scenario.UpdateScenarioRequest
	12, // 12: scenario.ScenarioSvc.DeleteScenario:input_type -> general.ResourceId
	13, // 13: scenario.ScenarioSvc.DeleteCollectionScenario:input_type -> general.ListOptions
	13, // 14: scenario.ScenarioSvc.ListScenario:input_type -> general.ListOptions
	12, // 15: scenario.ScenarioSvc.CopyScenario:input_type -> general.ResourceId
	12, // 16: scenario.ScenarioSvc.CreateScenario:output_type -> general.ResourceId
	0,  // 17: scenario.ScenarioSvc.GetScenario:output_type -> scenario.Scenario
	14, // 18: scenario.ScenarioSvc.UpdateScenario:output_type -> google.protobuf.Empty
	14, // 19: scenario.ScenarioSvc.DeleteScenario:output_type -> google.protobuf.Empty
	14, // 20: scenario.ScenarioSvc.DeleteCollectionScenario:output_type -> google.protobuf.Empty
	4,  // 21: scenario.ScenarioSvc.ListScenario:output_type -> scenario.ListScenariosResponse
	14, // 22: scenario.ScenarioSvc.CopyScenario:output_type -> google.protobuf.Empty
	16, // [16:23] is the sub-list for method output_type
	9,  // [9:16] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_scenario_scenario_proto_init() }
func file_scenario_scenario_proto_init() {
	if File_scenario_scenario_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_scenario_scenario_proto_rawDesc), len(file_scenario_scenario_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_scenario_scenario_proto_goTypes,
		DependencyIndexes: file_scenario_scenario_proto_depIdxs,
		MessageInfos:      file_scenario_scenario_proto_msgTypes,
	}.Build()
	File_scenario_scenario_proto = out.File
	file_scenario_scenario_proto_goTypes = nil
	file_scenario_scenario_proto_depIdxs = nil
}
