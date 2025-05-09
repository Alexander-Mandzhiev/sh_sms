// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.1
// source: apps/secret_rotation_history/secret_rotation_history.proto

package secrets_history

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type RotationHistory struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ClientId      string                 `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`          // UUID клиента (обязательно)
	AppId         int32                  `protobuf:"varint,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`                  // Положительное целое
	SecretType    string                 `protobuf:"bytes,3,opt,name=secret_type,json=secretType,proto3" json:"secret_type,omitempty"`    // "access" или "refresh"
	OldSecret     string                 `protobuf:"bytes,4,opt,name=old_secret,json=oldSecret,proto3" json:"old_secret,omitempty"`       // Длина до 512 символов
	NewSecret     string                 `protobuf:"bytes,5,opt,name=new_secret,json=newSecret,proto3" json:"new_secret,omitempty"`       // Длина до 512 символов
	RotatedBy     *string                `protobuf:"bytes,6,opt,name=rotated_by,json=rotatedBy,proto3,oneof" json:"rotated_by,omitempty"` // UUID инициатора (опционально)
	RotatedAt     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=rotated_at,json=rotatedAt,proto3" json:"rotated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RotationHistory) Reset() {
	*x = RotationHistory{}
	mi := &file_apps_secret_rotation_history_secret_rotation_history_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RotationHistory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RotationHistory) ProtoMessage() {}

func (x *RotationHistory) ProtoReflect() protoreflect.Message {
	mi := &file_apps_secret_rotation_history_secret_rotation_history_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RotationHistory.ProtoReflect.Descriptor instead.
func (*RotationHistory) Descriptor() ([]byte, []int) {
	return file_apps_secret_rotation_history_secret_rotation_history_proto_rawDescGZIP(), []int{0}
}

func (x *RotationHistory) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *RotationHistory) GetAppId() int32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *RotationHistory) GetSecretType() string {
	if x != nil {
		return x.SecretType
	}
	return ""
}

func (x *RotationHistory) GetOldSecret() string {
	if x != nil {
		return x.OldSecret
	}
	return ""
}

func (x *RotationHistory) GetNewSecret() string {
	if x != nil {
		return x.NewSecret
	}
	return ""
}

func (x *RotationHistory) GetRotatedBy() string {
	if x != nil && x.RotatedBy != nil {
		return *x.RotatedBy
	}
	return ""
}

func (x *RotationHistory) GetRotatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.RotatedAt
	}
	return nil
}

type GetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ClientId      string                 `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	AppId         int32                  `protobuf:"varint,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	SecretType    string                 `protobuf:"bytes,3,opt,name=secret_type,json=secretType,proto3" json:"secret_type,omitempty"`
	RotatedAt     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=rotated_at,json=rotatedAt,proto3" json:"rotated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	mi := &file_apps_secret_rotation_history_secret_rotation_history_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_secret_rotation_history_secret_rotation_history_proto_msgTypes[1]
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
	return file_apps_secret_rotation_history_secret_rotation_history_proto_rawDescGZIP(), []int{1}
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

func (x *GetRequest) GetSecretType() string {
	if x != nil {
		return x.SecretType
	}
	return ""
}

func (x *GetRequest) GetRotatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.RotatedAt
	}
	return nil
}

type ListRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Page          int64                  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Count         int64                  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Filter        *ListRequest_Filter    `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	mi := &file_apps_secret_rotation_history_secret_rotation_history_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_secret_rotation_history_secret_rotation_history_proto_msgTypes[2]
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
	return file_apps_secret_rotation_history_secret_rotation_history_proto_rawDescGZIP(), []int{2}
}

func (x *ListRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListRequest) GetCount() int64 {
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
	Rotations     []*RotationHistory     `protobuf:"bytes,1,rep,name=rotations,proto3" json:"rotations,omitempty"`
	TotalCount    int32                  `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	Page          int64                  `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Count         int64                  `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	mi := &file_apps_secret_rotation_history_secret_rotation_history_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apps_secret_rotation_history_secret_rotation_history_proto_msgTypes[3]
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
	return file_apps_secret_rotation_history_secret_rotation_history_proto_rawDescGZIP(), []int{3}
}

func (x *ListResponse) GetRotations() []*RotationHistory {
	if x != nil {
		return x.Rotations
	}
	return nil
}

func (x *ListResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *ListResponse) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type ListRequest_Filter struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ClientId      *string                `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3,oneof" json:"client_id,omitempty"`
	AppId         *int32                 `protobuf:"varint,2,opt,name=app_id,json=appId,proto3,oneof" json:"app_id,omitempty"`
	SecretType    *string                `protobuf:"bytes,3,opt,name=secret_type,json=secretType,proto3,oneof" json:"secret_type,omitempty"`
	RotatedBy     *string                `protobuf:"bytes,4,opt,name=rotated_by,json=rotatedBy,proto3,oneof" json:"rotated_by,omitempty"`
	RotatedAfter  *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=rotated_after,json=rotatedAfter,proto3,oneof" json:"rotated_after,omitempty"`
	RotatedBefore *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=rotated_before,json=rotatedBefore,proto3,oneof" json:"rotated_before,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRequest_Filter) Reset() {
	*x = ListRequest_Filter{}
	mi := &file_apps_secret_rotation_history_secret_rotation_history_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRequest_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest_Filter) ProtoMessage() {}

func (x *ListRequest_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_apps_secret_rotation_history_secret_rotation_history_proto_msgTypes[4]
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
	return file_apps_secret_rotation_history_secret_rotation_history_proto_rawDescGZIP(), []int{2, 0}
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

func (x *ListRequest_Filter) GetSecretType() string {
	if x != nil && x.SecretType != nil {
		return *x.SecretType
	}
	return ""
}

func (x *ListRequest_Filter) GetRotatedBy() string {
	if x != nil && x.RotatedBy != nil {
		return *x.RotatedBy
	}
	return ""
}

func (x *ListRequest_Filter) GetRotatedAfter() *timestamppb.Timestamp {
	if x != nil {
		return x.RotatedAfter
	}
	return nil
}

func (x *ListRequest_Filter) GetRotatedBefore() *timestamppb.Timestamp {
	if x != nil {
		return x.RotatedBefore
	}
	return nil
}

var File_apps_secret_rotation_history_secret_rotation_history_proto protoreflect.FileDescriptor

var file_apps_secret_rotation_history_secret_rotation_history_proto_rawDesc = string([]byte{
	0x0a, 0x3a, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x72, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x68,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x61, 0x70,
	0x70, 0x73, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x92, 0x02, 0x0a, 0x0f, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x6f, 0x6c, 0x64, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6f, 0x6c, 0x64, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6e,
	0x65, 0x77, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6e, 0x65, 0x77, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x22, 0x0a, 0x0a, 0x72, 0x6f,
	0x74, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x09, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x88, 0x01, 0x01, 0x12, 0x39,
	0x0a, 0x0a, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x72, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x72, 0x6f,
	0x74, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x22, 0x9c, 0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x39, 0x0a, 0x0a,
	0x72, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x72, 0x6f,
	0x74, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xf7, 0x03, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x40, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73,
	0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x1a, 0xfb, 0x02, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x20,
	0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x1a, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x48, 0x01, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x02, 0x52, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x09, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x65,
	0x64, 0x42, 0x79, 0x88, 0x01, 0x01, 0x12, 0x44, 0x0a, 0x0d, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x04, 0x52, 0x0c, 0x72, 0x6f, 0x74,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x66, 0x74, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x46, 0x0a, 0x0e,
	0x72, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x48, 0x05, 0x52, 0x0d, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x42, 0x65, 0x66, 0x6f, 0x72,
	0x65, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x42, 0x0e, 0x0a,
	0x0c, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x0d, 0x0a,
	0x0b, 0x5f, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x42, 0x10, 0x0a, 0x0e,
	0x5f, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x42, 0x11,
	0x0a, 0x0f, 0x5f, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x65, 0x66, 0x6f, 0x72,
	0x65, 0x22, 0x9e, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x43, 0x0a, 0x09, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x73, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x52, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x09, 0x72, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x32, 0xc8, 0x01, 0x0a, 0x16, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x56, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x2e, 0x61,
	0x70, 0x70, 0x73, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x5f, 0x68, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25,
	0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x5f, 0x68, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x56, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x73, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x70, 0x73,
	0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x29, 0x5a,
	0x27, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x5f, 0x68, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73,
	0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_apps_secret_rotation_history_secret_rotation_history_proto_rawDescOnce sync.Once
	file_apps_secret_rotation_history_secret_rotation_history_proto_rawDescData []byte
)

func file_apps_secret_rotation_history_secret_rotation_history_proto_rawDescGZIP() []byte {
	file_apps_secret_rotation_history_secret_rotation_history_proto_rawDescOnce.Do(func() {
		file_apps_secret_rotation_history_secret_rotation_history_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_apps_secret_rotation_history_secret_rotation_history_proto_rawDesc), len(file_apps_secret_rotation_history_secret_rotation_history_proto_rawDesc)))
	})
	return file_apps_secret_rotation_history_secret_rotation_history_proto_rawDescData
}

var file_apps_secret_rotation_history_secret_rotation_history_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_apps_secret_rotation_history_secret_rotation_history_proto_goTypes = []any{
	(*RotationHistory)(nil),       // 0: apps.secrets_history.RotationHistory
	(*GetRequest)(nil),            // 1: apps.secrets_history.GetRequest
	(*ListRequest)(nil),           // 2: apps.secrets_history.ListRequest
	(*ListResponse)(nil),          // 3: apps.secrets_history.ListResponse
	(*ListRequest_Filter)(nil),    // 4: apps.secrets_history.ListRequest.Filter
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_apps_secret_rotation_history_secret_rotation_history_proto_depIdxs = []int32{
	5, // 0: apps.secrets_history.RotationHistory.rotated_at:type_name -> google.protobuf.Timestamp
	5, // 1: apps.secrets_history.GetRequest.rotated_at:type_name -> google.protobuf.Timestamp
	4, // 2: apps.secrets_history.ListRequest.filter:type_name -> apps.secrets_history.ListRequest.Filter
	0, // 3: apps.secrets_history.ListResponse.rotations:type_name -> apps.secrets_history.RotationHistory
	5, // 4: apps.secrets_history.ListRequest.Filter.rotated_after:type_name -> google.protobuf.Timestamp
	5, // 5: apps.secrets_history.ListRequest.Filter.rotated_before:type_name -> google.protobuf.Timestamp
	1, // 6: apps.secrets_history.RotationHistoryService.GetRotation:input_type -> apps.secrets_history.GetRequest
	2, // 7: apps.secrets_history.RotationHistoryService.ListRotations:input_type -> apps.secrets_history.ListRequest
	0, // 8: apps.secrets_history.RotationHistoryService.GetRotation:output_type -> apps.secrets_history.RotationHistory
	3, // 9: apps.secrets_history.RotationHistoryService.ListRotations:output_type -> apps.secrets_history.ListResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_apps_secret_rotation_history_secret_rotation_history_proto_init() }
func file_apps_secret_rotation_history_secret_rotation_history_proto_init() {
	if File_apps_secret_rotation_history_secret_rotation_history_proto != nil {
		return
	}
	file_apps_secret_rotation_history_secret_rotation_history_proto_msgTypes[0].OneofWrappers = []any{}
	file_apps_secret_rotation_history_secret_rotation_history_proto_msgTypes[4].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_apps_secret_rotation_history_secret_rotation_history_proto_rawDesc), len(file_apps_secret_rotation_history_secret_rotation_history_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_secret_rotation_history_secret_rotation_history_proto_goTypes,
		DependencyIndexes: file_apps_secret_rotation_history_secret_rotation_history_proto_depIdxs,
		MessageInfos:      file_apps_secret_rotation_history_secret_rotation_history_proto_msgTypes,
	}.Build()
	File_apps_secret_rotation_history_secret_rotation_history_proto = out.File
	file_apps_secret_rotation_history_secret_rotation_history_proto_goTypes = nil
	file_apps_secret_rotation_history_secret_rotation_history_proto_depIdxs = nil
}
