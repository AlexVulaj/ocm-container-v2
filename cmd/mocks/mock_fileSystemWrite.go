// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/occ/cmd/run (interfaces: FileSystemWrite)

// Package mocks is a generated GoMock package.
package mocks

import (
	io "io"
	os "os"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFileSystemWrite is a mock of FileSystemWrite interface.
type MockFileSystemWrite struct {
	ctrl     *gomock.Controller
	recorder *MockFileSystemWriteMockRecorder
}

// MockFileSystemWriteMockRecorder is the mock recorder for MockFileSystemWrite.
type MockFileSystemWriteMockRecorder struct {
	mock *MockFileSystemWrite
}

// NewMockFileSystemWrite creates a new mock instance.
func NewMockFileSystemWrite(ctrl *gomock.Controller) *MockFileSystemWrite {
	mock := &MockFileSystemWrite{ctrl: ctrl}
	mock.recorder = &MockFileSystemWriteMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileSystemWrite) EXPECT() *MockFileSystemWriteMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockFileSystemWrite) Create(arg0 string) (*os.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*os.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockFileSystemWriteMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockFileSystemWrite)(nil).Create), arg0)
}

// Fprintln mocks base method.
func (m *MockFileSystemWrite) Fprintln(arg0 io.Writer, arg1 ...interface{}) (int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Fprintln", varargs...)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fprintln indicates an expected call of Fprintln.
func (mr *MockFileSystemWriteMockRecorder) Fprintln(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fprintln", reflect.TypeOf((*MockFileSystemWrite)(nil).Fprintln), varargs...)
}

// MkdirTemp mocks base method.
func (m *MockFileSystemWrite) MkdirTemp(arg0, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MkdirTemp", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MkdirTemp indicates an expected call of MkdirTemp.
func (mr *MockFileSystemWriteMockRecorder) MkdirTemp(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MkdirTemp", reflect.TypeOf((*MockFileSystemWrite)(nil).MkdirTemp), arg0, arg1)
}

// RemoveAll mocks base method.
func (m *MockFileSystemWrite) RemoveAll(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAll", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveAll indicates an expected call of RemoveAll.
func (mr *MockFileSystemWriteMockRecorder) RemoveAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAll", reflect.TypeOf((*MockFileSystemWrite)(nil).RemoveAll), arg0)
}
