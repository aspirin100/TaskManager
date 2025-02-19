// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/tasks/handlers/tests/create_task_test.go
//
// Generated by this command:
//
//	mockgen -source=./internal/tasks/handlers/tests/create_task_test.go -destination=./internal/tasks/handlers/tests/mocks/task_creator_mock.go
//

// Package mock_create_test is a generated GoMock package.
package mock_create_test

import (
	context "context"
	reflect "reflect"

	tasks "github.com/aspirin100/TaskManager/internal/tasks"
	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MocktaskCreator is a mock of taskCreator interface.
type MocktaskCreator struct {
	ctrl     *gomock.Controller
	recorder *MocktaskCreatorMockRecorder
	isgomock struct{}
}

// MocktaskCreatorMockRecorder is the mock recorder for MocktaskCreator.
type MocktaskCreatorMockRecorder struct {
	mock *MocktaskCreator
}

// NewMocktaskCreator creates a new mock instance.
func NewMocktaskCreator(ctrl *gomock.Controller) *MocktaskCreator {
	mock := &MocktaskCreator{ctrl: ctrl}
	mock.recorder = &MocktaskCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocktaskCreator) EXPECT() *MocktaskCreatorMockRecorder {
	return m.recorder
}

// CreateTask mocks base method.
func (m *MocktaskCreator) CreateTask(ctx context.Context, params tasks.CreateTaskRequest) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTask", ctx, params)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTask indicates an expected call of CreateTask.
func (mr *MocktaskCreatorMockRecorder) CreateTask(ctx, params any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTask", reflect.TypeOf((*MocktaskCreator)(nil).CreateTask), ctx, params)
}
