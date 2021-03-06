// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.0
// source: privacy.proto

package proto

import (
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

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId        string `protobuf:"bytes,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	BlockedUserId string `protobuf:"bytes,2,opt,name=BlockedUserId,proto3" json:"BlockedUserId,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_privacy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_privacy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_privacy_proto_rawDescGZIP(), []int{0}
}

func (x *Block) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Block) GetBlockedUserId() string {
	if x != nil {
		return x.BlockedUserId
	}
	return ""
}

type PrivacyMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	IsProfilePublic bool   `protobuf:"varint,2,opt,name=isProfilePublic,proto3" json:"isProfilePublic,omitempty"`
	IsDmPublic      bool   `protobuf:"varint,3,opt,name=isDmPublic,proto3" json:"isDmPublic,omitempty"`
	IsTagEnabled    bool   `protobuf:"varint,4,opt,name=isTagEnabled,proto3" json:"isTagEnabled,omitempty"`
}

func (x *PrivacyMessage) Reset() {
	*x = PrivacyMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_privacy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivacyMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivacyMessage) ProtoMessage() {}

func (x *PrivacyMessage) ProtoReflect() protoreflect.Message {
	mi := &file_privacy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivacyMessage.ProtoReflect.Descriptor instead.
func (*PrivacyMessage) Descriptor() ([]byte, []int) {
	return file_privacy_proto_rawDescGZIP(), []int{1}
}

func (x *PrivacyMessage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PrivacyMessage) GetIsProfilePublic() bool {
	if x != nil {
		return x.IsProfilePublic
	}
	return false
}

func (x *PrivacyMessage) GetIsDmPublic() bool {
	if x != nil {
		return x.IsDmPublic
	}
	return false
}

func (x *PrivacyMessage) GetIsTagEnabled() bool {
	if x != nil {
		return x.IsTagEnabled
	}
	return false
}

type CreatePrivacyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Privacy *PrivacyMessage `protobuf:"bytes,1,opt,name=privacy,proto3" json:"privacy,omitempty"`
}

func (x *CreatePrivacyRequest) Reset() {
	*x = CreatePrivacyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_privacy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePrivacyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePrivacyRequest) ProtoMessage() {}

func (x *CreatePrivacyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_privacy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePrivacyRequest.ProtoReflect.Descriptor instead.
func (*CreatePrivacyRequest) Descriptor() ([]byte, []int) {
	return file_privacy_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePrivacyRequest) GetPrivacy() *PrivacyMessage {
	if x != nil {
		return x.Privacy
	}
	return nil
}

type CreateBlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Block *Block `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
}

func (x *CreateBlockRequest) Reset() {
	*x = CreateBlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_privacy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBlockRequest) ProtoMessage() {}

func (x *CreateBlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_privacy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBlockRequest.ProtoReflect.Descriptor instead.
func (*CreateBlockRequest) Descriptor() ([]byte, []int) {
	return file_privacy_proto_rawDescGZIP(), []int{3}
}

func (x *CreateBlockRequest) GetBlock() *Block {
	if x != nil {
		return x.Block
	}
	return nil
}

type BooleanResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response bool `protobuf:"varint,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *BooleanResponse) Reset() {
	*x = BooleanResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_privacy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BooleanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BooleanResponse) ProtoMessage() {}

func (x *BooleanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_privacy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BooleanResponse.ProtoReflect.Descriptor instead.
func (*BooleanResponse) Descriptor() ([]byte, []int) {
	return file_privacy_proto_rawDescGZIP(), []int{4}
}

func (x *BooleanResponse) GetResponse() bool {
	if x != nil {
		return x.Response
	}
	return false
}

type PrivacyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *PrivacyRequest) Reset() {
	*x = PrivacyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_privacy_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivacyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivacyRequest) ProtoMessage() {}

func (x *PrivacyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_privacy_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivacyRequest.ProtoReflect.Descriptor instead.
func (*PrivacyRequest) Descriptor() ([]byte, []int) {
	return file_privacy_proto_rawDescGZIP(), []int{5}
}

func (x *PrivacyRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type StringArray struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *StringArray) Reset() {
	*x = StringArray{}
	if protoimpl.UnsafeEnabled {
		mi := &file_privacy_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringArray) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringArray) ProtoMessage() {}

func (x *StringArray) ProtoReflect() protoreflect.Message {
	mi := &file_privacy_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringArray.ProtoReflect.Descriptor instead.
func (*StringArray) Descriptor() ([]byte, []int) {
	return file_privacy_proto_rawDescGZIP(), []int{6}
}

func (x *StringArray) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type RequestIdPrivacy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RequestIdPrivacy) Reset() {
	*x = RequestIdPrivacy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_privacy_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestIdPrivacy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestIdPrivacy) ProtoMessage() {}

func (x *RequestIdPrivacy) ProtoReflect() protoreflect.Message {
	mi := &file_privacy_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestIdPrivacy.ProtoReflect.Descriptor instead.
func (*RequestIdPrivacy) Descriptor() ([]byte, []int) {
	return file_privacy_proto_rawDescGZIP(), []int{7}
}

func (x *RequestIdPrivacy) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type EmptyResponsePrivacy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyResponsePrivacy) Reset() {
	*x = EmptyResponsePrivacy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_privacy_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResponsePrivacy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResponsePrivacy) ProtoMessage() {}

func (x *EmptyResponsePrivacy) ProtoReflect() protoreflect.Message {
	mi := &file_privacy_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResponsePrivacy.ProtoReflect.Descriptor instead.
func (*EmptyResponsePrivacy) Descriptor() ([]byte, []int) {
	return file_privacy_proto_rawDescGZIP(), []int{8}
}

type EmptyRequestPrivacy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyRequestPrivacy) Reset() {
	*x = EmptyRequestPrivacy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_privacy_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyRequestPrivacy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRequestPrivacy) ProtoMessage() {}

func (x *EmptyRequestPrivacy) ProtoReflect() protoreflect.Message {
	mi := &file_privacy_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRequestPrivacy.ProtoReflect.Descriptor instead.
func (*EmptyRequestPrivacy) Descriptor() ([]byte, []int) {
	return file_privacy_proto_rawDescGZIP(), []int{9}
}

var File_privacy_proto protoreflect.FileDescriptor

var file_privacy_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x16, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x8e, 0x01, 0x0a, 0x0e,
	0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x28,
	0x0a, 0x0f, 0x69, 0x73, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69, 0x73, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x73, 0x44, 0x6d,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73,
	0x44, 0x6d, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x73, 0x54, 0x61,
	0x67, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c,
	0x69, 0x73, 0x54, 0x61, 0x67, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0x47, 0x0a, 0x14,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72,
	0x69, 0x76, 0x61, 0x63, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x63, 0x79, 0x22, 0x38, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x05, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x22,
	0x2d, 0x0a, 0x0f, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28,
	0x0a, 0x0e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x1f, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x22, 0x0a, 0x10, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x16, 0x0a,
	0x14, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72,
	0x69, 0x76, 0x61, 0x63, 0x79, 0x22, 0x15, 0x0a, 0x13, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x32, 0xdf, 0x06, 0x0a,
	0x07, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x12, 0x62, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72, 0x69, 0x76,
	0x61, 0x63, 0x79, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x0c, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x3a, 0x01, 0x2a, 0x12, 0x69, 0x0a, 0x0d,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x12, 0x1b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x76,
	0x61, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22,
	0x13, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x6b, 0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x22, 0x26, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x20, 0x22, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x63, 0x79, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x3a, 0x05, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x6f, 0x0a, 0x0b, 0x55, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x22, 0x28, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x22, 0x22, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63,
	0x79, 0x2f, 0x75, 0x6e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x3a, 0x05,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x68, 0x0a, 0x0e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x66,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x65,
	0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1d, 0x22, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79,
	0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x3a, 0x01, 0x2a, 0x12,
	0x70, 0x0a, 0x16, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21,
	0x22, 0x1c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x69,
	0x73, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x3a, 0x01,
	0x2a, 0x12, 0x62, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x1a,
	0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x41, 0x72,
	0x72, 0x61, 0x79, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x67, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79,
	0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x12,
	0x1d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x67, 0x65,
	0x74, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x2b,
	0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x77, 0x73,
	0x32, 0x30, 0x32, 0x31, 0x2d, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x67, 0x72, 0x61, 0x6d, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_privacy_proto_rawDescOnce sync.Once
	file_privacy_proto_rawDescData = file_privacy_proto_rawDesc
)

func file_privacy_proto_rawDescGZIP() []byte {
	file_privacy_proto_rawDescOnce.Do(func() {
		file_privacy_proto_rawDescData = protoimpl.X.CompressGZIP(file_privacy_proto_rawDescData)
	})
	return file_privacy_proto_rawDescData
}

var file_privacy_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_privacy_proto_goTypes = []interface{}{
	(*Block)(nil),                // 0: proto.Block
	(*PrivacyMessage)(nil),       // 1: proto.PrivacyMessage
	(*CreatePrivacyRequest)(nil), // 2: proto.CreatePrivacyRequest
	(*CreateBlockRequest)(nil),   // 3: proto.CreateBlockRequest
	(*BooleanResponse)(nil),      // 4: proto.BooleanResponse
	(*PrivacyRequest)(nil),       // 5: proto.PrivacyRequest
	(*StringArray)(nil),          // 6: proto.StringArray
	(*RequestIdPrivacy)(nil),     // 7: proto.RequestIdPrivacy
	(*EmptyResponsePrivacy)(nil), // 8: proto.EmptyResponsePrivacy
	(*EmptyRequestPrivacy)(nil),  // 9: proto.EmptyRequestPrivacy
}
var file_privacy_proto_depIdxs = []int32{
	1,  // 0: proto.CreatePrivacyRequest.privacy:type_name -> proto.PrivacyMessage
	0,  // 1: proto.CreateBlockRequest.block:type_name -> proto.Block
	2,  // 2: proto.Privacy.CreatePrivacy:input_type -> proto.CreatePrivacyRequest
	2,  // 3: proto.Privacy.UpdatePrivacy:input_type -> proto.CreatePrivacyRequest
	3,  // 4: proto.Privacy.BlockUser:input_type -> proto.CreateBlockRequest
	3,  // 5: proto.Privacy.UnBlockUser:input_type -> proto.CreateBlockRequest
	3,  // 6: proto.Privacy.CheckIfBlocked:input_type -> proto.CreateBlockRequest
	5,  // 7: proto.Privacy.CheckUserProfilePublic:input_type -> proto.PrivacyRequest
	7,  // 8: proto.Privacy.GetAllPublicUsers:input_type -> proto.RequestIdPrivacy
	7,  // 9: proto.Privacy.GetUserPrivacy:input_type -> proto.RequestIdPrivacy
	8,  // 10: proto.Privacy.CreatePrivacy:output_type -> proto.EmptyResponsePrivacy
	8,  // 11: proto.Privacy.UpdatePrivacy:output_type -> proto.EmptyResponsePrivacy
	8,  // 12: proto.Privacy.BlockUser:output_type -> proto.EmptyResponsePrivacy
	8,  // 13: proto.Privacy.UnBlockUser:output_type -> proto.EmptyResponsePrivacy
	4,  // 14: proto.Privacy.CheckIfBlocked:output_type -> proto.BooleanResponse
	4,  // 15: proto.Privacy.CheckUserProfilePublic:output_type -> proto.BooleanResponse
	6,  // 16: proto.Privacy.GetAllPublicUsers:output_type -> proto.StringArray
	1,  // 17: proto.Privacy.GetUserPrivacy:output_type -> proto.PrivacyMessage
	10, // [10:18] is the sub-list for method output_type
	2,  // [2:10] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_privacy_proto_init() }
func file_privacy_proto_init() {
	if File_privacy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_privacy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
		file_privacy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrivacyMessage); i {
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
		file_privacy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePrivacyRequest); i {
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
		file_privacy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBlockRequest); i {
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
		file_privacy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BooleanResponse); i {
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
		file_privacy_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrivacyRequest); i {
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
		file_privacy_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringArray); i {
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
		file_privacy_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestIdPrivacy); i {
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
		file_privacy_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyResponsePrivacy); i {
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
		file_privacy_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyRequestPrivacy); i {
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
			RawDescriptor: file_privacy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_privacy_proto_goTypes,
		DependencyIndexes: file_privacy_proto_depIdxs,
		MessageInfos:      file_privacy_proto_msgTypes,
	}.Build()
	File_privacy_proto = out.File
	file_privacy_proto_rawDesc = nil
	file_privacy_proto_goTypes = nil
	file_privacy_proto_depIdxs = nil
}
