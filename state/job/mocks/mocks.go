// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/cloudfoundry/bosh-init/state/job (interfaces: DependencyCompiler)

package mocks

import (
	gomock "github.com/cloudfoundry/bosh-init/internal/github.com/golang/mock/gomock"
	job "github.com/cloudfoundry/bosh-init/release/job"
	job0 "github.com/cloudfoundry/bosh-init/state/job"
	ui "github.com/cloudfoundry/bosh-init/ui"
)

// Mock of DependencyCompiler interface
type MockDependencyCompiler struct {
	ctrl     *gomock.Controller
	recorder *_MockDependencyCompilerRecorder
}

// Recorder for MockDependencyCompiler (not exported)
type _MockDependencyCompilerRecorder struct {
	mock *MockDependencyCompiler
}

func NewMockDependencyCompiler(ctrl *gomock.Controller) *MockDependencyCompiler {
	mock := &MockDependencyCompiler{ctrl: ctrl}
	mock.recorder = &_MockDependencyCompilerRecorder{mock}
	return mock
}

func (_m *MockDependencyCompiler) EXPECT() *_MockDependencyCompilerRecorder {
	return _m.recorder
}

func (_m *MockDependencyCompiler) Compile(_param0 []job.Job, _param1 ui.Stage) ([]job0.CompiledPackageRef, error) {
	ret := _m.ctrl.Call(_m, "Compile", _param0, _param1)
	ret0, _ := ret[0].([]job0.CompiledPackageRef)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDependencyCompilerRecorder) Compile(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Compile", arg0, arg1)
}
