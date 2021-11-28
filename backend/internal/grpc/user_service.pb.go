// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: user_service.proto

package grpc

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

type AuthToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token *string `protobuf:"bytes,1,req,name=token" json:"token,omitempty"`
}

func (x *AuthToken) Reset() {
	*x = AuthToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthToken) ProtoMessage() {}

func (x *AuthToken) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthToken.ProtoReflect.Descriptor instead.
func (*AuthToken) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{0}
}

func (x *AuthToken) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *string `protobuf:"bytes,1,req,name=result" json:"result,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{1}
}

func (x *Result) GetResult() string {
	if x != nil && x.Result != nil {
		return *x.Result
	}
	return ""
}

type ServiceName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Url  *string `protobuf:"bytes,2,req,name=url" json:"url,omitempty"`
}

func (x *ServiceName) Reset() {
	*x = ServiceName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceName) ProtoMessage() {}

func (x *ServiceName) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceName.ProtoReflect.Descriptor instead.
func (*ServiceName) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{2}
}

func (x *ServiceName) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *ServiceName) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

type AuthorizeParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token      *string `protobuf:"bytes,1,req,name=token" json:"token,omitempty"`
	UserUUID   *string `protobuf:"bytes,2,req,name=userUUID" json:"userUUID,omitempty"`
	Permission *string `protobuf:"bytes,3,req,name=permission" json:"permission,omitempty"`
}

func (x *AuthorizeParams) Reset() {
	*x = AuthorizeParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizeParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizeParams) ProtoMessage() {}

func (x *AuthorizeParams) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizeParams.ProtoReflect.Descriptor instead.
func (*AuthorizeParams) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{3}
}

func (x *AuthorizeParams) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

func (x *AuthorizeParams) GetUserUUID() string {
	if x != nil && x.UserUUID != nil {
		return *x.UserUUID
	}
	return ""
}

func (x *AuthorizeParams) GetPermission() string {
	if x != nil && x.Permission != nil {
		return *x.Permission
	}
	return ""
}

type SignupParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username *string `protobuf:"bytes,1,req,name=username" json:"username,omitempty"`
	Password *string `protobuf:"bytes,2,req,name=password" json:"password,omitempty"`
}

func (x *SignupParams) Reset() {
	*x = SignupParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignupParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignupParams) ProtoMessage() {}

func (x *SignupParams) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignupParams.ProtoReflect.Descriptor instead.
func (*SignupParams) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{4}
}

func (x *SignupParams) GetUsername() string {
	if x != nil && x.Username != nil {
		return *x.Username
	}
	return ""
}

func (x *SignupParams) GetPassword() string {
	if x != nil && x.Password != nil {
		return *x.Password
	}
	return ""
}

type LoginParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username *string `protobuf:"bytes,1,req,name=username" json:"username,omitempty"`
	Password *string `protobuf:"bytes,2,req,name=password" json:"password,omitempty"`
	Token    *string `protobuf:"bytes,3,req,name=token" json:"token,omitempty"`
}

func (x *LoginParams) Reset() {
	*x = LoginParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginParams) ProtoMessage() {}

func (x *LoginParams) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginParams.ProtoReflect.Descriptor instead.
func (*LoginParams) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{5}
}

func (x *LoginParams) GetUsername() string {
	if x != nil && x.Username != nil {
		return *x.Username
	}
	return ""
}

func (x *LoginParams) GetPassword() string {
	if x != nil && x.Password != nil {
		return *x.Password
	}
	return ""
}

func (x *LoginParams) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

type TokenWithPermission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token      *string `protobuf:"bytes,1,req,name=token" json:"token,omitempty"`
	Permission *string `protobuf:"bytes,2,req,name=permission" json:"permission,omitempty"`
}

func (x *TokenWithPermission) Reset() {
	*x = TokenWithPermission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenWithPermission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenWithPermission) ProtoMessage() {}

func (x *TokenWithPermission) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenWithPermission.ProtoReflect.Descriptor instead.
func (*TokenWithPermission) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{6}
}

func (x *TokenWithPermission) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

func (x *TokenWithPermission) GetPermission() string {
	if x != nil && x.Permission != nil {
		return *x.Permission
	}
	return ""
}

type TokenWithUrl struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token *string `protobuf:"bytes,1,req,name=token" json:"token,omitempty"`
	Url   *string `protobuf:"bytes,2,req,name=url" json:"url,omitempty"`
}

func (x *TokenWithUrl) Reset() {
	*x = TokenWithUrl{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenWithUrl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenWithUrl) ProtoMessage() {}

func (x *TokenWithUrl) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenWithUrl.ProtoReflect.Descriptor instead.
func (*TokenWithUrl) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{7}
}

func (x *TokenWithUrl) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

func (x *TokenWithUrl) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

var File_user_service_proto protoreflect.FileDescriptor

var file_user_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x21, 0x0a, 0x09, 0x41, 0x75,
	0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x20, 0x0a,
	0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x33, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x22, 0x63, 0x0a, 0x0f, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x02, 0x28, 0x09, 0x52, 0x0a, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x46, 0x0a, 0x0c, 0x53, 0x69, 0x67,
	0x6e, 0x75, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x22, 0x5b, 0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x02,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x03, 0x20, 0x02, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x4b,
	0x0a, 0x13, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52,
	0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x36, 0x0a, 0x0c, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x55, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x32, 0xa9, 0x04, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x12, 0x12, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x1a, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x07, 0x44, 0x72, 0x6f, 0x70, 0x6f, 0x75, 0x74, 0x12,
	0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x1a, 0x0c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00,
	0x12, 0x30, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x12, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x55, 0x72, 0x6c,
	0x22, 0x00, 0x12, 0x29, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x0f, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x0c, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x37, 0x0a,
	0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x1a, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x11, 0x55, 0x6e, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x0f, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x0c, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0c,
	0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x0c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x65, 0x12, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0c, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x12, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x19, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x57, 0x69,
	0x74, 0x68, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x0c, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x14,
	0x55, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x57, 0x69, 0x74, 0x68, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x1a,
	0x0c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42,
	0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x74,
	0x63, 0x68, 0x61, 0x6e, 0x2f, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x67, 0x72, 0x70, 0x63,
}

var (
	file_user_service_proto_rawDescOnce sync.Once
	file_user_service_proto_rawDescData = file_user_service_proto_rawDesc
)

func file_user_service_proto_rawDescGZIP() []byte {
	file_user_service_proto_rawDescOnce.Do(func() {
		file_user_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_service_proto_rawDescData)
	})
	return file_user_service_proto_rawDescData
}

var file_user_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_user_service_proto_goTypes = []interface{}{
	(*AuthToken)(nil),           // 0: grpc.AuthToken
	(*Result)(nil),              // 1: grpc.Result
	(*ServiceName)(nil),         // 2: grpc.ServiceName
	(*AuthorizeParams)(nil),     // 3: grpc.AuthorizeParams
	(*SignupParams)(nil),        // 4: grpc.SignupParams
	(*LoginParams)(nil),         // 5: grpc.LoginParams
	(*TokenWithPermission)(nil), // 6: grpc.TokenWithPermission
	(*TokenWithUrl)(nil),        // 7: grpc.TokenWithUrl
}
var file_user_service_proto_depIdxs = []int32{
	4,  // 0: grpc.UserService.Signup:input_type -> grpc.SignupParams
	0,  // 1: grpc.UserService.Dropout:input_type -> grpc.AuthToken
	5,  // 2: grpc.UserService.Login:input_type -> grpc.LoginParams
	0,  // 3: grpc.UserService.Logout:input_type -> grpc.AuthToken
	2,  // 4: grpc.UserService.RegisterService:input_type -> grpc.ServiceName
	0,  // 5: grpc.UserService.UnregisterService:input_type -> grpc.AuthToken
	6,  // 6: grpc.UserService.Authenticate:input_type -> grpc.TokenWithPermission
	3,  // 7: grpc.UserService.Authorize:input_type -> grpc.AuthorizeParams
	6,  // 8: grpc.UserService.RegisterPermission:input_type -> grpc.TokenWithPermission
	6,  // 9: grpc.UserService.UnregisterPermission:input_type -> grpc.TokenWithPermission
	0,  // 10: grpc.UserService.Signup:output_type -> grpc.AuthToken
	1,  // 11: grpc.UserService.Dropout:output_type -> grpc.Result
	7,  // 12: grpc.UserService.Login:output_type -> grpc.TokenWithUrl
	1,  // 13: grpc.UserService.Logout:output_type -> grpc.Result
	0,  // 14: grpc.UserService.RegisterService:output_type -> grpc.AuthToken
	1,  // 15: grpc.UserService.UnregisterService:output_type -> grpc.Result
	1,  // 16: grpc.UserService.Authenticate:output_type -> grpc.Result
	1,  // 17: grpc.UserService.Authorize:output_type -> grpc.Result
	1,  // 18: grpc.UserService.RegisterPermission:output_type -> grpc.Result
	1,  // 19: grpc.UserService.UnregisterPermission:output_type -> grpc.Result
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_user_service_proto_init() }
func file_user_service_proto_init() {
	if File_user_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthToken); i {
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
		file_user_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_user_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceName); i {
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
		file_user_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizeParams); i {
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
		file_user_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignupParams); i {
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
		file_user_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginParams); i {
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
		file_user_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenWithPermission); i {
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
		file_user_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenWithUrl); i {
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
			RawDescriptor: file_user_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_service_proto_goTypes,
		DependencyIndexes: file_user_service_proto_depIdxs,
		MessageInfos:      file_user_service_proto_msgTypes,
	}.Build()
	File_user_service_proto = out.File
	file_user_service_proto_rawDesc = nil
	file_user_service_proto_goTypes = nil
	file_user_service_proto_depIdxs = nil
}
