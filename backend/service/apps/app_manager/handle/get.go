package handle

import (
	sl "backend/pkg/logger"
	pb "backend/protos/gen/go/apps/app_manager"
	"backend/service/apps/models"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *serverAPI) GetApp(ctx context.Context, req *pb.AppIdentifier) (*pb.App, error) {
	const op = "grpc.handler.AppManager.Get"
	logger := s.logger.With(slog.String("op", op))
	logger.Debug("Get request received", slog.Any("request", req))

	if err := validateNoConflict(req); err != nil {
		logger.Warn("Conflicting identifiers", sl.Err(err, true))
		return nil, s.convertError(err)
	}

	if err := validateAtLeastOne(req); err != nil {
		logger.Warn("Missing identifier", sl.Err(err, true))
		return nil, s.convertError(err)
	}

	var app *models.App
	var err error

	if id := req.GetId(); id != 0 {
		if valErr := validateID(id); valErr != nil {
			logger.Warn("Invalid ID format", sl.Err(valErr, true), slog.Int("requested_id", int(id)))
			return nil, s.convertError(valErr)
		}

		app, err = s.service.GetByID(ctx, int(id))
		if err != nil {
			logger.Error("Failed to get app by ID", sl.Err(err, true), slog.Int("converted_id", int(id)))
			return nil, s.convertError(err)
		}

		logger.Info("Successfully retrieved app by ID", slog.Int("app_id", app.ID), slog.String("app_code", app.Code))
		return s.convertAppToProto(app), nil
	}

	if code := req.GetCode(); code != "" {
		if valErr := validateCode(code, 50); valErr != nil {
			logger.Warn("Invalid code format", sl.Err(valErr, true), slog.String("requested_code", code))
			return nil, s.convertError(valErr)
		}

		app, err = s.service.GetByCode(ctx, code)
		if err != nil {
			logger.Error("Failed to get app by code", sl.Err(err, true), slog.String("requested_code", code))
			return nil, s.convertError(err)
		}

		logger.Info("Successfully retrieved app by code", slog.Int("app_id", app.ID), slog.String("app_code", app.Code))
		return s.convertAppToProto(app), nil
	}

	logger.Error("Unexpected state in Get handler")
	return nil, status.Error(codes.Internal, "internal server error")
}
