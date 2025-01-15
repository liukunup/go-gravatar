// Code generated by MockGen. DO NOT EDIT.
// Source: internal/service/avatar.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	v1 "go-gravatar/api/v1"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAvatarService is a mock of AvatarService interface.
type MockAvatarService struct {
	ctrl     *gomock.Controller
	recorder *MockAvatarServiceMockRecorder
}

// MockAvatarServiceMockRecorder is the mock recorder for MockAvatarService.
type MockAvatarServiceMockRecorder struct {
	mock *MockAvatarService
}

// NewMockAvatarService creates a new mock instance.
func NewMockAvatarService(ctrl *gomock.Controller) *MockAvatarService {
	mock := &MockAvatarService{ctrl: ctrl}
	mock.recorder = &MockAvatarServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAvatarService) EXPECT() *MockAvatarServiceMockRecorder {
	return m.recorder
}

// DeleteAvatar mocks base method.
func (m *MockAvatarService) DeleteAvatar(ctx context.Context, userId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAvatar", ctx, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAvatar indicates an expected call of DeleteAvatar.
func (mr *MockAvatarServiceMockRecorder) DeleteAvatar(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAvatar", reflect.TypeOf((*MockAvatarService)(nil).DeleteAvatar), ctx, userId)
}

// GetAvatar mocks base method.
func (m *MockAvatarService) GetAvatar(ctx context.Context, req *v1.GetAvatarRequest) (*v1.GetAvatarResponseData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAvatar", ctx, req)
	ret0, _ := ret[0].(*v1.GetAvatarResponseData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAvatar indicates an expected call of GetAvatar.
func (mr *MockAvatarServiceMockRecorder) GetAvatar(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAvatar", reflect.TypeOf((*MockAvatarService)(nil).GetAvatar), ctx, req)
}

// UpdateAvatar mocks base method.
func (m *MockAvatarService) UpdateAvatar(ctx context.Context, userId string, req *v1.UpdateAvatarRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAvatar", ctx, userId, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAvatar indicates an expected call of UpdateAvatar.
func (mr *MockAvatarServiceMockRecorder) UpdateAvatar(ctx, userId, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAvatar", reflect.TypeOf((*MockAvatarService)(nil).UpdateAvatar), ctx, userId, req)
}
