// Code generated by MockGen. DO NOT EDIT.
// Source: ./gomoneypb/configuration/v1/configurationv1connect/configuration.connect.go

// Package configurationv1connect is a generated GoMock package.
package configurationv1connect

import (
	context "context"
	reflect "reflect"

	connect "connectrpc.com/connect"
	v1 "github.com/ft-t/go-money-pb/gen/gomoneypb/configuration/v1"
	gomock "github.com/golang/mock/gomock"
)

// MockConfigurationServiceClient is a mock of ConfigurationServiceClient interface.
type MockConfigurationServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockConfigurationServiceClientMockRecorder
}

// MockConfigurationServiceClientMockRecorder is the mock recorder for MockConfigurationServiceClient.
type MockConfigurationServiceClientMockRecorder struct {
	mock *MockConfigurationServiceClient
}

// NewMockConfigurationServiceClient creates a new mock instance.
func NewMockConfigurationServiceClient(ctrl *gomock.Controller) *MockConfigurationServiceClient {
	mock := &MockConfigurationServiceClient{ctrl: ctrl}
	mock.recorder = &MockConfigurationServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfigurationServiceClient) EXPECT() *MockConfigurationServiceClientMockRecorder {
	return m.recorder
}

// GetConfiguration mocks base method.
func (m *MockConfigurationServiceClient) GetConfiguration(arg0 context.Context, arg1 *connect.Request[v1.GetConfigurationRequest]) (*connect.Response[v1.GetConfigurationResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfiguration", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.GetConfigurationResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfiguration indicates an expected call of GetConfiguration.
func (mr *MockConfigurationServiceClientMockRecorder) GetConfiguration(arg0, arg1 interface{}) *ConfigurationServiceClientGetConfigurationCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfiguration", reflect.TypeOf((*MockConfigurationServiceClient)(nil).GetConfiguration), arg0, arg1)
	return &ConfigurationServiceClientGetConfigurationCall{Call: call}
}

// ConfigurationServiceClientGetConfigurationCall wrap *gomock.Call
type ConfigurationServiceClientGetConfigurationCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ConfigurationServiceClientGetConfigurationCall) Return(arg0 *connect.Response[v1.GetConfigurationResponse], arg1 error) *ConfigurationServiceClientGetConfigurationCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ConfigurationServiceClientGetConfigurationCall) Do(f func(context.Context, *connect.Request[v1.GetConfigurationRequest]) (*connect.Response[v1.GetConfigurationResponse], error)) *ConfigurationServiceClientGetConfigurationCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ConfigurationServiceClientGetConfigurationCall) DoAndReturn(f func(context.Context, *connect.Request[v1.GetConfigurationRequest]) (*connect.Response[v1.GetConfigurationResponse], error)) *ConfigurationServiceClientGetConfigurationCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockConfigurationServiceHandler is a mock of ConfigurationServiceHandler interface.
type MockConfigurationServiceHandler struct {
	ctrl     *gomock.Controller
	recorder *MockConfigurationServiceHandlerMockRecorder
}

// MockConfigurationServiceHandlerMockRecorder is the mock recorder for MockConfigurationServiceHandler.
type MockConfigurationServiceHandlerMockRecorder struct {
	mock *MockConfigurationServiceHandler
}

// NewMockConfigurationServiceHandler creates a new mock instance.
func NewMockConfigurationServiceHandler(ctrl *gomock.Controller) *MockConfigurationServiceHandler {
	mock := &MockConfigurationServiceHandler{ctrl: ctrl}
	mock.recorder = &MockConfigurationServiceHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfigurationServiceHandler) EXPECT() *MockConfigurationServiceHandlerMockRecorder {
	return m.recorder
}

// GetConfiguration mocks base method.
func (m *MockConfigurationServiceHandler) GetConfiguration(arg0 context.Context, arg1 *connect.Request[v1.GetConfigurationRequest]) (*connect.Response[v1.GetConfigurationResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfiguration", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.GetConfigurationResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfiguration indicates an expected call of GetConfiguration.
func (mr *MockConfigurationServiceHandlerMockRecorder) GetConfiguration(arg0, arg1 interface{}) *ConfigurationServiceHandlerGetConfigurationCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfiguration", reflect.TypeOf((*MockConfigurationServiceHandler)(nil).GetConfiguration), arg0, arg1)
	return &ConfigurationServiceHandlerGetConfigurationCall{Call: call}
}

// ConfigurationServiceHandlerGetConfigurationCall wrap *gomock.Call
type ConfigurationServiceHandlerGetConfigurationCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ConfigurationServiceHandlerGetConfigurationCall) Return(arg0 *connect.Response[v1.GetConfigurationResponse], arg1 error) *ConfigurationServiceHandlerGetConfigurationCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ConfigurationServiceHandlerGetConfigurationCall) Do(f func(context.Context, *connect.Request[v1.GetConfigurationRequest]) (*connect.Response[v1.GetConfigurationResponse], error)) *ConfigurationServiceHandlerGetConfigurationCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ConfigurationServiceHandlerGetConfigurationCall) DoAndReturn(f func(context.Context, *connect.Request[v1.GetConfigurationRequest]) (*connect.Response[v1.GetConfigurationResponse], error)) *ConfigurationServiceHandlerGetConfigurationCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
