package groups_handle

import (
	"backend/pkg/models/groups"
	"backend/pkg/utils"
	"backend/protos/gen/go/private_school"
	"context"
	"log/slog"
)

func (s *ServerAPI) ListGroups(ctx context.Context, req *private_school_v1.ListGroupsRequest) (*private_school_v1.ListGroupsResponse, error) {
	const op = "grpc.GroupService.ListGroups"
	logger := s.Logger.With(slog.String("op", op), slog.String("client_id", req.GetClientId()), slog.Int("page_size", int(req.GetPageSize())), slog.String("trace_id", utils.TraceIDFromContext(ctx)))
	if req.Cursor != nil {
		logger = logger.With(slog.Int64("cursor", req.GetCursor()))
	}
	if req.NameFilter != nil {
		logger = logger.With(slog.String("name_filter", req.GetNameFilter()))
	}
	logger.Debug("list groups request received")

	if ctx.Err() != nil {
		logger.Warn("request canceled before processing")
		return nil, s.convertError(ctx.Err())
	}

	listParams, err := groups_models.ListGroupsParamsFromProto(req)
	if err != nil {
		logger.Warn("invalid request parameters", "error", err)
		return nil, s.convertError(err)
	}

	response, err := s.Service.ListGroups(ctx, listParams)
	if err != nil {
		logger.Error("failed to list groups", "error", err, "client_id", req.GetClientId())
		return nil, s.convertError(err)
	}

	protoResponse := response.ToProto()
	logger.Info("groups listed successfully", "count", len(protoResponse.Groups), "next_cursor", protoResponse.NextCursor, "has_next", protoResponse.NextCursor > 0)
	return protoResponse, nil
}
