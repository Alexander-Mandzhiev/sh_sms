package app_manager_handle

import (
	apps "backend/service/clients/service/apps_manager"
	"log/slog"
)

func validateIdentifier(logger *slog.Logger, id *int32, code *string) error {
	switch {
	case id == nil && code == nil:
		logger.Error("No identifier provided")
		return apps.ErrIdentifierRequired

	case id != nil && code != nil:
		logger.Error("Conflicting identifiers", slog.Int("id", int(*id)), slog.String("code", *code))
		return apps.ErrConflictParams

	case id != nil && *id <= 0:
		logger.Error("Invalid ID", slog.Int("id", int(*id)))
		return apps.ErrInvalidID

	case code != nil && *code == "":
		logger.Error("Empty code")
		return apps.ErrEmptyCode

	default:
		return nil
	}
}
