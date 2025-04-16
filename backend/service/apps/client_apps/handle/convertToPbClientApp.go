package handle

import (
	pb "backend/protos/gen/go/apps/clients_apps"
	"backend/service/apps/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *serverAPI) convertToPbClientApp(app *models.ClientApp) *pb.ClientApp {
	return &pb.ClientApp{
		ClientId:  app.ClientID,
		AppId:     int32(app.AppID),
		IsActive:  app.IsActive,
		CreatedAt: timestamppb.New(app.CreatedAt),
		UpdatedAt: timestamppb.New(app.UpdatedAt),
	}
}

func (s *serverAPI) convertToPbClientApps(apps []*models.ClientApp) []*pb.ClientApp {
	result := make([]*pb.ClientApp, 0, len(apps))
	for _, app := range apps {
		pbApp := s.convertToPbClientApp(app)
		result = append(result, pbApp)
	}
	return result
}
