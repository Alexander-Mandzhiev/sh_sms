package groups_handle

import (
	"backend/pkg/models/groups"
	"backend/pkg/utils"
	"backend/protos/gen/go/private_school"
	"context"
	"log/slog"
)

func (s *ServerAPI) CreateGroup(ctx context.Context, req *private_school_v1.CreateGroupRequest) (*private_school_v1.GroupResponse, error) {
	const op = "grpc.GroupService.CreateGroup"
	logger := s.Logger.With(slog.String("op", op), slog.String("client_id", req.GetClientId()), slog.String("trace_id", utils.TraceIDFromContext(ctx)))
	if req.CuratorId != nil && *req.CuratorId != "" {
		logger = logger.With(slog.String("curator_id", *req.CuratorId))
	}
	logger.Debug("create group request received", "name", req.GetName(), "client_id", req.GetClientId())

	if ctx.Err() != nil {
		logger.Warn("request canceled before processing")
		return nil, s.convertError(ctx.Err())
	}

	createGroup, err := groups_models.CreateGroupFromProto(req)
	if err != nil {
		logger.Warn("invalid request parameters", "error", err)
		return nil, s.convertError(err)
	}

	group, err := s.Service.CreateGroup(ctx, createGroup)
	if err != nil {
		logger.Error("group creation failed", "name", req.GetName(), "error", err)
		return nil, s.convertError(err)
	}

	response := group.ToProto()
	logger.Info("group created successfully", "public_id", group.PublicID.String(), "internal_id", group.InternalID, "name", group.Name, "curator_id", group.CuratorID)
	return response, nil
}
