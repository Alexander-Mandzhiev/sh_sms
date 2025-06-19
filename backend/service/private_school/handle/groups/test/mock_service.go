package groups_handle_test

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"

	"backend/pkg/models/groups"
)

type MockGroupsService struct {
	ctrl     *gomock.Controller
	recorder *MockGroupsServiceRecorder
}

type MockGroupsServiceRecorder struct {
	mock *MockGroupsService
}

func NewMockGroupsService(ctrl *gomock.Controller) *MockGroupsService {
	mock := &MockGroupsService{ctrl: ctrl}
	mock.recorder = &MockGroupsServiceRecorder{mock}
	return mock
}

func (m *MockGroupsService) EXPECT() *MockGroupsServiceRecorder {
	return m.recorder
}

func (m *MockGroupsService) CreateGroup(ctx context.Context, req *groups_models.CreateGroup) (*groups_models.Group, error) {
	m.ctrl.T.Helper()
	args := m.ctrl.Call(m, "CreateGroup", ctx, req)

	var result *groups_models.Group
	if args[0] != nil {
		result = args[0].(*groups_models.Group)
	}

	var err error
	if args[1] != nil {
		err = args[1].(error)
	}

	return result, err
}

func (mr *MockGroupsServiceRecorder) CreateGroup(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCall(mr.mock, "CreateGroup", ctx, req)
}

func (m *MockGroupsService) GetGroup(ctx context.Context, publicID, clientID uuid.UUID) (*groups_models.Group, error) {
	m.ctrl.T.Helper()
	args := m.ctrl.Call(m, "GetGroup", ctx, publicID, clientID)

	var result *groups_models.Group
	if args[0] != nil {
		result = args[0].(*groups_models.Group)
	}

	var err error
	if args[1] != nil {
		err = args[1].(error)
	}

	return result, err
}

func (mr *MockGroupsServiceRecorder) GetGroup(ctx, publicID, clientID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCall(mr.mock, "GetGroup", ctx, publicID, clientID)
}

func (m *MockGroupsService) ListGroups(ctx context.Context, req *groups_models.ListGroupsRequest) (*groups_models.GroupListResponse, error) {
	m.ctrl.T.Helper()
	args := m.ctrl.Call(m, "ListGroups", ctx, req)

	var response *groups_models.GroupListResponse
	if args[0] != nil {
		response = args[0].(*groups_models.GroupListResponse)
	}

	var err error
	if args[1] != nil {
		err = args[1].(error)
	}

	return response, err
}

func (mr *MockGroupsServiceRecorder) ListGroups(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCall(mr.mock, "ListGroups", ctx, req)
}

func (m *MockGroupsService) UpdateGroup(ctx context.Context, req *groups_models.UpdateGroup) (*groups_models.Group, error) {
	m.ctrl.T.Helper()
	args := m.ctrl.Call(m, "UpdateGroup", ctx, req)

	var result *groups_models.Group
	if args[0] != nil {
		result = args[0].(*groups_models.Group)
	}

	var err error
	if args[1] != nil {
		err = args[1].(error)
	}

	return result, err
}

func (mr *MockGroupsServiceRecorder) UpdateGroup(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCall(mr.mock, "UpdateGroup", ctx, req)
}

func (m *MockGroupsService) DeleteGroup(ctx context.Context, publicID, clientID uuid.UUID) error {
	m.ctrl.T.Helper()
	args := m.ctrl.Call(m, "DeleteGroup", ctx, publicID, clientID)

	if args[0] != nil {
		return args[0].(error)
	}
	return nil
}

func (mr *MockGroupsServiceRecorder) DeleteGroup(ctx, publicID, clientID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCall(mr.mock, "DeleteGroup", ctx, publicID, clientID)
}
