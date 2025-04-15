package handle

import (
	pb "backend/protos/gen/go/apps/app_manager"
	"backend/service/apps/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *serverAPI) convertAppToProto(app *models.App) *pb.App {
	pbApp := &pb.App{
		Id:        int32(app.ID),
		Code:      app.Code,
		Name:      app.Name,
		IsActive:  app.IsActive,
		CreatedAt: timestamppb.New(app.CreatedAt),
		UpdatedAt: timestamppb.New(app.UpdatedAt),
	}

	if app.Description != "" {
		pbApp.Description = app.Description
	}

	return pbApp
}
