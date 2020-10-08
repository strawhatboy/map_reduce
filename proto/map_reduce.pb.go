// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: map_reduce.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type DataProvider int32

const (
	DataProvider_raw   DataProvider = 0
	DataProvider_audio DataProvider = 1
	DataProvider_csv   DataProvider = 2
)

// Enum value maps for DataProvider.
var (
	DataProvider_name = map[int32]string{
		0: "raw",
		1: "audio",
		2: "csv",
	}
	DataProvider_value = map[string]int32{
		"raw":   0,
		"audio": 1,
		"csv":   2,
	}
)

func (x DataProvider) Enum() *DataProvider {
	p := new(DataProvider)
	*p = x
	return p
}

func (x DataProvider) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataProvider) Descriptor() protoreflect.EnumDescriptor {
	return file_map_reduce_proto_enumTypes[0].Descriptor()
}

func (DataProvider) Type() protoreflect.EnumType {
	return &file_map_reduce_proto_enumTypes[0]
}

func (x DataProvider) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataProvider.Descriptor instead.
func (DataProvider) EnumDescriptor() ([]byte, []int) {
	return file_map_reduce_proto_rawDescGZIP(), []int{0}
}

type ClientStatus int32

const (
	ClientStatus_idle            ClientStatus = 0
	ClientStatus_working_mapper  ClientStatus = 1
	ClientStatus_working_reducer ClientStatus = 2
	ClientStatus_unknown         ClientStatus = 3
)

// Enum value maps for ClientStatus.
var (
	ClientStatus_name = map[int32]string{
		0: "idle",
		1: "working_mapper",
		2: "working_reducer",
		3: "unknown",
	}
	ClientStatus_value = map[string]int32{
		"idle":            0,
		"working_mapper":  1,
		"working_reducer": 2,
		"unknown":         3,
	}
)

func (x ClientStatus) Enum() *ClientStatus {
	p := new(ClientStatus)
	*p = x
	return p
}

func (x ClientStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClientStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_map_reduce_proto_enumTypes[1].Descriptor()
}

func (ClientStatus) Type() protoreflect.EnumType {
	return &file_map_reduce_proto_enumTypes[1]
}

func (x ClientStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClientStatus.Descriptor instead.
func (ClientStatus) EnumDescriptor() ([]byte, []int) {
	return file_map_reduce_proto_rawDescGZIP(), []int{1}
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_map_reduce_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_map_reduce_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_map_reduce_proto_rawDescGZIP(), []int{0}
}

type MapRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AssignedId   int64        `protobuf:"varint,1,opt,name=assigned_id,json=assignedId,proto3" json:"assigned_id,omitempty"`
	JobId        string       `protobuf:"bytes,2,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	InputFile    string       `protobuf:"bytes,3,opt,name=input_file,json=inputFile,proto3" json:"input_file,omitempty"`
	Script       string       `protobuf:"bytes,4,opt,name=script,proto3" json:"script,omitempty"`
	DataProvider DataProvider `protobuf:"varint,5,opt,name=data_provider,json=dataProvider,proto3,enum=DataProvider" json:"data_provider,omitempty"`
	ReducerCount int64        `protobuf:"varint,6,opt,name=reducer_count,json=reducerCount,proto3" json:"reducer_count,omitempty"`
}

func (x *MapRequest) Reset() {
	*x = MapRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_map_reduce_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapRequest) ProtoMessage() {}

func (x *MapRequest) ProtoReflect() protoreflect.Message {
	mi := &file_map_reduce_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapRequest.ProtoReflect.Descriptor instead.
func (*MapRequest) Descriptor() ([]byte, []int) {
	return file_map_reduce_proto_rawDescGZIP(), []int{1}
}

func (x *MapRequest) GetAssignedId() int64 {
	if x != nil {
		return x.AssignedId
	}
	return 0
}

func (x *MapRequest) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *MapRequest) GetInputFile() string {
	if x != nil {
		return x.InputFile
	}
	return ""
}

func (x *MapRequest) GetScript() string {
	if x != nil {
		return x.Script
	}
	return ""
}

func (x *MapRequest) GetDataProvider() DataProvider {
	if x != nil {
		return x.DataProvider
	}
	return DataProvider_raw
}

func (x *MapRequest) GetReducerCount() int64 {
	if x != nil {
		return x.ReducerCount
	}
	return 0
}

type ReduceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AssignedId int64  `protobuf:"varint,1,opt,name=assigned_id,json=assignedId,proto3" json:"assigned_id,omitempty"`
	JobId      string `protobuf:"bytes,2,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	InputFile  string `protobuf:"bytes,3,opt,name=input_file,json=inputFile,proto3" json:"input_file,omitempty"`
	Script     string `protobuf:"bytes,4,opt,name=script,proto3" json:"script,omitempty"`
}

func (x *ReduceRequest) Reset() {
	*x = ReduceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_map_reduce_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReduceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReduceRequest) ProtoMessage() {}

func (x *ReduceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_map_reduce_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReduceRequest.ProtoReflect.Descriptor instead.
func (*ReduceRequest) Descriptor() ([]byte, []int) {
	return file_map_reduce_proto_rawDescGZIP(), []int{2}
}

func (x *ReduceRequest) GetAssignedId() int64 {
	if x != nil {
		return x.AssignedId
	}
	return 0
}

func (x *ReduceRequest) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *ReduceRequest) GetInputFile() string {
	if x != nil {
		return x.InputFile
	}
	return ""
}

func (x *ReduceRequest) GetScript() string {
	if x != nil {
		return x.Script
	}
	return ""
}

type CommonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok  bool   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *CommonResponse) Reset() {
	*x = CommonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_map_reduce_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonResponse) ProtoMessage() {}

func (x *CommonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_map_reduce_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonResponse.ProtoReflect.Descriptor instead.
func (*CommonResponse) Descriptor() ([]byte, []int) {
	return file_map_reduce_proto_rawDescGZIP(), []int{3}
}

func (x *CommonResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *CommonResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

// status:
// 1: idle
// 2: working as mapper
// 3: working as reducer
// 4: unknown
type StatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status ClientStatus `protobuf:"varint,1,opt,name=status,proto3,enum=ClientStatus" json:"status,omitempty"`
}

func (x *StatusResponse) Reset() {
	*x = StatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_map_reduce_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResponse) ProtoMessage() {}

func (x *StatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_map_reduce_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResponse.ProtoReflect.Descriptor instead.
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return file_map_reduce_proto_rawDescGZIP(), []int{4}
}

func (x *StatusResponse) GetStatus() ClientStatus {
	if x != nil {
		return x.Status
	}
	return ClientStatus_idle
}

type MapPair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	First  string `protobuf:"bytes,1,opt,name=first,proto3" json:"first,omitempty"`
	Second int64  `protobuf:"varint,2,opt,name=second,proto3" json:"second,omitempty"`
}

func (x *MapPair) Reset() {
	*x = MapPair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_map_reduce_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapPair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapPair) ProtoMessage() {}

func (x *MapPair) ProtoReflect() protoreflect.Message {
	mi := &file_map_reduce_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapPair.ProtoReflect.Descriptor instead.
func (*MapPair) Descriptor() ([]byte, []int) {
	return file_map_reduce_proto_rawDescGZIP(), []int{5}
}

func (x *MapPair) GetFirst() string {
	if x != nil {
		return x.First
	}
	return ""
}

func (x *MapPair) GetSecond() int64 {
	if x != nil {
		return x.Second
	}
	return 0
}

type ReduceSliceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReduceId int64 `protobuf:"varint,1,opt,name=reduce_id,json=reduceId,proto3" json:"reduce_id,omitempty"` // start from 0
}

func (x *ReduceSliceRequest) Reset() {
	*x = ReduceSliceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_map_reduce_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReduceSliceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReduceSliceRequest) ProtoMessage() {}

func (x *ReduceSliceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_map_reduce_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReduceSliceRequest.ProtoReflect.Descriptor instead.
func (*ReduceSliceRequest) Descriptor() ([]byte, []int) {
	return file_map_reduce_proto_rawDescGZIP(), []int{6}
}

func (x *ReduceSliceRequest) GetReduceId() int64 {
	if x != nil {
		return x.ReduceId
	}
	return 0
}

type ReduceSliceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pairs []*MapPair `protobuf:"bytes,1,rep,name=pairs,proto3" json:"pairs,omitempty"`
}

func (x *ReduceSliceResponse) Reset() {
	*x = ReduceSliceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_map_reduce_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReduceSliceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReduceSliceResponse) ProtoMessage() {}

func (x *ReduceSliceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_map_reduce_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReduceSliceResponse.ProtoReflect.Descriptor instead.
func (*ReduceSliceResponse) Descriptor() ([]byte, []int) {
	return file_map_reduce_proto_rawDescGZIP(), []int{7}
}

func (x *ReduceSliceResponse) GetPairs() []*MapPair {
	if x != nil {
		return x.Pairs
	}
	return nil
}

type ResetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MapperCount  int64 `protobuf:"varint,1,opt,name=mapper_count,json=mapperCount,proto3" json:"mapper_count,omitempty"`
	ReducerCount int64 `protobuf:"varint,2,opt,name=reducer_count,json=reducerCount,proto3" json:"reducer_count,omitempty"`
}

func (x *ResetRequest) Reset() {
	*x = ResetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_map_reduce_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetRequest) ProtoMessage() {}

func (x *ResetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_map_reduce_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetRequest.ProtoReflect.Descriptor instead.
func (*ResetRequest) Descriptor() ([]byte, []int) {
	return file_map_reduce_proto_rawDescGZIP(), []int{8}
}

func (x *ResetRequest) GetMapperCount() int64 {
	if x != nil {
		return x.MapperCount
	}
	return 0
}

func (x *ResetRequest) GetReducerCount() int64 {
	if x != nil {
		return x.ReducerCount
	}
	return 0
}

type JobDoneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId           string `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	MapperReducerId int64  `protobuf:"varint,2,opt,name=mapper_reducer_id,json=mapperReducerId,proto3" json:"mapper_reducer_id,omitempty"` // start from 0
	ResultPath      string `protobuf:"bytes,3,opt,name=result_path,json=resultPath,proto3" json:"result_path,omitempty"`
}

func (x *JobDoneRequest) Reset() {
	*x = JobDoneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_map_reduce_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobDoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobDoneRequest) ProtoMessage() {}

func (x *JobDoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_map_reduce_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobDoneRequest.ProtoReflect.Descriptor instead.
func (*JobDoneRequest) Descriptor() ([]byte, []int) {
	return file_map_reduce_proto_rawDescGZIP(), []int{9}
}

func (x *JobDoneRequest) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *JobDoneRequest) GetMapperReducerId() int64 {
	if x != nil {
		return x.MapperReducerId
	}
	return 0
}

func (x *JobDoneRequest) GetResultPath() string {
	if x != nil {
		return x.ResultPath
	}
	return ""
}

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId       string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"` // guid
	ClientEndpoint string `protobuf:"bytes,2,opt,name=client_endpoint,json=clientEndpoint,proto3" json:"client_endpoint,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_map_reduce_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_map_reduce_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_map_reduce_proto_rawDescGZIP(), []int{10}
}

func (x *RegisterRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *RegisterRequest) GetClientEndpoint() string {
	if x != nil {
		return x.ClientEndpoint
	}
	return ""
}

var File_map_reduce_proto protoreflect.FileDescriptor

var file_map_reduce_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6d, 0x61, 0x70, 0x5f, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0xd4, 0x01, 0x0a, 0x0a,
	0x4d, 0x61, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x6a,
	0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x46, 0x69, 0x6c,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x12, 0x32, 0x0a, 0x0d, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0d, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52,
	0x0c, 0x64, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x23, 0x0a,
	0x0d, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x72, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x7e, 0x0a, 0x0d, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x64, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x22, 0x32, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x02, 0x6f, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x37, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x37, 0x0a, 0x07, 0x4d, 0x61, 0x70, 0x50, 0x61, 0x69, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x22, 0x31, 0x0a, 0x12, 0x52, 0x65, 0x64, 0x75,
	0x63, 0x65, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x49, 0x64, 0x22, 0x35, 0x0a, 0x13, 0x52,
	0x65, 0x64, 0x75, 0x63, 0x65, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1e, 0x0a, 0x05, 0x70, 0x61, 0x69, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x08, 0x2e, 0x4d, 0x61, 0x70, 0x50, 0x61, 0x69, 0x72, 0x52, 0x05, 0x70, 0x61, 0x69,
	0x72, 0x73, 0x22, 0x56, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x72,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x72, 0x65,
	0x64, 0x75, 0x63, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x74, 0x0a, 0x0e, 0x4a, 0x6f,
	0x62, 0x44, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06,
	0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f,
	0x62, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x72, 0x65,
	0x64, 0x75, 0x63, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f,
	0x6d, 0x61, 0x70, 0x70, 0x65, 0x72, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x50, 0x61, 0x74, 0x68,
	0x22, 0x57, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2a, 0x2b, 0x0a, 0x0c, 0x44, 0x61, 0x74,
	0x61, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x07, 0x0a, 0x03, 0x72, 0x61, 0x77,
	0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x10, 0x01, 0x12, 0x07, 0x0a,
	0x03, 0x63, 0x73, 0x76, 0x10, 0x02, 0x2a, 0x4e, 0x0a, 0x0c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x69, 0x64, 0x6c, 0x65, 0x10, 0x00,
	0x12, 0x12, 0x0a, 0x0e, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x5f,
	0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x72, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x75, 0x6e, 0x6b,
	0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x03, 0x32, 0xa5, 0x02, 0x0a, 0x0e, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x03, 0x6d, 0x61, 0x70,
	0x12, 0x0b, 0x2e, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x2b, 0x0a, 0x06, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x12, 0x0e, 0x2e, 0x52, 0x65, 0x64,
	0x75, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2e, 0x0a,
	0x08, 0x6d, 0x61, 0x70, 0x5f, 0x64, 0x6f, 0x6e, 0x65, 0x12, 0x0f, 0x2e, 0x4a, 0x6f, 0x62, 0x44,
	0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x23, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x0f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x3f, 0x0a, 0x10, 0x67, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65,
	0x5f, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x12, 0x13, 0x2e, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x53,
	0x6c, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x52, 0x65,
	0x64, 0x75, 0x63, 0x65, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x05, 0x72, 0x65, 0x73, 0x65, 0x74, 0x12, 0x0d, 0x2e, 0x52,
	0x65, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0xa4,
	0x01, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x2e, 0x0a, 0x08, 0x6d, 0x61, 0x70, 0x5f, 0x64, 0x6f, 0x6e, 0x65, 0x12, 0x0f, 0x2e,
	0x4a, 0x6f, 0x62, 0x44, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x31, 0x0a, 0x0b, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x5f, 0x64, 0x6f, 0x6e, 0x65,
	0x12, 0x0f, 0x2e, 0x4a, 0x6f, 0x62, 0x44, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x12, 0x10, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x61, 0x77, 0x68, 0x61, 0x74, 0x62, 0x6f, 0x79, 0x2f,
	0x6d, 0x61, 0x70, 0x5f, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_map_reduce_proto_rawDescOnce sync.Once
	file_map_reduce_proto_rawDescData = file_map_reduce_proto_rawDesc
)

func file_map_reduce_proto_rawDescGZIP() []byte {
	file_map_reduce_proto_rawDescOnce.Do(func() {
		file_map_reduce_proto_rawDescData = protoimpl.X.CompressGZIP(file_map_reduce_proto_rawDescData)
	})
	return file_map_reduce_proto_rawDescData
}

var file_map_reduce_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_map_reduce_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_map_reduce_proto_goTypes = []interface{}{
	(DataProvider)(0),           // 0: DataProvider
	(ClientStatus)(0),           // 1: ClientStatus
	(*Empty)(nil),               // 2: Empty
	(*MapRequest)(nil),          // 3: MapRequest
	(*ReduceRequest)(nil),       // 4: ReduceRequest
	(*CommonResponse)(nil),      // 5: CommonResponse
	(*StatusResponse)(nil),      // 6: StatusResponse
	(*MapPair)(nil),             // 7: MapPair
	(*ReduceSliceRequest)(nil),  // 8: ReduceSliceRequest
	(*ReduceSliceResponse)(nil), // 9: ReduceSliceResponse
	(*ResetRequest)(nil),        // 10: ResetRequest
	(*JobDoneRequest)(nil),      // 11: JobDoneRequest
	(*RegisterRequest)(nil),     // 12: RegisterRequest
}
var file_map_reduce_proto_depIdxs = []int32{
	0,  // 0: MapRequest.data_provider:type_name -> DataProvider
	1,  // 1: StatusResponse.status:type_name -> ClientStatus
	7,  // 2: ReduceSliceResponse.pairs:type_name -> MapPair
	3,  // 3: Client_Service.map:input_type -> MapRequest
	4,  // 4: Client_Service.reduce:input_type -> ReduceRequest
	11, // 5: Client_Service.map_done:input_type -> JobDoneRequest
	2,  // 6: Client_Service.status:input_type -> Empty
	8,  // 7: Client_Service.get_reduce_slice:input_type -> ReduceSliceRequest
	10, // 8: Client_Service.reset:input_type -> ResetRequest
	11, // 9: Server_Service.map_done:input_type -> JobDoneRequest
	11, // 10: Server_Service.reduce_done:input_type -> JobDoneRequest
	12, // 11: Server_Service.register:input_type -> RegisterRequest
	5,  // 12: Client_Service.map:output_type -> CommonResponse
	5,  // 13: Client_Service.reduce:output_type -> CommonResponse
	5,  // 14: Client_Service.map_done:output_type -> CommonResponse
	6,  // 15: Client_Service.status:output_type -> StatusResponse
	9,  // 16: Client_Service.get_reduce_slice:output_type -> ReduceSliceResponse
	5,  // 17: Client_Service.reset:output_type -> CommonResponse
	5,  // 18: Server_Service.map_done:output_type -> CommonResponse
	5,  // 19: Server_Service.reduce_done:output_type -> CommonResponse
	5,  // 20: Server_Service.register:output_type -> CommonResponse
	12, // [12:21] is the sub-list for method output_type
	3,  // [3:12] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_map_reduce_proto_init() }
func file_map_reduce_proto_init() {
	if File_map_reduce_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_map_reduce_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_map_reduce_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapRequest); i {
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
		file_map_reduce_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReduceRequest); i {
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
		file_map_reduce_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonResponse); i {
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
		file_map_reduce_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusResponse); i {
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
		file_map_reduce_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapPair); i {
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
		file_map_reduce_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReduceSliceRequest); i {
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
		file_map_reduce_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReduceSliceResponse); i {
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
		file_map_reduce_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetRequest); i {
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
		file_map_reduce_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobDoneRequest); i {
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
		file_map_reduce_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
			RawDescriptor: file_map_reduce_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_map_reduce_proto_goTypes,
		DependencyIndexes: file_map_reduce_proto_depIdxs,
		EnumInfos:         file_map_reduce_proto_enumTypes,
		MessageInfos:      file_map_reduce_proto_msgTypes,
	}.Build()
	File_map_reduce_proto = out.File
	file_map_reduce_proto_rawDesc = nil
	file_map_reduce_proto_goTypes = nil
	file_map_reduce_proto_depIdxs = nil
}
