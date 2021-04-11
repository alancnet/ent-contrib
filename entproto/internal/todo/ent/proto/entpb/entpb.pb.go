// Code generated by entproto. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: entpb/entpb.proto

package entpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Todo_Status int32

const (
	Todo_PENDING     Todo_Status = 0
	Todo_IN_PROGRESS Todo_Status = 1
	Todo_DONE        Todo_Status = 2
)

// Enum value maps for Todo_Status.
var (
	Todo_Status_name = map[int32]string{
		0: "PENDING",
		1: "IN_PROGRESS",
		2: "DONE",
	}
	Todo_Status_value = map[string]int32{
		"PENDING":     0,
		"IN_PROGRESS": 1,
		"DONE":        2,
	}
)

func (x Todo_Status) Enum() *Todo_Status {
	p := new(Todo_Status)
	*p = x
	return p
}

func (x Todo_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Todo_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_entpb_entpb_proto_enumTypes[0].Descriptor()
}

func (Todo_Status) Type() protoreflect.EnumType {
	return &file_entpb_entpb_proto_enumTypes[0]
}

func (x Todo_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Todo_Status.Descriptor instead.
func (Todo_Status) EnumDescriptor() ([]byte, []int) {
	return file_entpb_entpb_proto_rawDescGZIP(), []int{6, 0}
}

type User_Status int32

const (
	User_STATUS_UNSPECIFIED User_Status = 0
	User_PENDING            User_Status = 1
	User_ACTIVE             User_Status = 2
)

// Enum value maps for User_Status.
var (
	User_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "PENDING",
		2: "ACTIVE",
	}
	User_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"PENDING":            1,
		"ACTIVE":             2,
	}
)

func (x User_Status) Enum() *User_Status {
	p := new(User_Status)
	*p = x
	return p
}

func (x User_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (User_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_entpb_entpb_proto_enumTypes[1].Descriptor()
}

func (User_Status) Type() protoreflect.EnumType {
	return &file_entpb_entpb_proto_enumTypes[1]
}

func (x User_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use User_Status.Descriptor instead.
func (User_Status) EnumDescriptor() ([]byte, []int) {
	return file_entpb_entpb_proto_rawDescGZIP(), []int{7, 0}
}

type Attachment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	User *User  `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *Attachment) Reset() {
	*x = Attachment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entpb_entpb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attachment) ProtoMessage() {}

func (x *Attachment) ProtoReflect() protoreflect.Message {
	mi := &file_entpb_entpb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attachment.ProtoReflect.Descriptor instead.
func (*Attachment) Descriptor() ([]byte, []int) {
	return file_entpb_entpb_proto_rawDescGZIP(), []int{0}
}

func (x *Attachment) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Attachment) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type CreateAttachmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Attachment *Attachment `protobuf:"bytes,1,opt,name=attachment,proto3" json:"attachment,omitempty"`
}

func (x *CreateAttachmentRequest) Reset() {
	*x = CreateAttachmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entpb_entpb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAttachmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAttachmentRequest) ProtoMessage() {}

func (x *CreateAttachmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_entpb_entpb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAttachmentRequest.ProtoReflect.Descriptor instead.
func (*CreateAttachmentRequest) Descriptor() ([]byte, []int) {
	return file_entpb_entpb_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAttachmentRequest) GetAttachment() *Attachment {
	if x != nil {
		return x.Attachment
	}
	return nil
}

type GetAttachmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetAttachmentRequest) Reset() {
	*x = GetAttachmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entpb_entpb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAttachmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAttachmentRequest) ProtoMessage() {}

func (x *GetAttachmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_entpb_entpb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAttachmentRequest.ProtoReflect.Descriptor instead.
func (*GetAttachmentRequest) Descriptor() ([]byte, []int) {
	return file_entpb_entpb_proto_rawDescGZIP(), []int{2}
}

func (x *GetAttachmentRequest) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

type UpdateAttachmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Attachment *Attachment `protobuf:"bytes,1,opt,name=attachment,proto3" json:"attachment,omitempty"`
}

func (x *UpdateAttachmentRequest) Reset() {
	*x = UpdateAttachmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entpb_entpb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAttachmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAttachmentRequest) ProtoMessage() {}

func (x *UpdateAttachmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_entpb_entpb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAttachmentRequest.ProtoReflect.Descriptor instead.
func (*UpdateAttachmentRequest) Descriptor() ([]byte, []int) {
	return file_entpb_entpb_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateAttachmentRequest) GetAttachment() *Attachment {
	if x != nil {
		return x.Attachment
	}
	return nil
}

type DeleteAttachmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteAttachmentRequest) Reset() {
	*x = DeleteAttachmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entpb_entpb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAttachmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAttachmentRequest) ProtoMessage() {}

func (x *DeleteAttachmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_entpb_entpb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAttachmentRequest.ProtoReflect.Descriptor instead.
func (*DeleteAttachmentRequest) Descriptor() ([]byte, []int) {
	return file_entpb_entpb_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteAttachmentRequest) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Users []*User `protobuf:"bytes,3,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entpb_entpb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_entpb_entpb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_entpb_entpb_proto_rawDescGZIP(), []int{5}
}

func (x *Group) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Group) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Group) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type Todo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Task   string      `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
	Status Todo_Status `protobuf:"varint,3,opt,name=status,proto3,enum=entpb.Todo_Status" json:"status,omitempty"`
	User   *User       `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *Todo) Reset() {
	*x = Todo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entpb_entpb_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Todo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Todo) ProtoMessage() {}

func (x *Todo) ProtoReflect() protoreflect.Message {
	mi := &file_entpb_entpb_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Todo.ProtoReflect.Descriptor instead.
func (*Todo) Descriptor() ([]byte, []int) {
	return file_entpb_entpb_proto_rawDescGZIP(), []int{6}
}

func (x *Todo) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Todo) GetTask() string {
	if x != nil {
		return x.Task
	}
	return ""
}

func (x *Todo) GetStatus() Todo_Status {
	if x != nil {
		return x.Status
	}
	return Todo_PENDING
}

func (x *Todo) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserName   string                 `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Joined     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=joined,proto3" json:"joined,omitempty"`
	Points     uint32                 `protobuf:"varint,4,opt,name=points,proto3" json:"points,omitempty"`
	Exp        uint64                 `protobuf:"varint,5,opt,name=exp,proto3" json:"exp,omitempty"`
	Status     User_Status            `protobuf:"varint,6,opt,name=status,proto3,enum=entpb.User_Status" json:"status,omitempty"`
	ExternalId int32                  `protobuf:"varint,8,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
	CrmId      []byte                 `protobuf:"bytes,9,opt,name=crm_id,json=crmId,proto3" json:"crm_id,omitempty"`
	Banned     bool                   `protobuf:"varint,10,opt,name=banned,proto3" json:"banned,omitempty"`
	CustomPb   uint64                 `protobuf:"varint,12,opt,name=custom_pb,json=customPb,proto3" json:"custom_pb,omitempty"`
	Group      *Group                 `protobuf:"bytes,7,opt,name=group,proto3" json:"group,omitempty"`
	Attachment *Attachment            `protobuf:"bytes,11,opt,name=attachment,proto3" json:"attachment,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entpb_entpb_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_entpb_entpb_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_entpb_entpb_proto_rawDescGZIP(), []int{7}
}

func (x *User) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *User) GetJoined() *timestamppb.Timestamp {
	if x != nil {
		return x.Joined
	}
	return nil
}

func (x *User) GetPoints() uint32 {
	if x != nil {
		return x.Points
	}
	return 0
}

func (x *User) GetExp() uint64 {
	if x != nil {
		return x.Exp
	}
	return 0
}

func (x *User) GetStatus() User_Status {
	if x != nil {
		return x.Status
	}
	return User_STATUS_UNSPECIFIED
}

func (x *User) GetExternalId() int32 {
	if x != nil {
		return x.ExternalId
	}
	return 0
}

func (x *User) GetCrmId() []byte {
	if x != nil {
		return x.CrmId
	}
	return nil
}

func (x *User) GetBanned() bool {
	if x != nil {
		return x.Banned
	}
	return false
}

func (x *User) GetCustomPb() uint64 {
	if x != nil {
		return x.CustomPb
	}
	return 0
}

func (x *User) GetGroup() *Group {
	if x != nil {
		return x.Group
	}
	return nil
}

func (x *User) GetAttachment() *Attachment {
	if x != nil {
		return x.Attachment
	}
	return nil
}

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entpb_entpb_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_entpb_entpb_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_entpb_entpb_proto_rawDescGZIP(), []int{8}
}

func (x *CreateUserRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type GetUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entpb_entpb_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_entpb_entpb_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_entpb_entpb_proto_rawDescGZIP(), []int{9}
}

func (x *GetUserRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UpdateUserRequest) Reset() {
	*x = UpdateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entpb_entpb_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserRequest) ProtoMessage() {}

func (x *UpdateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_entpb_entpb_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return file_entpb_entpb_proto_rawDescGZIP(), []int{10}
}

func (x *UpdateUserRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type DeleteUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteUserRequest) Reset() {
	*x = DeleteUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entpb_entpb_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserRequest) ProtoMessage() {}

func (x *DeleteUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_entpb_entpb_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return file_entpb_entpb_proto_rawDescGZIP(), []int{11}
}

func (x *DeleteUserRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_entpb_entpb_proto protoreflect.FileDescriptor

var file_entpb_entpb_proto_rawDesc = []byte{
	0x0a, 0x11, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3d, 0x0a, 0x0a, 0x41, 0x74, 0x74, 0x61,
	0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x4c, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x31, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x26, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74, 0x61,
	0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4c, 0x0a,
	0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x61,
	0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x65,
	0x6e, 0x74, 0x70, 0x62, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x29, 0x0a, 0x17, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4e, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0xa9, 0x01, 0x0a, 0x04, 0x54, 0x6f, 0x64, 0x6f, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x61, 0x73, 0x6b, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x54, 0x6f, 0x64, 0x6f,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1f, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x22, 0x30, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45,
	0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x5f, 0x50, 0x52,
	0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x4f, 0x4e, 0x45,
	0x10, 0x02, 0x22, 0xbc, 0x03, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x6a, 0x6f, 0x69, 0x6e,
	0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x06, 0x6a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x03, 0x65, 0x78, 0x70, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x63, 0x72, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x63, 0x72, 0x6d, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x61,
	0x6e, 0x6e, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x62, 0x61, 0x6e, 0x6e,
	0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x70, 0x62, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x50, 0x62, 0x12,
	0x22, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x31, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e,
	0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x61,
	0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x39, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44,
	0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10,
	0x02, 0x22, 0x34, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x34, 0x0a, 0x11, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x65,
	0x6e, 0x74, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22,
	0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x32, 0x86, 0x02, 0x0a, 0x11, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x41, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x35, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1b,
	0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x65, 0x6e,
	0x74, 0x70, 0x62, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x3b,
	0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62,
	0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x40, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xd6, 0x01,
	0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0b, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x29,
	0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x15, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x65,
	0x6e, 0x74, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x06, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e,
	0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x3a, 0x0a, 0x06, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x39, 0x5a, 0x37, 0x65, 0x6e, 0x74, 0x67, 0x6f, 0x2e,
	0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x2f, 0x65, 0x6e, 0x74, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x6f, 0x64,
	0x6f, 0x2f, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6e, 0x74, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_entpb_entpb_proto_rawDescOnce sync.Once
	file_entpb_entpb_proto_rawDescData = file_entpb_entpb_proto_rawDesc
)

func file_entpb_entpb_proto_rawDescGZIP() []byte {
	file_entpb_entpb_proto_rawDescOnce.Do(func() {
		file_entpb_entpb_proto_rawDescData = protoimpl.X.CompressGZIP(file_entpb_entpb_proto_rawDescData)
	})
	return file_entpb_entpb_proto_rawDescData
}

var file_entpb_entpb_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_entpb_entpb_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_entpb_entpb_proto_goTypes = []interface{}{
	(Todo_Status)(0),                // 0: entpb.Todo.Status
	(User_Status)(0),                // 1: entpb.User.Status
	(*Attachment)(nil),              // 2: entpb.Attachment
	(*CreateAttachmentRequest)(nil), // 3: entpb.CreateAttachmentRequest
	(*GetAttachmentRequest)(nil),    // 4: entpb.GetAttachmentRequest
	(*UpdateAttachmentRequest)(nil), // 5: entpb.UpdateAttachmentRequest
	(*DeleteAttachmentRequest)(nil), // 6: entpb.DeleteAttachmentRequest
	(*Group)(nil),                   // 7: entpb.Group
	(*Todo)(nil),                    // 8: entpb.Todo
	(*User)(nil),                    // 9: entpb.User
	(*CreateUserRequest)(nil),       // 10: entpb.CreateUserRequest
	(*GetUserRequest)(nil),          // 11: entpb.GetUserRequest
	(*UpdateUserRequest)(nil),       // 12: entpb.UpdateUserRequest
	(*DeleteUserRequest)(nil),       // 13: entpb.DeleteUserRequest
	(*timestamppb.Timestamp)(nil),   // 14: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),           // 15: google.protobuf.Empty
}
var file_entpb_entpb_proto_depIdxs = []int32{
	9,  // 0: entpb.Attachment.user:type_name -> entpb.User
	2,  // 1: entpb.CreateAttachmentRequest.attachment:type_name -> entpb.Attachment
	2,  // 2: entpb.UpdateAttachmentRequest.attachment:type_name -> entpb.Attachment
	9,  // 3: entpb.Group.users:type_name -> entpb.User
	0,  // 4: entpb.Todo.status:type_name -> entpb.Todo.Status
	9,  // 5: entpb.Todo.user:type_name -> entpb.User
	14, // 6: entpb.User.joined:type_name -> google.protobuf.Timestamp
	1,  // 7: entpb.User.status:type_name -> entpb.User.Status
	7,  // 8: entpb.User.group:type_name -> entpb.Group
	2,  // 9: entpb.User.attachment:type_name -> entpb.Attachment
	9,  // 10: entpb.CreateUserRequest.user:type_name -> entpb.User
	9,  // 11: entpb.UpdateUserRequest.user:type_name -> entpb.User
	3,  // 12: entpb.AttachmentService.Create:input_type -> entpb.CreateAttachmentRequest
	4,  // 13: entpb.AttachmentService.Get:input_type -> entpb.GetAttachmentRequest
	5,  // 14: entpb.AttachmentService.Update:input_type -> entpb.UpdateAttachmentRequest
	6,  // 15: entpb.AttachmentService.Delete:input_type -> entpb.DeleteAttachmentRequest
	10, // 16: entpb.UserService.Create:input_type -> entpb.CreateUserRequest
	11, // 17: entpb.UserService.Get:input_type -> entpb.GetUserRequest
	12, // 18: entpb.UserService.Update:input_type -> entpb.UpdateUserRequest
	13, // 19: entpb.UserService.Delete:input_type -> entpb.DeleteUserRequest
	2,  // 20: entpb.AttachmentService.Create:output_type -> entpb.Attachment
	2,  // 21: entpb.AttachmentService.Get:output_type -> entpb.Attachment
	2,  // 22: entpb.AttachmentService.Update:output_type -> entpb.Attachment
	15, // 23: entpb.AttachmentService.Delete:output_type -> google.protobuf.Empty
	9,  // 24: entpb.UserService.Create:output_type -> entpb.User
	9,  // 25: entpb.UserService.Get:output_type -> entpb.User
	9,  // 26: entpb.UserService.Update:output_type -> entpb.User
	15, // 27: entpb.UserService.Delete:output_type -> google.protobuf.Empty
	20, // [20:28] is the sub-list for method output_type
	12, // [12:20] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_entpb_entpb_proto_init() }
func file_entpb_entpb_proto_init() {
	if File_entpb_entpb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_entpb_entpb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attachment); i {
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
		file_entpb_entpb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAttachmentRequest); i {
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
		file_entpb_entpb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAttachmentRequest); i {
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
		file_entpb_entpb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAttachmentRequest); i {
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
		file_entpb_entpb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAttachmentRequest); i {
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
		file_entpb_entpb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Group); i {
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
		file_entpb_entpb_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Todo); i {
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
		file_entpb_entpb_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_entpb_entpb_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
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
		file_entpb_entpb_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRequest); i {
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
		file_entpb_entpb_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserRequest); i {
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
		file_entpb_entpb_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserRequest); i {
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
			RawDescriptor: file_entpb_entpb_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_entpb_entpb_proto_goTypes,
		DependencyIndexes: file_entpb_entpb_proto_depIdxs,
		EnumInfos:         file_entpb_entpb_proto_enumTypes,
		MessageInfos:      file_entpb_entpb_proto_msgTypes,
	}.Build()
	File_entpb_entpb_proto = out.File
	file_entpb_entpb_proto_rawDesc = nil
	file_entpb_entpb_proto_goTypes = nil
	file_entpb_entpb_proto_depIdxs = nil
}
