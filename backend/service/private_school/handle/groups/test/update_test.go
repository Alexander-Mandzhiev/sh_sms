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

func TestUpdateGroup_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := groups_handle_test.NewMockGroupsService(ctrl)
	handler := createTestHandler(mockService)

	clientID := uuid.New()
	groupID := uuid.New()
	curatorID := uuid.New().String()
	req := &private_school_v1.UpdateGroupRequest{
		Id:        groupID.String(),
		ClientId:  clientID.String(),
		Name:      "Updated Group",
		CuratorId: curatorID,
	}

	expectedGroup := &groups_models.Group{
		PublicID: groupID,
		ClientID: clientID,
		Name:     "Updated Group",
	}

	mockService.EXPECT().UpdateGroup(gomock.Any(), gomock.AssignableToTypeOf(&groups_models.UpdateGroup{})).Return(expectedGroup, nil)
	resp, err := handler.UpdateGroup(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, groupID.String(), resp.Id)
	assert.Equal(t, "Updated Group", resp.Name)
}

func TestUpdateGroup_EmptyName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := groups_handle_test.NewMockGroupsService(ctrl)
	handler := createTestHandler(mockService)

	req := &private_school_v1.UpdateGroupRequest{
		Id:        uuid.New().String(),
		ClientId:  uuid.New().String(),
		Name:      "",
		CuratorId: uuid.New().String(),
	}

	mockService.EXPECT().UpdateGroup(gomock.Any(), gomock.Any()).Times(0)
	resp, err := handler.UpdateGroup(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, resp)
	groups_handle_test.AssertGRPCError(t, err, codes.InvalidArgument, "group name is required")
}

func TestUpdateGroup_LongName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := groups_handle_test.NewMockGroupsService(ctrl)
	handler := createTestHandler(mockService)

	longName := "This is a very long group name that exceeds the maximum allowed length of 100 characters by a significant margin"
	req := &private_school_v1.UpdateGroupRequest{
		Id:        uuid.New().String(),
		ClientId:  uuid.New().String(),
		Name:      longName,
		CuratorId: uuid.New().String(),
	}

	mockService.EXPECT().UpdateGroup(gomock.Any(), gomock.Any()).Times(0)
	resp, err := handler.UpdateGroup(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, resp)
	groups_handle_test.AssertGRPCError(t, err, codes.InvalidArgument, "group name exceeds maximum length (100 characters)")
}

func TestUpdateGroup_InvalidGroupID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := groups_handle_test.NewMockGroupsService(ctrl)
	handler := createTestHandler(mockService)

	req := &private_school_v1.UpdateGroupRequest{
		Id:        "invalid-uuid",
		ClientId:  uuid.New().String(),
		Name:      "Updated Group",
		CuratorId: uuid.New().String(),
	}

	mockService.EXPECT().UpdateGroup(gomock.Any(), gomock.Any()).Times(0)
	resp, err := handler.UpdateGroup(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, resp)
	groups_handle_test.AssertGRPCError(t, err, codes.InvalidArgument, "invalid group ID format")
}

func TestUpdateGroup_InvalidClientID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := groups_handle_test.NewMockGroupsService(ctrl)
	handler := createTestHandler(mockService)

	req := &private_school_v1.UpdateGroupRequest{
		Id:        uuid.New().String(),
		ClientId:  "invalid-uuid",
		Name:      "Updated Group",
		CuratorId: uuid.New().String(),
	}

	mockService.EXPECT().UpdateGroup(gomock.Any(), gomock.Any()).Times(0)
	resp, err := handler.UpdateGroup(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, resp)
	groups_handle_test.AssertGRPCError(t, err, codes.InvalidArgument, "invalid client ID format")
}

func TestUpdateGroup_InvalidCuratorID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := groups_handle_test.NewMockGroupsService(ctrl)
	handler := createTestHandler(mockService)

	req := &private_school_v1.UpdateGroupRequest{
		Id:        uuid.New().String(),
		ClientId:  uuid.New().String(),
		Name:      "Updated Group",
		CuratorId: "invalid-uuid",
	}

	mockService.EXPECT().UpdateGroup(gomock.Any(), gomock.Any()).Times(0)
	resp, err := handler.UpdateGroup(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, resp)
	groups_handle_test.AssertGRPCError(t, err, codes.InvalidArgument, "invalid curator ID format")
}

func TestUpdateGroup_EmptyCuratorID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := groups_handle_test.NewMockGroupsService(ctrl)
	handler := createTestHandler(mockService)

	clientID := uuid.New()
	groupID := uuid.New()
	req := &private_school_v1.UpdateGroupRequest{
		Id:        groupID.String(),
		ClientId:  clientID.String(),
		Name:      "Updated Group",
		CuratorId: "",
	}

	expectedGroup := &groups_models.Group{
		PublicID:  groupID,
		ClientID:  clientID,
		Name:      "Updated Group",
		CuratorID: nil,
	}

	mockService.EXPECT().UpdateGroup(gomock.Any(), gomock.AssignableToTypeOf(&groups_models.UpdateGroup{})).Return(expectedGroup, nil)
	resp, err := handler.UpdateGroup(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, groupID.String(), resp.Id)
	assert.Equal(t, "Updated Group", resp.Name)
	assert.Nil(t, resp.CuratorId)
}

func TestUpdateGroup_DuplicateName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := groups_handle_test.NewMockGroupsService(ctrl)
	handler := createTestHandler(mockService)

	clientID := uuid.New()
	groupID := uuid.New()
	req := &private_school_v1.UpdateGroupRequest{
		Id:        groupID.String(),
		ClientId:  clientID.String(),
		Name:      "Duplicate Group",
		CuratorId: uuid.New().String(),
	}

	mockService.EXPECT().UpdateGroup(gomock.Any(), gomock.AssignableToTypeOf(&groups_models.UpdateGroup{})).Return(nil, groups_models.ErrDuplicateGroupName)
	resp, err := handler.UpdateGroup(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, resp)
	groups_handle_test.AssertGRPCError(t, err, codes.AlreadyExists, "group name already exists for this client")
}

func TestUpdateGroup_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := groups_handle_test.NewMockGroupsService(ctrl)
	handler := createTestHandler(mockService)

	clientID := uuid.New()
	groupID := uuid.New()
	req := &private_school_v1.UpdateGroupRequest{
		Id:        groupID.String(),
		ClientId:  clientID.String(),
		Name:      "Updated Group",
		CuratorId: uuid.New().String(),
	}

	mockService.EXPECT().UpdateGroup(gomock.Any(), gomock.AssignableToTypeOf(&groups_models.UpdateGroup{})).Return(nil, groups_models.ErrGroupNotFound)
	resp, err := handler.UpdateGroup(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, resp)
	groups_handle_test.AssertGRPCError(t, err, codes.NotFound, "group not found")
}

func TestUpdateGroup_InternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := groups_handle_test.NewMockGroupsService(ctrl)
	handler := createTestHandler(mockService)

	clientID := uuid.New()
	groupID := uuid.New()
	req := &private_school_v1.UpdateGroupRequest{
		Id:        groupID.String(),
		ClientId:  clientID.String(),
		Name:      "Updated Group",
		CuratorId: uuid.New().String(),
	}

	mockService.EXPECT().UpdateGroup(gomock.Any(), gomock.AssignableToTypeOf(&groups_models.UpdateGroup{})).Return(nil, errors.New("database connection failed"))
	resp, err := handler.UpdateGroup(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, resp)
	groups_handle_test.AssertGRPCError(t, err, codes.Internal, "internal server error")
}

func TestUpdateGroup_ContextCanceled(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := groups_handle_test.NewMockGroupsService(ctrl)
	handler := createTestHandler(mockService)

	clientID := uuid.New()
	groupID := uuid.New()
	req := &private_school_v1.UpdateGroupRequest{Id: groupID.String(), ClientId: clientID.String(), Name: "Updated Group", CuratorId: uuid.New().String()}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	mockService.EXPECT().UpdateGroup(gomock.Any(), gomock.Any()).Times(0)
	resp, err := handler.UpdateGroup(ctx, req)
	assert.Error(t, err)
	assert.Nil(t, resp)
	groups_handle_test.AssertGRPCError(t, err, codes.Canceled, "request canceled")
}
