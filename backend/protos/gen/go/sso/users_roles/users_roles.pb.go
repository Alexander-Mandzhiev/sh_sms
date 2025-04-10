// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.1
// source: sso/users_roles/users_roles.proto

package user_roles

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

// Запрос на назначение роли
type AssignRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ClientId      string                 `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`          // Соответствует полю client_id в БД
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`                // UUID пользователя
	RoleId        string                 `protobuf:"bytes,3,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`                // UUID роли
	AssignedBy    string                 `protobuf:"bytes,4,opt,name=assigned_by,json=assignedBy,proto3" json:"assigned_by,omitempty"`    // UUID администратора, назначившего роль
	ExpiresAt     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=expires_at,json=expiresAt,proto3,oneof" json:"expires_at,omitempty"` // Опциональное время истечения
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AssignRequest) Reset() {
	*x = AssignRequest{}
	mi := &file_sso_users_roles_users_roles_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AssignRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignRequest) ProtoMessage() {}

func (x *AssignRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_users_roles_users_roles_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignRequest.ProtoReflect.Descriptor instead.
func (*AssignRequest) Descriptor() ([]byte, []int) {
	return file_sso_users_roles_users_roles_proto_rawDescGZIP(), []int{0}
}

func (x *AssignRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *AssignRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AssignRequest) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *AssignRequest) GetAssignedBy() string {
	if x != nil {
		return x.AssignedBy
	}
	return ""
}

func (x *AssignRequest) GetExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}

// Запрос на отзыв роли
type RevokeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ClientId      string                 `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RoleId        string                 `protobuf:"bytes,3,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	RevokedBy     string                 `protobuf:"bytes,4,opt,name=revoked_by,json=revokedBy,proto3" json:"revoked_by,omitempty"` // UUID администратора, отозвавшего роль
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RevokeRequest) Reset() {
	*x = RevokeRequest{}
	mi := &file_sso_users_roles_users_roles_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RevokeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevokeRequest) ProtoMessage() {}

func (x *RevokeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_users_roles_users_roles_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevokeRequest.ProtoReflect.Descriptor instead.
func (*RevokeRequest) Descriptor() ([]byte, []int) {
	return file_sso_users_roles_users_roles_proto_rawDescGZIP(), []int{1}
}

func (x *RevokeRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *RevokeRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RevokeRequest) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *RevokeRequest) GetRevokedBy() string {
	if x != nil {
		return x.RevokedBy
	}
	return ""
}

// Общий ответ для операций
type OperationResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Timestamp     *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Message       string                 `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"` // Детализация ошибки при success=false
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OperationResponse) Reset() {
	*x = OperationResponse{}
	mi := &file_sso_users_roles_users_roles_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OperationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationResponse) ProtoMessage() {}

func (x *OperationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_users_roles_users_roles_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationResponse.ProtoReflect.Descriptor instead.
func (*OperationResponse) Descriptor() ([]byte, []int) {
	return file_sso_users_roles_users_roles_proto_rawDescGZIP(), []int{2}
}

func (x *OperationResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *OperationResponse) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *OperationResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Ответ с данными о назначении
type UserRoleResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AssignmentId  string                 `protobuf:"bytes,1,opt,name=assignment_id,json=assignmentId,proto3" json:"assignment_id,omitempty"` // UUID назначения (опционально)
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	ExpiresAt     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=expires_at,json=expiresAt,proto3,oneof" json:"expires_at,omitempty"`
	AssignedBy    string                 `protobuf:"bytes,4,opt,name=assigned_by,json=assignedBy,proto3" json:"assigned_by,omitempty"` // UUID назначившего
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserRoleResponse) Reset() {
	*x = UserRoleResponse{}
	mi := &file_sso_users_roles_users_roles_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserRoleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRoleResponse) ProtoMessage() {}

func (x *UserRoleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_users_roles_users_roles_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRoleResponse.ProtoReflect.Descriptor instead.
func (*UserRoleResponse) Descriptor() ([]byte, []int) {
	return file_sso_users_roles_users_roles_proto_rawDescGZIP(), []int{3}
}

func (x *UserRoleResponse) GetAssignmentId() string {
	if x != nil {
		return x.AssignmentId
	}
	return ""
}

func (x *UserRoleResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *UserRoleResponse) GetExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}

func (x *UserRoleResponse) GetAssignedBy() string {
	if x != nil {
		return x.AssignedBy
	}
	return ""
}

// Запрос списка ролей
type ListRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	ClientId       string                 `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`                    // Обязательный контекст
	UserId         string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`                          // Фильтр по конкретному пользователю
	Page           int32                  `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`                                           // Начинается с 1 (default: 1)
	Count          int32                  `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`                                         // 1-200 элементов (default: 50)
	IncludeExpired bool                   `protobuf:"varint,5,opt,name=include_expired,json=includeExpired,proto3" json:"include_expired,omitempty"` // Включать истекшие назначения
	IncludeDeleted bool                   `protobuf:"varint,6,opt,name=include_deleted,json=includeDeleted,proto3" json:"include_deleted,omitempty"` // Включать удаленные роли
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	mi := &file_sso_users_roles_users_roles_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_users_roles_users_roles_proto_msgTypes[4]
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
	return file_sso_users_roles_users_roles_proto_rawDescGZIP(), []int{4}
}

func (x *ListRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *ListRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
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

func (x *ListRequest) GetIncludeExpired() bool {
	if x != nil {
		return x.IncludeExpired
	}
	return false
}

func (x *ListRequest) GetIncludeDeleted() bool {
	if x != nil {
		return x.IncludeDeleted
	}
	return false
}

// Ответ со списком ролей
type ListResponse struct {
	state         protoimpl.MessageState         `protogen:"open.v1"`
	Roles         []*ListResponse_RoleAssignment `protobuf:"bytes,1,rep,name=roles,proto3" json:"roles,omitempty"`
	TotalCount    int32                          `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"` // Всего назначений (без учета пагинации)
	CurrentPage   int32                          `protobuf:"varint,3,opt,name=current_page,json=currentPage,proto3" json:"current_page,omitempty"`
	ItemsPerPage  int32                          `protobuf:"varint,4,opt,name=items_per_page,json=itemsPerPage,proto3" json:"items_per_page,omitempty"` // Синхронизировано с count из запроса
	HasMore       bool                           `protobuf:"varint,5,opt,name=has_more,json=hasMore,proto3" json:"has_more,omitempty"`                  // Признак наличия следующих страниц
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	mi := &file_sso_users_roles_users_roles_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_users_roles_users_roles_proto_msgTypes[5]
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
	return file_sso_users_roles_users_roles_proto_rawDescGZIP(), []int{5}
}

func (x *ListResponse) GetRoles() []*ListResponse_RoleAssignment {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *ListResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *ListResponse) GetCurrentPage() int32 {
	if x != nil {
		return x.CurrentPage
	}
	return 0
}

func (x *ListResponse) GetItemsPerPage() int32 {
	if x != nil {
		return x.ItemsPerPage
	}
	return 0
}

func (x *ListResponse) GetHasMore() bool {
	if x != nil {
		return x.HasMore
	}
	return false
}

type ListResponse_RoleAssignment struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	RoleId          string                 `protobuf:"bytes,1,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	RoleName        string                 `protobuf:"bytes,2,opt,name=role_name,json=roleName,proto3" json:"role_name,omitempty"`                      // Актуальное имя роли из таблицы roles
	RoleDescription string                 `protobuf:"bytes,3,opt,name=role_description,json=roleDescription,proto3" json:"role_description,omitempty"` // Описание роли
	AssignedAt      *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=assigned_at,json=assignedAt,proto3" json:"assigned_at,omitempty"`
	ExpiresAt       *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=expires_at,json=expiresAt,proto3,oneof" json:"expires_at,omitempty"`
	AssignedBy      string                 `protobuf:"bytes,6,opt,name=assigned_by,json=assignedBy,proto3" json:"assigned_by,omitempty"` // UUID назначившего
	IsActive        bool                   `protobuf:"varint,7,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`      // Рассчитанное поле (expires_at > now)
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *ListResponse_RoleAssignment) Reset() {
	*x = ListResponse_RoleAssignment{}
	mi := &file_sso_users_roles_users_roles_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListResponse_RoleAssignment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse_RoleAssignment) ProtoMessage() {}

func (x *ListResponse_RoleAssignment) ProtoReflect() protoreflect.Message {
	mi := &file_sso_users_roles_users_roles_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse_RoleAssignment.ProtoReflect.Descriptor instead.
func (*ListResponse_RoleAssignment) Descriptor() ([]byte, []int) {
	return file_sso_users_roles_users_roles_proto_rawDescGZIP(), []int{5, 0}
}

func (x *ListResponse_RoleAssignment) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *ListResponse_RoleAssignment) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

func (x *ListResponse_RoleAssignment) GetRoleDescription() string {
	if x != nil {
		return x.RoleDescription
	}
	return ""
}

func (x *ListResponse_RoleAssignment) GetAssignedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.AssignedAt
	}
	return nil
}

func (x *ListResponse_RoleAssignment) GetExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}

func (x *ListResponse_RoleAssignment) GetAssignedBy() string {
	if x != nil {
		return x.AssignedBy
	}
	return ""
}

func (x *ListResponse_RoleAssignment) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

var File_sso_users_roles_users_roles_proto protoreflect.FileDescriptor

var file_sso_users_roles_users_roles_proto_rawDesc = string([]byte{
	0x0a, 0x21, 0x73, 0x73, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x72, 0x6f, 0x6c, 0x65,
	0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x73, 0x73, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x6f,
	0x6c, 0x65, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xce, 0x01, 0x0a, 0x0d, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65,
	0x64, 0x5f, 0x62, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x65, 0x64, 0x42, 0x79, 0x12, 0x3e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x73, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x73, 0x41, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x73, 0x5f, 0x61, 0x74, 0x22, 0x7d, 0x0a, 0x0d, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x64,
	0x5f, 0x62, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x76, 0x6f, 0x6b,
	0x65, 0x64, 0x42, 0x79, 0x22, 0x81, 0x01, 0x0a, 0x11, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xe2, 0x01, 0x0a, 0x10, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a,
	0x0d, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3e, 0x0a,
	0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52,
	0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a,
	0x0b, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x42, 0x79, 0x42, 0x0d,
	0x0a, 0x0b, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x22, 0xbf, 0x01,
	0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x27, 0x0a,
	0x0f, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x45,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64,
	0x65, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0e, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22,
	0x94, 0x04, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x41, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2b, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x73,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x6f,
	0x6c, 0x65, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x72, 0x6f,
	0x6c, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x5f, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0c, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x50, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x68, 0x61, 0x73, 0x5f, 0x6d, 0x6f, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x68, 0x61, 0x73, 0x4d, 0x6f, 0x72, 0x65, 0x1a, 0xbb, 0x02, 0x0a, 0x0e, 0x52, 0x6f, 0x6c,
	0x65, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x72,
	0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f,
	0x6c, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x6f, 0x6c,
	0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x0b,
	0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x61,
	0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3e, 0x0a, 0x0a, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x09, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x73, 0x73,
	0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73,
	0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69,
	0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x32, 0xf8, 0x01, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x6f, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x0a, 0x41, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x1d, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0a, 0x52, 0x65, 0x76,
	0x6f, 0x6b, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x1d, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x09, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x1b, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72,
	0x6f, 0x6c, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x1e, 0x5a, 0x1c, 0x73, 0x73, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x6f,
	0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x6f, 0x6c, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_sso_users_roles_users_roles_proto_rawDescOnce sync.Once
	file_sso_users_roles_users_roles_proto_rawDescData []byte
)

func file_sso_users_roles_users_roles_proto_rawDescGZIP() []byte {
	file_sso_users_roles_users_roles_proto_rawDescOnce.Do(func() {
		file_sso_users_roles_users_roles_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_sso_users_roles_users_roles_proto_rawDesc), len(file_sso_users_roles_users_roles_proto_rawDesc)))
	})
	return file_sso_users_roles_users_roles_proto_rawDescData
}

var file_sso_users_roles_users_roles_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_sso_users_roles_users_roles_proto_goTypes = []any{
	(*AssignRequest)(nil),               // 0: sso.user_roles.AssignRequest
	(*RevokeRequest)(nil),               // 1: sso.user_roles.RevokeRequest
	(*OperationResponse)(nil),           // 2: sso.user_roles.OperationResponse
	(*UserRoleResponse)(nil),            // 3: sso.user_roles.UserRoleResponse
	(*ListRequest)(nil),                 // 4: sso.user_roles.ListRequest
	(*ListResponse)(nil),                // 5: sso.user_roles.ListResponse
	(*ListResponse_RoleAssignment)(nil), // 6: sso.user_roles.ListResponse.RoleAssignment
	(*timestamppb.Timestamp)(nil),       // 7: google.protobuf.Timestamp
}
var file_sso_users_roles_users_roles_proto_depIdxs = []int32{
	7,  // 0: sso.user_roles.AssignRequest.expires_at:type_name -> google.protobuf.Timestamp
	7,  // 1: sso.user_roles.OperationResponse.timestamp:type_name -> google.protobuf.Timestamp
	7,  // 2: sso.user_roles.UserRoleResponse.created_at:type_name -> google.protobuf.Timestamp
	7,  // 3: sso.user_roles.UserRoleResponse.expires_at:type_name -> google.protobuf.Timestamp
	6,  // 4: sso.user_roles.ListResponse.roles:type_name -> sso.user_roles.ListResponse.RoleAssignment
	7,  // 5: sso.user_roles.ListResponse.RoleAssignment.assigned_at:type_name -> google.protobuf.Timestamp
	7,  // 6: sso.user_roles.ListResponse.RoleAssignment.expires_at:type_name -> google.protobuf.Timestamp
	0,  // 7: sso.user_roles.UserRoleService.AssignRole:input_type -> sso.user_roles.AssignRequest
	1,  // 8: sso.user_roles.UserRoleService.RevokeRole:input_type -> sso.user_roles.RevokeRequest
	4,  // 9: sso.user_roles.UserRoleService.ListRoles:input_type -> sso.user_roles.ListRequest
	3,  // 10: sso.user_roles.UserRoleService.AssignRole:output_type -> sso.user_roles.UserRoleResponse
	2,  // 11: sso.user_roles.UserRoleService.RevokeRole:output_type -> sso.user_roles.OperationResponse
	5,  // 12: sso.user_roles.UserRoleService.ListRoles:output_type -> sso.user_roles.ListResponse
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_sso_users_roles_users_roles_proto_init() }
func file_sso_users_roles_users_roles_proto_init() {
	if File_sso_users_roles_users_roles_proto != nil {
		return
	}
	file_sso_users_roles_users_roles_proto_msgTypes[0].OneofWrappers = []any{}
	file_sso_users_roles_users_roles_proto_msgTypes[3].OneofWrappers = []any{}
	file_sso_users_roles_users_roles_proto_msgTypes[6].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_sso_users_roles_users_roles_proto_rawDesc), len(file_sso_users_roles_users_roles_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sso_users_roles_users_roles_proto_goTypes,
		DependencyIndexes: file_sso_users_roles_users_roles_proto_depIdxs,
		MessageInfos:      file_sso_users_roles_users_roles_proto_msgTypes,
	}.Build()
	File_sso_users_roles_users_roles_proto = out.File
	file_sso_users_roles_users_roles_proto_goTypes = nil
	file_sso_users_roles_users_roles_proto_depIdxs = nil
}
