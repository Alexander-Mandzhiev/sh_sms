package students_models

import (
	"backend/protos/gen/go/private_school"
)

type ListStudentsResponse struct {
	Students   []*Student
	NextCursor *Cursor
}

func ListStudentsResponseFromProto(protoResp *private_school_v1.ListStudentsResponse) (*ListStudentsResponse, error) {
	resp := &ListStudentsResponse{
		Students: make([]*Student, 0, len(protoResp.Students)),
	}

	if protoResp.NextCursor != nil {
		cursor, err := CursorFromProto(protoResp.NextCursor)
		if err != nil {
			return nil, err
		}
		resp.NextCursor = cursor
	}

	for _, protoStudent := range protoResp.Students {
		student, err := StudentFromProto(protoStudent)
		if err != nil {
			return nil, err
		}
		resp.Students = append(resp.Students, student)
	}

	return resp, nil
}

func (resp *ListStudentsResponse) ToProto() *private_school_v1.ListStudentsResponse {
	protoStudents := make([]*private_school_v1.StudentResponse, len(resp.Students))
	for i, student := range resp.Students {
		protoStudents[i] = student.StudentToProto()
	}

	protoResp := &private_school_v1.ListStudentsResponse{
		Students: protoStudents,
	}

	if resp.NextCursor != nil {
		protoResp.NextCursor = resp.NextCursor.ToProto()
	}

	return protoResp
}
