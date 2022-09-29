// Code generated by MockGen. DO NOT EDIT.
// Source: ./main.go

// Package mock_usecases is a generated GoMock package.
package mock_usecases

import (
	reflect "reflect"

	model "github.com/dougmendes/gondalf/pkg/instances/model"
	gomock "github.com/golang/mock/gomock"
)

// MockInstanceSolver is a mock of InstanceSolver interface.
type MockInstanceSolver struct {
	ctrl     *gomock.Controller
	recorder *MockInstanceSolverMockRecorder
}

// MockInstanceSolverMockRecorder is the mock recorder for MockInstanceSolver.
type MockInstanceSolverMockRecorder struct {
	mock *MockInstanceSolver
}

// NewMockInstanceSolver creates a new mock instance.
func NewMockInstanceSolver(ctrl *gomock.Controller) *MockInstanceSolver {
	mock := &MockInstanceSolver{ctrl: ctrl}
	mock.recorder = &MockInstanceSolverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInstanceSolver) EXPECT() *MockInstanceSolverMockRecorder {
	return m.recorder
}

// GetAllInstances mocks base method.
func (m *MockInstanceSolver) GetAllInstances(region string) (*model.InstanceList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllInstances", region)
	ret0, _ := ret[0].(*model.InstanceList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllInstances indicates an expected call of GetAllInstances.
func (mr *MockInstanceSolverMockRecorder) GetAllInstances(region interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllInstances", reflect.TypeOf((*MockInstanceSolver)(nil).GetAllInstances), region)
}
