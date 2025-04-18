package handle

import (
	"backend/protos/gen/go/sso/users"
	"backend/service/utils"
	"context"
	"log/slog"
)

func (s *serverAPI) List(ctx context.Context, req *users.ListRequest) (*users.ListResponse, error) {
	const op = "grpc.user.List"
	logger := s.logger.With(slog.String("op", op))
	logger.Debug("starting users listing")

	listReq, err := s.convertListRequest(req)
	if err != nil {
		logger.Warn("invalid request parameters", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	if err = utils.ValidatePagination(listReq.Page, listReq.Count); err != nil {
		logger.Warn("invalid pagination", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	if listReq.ClientID != nil {
		if err = utils.ValidateClientID(listReq.ClientID.String()); err != nil {
			logger.Warn("invalid client ID", slog.Any("error", err))
			return nil, s.convertError(err)
		}
	}

	usersList, total, err := s.service.List(ctx, *listReq)
	if err != nil {
		logger.Error("failed to list users", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	pbUsers := make([]*users.User, 0, len(usersList))
	for _, user := range usersList {
		pbUsers = append(pbUsers, convertUserToProto(user))
	}

	logger.Info("successfully retrieved users", slog.Int("total", total), slog.Int("returned", len(usersList)))
	return &users.ListResponse{
		Users:       pbUsers,
		TotalCount:  int32(total),
		CurrentPage: req.Page,
	}, nil
}
