package library_models

import (
	library "backend/protos/gen/go/library"
	"fmt"
	"github.com/google/uuid"
	"regexp"
	"strings"
)

const MaxDescriptionLength = 2000

type UpdateBookParams struct {
	ID          int64   `json:"id"`
	ClientID    string  `json:"client_id"`
	Title       *string `json:"title,omitempty"`
	Author      *string `json:"author,omitempty"`
	Description *string `json:"description,omitempty"`
	SubjectID   *int32  `json:"subject_id,omitempty"`
	ClassID     *int32  `json:"class_id,omitempty"`
}

func UpdateParamsFromProto(req *library.UpdateBookRequest) (*UpdateBookParams, error) {
	params := &UpdateBookParams{
		ID:          req.GetId(),
		ClientID:    req.GetClientId(),
		Title:       req.Title,
		Author:      req.Author,
		Description: req.Description,
		SubjectID:   req.SubjectId,
		ClassID:     req.ClassId,
	}

	if err := params.Validate(); err != nil {
		return nil, err
	}

	return params, nil
}

func (p *UpdateBookParams) Validate() error {
	if p.ID <= 0 {
		return ErrInvalidID
	}

	if _, err := uuid.Parse(p.ClientID); err != nil {
		return fmt.Errorf("%w: %s", ErrBookInvalidClientID, p.ClientID)
	}

	if p.Title != nil {
		if strings.TrimSpace(*p.Title) == "" {
			return fmt.Errorf("%w: %s", ErrBookInvalidTitle, p.Title)
		}
	}

	if p.Author != nil {
		if strings.TrimSpace(*p.Author) == "" {
			return fmt.Errorf("%w: %s", ErrBookInvalidAuthor, p.Author)
		}
	}

	if p.Description != nil {
		if len(*p.Description) > MaxDescriptionLength {
			return fmt.Errorf("%w: %d > %d",
				ErrBookDescriptionLong,
				len(*p.Description),
				MaxDescriptionLength)
		}
	}

	if p.SubjectID != nil && *p.SubjectID <= 0 {
		return fmt.Errorf("%w: %d", ErrBookInvalidSubjectID, *p.SubjectID)
	}

	if p.ClassID != nil && (*p.ClassID < 1 || *p.ClassID > 11) {
		return fmt.Errorf("%w: %d", ErrBookInvalidClassID, *p.ClassID)
	}

	return nil
}

func (p *UpdateBookParams) Sanitize() {
	if p.Title != nil {
		*p.Title = strings.TrimSpace(*p.Title)
		*p.Title = regexp.MustCompile(`\s+`).ReplaceAllString(*p.Title, " ")
	}

	if p.Author != nil {
		*p.Author = strings.TrimSpace(*p.Author)
		*p.Author = regexp.MustCompile(`\s+`).ReplaceAllString(*p.Author, " ")
		*p.Author = normalizeAuthorName(*p.Author)
	}

	if p.Description != nil {
		*p.Description = strings.TrimSpace(*p.Description)
	}
}

func normalizeAuthorName(name string) string {
	parts := strings.Fields(name)
	for i, part := range parts {
		if len(part) > 0 {
			parts[i] = strings.ToUpper(part[:1]) + strings.ToLower(part[1:])
		}
	}
	return strings.Join(parts, " ")
}

func (p *UpdateBookParams) ToUpdateRequestProto() *library.UpdateBookRequest {
	return &library.UpdateBookRequest{
		Id:          p.ID,
		ClientId:    p.ClientID,
		Title:       p.Title,
		Author:      p.Author,
		Description: p.Description,
		SubjectId:   p.SubjectID,
		ClassId:     p.ClassID,
	}
}
