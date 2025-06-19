package groups_handle_test_test

import (
	"backend/pkg/models/groups"
	"backend/protos/gen/go/private_school"
	"backend/service/private_school/handle/groups/test"
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"testing"
)

func TestDeleteGroup_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := groups_handle_test.NewMockGroupsService(ctrl)
	handler := createTestHandler(mockService)

	clientID := uuid.New()
	groupID := uuid.New()
	req := &private_school_v1.GroupRequest{Id: groupID.String(), ClientId: clientID.String()}
	mockService.EXPECT().DeleteGroup(gomock.Any(), groupID, clientID).Return(nil)
	_, err := handler.DeleteGroup(context.Background(), req)
	assert.NoError(t, err)
}

func TestDeleteGroup_InvalidGroupID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := groups_handle_test.NewMockGroupsService(ctrl)
	handler := createTestHandler(mockService)

	req := &private_school_v1.GroupRequest{Id: "invalid-uuid", ClientId: uuid.New().String()}
	mockService.EXPECT().DeleteGroup(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
	_, err := handler.DeleteGroup(context.Background(), req)
	assert.Error(t, err)
	groups_handle_test.AssertGRPCError(t, err, codes.InvalidArgument, "invalid group ID format")
}

func TestDeleteGroup_InvalidClientID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := groups_handle_test.NewMockGroupsService(ctrl)
	handler := createTestHandler(mockService)

	req := &private_school_v1.GroupRequest{Id: uuid.New().String(), ClientId: "invalid-uuid"}
	mockService.EXPECT().DeleteGroup(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
	_, err := handler.DeleteGroup(context.Background(), req)
	assert.Error(t, err)
	groups_handle_test.AssertGRPCError(t, err, codes.InvalidArgument, "invalid client ID format")
}

func TestDeleteGroup_GroupNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := groups_handle_test.NewMockGroupsService(ctrl)
	handler := createTestHandler(mockService)

	clientID := uuid.New()
	groupID := uuid.New()
	req := &private_school_v1.GroupRequest{Id: groupID.String(), ClientId: clientID.String()}
	mockService.EXPECT().DeleteGroup(gomock.Any(), groupID, clientID).Return(groups_models.ErrGroupNotFound)
	_, err := handler.DeleteGroup(context.Background(), req)
	assert.Error(t, err)
	groups_handle_test.AssertGRPCError(t, err, codes.NotFound, "group not found")
}

func TestDeleteGroup_DependencyError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := groups_handle_test.NewMockGroupsService(ctrl)
	handler := createTestHandler(mockService)

	clientID := uuid.New()
	groupID := uuid.New()
	req := &private_school_v1.GroupRequest{Id: groupID.String(), ClientId: clientID.String()}
	mockService.EXPECT().DeleteGroup(gomock.Any(), groupID, clientID).Return(groups_models.ErrDependentRecordsExist)
	_, err := handler.DeleteGroup(context.Background(), req)
	assert.Error(t, err)
	groups_handle_test.AssertGRPCError(t, err, codes.FailedPrecondition, "cannot perform operation due to existing dependencies")
}

func TestDeleteGroup_ForeignKeyViolation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := groups_handle_test.NewMockGroupsService(ctrl)
	handler := createTestHandler(mockService)

	clientID := uuid.New()
	groupID := uuid.New()
	req := &private_school_v1.GroupRequest{Id: groupID.String(), ClientId: clientID.String()}
	mockService.EXPECT().DeleteGroup(gomock.Any(), groupID, clientID).Return(groups_models.ErrForeignKeyViolation)
	_, err := handler.DeleteGroup(context.Background(), req)
	assert.Error(t, err)
	groups_handle_test.AssertGRPCError(t, err, codes.FailedPrecondition, "cannot perform operation due to existing dependencies")
}

func TestDeleteGroup_InternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := groups_handle_test.NewMockGroupsService(ctrl)
	handler := createTestHandler(mockService)

	clientID := uuid.New()
	groupID := uuid.New()
	req := &private_school_v1.GroupRequest{Id: groupID.String(), ClientId: clientID.String()}
	mockService.EXPECT().DeleteGroup(gomock.Any(), groupID, clientID).Return(errors.New("database connection failed"))
	_, err := handler.DeleteGroup(context.Background(), req)
	assert.Error(t, err)
	groups_handle_test.AssertGRPCError(t, err, codes.Internal, "internal server error")
}

func TestDeleteGroup_ContextCanceled(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := groups_handle_test.NewMockGroupsService(ctrl)
	handler := createTestHandler(mockService)

	clientID := uuid.New()
	groupID := uuid.New()
	req := &private_school_v1.GroupRequest{Id: groupID.String(), ClientId: clientID.String()}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	mockService.EXPECT().DeleteGroup(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
	_, err := handler.DeleteGroup(ctx, req)
	assert.Error(t, err)
	groups_handle_test.AssertGRPCError(t, err, codes.Canceled, "request canceled")
}
