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

func TestListGroups_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := groups_handle_test.NewMockGroupsService(ctrl)
	handler := createTestHandler(mockService)

	clientID := uuid.New()
	cursor := int64(100)
	req := &private_school_v1.ListGroupsRequest{
		ClientId: clientID.String(),
		PageSize: 10,
		Cursor:   &cursor,
	}

	expectedResponse := &groups_models.GroupListResponse{
		Groups: []*groups_models.Group{
			{PublicID: uuid.New(), Name: "Group 1"},
			{PublicID: uuid.New(), Name: "Group 2"},
		},
		NextCursor: 90,
	}

	mockService.EXPECT().ListGroups(gomock.Any(), gomock.AssignableToTypeOf(&groups_models.ListGroupsRequest{})).Return(expectedResponse, nil)
	resp, err := handler.ListGroups(context.Background(), req)
	assert.NoError(t, err)
	assert.Len(t, resp.Groups, 2)
	assert.Equal(t, int64(90), resp.NextCursor)
}

func TestListGroups_EmptyResult(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := groups_handle_test.NewMockGroupsService(ctrl)
	handler := createTestHandler(mockService)

	clientID := uuid.New()
	req := &private_school_v1.ListGroupsRequest{ClientId: clientID.String(), PageSize: 10}
	expectedResponse := &groups_models.GroupListResponse{Groups: []*groups_models.Group{}, NextCursor: 0}
	mockService.EXPECT().ListGroups(gomock.Any(), gomock.AssignableToTypeOf(&groups_models.ListGroupsRequest{})).Return(expectedResponse, nil)
	resp, err := handler.ListGroups(context.Background(), req)
	assert.NoError(t, err)
	assert.Len(t, resp.Groups, 0)
	assert.Equal(t, int64(0), resp.NextCursor)
}

func TestListGroups_WithNameFilter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := groups_handle_test.NewMockGroupsService(ctrl)
	handler := createTestHandler(mockService)

	clientID := uuid.New()
	nameFilter := "advanced"
	req := &private_school_v1.ListGroupsRequest{ClientId: clientID.String(), PageSize: 10, NameFilter: &nameFilter}
	expectedResponse := &groups_models.GroupListResponse{Groups: []*groups_models.Group{{PublicID: uuid.New(), Name: "Advanced Group"}}, NextCursor: 0}
	mockService.EXPECT().ListGroups(gomock.Any(), gomock.AssignableToTypeOf(&groups_models.ListGroupsRequest{})).Return(expectedResponse, nil)
	resp, err := handler.ListGroups(context.Background(), req)
	assert.NoError(t, err)
	assert.Len(t, resp.Groups, 1)
	assert.Equal(t, "Advanced Group", resp.Groups[0].Name)
}

func TestListGroups_InvalidPageSize(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := groups_handle_test.NewMockGroupsService(ctrl)
	handler := createTestHandler(mockService)

	req1 := &private_school_v1.ListGroupsRequest{ClientId: uuid.New().String(), PageSize: 0}
	req2 := &private_school_v1.ListGroupsRequest{ClientId: uuid.New().String(), PageSize: 1001}
	mockService.EXPECT().ListGroups(gomock.Any(), gomock.Any()).Times(0)

	resp1, err1 := handler.ListGroups(context.Background(), req1)
	assert.Error(t, err1)
	assert.Nil(t, resp1)
	groups_handle_test.AssertGRPCError(t, err1, codes.InvalidArgument, "page size must be between 1 and 100")

	resp2, err2 := handler.ListGroups(context.Background(), req2)
	assert.Error(t, err2)
	assert.Nil(t, resp2)
	groups_handle_test.AssertGRPCError(t, err2, codes.InvalidArgument, "page size must be between 1 and 100")
}

func TestListGroups_InvalidClientID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := groups_handle_test.NewMockGroupsService(ctrl)
	handler := createTestHandler(mockService)

	req := &private_school_v1.ListGroupsRequest{ClientId: "invalid-uuid", PageSize: 10}
	mockService.EXPECT().ListGroups(gomock.Any(), gomock.Any()).Times(0)
	resp, err := handler.ListGroups(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, resp)
	groups_handle_test.AssertGRPCError(t, err, codes.InvalidArgument, "invalid client ID format")
}

func TestListGroups_InvalidCursor(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := groups_handle_test.NewMockGroupsService(ctrl)
	handler := createTestHandler(mockService)

	clientID := uuid.New()
	negativeCursor := int64(-10)
	req := &private_school_v1.ListGroupsRequest{ClientId: clientID.String(), PageSize: 10, Cursor: &negativeCursor}
	mockService.EXPECT().ListGroups(gomock.Any(), gomock.Any()).Times(0)
	resp, err := handler.ListGroups(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, resp)
	groups_handle_test.AssertGRPCError(t, err, codes.InvalidArgument, "invalid cursor value")
}

func TestListGroups_LongNameFilter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := groups_handle_test.NewMockGroupsService(ctrl)
	handler := createTestHandler(mockService)

	clientID := uuid.New()
	longNameFilter := "this is a very long filter string that exceeds the maximum allowed length of 100 characters by a significant margin"
	req := &private_school_v1.ListGroupsRequest{ClientId: clientID.String(), PageSize: 10, NameFilter: &longNameFilter}
	mockService.EXPECT().ListGroups(gomock.Any(), gomock.Any()).Times(0)
	resp, err := handler.ListGroups(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, resp)
	groups_handle_test.AssertGRPCError(t, err, codes.InvalidArgument, "filter value exceeds maximum length")
}

func TestListGroups_DBError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := groups_handle_test.NewMockGroupsService(ctrl)
	handler := createTestHandler(mockService)

	clientID := uuid.New()
	req := &private_school_v1.ListGroupsRequest{ClientId: clientID.String(), PageSize: 10}

	mockService.EXPECT().ListGroups(gomock.Any(), gomock.AssignableToTypeOf(&groups_models.ListGroupsRequest{})).Return(nil, errors.New("query execution failed"))
	resp, err := handler.ListGroups(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, resp)
	groups_handle_test.AssertGRPCError(t, err, codes.Internal, "internal server error")
}

func TestListGroups_ContextCanceled(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := groups_handle_test.NewMockGroupsService(ctrl)
	handler := createTestHandler(mockService)

	clientID := uuid.New()
	req := &private_school_v1.ListGroupsRequest{ClientId: clientID.String(), PageSize: 10}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	mockService.EXPECT().ListGroups(gomock.Any(), gomock.Any()).Times(0)
	resp, err := handler.ListGroups(ctx, req)
	assert.Error(t, err)
	assert.Nil(t, resp)
	groups_handle_test.AssertGRPCError(t, err, codes.Canceled, "request canceled")
}

func TestListGroups_LastPage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := groups_handle_test.NewMockGroupsService(ctrl)
	handler := createTestHandler(mockService)

	clientID := uuid.New()
	cursor := int64(50)
	req := &private_school_v1.ListGroupsRequest{ClientId: clientID.String(), PageSize: 10, Cursor: &cursor}

	groups := make([]*groups_models.Group, 10)
	for i := range groups {
		groups[i] = &groups_models.Group{PublicID: uuid.New(), Name: "Group"}
	}

	expectedResponse := &groups_models.GroupListResponse{Groups: groups, NextCursor: 0}
	mockService.EXPECT().ListGroups(gomock.Any(), gomock.AssignableToTypeOf(&groups_models.ListGroupsRequest{})).Return(expectedResponse, nil)
	resp, err := handler.ListGroups(context.Background(), req)

	assert.NoError(t, err)
	assert.Len(t, resp.Groups, 10)
	assert.Equal(t, int64(0), resp.NextCursor)
}

func TestListGroupsParamsFromProto_InvalidCursor(t *testing.T) {
	negativeCursor := int64(-10)
	req := &private_school_v1.ListGroupsRequest{
		ClientId: uuid.New().String(),
		PageSize: 10,
		Cursor:   &negativeCursor,
	}

	_, err := groups_models.ListGroupsParamsFromProto(req)
	assert.Error(t, err)
	assert.Equal(t, groups_models.ErrInvalidCursor, err)
}
