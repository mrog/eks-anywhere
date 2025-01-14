// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/controller/clusters/validations.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	v1alpha1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	gomock "github.com/golang/mock/gomock"
)

// MockIPUniquenessValidator is a mock of IPUniquenessValidator interface.
type MockIPUniquenessValidator struct {
	ctrl     *gomock.Controller
	recorder *MockIPUniquenessValidatorMockRecorder
}

// MockIPUniquenessValidatorMockRecorder is the mock recorder for MockIPUniquenessValidator.
type MockIPUniquenessValidatorMockRecorder struct {
	mock *MockIPUniquenessValidator
}

// NewMockIPUniquenessValidator creates a new mock instance.
func NewMockIPUniquenessValidator(ctrl *gomock.Controller) *MockIPUniquenessValidator {
	mock := &MockIPUniquenessValidator{ctrl: ctrl}
	mock.recorder = &MockIPUniquenessValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIPUniquenessValidator) EXPECT() *MockIPUniquenessValidatorMockRecorder {
	return m.recorder
}

// ValidateControlPlaneIPUniqueness mocks base method.
func (m *MockIPUniquenessValidator) ValidateControlPlaneIPUniqueness(cluster *v1alpha1.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateControlPlaneIPUniqueness", cluster)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateControlPlaneIPUniqueness indicates an expected call of ValidateControlPlaneIPUniqueness.
func (mr *MockIPUniquenessValidatorMockRecorder) ValidateControlPlaneIPUniqueness(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateControlPlaneIPUniqueness", reflect.TypeOf((*MockIPUniquenessValidator)(nil).ValidateControlPlaneIPUniqueness), cluster)
}
