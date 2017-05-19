// Code generated by protoc-gen-go.
// source: job_desc.proto
// DO NOT EDIT!

package firmament

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type JobDescriptor_JobState int32

const (
	JobDescriptor_NEW       JobDescriptor_JobState = 0
	JobDescriptor_CREATED   JobDescriptor_JobState = 1
	JobDescriptor_RUNNING   JobDescriptor_JobState = 2
	JobDescriptor_COMPLETED JobDescriptor_JobState = 3
	JobDescriptor_FAILED    JobDescriptor_JobState = 4
	JobDescriptor_ABORTED   JobDescriptor_JobState = 5
	JobDescriptor_UNKNOWN   JobDescriptor_JobState = 6
)

var JobDescriptor_JobState_name = map[int32]string{
	0: "NEW",
	1: "CREATED",
	2: "RUNNING",
	3: "COMPLETED",
	4: "FAILED",
	5: "ABORTED",
	6: "UNKNOWN",
}
var JobDescriptor_JobState_value = map[string]int32{
	"NEW":       0,
	"CREATED":   1,
	"RUNNING":   2,
	"COMPLETED": 3,
	"FAILED":    4,
	"ABORTED":   5,
	"UNKNOWN":   6,
}

func (x JobDescriptor_JobState) String() string {
	return proto.EnumName(JobDescriptor_JobState_name, int32(x))
}
func (JobDescriptor_JobState) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0, 0} }

type JobDescriptor struct {
	Uuid      string                 `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	Name      string                 `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	State     JobDescriptor_JobState `protobuf:"varint,3,opt,name=state,enum=firmament.JobDescriptor_JobState" json:"state,omitempty"`
	RootTask  *TaskDescriptor        `protobuf:"bytes,4,opt,name=root_task,json=rootTask" json:"root_task,omitempty"`
	OutputIds [][]byte               `protobuf:"bytes,5,rep,name=output_ids,json=outputIds,proto3" json:"output_ids,omitempty"`
}

func (m *JobDescriptor) Reset()                    { *m = JobDescriptor{} }
func (m *JobDescriptor) String() string            { return proto.CompactTextString(m) }
func (*JobDescriptor) ProtoMessage()               {}
func (*JobDescriptor) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *JobDescriptor) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *JobDescriptor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *JobDescriptor) GetState() JobDescriptor_JobState {
	if m != nil {
		return m.State
	}
	return JobDescriptor_NEW
}

func (m *JobDescriptor) GetRootTask() *TaskDescriptor {
	if m != nil {
		return m.RootTask
	}
	return nil
}

func (m *JobDescriptor) GetOutputIds() [][]byte {
	if m != nil {
		return m.OutputIds
	}
	return nil
}

func init() {
	proto.RegisterType((*JobDescriptor)(nil), "firmament.JobDescriptor")
	proto.RegisterEnum("firmament.JobDescriptor_JobState", JobDescriptor_JobState_name, JobDescriptor_JobState_value)
}

func init() { proto.RegisterFile("job_desc.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x90, 0x4f, 0x4f, 0x83, 0x40,
	0x10, 0xc5, 0xe5, 0x6f, 0xcb, 0xd4, 0x56, 0xb2, 0x27, 0x34, 0x31, 0xc1, 0x9e, 0x38, 0x71, 0xa8,
	0x89, 0x9e, 0xb1, 0xa0, 0x41, 0xeb, 0x62, 0xd6, 0x36, 0x3d, 0x12, 0x28, 0x98, 0x60, 0x43, 0x97,
	0xec, 0x2e, 0x9f, 0xca, 0x2f, 0x69, 0xa6, 0x44, 0xad, 0xb7, 0x79, 0xef, 0xfd, 0x66, 0xdf, 0x66,
	0x60, 0xf6, 0xc9, 0xcb, 0xbc, 0xaa, 0xe5, 0x2e, 0xec, 0x04, 0x57, 0x9c, 0x38, 0x1f, 0x8d, 0x68,
	0x8b, 0xb6, 0x3e, 0xa8, 0xab, 0x0b, 0x55, 0xc8, 0xfd, 0x49, 0x36, 0xff, 0xd2, 0x61, 0xfa, 0xcc,
	0xcb, 0xb8, 0x96, 0x3b, 0xd1, 0x74, 0x8a, 0x0b, 0x42, 0xc0, 0xec, 0xfb, 0xa6, 0xf2, 0x34, 0x5f,
	0x0b, 0x1c, 0x76, 0x9c, 0xd1, 0x3b, 0x14, 0x6d, 0xed, 0xe9, 0x83, 0x87, 0x33, 0xb9, 0x07, 0x4b,
	0xaa, 0x42, 0xd5, 0x9e, 0xe1, 0x6b, 0xc1, 0x6c, 0x71, 0x13, 0xfe, 0xb6, 0x84, 0xff, 0x1e, 0x44,
	0xf5, 0x8e, 0x20, 0x1b, 0x78, 0x72, 0x07, 0x8e, 0xe0, 0x5c, 0xe5, 0xf8, 0x15, 0xcf, 0xf4, 0xb5,
	0x60, 0xb2, 0xb8, 0x3c, 0x59, 0x5e, 0x17, 0x72, 0xff, 0xb7, 0xcd, 0xc6, 0xc8, 0xa2, 0x47, 0xae,
	0x01, 0x78, 0xaf, 0xba, 0x5e, 0xe5, 0x4d, 0x25, 0x3d, 0xcb, 0x37, 0x82, 0x73, 0xe6, 0x0c, 0x4e,
	0x5a, 0xc9, 0x79, 0x09, 0xe3, 0x9f, 0x26, 0x32, 0x02, 0x83, 0x26, 0x5b, 0xf7, 0x8c, 0x4c, 0x60,
	0xb4, 0x64, 0x49, 0xb4, 0x4e, 0x62, 0x57, 0x43, 0xc1, 0x36, 0x94, 0xa6, 0xf4, 0xc9, 0xd5, 0xc9,
	0x14, 0x9c, 0x65, 0xf6, 0xfa, 0xb6, 0x4a, 0x30, 0x33, 0x08, 0x80, 0xfd, 0x18, 0xa5, 0xab, 0x24,
	0x76, 0x4d, 0xe4, 0xa2, 0x87, 0x8c, 0x61, 0x60, 0xa1, 0xd8, 0xd0, 0x17, 0x9a, 0x6d, 0xa9, 0x6b,
	0x97, 0xf6, 0xf1, 0x68, 0xb7, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbf, 0xd7, 0x73, 0x42, 0x62,
	0x01, 0x00, 0x00,
}
