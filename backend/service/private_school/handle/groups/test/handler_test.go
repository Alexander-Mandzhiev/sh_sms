package groups_handle_test_test

import (
	"backend/service/private_school/handle/groups"
	"log/slog"
)

func createTestHandler(service groups_handle.GroupsService) *groups_handle.ServerAPI {
	return groups_handle.NewServerAPI(service, slog.Default())
}
