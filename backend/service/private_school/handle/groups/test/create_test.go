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

func TestCreateGroup_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := groups_handle_test.NewMockGroupsService(ctrl)
	handler := createTestHandler(mockService)

	clientID := uuid.New()
	req := &private_school_v1.CreateGroupRequest{ClientId: clientID.String(), Name: "Test Group"}
	expectedGroup := &groups_models.Group{PublicID: uuid.New(), ClientID: clientID, Name: "Test Group"}
	mockService.EXPECT().CreateGroup(gomock.Any(), gomock.AssignableToTypeOf(&groups_models.CreateGroup{})).Return(expectedGroup, nil)
	resp, err := handler.CreateGroup(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedGroup.PublicID.String(), resp.Id)
	assert.Equal(t, "Test Group", resp.Name)
}

func TestCreateGroup_ValidationError_EmptyName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := groups_handle_test.NewMockGroupsService(ctrl)
	handler := createTestHandler(mockService)

	req := &private_school_v1.CreateGroupRequest{ClientId: uuid.New().String(), Name: ""}
	resp, err := handler.CreateGroup(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, resp)
	groups_handle_test.AssertGRPCError(t, err, codes.InvalidArgument, "group name is required")
}

func TestCreateGroup_ValidationError_LongName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := groups_handle_test.NewMockGroupsService(ctrl)
	handler := createTestHandler(mockService)

	longName := "This is a very long group name that exceeds the maximum allowed length of 100 characters by a significant margin"
	req := &private_school_v1.CreateGroupRequest{ClientId: uuid.New().String(), Name: longName}
	resp, err := handler.CreateGroup(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, resp)
	groups_handle_test.AssertGRPCError(t, err, codes.InvalidArgument, "group name exceeds maximum length (100 characters)")
}

func TestCreateGroup_InvalidClientID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := groups_handle_test.NewMockGroupsService(ctrl)
	handler := createTestHandler(mockService)

	req := &private_school_v1.CreateGroupRequest{ClientId: "invalid-uuid", Name: "Test Group"}
	resp, err := handler.CreateGroup(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, resp)
	groups_handle_test.AssertGRPCError(t, err, codes.InvalidArgument, "invalid client ID format")
}

func TestCreateGroup_InvalidCuratorID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := groups_handle_test.NewMockGroupsService(ctrl)
	handler := createTestHandler(mockService)

	invalidCuratorID := "invalid-uuid"
	req := &private_school_v1.CreateGroupRequest{ClientId: uuid.New().String(), Name: "Test Group", CuratorId: &invalidCuratorID}
	resp, err := handler.CreateGroup(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, resp)
	groups_handle_test.AssertGRPCError(t, err, codes.InvalidArgument, "invalid curator ID format")
}

func TestCreateGroup_DuplicateError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := groups_handle_test.NewMockGroupsService(ctrl)
	handler := createTestHandler(mockService)

	clientID := uuid.New()
	req := &private_school_v1.CreateGroupRequest{ClientId: clientID.String(), Name: "Duplicate Group"}
	mockService.EXPECT().CreateGroup(gomock.Any(), gomock.AssignableToTypeOf(&groups_models.CreateGroup{})).Return(nil, groups_models.ErrDuplicateGroupName)
	resp, err := handler.CreateGroup(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, resp)
	groups_handle_test.AssertGRPCError(t, err, codes.AlreadyExists, "group name already exists for this client")
}

func TestCreateGroup_InternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := groups_handle_test.NewMockGroupsService(ctrl)
	handler := createTestHandler(mockService)

	clientID := uuid.New()
	req := &private_school_v1.CreateGroupRequest{
		ClientId: clientID.String(),
		Name:     "Test Group",
	}

	mockService.EXPECT().CreateGroup(gomock.Any(), gomock.AssignableToTypeOf(&groups_models.CreateGroup{})).Return(nil, errors.New("database connection failed"))
	resp, err := handler.CreateGroup(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, resp)
	groups_handle_test.AssertGRPCError(t, err, codes.Internal, "internal server error")
}

func TestCreateGroup_WithCurator_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := groups_handle_test.NewMockGroupsService(ctrl)
	handler := createTestHandler(mockService)

	clientID := uuid.New()
	curatorID := uuid.New()
	curatorStringID := curatorID.String()

	req := &private_school_v1.CreateGroupRequest{
		ClientId:  clientID.String(),
		Name:      "Test Group with Curator",
		CuratorId: &curatorStringID,
	}

	expectedGroup := &groups_models.Group{
		PublicID:  uuid.New(),
		ClientID:  clientID,
		Name:      "Test Group with Curator",
		CuratorID: &curatorID,
	}

	mockService.EXPECT().CreateGroup(gomock.Any(), gomock.AssignableToTypeOf(&groups_models.CreateGroup{})).Return(expectedGroup, nil)
	resp, err := handler.CreateGroup(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedGroup.PublicID.String(), resp.Id)
	assert.Equal(t, "Test Group with Curator", resp.Name)
	assert.NotNil(t, resp.CuratorId)
	assert.Equal(t, curatorID.String(), *resp.CuratorId)
}
