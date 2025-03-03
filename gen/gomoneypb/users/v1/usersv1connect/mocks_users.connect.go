// Code generated by MockGen. DO NOT EDIT.
// Source: ./gomoneypb/users/v1/usersv1connect/users.connect.go

// Package usersv1connect is a generated GoMock package.
package usersv1connect

import (
	context "context"
	reflect "reflect"

	connect "connectrpc.com/connect"
	v1 "github.com/ft-t/go-money-pb/gen/gomoneypb/users/v1"
	gomock "github.com/golang/mock/gomock"
)

// MockUsersServiceClient is a mock of UsersServiceClient interface.
type MockUsersServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockUsersServiceClientMockRecorder
}

// MockUsersServiceClientMockRecorder is the mock recorder for MockUsersServiceClient.
type MockUsersServiceClientMockRecorder struct {
	mock *MockUsersServiceClient
}

// NewMockUsersServiceClient creates a new mock instance.
func NewMockUsersServiceClient(ctrl *gomock.Controller) *MockUsersServiceClient {
	mock := &MockUsersServiceClient{ctrl: ctrl}
	mock.recorder = &MockUsersServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsersServiceClient) EXPECT() *MockUsersServiceClientMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUsersServiceClient) Create(arg0 context.Context, arg1 *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.CreateResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUsersServiceClientMockRecorder) Create(arg0, arg1 interface{}) *UsersServiceClientCreateCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUsersServiceClient)(nil).Create), arg0, arg1)
	return &UsersServiceClientCreateCall{Call: call}
}

// UsersServiceClientCreateCall wrap *gomock.Call
type UsersServiceClientCreateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *UsersServiceClientCreateCall) Return(arg0 *connect.Response[v1.CreateResponse], arg1 error) *UsersServiceClientCreateCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *UsersServiceClientCreateCall) Do(f func(context.Context, *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error)) *UsersServiceClientCreateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *UsersServiceClientCreateCall) DoAndReturn(f func(context.Context, *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error)) *UsersServiceClientCreateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Login mocks base method.
func (m *MockUsersServiceClient) Login(arg0 context.Context, arg1 *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.LoginResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockUsersServiceClientMockRecorder) Login(arg0, arg1 interface{}) *UsersServiceClientLoginCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUsersServiceClient)(nil).Login), arg0, arg1)
	return &UsersServiceClientLoginCall{Call: call}
}

// UsersServiceClientLoginCall wrap *gomock.Call
type UsersServiceClientLoginCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *UsersServiceClientLoginCall) Return(arg0 *connect.Response[v1.LoginResponse], arg1 error) *UsersServiceClientLoginCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *UsersServiceClientLoginCall) Do(f func(context.Context, *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error)) *UsersServiceClientLoginCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *UsersServiceClientLoginCall) DoAndReturn(f func(context.Context, *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error)) *UsersServiceClientLoginCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockUsersServiceHandler is a mock of UsersServiceHandler interface.
type MockUsersServiceHandler struct {
	ctrl     *gomock.Controller
	recorder *MockUsersServiceHandlerMockRecorder
}

// MockUsersServiceHandlerMockRecorder is the mock recorder for MockUsersServiceHandler.
type MockUsersServiceHandlerMockRecorder struct {
	mock *MockUsersServiceHandler
}

// NewMockUsersServiceHandler creates a new mock instance.
func NewMockUsersServiceHandler(ctrl *gomock.Controller) *MockUsersServiceHandler {
	mock := &MockUsersServiceHandler{ctrl: ctrl}
	mock.recorder = &MockUsersServiceHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsersServiceHandler) EXPECT() *MockUsersServiceHandlerMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUsersServiceHandler) Create(arg0 context.Context, arg1 *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.CreateResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUsersServiceHandlerMockRecorder) Create(arg0, arg1 interface{}) *UsersServiceHandlerCreateCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUsersServiceHandler)(nil).Create), arg0, arg1)
	return &UsersServiceHandlerCreateCall{Call: call}
}

// UsersServiceHandlerCreateCall wrap *gomock.Call
type UsersServiceHandlerCreateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *UsersServiceHandlerCreateCall) Return(arg0 *connect.Response[v1.CreateResponse], arg1 error) *UsersServiceHandlerCreateCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *UsersServiceHandlerCreateCall) Do(f func(context.Context, *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error)) *UsersServiceHandlerCreateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *UsersServiceHandlerCreateCall) DoAndReturn(f func(context.Context, *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error)) *UsersServiceHandlerCreateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Login mocks base method.
func (m *MockUsersServiceHandler) Login(arg0 context.Context, arg1 *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.LoginResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockUsersServiceHandlerMockRecorder) Login(arg0, arg1 interface{}) *UsersServiceHandlerLoginCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUsersServiceHandler)(nil).Login), arg0, arg1)
	return &UsersServiceHandlerLoginCall{Call: call}
}

// UsersServiceHandlerLoginCall wrap *gomock.Call
type UsersServiceHandlerLoginCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *UsersServiceHandlerLoginCall) Return(arg0 *connect.Response[v1.LoginResponse], arg1 error) *UsersServiceHandlerLoginCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *UsersServiceHandlerLoginCall) Do(f func(context.Context, *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error)) *UsersServiceHandlerLoginCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *UsersServiceHandlerLoginCall) DoAndReturn(f func(context.Context, *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error)) *UsersServiceHandlerLoginCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
