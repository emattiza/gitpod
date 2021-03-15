// Copyright (c) 2021 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.5
// source: workspace.proto

package api

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

type PrepareForUserNSRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PrepareForUserNSRequest) Reset() {
	*x = PrepareForUserNSRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workspace_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrepareForUserNSRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrepareForUserNSRequest) ProtoMessage() {}

func (x *PrepareForUserNSRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workspace_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrepareForUserNSRequest.ProtoReflect.Descriptor instead.
func (*PrepareForUserNSRequest) Descriptor() ([]byte, []int) {
	return file_workspace_proto_rawDescGZIP(), []int{0}
}

type PrepareForUserNSResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PrepareForUserNSResponse) Reset() {
	*x = PrepareForUserNSResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workspace_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrepareForUserNSResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrepareForUserNSResponse) ProtoMessage() {}

func (x *PrepareForUserNSResponse) ProtoReflect() protoreflect.Message {
	mi := &file_workspace_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrepareForUserNSResponse.ProtoReflect.Descriptor instead.
func (*PrepareForUserNSResponse) Descriptor() ([]byte, []int) {
	return file_workspace_proto_rawDescGZIP(), []int{1}
}

type WriteIDMappingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message   string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	ErrorCode uint32 `protobuf:"varint,2,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
}

func (x *WriteIDMappingResponse) Reset() {
	*x = WriteIDMappingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workspace_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteIDMappingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteIDMappingResponse) ProtoMessage() {}

func (x *WriteIDMappingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_workspace_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteIDMappingResponse.ProtoReflect.Descriptor instead.
func (*WriteIDMappingResponse) Descriptor() ([]byte, []int) {
	return file_workspace_proto_rawDescGZIP(), []int{2}
}

func (x *WriteIDMappingResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *WriteIDMappingResponse) GetErrorCode() uint32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

type WriteIDMappingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid     int64                            `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	Gid     bool                             `protobuf:"varint,2,opt,name=gid,proto3" json:"gid,omitempty"`
	Mapping []*WriteIDMappingRequest_Mapping `protobuf:"bytes,3,rep,name=mapping,proto3" json:"mapping,omitempty"`
}

func (x *WriteIDMappingRequest) Reset() {
	*x = WriteIDMappingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workspace_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteIDMappingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteIDMappingRequest) ProtoMessage() {}

func (x *WriteIDMappingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workspace_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteIDMappingRequest.ProtoReflect.Descriptor instead.
func (*WriteIDMappingRequest) Descriptor() ([]byte, []int) {
	return file_workspace_proto_rawDescGZIP(), []int{3}
}

func (x *WriteIDMappingRequest) GetPid() int64 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *WriteIDMappingRequest) GetGid() bool {
	if x != nil {
		return x.Gid
	}
	return false
}

func (x *WriteIDMappingRequest) GetMapping() []*WriteIDMappingRequest_Mapping {
	if x != nil {
		return x.Mapping
	}
	return nil
}

type MountProcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target string `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	Pid    int64  `protobuf:"varint,2,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (x *MountProcRequest) Reset() {
	*x = MountProcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workspace_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MountProcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MountProcRequest) ProtoMessage() {}

func (x *MountProcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workspace_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MountProcRequest.ProtoReflect.Descriptor instead.
func (*MountProcRequest) Descriptor() ([]byte, []int) {
	return file_workspace_proto_rawDescGZIP(), []int{4}
}

func (x *MountProcRequest) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *MountProcRequest) GetPid() int64 {
	if x != nil {
		return x.Pid
	}
	return 0
}

type MountProcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MountProcResponse) Reset() {
	*x = MountProcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workspace_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MountProcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MountProcResponse) ProtoMessage() {}

func (x *MountProcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_workspace_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MountProcResponse.ProtoReflect.Descriptor instead.
func (*MountProcResponse) Descriptor() ([]byte, []int) {
	return file_workspace_proto_rawDescGZIP(), []int{5}
}

type UmountProcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target string `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	Pid    int64  `protobuf:"varint,2,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (x *UmountProcRequest) Reset() {
	*x = UmountProcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workspace_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UmountProcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UmountProcRequest) ProtoMessage() {}

func (x *UmountProcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workspace_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UmountProcRequest.ProtoReflect.Descriptor instead.
func (*UmountProcRequest) Descriptor() ([]byte, []int) {
	return file_workspace_proto_rawDescGZIP(), []int{6}
}

func (x *UmountProcRequest) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *UmountProcRequest) GetPid() int64 {
	if x != nil {
		return x.Pid
	}
	return 0
}

type UmountProcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UmountProcResponse) Reset() {
	*x = UmountProcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workspace_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UmountProcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UmountProcResponse) ProtoMessage() {}

func (x *UmountProcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_workspace_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UmountProcResponse.ProtoReflect.Descriptor instead.
func (*UmountProcResponse) Descriptor() ([]byte, []int) {
	return file_workspace_proto_rawDescGZIP(), []int{7}
}

type TeardownRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TeardownRequest) Reset() {
	*x = TeardownRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workspace_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeardownRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeardownRequest) ProtoMessage() {}

func (x *TeardownRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workspace_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeardownRequest.ProtoReflect.Descriptor instead.
func (*TeardownRequest) Descriptor() ([]byte, []int) {
	return file_workspace_proto_rawDescGZIP(), []int{8}
}

type TeardownResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *TeardownResponse) Reset() {
	*x = TeardownResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workspace_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeardownResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeardownResponse) ProtoMessage() {}

func (x *TeardownResponse) ProtoReflect() protoreflect.Message {
	mi := &file_workspace_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeardownResponse.ProtoReflect.Descriptor instead.
func (*TeardownResponse) Descriptor() ([]byte, []int) {
	return file_workspace_proto_rawDescGZIP(), []int{9}
}

func (x *TeardownResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type WriteIDMappingRequest_Mapping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContainerId uint32 `protobuf:"varint,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	HostId      uint32 `protobuf:"varint,2,opt,name=host_id,json=hostId,proto3" json:"host_id,omitempty"`
	Size        uint32 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *WriteIDMappingRequest_Mapping) Reset() {
	*x = WriteIDMappingRequest_Mapping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workspace_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteIDMappingRequest_Mapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteIDMappingRequest_Mapping) ProtoMessage() {}

func (x *WriteIDMappingRequest_Mapping) ProtoReflect() protoreflect.Message {
	mi := &file_workspace_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteIDMappingRequest_Mapping.ProtoReflect.Descriptor instead.
func (*WriteIDMappingRequest_Mapping) Descriptor() ([]byte, []int) {
	return file_workspace_proto_rawDescGZIP(), []int{3, 0}
}

func (x *WriteIDMappingRequest_Mapping) GetContainerId() uint32 {
	if x != nil {
		return x.ContainerId
	}
	return 0
}

func (x *WriteIDMappingRequest_Mapping) GetHostId() uint32 {
	if x != nil {
		return x.HostId
	}
	return 0
}

func (x *WriteIDMappingRequest_Mapping) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

var File_workspace_proto protoreflect.FileDescriptor

var file_workspace_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x69, 0x77, 0x73, 0x22, 0x19, 0x0a, 0x17, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72,
	0x65, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x53, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x1a, 0x0a, 0x18, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x46, 0x6f, 0x72, 0x55,
	0x73, 0x65, 0x72, 0x4e, 0x53, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x51, 0x0a,
	0x16, 0x57, 0x72, 0x69, 0x74, 0x65, 0x49, 0x44, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x22, 0xd4, 0x01, 0x0a, 0x15, 0x57, 0x72, 0x69, 0x74, 0x65, 0x49, 0x44, 0x4d, 0x61, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x67, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x67, 0x69, 0x64, 0x12, 0x3c,
	0x0a, 0x07, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x69, 0x77, 0x73, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x49, 0x44, 0x4d, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x61, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x52, 0x07, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x1a, 0x59, 0x0a, 0x07,
	0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x68, 0x6f,
	0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x68, 0x6f, 0x73,
	0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x3c, 0x0a, 0x10, 0x4d, 0x6f, 0x75, 0x6e, 0x74,
	0x50, 0x72, 0x6f, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x70, 0x69, 0x64, 0x22, 0x13, 0x0a, 0x11, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72,
	0x6f, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3d, 0x0a, 0x11, 0x55, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x70, 0x69, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x55, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x11, 0x0a, 0x0f, 0x54, 0x65, 0x61, 0x72, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x2c, 0x0a, 0x10, 0x54, 0x65, 0x61, 0x72, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x32, 0xee, 0x02, 0x0a, 0x12, 0x49, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x10, 0x50, 0x72, 0x65, 0x70, 0x61,
	0x72, 0x65, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x53, 0x12, 0x1c, 0x2e, 0x69, 0x77,
	0x73, 0x2e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72,
	0x4e, 0x53, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x69, 0x77, 0x73, 0x2e,
	0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x53,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0e, 0x57, 0x72,
	0x69, 0x74, 0x65, 0x49, 0x44, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x2e, 0x69,
	0x77, 0x73, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x49, 0x44, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x69, 0x77, 0x73, 0x2e, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x49, 0x44, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x09, 0x4d, 0x6f, 0x75, 0x6e, 0x74,
	0x50, 0x72, 0x6f, 0x63, 0x12, 0x15, 0x2e, 0x69, 0x77, 0x73, 0x2e, 0x4d, 0x6f, 0x75, 0x6e, 0x74,
	0x50, 0x72, 0x6f, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x69, 0x77,
	0x73, 0x2e, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0a, 0x55, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x50,
	0x72, 0x6f, 0x63, 0x12, 0x16, 0x2e, 0x69, 0x77, 0x73, 0x2e, 0x55, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x50, 0x72, 0x6f, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x69, 0x77,
	0x73, 0x2e, 0x55, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x08, 0x54, 0x65, 0x61, 0x72, 0x64, 0x6f,
	0x77, 0x6e, 0x12, 0x14, 0x2e, 0x69, 0x77, 0x73, 0x2e, 0x54, 0x65, 0x61, 0x72, 0x64, 0x6f, 0x77,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x69, 0x77, 0x73, 0x2e, 0x54,
	0x65, 0x61, 0x72, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64,
	0x2f, 0x77, 0x73, 0x2d, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_workspace_proto_rawDescOnce sync.Once
	file_workspace_proto_rawDescData = file_workspace_proto_rawDesc
)

func file_workspace_proto_rawDescGZIP() []byte {
	file_workspace_proto_rawDescOnce.Do(func() {
		file_workspace_proto_rawDescData = protoimpl.X.CompressGZIP(file_workspace_proto_rawDescData)
	})
	return file_workspace_proto_rawDescData
}

var file_workspace_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_workspace_proto_goTypes = []interface{}{
	(*PrepareForUserNSRequest)(nil),       // 0: iws.PrepareForUserNSRequest
	(*PrepareForUserNSResponse)(nil),      // 1: iws.PrepareForUserNSResponse
	(*WriteIDMappingResponse)(nil),        // 2: iws.WriteIDMappingResponse
	(*WriteIDMappingRequest)(nil),         // 3: iws.WriteIDMappingRequest
	(*MountProcRequest)(nil),              // 4: iws.MountProcRequest
	(*MountProcResponse)(nil),             // 5: iws.MountProcResponse
	(*UmountProcRequest)(nil),             // 6: iws.UmountProcRequest
	(*UmountProcResponse)(nil),            // 7: iws.UmountProcResponse
	(*TeardownRequest)(nil),               // 8: iws.TeardownRequest
	(*TeardownResponse)(nil),              // 9: iws.TeardownResponse
	(*WriteIDMappingRequest_Mapping)(nil), // 10: iws.WriteIDMappingRequest.Mapping
}
var file_workspace_proto_depIdxs = []int32{
	10, // 0: iws.WriteIDMappingRequest.mapping:type_name -> iws.WriteIDMappingRequest.Mapping
	0,  // 1: iws.InWorkspaceService.PrepareForUserNS:input_type -> iws.PrepareForUserNSRequest
	3,  // 2: iws.InWorkspaceService.WriteIDMapping:input_type -> iws.WriteIDMappingRequest
	4,  // 3: iws.InWorkspaceService.MountProc:input_type -> iws.MountProcRequest
	6,  // 4: iws.InWorkspaceService.UmountProc:input_type -> iws.UmountProcRequest
	8,  // 5: iws.InWorkspaceService.Teardown:input_type -> iws.TeardownRequest
	1,  // 6: iws.InWorkspaceService.PrepareForUserNS:output_type -> iws.PrepareForUserNSResponse
	2,  // 7: iws.InWorkspaceService.WriteIDMapping:output_type -> iws.WriteIDMappingResponse
	5,  // 8: iws.InWorkspaceService.MountProc:output_type -> iws.MountProcResponse
	7,  // 9: iws.InWorkspaceService.UmountProc:output_type -> iws.UmountProcResponse
	9,  // 10: iws.InWorkspaceService.Teardown:output_type -> iws.TeardownResponse
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_workspace_proto_init() }
func file_workspace_proto_init() {
	if File_workspace_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_workspace_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrepareForUserNSRequest); i {
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
		file_workspace_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrepareForUserNSResponse); i {
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
		file_workspace_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteIDMappingResponse); i {
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
		file_workspace_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteIDMappingRequest); i {
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
		file_workspace_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MountProcRequest); i {
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
		file_workspace_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MountProcResponse); i {
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
		file_workspace_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UmountProcRequest); i {
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
		file_workspace_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UmountProcResponse); i {
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
		file_workspace_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeardownRequest); i {
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
		file_workspace_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeardownResponse); i {
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
		file_workspace_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteIDMappingRequest_Mapping); i {
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
			RawDescriptor: file_workspace_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_workspace_proto_goTypes,
		DependencyIndexes: file_workspace_proto_depIdxs,
		MessageInfos:      file_workspace_proto_msgTypes,
	}.Build()
	File_workspace_proto = out.File
	file_workspace_proto_rawDesc = nil
	file_workspace_proto_goTypes = nil
	file_workspace_proto_depIdxs = nil
}