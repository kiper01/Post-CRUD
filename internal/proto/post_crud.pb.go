// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: proto/post_crud.proto

package Post_CRUD

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

type PostValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Code  uint32 `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Name  string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	River string `protobuf:"bytes,4,opt,name=river,proto3" json:"river,omitempty"`
}

func (x *PostValue) Reset() {
	*x = PostValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_post_crud_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostValue) ProtoMessage() {}

func (x *PostValue) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_crud_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostValue.ProtoReflect.Descriptor instead.
func (*PostValue) Descriptor() ([]byte, []int) {
	return file_proto_post_crud_proto_rawDescGZIP(), []int{0}
}

func (x *PostValue) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PostValue) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *PostValue) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PostValue) GetRiver() string {
	if x != nil {
		return x.River
	}
	return ""
}

type AddPostValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	River string `protobuf:"bytes,3,opt,name=river,proto3" json:"river,omitempty"`
}

func (x *AddPostValueRequest) Reset() {
	*x = AddPostValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_post_crud_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPostValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPostValueRequest) ProtoMessage() {}

func (x *AddPostValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_crud_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPostValueRequest.ProtoReflect.Descriptor instead.
func (*AddPostValueRequest) Descriptor() ([]byte, []int) {
	return file_proto_post_crud_proto_rawDescGZIP(), []int{1}
}

func (x *AddPostValueRequest) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *AddPostValueRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddPostValueRequest) GetRiver() string {
	if x != nil {
		return x.River
	}
	return ""
}

type AddPostValueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostValue *PostValue `protobuf:"bytes,1,opt,name=post_value,json=postValue,proto3" json:"post_value,omitempty"`
}

func (x *AddPostValueResponse) Reset() {
	*x = AddPostValueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_post_crud_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPostValueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPostValueResponse) ProtoMessage() {}

func (x *AddPostValueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_crud_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPostValueResponse.ProtoReflect.Descriptor instead.
func (*AddPostValueResponse) Descriptor() ([]byte, []int) {
	return file_proto_post_crud_proto_rawDescGZIP(), []int{2}
}

func (x *AddPostValueResponse) GetPostValue() *PostValue {
	if x != nil {
		return x.PostValue
	}
	return nil
}

type RemovePostValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RemovePostValueRequest) Reset() {
	*x = RemovePostValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_post_crud_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemovePostValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemovePostValueRequest) ProtoMessage() {}

func (x *RemovePostValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_crud_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemovePostValueRequest.ProtoReflect.Descriptor instead.
func (*RemovePostValueRequest) Descriptor() ([]byte, []int) {
	return file_proto_post_crud_proto_rawDescGZIP(), []int{3}
}

func (x *RemovePostValueRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RemovePostValueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *RemovePostValueResponse) Reset() {
	*x = RemovePostValueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_post_crud_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemovePostValueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemovePostValueResponse) ProtoMessage() {}

func (x *RemovePostValueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_crud_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemovePostValueResponse.ProtoReflect.Descriptor instead.
func (*RemovePostValueResponse) Descriptor() ([]byte, []int) {
	return file_proto_post_crud_proto_rawDescGZIP(), []int{4}
}

func (x *RemovePostValueResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type UpdatePostValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostValue []*PostValue `protobuf:"bytes,1,rep,name=post_value,json=postValue,proto3" json:"post_value,omitempty"`
}

func (x *UpdatePostValueRequest) Reset() {
	*x = UpdatePostValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_post_crud_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePostValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePostValueRequest) ProtoMessage() {}

func (x *UpdatePostValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_crud_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePostValueRequest.ProtoReflect.Descriptor instead.
func (*UpdatePostValueRequest) Descriptor() ([]byte, []int) {
	return file_proto_post_crud_proto_rawDescGZIP(), []int{5}
}

func (x *UpdatePostValueRequest) GetPostValue() []*PostValue {
	if x != nil {
		return x.PostValue
	}
	return nil
}

type UpdatePostValueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostValue []*PostValue `protobuf:"bytes,1,rep,name=post_value,json=postValue,proto3" json:"post_value,omitempty"`
}

func (x *UpdatePostValueResponse) Reset() {
	*x = UpdatePostValueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_post_crud_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePostValueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePostValueResponse) ProtoMessage() {}

func (x *UpdatePostValueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_crud_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePostValueResponse.ProtoReflect.Descriptor instead.
func (*UpdatePostValueResponse) Descriptor() ([]byte, []int) {
	return file_proto_post_crud_proto_rawDescGZIP(), []int{6}
}

func (x *UpdatePostValueResponse) GetPostValue() []*PostValue {
	if x != nil {
		return x.PostValue
	}
	return nil
}

type GetPostValuesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Page uint32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *GetPostValuesRequest) Reset() {
	*x = GetPostValuesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_post_crud_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostValuesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostValuesRequest) ProtoMessage() {}

func (x *GetPostValuesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_crud_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostValuesRequest.ProtoReflect.Descriptor instead.
func (*GetPostValuesRequest) Descriptor() ([]byte, []int) {
	return file_proto_post_crud_proto_rawDescGZIP(), []int{7}
}

func (x *GetPostValuesRequest) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetPostValuesRequest) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

type GetPostValuesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page      uint32       `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	MaxPage   uint32       `protobuf:"varint,2,opt,name=max_page,json=maxPage,proto3" json:"max_page,omitempty"`
	PostValue []*PostValue `protobuf:"bytes,3,rep,name=post_value,json=postValue,proto3" json:"post_value,omitempty"`
}

func (x *GetPostValuesResponse) Reset() {
	*x = GetPostValuesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_post_crud_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostValuesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostValuesResponse) ProtoMessage() {}

func (x *GetPostValuesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_crud_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostValuesResponse.ProtoReflect.Descriptor instead.
func (*GetPostValuesResponse) Descriptor() ([]byte, []int) {
	return file_proto_post_crud_proto_rawDescGZIP(), []int{8}
}

func (x *GetPostValuesResponse) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetPostValuesResponse) GetMaxPage() uint32 {
	if x != nil {
		return x.MaxPage
	}
	return 0
}

func (x *GetPostValuesResponse) GetPostValue() []*PostValue {
	if x != nil {
		return x.PostValue
	}
	return nil
}

type GetPostValuesByCodeOrIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Code uint32 `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *GetPostValuesByCodeOrIdRequest) Reset() {
	*x = GetPostValuesByCodeOrIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_post_crud_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostValuesByCodeOrIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostValuesByCodeOrIdRequest) ProtoMessage() {}

func (x *GetPostValuesByCodeOrIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_crud_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostValuesByCodeOrIdRequest.ProtoReflect.Descriptor instead.
func (*GetPostValuesByCodeOrIdRequest) Descriptor() ([]byte, []int) {
	return file_proto_post_crud_proto_rawDescGZIP(), []int{9}
}

func (x *GetPostValuesByCodeOrIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetPostValuesByCodeOrIdRequest) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type GetPostValuesByCodeOrIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostValue []*PostValue `protobuf:"bytes,1,rep,name=post_value,json=postValue,proto3" json:"post_value,omitempty"`
}

func (x *GetPostValuesByCodeOrIdResponse) Reset() {
	*x = GetPostValuesByCodeOrIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_post_crud_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostValuesByCodeOrIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostValuesByCodeOrIdResponse) ProtoMessage() {}

func (x *GetPostValuesByCodeOrIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_crud_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostValuesByCodeOrIdResponse.ProtoReflect.Descriptor instead.
func (*GetPostValuesByCodeOrIdResponse) Descriptor() ([]byte, []int) {
	return file_proto_post_crud_proto_rawDescGZIP(), []int{10}
}

func (x *GetPostValuesByCodeOrIdResponse) GetPostValue() []*PostValue {
	if x != nil {
		return x.PostValue
	}
	return nil
}

var File_proto_post_crud_proto protoreflect.FileDescriptor

var file_proto_post_crud_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x63, 0x72, 0x75,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x63, 0x72, 0x75,
	0x64, 0x22, 0x59, 0x0a, 0x09, 0x50, 0x6f, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x69, 0x76, 0x65, 0x72, 0x22, 0x53, 0x0a, 0x13,
	0x41, 0x64, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x22, 0x4a, 0x0a, 0x14, 0x41, 0x64, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x0a, 0x70, 0x6f, 0x73,
	0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x70, 0x6f, 0x73, 0x74, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x28, 0x0a,
	0x16, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x33, 0x0a, 0x17, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x50, 0x6f, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x4c, 0x0a, 0x16,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x6f, 0x73,
	0x74, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x09, 0x70, 0x6f, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x4d, 0x0a, 0x17, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x6f, 0x73, 0x74,
	0x63, 0x72, 0x75, 0x64, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09,
	0x70, 0x6f, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x3e, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x50, 0x6f, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x7a, 0x0a, 0x15, 0x47, 0x65, 0x74,
	0x50, 0x6f, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x61, 0x78, 0x5f, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x50, 0x61, 0x67,
	0x65, 0x12, 0x32, 0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x63, 0x72, 0x75, 0x64,
	0x2e, 0x50, 0x6f, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x74,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x44, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x42, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x4f, 0x72, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x55, 0x0a, 0x1f, 0x47,
	0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x42, 0x79, 0x43, 0x6f,
	0x64, 0x65, 0x4f, 0x72, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32,
	0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x50, 0x6f,
	0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x74, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x32, 0xd2, 0x03, 0x0a, 0x0f, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x50, 0x6f, 0x73,
	0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x63, 0x72, 0x75,
	0x64, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x63, 0x72, 0x75, 0x64,
	0x2e, 0x41, 0x64, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x0f, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50,
	0x6f, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x20, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x63,
	0x72, 0x75, 0x64, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x6f, 0x73,
	0x74, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x6f, 0x73, 0x74,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a,
	0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x20, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x1e, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x63, 0x72, 0x75,
	0x64, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x63, 0x72, 0x75,
	0x64, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6e, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x50, 0x6f,
	0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x42, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x4f, 0x72,
	0x49, 0x64, 0x12, 0x28, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x6f, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x42, 0x79, 0x43, 0x6f, 0x64,
	0x65, 0x4f, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x70,
	0x6f, 0x73, 0x74, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x42, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x4f, 0x72, 0x49, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1e, 0x5a, 0x1c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x69, 0x70, 0x65, 0x72, 0x30, 0x31, 0x2f, 0x50, 0x6f,
	0x73, 0x74, 0x2d, 0x43, 0x52, 0x55, 0x44, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_post_crud_proto_rawDescOnce sync.Once
	file_proto_post_crud_proto_rawDescData = file_proto_post_crud_proto_rawDesc
)

func file_proto_post_crud_proto_rawDescGZIP() []byte {
	file_proto_post_crud_proto_rawDescOnce.Do(func() {
		file_proto_post_crud_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_post_crud_proto_rawDescData)
	})
	return file_proto_post_crud_proto_rawDescData
}

var file_proto_post_crud_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_post_crud_proto_goTypes = []interface{}{
	(*PostValue)(nil),                       // 0: postcrud.PostValue
	(*AddPostValueRequest)(nil),             // 1: postcrud.AddPostValueRequest
	(*AddPostValueResponse)(nil),            // 2: postcrud.AddPostValueResponse
	(*RemovePostValueRequest)(nil),          // 3: postcrud.RemovePostValueRequest
	(*RemovePostValueResponse)(nil),         // 4: postcrud.RemovePostValueResponse
	(*UpdatePostValueRequest)(nil),          // 5: postcrud.UpdatePostValueRequest
	(*UpdatePostValueResponse)(nil),         // 6: postcrud.UpdatePostValueResponse
	(*GetPostValuesRequest)(nil),            // 7: postcrud.GetPostValuesRequest
	(*GetPostValuesResponse)(nil),           // 8: postcrud.GetPostValuesResponse
	(*GetPostValuesByCodeOrIdRequest)(nil),  // 9: postcrud.GetPostValuesByCodeOrIdRequest
	(*GetPostValuesByCodeOrIdResponse)(nil), // 10: postcrud.GetPostValuesByCodeOrIdResponse
}
var file_proto_post_crud_proto_depIdxs = []int32{
	0,  // 0: postcrud.AddPostValueResponse.post_value:type_name -> postcrud.PostValue
	0,  // 1: postcrud.UpdatePostValueRequest.post_value:type_name -> postcrud.PostValue
	0,  // 2: postcrud.UpdatePostValueResponse.post_value:type_name -> postcrud.PostValue
	0,  // 3: postcrud.GetPostValuesResponse.post_value:type_name -> postcrud.PostValue
	0,  // 4: postcrud.GetPostValuesByCodeOrIdResponse.post_value:type_name -> postcrud.PostValue
	1,  // 5: postcrud.PostInfoService.AddPostValue:input_type -> postcrud.AddPostValueRequest
	3,  // 6: postcrud.PostInfoService.RemovePostValue:input_type -> postcrud.RemovePostValueRequest
	5,  // 7: postcrud.PostInfoService.UpdatePostValue:input_type -> postcrud.UpdatePostValueRequest
	7,  // 8: postcrud.PostInfoService.GetPostValues:input_type -> postcrud.GetPostValuesRequest
	9,  // 9: postcrud.PostInfoService.GetPostValuesByCodeOrId:input_type -> postcrud.GetPostValuesByCodeOrIdRequest
	2,  // 10: postcrud.PostInfoService.AddPostValue:output_type -> postcrud.AddPostValueResponse
	4,  // 11: postcrud.PostInfoService.RemovePostValue:output_type -> postcrud.RemovePostValueResponse
	6,  // 12: postcrud.PostInfoService.UpdatePostValue:output_type -> postcrud.UpdatePostValueResponse
	8,  // 13: postcrud.PostInfoService.GetPostValues:output_type -> postcrud.GetPostValuesResponse
	10, // 14: postcrud.PostInfoService.GetPostValuesByCodeOrId:output_type -> postcrud.GetPostValuesByCodeOrIdResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_proto_post_crud_proto_init() }
func file_proto_post_crud_proto_init() {
	if File_proto_post_crud_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_post_crud_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostValue); i {
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
		file_proto_post_crud_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPostValueRequest); i {
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
		file_proto_post_crud_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPostValueResponse); i {
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
		file_proto_post_crud_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemovePostValueRequest); i {
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
		file_proto_post_crud_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemovePostValueResponse); i {
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
		file_proto_post_crud_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePostValueRequest); i {
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
		file_proto_post_crud_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePostValueResponse); i {
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
		file_proto_post_crud_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostValuesRequest); i {
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
		file_proto_post_crud_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostValuesResponse); i {
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
		file_proto_post_crud_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostValuesByCodeOrIdRequest); i {
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
		file_proto_post_crud_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostValuesByCodeOrIdResponse); i {
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
			RawDescriptor: file_proto_post_crud_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_post_crud_proto_goTypes,
		DependencyIndexes: file_proto_post_crud_proto_depIdxs,
		MessageInfos:      file_proto_post_crud_proto_msgTypes,
	}.Build()
	File_proto_post_crud_proto = out.File
	file_proto_post_crud_proto_rawDesc = nil
	file_proto_post_crud_proto_goTypes = nil
	file_proto_post_crud_proto_depIdxs = nil
}
