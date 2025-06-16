package private_school_models

import "backend/protos/gen/go/private_school"

type ListTeachersResponse struct {
	Teachers   []*Teacher
	NextCursor *Cursor
}

func ListTeachersResponseFromProto(protoResp *private_school_v1.ListTeachersResponse) *ListTeachersResponse {
	teachers := make([]*Teacher, 0, len(protoResp.Teachers))
	for _, teacherProto := range protoResp.Teachers {
		teacher, err := TeacherFromProto(teacherProto)
		if err != nil {
			return nil
		}
		teachers = append(teachers, teacher)
	}

	return &ListTeachersResponse{
		Teachers:   teachers,
		NextCursor: CursorFromProto(protoResp.NextCursor),
	}
}

func ListTeachersResponseToProto(resp *ListTeachersResponse) *private_school_v1.ListTeachersResponse {
	teachersProto := make([]*private_school_v1.TeacherResponse, 0, len(resp.Teachers))
	for _, teacher := range resp.Teachers {
		teachersProto = append(teachersProto, teacher.TeacherToProto())
	}

	return &private_school_v1.ListTeachersResponse{
		Teachers:   teachersProto,
		NextCursor: CursorToProto(resp.NextCursor),
	}
}
