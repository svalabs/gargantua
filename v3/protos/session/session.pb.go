// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: session/session.proto

package sessionpb

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

type Session struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid           string                 `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Scenario      string                 `protobuf:"bytes,3,opt,name=scenario,proto3" json:"scenario,omitempty"`
	Course        string                 `protobuf:"bytes,4,opt,name=course,proto3" json:"course,omitempty"`
	KeepCourseVm  bool                   `protobuf:"varint,5,opt,name=keep_course_vm,json=keepCourseVm,proto3" json:"keep_course_vm,omitempty"`
	User          string                 `protobuf:"bytes,6,opt,name=user,proto3" json:"user,omitempty"`
	VmClaim       []string               `protobuf:"bytes,7,rep,name=vm_claim,json=vmClaim,proto3" json:"vm_claim,omitempty"`
	AccessCode    string                 `protobuf:"bytes,8,opt,name=access_code,json=accessCode,proto3" json:"access_code,omitempty"`
	Labels        map[string]string      `protobuf:"bytes,9,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Status        *SessionStatus         `protobuf:"bytes,10,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Session) Reset() {
	*x = Session{}
	mi := &file_session_session_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_session_session_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session.ProtoReflect.Descriptor instead.
func (*Session) Descriptor() ([]byte, []int) {
	return file_session_session_proto_rawDescGZIP(), []int{0}
}

func (x *Session) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Session) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Session) GetScenario() string {
	if x != nil {
		return x.Scenario
	}
	return ""
}

func (x *Session) GetCourse() string {
	if x != nil {
		return x.Course
	}
	return ""
}

func (x *Session) GetKeepCourseVm() bool {
	if x != nil {
		return x.KeepCourseVm
	}
	return false
}

func (x *Session) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *Session) GetVmClaim() []string {
	if x != nil {
		return x.VmClaim
	}
	return nil
}

func (x *Session) GetAccessCode() string {
	if x != nil {
		return x.AccessCode
	}
	return ""
}

func (x *Session) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Session) GetStatus() *SessionStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

type CreateSessionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Scenario      string                 `protobuf:"bytes,1,opt,name=scenario,proto3" json:"scenario,omitempty"`
	Course        string                 `protobuf:"bytes,2,opt,name=course,proto3" json:"course,omitempty"`
	KeepCourseVm  bool                   `protobuf:"varint,3,opt,name=keep_course_vm,json=keepCourseVm,proto3" json:"keep_course_vm,omitempty"`
	User          string                 `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	VmClaim       []string               `protobuf:"bytes,5,rep,name=vm_claim,json=vmClaim,proto3" json:"vm_claim,omitempty"`
	AccessCode    string                 `protobuf:"bytes,6,opt,name=access_code,json=accessCode,proto3" json:"access_code,omitempty"`
	Labels        map[string]string      `protobuf:"bytes,7,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateSessionRequest) Reset() {
	*x = CreateSessionRequest{}
	mi := &file_session_session_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSessionRequest) ProtoMessage() {}

func (x *CreateSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_session_session_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSessionRequest.ProtoReflect.Descriptor instead.
func (*CreateSessionRequest) Descriptor() ([]byte, []int) {
	return file_session_session_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSessionRequest) GetScenario() string {
	if x != nil {
		return x.Scenario
	}
	return ""
}

func (x *CreateSessionRequest) GetCourse() string {
	if x != nil {
		return x.Course
	}
	return ""
}

func (x *CreateSessionRequest) GetKeepCourseVm() bool {
	if x != nil {
		return x.KeepCourseVm
	}
	return false
}

func (x *CreateSessionRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *CreateSessionRequest) GetVmClaim() []string {
	if x != nil {
		return x.VmClaim
	}
	return nil
}

func (x *CreateSessionRequest) GetAccessCode() string {
	if x != nil {
		return x.AccessCode
	}
	return ""
}

func (x *CreateSessionRequest) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

// Currently sessions are bound to a course, user and access code
// Thus, only the scenario for a session can be updated
type UpdateSessionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Scenario      string                 `protobuf:"bytes,2,opt,name=scenario,proto3" json:"scenario,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateSessionRequest) Reset() {
	*x = UpdateSessionRequest{}
	mi := &file_session_session_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSessionRequest) ProtoMessage() {}

func (x *UpdateSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_session_session_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSessionRequest.ProtoReflect.Descriptor instead.
func (*UpdateSessionRequest) Descriptor() ([]byte, []int) {
	return file_session_session_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateSessionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateSessionRequest) GetScenario() string {
	if x != nil {
		return x.Scenario
	}
	return ""
}

type UpdateSessionStatusRequest struct {
	state          protoimpl.MessageState  `protogen:"open.v1"`
	Id             string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Paused         *wrapperspb.BoolValue   `protobuf:"bytes,2,opt,name=paused,proto3" json:"paused,omitempty"`
	PausedTime     *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=paused_time,json=pausedTime,proto3" json:"paused_time,omitempty"`
	Active         *wrapperspb.BoolValue   `protobuf:"bytes,4,opt,name=active,proto3" json:"active,omitempty"`
	Finished       *wrapperspb.BoolValue   `protobuf:"bytes,5,opt,name=finished,proto3" json:"finished,omitempty"`
	StartTime      string                  `protobuf:"bytes,6,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	ExpirationTime string                  `protobuf:"bytes,7,opt,name=expiration_time,json=expirationTime,proto3" json:"expiration_time,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *UpdateSessionStatusRequest) Reset() {
	*x = UpdateSessionStatusRequest{}
	mi := &file_session_session_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateSessionStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSessionStatusRequest) ProtoMessage() {}

func (x *UpdateSessionStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_session_session_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSessionStatusRequest.ProtoReflect.Descriptor instead.
func (*UpdateSessionStatusRequest) Descriptor() ([]byte, []int) {
	return file_session_session_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateSessionStatusRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateSessionStatusRequest) GetPaused() *wrapperspb.BoolValue {
	if x != nil {
		return x.Paused
	}
	return nil
}

func (x *UpdateSessionStatusRequest) GetPausedTime() *wrapperspb.StringValue {
	if x != nil {
		return x.PausedTime
	}
	return nil
}

func (x *UpdateSessionStatusRequest) GetActive() *wrapperspb.BoolValue {
	if x != nil {
		return x.Active
	}
	return nil
}

func (x *UpdateSessionStatusRequest) GetFinished() *wrapperspb.BoolValue {
	if x != nil {
		return x.Finished
	}
	return nil
}

func (x *UpdateSessionStatusRequest) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *UpdateSessionStatusRequest) GetExpirationTime() string {
	if x != nil {
		return x.ExpirationTime
	}
	return ""
}

type SessionStatus struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Paused         bool                   `protobuf:"varint,1,opt,name=paused,proto3" json:"paused,omitempty"`
	PausedTime     string                 `protobuf:"bytes,2,opt,name=paused_time,json=pausedTime,proto3" json:"paused_time,omitempty"`
	Active         bool                   `protobuf:"varint,3,opt,name=active,proto3" json:"active,omitempty"`
	Finished       bool                   `protobuf:"varint,4,opt,name=finished,proto3" json:"finished,omitempty"`
	StartTime      string                 `protobuf:"bytes,5,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	ExpirationTime string                 `protobuf:"bytes,6,opt,name=expiration_time,json=expirationTime,proto3" json:"expiration_time,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *SessionStatus) Reset() {
	*x = SessionStatus{}
	mi := &file_session_session_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SessionStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionStatus) ProtoMessage() {}

func (x *SessionStatus) ProtoReflect() protoreflect.Message {
	mi := &file_session_session_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionStatus.ProtoReflect.Descriptor instead.
func (*SessionStatus) Descriptor() ([]byte, []int) {
	return file_session_session_proto_rawDescGZIP(), []int{4}
}

func (x *SessionStatus) GetPaused() bool {
	if x != nil {
		return x.Paused
	}
	return false
}

func (x *SessionStatus) GetPausedTime() string {
	if x != nil {
		return x.PausedTime
	}
	return ""
}

func (x *SessionStatus) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *SessionStatus) GetFinished() bool {
	if x != nil {
		return x.Finished
	}
	return false
}

func (x *SessionStatus) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *SessionStatus) GetExpirationTime() string {
	if x != nil {
		return x.ExpirationTime
	}
	return ""
}

type ListSessionsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Sessions      []*Session             `protobuf:"bytes,1,rep,name=sessions,proto3" json:"sessions,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListSessionsResponse) Reset() {
	*x = ListSessionsResponse{}
	mi := &file_session_session_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListSessionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSessionsResponse) ProtoMessage() {}

func (x *ListSessionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_session_session_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSessionsResponse.ProtoReflect.Descriptor instead.
func (*ListSessionsResponse) Descriptor() ([]byte, []int) {
	return file_session_session_proto_rawDescGZIP(), []int{5}
}

func (x *ListSessionsResponse) GetSessions() []*Session {
	if x != nil {
		return x.Sessions
	}
	return nil
}

var File_session_session_proto protoreflect.FileDescriptor

var file_session_session_proto_rawDesc = string([]byte{
	0x0a, 0x15, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x1a, 0x15, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf6, 0x02, 0x0a, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x12, 0x16,
	0x0a, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x6b, 0x65, 0x65, 0x70, 0x5f, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x76, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c,
	0x6b, 0x65, 0x65, 0x70, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x56, 0x6d, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x12, 0x19, 0x0a, 0x08, 0x76, 0x6d, 0x5f, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x76, 0x6d, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x12, 0x1f, 0x0a, 0x0b, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x34, 0x0a, 0x06,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xbe, 0x02,
	0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72,
	0x69, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72,
	0x69, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x6b, 0x65,
	0x65, 0x70, 0x5f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x76, 0x6d, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0c, 0x6b, 0x65, 0x65, 0x70, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x56, 0x6d,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x6d, 0x5f, 0x63, 0x6c, 0x61, 0x69, 0x6d,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x76, 0x6d, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x12,
	0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x41, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x29, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x42,
	0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72,
	0x69, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72,
	0x69, 0x6f, 0x22, 0xd3, 0x02, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x32, 0x0a, 0x06, 0x70, 0x61, 0x75, 0x73, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x70,
	0x61, 0x75, 0x73, 0x65, 0x64, 0x12, 0x3d, 0x0a, 0x0b, 0x70, 0x61, 0x75, 0x73, 0x65, 0x64, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x70, 0x61, 0x75, 0x73, 0x65, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x36, 0x0a, 0x08, 0x66, 0x69, 0x6e, 0x69,
	0x73, 0x68, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f,
	0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x27, 0x0a, 0x0f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xc4, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61,
	0x75, 0x73, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x70, 0x61, 0x75, 0x73,
	0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x75, 0x73, 0x65, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66,
	0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x66,
	0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x22,
	0x44, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0xed, 0x03, 0x0a, 0x0a, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x53, 0x76, 0x63, 0x12, 0x43, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x13, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x6c, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x46,
	0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x1d, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x52, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x2e,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3c, 0x0a, 0x0d, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x13, 0x2e, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x47, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x14, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x42, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x6f, 0x62, 0x62, 0x79, 0x66, 0x61, 0x72, 0x6d, 0x2f, 0x67, 0x61,
	0x72, 0x67, 0x61, 0x6e, 0x74, 0x75, 0x61, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x3b, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_session_session_proto_rawDescOnce sync.Once
	file_session_session_proto_rawDescData []byte
)

func file_session_session_proto_rawDescGZIP() []byte {
	file_session_session_proto_rawDescOnce.Do(func() {
		file_session_session_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_session_session_proto_rawDesc), len(file_session_session_proto_rawDesc)))
	})
	return file_session_session_proto_rawDescData
}

var file_session_session_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_session_session_proto_goTypes = []any{
	(*Session)(nil),                    // 0: session.Session
	(*CreateSessionRequest)(nil),       // 1: session.CreateSessionRequest
	(*UpdateSessionRequest)(nil),       // 2: session.UpdateSessionRequest
	(*UpdateSessionStatusRequest)(nil), // 3: session.UpdateSessionStatusRequest
	(*SessionStatus)(nil),              // 4: session.SessionStatus
	(*ListSessionsResponse)(nil),       // 5: session.ListSessionsResponse
	nil,                                // 6: session.Session.LabelsEntry
	nil,                                // 7: session.CreateSessionRequest.LabelsEntry
	(*wrapperspb.BoolValue)(nil),       // 8: google.protobuf.BoolValue
	(*wrapperspb.StringValue)(nil),     // 9: google.protobuf.StringValue
	(*general.GetRequest)(nil),         // 10: general.GetRequest
	(*general.ResourceId)(nil),         // 11: general.ResourceId
	(*general.ListOptions)(nil),        // 12: general.ListOptions
	(*emptypb.Empty)(nil),              // 13: google.protobuf.Empty
}
var file_session_session_proto_depIdxs = []int32{
	6,  // 0: session.Session.labels:type_name -> session.Session.LabelsEntry
	4,  // 1: session.Session.status:type_name -> session.SessionStatus
	7,  // 2: session.CreateSessionRequest.labels:type_name -> session.CreateSessionRequest.LabelsEntry
	8,  // 3: session.UpdateSessionStatusRequest.paused:type_name -> google.protobuf.BoolValue
	9,  // 4: session.UpdateSessionStatusRequest.paused_time:type_name -> google.protobuf.StringValue
	8,  // 5: session.UpdateSessionStatusRequest.active:type_name -> google.protobuf.BoolValue
	8,  // 6: session.UpdateSessionStatusRequest.finished:type_name -> google.protobuf.BoolValue
	0,  // 7: session.ListSessionsResponse.sessions:type_name -> session.Session
	1,  // 8: session.SessionSvc.CreateSession:input_type -> session.CreateSessionRequest
	10, // 9: session.SessionSvc.GetSession:input_type -> general.GetRequest
	2,  // 10: session.SessionSvc.UpdateSession:input_type -> session.UpdateSessionRequest
	3,  // 11: session.SessionSvc.UpdateSessionStatus:input_type -> session.UpdateSessionStatusRequest
	11, // 12: session.SessionSvc.DeleteSession:input_type -> general.ResourceId
	12, // 13: session.SessionSvc.DeleteCollectionSession:input_type -> general.ListOptions
	12, // 14: session.SessionSvc.ListSession:input_type -> general.ListOptions
	11, // 15: session.SessionSvc.CreateSession:output_type -> general.ResourceId
	0,  // 16: session.SessionSvc.GetSession:output_type -> session.Session
	13, // 17: session.SessionSvc.UpdateSession:output_type -> google.protobuf.Empty
	13, // 18: session.SessionSvc.UpdateSessionStatus:output_type -> google.protobuf.Empty
	13, // 19: session.SessionSvc.DeleteSession:output_type -> google.protobuf.Empty
	13, // 20: session.SessionSvc.DeleteCollectionSession:output_type -> google.protobuf.Empty
	5,  // 21: session.SessionSvc.ListSession:output_type -> session.ListSessionsResponse
	15, // [15:22] is the sub-list for method output_type
	8,  // [8:15] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_session_session_proto_init() }
func file_session_session_proto_init() {
	if File_session_session_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_session_session_proto_rawDesc), len(file_session_session_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_session_session_proto_goTypes,
		DependencyIndexes: file_session_session_proto_depIdxs,
		MessageInfos:      file_session_session_proto_msgTypes,
	}.Build()
	File_session_session_proto = out.File
	file_session_session_proto_goTypes = nil
	file_session_session_proto_depIdxs = nil
}
