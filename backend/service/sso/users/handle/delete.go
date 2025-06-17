package handle

import (
	"backend/pkg/utils"
	"context"
	"log/slog"

	"backend/protos/gen/go/sso/users"
)

func (s *serverAPI) DeleteUser(ctx context.Context, req *users.DeleteRequest) (*users.SuccessResponse, error) {
	const op = "grpc.user.Delete"
	logger := s.logger.With(slog.String("op", op), slog.String("user_id", req.GetId()), slog.String("client_id", req.GetClientId()), slog.Bool("permanent", req.GetPermanent()))
	logger.Debug("attempting to delete user")

	clientID, err := utils.ValidateStringAndReturnUUID(req.GetClientId())
	if err != nil {
		logger.Warn("invalid client_id", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	userID, err := utils.ValidateStringAndReturnUUID(req.GetId())
	if err != nil {
		logger.Warn("invalid user_id", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	if err = s.service.Delete(ctx, clientID, userID, req.GetPermanent()); err != nil {
		logger.Error("delete operation failed", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	logger.Info("user deleted successfully")
	return &users.SuccessResponse{Success: true}, nil
}
