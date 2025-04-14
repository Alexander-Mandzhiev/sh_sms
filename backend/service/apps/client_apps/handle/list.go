package handle

import (
	pb "backend/protos/gen/go/apps/clients_apps"
	"context"
	"log/slog"
)

func (s *serverAPI) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	const op = "handler.List"
	logger := s.logger.With(slog.String("op", op))
	logger.Debug("starting operation", slog.Any("request", req))

	if err := validatePagination(req.Page, req.Count); err != nil {
		return nil, err
	}

	if req.ClientId != nil {
		if err := validateClientID(*req.ClientId); err != nil {
			return nil, err
		}
	}

	if req.AppId != nil {
		if err := validateAppID(*req.AppId); err != nil {
			return nil, err
		}
	}

	resp, err := s.service.List(ctx, req)
	if err != nil {
		return nil, s.handleError(op, err)
	}

	logger.Debug("operation completed", slog.Int("count", len(resp.ClientApps)))
	return resp, nil
}
