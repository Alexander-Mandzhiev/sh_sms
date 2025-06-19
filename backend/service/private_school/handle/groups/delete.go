package groups_handle

import (
	"backend/pkg/models/groups"
	"backend/pkg/utils"
	"backend/protos/gen/go/private_school"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"log/slog"
)

func (s *ServerAPI) DeleteGroup(ctx context.Context, req *private_school_v1.GroupRequest) (*emptypb.Empty, error) {
	const op = "grpc.GroupService.DeleteGroup"
	logger := s.Logger.With(slog.String("op", op), slog.String("group_id", req.GetId()), slog.String("client_id", req.GetClientId()), slog.String("trace_id", utils.TraceIDFromContext(ctx)))
	logger.Debug("delete group request received")

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

	err = s.Service.DeleteGroup(ctx, groupID, clientID)
	if err != nil {
		logger.Error("failed to delete group", "group_id", req.GetId(), "error", err)
		return nil, s.convertError(err)
	}

	logger.Info("group deleted successfully", "public_id", groupID.String(), "client_id", clientID.String())
	return &emptypb.Empty{}, nil
}
