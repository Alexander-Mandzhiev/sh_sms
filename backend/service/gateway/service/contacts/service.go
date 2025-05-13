package contacts_service

import (
	"context"
	"log/slog"

	"backend/protos/gen/go/clients/contacts"
)

type ContactService interface {
	CreateContact(ctx context.Context, req *contacts.CreateRequest) (*contacts.Contact, error)
	GetContact(ctx context.Context, req *contacts.GetRequest) (*contacts.Contact, error)
	UpdateContact(ctx context.Context, req *contacts.UpdateRequest) (*contacts.Contact, error)
	DeleteContact(ctx context.Context, req *contacts.DeleteRequest) (*contacts.DeleteResponse, error)
	ListContacts(ctx context.Context, req *contacts.ListRequest) (*contacts.ListResponse, error)
}

type contactService struct {
	client contacts.ContactServiceClient
	logger *slog.Logger
}

func NewContactService(client contacts.ContactServiceClient, logger *slog.Logger) ContactService {
	return &contactService{
		client: client,
		logger: logger.With("service", "contact"),
	}
}

func (s *contactService) CreateContact(ctx context.Context, req *contacts.CreateRequest) (*contacts.Contact, error) {
	s.logger.Debug("creating contact", "client_id", req.ClientId, "full_name", req.FullName, "email", req.GetEmail(), "phone", req.GetPhone())
	return s.client.CreateContact(ctx, req)
}

func (s *contactService) GetContact(ctx context.Context, req *contacts.GetRequest) (*contacts.Contact, error) {
	s.logger.Debug("getting contact", "contact_id", req.Id)
	return s.client.GetContact(ctx, req)
}

func (s *contactService) UpdateContact(ctx context.Context, req *contacts.UpdateRequest) (*contacts.Contact, error) {
	s.logger.Debug("updating contact", "contact_id", req.Id, "updated_fields", getUpdatedContactFields(req))
	return s.client.UpdateContact(ctx, req)
}

func (s *contactService) DeleteContact(ctx context.Context, req *contacts.DeleteRequest) (*contacts.DeleteResponse, error) {
	s.logger.Debug("deleting contact", "contact_id", req.Id)
	return s.client.DeleteContact(ctx, req)
}

func (s *contactService) ListContacts(ctx context.Context, req *contacts.ListRequest) (*contacts.ListResponse, error) {
	s.logger.Debug("listing contacts", "client_id", req.ClientId, "page", req.Page, "page_size", req.PageSize, "search", req.GetSearch(), "active_only", req.GetActiveOnly())
	return s.client.ListContacts(ctx, req)
}

func getUpdatedContactFields(req *contacts.UpdateRequest) []string {
	fields := make([]string, 0)
	if req.FullName != nil {
		fields = append(fields, "full_name")
	}
	if req.Position != nil {
		fields = append(fields, "position")
	}
	if req.Email != nil {
		fields = append(fields, "email")
	}
	if req.Phone != nil {
		fields = append(fields, "phone")
	}
	if req.IsActive != nil {
		fields = append(fields, "is_active")
	}
	return fields
}
