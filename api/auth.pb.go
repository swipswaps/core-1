// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.4
// source: auth.proto

package api

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type IsAuthorizedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace    string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Verb         string `protobuf:"bytes,2,opt,name=verb,proto3" json:"verb,omitempty"`
	Group        string `protobuf:"bytes,3,opt,name=group,proto3" json:"group,omitempty"`
	Resource     string `protobuf:"bytes,4,opt,name=resource,proto3" json:"resource,omitempty"`
	ResourceName string `protobuf:"bytes,5,opt,name=resourceName,proto3" json:"resourceName,omitempty"`
}

func (x *IsAuthorizedRequest) Reset() {
	*x = IsAuthorizedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsAuthorizedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsAuthorizedRequest) ProtoMessage() {}

func (x *IsAuthorizedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsAuthorizedRequest.ProtoReflect.Descriptor instead.
func (*IsAuthorizedRequest) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{0}
}

func (x *IsAuthorizedRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *IsAuthorizedRequest) GetVerb() string {
	if x != nil {
		return x.Verb
	}
	return ""
}

func (x *IsAuthorizedRequest) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *IsAuthorizedRequest) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *IsAuthorizedRequest) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

type IsAuthorizedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Authorized bool `protobuf:"varint,1,opt,name=authorized,proto3" json:"authorized,omitempty"`
}

func (x *IsAuthorizedResponse) Reset() {
	*x = IsAuthorizedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsAuthorizedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsAuthorizedResponse) ProtoMessage() {}

func (x *IsAuthorizedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsAuthorizedResponse.ProtoReflect.Descriptor instead.
func (*IsAuthorizedResponse) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{1}
}

func (x *IsAuthorizedResponse) GetAuthorized() bool {
	if x != nil {
		return x.Authorized
	}
	return false
}

type TokenWrapper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *TokenWrapper) Reset() {
	*x = TokenWrapper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenWrapper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenWrapper) ProtoMessage() {}

func (x *TokenWrapper) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenWrapper.ProtoReflect.Descriptor instead.
func (*TokenWrapper) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{2}
}

func (x *TokenWrapper) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type IsValidTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token *TokenWrapper `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *IsValidTokenRequest) Reset() {
	*x = IsValidTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsValidTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsValidTokenRequest) ProtoMessage() {}

func (x *IsValidTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsValidTokenRequest.ProtoReflect.Descriptor instead.
func (*IsValidTokenRequest) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{3}
}

func (x *IsValidTokenRequest) GetToken() *TokenWrapper {
	if x != nil {
		return x.Token
	}
	return nil
}

type IsValidTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
}

func (x *IsValidTokenResponse) Reset() {
	*x = IsValidTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsValidTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsValidTokenResponse) ProtoMessage() {}

func (x *IsValidTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsValidTokenResponse.ProtoReflect.Descriptor instead.
func (*IsValidTokenResponse) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{4}
}

func (x *IsValidTokenResponse) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

type IsWorkspaceAuthenticatedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FQDN               string `protobuf:"bytes,1,opt,name=FQDN,proto3" json:"FQDN,omitempty"`
	XOriginalMethod    string `protobuf:"bytes,2,opt,name=XOriginalMethod,proto3" json:"XOriginalMethod,omitempty"`
	XOriginalAuthority string `protobuf:"bytes,3,opt,name=XOriginalAuthority,proto3" json:"XOriginalAuthority,omitempty"`
	XOriginalUri       string `protobuf:"bytes,4,opt,name=XOriginalUri,proto3" json:"XOriginalUri,omitempty"`
}

func (x *IsWorkspaceAuthenticatedRequest) Reset() {
	*x = IsWorkspaceAuthenticatedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsWorkspaceAuthenticatedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsWorkspaceAuthenticatedRequest) ProtoMessage() {}

func (x *IsWorkspaceAuthenticatedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsWorkspaceAuthenticatedRequest.ProtoReflect.Descriptor instead.
func (*IsWorkspaceAuthenticatedRequest) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{5}
}

func (x *IsWorkspaceAuthenticatedRequest) GetFQDN() string {
	if x != nil {
		return x.FQDN
	}
	return ""
}

func (x *IsWorkspaceAuthenticatedRequest) GetXOriginalMethod() string {
	if x != nil {
		return x.XOriginalMethod
	}
	return ""
}

func (x *IsWorkspaceAuthenticatedRequest) GetXOriginalAuthority() string {
	if x != nil {
		return x.XOriginalAuthority
	}
	return ""
}

func (x *IsWorkspaceAuthenticatedRequest) GetXOriginalUri() string {
	if x != nil {
		return x.XOriginalUri
	}
	return ""
}

var File_auth_proto protoreflect.FileDescriptor

var file_auth_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70,
	0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x01, 0x0a,
	0x13, 0x49, 0x73, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x65, 0x72, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x76, 0x65, 0x72, 0x62, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x36, 0x0a, 0x14,
	0x49, 0x73, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x65, 0x64, 0x22, 0x24, 0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x57, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x3e, 0x0a, 0x13, 0x49, 0x73,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x27, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x57, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x2e, 0x0a, 0x14, 0x49, 0x73,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0xb3, 0x01, 0x0a, 0x1f, 0x49,
	0x73, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x46, 0x51, 0x44, 0x4e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x46, 0x51,
	0x44, 0x4e, 0x12, 0x28, 0x0a, 0x0f, 0x58, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x58, 0x4f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x2e, 0x0a, 0x12,
	0x58, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x58, 0x4f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x61, 0x6c, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x22, 0x0a, 0x0c,
	0x58, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x69, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x58, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x69,
	0x32, 0xe7, 0x02, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x6c, 0x0a, 0x0c, 0x49, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x49, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x22, 0x18, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x3a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x82,
	0x01, 0x0a, 0x18, 0x49, 0x73, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x41, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x12, 0x24, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x49, 0x73, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x41, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x22, 0x42, 0x20, 0x0a, 0x04, 0x48, 0x45, 0x41, 0x44, 0x12, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x65, 0x0a, 0x0c, 0x49, 0x73, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x65, 0x64, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x73, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x49, 0x73, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a,
	0x22, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_auth_proto_rawDescOnce sync.Once
	file_auth_proto_rawDescData = file_auth_proto_rawDesc
)

func file_auth_proto_rawDescGZIP() []byte {
	file_auth_proto_rawDescOnce.Do(func() {
		file_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_proto_rawDescData)
	})
	return file_auth_proto_rawDescData
}

var file_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_auth_proto_goTypes = []interface{}{
	(*IsAuthorizedRequest)(nil),             // 0: api.IsAuthorizedRequest
	(*IsAuthorizedResponse)(nil),            // 1: api.IsAuthorizedResponse
	(*TokenWrapper)(nil),                    // 2: api.TokenWrapper
	(*IsValidTokenRequest)(nil),             // 3: api.IsValidTokenRequest
	(*IsValidTokenResponse)(nil),            // 4: api.IsValidTokenResponse
	(*IsWorkspaceAuthenticatedRequest)(nil), // 5: api.IsWorkspaceAuthenticatedRequest
	(*empty.Empty)(nil),                     // 6: google.protobuf.Empty
}
var file_auth_proto_depIdxs = []int32{
	2, // 0: api.IsValidTokenRequest.token:type_name -> api.TokenWrapper
	3, // 1: api.AuthService.IsValidToken:input_type -> api.IsValidTokenRequest
	5, // 2: api.AuthService.IsWorkspaceAuthenticated:input_type -> api.IsWorkspaceAuthenticatedRequest
	0, // 3: api.AuthService.IsAuthorized:input_type -> api.IsAuthorizedRequest
	4, // 4: api.AuthService.IsValidToken:output_type -> api.IsValidTokenResponse
	6, // 5: api.AuthService.IsWorkspaceAuthenticated:output_type -> google.protobuf.Empty
	1, // 6: api.AuthService.IsAuthorized:output_type -> api.IsAuthorizedResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_auth_proto_init() }
func file_auth_proto_init() {
	if File_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsAuthorizedRequest); i {
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
		file_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsAuthorizedResponse); i {
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
		file_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenWrapper); i {
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
		file_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsValidTokenRequest); i {
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
		file_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsValidTokenResponse); i {
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
		file_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsWorkspaceAuthenticatedRequest); i {
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
			RawDescriptor: file_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_proto_goTypes,
		DependencyIndexes: file_auth_proto_depIdxs,
		MessageInfos:      file_auth_proto_msgTypes,
	}.Build()
	File_auth_proto = out.File
	file_auth_proto_rawDesc = nil
	file_auth_proto_goTypes = nil
	file_auth_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthServiceClient interface {
	IsValidToken(ctx context.Context, in *IsValidTokenRequest, opts ...grpc.CallOption) (*IsValidTokenResponse, error)
	IsWorkspaceAuthenticated(ctx context.Context, in *IsWorkspaceAuthenticatedRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	IsAuthorized(ctx context.Context, in *IsAuthorizedRequest, opts ...grpc.CallOption) (*IsAuthorizedResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) IsValidToken(ctx context.Context, in *IsValidTokenRequest, opts ...grpc.CallOption) (*IsValidTokenResponse, error) {
	out := new(IsValidTokenResponse)
	err := c.cc.Invoke(ctx, "/api.AuthService/IsValidToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) IsWorkspaceAuthenticated(ctx context.Context, in *IsWorkspaceAuthenticatedRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.AuthService/IsWorkspaceAuthenticated", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) IsAuthorized(ctx context.Context, in *IsAuthorizedRequest, opts ...grpc.CallOption) (*IsAuthorizedResponse, error) {
	out := new(IsAuthorizedResponse)
	err := c.cc.Invoke(ctx, "/api.AuthService/IsAuthorized", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	IsValidToken(context.Context, *IsValidTokenRequest) (*IsValidTokenResponse, error)
	IsWorkspaceAuthenticated(context.Context, *IsWorkspaceAuthenticatedRequest) (*empty.Empty, error)
	IsAuthorized(context.Context, *IsAuthorizedRequest) (*IsAuthorizedResponse, error)
}

// UnimplementedAuthServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (*UnimplementedAuthServiceServer) IsValidToken(context.Context, *IsValidTokenRequest) (*IsValidTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsValidToken not implemented")
}
func (*UnimplementedAuthServiceServer) IsWorkspaceAuthenticated(context.Context, *IsWorkspaceAuthenticatedRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsWorkspaceAuthenticated not implemented")
}
func (*UnimplementedAuthServiceServer) IsAuthorized(context.Context, *IsAuthorizedRequest) (*IsAuthorizedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsAuthorized not implemented")
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_IsValidToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsValidTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).IsValidToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AuthService/IsValidToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).IsValidToken(ctx, req.(*IsValidTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_IsWorkspaceAuthenticated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsWorkspaceAuthenticatedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).IsWorkspaceAuthenticated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AuthService/IsWorkspaceAuthenticated",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).IsWorkspaceAuthenticated(ctx, req.(*IsWorkspaceAuthenticatedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_IsAuthorized_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsAuthorizedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).IsAuthorized(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AuthService/IsAuthorized",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).IsAuthorized(ctx, req.(*IsAuthorizedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsValidToken",
			Handler:    _AuthService_IsValidToken_Handler,
		},
		{
			MethodName: "IsWorkspaceAuthenticated",
			Handler:    _AuthService_IsWorkspaceAuthenticated_Handler,
		},
		{
			MethodName: "IsAuthorized",
			Handler:    _AuthService_IsAuthorized_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
