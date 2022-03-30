// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/server/api/persistence/v1/tasks.proto

package persistence

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
	time "time"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	v1 "go.temporal.io/api/enums/v1"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// task column
type AllocatedTaskInfo struct {
	Data   *TaskInfo `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	TaskId int64     `protobuf:"varint,2,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
}

func (m *AllocatedTaskInfo) Reset()      { *m = AllocatedTaskInfo{} }
func (*AllocatedTaskInfo) ProtoMessage() {}
func (*AllocatedTaskInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9c734e3b35cf986, []int{0}
}
func (m *AllocatedTaskInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AllocatedTaskInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AllocatedTaskInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AllocatedTaskInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllocatedTaskInfo.Merge(m, src)
}
func (m *AllocatedTaskInfo) XXX_Size() int {
	return m.Size()
}
func (m *AllocatedTaskInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AllocatedTaskInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AllocatedTaskInfo proto.InternalMessageInfo

func (m *AllocatedTaskInfo) GetData() *TaskInfo {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *AllocatedTaskInfo) GetTaskId() int64 {
	if m != nil {
		return m.TaskId
	}
	return 0
}

type TaskInfo struct {
	NamespaceId string     `protobuf:"bytes,1,opt,name=namespace_id,json=namespaceId,proto3" json:"namespace_id,omitempty"`
	WorkflowId  string     `protobuf:"bytes,2,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id,omitempty"`
	RunId       string     `protobuf:"bytes,3,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
	ScheduleId  int64      `protobuf:"varint,4,opt,name=schedule_id,json=scheduleId,proto3" json:"schedule_id,omitempty"`
	CreateTime  *time.Time `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3,stdtime" json:"create_time,omitempty"`
	ExpiryTime  *time.Time `protobuf:"bytes,6,opt,name=expiry_time,json=expiryTime,proto3,stdtime" json:"expiry_time,omitempty"`
	Clock       int64      `protobuf:"varint,7,opt,name=clock,proto3" json:"clock,omitempty"`
}

func (m *TaskInfo) Reset()      { *m = TaskInfo{} }
func (*TaskInfo) ProtoMessage() {}
func (*TaskInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9c734e3b35cf986, []int{1}
}
func (m *TaskInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TaskInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TaskInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TaskInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskInfo.Merge(m, src)
}
func (m *TaskInfo) XXX_Size() int {
	return m.Size()
}
func (m *TaskInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TaskInfo proto.InternalMessageInfo

func (m *TaskInfo) GetNamespaceId() string {
	if m != nil {
		return m.NamespaceId
	}
	return ""
}

func (m *TaskInfo) GetWorkflowId() string {
	if m != nil {
		return m.WorkflowId
	}
	return ""
}

func (m *TaskInfo) GetRunId() string {
	if m != nil {
		return m.RunId
	}
	return ""
}

func (m *TaskInfo) GetScheduleId() int64 {
	if m != nil {
		return m.ScheduleId
	}
	return 0
}

func (m *TaskInfo) GetCreateTime() *time.Time {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *TaskInfo) GetExpiryTime() *time.Time {
	if m != nil {
		return m.ExpiryTime
	}
	return nil
}

func (m *TaskInfo) GetClock() int64 {
	if m != nil {
		return m.Clock
	}
	return 0
}

// task_queue column
type TaskQueueInfo struct {
	NamespaceId    string           `protobuf:"bytes,1,opt,name=namespace_id,json=namespaceId,proto3" json:"namespace_id,omitempty"`
	Name           string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	TaskType       v1.TaskQueueType `protobuf:"varint,3,opt,name=task_type,json=taskType,proto3,enum=temporal.api.enums.v1.TaskQueueType" json:"task_type,omitempty"`
	Kind           v1.TaskQueueKind `protobuf:"varint,4,opt,name=kind,proto3,enum=temporal.api.enums.v1.TaskQueueKind" json:"kind,omitempty"`
	AckLevel       int64            `protobuf:"varint,5,opt,name=ack_level,json=ackLevel,proto3" json:"ack_level,omitempty"`
	ExpiryTime     *time.Time       `protobuf:"bytes,6,opt,name=expiry_time,json=expiryTime,proto3,stdtime" json:"expiry_time,omitempty"`
	LastUpdateTime *time.Time       `protobuf:"bytes,7,opt,name=last_update_time,json=lastUpdateTime,proto3,stdtime" json:"last_update_time,omitempty"`
}

func (m *TaskQueueInfo) Reset()      { *m = TaskQueueInfo{} }
func (*TaskQueueInfo) ProtoMessage() {}
func (*TaskQueueInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9c734e3b35cf986, []int{2}
}
func (m *TaskQueueInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TaskQueueInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TaskQueueInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TaskQueueInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskQueueInfo.Merge(m, src)
}
func (m *TaskQueueInfo) XXX_Size() int {
	return m.Size()
}
func (m *TaskQueueInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskQueueInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TaskQueueInfo proto.InternalMessageInfo

func (m *TaskQueueInfo) GetNamespaceId() string {
	if m != nil {
		return m.NamespaceId
	}
	return ""
}

func (m *TaskQueueInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TaskQueueInfo) GetTaskType() v1.TaskQueueType {
	if m != nil {
		return m.TaskType
	}
	return v1.TASK_QUEUE_TYPE_UNSPECIFIED
}

func (m *TaskQueueInfo) GetKind() v1.TaskQueueKind {
	if m != nil {
		return m.Kind
	}
	return v1.TASK_QUEUE_KIND_UNSPECIFIED
}

func (m *TaskQueueInfo) GetAckLevel() int64 {
	if m != nil {
		return m.AckLevel
	}
	return 0
}

func (m *TaskQueueInfo) GetExpiryTime() *time.Time {
	if m != nil {
		return m.ExpiryTime
	}
	return nil
}

func (m *TaskQueueInfo) GetLastUpdateTime() *time.Time {
	if m != nil {
		return m.LastUpdateTime
	}
	return nil
}

func init() {
	proto.RegisterType((*AllocatedTaskInfo)(nil), "temporal.server.api.persistence.v1.AllocatedTaskInfo")
	proto.RegisterType((*TaskInfo)(nil), "temporal.server.api.persistence.v1.TaskInfo")
	proto.RegisterType((*TaskQueueInfo)(nil), "temporal.server.api.persistence.v1.TaskQueueInfo")
}

func init() {
	proto.RegisterFile("temporal/server/api/persistence/v1/tasks.proto", fileDescriptor_f9c734e3b35cf986)
}

var fileDescriptor_f9c734e3b35cf986 = []byte{
	// 549 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x3f, 0x6f, 0xd3, 0x40,
	0x1c, 0xb5, 0xf3, 0x3f, 0x17, 0xa8, 0xc0, 0x02, 0x11, 0x05, 0xe9, 0xd2, 0x46, 0x08, 0x75, 0x40,
	0x67, 0xb5, 0x30, 0x20, 0xb1, 0x90, 0x6e, 0x01, 0x16, 0xac, 0xb0, 0xb0, 0x44, 0x57, 0xdf, 0x2f,
	0xc1, 0xd8, 0xb9, 0x3b, 0x7c, 0xe7, 0x94, 0x6c, 0x7c, 0x84, 0x7e, 0x06, 0x26, 0x3e, 0x0a, 0x63,
	0xc6, 0x6e, 0x10, 0x67, 0x61, 0xa3, 0x1f, 0x01, 0xdd, 0xb9, 0x4e, 0xbb, 0x20, 0x32, 0xb0, 0xdd,
	0xef, 0xcf, 0x7b, 0xbf, 0x97, 0xf7, 0x22, 0x23, 0xa2, 0x61, 0x2e, 0x45, 0x4a, 0x13, 0x5f, 0x41,
	0xba, 0x80, 0xd4, 0xa7, 0x32, 0xf2, 0x25, 0xa4, 0x2a, 0x52, 0x1a, 0x78, 0x08, 0xfe, 0xe2, 0xc8,
	0xd7, 0x54, 0xc5, 0x8a, 0xc8, 0x54, 0x68, 0xe1, 0x0d, 0xca, 0x7d, 0x52, 0xec, 0x13, 0x2a, 0x23,
	0x72, 0x63, 0x9f, 0x2c, 0x8e, 0x7a, 0xfd, 0x99, 0x10, 0xb3, 0x04, 0x7c, 0x8b, 0x38, 0xcd, 0xa6,
	0xbe, 0x8e, 0xe6, 0xa0, 0x34, 0x9d, 0xcb, 0x82, 0xa4, 0x77, 0xc0, 0x40, 0x02, 0x67, 0xc0, 0xc3,
	0x08, 0x94, 0x3f, 0x13, 0x33, 0x61, 0xfb, 0xf6, 0x75, 0xb5, 0xf2, 0x78, 0xab, 0xcb, 0x08, 0x02,
	0x9e, 0xcd, 0x55, 0x29, 0x65, 0xf2, 0x29, 0x83, 0x0c, 0x8a, 0xbd, 0x01, 0x47, 0x77, 0x87, 0x49,
	0x22, 0x42, 0xaa, 0x81, 0x8d, 0xa9, 0x8a, 0x47, 0x7c, 0x2a, 0xbc, 0x97, 0xa8, 0xc6, 0xa8, 0xa6,
	0x5d, 0x77, 0xdf, 0x3d, 0xec, 0x1c, 0x3f, 0x21, 0xff, 0xd6, 0x4c, 0x4a, 0x6c, 0x60, 0x91, 0xde,
	0x03, 0xd4, 0xb4, 0xa7, 0x22, 0xd6, 0xad, 0xec, 0xbb, 0x87, 0xd5, 0xa0, 0x61, 0xca, 0x11, 0x1b,
	0x7c, 0xad, 0xa0, 0xd6, 0xf6, 0xce, 0x01, 0xba, 0xc5, 0xe9, 0x1c, 0x94, 0xa4, 0x21, 0x98, 0x55,
	0x73, 0xaf, 0x1d, 0x74, 0xb6, 0xbd, 0x11, 0xf3, 0xfa, 0xa8, 0x73, 0x26, 0xd2, 0x78, 0x9a, 0x88,
	0xb3, 0x92, 0xac, 0x1d, 0xa0, 0xb2, 0x35, 0x62, 0xde, 0x7d, 0xd4, 0x48, 0x33, 0x6e, 0x66, 0x55,
	0x3b, 0xab, 0xa7, 0x19, 0x2f, 0x70, 0x2a, 0xfc, 0x00, 0x2c, 0x4b, 0x2c, 0x73, 0xcd, 0x8a, 0x40,
	0x65, 0x6b, 0xc4, 0xbc, 0x21, 0xea, 0x84, 0x29, 0x50, 0x0d, 0x13, 0xe3, 0x6e, 0xb7, 0x6e, 0x7f,
	0x6a, 0x8f, 0x14, 0xd6, 0x93, 0xd2, 0x7a, 0x32, 0x2e, 0xad, 0x3f, 0xa9, 0x9d, 0xff, 0xe8, 0xbb,
	0x01, 0x2a, 0x40, 0xa6, 0x6d, 0x28, 0xe0, 0xb3, 0x8c, 0xd2, 0x65, 0x41, 0xd1, 0xd8, 0x95, 0xa2,
	0x00, 0x59, 0x8a, 0x7b, 0xa8, 0x1e, 0x26, 0x22, 0x8c, 0xbb, 0x4d, 0x2b, 0xb0, 0x28, 0x06, 0xbf,
	0x2b, 0xe8, 0xb6, 0x31, 0xe9, 0xad, 0x09, 0x6a, 0x57, 0xa7, 0x3c, 0x54, 0x33, 0xe5, 0x95, 0x45,
	0xf6, 0xed, 0x0d, 0x51, 0xdb, 0xc6, 0xa0, 0x97, 0x12, 0xac, 0x3f, 0x7b, 0xc7, 0x8f, 0xae, 0xd3,
	0x34, 0x31, 0xda, 0x7f, 0x46, 0x19, 0xa0, 0xbd, 0x37, 0x5e, 0x4a, 0x08, 0x5a, 0x06, 0x66, 0x5e,
	0xde, 0x73, 0x54, 0x8b, 0x23, 0x5e, 0x38, 0xb8, 0x03, 0xfa, 0x75, 0xc4, 0x59, 0x60, 0x11, 0xde,
	0x43, 0xd4, 0xa6, 0x61, 0x3c, 0x49, 0x60, 0x01, 0x89, 0xf5, 0xb7, 0x1a, 0xb4, 0x68, 0x18, 0xbf,
	0x31, 0xf5, 0xff, 0xf0, 0xee, 0x15, 0xba, 0x93, 0x50, 0xa5, 0x27, 0x99, 0x64, 0xdb, 0x18, 0x9b,
	0x3b, 0xf2, 0xec, 0x19, 0xe4, 0x3b, 0x0b, 0x34, 0xa3, 0x93, 0x8f, 0xab, 0x35, 0x76, 0x2e, 0xd6,
	0xd8, 0xb9, 0x5c, 0x63, 0xf7, 0x4b, 0x8e, 0xdd, 0x6f, 0x39, 0x76, 0xbf, 0xe7, 0xd8, 0x5d, 0xe5,
	0xd8, 0xfd, 0x99, 0x63, 0xf7, 0x57, 0x8e, 0x9d, 0xcb, 0x1c, 0xbb, 0xe7, 0x1b, 0xec, 0xac, 0x36,
	0xd8, 0xb9, 0xd8, 0x60, 0xe7, 0xfd, 0xb3, 0x99, 0xb8, 0xf6, 0x23, 0x12, 0x7f, 0xff, 0x04, 0xbc,
	0xb8, 0x51, 0x9e, 0x36, 0xac, 0xaa, 0xa7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x63, 0x5b, 0x1a,
	0x8f, 0x3b, 0x04, 0x00, 0x00,
}

func (this *AllocatedTaskInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AllocatedTaskInfo)
	if !ok {
		that2, ok := that.(AllocatedTaskInfo)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Data.Equal(that1.Data) {
		return false
	}
	if this.TaskId != that1.TaskId {
		return false
	}
	return true
}
func (this *TaskInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TaskInfo)
	if !ok {
		that2, ok := that.(TaskInfo)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.NamespaceId != that1.NamespaceId {
		return false
	}
	if this.WorkflowId != that1.WorkflowId {
		return false
	}
	if this.RunId != that1.RunId {
		return false
	}
	if this.ScheduleId != that1.ScheduleId {
		return false
	}
	if that1.CreateTime == nil {
		if this.CreateTime != nil {
			return false
		}
	} else if !this.CreateTime.Equal(*that1.CreateTime) {
		return false
	}
	if that1.ExpiryTime == nil {
		if this.ExpiryTime != nil {
			return false
		}
	} else if !this.ExpiryTime.Equal(*that1.ExpiryTime) {
		return false
	}
	if this.Clock != that1.Clock {
		return false
	}
	return true
}
func (this *TaskQueueInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TaskQueueInfo)
	if !ok {
		that2, ok := that.(TaskQueueInfo)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.NamespaceId != that1.NamespaceId {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.TaskType != that1.TaskType {
		return false
	}
	if this.Kind != that1.Kind {
		return false
	}
	if this.AckLevel != that1.AckLevel {
		return false
	}
	if that1.ExpiryTime == nil {
		if this.ExpiryTime != nil {
			return false
		}
	} else if !this.ExpiryTime.Equal(*that1.ExpiryTime) {
		return false
	}
	if that1.LastUpdateTime == nil {
		if this.LastUpdateTime != nil {
			return false
		}
	} else if !this.LastUpdateTime.Equal(*that1.LastUpdateTime) {
		return false
	}
	return true
}
func (this *AllocatedTaskInfo) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&persistence.AllocatedTaskInfo{")
	if this.Data != nil {
		s = append(s, "Data: "+fmt.Sprintf("%#v", this.Data)+",\n")
	}
	s = append(s, "TaskId: "+fmt.Sprintf("%#v", this.TaskId)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *TaskInfo) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 11)
	s = append(s, "&persistence.TaskInfo{")
	s = append(s, "NamespaceId: "+fmt.Sprintf("%#v", this.NamespaceId)+",\n")
	s = append(s, "WorkflowId: "+fmt.Sprintf("%#v", this.WorkflowId)+",\n")
	s = append(s, "RunId: "+fmt.Sprintf("%#v", this.RunId)+",\n")
	s = append(s, "ScheduleId: "+fmt.Sprintf("%#v", this.ScheduleId)+",\n")
	s = append(s, "CreateTime: "+fmt.Sprintf("%#v", this.CreateTime)+",\n")
	s = append(s, "ExpiryTime: "+fmt.Sprintf("%#v", this.ExpiryTime)+",\n")
	s = append(s, "Clock: "+fmt.Sprintf("%#v", this.Clock)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *TaskQueueInfo) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 11)
	s = append(s, "&persistence.TaskQueueInfo{")
	s = append(s, "NamespaceId: "+fmt.Sprintf("%#v", this.NamespaceId)+",\n")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "TaskType: "+fmt.Sprintf("%#v", this.TaskType)+",\n")
	s = append(s, "Kind: "+fmt.Sprintf("%#v", this.Kind)+",\n")
	s = append(s, "AckLevel: "+fmt.Sprintf("%#v", this.AckLevel)+",\n")
	s = append(s, "ExpiryTime: "+fmt.Sprintf("%#v", this.ExpiryTime)+",\n")
	s = append(s, "LastUpdateTime: "+fmt.Sprintf("%#v", this.LastUpdateTime)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringTasks(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *AllocatedTaskInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AllocatedTaskInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AllocatedTaskInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TaskId != 0 {
		i = encodeVarintTasks(dAtA, i, uint64(m.TaskId))
		i--
		dAtA[i] = 0x10
	}
	if m.Data != nil {
		{
			size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTasks(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TaskInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TaskInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TaskInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Clock != 0 {
		i = encodeVarintTasks(dAtA, i, uint64(m.Clock))
		i--
		dAtA[i] = 0x38
	}
	if m.ExpiryTime != nil {
		n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.ExpiryTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.ExpiryTime):])
		if err2 != nil {
			return 0, err2
		}
		i -= n2
		i = encodeVarintTasks(dAtA, i, uint64(n2))
		i--
		dAtA[i] = 0x32
	}
	if m.CreateTime != nil {
		n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.CreateTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.CreateTime):])
		if err3 != nil {
			return 0, err3
		}
		i -= n3
		i = encodeVarintTasks(dAtA, i, uint64(n3))
		i--
		dAtA[i] = 0x2a
	}
	if m.ScheduleId != 0 {
		i = encodeVarintTasks(dAtA, i, uint64(m.ScheduleId))
		i--
		dAtA[i] = 0x20
	}
	if len(m.RunId) > 0 {
		i -= len(m.RunId)
		copy(dAtA[i:], m.RunId)
		i = encodeVarintTasks(dAtA, i, uint64(len(m.RunId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.WorkflowId) > 0 {
		i -= len(m.WorkflowId)
		copy(dAtA[i:], m.WorkflowId)
		i = encodeVarintTasks(dAtA, i, uint64(len(m.WorkflowId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.NamespaceId) > 0 {
		i -= len(m.NamespaceId)
		copy(dAtA[i:], m.NamespaceId)
		i = encodeVarintTasks(dAtA, i, uint64(len(m.NamespaceId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TaskQueueInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TaskQueueInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TaskQueueInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LastUpdateTime != nil {
		n4, err4 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.LastUpdateTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.LastUpdateTime):])
		if err4 != nil {
			return 0, err4
		}
		i -= n4
		i = encodeVarintTasks(dAtA, i, uint64(n4))
		i--
		dAtA[i] = 0x3a
	}
	if m.ExpiryTime != nil {
		n5, err5 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.ExpiryTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.ExpiryTime):])
		if err5 != nil {
			return 0, err5
		}
		i -= n5
		i = encodeVarintTasks(dAtA, i, uint64(n5))
		i--
		dAtA[i] = 0x32
	}
	if m.AckLevel != 0 {
		i = encodeVarintTasks(dAtA, i, uint64(m.AckLevel))
		i--
		dAtA[i] = 0x28
	}
	if m.Kind != 0 {
		i = encodeVarintTasks(dAtA, i, uint64(m.Kind))
		i--
		dAtA[i] = 0x20
	}
	if m.TaskType != 0 {
		i = encodeVarintTasks(dAtA, i, uint64(m.TaskType))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTasks(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.NamespaceId) > 0 {
		i -= len(m.NamespaceId)
		copy(dAtA[i:], m.NamespaceId)
		i = encodeVarintTasks(dAtA, i, uint64(len(m.NamespaceId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTasks(dAtA []byte, offset int, v uint64) int {
	offset -= sovTasks(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AllocatedTaskInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovTasks(uint64(l))
	}
	if m.TaskId != 0 {
		n += 1 + sovTasks(uint64(m.TaskId))
	}
	return n
}

func (m *TaskInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.NamespaceId)
	if l > 0 {
		n += 1 + l + sovTasks(uint64(l))
	}
	l = len(m.WorkflowId)
	if l > 0 {
		n += 1 + l + sovTasks(uint64(l))
	}
	l = len(m.RunId)
	if l > 0 {
		n += 1 + l + sovTasks(uint64(l))
	}
	if m.ScheduleId != 0 {
		n += 1 + sovTasks(uint64(m.ScheduleId))
	}
	if m.CreateTime != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.CreateTime)
		n += 1 + l + sovTasks(uint64(l))
	}
	if m.ExpiryTime != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.ExpiryTime)
		n += 1 + l + sovTasks(uint64(l))
	}
	if m.Clock != 0 {
		n += 1 + sovTasks(uint64(m.Clock))
	}
	return n
}

func (m *TaskQueueInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.NamespaceId)
	if l > 0 {
		n += 1 + l + sovTasks(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTasks(uint64(l))
	}
	if m.TaskType != 0 {
		n += 1 + sovTasks(uint64(m.TaskType))
	}
	if m.Kind != 0 {
		n += 1 + sovTasks(uint64(m.Kind))
	}
	if m.AckLevel != 0 {
		n += 1 + sovTasks(uint64(m.AckLevel))
	}
	if m.ExpiryTime != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.ExpiryTime)
		n += 1 + l + sovTasks(uint64(l))
	}
	if m.LastUpdateTime != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.LastUpdateTime)
		n += 1 + l + sovTasks(uint64(l))
	}
	return n
}

func sovTasks(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTasks(x uint64) (n int) {
	return sovTasks(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *AllocatedTaskInfo) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AllocatedTaskInfo{`,
		`Data:` + strings.Replace(this.Data.String(), "TaskInfo", "TaskInfo", 1) + `,`,
		`TaskId:` + fmt.Sprintf("%v", this.TaskId) + `,`,
		`}`,
	}, "")
	return s
}
func (this *TaskInfo) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TaskInfo{`,
		`NamespaceId:` + fmt.Sprintf("%v", this.NamespaceId) + `,`,
		`WorkflowId:` + fmt.Sprintf("%v", this.WorkflowId) + `,`,
		`RunId:` + fmt.Sprintf("%v", this.RunId) + `,`,
		`ScheduleId:` + fmt.Sprintf("%v", this.ScheduleId) + `,`,
		`CreateTime:` + strings.Replace(fmt.Sprintf("%v", this.CreateTime), "Timestamp", "types.Timestamp", 1) + `,`,
		`ExpiryTime:` + strings.Replace(fmt.Sprintf("%v", this.ExpiryTime), "Timestamp", "types.Timestamp", 1) + `,`,
		`Clock:` + fmt.Sprintf("%v", this.Clock) + `,`,
		`}`,
	}, "")
	return s
}
func (this *TaskQueueInfo) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TaskQueueInfo{`,
		`NamespaceId:` + fmt.Sprintf("%v", this.NamespaceId) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`TaskType:` + fmt.Sprintf("%v", this.TaskType) + `,`,
		`Kind:` + fmt.Sprintf("%v", this.Kind) + `,`,
		`AckLevel:` + fmt.Sprintf("%v", this.AckLevel) + `,`,
		`ExpiryTime:` + strings.Replace(fmt.Sprintf("%v", this.ExpiryTime), "Timestamp", "types.Timestamp", 1) + `,`,
		`LastUpdateTime:` + strings.Replace(fmt.Sprintf("%v", this.LastUpdateTime), "Timestamp", "types.Timestamp", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringTasks(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *AllocatedTaskInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTasks
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AllocatedTaskInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AllocatedTaskInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTasks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTasks
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTasks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &TaskInfo{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaskId", wireType)
			}
			m.TaskId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTasks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TaskId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTasks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTasks
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTasks
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TaskInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTasks
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TaskInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TaskInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NamespaceId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTasks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTasks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTasks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NamespaceId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WorkflowId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTasks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTasks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTasks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WorkflowId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RunId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTasks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTasks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTasks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RunId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScheduleId", wireType)
			}
			m.ScheduleId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTasks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ScheduleId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTasks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTasks
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTasks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CreateTime == nil {
				m.CreateTime = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.CreateTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiryTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTasks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTasks
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTasks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ExpiryTime == nil {
				m.ExpiryTime = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.ExpiryTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Clock", wireType)
			}
			m.Clock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTasks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Clock |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTasks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTasks
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTasks
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TaskQueueInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTasks
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TaskQueueInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TaskQueueInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NamespaceId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTasks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTasks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTasks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NamespaceId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTasks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTasks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTasks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaskType", wireType)
			}
			m.TaskType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTasks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TaskType |= v1.TaskQueueType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kind", wireType)
			}
			m.Kind = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTasks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Kind |= v1.TaskQueueKind(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AckLevel", wireType)
			}
			m.AckLevel = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTasks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AckLevel |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiryTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTasks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTasks
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTasks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ExpiryTime == nil {
				m.ExpiryTime = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.ExpiryTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastUpdateTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTasks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTasks
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTasks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LastUpdateTime == nil {
				m.LastUpdateTime = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.LastUpdateTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTasks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTasks
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTasks
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTasks(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTasks
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTasks
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTasks
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTasks
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTasks
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTasks
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTasks        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTasks          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTasks = fmt.Errorf("proto: unexpected end of group")
)
