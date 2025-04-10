// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: grpc.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// 状态码
type TaskResponse_Status int32

const (
	TaskResponse_SUCCESS TaskResponse_Status = 0
	TaskResponse_FAILED  TaskResponse_Status = 1
)

// Enum value maps for TaskResponse_Status.
var (
	TaskResponse_Status_name = map[int32]string{
		0: "SUCCESS",
		1: "FAILED",
	}
	TaskResponse_Status_value = map[string]int32{
		"SUCCESS": 0,
		"FAILED":  1,
	}
)

func (x TaskResponse_Status) Enum() *TaskResponse_Status {
	p := new(TaskResponse_Status)
	*p = x
	return p
}

func (x TaskResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_grpc_proto_enumTypes[0].Descriptor()
}

func (TaskResponse_Status) Type() protoreflect.EnumType {
	return &file_grpc_proto_enumTypes[0]
}

func (x TaskResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskResponse_Status.Descriptor instead.
func (TaskResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{2, 0}
}

type Record struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AgentId       string                 `protobuf:"bytes,1,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	AgentVersion  string                 `protobuf:"bytes,2,opt,name=agent_version,json=agentVersion,proto3" json:"agent_version,omitempty"`
	Hostname      string                 `protobuf:"bytes,3,opt,name=hostname,proto3" json:"hostname,omitempty"`
	ClousterName  string                 `protobuf:"bytes,4,opt,name=clouster_name,json=clousterName,proto3" json:"clouster_name,omitempty"`
	NodeType      string                 `protobuf:"bytes,5,opt,name=node_type,json=nodeType,proto3" json:"node_type,omitempty"`
	Environment   []string               `protobuf:"bytes,6,rep,name=environment,proto3" json:"environment,omitempty"`
	PrivateIpv4   []string               `protobuf:"bytes,7,rep,name=private_ipv4,json=privateIpv4,proto3" json:"private_ipv4,omitempty"`
	PublicIpv4    []string               `protobuf:"bytes,8,rep,name=public_ipv4,json=publicIpv4,proto3" json:"public_ipv4,omitempty"`
	Platform      string                 `protobuf:"bytes,9,opt,name=platform,proto3" json:"platform,omitempty"`
	KernelVersion string                 `protobuf:"bytes,10,opt,name=kernel_version,json=kernelVersion,proto3" json:"kernel_version,omitempty"`
	Data          string                 `protobuf:"bytes,11,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Record) Reset() {
	*x = Record{}
	mi := &file_grpc_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record.ProtoReflect.Descriptor instead.
func (*Record) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *Record) GetAgentId() string {
	if x != nil {
		return x.AgentId
	}
	return ""
}

func (x *Record) GetAgentVersion() string {
	if x != nil {
		return x.AgentVersion
	}
	return ""
}

func (x *Record) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *Record) GetClousterName() string {
	if x != nil {
		return x.ClousterName
	}
	return ""
}

func (x *Record) GetNodeType() string {
	if x != nil {
		return x.NodeType
	}
	return ""
}

func (x *Record) GetEnvironment() []string {
	if x != nil {
		return x.Environment
	}
	return nil
}

func (x *Record) GetPrivateIpv4() []string {
	if x != nil {
		return x.PrivateIpv4
	}
	return nil
}

func (x *Record) GetPublicIpv4() []string {
	if x != nil {
		return x.PublicIpv4
	}
	return nil
}

func (x *Record) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *Record) GetKernelVersion() string {
	if x != nil {
		return x.KernelVersion
	}
	return ""
}

func (x *Record) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type Task struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TaskId        string                 `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`                                                                                       // 任务ID
	JobId         string                 `protobuf:"bytes,2,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`                                                                                          // 任务ID
	TaskStatus    string                 `protobuf:"bytes,3,opt,name=task_status,json=taskStatus,proto3" json:"task_status,omitempty"`                                                                           // 任务状态
	CreateTime    string                 `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`                                                                           // 任务执行时间
	AgentId       string                 `protobuf:"bytes,5,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`                                                                                    // agent ID
	TaskType      string                 `protobuf:"bytes,6,opt,name=task_type,json=taskType,proto3" json:"task_type,omitempty"`                                                                                 // 任务类型
	TaskParams    map[string]string      `protobuf:"bytes,7,rep,name=task_params,json=taskParams,proto3" json:"task_params,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"` // 任务参数
	Report        []byte                 `protobuf:"bytes,8,opt,name=report,proto3" json:"report,omitempty"`                                                                                                     // 任务结果
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Task) Reset() {
	*x = Task{}
	mi := &file_grpc_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[1]
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
	return file_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *Task) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *Task) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *Task) GetTaskStatus() string {
	if x != nil {
		return x.TaskStatus
	}
	return ""
}

func (x *Task) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *Task) GetAgentId() string {
	if x != nil {
		return x.AgentId
	}
	return ""
}

func (x *Task) GetTaskType() string {
	if x != nil {
		return x.TaskType
	}
	return ""
}

func (x *Task) GetTaskParams() map[string]string {
	if x != nil {
		return x.TaskParams
	}
	return nil
}

func (x *Task) GetReport() []byte {
	if x != nil {
		return x.Report
	}
	return nil
}

// server --> agent,上传结果响应
type TaskResponse struct {
	state    protoimpl.MessageState `protogen:"open.v1"`
	TaskId   string                 `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`                                                             // 任务ID
	TaskType string                 `protobuf:"bytes,2,opt,name=task_type,json=taskType,proto3" json:"task_type,omitempty"`                                                       // 任务类型
	Params   map[string]string      `protobuf:"bytes,3,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"` // 任务参数
	Status   TaskResponse_Status    `protobuf:"varint,4,opt,name=status,proto3,enum=filetransfer.TaskResponse_Status" json:"status,omitempty"`
	// 可选的错误信息
	Message       string `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TaskResponse) Reset() {
	*x = TaskResponse{}
	mi := &file_grpc_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskResponse) ProtoMessage() {}

func (x *TaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskResponse.ProtoReflect.Descriptor instead.
func (*TaskResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{2}
}

func (x *TaskResponse) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *TaskResponse) GetTaskType() string {
	if x != nil {
		return x.TaskType
	}
	return ""
}

func (x *TaskResponse) GetParams() map[string]string {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *TaskResponse) GetStatus() TaskResponse_Status {
	if x != nil {
		return x.Status
	}
	return TaskResponse_SUCCESS
}

func (x *TaskResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_grpc_proto protoreflect.FileDescriptor

const file_grpc_proto_rawDesc = "" +
	"\n" +
	"\n" +
	"grpc.proto\x12\ffiletransfer\"\xe3\x02\n" +
	"\x06Record\x12\x19\n" +
	"\bagent_id\x18\x01 \x01(\tR\aagentId\x12#\n" +
	"\ragent_version\x18\x02 \x01(\tR\fagentVersion\x12\x1a\n" +
	"\bhostname\x18\x03 \x01(\tR\bhostname\x12#\n" +
	"\rclouster_name\x18\x04 \x01(\tR\fclousterName\x12\x1b\n" +
	"\tnode_type\x18\x05 \x01(\tR\bnodeType\x12 \n" +
	"\venvironment\x18\x06 \x03(\tR\venvironment\x12!\n" +
	"\fprivate_ipv4\x18\a \x03(\tR\vprivateIpv4\x12\x1f\n" +
	"\vpublic_ipv4\x18\b \x03(\tR\n" +
	"publicIpv4\x12\x1a\n" +
	"\bplatform\x18\t \x01(\tR\bplatform\x12%\n" +
	"\x0ekernel_version\x18\n" +
	" \x01(\tR\rkernelVersion\x12\x12\n" +
	"\x04data\x18\v \x01(\tR\x04data\"\xcc\x02\n" +
	"\x04Task\x12\x17\n" +
	"\atask_id\x18\x01 \x01(\tR\x06taskId\x12\x15\n" +
	"\x06job_id\x18\x02 \x01(\tR\x05jobId\x12\x1f\n" +
	"\vtask_status\x18\x03 \x01(\tR\n" +
	"taskStatus\x12\x1f\n" +
	"\vcreate_time\x18\x04 \x01(\tR\n" +
	"createTime\x12\x19\n" +
	"\bagent_id\x18\x05 \x01(\tR\aagentId\x12\x1b\n" +
	"\ttask_type\x18\x06 \x01(\tR\btaskType\x12C\n" +
	"\vtask_params\x18\a \x03(\v2\".filetransfer.Task.TaskParamsEntryR\n" +
	"taskParams\x12\x16\n" +
	"\x06report\x18\b \x01(\fR\x06report\x1a=\n" +
	"\x0fTaskParamsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"\xb7\x02\n" +
	"\fTaskResponse\x12\x17\n" +
	"\atask_id\x18\x01 \x01(\tR\x06taskId\x12\x1b\n" +
	"\ttask_type\x18\x02 \x01(\tR\btaskType\x12>\n" +
	"\x06params\x18\x03 \x03(\v2&.filetransfer.TaskResponse.ParamsEntryR\x06params\x129\n" +
	"\x06status\x18\x04 \x01(\x0e2!.filetransfer.TaskResponse.StatusR\x06status\x12\x18\n" +
	"\amessage\x18\x05 \x01(\tR\amessage\x1a9\n" +
	"\vParamsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"!\n" +
	"\x06Status\x12\v\n" +
	"\aSUCCESS\x10\x00\x12\n" +
	"\n" +
	"\x06FAILED\x10\x012J\n" +
	"\fTaskTransfer\x12:\n" +
	"\x06Upload\x12\x12.filetransfer.Task\x1a\x1a.filetransfer.TaskResponse(\x012J\n" +
	"\x0eRecordTransfer\x128\n" +
	"\bTransfer\x12\x14.filetransfer.Record\x1a\x12.filetransfer.Task(\x010\x01B+Z)github.com/colin-404/cloudog-common/protob\x06proto3"

var (
	file_grpc_proto_rawDescOnce sync.Once
	file_grpc_proto_rawDescData []byte
)

func file_grpc_proto_rawDescGZIP() []byte {
	file_grpc_proto_rawDescOnce.Do(func() {
		file_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_grpc_proto_rawDesc), len(file_grpc_proto_rawDesc)))
	})
	return file_grpc_proto_rawDescData
}

var file_grpc_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_grpc_proto_goTypes = []any{
	(TaskResponse_Status)(0), // 0: filetransfer.TaskResponse.Status
	(*Record)(nil),           // 1: filetransfer.Record
	(*Task)(nil),             // 2: filetransfer.Task
	(*TaskResponse)(nil),     // 3: filetransfer.TaskResponse
	nil,                      // 4: filetransfer.Task.TaskParamsEntry
	nil,                      // 5: filetransfer.TaskResponse.ParamsEntry
}
var file_grpc_proto_depIdxs = []int32{
	4, // 0: filetransfer.Task.task_params:type_name -> filetransfer.Task.TaskParamsEntry
	5, // 1: filetransfer.TaskResponse.params:type_name -> filetransfer.TaskResponse.ParamsEntry
	0, // 2: filetransfer.TaskResponse.status:type_name -> filetransfer.TaskResponse.Status
	2, // 3: filetransfer.TaskTransfer.Upload:input_type -> filetransfer.Task
	1, // 4: filetransfer.RecordTransfer.Transfer:input_type -> filetransfer.Record
	3, // 5: filetransfer.TaskTransfer.Upload:output_type -> filetransfer.TaskResponse
	2, // 6: filetransfer.RecordTransfer.Transfer:output_type -> filetransfer.Task
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_grpc_proto_init() }
func file_grpc_proto_init() {
	if File_grpc_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_grpc_proto_rawDesc), len(file_grpc_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_grpc_proto_goTypes,
		DependencyIndexes: file_grpc_proto_depIdxs,
		EnumInfos:         file_grpc_proto_enumTypes,
		MessageInfos:      file_grpc_proto_msgTypes,
	}.Build()
	File_grpc_proto = out.File
	file_grpc_proto_goTypes = nil
	file_grpc_proto_depIdxs = nil
}
