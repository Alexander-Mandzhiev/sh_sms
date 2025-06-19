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

func TestGetGroup_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := groups_handle_test.NewMockGroupsService(ctrl)
	handler := createTestHandler(mockService)

	clientID := uuid.New()
	groupID := uuid.New()
	req := &private_school_v1.GroupRequest{Id: groupID.String(), ClientId: clientID.String()}

	expectedGroup := &groups_models.Group{
		PublicID: groupID,
		ClientID: clientID,
		Name:     "Test Group",
	}

	mockService.EXPECT().GetGroup(gomock.Any(), groupID, clientID).Return(expectedGroup, nil)
	resp, err := handler.GetGroup(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, groupID.String(), resp.Id)
	assert.Equal(t, "Test Group", resp.Name)
}

func TestGetGroup_WithCurator_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := groups_handle_test.NewMockGroupsService(ctrl)
	handler := createTestHandler(mockService)

	clientID := uuid.New()
	groupID := uuid.New()
	curatorID := uuid.New()
	req := &private_school_v1.GroupRequest{Id: groupID.String(), ClientId: clientID.String()}

	expectedGroup := &groups_models.Group{
		PublicID:  groupID,
		ClientID:  clientID,
		Name:      "Test Group with Curator",
		CuratorID: &curatorID,
	}

	mockService.EXPECT().GetGroup(gomock.Any(), groupID, clientID).Return(expectedGroup, nil)
	resp, err := handler.GetGroup(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, groupID.String(), resp.Id)
	assert.Equal(t, "Test Group with Curator", resp.Name)
	if assert.NotNil(t, resp.CuratorId) {
		assert.Equal(t, curatorID.String(), *resp.CuratorId)
	}
}

func TestGetGroup_WithoutCurator_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := groups_handle_test.NewMockGroupsService(ctrl)
	handler := createTestHandler(mockService)

	clientID := uuid.New()
	groupID := uuid.New()
	req := &private_school_v1.GroupRequest{Id: groupID.String(), ClientId: clientID.String()}

	expectedGroup := &groups_models.Group{
		PublicID:  groupID,
		ClientID:  clientID,
		Name:      "Test Group without Curator",
		CuratorID: nil,
	}

	mockService.EXPECT().GetGroup(gomock.Any(), groupID, clientID).Return(expectedGroup, nil)
	resp, err := handler.GetGroup(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, groupID.String(), resp.Id)
	assert.Equal(t, "Test Group without Curator", resp.Name)
	assert.Nil(t, resp.CuratorId)
}

func TestGetGroup_InvalidGroupID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := groups_handle_test.NewMockGroupsService(ctrl)
	handler := createTestHandler(mockService)

	req := &private_school_v1.GroupRequest{Id: "invalid-uuid", ClientId: uuid.New().String()}
	mockService.EXPECT().GetGroup(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
	resp, err := handler.GetGroup(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, resp)
	groups_handle_test.AssertGRPCError(t, err, codes.InvalidArgument, "invalid group ID format")
}

func TestGetGroup_InvalidClientID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := groups_handle_test.NewMockGroupsService(ctrl)
	handler := createTestHandler(mockService)

	req := &private_school_v1.GroupRequest{Id: uuid.New().String(), ClientId: "invalid-uuid"}
	mockService.EXPECT().GetGroup(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
	resp, err := handler.GetGroup(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, resp)
	groups_handle_test.AssertGRPCError(t, err, codes.InvalidArgument, "invalid client ID format")
}

func TestGetGroup_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := groups_handle_test.NewMockGroupsService(ctrl)
	handler := createTestHandler(mockService)

	clientID := uuid.New()
	groupID := uuid.New()
	req := &private_school_v1.GroupRequest{Id: groupID.String(), ClientId: clientID.String()}

	mockService.EXPECT().GetGroup(gomock.Any(), groupID, clientID).Return(nil, groups_models.ErrGroupNotFound)
	resp, err := handler.GetGroup(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, resp)
	groups_handle_test.AssertGRPCError(t, err, codes.NotFound, "group not found")
}

func TestGetGroup_DBError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := groups_handle_test.NewMockGroupsService(ctrl)
	handler := createTestHandler(mockService)

	clientID := uuid.New()
	groupID := uuid.New()
	req := &private_school_v1.GroupRequest{Id: groupID.String(), ClientId: clientID.String()}

	mockService.EXPECT().GetGroup(gomock.Any(), groupID, clientID).Return(nil, errors.New("connection timeout"))
	resp, err := handler.GetGroup(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, resp)
	groups_handle_test.AssertGRPCError(t, err, codes.Internal, "internal server error")
}

func TestGetGroup_ContextCanceled(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := groups_handle_test.NewMockGroupsService(ctrl)
	handler := createTestHandler(mockService)

	clientID := uuid.New()
	groupID := uuid.New()
	req := &private_school_v1.GroupRequest{Id: groupID.String(), ClientId: clientID.String()}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	mockService.EXPECT().GetGroup(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
	resp, err := handler.GetGroup(ctx, req)
	assert.Error(t, err)
	assert.Nil(t, resp)
	groups_handle_test.AssertGRPCError(t, err, codes.Canceled, "request canceled")
}
