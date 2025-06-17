package students_models

import "backend/protos/gen/go/private_school"

type ListStudentsResponse struct {
	Students   []*Student
	NextCursor string
}

func ListStudentsResponseFromProto(protoResp *private_school_v1.ListStudentsResponse) (*ListStudentsResponse, error) {
	resp := &ListStudentsResponse{
		NextCursor: "",
	}

	if protoResp.NextCursor != nil {
		resp.NextCursor = *protoResp.NextCursor
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
	protoStudents := make([]*private_school_v1.StudentResponse, 0, len(resp.Students))
	for _, student := range resp.Students {
		protoStudents = append(protoStudents, student.StudentToProto())
	}

	protoResp := &private_school_v1.ListStudentsResponse{Students: protoStudents}

	if resp.NextCursor != "" {
		protoResp.NextCursor = &resp.NextCursor
	}

	return protoResp
}
