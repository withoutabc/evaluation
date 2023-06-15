// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: job.proto

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

type CountJobReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level int32 `protobuf:"varint,1,opt,name=level,proto3" json:"level,omitempty"`
	Year  int32 `protobuf:"varint,2,opt,name=year,proto3" json:"year,omitempty"`
	Month int32 `protobuf:"varint,3,opt,name=month,proto3" json:"month,omitempty"`
	Set   int32 `protobuf:"varint,4,opt,name=set,proto3" json:"set,omitempty"`
}

func (x *CountJobReq) Reset() {
	*x = CountJobReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountJobReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountJobReq) ProtoMessage() {}

func (x *CountJobReq) ProtoReflect() protoreflect.Message {
	mi := &file_job_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountJobReq.ProtoReflect.Descriptor instead.
func (*CountJobReq) Descriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{0}
}

func (x *CountJobReq) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *CountJobReq) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *CountJobReq) GetMonth() int32 {
	if x != nil {
		return x.Month
	}
	return 0
}

func (x *CountJobReq) GetSet() int32 {
	if x != nil {
		return x.Set
	}
	return 0
}

type CountJobResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	JobName    string `protobuf:"bytes,3,opt,name=job_name,json=jobName,proto3" json:"job_name,omitempty"`
}

func (x *CountJobResp) Reset() {
	*x = CountJobResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountJobResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountJobResp) ProtoMessage() {}

func (x *CountJobResp) ProtoReflect() protoreflect.Message {
	mi := &file_job_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountJobResp.ProtoReflect.Descriptor instead.
func (*CountJobResp) Descriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{1}
}

func (x *CountJobResp) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *CountJobResp) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *CountJobResp) GetJobName() string {
	if x != nil {
		return x.JobName
	}
	return ""
}

type JoinJobReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Begin int32 `protobuf:"varint,1,opt,name=begin,proto3" json:"begin,omitempty"`
	End   int32 `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
	Level int32 `protobuf:"varint,3,opt,name=level,proto3" json:"level,omitempty"`
}

func (x *JoinJobReq) Reset() {
	*x = JoinJobReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinJobReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinJobReq) ProtoMessage() {}

func (x *JoinJobReq) ProtoReflect() protoreflect.Message {
	mi := &file_job_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinJobReq.ProtoReflect.Descriptor instead.
func (*JoinJobReq) Descriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{2}
}

func (x *JoinJobReq) GetBegin() int32 {
	if x != nil {
		return x.Begin
	}
	return 0
}

func (x *JoinJobReq) GetEnd() int32 {
	if x != nil {
		return x.End
	}
	return 0
}

func (x *JoinJobReq) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

type JoinJobResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	JobName    string `protobuf:"bytes,3,opt,name=job_name,json=jobName,proto3" json:"job_name,omitempty"`
}

func (x *JoinJobResp) Reset() {
	*x = JoinJobResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinJobResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinJobResp) ProtoMessage() {}

func (x *JoinJobResp) ProtoReflect() protoreflect.Message {
	mi := &file_job_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinJobResp.ProtoReflect.Descriptor instead.
func (*JoinJobResp) Descriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{3}
}

func (x *JoinJobResp) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *JoinJobResp) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *JoinJobResp) GetJobName() string {
	if x != nil {
		return x.JobName
	}
	return ""
}

type ViewJobsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ViewJobsReq) Reset() {
	*x = ViewJobsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewJobsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewJobsReq) ProtoMessage() {}

func (x *ViewJobsReq) ProtoReflect() protoreflect.Message {
	mi := &file_job_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewJobsReq.ProtoReflect.Descriptor instead.
func (*ViewJobsReq) Descriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{4}
}

func (x *ViewJobsReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Jobs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status       string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	FinalStatus  string `protobuf:"bytes,4,opt,name=final_status,json=finalStatus,proto3" json:"final_status,omitempty"`
	Progress     int32  `protobuf:"varint,5,opt,name=progress,proto3" json:"progress,omitempty"`
	TrackingUrl  string `protobuf:"bytes,6,opt,name=tracking_url,json=trackingUrl,proto3" json:"tracking_url,omitempty"`
	StartedTime  string `protobuf:"bytes,7,opt,name=started_time,json=startedTime,proto3" json:"started_time,omitempty"`
	LaunchTime   string `protobuf:"bytes,8,opt,name=launch_time,json=launchTime,proto3" json:"launch_time,omitempty"`
	FinishedTime string `protobuf:"bytes,9,opt,name=finished_time,json=finishedTime,proto3" json:"finished_time,omitempty"`
	ElapsedTime  string `protobuf:"bytes,10,opt,name=elapsed_time,json=elapsedTime,proto3" json:"elapsed_time,omitempty"`
}

func (x *Jobs) Reset() {
	*x = Jobs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Jobs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Jobs) ProtoMessage() {}

func (x *Jobs) ProtoReflect() protoreflect.Message {
	mi := &file_job_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Jobs.ProtoReflect.Descriptor instead.
func (*Jobs) Descriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{5}
}

func (x *Jobs) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Jobs) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Jobs) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Jobs) GetFinalStatus() string {
	if x != nil {
		return x.FinalStatus
	}
	return ""
}

func (x *Jobs) GetProgress() int32 {
	if x != nil {
		return x.Progress
	}
	return 0
}

func (x *Jobs) GetTrackingUrl() string {
	if x != nil {
		return x.TrackingUrl
	}
	return ""
}

func (x *Jobs) GetStartedTime() string {
	if x != nil {
		return x.StartedTime
	}
	return ""
}

func (x *Jobs) GetLaunchTime() string {
	if x != nil {
		return x.LaunchTime
	}
	return ""
}

func (x *Jobs) GetFinishedTime() string {
	if x != nil {
		return x.FinishedTime
	}
	return ""
}

func (x *Jobs) GetElapsedTime() string {
	if x != nil {
		return x.ElapsedTime
	}
	return ""
}

type ViewJobsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32   `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMsg  string  `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	Jobs       []*Jobs `protobuf:"bytes,3,rep,name=Jobs,proto3" json:"Jobs,omitempty"`
}

func (x *ViewJobsResp) Reset() {
	*x = ViewJobsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewJobsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewJobsResp) ProtoMessage() {}

func (x *ViewJobsResp) ProtoReflect() protoreflect.Message {
	mi := &file_job_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewJobsResp.ProtoReflect.Descriptor instead.
func (*ViewJobsResp) Descriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{6}
}

func (x *ViewJobsResp) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *ViewJobsResp) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *ViewJobsResp) GetJobs() []*Jobs {
	if x != nil {
		return x.Jobs
	}
	return nil
}

type ViewWordReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level int32 `protobuf:"varint,1,opt,name=level,proto3" json:"level,omitempty"`
	Begin int32 `protobuf:"varint,2,opt,name=begin,proto3" json:"begin,omitempty"`
	End   int32 `protobuf:"varint,3,opt,name=end,proto3" json:"end,omitempty"`
	Count int32 `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ViewWordReq) Reset() {
	*x = ViewWordReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewWordReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewWordReq) ProtoMessage() {}

func (x *ViewWordReq) ProtoReflect() protoreflect.Message {
	mi := &file_job_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewWordReq.ProtoReflect.Descriptor instead.
func (*ViewWordReq) Descriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{7}
}

func (x *ViewWordReq) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *ViewWordReq) GetBegin() int32 {
	if x != nil {
		return x.Begin
	}
	return 0
}

func (x *ViewWordReq) GetEnd() int32 {
	if x != nil {
		return x.End
	}
	return 0
}

func (x *ViewWordReq) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type Words struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Word  string `protobuf:"bytes,1,opt,name=word,proto3" json:"word,omitempty"`
	Count int64  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *Words) Reset() {
	*x = Words{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Words) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Words) ProtoMessage() {}

func (x *Words) ProtoReflect() protoreflect.Message {
	mi := &file_job_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Words.ProtoReflect.Descriptor instead.
func (*Words) Descriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{8}
}

func (x *Words) GetWord() string {
	if x != nil {
		return x.Word
	}
	return ""
}

func (x *Words) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type ViewWordResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32    `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMsg  string   `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	Words      []*Words `protobuf:"bytes,3,rep,name=Words,proto3" json:"Words,omitempty"`
}

func (x *ViewWordResp) Reset() {
	*x = ViewWordResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewWordResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewWordResp) ProtoMessage() {}

func (x *ViewWordResp) ProtoReflect() protoreflect.Message {
	mi := &file_job_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewWordResp.ProtoReflect.Descriptor instead.
func (*ViewWordResp) Descriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{9}
}

func (x *ViewWordResp) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *ViewWordResp) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *ViewWordResp) GetWords() []*Words {
	if x != nil {
		return x.Words
	}
	return nil
}

var File_job_proto protoreflect.FileDescriptor

var file_job_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x66, 0x69, 0x6c,
	0x65, 0x22, 0x5f, 0x0a, 0x0b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f,
	0x6e, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68,
	0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x73,
	0x65, 0x74, 0x22, 0x69, 0x0a, 0x0c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d,
	0x73, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x6a, 0x6f, 0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6a, 0x6f, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x4a, 0x0a,
	0x0a, 0x4a, 0x6f, 0x69, 0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x62,
	0x65, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x62, 0x65, 0x67, 0x69,
	0x6e, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x65, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x68, 0x0a, 0x0b, 0x4a, 0x6f, 0x69,
	0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x6a, 0x6f, 0x62, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6a, 0x6f, 0x62, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x1d, 0x0a, 0x0b, 0x56, 0x69, 0x65, 0x77, 0x4a, 0x6f, 0x62, 0x73, 0x52,
	0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0xb0, 0x02, 0x0a, 0x04, 0x4a, 0x6f, 0x62, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x69, 0x6e, 0x61, 0x6c,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66,
	0x69, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72,
	0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72,
	0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69,
	0x6e, 0x67, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x72,
	0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x55, 0x72, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x23, 0x0a,
	0x0d, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6c, 0x61, 0x70, 0x73, 0x65,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x6e, 0x0a, 0x0c, 0x56, 0x69, 0x65, 0x77, 0x4a, 0x6f, 0x62,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x1e, 0x0a, 0x04, 0x4a, 0x6f, 0x62, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x4a, 0x6f, 0x62, 0x73, 0x52,
	0x04, 0x4a, 0x6f, 0x62, 0x73, 0x22, 0x61, 0x0a, 0x0b, 0x56, 0x69, 0x65, 0x77, 0x57, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x65,
	0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x62, 0x65, 0x67, 0x69, 0x6e,
	0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x65,
	0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x31, 0x0a, 0x05, 0x57, 0x6f, 0x72, 0x64,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x71, 0x0a, 0x0c, 0x56,
	0x69, 0x65, 0x77, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x21, 0x0a, 0x05, 0x57,
	0x6f, 0x72, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x66, 0x69, 0x6c,
	0x65, 0x2e, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x05, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x32, 0xce,
	0x01, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x12, 0x31, 0x0a, 0x08, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4a,
	0x6f, 0x62, 0x12, 0x11, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4a,
	0x6f, 0x62, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2e, 0x0a, 0x07, 0x4a, 0x6f, 0x69,
	0x6e, 0x4a, 0x6f, 0x62, 0x12, 0x10, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x4a, 0x6f, 0x69, 0x6e,
	0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x4a, 0x6f,
	0x69, 0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x12, 0x31, 0x0a, 0x08, 0x56, 0x69, 0x65,
	0x77, 0x4a, 0x6f, 0x62, 0x73, 0x12, 0x11, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x56, 0x69, 0x65,
	0x77, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e,
	0x56, 0x69, 0x65, 0x77, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x31, 0x0a, 0x08,
	0x56, 0x69, 0x65, 0x77, 0x57, 0x6f, 0x72, 0x64, 0x12, 0x11, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e,
	0x56, 0x69, 0x65, 0x77, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x66, 0x69,
	0x6c, 0x65, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x42,
	0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_job_proto_rawDescOnce sync.Once
	file_job_proto_rawDescData = file_job_proto_rawDesc
)

func file_job_proto_rawDescGZIP() []byte {
	file_job_proto_rawDescOnce.Do(func() {
		file_job_proto_rawDescData = protoimpl.X.CompressGZIP(file_job_proto_rawDescData)
	})
	return file_job_proto_rawDescData
}

var file_job_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_job_proto_goTypes = []interface{}{
	(*CountJobReq)(nil),  // 0: file.CountJobReq
	(*CountJobResp)(nil), // 1: file.CountJobResp
	(*JoinJobReq)(nil),   // 2: file.JoinJobReq
	(*JoinJobResp)(nil),  // 3: file.JoinJobResp
	(*ViewJobsReq)(nil),  // 4: file.ViewJobsReq
	(*Jobs)(nil),         // 5: file.Jobs
	(*ViewJobsResp)(nil), // 6: file.ViewJobsResp
	(*ViewWordReq)(nil),  // 7: file.ViewWordReq
	(*Words)(nil),        // 8: file.Words
	(*ViewWordResp)(nil), // 9: file.ViewWordResp
}
var file_job_proto_depIdxs = []int32{
	5, // 0: file.ViewJobsResp.Jobs:type_name -> file.Jobs
	8, // 1: file.ViewWordResp.Words:type_name -> file.Words
	0, // 2: file.job.CountJob:input_type -> file.CountJobReq
	2, // 3: file.job.JoinJob:input_type -> file.JoinJobReq
	4, // 4: file.job.ViewJobs:input_type -> file.ViewJobsReq
	7, // 5: file.job.ViewWord:input_type -> file.ViewWordReq
	1, // 6: file.job.CountJob:output_type -> file.CountJobResp
	3, // 7: file.job.JoinJob:output_type -> file.JoinJobResp
	6, // 8: file.job.ViewJobs:output_type -> file.ViewJobsResp
	9, // 9: file.job.ViewWord:output_type -> file.ViewWordResp
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_job_proto_init() }
func file_job_proto_init() {
	if File_job_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_job_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountJobReq); i {
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
		file_job_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountJobResp); i {
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
		file_job_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinJobReq); i {
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
		file_job_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinJobResp); i {
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
		file_job_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewJobsReq); i {
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
		file_job_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Jobs); i {
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
		file_job_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewJobsResp); i {
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
		file_job_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewWordReq); i {
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
		file_job_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Words); i {
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
		file_job_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewWordResp); i {
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
			RawDescriptor: file_job_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_job_proto_goTypes,
		DependencyIndexes: file_job_proto_depIdxs,
		MessageInfos:      file_job_proto_msgTypes,
	}.Build()
	File_job_proto = out.File
	file_job_proto_rawDesc = nil
	file_job_proto_goTypes = nil
	file_job_proto_depIdxs = nil
}
