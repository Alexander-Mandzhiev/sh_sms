package library_models

import (
	library "backend/protos/gen/go/library"
	"fmt"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

type CreateBookParams struct {
	ClientID    uuid.UUID `json:"client_id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Description string    `json:"description"`
	SubjectID   int32     `json:"subject_id"`
	ClassID     int32     `json:"class_id"`
}

func CreateParamsFromProto(req *library.CreateBookRequest) (*CreateBookParams, error) {
	clientID, err := uuid.Parse(req.ClientId)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrInvalidClientID, req.ClientId)
	}

	params := &CreateBookParams{
		ClientID:    clientID,
		Title:       req.GetTitle(),
		Author:      req.GetAuthor(),
		Description: req.GetDescription(),
		SubjectID:   req.GetSubjectId(),
		ClassID:     req.GetClassId(),
	}

	if err = params.Validate(); err != nil {
		return nil, err
	}

	return params, nil
}

func (p *CreateBookParams) Validate() error {
	if strings.TrimSpace(p.Title) == "" {
		return ErrEmptyTitle
	}

	if strings.TrimSpace(p.Author) == "" {
		return ErrEmptyAuthor
	}

	if p.SubjectID <= 0 {
		return fmt.Errorf("%w: %d", ErrInvalidSubject, p.SubjectID)
	}

	if p.ClassID < 1 || p.ClassID > 11 {
		return fmt.Errorf("%w: %d", ErrInvalidClass, p.ClassID)
	}

	return nil
}

func (p *CreateBookParams) Sanitize() {
	p.Title = strings.TrimSpace(p.Title)
	p.Author = strings.TrimSpace(p.Author)
	p.Description = strings.TrimSpace(p.Description)

	p.Title = regexp.MustCompile(`\s+`).ReplaceAllString(p.Title, " ")
	p.Author = regexp.MustCompile(`\s+`).ReplaceAllString(p.Author, " ")
}

func (p *CreateBookParams) ToCreateRequestProto() *library.CreateBookRequest {
	return &library.CreateBookRequest{
		ClientId:    p.ClientID.String(),
		Title:       p.Title,
		Author:      p.Author,
		Description: &p.Description,
		SubjectId:   p.SubjectID,
		ClassId:     p.ClassID,
	}
}
