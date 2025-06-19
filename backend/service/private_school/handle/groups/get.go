package groups_handle

import (
	groups_models "backend/pkg/models/groups"
	"backend/pkg/utils"
	private_school_v1 "backend/protos/gen/go/private_school"
	"context"
	"log/slog"
)

func (s *serverAPI) GetGroup(ctx context.Context, req *private_school_v1.GroupRequest) (*private_school_v1.GroupResponse, error) {
	const op = "grpc.GroupService.GetGroup"
	logger := s.logger.With(slog.String("op", op), slog.String("group_id", req.GetId()), slog.String("client_id", req.GetClientId()), slog.String("trace_id", utils.TraceIDFromContext(ctx)))
	logger.Debug("get group request received")

	groupID, err := utils.ValidateStringAndReturnUUID(req.GetId())
	if err != nil {
		logger.Warn("invalid group ID format", "error", err)
		return nil, s.convertError(groups_models.ErrInvalidGroupID)
	}

	clientID, err := utils.ValidateStringAndReturnUUID(req.GetClientId())
	if err != nil {
		logger.Warn("invalid client ID format", "error", err)
		return nil, s.convertError(groups_models.ErrInvalidClientID)
	}

	if ctx.Err() != nil {
		logger.Warn("request canceled before processing")
		return nil, s.convertError(ctx.Err())
	}

	group, err := s.service.GetGroup(ctx, groupID, clientID)
	if err != nil {
		logger.Error("failed to get group", "group_id", req.GetId(), "error", err)
		return nil, s.convertError(err)
	}

	response := group.ToProto()
	curatorInfo := "none"
	if group.CuratorID != nil {
		curatorInfo = group.CuratorID.String()
	}
	logger.Info("group retrieved successfully", "public_id", group.PublicID.String(), "name", group.Name, "curator_id", curatorInfo)
	return response, nil
}
