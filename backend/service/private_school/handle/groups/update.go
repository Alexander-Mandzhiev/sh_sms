package groups_handle

import (
	"backend/pkg/models/groups"
	"backend/pkg/utils"
	"backend/protos/gen/go/private_school"
	"context"
	"log/slog"
)

func (s *ServerAPI) UpdateGroup(ctx context.Context, req *private_school_v1.UpdateGroupRequest) (*private_school_v1.GroupResponse, error) {
	const op = "grpc.GroupService.UpdateGroup"
	logger := s.Logger.With(slog.String("op", op), slog.String("group_id", req.GetId()), slog.String("client_id", req.GetClientId()), slog.String("trace_id", utils.TraceIDFromContext(ctx)))
	logger.Debug("update group request received", "name", req.GetName(), "curator_id", req.GetCuratorId())

	if ctx.Err() != nil {
		logger.Warn("request canceled before processing")
		return nil, s.convertError(ctx.Err())
	}

	updateGroup, err := groups_models.UpdateGroupFromProto(req)
	if err != nil {
		logger.Warn("invalid request parameters", "error", err)
		return nil, s.convertError(err)
	}

	group, err := s.Service.UpdateGroup(ctx, updateGroup)
	if err != nil {
		logger.Error("group update failed", "group_id", req.GetId(), "error", err)
		return nil, s.convertError(err)
	}

	response := group.ToProto()

	curatorInfo := "none"
	if group.CuratorID != nil {
		curatorInfo = group.CuratorID.String()
	}

	logger.Info("group updated successfully", "public_id", group.PublicID.String(), "name", group.Name, "curator_id", curatorInfo)
	return response, nil
}
