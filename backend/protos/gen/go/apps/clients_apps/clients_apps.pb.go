// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.1
// source: apps/clients_apps/clients_apps.proto

package client_apps

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ClientApp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ClientId      string                 `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"` // UUID формата (обязательно)
	AppId         int32                  `protobuf:"varint,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`         // Положительное целое (существующий apps.id)
	IsActive      bool                   `protobuf:"varint,3,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"` // Автоматически генерируется
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"` // Автоматически обновляется
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ClientApp) Reset() {
	*x = ClientApp{}
	mi := &file_apps_clients_apps_clients_apps_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClientApp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientApp) ProtoMessage() {}

func (x *ClientApp) ProtoReflect() protoreflect.Message {
	mi := &file_apps_clients_apps_clients_apps_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientApp.ProtoReflect.Descriptor instead.
func (*ClientApp) Descriptor() ([]byte, []int) {
	return file_apps_clients_apps_clients_apps_proto_rawDescGZIP(), []int{0}
}

func (x *ClientApp) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *ClientApp) GetAppId() int32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *ClientApp) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *ClientApp) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ClientApp) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

// Create
type CreateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ClientId      string                 `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`        // UUID, обязательное
	AppId         int32                  `protobuf:"varint,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`                // > 0, обязательное
	IsActive      *bool                  `protobuf:"varint,3,opt,name=is_active,json=isActive,proto3,oneof" json:"is_active,omitempty"` // Дефолт: true
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	mi := &file_apps_clients_apps_clients_apps_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_clients_apps_clients_apps_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_apps_clients_apps_clients_apps_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *CreateRequest) GetAppId() int32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *CreateRequest) GetIsActive() bool {
	if x != nil && x.IsActive != nil {
		return *x.IsActive
	}
	return false
}

// Read
type GetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ClientId      string                 `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"` // UUID, обязательное
	AppId         int32                  `protobuf:"varint,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`         // > 0, обязательное
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	mi := &file_apps_clients_apps_clients_apps_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_clients_apps_clients_apps_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_apps_clients_apps_clients_apps_proto_rawDescGZIP(), []int{2}
}

func (x *GetRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *GetRequest) GetAppId() int32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

// Update
type UpdateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ClientId      string                 `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"` // UUID, обязательное
	AppId         int32                  `protobuf:"varint,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`         // > 0, обязательное
	IsActive      *bool                  `protobuf:"varint,3,opt,name=is_active,json=isActive,proto3,oneof" json:"is_active,omitempty"`
	UpdateMask    *fieldmaskpb.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"` // Допустимые поля: "is_active"
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	mi := &file_apps_clients_apps_clients_apps_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_clients_apps_clients_apps_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_apps_clients_apps_clients_apps_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *UpdateRequest) GetAppId() int32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *UpdateRequest) GetIsActive() bool {
	if x != nil && x.IsActive != nil {
		return *x.IsActive
	}
	return false
}

func (x *UpdateRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

// Delete
type DeleteRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ClientId      string                 `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"` // UUID, обязательное
	AppId         int32                  `protobuf:"varint,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`         // > 0, обязательное
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	mi := &file_apps_clients_apps_clients_apps_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_clients_apps_clients_apps_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_apps_clients_apps_clients_apps_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *DeleteRequest) GetAppId() int32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

type DeleteResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	mi := &file_apps_clients_apps_clients_apps_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apps_clients_apps_clients_apps_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_apps_clients_apps_clients_apps_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// List
type ListRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Page          int32                  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`   // >= 1
	Count         int32                  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"` // 1-100
	Filter        *ListRequest_Filter    `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	mi := &file_apps_clients_apps_clients_apps_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_clients_apps_clients_apps_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_apps_clients_apps_clients_apps_proto_rawDescGZIP(), []int{6}
}

func (x *ListRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListRequest) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ListRequest) GetFilter() *ListRequest_Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type ListResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ClientApps    []*ClientApp           `protobuf:"bytes,1,rep,name=client_apps,json=clientApps,proto3" json:"client_apps,omitempty"`
	TotalCount    int32                  `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	Page          int32                  `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Count         int32                  `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	mi := &file_apps_clients_apps_clients_apps_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apps_clients_apps_clients_apps_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_apps_clients_apps_clients_apps_proto_rawDescGZIP(), []int{7}
}

func (x *ListResponse) GetClientApps() []*ClientApp {
	if x != nil {
		return x.ClientApps
	}
	return nil
}

func (x *ListResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *ListResponse) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type ListRequest_Filter struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ClientId      *string                `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3,oneof" json:"client_id,omitempty"` // UUID
	AppId         *int32                 `protobuf:"varint,2,opt,name=app_id,json=appId,proto3,oneof" json:"app_id,omitempty"`         // > 0
	IsActive      *bool                  `protobuf:"varint,3,opt,name=is_active,json=isActive,proto3,oneof" json:"is_active,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRequest_Filter) Reset() {
	*x = ListRequest_Filter{}
	mi := &file_apps_clients_apps_clients_apps_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRequest_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest_Filter) ProtoMessage() {}

func (x *ListRequest_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_apps_clients_apps_clients_apps_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest_Filter.ProtoReflect.Descriptor instead.
func (*ListRequest_Filter) Descriptor() ([]byte, []int) {
	return file_apps_clients_apps_clients_apps_proto_rawDescGZIP(), []int{6, 0}
}

func (x *ListRequest_Filter) GetClientId() string {
	if x != nil && x.ClientId != nil {
		return *x.ClientId
	}
	return ""
}

func (x *ListRequest_Filter) GetAppId() int32 {
	if x != nil && x.AppId != nil {
		return *x.AppId
	}
	return 0
}

func (x *ListRequest_Filter) GetIsActive() bool {
	if x != nil && x.IsActive != nil {
		return *x.IsActive
	}
	return false
}

var File_apps_clients_apps_clients_apps_proto protoreflect.FileDescriptor

var file_apps_clients_apps_clients_apps_proto_rawDesc = string([]byte{
	0x0a, 0x24, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x61,
	0x70, 0x70, 0x73, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x61, 0x70, 0x70, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x61, 0x70, 0x70, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd2, 0x01, 0x0a, 0x09,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x70, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0x73, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x15,
	0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x69, 0x73, 0x5f, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x40, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x22, 0xb0, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x20, 0x0a,
	0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x48, 0x00, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b,
	0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x42, 0x0c, 0x0a, 0x0a,
	0x5f, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x43, 0x0a, 0x0d, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x22,
	0x2a, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x87, 0x02, 0x0a, 0x0b,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3c, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x1a, 0x8f, 0x01, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x20,
	0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x1a, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x48, 0x01, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09,
	0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48,
	0x02, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0c,
	0x0a, 0x0a, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x42, 0x09, 0x0a, 0x07,
	0x5f, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x69, 0x73, 0x5f, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x97, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x61, 0x70, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x70,
	0x70, 0x73, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x70, 0x70, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x41, 0x70, 0x70, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32,
	0xf9, 0x02, 0x0a, 0x11, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x41, 0x70, 0x70, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x1f, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x70,
	0x70, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x61,
	0x70, 0x70, 0x73, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x70, 0x70, 0x12, 0x40, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x70, 0x70, 0x12,
	0x46, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x2e, 0x61, 0x70, 0x70, 0x73,
	0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x70,
	0x73, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x41, 0x70, 0x70, 0x12, 0x4b, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x1f, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x61, 0x70, 0x70, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1d, 0x2e, 0x61,
	0x70, 0x70, 0x73, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x70, 0x70, 0x73, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70,
	0x70, 0x73, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x21, 0x5a, 0x1f, 0x61,
	0x70, 0x70, 0x73, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x70, 0x70, 0x73, 0x2e,
	0x76, 0x31, 0x3b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x70, 0x70, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_apps_clients_apps_clients_apps_proto_rawDescOnce sync.Once
	file_apps_clients_apps_clients_apps_proto_rawDescData []byte
)

func file_apps_clients_apps_clients_apps_proto_rawDescGZIP() []byte {
	file_apps_clients_apps_clients_apps_proto_rawDescOnce.Do(func() {
		file_apps_clients_apps_clients_apps_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_apps_clients_apps_clients_apps_proto_rawDesc), len(file_apps_clients_apps_clients_apps_proto_rawDesc)))
	})
	return file_apps_clients_apps_clients_apps_proto_rawDescData
}

var file_apps_clients_apps_clients_apps_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_apps_clients_apps_clients_apps_proto_goTypes = []any{
	(*ClientApp)(nil),             // 0: apps.client_apps.ClientApp
	(*CreateRequest)(nil),         // 1: apps.client_apps.CreateRequest
	(*GetRequest)(nil),            // 2: apps.client_apps.GetRequest
	(*UpdateRequest)(nil),         // 3: apps.client_apps.UpdateRequest
	(*DeleteRequest)(nil),         // 4: apps.client_apps.DeleteRequest
	(*DeleteResponse)(nil),        // 5: apps.client_apps.DeleteResponse
	(*ListRequest)(nil),           // 6: apps.client_apps.ListRequest
	(*ListResponse)(nil),          // 7: apps.client_apps.ListResponse
	(*ListRequest_Filter)(nil),    // 8: apps.client_apps.ListRequest.Filter
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
	(*fieldmaskpb.FieldMask)(nil), // 10: google.protobuf.FieldMask
}
var file_apps_clients_apps_clients_apps_proto_depIdxs = []int32{
	9,  // 0: apps.client_apps.ClientApp.created_at:type_name -> google.protobuf.Timestamp
	9,  // 1: apps.client_apps.ClientApp.updated_at:type_name -> google.protobuf.Timestamp
	10, // 2: apps.client_apps.UpdateRequest.update_mask:type_name -> google.protobuf.FieldMask
	8,  // 3: apps.client_apps.ListRequest.filter:type_name -> apps.client_apps.ListRequest.Filter
	0,  // 4: apps.client_apps.ListResponse.client_apps:type_name -> apps.client_apps.ClientApp
	1,  // 5: apps.client_apps.ClientsAppService.Create:input_type -> apps.client_apps.CreateRequest
	2,  // 6: apps.client_apps.ClientsAppService.Get:input_type -> apps.client_apps.GetRequest
	3,  // 7: apps.client_apps.ClientsAppService.Update:input_type -> apps.client_apps.UpdateRequest
	4,  // 8: apps.client_apps.ClientsAppService.Delete:input_type -> apps.client_apps.DeleteRequest
	6,  // 9: apps.client_apps.ClientsAppService.List:input_type -> apps.client_apps.ListRequest
	0,  // 10: apps.client_apps.ClientsAppService.Create:output_type -> apps.client_apps.ClientApp
	0,  // 11: apps.client_apps.ClientsAppService.Get:output_type -> apps.client_apps.ClientApp
	0,  // 12: apps.client_apps.ClientsAppService.Update:output_type -> apps.client_apps.ClientApp
	5,  // 13: apps.client_apps.ClientsAppService.Delete:output_type -> apps.client_apps.DeleteResponse
	7,  // 14: apps.client_apps.ClientsAppService.List:output_type -> apps.client_apps.ListResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_apps_clients_apps_clients_apps_proto_init() }
func file_apps_clients_apps_clients_apps_proto_init() {
	if File_apps_clients_apps_clients_apps_proto != nil {
		return
	}
	file_apps_clients_apps_clients_apps_proto_msgTypes[1].OneofWrappers = []any{}
	file_apps_clients_apps_clients_apps_proto_msgTypes[3].OneofWrappers = []any{}
	file_apps_clients_apps_clients_apps_proto_msgTypes[8].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_apps_clients_apps_clients_apps_proto_rawDesc), len(file_apps_clients_apps_clients_apps_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_clients_apps_clients_apps_proto_goTypes,
		DependencyIndexes: file_apps_clients_apps_clients_apps_proto_depIdxs,
		MessageInfos:      file_apps_clients_apps_clients_apps_proto_msgTypes,
	}.Build()
	File_apps_clients_apps_clients_apps_proto = out.File
	file_apps_clients_apps_clients_apps_proto_goTypes = nil
	file_apps_clients_apps_clients_apps_proto_depIdxs = nil
}
