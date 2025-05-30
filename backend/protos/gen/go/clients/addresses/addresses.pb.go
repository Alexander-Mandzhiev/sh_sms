// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.1
// source: clients/addresses/addresses.proto

package addresses

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type Address struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ClientId      string                 `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Country       string                 `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	Region        *string                `protobuf:"bytes,4,opt,name=region,proto3,oneof" json:"region,omitempty"`
	City          string                 `protobuf:"bytes,5,opt,name=city,proto3" json:"city,omitempty"`
	District      *string                `protobuf:"bytes,6,opt,name=district,proto3,oneof" json:"district,omitempty"`
	MicroDistrict *string                `protobuf:"bytes,7,opt,name=micro_district,json=microDistrict,proto3,oneof" json:"micro_district,omitempty"`
	Street        string                 `protobuf:"bytes,8,opt,name=street,proto3" json:"street,omitempty"`
	HouseNumber   string                 `protobuf:"bytes,9,opt,name=house_number,json=houseNumber,proto3" json:"house_number,omitempty"`
	Apartment     *string                `protobuf:"bytes,10,opt,name=apartment,proto3,oneof" json:"apartment,omitempty"`
	PostalCode    *string                `protobuf:"bytes,11,opt,name=postal_code,json=postalCode,proto3,oneof" json:"postal_code,omitempty"`
	Latitude      *float64               `protobuf:"fixed64,12,opt,name=latitude,proto3,oneof" json:"latitude,omitempty"`
	Longitude     *float64               `protobuf:"fixed64,13,opt,name=longitude,proto3,oneof" json:"longitude,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Address) Reset() {
	*x = Address{}
	mi := &file_clients_addresses_addresses_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_clients_addresses_addresses_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_clients_addresses_addresses_proto_rawDescGZIP(), []int{0}
}

func (x *Address) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Address) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *Address) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Address) GetRegion() string {
	if x != nil && x.Region != nil {
		return *x.Region
	}
	return ""
}

func (x *Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Address) GetDistrict() string {
	if x != nil && x.District != nil {
		return *x.District
	}
	return ""
}

func (x *Address) GetMicroDistrict() string {
	if x != nil && x.MicroDistrict != nil {
		return *x.MicroDistrict
	}
	return ""
}

func (x *Address) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *Address) GetHouseNumber() string {
	if x != nil {
		return x.HouseNumber
	}
	return ""
}

func (x *Address) GetApartment() string {
	if x != nil && x.Apartment != nil {
		return *x.Apartment
	}
	return ""
}

func (x *Address) GetPostalCode() string {
	if x != nil && x.PostalCode != nil {
		return *x.PostalCode
	}
	return ""
}

func (x *Address) GetLatitude() float64 {
	if x != nil && x.Latitude != nil {
		return *x.Latitude
	}
	return 0
}

func (x *Address) GetLongitude() float64 {
	if x != nil && x.Longitude != nil {
		return *x.Longitude
	}
	return 0
}

type CreateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ClientId      string                 `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Country       string                 `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	City          string                 `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	Region        *string                `protobuf:"bytes,4,opt,name=region,proto3,oneof" json:"region,omitempty"`
	District      *string                `protobuf:"bytes,5,opt,name=district,proto3,oneof" json:"district,omitempty"`
	MicroDistrict *string                `protobuf:"bytes,6,opt,name=micro_district,json=microDistrict,proto3,oneof" json:"micro_district,omitempty"`
	Street        string                 `protobuf:"bytes,7,opt,name=street,proto3" json:"street,omitempty"`
	HouseNumber   string                 `protobuf:"bytes,8,opt,name=house_number,json=houseNumber,proto3" json:"house_number,omitempty"`
	Apartment     *string                `protobuf:"bytes,9,opt,name=apartment,proto3,oneof" json:"apartment,omitempty"`
	PostalCode    *string                `protobuf:"bytes,10,opt,name=postal_code,json=postalCode,proto3,oneof" json:"postal_code,omitempty"`
	Latitude      *float64               `protobuf:"fixed64,11,opt,name=latitude,proto3,oneof" json:"latitude,omitempty"`
	Longitude     *float64               `protobuf:"fixed64,12,opt,name=longitude,proto3,oneof" json:"longitude,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	mi := &file_clients_addresses_addresses_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_clients_addresses_addresses_proto_msgTypes[1]
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
	return file_clients_addresses_addresses_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *CreateRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *CreateRequest) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *CreateRequest) GetRegion() string {
	if x != nil && x.Region != nil {
		return *x.Region
	}
	return ""
}

func (x *CreateRequest) GetDistrict() string {
	if x != nil && x.District != nil {
		return *x.District
	}
	return ""
}

func (x *CreateRequest) GetMicroDistrict() string {
	if x != nil && x.MicroDistrict != nil {
		return *x.MicroDistrict
	}
	return ""
}

func (x *CreateRequest) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *CreateRequest) GetHouseNumber() string {
	if x != nil {
		return x.HouseNumber
	}
	return ""
}

func (x *CreateRequest) GetApartment() string {
	if x != nil && x.Apartment != nil {
		return *x.Apartment
	}
	return ""
}

func (x *CreateRequest) GetPostalCode() string {
	if x != nil && x.PostalCode != nil {
		return *x.PostalCode
	}
	return ""
}

func (x *CreateRequest) GetLatitude() float64 {
	if x != nil && x.Latitude != nil {
		return *x.Latitude
	}
	return 0
}

func (x *CreateRequest) GetLongitude() float64 {
	if x != nil && x.Longitude != nil {
		return *x.Longitude
	}
	return 0
}

type UpdateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ClientId      string                 `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Country       *string                `protobuf:"bytes,3,opt,name=country,proto3,oneof" json:"country,omitempty"`
	Region        *string                `protobuf:"bytes,4,opt,name=region,proto3,oneof" json:"region,omitempty"`
	City          *string                `protobuf:"bytes,5,opt,name=city,proto3,oneof" json:"city,omitempty"`
	District      *string                `protobuf:"bytes,6,opt,name=district,proto3,oneof" json:"district,omitempty"`
	MicroDistrict *string                `protobuf:"bytes,7,opt,name=micro_district,json=microDistrict,proto3,oneof" json:"micro_district,omitempty"`
	Street        *string                `protobuf:"bytes,8,opt,name=street,proto3,oneof" json:"street,omitempty"`
	HouseNumber   *string                `protobuf:"bytes,9,opt,name=house_number,json=houseNumber,proto3,oneof" json:"house_number,omitempty"`
	Apartment     *string                `protobuf:"bytes,10,opt,name=apartment,proto3,oneof" json:"apartment,omitempty"`
	PostalCode    *string                `protobuf:"bytes,11,opt,name=postal_code,json=postalCode,proto3,oneof" json:"postal_code,omitempty"`
	Latitude      *float64               `protobuf:"fixed64,12,opt,name=latitude,proto3,oneof" json:"latitude,omitempty"`
	Longitude     *float64               `protobuf:"fixed64,13,opt,name=longitude,proto3,oneof" json:"longitude,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	mi := &file_clients_addresses_addresses_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_clients_addresses_addresses_proto_msgTypes[2]
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
	return file_clients_addresses_addresses_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *UpdateRequest) GetCountry() string {
	if x != nil && x.Country != nil {
		return *x.Country
	}
	return ""
}

func (x *UpdateRequest) GetRegion() string {
	if x != nil && x.Region != nil {
		return *x.Region
	}
	return ""
}

func (x *UpdateRequest) GetCity() string {
	if x != nil && x.City != nil {
		return *x.City
	}
	return ""
}

func (x *UpdateRequest) GetDistrict() string {
	if x != nil && x.District != nil {
		return *x.District
	}
	return ""
}

func (x *UpdateRequest) GetMicroDistrict() string {
	if x != nil && x.MicroDistrict != nil {
		return *x.MicroDistrict
	}
	return ""
}

func (x *UpdateRequest) GetStreet() string {
	if x != nil && x.Street != nil {
		return *x.Street
	}
	return ""
}

func (x *UpdateRequest) GetHouseNumber() string {
	if x != nil && x.HouseNumber != nil {
		return *x.HouseNumber
	}
	return ""
}

func (x *UpdateRequest) GetApartment() string {
	if x != nil && x.Apartment != nil {
		return *x.Apartment
	}
	return ""
}

func (x *UpdateRequest) GetPostalCode() string {
	if x != nil && x.PostalCode != nil {
		return *x.PostalCode
	}
	return ""
}

func (x *UpdateRequest) GetLatitude() float64 {
	if x != nil && x.Latitude != nil {
		return *x.Latitude
	}
	return 0
}

func (x *UpdateRequest) GetLongitude() float64 {
	if x != nil && x.Longitude != nil {
		return *x.Longitude
	}
	return 0
}

type GetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	mi := &file_clients_addresses_addresses_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_clients_addresses_addresses_proto_msgTypes[3]
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
	return file_clients_addresses_addresses_proto_rawDescGZIP(), []int{3}
}

func (x *GetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	mi := &file_clients_addresses_addresses_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_clients_addresses_addresses_proto_msgTypes[4]
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
	return file_clients_addresses_addresses_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ClientId      string                 `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Page          int32                  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Count         int32                  `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	mi := &file_clients_addresses_addresses_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_clients_addresses_addresses_proto_msgTypes[5]
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
	return file_clients_addresses_addresses_proto_rawDescGZIP(), []int{5}
}

func (x *ListRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
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

type ListResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Addresses     []*Address             `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty"`
	TotalCount    int32                  `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	mi := &file_clients_addresses_addresses_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_clients_addresses_addresses_proto_msgTypes[6]
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
	return file_clients_addresses_addresses_proto_rawDescGZIP(), []int{6}
}

func (x *ListResponse) GetAddresses() []*Address {
	if x != nil {
		return x.Addresses
	}
	return nil
}

func (x *ListResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
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
	mi := &file_clients_addresses_addresses_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_clients_addresses_addresses_proto_msgTypes[7]
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
	return file_clients_addresses_addresses_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_clients_addresses_addresses_proto protoreflect.FileDescriptor

var file_clients_addresses_addresses_proto_rawDesc = string([]byte{
	0x0a, 0x21, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x65, 0x73, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x11, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x22, 0xfa, 0x03, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x06, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x1f, 0x0a, 0x08, 0x64, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x08,
	0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a, 0x0e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0d, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x44, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x63, 0x74, 0x88, 0x01, 0x01, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x12,
	0x21, 0x0a, 0x0c, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x21, 0x0a, 0x09, 0x61, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x09, 0x61, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x0a, 0x70, 0x6f,
	0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x6c,
	0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x01, 0x48, 0x05, 0x52,
	0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09,
	0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x01, 0x48,
	0x06, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x88, 0x01, 0x01, 0x42,
	0x09, 0x0a, 0x07, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x64,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x61,
	0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x70, 0x6f, 0x73,
	0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6c, 0x61, 0x74,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x22, 0xf0, 0x03, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79,
	0x12, 0x1b, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a,
	0x08, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x01, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x88, 0x01, 0x01, 0x12, 0x2a,
	0x0a, 0x0e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0d, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x44,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x88, 0x01, 0x01, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x72, 0x65, 0x65, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65,
	0x65, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x09, 0x61, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x09, 0x61, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x74,
	0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52,
	0x0a, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f,
	0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01,
	0x48, 0x05, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x21, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x01, 0x48, 0x06, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x88,
	0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x42, 0x0b, 0x0a,
	0x09, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x42, 0x0c, 0x0a,
	0x0a, 0x5f, 0x61, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x0e, 0x0a, 0x0c, 0x5f,
	0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f,
	0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6c, 0x6f, 0x6e,
	0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0xc5, 0x04, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x88, 0x01,
	0x01, 0x12, 0x17, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x02, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x64, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x08,
	0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a, 0x0e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x0d, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x44, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x63, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65,
	0x74, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x06, 0x52, 0x0b, 0x68, 0x6f,
	0x75, 0x73, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09,
	0x61, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x07, 0x52, 0x09, 0x61, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12,
	0x24, 0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x08, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f,
	0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x01, 0x48, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x01, 0x48, 0x0a, 0x52, 0x09, 0x6c, 0x6f, 0x6e,
	0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x42, 0x07, 0x0a, 0x05, 0x5f, 0x63, 0x69, 0x74, 0x79, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x64, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x5f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74,
	0x72, 0x65, 0x65, 0x74, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x61, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0x1c,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1f, 0x0a, 0x0d,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x54, 0x0a,
	0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x69, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x1f, 0x0a,
	0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2a,
	0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x9e, 0x03, 0x0a, 0x0e, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a,
	0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x20,
	0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x65, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x47, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x2e, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x2e, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x4d, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x20, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x2e, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x54, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x20, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0c, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x79, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x2e, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x16, 0x5a, 0x14, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x3b, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_clients_addresses_addresses_proto_rawDescOnce sync.Once
	file_clients_addresses_addresses_proto_rawDescData []byte
)

func file_clients_addresses_addresses_proto_rawDescGZIP() []byte {
	file_clients_addresses_addresses_proto_rawDescOnce.Do(func() {
		file_clients_addresses_addresses_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_clients_addresses_addresses_proto_rawDesc), len(file_clients_addresses_addresses_proto_rawDesc)))
	})
	return file_clients_addresses_addresses_proto_rawDescData
}

var file_clients_addresses_addresses_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_clients_addresses_addresses_proto_goTypes = []any{
	(*Address)(nil),        // 0: clients.addresses.Address
	(*CreateRequest)(nil),  // 1: clients.addresses.CreateRequest
	(*UpdateRequest)(nil),  // 2: clients.addresses.UpdateRequest
	(*GetRequest)(nil),     // 3: clients.addresses.GetRequest
	(*DeleteRequest)(nil),  // 4: clients.addresses.DeleteRequest
	(*ListRequest)(nil),    // 5: clients.addresses.ListRequest
	(*ListResponse)(nil),   // 6: clients.addresses.ListResponse
	(*DeleteResponse)(nil), // 7: clients.addresses.DeleteResponse
}
var file_clients_addresses_addresses_proto_depIdxs = []int32{
	0, // 0: clients.addresses.ListResponse.addresses:type_name -> clients.addresses.Address
	1, // 1: clients.addresses.AddressService.CreateAddress:input_type -> clients.addresses.CreateRequest
	3, // 2: clients.addresses.AddressService.GetAddress:input_type -> clients.addresses.GetRequest
	2, // 3: clients.addresses.AddressService.UpdateAddress:input_type -> clients.addresses.UpdateRequest
	4, // 4: clients.addresses.AddressService.DeleteAddress:input_type -> clients.addresses.DeleteRequest
	5, // 5: clients.addresses.AddressService.ListByClient:input_type -> clients.addresses.ListRequest
	0, // 6: clients.addresses.AddressService.CreateAddress:output_type -> clients.addresses.Address
	0, // 7: clients.addresses.AddressService.GetAddress:output_type -> clients.addresses.Address
	0, // 8: clients.addresses.AddressService.UpdateAddress:output_type -> clients.addresses.Address
	7, // 9: clients.addresses.AddressService.DeleteAddress:output_type -> clients.addresses.DeleteResponse
	6, // 10: clients.addresses.AddressService.ListByClient:output_type -> clients.addresses.ListResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_clients_addresses_addresses_proto_init() }
func file_clients_addresses_addresses_proto_init() {
	if File_clients_addresses_addresses_proto != nil {
		return
	}
	file_clients_addresses_addresses_proto_msgTypes[0].OneofWrappers = []any{}
	file_clients_addresses_addresses_proto_msgTypes[1].OneofWrappers = []any{}
	file_clients_addresses_addresses_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_clients_addresses_addresses_proto_rawDesc), len(file_clients_addresses_addresses_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_clients_addresses_addresses_proto_goTypes,
		DependencyIndexes: file_clients_addresses_addresses_proto_depIdxs,
		MessageInfos:      file_clients_addresses_addresses_proto_msgTypes,
	}.Build()
	File_clients_addresses_addresses_proto = out.File
	file_clients_addresses_addresses_proto_goTypes = nil
	file_clients_addresses_addresses_proto_depIdxs = nil
}
