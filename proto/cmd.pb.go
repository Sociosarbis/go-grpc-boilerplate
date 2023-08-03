// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: proto/cmd.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Cmd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Script string `protobuf:"bytes,1,opt,name=script,proto3" json:"script,omitempty"`
	Wd     string `protobuf:"bytes,2,opt,name=wd,proto3" json:"wd,omitempty"`
}

func (x *Cmd) Reset() {
	*x = Cmd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cmd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cmd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cmd) ProtoMessage() {}

func (x *Cmd) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cmd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cmd.ProtoReflect.Descriptor instead.
func (*Cmd) Descriptor() ([]byte, []int) {
	return file_proto_cmd_proto_rawDescGZIP(), []int{0}
}

func (x *Cmd) GetScript() string {
	if x != nil {
		return x.Script
	}
	return ""
}

func (x *Cmd) GetWd() string {
	if x != nil {
		return x.Wd
	}
	return ""
}

type CmdCallRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Output string `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
	Type   int32  `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *CmdCallRes) Reset() {
	*x = CmdCallRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cmd_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdCallRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdCallRes) ProtoMessage() {}

func (x *CmdCallRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cmd_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdCallRes.ProtoReflect.Descriptor instead.
func (*CmdCallRes) Descriptor() ([]byte, []int) {
	return file_proto_cmd_proto_rawDescGZIP(), []int{1}
}

func (x *CmdCallRes) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

func (x *CmdCallRes) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

type CmdListFolderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Folder string `protobuf:"bytes,1,opt,name=folder,proto3" json:"folder,omitempty"`
}

func (x *CmdListFolderReq) Reset() {
	*x = CmdListFolderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cmd_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdListFolderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdListFolderReq) ProtoMessage() {}

func (x *CmdListFolderReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cmd_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdListFolderReq.ProtoReflect.Descriptor instead.
func (*CmdListFolderReq) Descriptor() ([]byte, []int) {
	return file_proto_cmd_proto_rawDescGZIP(), []int{2}
}

func (x *CmdListFolderReq) GetFolder() string {
	if x != nil {
		return x.Folder
	}
	return ""
}

type CmdItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *CmdItem) Reset() {
	*x = CmdItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cmd_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdItem) ProtoMessage() {}

func (x *CmdItem) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cmd_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdItem.ProtoReflect.Descriptor instead.
func (*CmdItem) Descriptor() ([]byte, []int) {
	return file_proto_cmd_proto_rawDescGZIP(), []int{3}
}

func (x *CmdItem) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CmdItem) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint32     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Items []*CmdItem `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cmd_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cmd_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_proto_cmd_proto_rawDescGZIP(), []int{4}
}

func (x *Command) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Command) GetItems() []*CmdItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type CmdAddReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*CmdItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *CmdAddReq) Reset() {
	*x = CmdAddReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cmd_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdAddReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdAddReq) ProtoMessage() {}

func (x *CmdAddReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cmd_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdAddReq.ProtoReflect.Descriptor instead.
func (*CmdAddReq) Descriptor() ([]byte, []int) {
	return file_proto_cmd_proto_rawDescGZIP(), []int{5}
}

func (x *CmdAddReq) GetItems() []*CmdItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type CmdAddRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CmdAddRes) Reset() {
	*x = CmdAddRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cmd_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdAddRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdAddRes) ProtoMessage() {}

func (x *CmdAddRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cmd_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdAddRes.ProtoReflect.Descriptor instead.
func (*CmdAddRes) Descriptor() ([]byte, []int) {
	return file_proto_cmd_proto_rawDescGZIP(), []int{6}
}

func (x *CmdAddRes) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CmdDeleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CmdDeleteReq) Reset() {
	*x = CmdDeleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cmd_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdDeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdDeleteReq) ProtoMessage() {}

func (x *CmdDeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cmd_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdDeleteReq.ProtoReflect.Descriptor instead.
func (*CmdDeleteReq) Descriptor() ([]byte, []int) {
	return file_proto_cmd_proto_rawDescGZIP(), []int{7}
}

func (x *CmdDeleteReq) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CmdUpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint32     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Items []*CmdItem `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *CmdUpdateReq) Reset() {
	*x = CmdUpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cmd_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdUpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdUpdateReq) ProtoMessage() {}

func (x *CmdUpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cmd_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdUpdateReq.ProtoReflect.Descriptor instead.
func (*CmdUpdateReq) Descriptor() ([]byte, []int) {
	return file_proto_cmd_proto_rawDescGZIP(), []int{8}
}

func (x *CmdUpdateReq) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CmdUpdateReq) GetItems() []*CmdItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type CmdListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page uint32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size uint32 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *CmdListReq) Reset() {
	*x = CmdListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cmd_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdListReq) ProtoMessage() {}

func (x *CmdListReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cmd_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdListReq.ProtoReflect.Descriptor instead.
func (*CmdListReq) Descriptor() ([]byte, []int) {
	return file_proto_cmd_proto_rawDescGZIP(), []int{9}
}

func (x *CmdListReq) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *CmdListReq) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type CmdListRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count uint32     `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Items []*Command `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *CmdListRes) Reset() {
	*x = CmdListRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cmd_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdListRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdListRes) ProtoMessage() {}

func (x *CmdListRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cmd_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdListRes.ProtoReflect.Descriptor instead.
func (*CmdListRes) Descriptor() ([]byte, []int) {
	return file_proto_cmd_proto_rawDescGZIP(), []int{10}
}

func (x *CmdListRes) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *CmdListRes) GetItems() []*Command {
	if x != nil {
		return x.Items
	}
	return nil
}

type FolderItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	IsFolder bool   `protobuf:"varint,2,opt,name=is_folder,json=isFolder,proto3" json:"is_folder,omitempty"`
}

func (x *FolderItem) Reset() {
	*x = FolderItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cmd_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FolderItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FolderItem) ProtoMessage() {}

func (x *FolderItem) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cmd_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FolderItem.ProtoReflect.Descriptor instead.
func (*FolderItem) Descriptor() ([]byte, []int) {
	return file_proto_cmd_proto_rawDescGZIP(), []int{11}
}

func (x *FolderItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FolderItem) GetIsFolder() bool {
	if x != nil {
		return x.IsFolder
	}
	return false
}

type CmdListFolderRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*FolderItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *CmdListFolderRes) Reset() {
	*x = CmdListFolderRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cmd_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdListFolderRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdListFolderRes) ProtoMessage() {}

func (x *CmdListFolderRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cmd_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdListFolderRes.ProtoReflect.Descriptor instead.
func (*CmdListFolderRes) Descriptor() ([]byte, []int) {
	return file_proto_cmd_proto_rawDescGZIP(), []int{12}
}

func (x *CmdListFolderRes) GetItems() []*FolderItem {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_proto_cmd_proto protoreflect.FileDescriptor

var file_proto_cmd_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6d, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x03, 0x43, 0x6d, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x77, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x77, 0x64, 0x22, 0x38, 0x0a, 0x0a, 0x43, 0x6d, 0x64, 0x43, 0x61, 0x6c, 0x6c, 0x52,
	0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x2a,
	0x0a, 0x10, 0x43, 0x6d, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x22, 0x33, 0x0a, 0x07, 0x43, 0x6d,
	0x64, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x3f, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x6d, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x22, 0x31, 0x0a, 0x09, 0x43, 0x6d, 0x64, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x12, 0x24, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6d, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x22, 0x1b, 0x0a, 0x09, 0x43, 0x6d, 0x64, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x1e, 0x0a, 0x0c, 0x43, 0x6d, 0x64, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x44, 0x0a, 0x0c, 0x43, 0x6d, 0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x24, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6d, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x34, 0x0a, 0x0a, 0x43, 0x6d, 0x64, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x48, 0x0a, 0x0a,
	0x43, 0x6d, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x24, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x3d, 0x0a, 0x0a, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x66,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x46,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x22, 0x3b, 0x0a, 0x10, 0x43, 0x6d, 0x64, 0x4c, 0x69, 0x73, 0x74,
	0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x32, 0xce, 0x02, 0x0a, 0x0a, 0x43, 0x6d, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x2a, 0x0a, 0x07, 0x43, 0x6d, 0x64, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x0a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6d, 0x64, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x6d, 0x64, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x30, 0x01, 0x12, 0x41, 0x0a,
	0x0d, 0x43, 0x6d, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x17,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6d, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x6d, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x12, 0x2c, 0x0a, 0x06, 0x43, 0x6d, 0x64, 0x41, 0x64, 0x64, 0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x6d, 0x64, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6d, 0x64, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x12, 0x2f,
	0x0a, 0x07, 0x43, 0x6d, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x6d, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6d, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x12,
	0x38, 0x0a, 0x09, 0x43, 0x6d, 0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6d, 0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x38, 0x0a, 0x09, 0x43, 0x6d, 0x64,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x6d, 0x64, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_cmd_proto_rawDescOnce sync.Once
	file_proto_cmd_proto_rawDescData = file_proto_cmd_proto_rawDesc
)

func file_proto_cmd_proto_rawDescGZIP() []byte {
	file_proto_cmd_proto_rawDescOnce.Do(func() {
		file_proto_cmd_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_cmd_proto_rawDescData)
	})
	return file_proto_cmd_proto_rawDescData
}

var file_proto_cmd_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_proto_cmd_proto_goTypes = []interface{}{
	(*Cmd)(nil),              // 0: proto.Cmd
	(*CmdCallRes)(nil),       // 1: proto.CmdCallRes
	(*CmdListFolderReq)(nil), // 2: proto.CmdListFolderReq
	(*CmdItem)(nil),          // 3: proto.CmdItem
	(*Command)(nil),          // 4: proto.Command
	(*CmdAddReq)(nil),        // 5: proto.CmdAddReq
	(*CmdAddRes)(nil),        // 6: proto.CmdAddRes
	(*CmdDeleteReq)(nil),     // 7: proto.CmdDeleteReq
	(*CmdUpdateReq)(nil),     // 8: proto.CmdUpdateReq
	(*CmdListReq)(nil),       // 9: proto.CmdListReq
	(*CmdListRes)(nil),       // 10: proto.CmdListRes
	(*FolderItem)(nil),       // 11: proto.FolderItem
	(*CmdListFolderRes)(nil), // 12: proto.CmdListFolderRes
	(*emptypb.Empty)(nil),    // 13: google.protobuf.Empty
}
var file_proto_cmd_proto_depIdxs = []int32{
	3,  // 0: proto.Command.items:type_name -> proto.CmdItem
	3,  // 1: proto.CmdAddReq.items:type_name -> proto.CmdItem
	3,  // 2: proto.CmdUpdateReq.items:type_name -> proto.CmdItem
	4,  // 3: proto.CmdListRes.items:type_name -> proto.Command
	11, // 4: proto.CmdListFolderRes.items:type_name -> proto.FolderItem
	0,  // 5: proto.CmdService.CmdCall:input_type -> proto.Cmd
	2,  // 6: proto.CmdService.CmdListFolder:input_type -> proto.CmdListFolderReq
	5,  // 7: proto.CmdService.CmdAdd:input_type -> proto.CmdAddReq
	9,  // 8: proto.CmdService.CmdList:input_type -> proto.CmdListReq
	8,  // 9: proto.CmdService.CmdUpdate:input_type -> proto.CmdUpdateReq
	7,  // 10: proto.CmdService.CmdDelete:input_type -> proto.CmdDeleteReq
	1,  // 11: proto.CmdService.CmdCall:output_type -> proto.CmdCallRes
	12, // 12: proto.CmdService.CmdListFolder:output_type -> proto.CmdListFolderRes
	6,  // 13: proto.CmdService.CmdAdd:output_type -> proto.CmdAddRes
	10, // 14: proto.CmdService.CmdList:output_type -> proto.CmdListRes
	13, // 15: proto.CmdService.CmdUpdate:output_type -> google.protobuf.Empty
	13, // 16: proto.CmdService.CmdDelete:output_type -> google.protobuf.Empty
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_proto_cmd_proto_init() }
func file_proto_cmd_proto_init() {
	if File_proto_cmd_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_cmd_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cmd); i {
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
		file_proto_cmd_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmdCallRes); i {
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
		file_proto_cmd_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmdListFolderReq); i {
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
		file_proto_cmd_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmdItem); i {
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
		file_proto_cmd_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command); i {
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
		file_proto_cmd_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmdAddReq); i {
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
		file_proto_cmd_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmdAddRes); i {
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
		file_proto_cmd_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmdDeleteReq); i {
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
		file_proto_cmd_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmdUpdateReq); i {
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
		file_proto_cmd_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmdListReq); i {
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
		file_proto_cmd_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmdListRes); i {
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
		file_proto_cmd_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FolderItem); i {
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
		file_proto_cmd_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmdListFolderRes); i {
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
			RawDescriptor: file_proto_cmd_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_cmd_proto_goTypes,
		DependencyIndexes: file_proto_cmd_proto_depIdxs,
		MessageInfos:      file_proto_cmd_proto_msgTypes,
	}.Build()
	File_proto_cmd_proto = out.File
	file_proto_cmd_proto_rawDesc = nil
	file_proto_cmd_proto_goTypes = nil
	file_proto_cmd_proto_depIdxs = nil
}
