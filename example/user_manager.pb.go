// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: user_manager.proto

package compete

import (
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateUserInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *CreateUserInput) Reset() {
	*x = CreateUserInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserInput) ProtoMessage() {}

func (x *CreateUserInput) ProtoReflect() protoreflect.Message {
	mi := &file_user_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserInput.ProtoReflect.Descriptor instead.
func (*CreateUserInput) Descriptor() ([]byte, []int) {
	return file_user_manager_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserInput) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateUserInput) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type CreateUserOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateUserOutput) Reset() {
	*x = CreateUserOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserOutput) ProtoMessage() {}

func (x *CreateUserOutput) ProtoReflect() protoreflect.Message {
	mi := &file_user_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserOutput.ProtoReflect.Descriptor instead.
func (*CreateUserOutput) Descriptor() ([]byte, []int) {
	return file_user_manager_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserOutput) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteUserInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *DeleteUserInput) Reset() {
	*x = DeleteUserInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_manager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserInput) ProtoMessage() {}

func (x *DeleteUserInput) ProtoReflect() protoreflect.Message {
	mi := &file_user_manager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserInput.ProtoReflect.Descriptor instead.
func (*DeleteUserInput) Descriptor() ([]byte, []int) {
	return file_user_manager_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteUserInput) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type DeleteUserOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteUserOutput) Reset() {
	*x = DeleteUserOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_manager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserOutput) ProtoMessage() {}

func (x *DeleteUserOutput) ProtoReflect() protoreflect.Message {
	mi := &file_user_manager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserOutput.ProtoReflect.Descriptor instead.
func (*DeleteUserOutput) Descriptor() ([]byte, []int) {
	return file_user_manager_proto_rawDescGZIP(), []int{3}
}

type GetCommentsInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Offset int32  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Size   int32  `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Sort   string `protobuf:"bytes,4,opt,name=sort,proto3" json:"sort,omitempty"`
	Order  string `protobuf:"bytes,5,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *GetCommentsInput) Reset() {
	*x = GetCommentsInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_manager_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommentsInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentsInput) ProtoMessage() {}

func (x *GetCommentsInput) ProtoReflect() protoreflect.Message {
	mi := &file_user_manager_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentsInput.ProtoReflect.Descriptor instead.
func (*GetCommentsInput) Descriptor() ([]byte, []int) {
	return file_user_manager_proto_rawDescGZIP(), []int{4}
}

func (x *GetCommentsInput) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetCommentsInput) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetCommentsInput) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *GetCommentsInput) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

func (x *GetCommentsInput) GetOrder() string {
	if x != nil {
		return x.Order
	}
	return ""
}

type GetCommentsOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comments []*Comment `protobuf:"bytes,1,rep,name=comments,proto3" json:"comments,omitempty"`
}

func (x *GetCommentsOutput) Reset() {
	*x = GetCommentsOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_manager_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommentsOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentsOutput) ProtoMessage() {}

func (x *GetCommentsOutput) ProtoReflect() protoreflect.Message {
	mi := &file_user_manager_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentsOutput.ProtoReflect.Descriptor instead.
func (*GetCommentsOutput) Descriptor() ([]byte, []int) {
	return file_user_manager_proto_rawDescGZIP(), []int{5}
}

func (x *GetCommentsOutput) GetComments() []*Comment {
	if x != nil {
		return x.Comments
	}
	return nil
}

type CreateCommentInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Comment *Comment `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *CreateCommentInput) Reset() {
	*x = CreateCommentInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_manager_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCommentInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCommentInput) ProtoMessage() {}

func (x *CreateCommentInput) ProtoReflect() protoreflect.Message {
	mi := &file_user_manager_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCommentInput.ProtoReflect.Descriptor instead.
func (*CreateCommentInput) Descriptor() ([]byte, []int) {
	return file_user_manager_proto_rawDescGZIP(), []int{6}
}

func (x *CreateCommentInput) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateCommentInput) GetComment() *Comment {
	if x != nil {
		return x.Comment
	}
	return nil
}

type CreateCommentOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommentId string `protobuf:"bytes,1,opt,name=comment_id,json=commentId,proto3" json:"comment_id,omitempty"`
}

func (x *CreateCommentOutput) Reset() {
	*x = CreateCommentOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_manager_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCommentOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCommentOutput) ProtoMessage() {}

func (x *CreateCommentOutput) ProtoReflect() protoreflect.Message {
	mi := &file_user_manager_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCommentOutput.ProtoReflect.Descriptor instead.
func (*CreateCommentOutput) Descriptor() ([]byte, []int) {
	return file_user_manager_proto_rawDescGZIP(), []int{7}
}

func (x *CreateCommentOutput) GetCommentId() string {
	if x != nil {
		return x.CommentId
	}
	return ""
}

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_manager_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_user_manager_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_user_manager_proto_rawDescGZIP(), []int{8}
}

func (x *Comment) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Comment) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_user_manager_proto protoreflect.FileDescriptor

var file_user_manager_proto_rawDesc = []byte{
	0x0a, 0x12, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x43, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x22, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2a, 0x0a, 0x0f, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x12, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x81, 0x01, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22,
	0x39, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x12, 0x24, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x51, 0x0a, 0x12, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x34, 0x0a,
	0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x22, 0x39, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x32, 0xe1,
	0x02, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x44,
	0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x11,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x22, 0x06, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x3a, 0x01, 0x2a, 0x12, 0x4b, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x10, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x1a, 0x11, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x2a,
	0x10, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x7d, 0x12, 0x57, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x1a, 0x12, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12,
	0x19, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x7d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x66, 0x0a, 0x0d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x13, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x1a, 0x14, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x22, 0x19,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x65, 0x3b, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_manager_proto_rawDescOnce sync.Once
	file_user_manager_proto_rawDescData = file_user_manager_proto_rawDesc
)

func file_user_manager_proto_rawDescGZIP() []byte {
	file_user_manager_proto_rawDescOnce.Do(func() {
		file_user_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_manager_proto_rawDescData)
	})
	return file_user_manager_proto_rawDescData
}

var file_user_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_user_manager_proto_goTypes = []interface{}{
	(*CreateUserInput)(nil),     // 0: CreateUserInput
	(*CreateUserOutput)(nil),    // 1: CreateUserOutput
	(*DeleteUserInput)(nil),     // 2: DeleteUserInput
	(*DeleteUserOutput)(nil),    // 3: DeleteUserOutput
	(*GetCommentsInput)(nil),    // 4: GetCommentsInput
	(*GetCommentsOutput)(nil),   // 5: GetCommentsOutput
	(*CreateCommentInput)(nil),  // 6: CreateCommentInput
	(*CreateCommentOutput)(nil), // 7: CreateCommentOutput
	(*Comment)(nil),             // 8: Comment
}
var file_user_manager_proto_depIdxs = []int32{
	8, // 0: GetCommentsOutput.comments:type_name -> Comment
	8, // 1: CreateCommentInput.comment:type_name -> Comment
	0, // 2: UserManager.CreateUser:input_type -> CreateUserInput
	2, // 3: UserManager.DeleteUser:input_type -> DeleteUserInput
	4, // 4: UserManager.GetComments:input_type -> GetCommentsInput
	6, // 5: UserManager.CreateComment:input_type -> CreateCommentInput
	1, // 6: UserManager.CreateUser:output_type -> CreateUserOutput
	3, // 7: UserManager.DeleteUser:output_type -> DeleteUserOutput
	5, // 8: UserManager.GetComments:output_type -> GetCommentsOutput
	7, // 9: UserManager.CreateComment:output_type -> CreateCommentOutput
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_user_manager_proto_init() }
func file_user_manager_proto_init() {
	if File_user_manager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserInput); i {
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
		file_user_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserOutput); i {
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
		file_user_manager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserInput); i {
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
		file_user_manager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserOutput); i {
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
		file_user_manager_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCommentsInput); i {
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
		file_user_manager_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCommentsOutput); i {
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
		file_user_manager_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCommentInput); i {
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
		file_user_manager_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCommentOutput); i {
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
		file_user_manager_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
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
			RawDescriptor: file_user_manager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_manager_proto_goTypes,
		DependencyIndexes: file_user_manager_proto_depIdxs,
		MessageInfos:      file_user_manager_proto_msgTypes,
	}.Build()
	File_user_manager_proto = out.File
	file_user_manager_proto_rawDesc = nil
	file_user_manager_proto_goTypes = nil
	file_user_manager_proto_depIdxs = nil
}
