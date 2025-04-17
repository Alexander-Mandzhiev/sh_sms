package handle

import (
	pb "backend/protos/gen/go/apps/secrets"
	"backend/service/apps/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertRotationHistoryToPB(h *models.RotationHistory) *pb.RotationHistory {
	pbHistory := &pb.RotationHistory{
		ClientId:   h.ClientID,
		AppId:      int32(h.AppID),
		SecretType: h.SecretType,
		OldSecret:  h.OldSecret,
		NewSecret:  h.NewSecret,
		RotatedAt:  timestamppb.New(h.RotatedAt),
	}

	if h.RotatedBy != "" {
		pbHistory.RotatedBy = &h.RotatedBy
	}

	return pbHistory
}

func convertRotationsToPB(histories []*models.RotationHistory) []*pb.RotationHistory {
	pbHistories := make([]*pb.RotationHistory, 0, len(histories))
	for _, h := range histories {
		pbHistories = append(pbHistories, convertRotationHistoryToPB(h))
	}
	return pbHistories
}
