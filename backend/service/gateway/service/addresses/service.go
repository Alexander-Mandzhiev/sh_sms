package addresses_service

import (
	"context"
	"log/slog"

	"backend/protos/gen/go/clients/addresses"
)

type AddressService interface {
	CreateAddress(ctx context.Context, req *addresses.CreateRequest) (*addresses.Address, error)
	GetAddress(ctx context.Context, req *addresses.GetRequest) (*addresses.Address, error)
	UpdateAddress(ctx context.Context, req *addresses.UpdateRequest) (*addresses.Address, error)
	DeleteAddress(ctx context.Context, req *addresses.DeleteRequest) (*addresses.DeleteResponse, error)
	ListByClient(ctx context.Context, req *addresses.ListRequest) (*addresses.ListResponse, error)
}

type addressService struct {
	client addresses.AddressServiceClient
	logger *slog.Logger
}

func NewAddressService(client addresses.AddressServiceClient, logger *slog.Logger) AddressService {
	return &addressService{
		client: client,
		logger: logger.With("service", "address"),
	}
}

func (s *addressService) CreateAddress(ctx context.Context, req *addresses.CreateRequest) (*addresses.Address, error) {
	s.logger.Debug("creating address", "client_id", req.ClientId, "country", req.Country, "city", req.City, "street", req.Street)
	return s.client.CreateAddress(ctx, req)
}

func (s *addressService) GetAddress(ctx context.Context, req *addresses.GetRequest) (*addresses.Address, error) {
	s.logger.Debug("getting address", "address_id", req.Id)
	return s.client.GetAddress(ctx, req)
}

func (s *addressService) UpdateAddress(ctx context.Context, req *addresses.UpdateRequest) (*addresses.Address, error) {
	s.logger.Debug("updating address", "address_id", req.Id, "updated_fields", getUpdatedAddressFields(req))
	return s.client.UpdateAddress(ctx, req)
}

func (s *addressService) DeleteAddress(ctx context.Context, req *addresses.DeleteRequest) (*addresses.DeleteResponse, error) {
	s.logger.Debug("deleting address", "address_id", req.Id)
	return s.client.DeleteAddress(ctx, req)
}

func (s *addressService) ListByClient(ctx context.Context, req *addresses.ListRequest) (*addresses.ListResponse, error) {
	s.logger.Debug("listing addresses by client", "client_id", req.ClientId, "page", req.Page, "count", req.Count)
	return s.client.ListByClient(ctx, req)
}

func getUpdatedAddressFields(req *addresses.UpdateRequest) []string {
	fields := make([]string, 0)
	if req.Country != nil {
		fields = append(fields, "country")
	}
	if req.Region != nil {
		fields = append(fields, "region")
	}
	if req.City != nil {
		fields = append(fields, "city")
	}
	if req.District != nil {
		fields = append(fields, "district")
	}
	if req.MicroDistrict != nil {
		fields = append(fields, "micro_district")
	}
	if req.Street != nil {
		fields = append(fields, "street")
	}
	if req.HouseNumber != nil {
		fields = append(fields, "house_number")
	}
	if req.Apartment != nil {
		fields = append(fields, "apartment")
	}
	if req.PostalCode != nil {
		fields = append(fields, "postal_code")
	}
	if req.Latitude != nil {
		fields = append(fields, "latitude")
	}
	if req.Longitude != nil {
		fields = append(fields, "longitude")
	}
	return fields
}
