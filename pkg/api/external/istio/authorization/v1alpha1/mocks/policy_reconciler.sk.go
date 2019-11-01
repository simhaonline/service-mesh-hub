// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/api/external/istio/authorization/v1alpha1/policy_reconciler.sk.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/mesh-projects/pkg/api/external/istio/authorization/v1alpha1"
	clients "github.com/solo-io/solo-kit/pkg/api/v1/clients"
	reflect "reflect"
)

// MockPolicyReconciler is a mock of PolicyReconciler interface
type MockPolicyReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockPolicyReconcilerMockRecorder
}

// MockPolicyReconcilerMockRecorder is the mock recorder for MockPolicyReconciler
type MockPolicyReconcilerMockRecorder struct {
	mock *MockPolicyReconciler
}

// NewMockPolicyReconciler creates a new mock instance
func NewMockPolicyReconciler(ctrl *gomock.Controller) *MockPolicyReconciler {
	mock := &MockPolicyReconciler{ctrl: ctrl}
	mock.recorder = &MockPolicyReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPolicyReconciler) EXPECT() *MockPolicyReconcilerMockRecorder {
	return m.recorder
}

// Reconcile mocks base method
func (m *MockPolicyReconciler) Reconcile(namespace string, desiredResources v1alpha1.PolicyList, transition v1alpha1.TransitionPolicyFunc, opts clients.ListOpts) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reconcile", namespace, desiredResources, transition, opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reconcile indicates an expected call of Reconcile
func (mr *MockPolicyReconcilerMockRecorder) Reconcile(namespace, desiredResources, transition, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reconcile", reflect.TypeOf((*MockPolicyReconciler)(nil).Reconcile), namespace, desiredResources, transition, opts)
}
