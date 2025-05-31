package library_models

import (
	library "backend/protos/gen/go/library"
	"strings"
)

type UpdateSubjectParams struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func UpdateSubjectParamsFromProto(req *library.UpdateSubjectRequest) (*UpdateSubjectParams, error) {
	params := &UpdateSubjectParams{
		ID:   req.GetId(),
		Name: req.GetName(),
	}

	if err := params.Validate(); err != nil {
		return nil, err
	}

	params.Sanitize()
	return params, nil
}

func (p *UpdateSubjectParams) Validate() error {
	if p.ID <= 0 {
		return ErrInvalidSubjectID
	}
	if strings.TrimSpace(p.Name) == "" {
		return ErrEmptyName
	}
	return nil
}

func (p *UpdateSubjectParams) Sanitize() {
	p.Name = strings.TrimSpace(p.Name)
	p.Name = cleanSpaces(p.Name)
}

func cleanSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
