package library_models

import (
	library "backend/protos/gen/go/library"
)

type ListSubjectsResponse struct {
	Subjects []*Subject `json:"subjects"`
}

func ListSubjectsResponseFromProto(resp *library.ListSubjectsResponse) *ListSubjectsResponse {
	subjects := make([]*Subject, 0, len(resp.GetSubjects()))
	for _, s := range resp.GetSubjects() {
		subjects = append(subjects, SubjectFromProto(s))
	}
	return &ListSubjectsResponse{Subjects: subjects}
}

func (r *ListSubjectsResponse) ToProto() *library.ListSubjectsResponse {
	subjects := make([]*library.Subject, 0, len(r.Subjects))
	for _, s := range r.Subjects {
		subjects = append(subjects, s.ToProto())
	}
	return &library.ListSubjectsResponse{Subjects: subjects}
}
