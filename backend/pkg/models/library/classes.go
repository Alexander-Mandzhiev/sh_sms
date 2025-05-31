package library_models

import (
	library "backend/protos/gen/go/library"
)

type Class struct {
	ID    int32 `json:"id"`
	Grade int32 `json:"grade"`
}

func ProtoToClass(pb *library.Class) *Class {
	return &Class{
		ID:    pb.GetId(),
		Grade: pb.GetGrade(),
	}
}

func ClassToProto(c *Class) *library.Class {
	return &library.Class{
		Id:    c.ID,
		Grade: c.Grade,
	}
}

func ListProtoToClasses(pbList *library.ListClassesResponse) []*Class {
	classes := make([]*Class, 0, len(pbList.GetClasses()))
	for _, pbClass := range pbList.GetClasses() {
		classes = append(classes, ProtoToClass(pbClass))
	}
	return classes
}
