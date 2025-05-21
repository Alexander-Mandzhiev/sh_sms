package handle

import (
	"backend/pkg/utils"
	"backend/protos/gen/go/sso/users"
	"backend/service/sso/models"
	"context"
	"github.com/google/uuid"
	"log/slog"
)

func (s *serverAPI) BatchGetUsers(ctx context.Context, req *users.BatchGetRequest) (*users.BatchGetResponse, error) {
	const op = "grpc.user.BatchGetUsers"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", req.GetClientId()), slog.Int("user_count", len(req.GetUserIds())))
	logger.Debug("processing batch users request")

	clientID, err := utils.ValidateAndReturnUUID(req.GetClientId())
	if err != nil {
		logger.Warn("invalid client_id", slog.Any("error", err))
		return nil, s.convertError(ErrInvalidArgument)
	}

	var validIDs []uuid.UUID
	var invalidIDs []string
	for _, idStr := range req.GetUserIds() {
		var id uuid.UUID
		id, err = utils.ValidateAndReturnUUID(idStr)
		if err != nil {
			invalidIDs = append(invalidIDs, idStr)
			continue
		}
		validIDs = append(validIDs, id)
	}

	if len(invalidIDs) > 0 {
		logger.Warn("invalid user_ids detected", slog.Any("invalid_ids", invalidIDs), slog.Int("total_requested", len(req.GetUserIds())))
		return nil, s.convertError(ErrInvalidArgument)
	}

	if len(validIDs) == 0 {
		logger.Warn("empty valid user_ids")
		return nil, s.convertError(ErrInvalidArgument)
	}

	if len(validIDs) > 1000 {
		logger.Warn("too many user_ids requested", slog.Int("requested", len(validIDs)), slog.Int("max_allowed", 1000))
		return nil, s.convertError(ErrInvalidArgument)
	}

	us, missingIDs, err := s.service.BatchGetUsers(ctx, clientID, validIDs, req.GetIncludeInactive())
	if err != nil {
		logger.Error("failed to batch get users", slog.Any("error", err), slog.Int("attempt", 1))
		return nil, s.convertError(err)
	}

	pbUsers := make([]*users.User, 0, len(us))
	for _, u := range us {
		pbUsers = append(pbUsers, models.ConvertUserToProto(u))
	}

	pbMissing := make([]string, 0, len(missingIDs))
	for _, id := range missingIDs {
		pbMissing = append(pbMissing, id.String())
	}

	logger.Info("batch users request processed", slog.Int("found", len(pbUsers)), slog.Int("missing", len(pbMissing)))

	return &users.BatchGetResponse{
		Users:      pbUsers,
		MissingIds: pbMissing,
	}, nil
}
