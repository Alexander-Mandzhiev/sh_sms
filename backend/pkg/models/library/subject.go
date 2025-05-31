package library_models

import (
	library "backend/protos/gen/go/library"
)

type Subject struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func SubjectFromProto(subj *library.Subject) *Subject {
	return &Subject{
		ID:   subj.GetId(),
		Name: subj.GetName(),
	}
}

func (s *Subject) ToProto() *library.Subject {
	return &library.Subject{
		Id:   s.ID,
		Name: s.Name,
	}
}
