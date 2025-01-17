// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	models "MessageProcessing/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockMessage is a mock of Message interface.
type MockMessage struct {
	ctrl     *gomock.Controller
	recorder *MockMessageMockRecorder
}

// MockMessageMockRecorder is the mock recorder for MockMessage.
type MockMessageMockRecorder struct {
	mock *MockMessage
}

// NewMockMessage creates a new mock instance.
func NewMockMessage(ctrl *gomock.Controller) *MockMessage {
	mock := &MockMessage{ctrl: ctrl}
	mock.recorder = &MockMessageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessage) EXPECT() *MockMessageMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockMessage) Create(message models.Message) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", message)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockMessageMockRecorder) Create(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMessage)(nil).Create), message)
}

// GetCompMessages mocks base method.
func (m *MockMessage) GetCompMessages() ([]models.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCompMessages")
	ret0, _ := ret[0].([]models.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCompMessages indicates an expected call of GetCompMessages.
func (mr *MockMessageMockRecorder) GetCompMessages() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCompMessages", reflect.TypeOf((*MockMessage)(nil).GetCompMessages))
}

// GetCurMessages mocks base method.
func (m *MockMessage) GetCurMessages() ([]models.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurMessages")
	ret0, _ := ret[0].([]models.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurMessages indicates an expected call of GetCurMessages.
func (mr *MockMessageMockRecorder) GetCurMessages() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurMessages", reflect.TypeOf((*MockMessage)(nil).GetCurMessages))
}

// ScanAndResend mocks base method.
func (m *MockMessage) ScanAndResend() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ScanAndResend")
}

// ScanAndResend indicates an expected call of ScanAndResend.
func (mr *MockMessageMockRecorder) ScanAndResend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScanAndResend", reflect.TypeOf((*MockMessage)(nil).ScanAndResend))
}

// SendToKafka mocks base method.
func (m *MockMessage) SendToKafka(message models.Message) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendToKafka", message)
}

// SendToKafka indicates an expected call of SendToKafka.
func (mr *MockMessageMockRecorder) SendToKafka(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendToKafka", reflect.TypeOf((*MockMessage)(nil).SendToKafka), message)
}
