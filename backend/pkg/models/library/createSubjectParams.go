package library_models

import (
	library "backend/protos/gen/go/library"
	"strings"
)

type CreateSubjectParams struct {
	Name string `json:"name"`
}

func CreateSubjectParamsFromProto(req *library.CreateSubjectRequest) (*CreateSubjectParams, error) {
	params := &CreateSubjectParams{
		Name: req.GetName(),
	}

	if err := params.Validate(); err != nil {
		return nil, err
	}

	params.Sanitize()
	return params, nil
}

func (p *CreateSubjectParams) Validate() error {
	if strings.TrimSpace(p.Name) == "" {
		return ErrEmptyName
	}
	return nil
}

func (p *CreateSubjectParams) Sanitize() {
	p.Name = strings.TrimSpace(p.Name)
	p.Name = cleanSpaces(p.Name)
}
