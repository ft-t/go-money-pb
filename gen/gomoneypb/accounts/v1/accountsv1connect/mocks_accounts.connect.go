// Code generated by MockGen. DO NOT EDIT.
// Source: ./gomoneypb/accounts/v1/accountsv1connect/accounts.connect.go

// Package accountsv1connect is a generated GoMock package.
package accountsv1connect

import (
	context "context"
	reflect "reflect"

	connect "connectrpc.com/connect"
	v1 "github.com/ft-t/go-money-pb/gen/gomoneypb/accounts/v1"
	gomock "github.com/golang/mock/gomock"
)

// MockAccountsServiceClient is a mock of AccountsServiceClient interface.
type MockAccountsServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockAccountsServiceClientMockRecorder
}

// MockAccountsServiceClientMockRecorder is the mock recorder for MockAccountsServiceClient.
type MockAccountsServiceClientMockRecorder struct {
	mock *MockAccountsServiceClient
}

// NewMockAccountsServiceClient creates a new mock instance.
func NewMockAccountsServiceClient(ctrl *gomock.Controller) *MockAccountsServiceClient {
	mock := &MockAccountsServiceClient{ctrl: ctrl}
	mock.recorder = &MockAccountsServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountsServiceClient) EXPECT() *MockAccountsServiceClientMockRecorder {
	return m.recorder
}

// CreateAccount mocks base method.
func (m *MockAccountsServiceClient) CreateAccount(arg0 context.Context, arg1 *connect.Request[v1.CreateAccountRequest]) (*connect.Response[v1.CreateAccountResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.CreateAccountResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockAccountsServiceClientMockRecorder) CreateAccount(arg0, arg1 interface{}) *AccountsServiceClientCreateAccountCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockAccountsServiceClient)(nil).CreateAccount), arg0, arg1)
	return &AccountsServiceClientCreateAccountCall{Call: call}
}

// AccountsServiceClientCreateAccountCall wrap *gomock.Call
type AccountsServiceClientCreateAccountCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *AccountsServiceClientCreateAccountCall) Return(arg0 *connect.Response[v1.CreateAccountResponse], arg1 error) *AccountsServiceClientCreateAccountCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *AccountsServiceClientCreateAccountCall) Do(f func(context.Context, *connect.Request[v1.CreateAccountRequest]) (*connect.Response[v1.CreateAccountResponse], error)) *AccountsServiceClientCreateAccountCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *AccountsServiceClientCreateAccountCall) DoAndReturn(f func(context.Context, *connect.Request[v1.CreateAccountRequest]) (*connect.Response[v1.CreateAccountResponse], error)) *AccountsServiceClientCreateAccountCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListAccounts mocks base method.
func (m *MockAccountsServiceClient) ListAccounts(arg0 context.Context, arg1 *connect.Request[v1.ListAccountsRequest]) (*connect.Response[v1.ListAccountsResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccounts", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.ListAccountsResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccounts indicates an expected call of ListAccounts.
func (mr *MockAccountsServiceClientMockRecorder) ListAccounts(arg0, arg1 interface{}) *AccountsServiceClientListAccountsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccounts", reflect.TypeOf((*MockAccountsServiceClient)(nil).ListAccounts), arg0, arg1)
	return &AccountsServiceClientListAccountsCall{Call: call}
}

// AccountsServiceClientListAccountsCall wrap *gomock.Call
type AccountsServiceClientListAccountsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *AccountsServiceClientListAccountsCall) Return(arg0 *connect.Response[v1.ListAccountsResponse], arg1 error) *AccountsServiceClientListAccountsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *AccountsServiceClientListAccountsCall) Do(f func(context.Context, *connect.Request[v1.ListAccountsRequest]) (*connect.Response[v1.ListAccountsResponse], error)) *AccountsServiceClientListAccountsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *AccountsServiceClientListAccountsCall) DoAndReturn(f func(context.Context, *connect.Request[v1.ListAccountsRequest]) (*connect.Response[v1.ListAccountsResponse], error)) *AccountsServiceClientListAccountsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockAccountsServiceHandler is a mock of AccountsServiceHandler interface.
type MockAccountsServiceHandler struct {
	ctrl     *gomock.Controller
	recorder *MockAccountsServiceHandlerMockRecorder
}

// MockAccountsServiceHandlerMockRecorder is the mock recorder for MockAccountsServiceHandler.
type MockAccountsServiceHandlerMockRecorder struct {
	mock *MockAccountsServiceHandler
}

// NewMockAccountsServiceHandler creates a new mock instance.
func NewMockAccountsServiceHandler(ctrl *gomock.Controller) *MockAccountsServiceHandler {
	mock := &MockAccountsServiceHandler{ctrl: ctrl}
	mock.recorder = &MockAccountsServiceHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountsServiceHandler) EXPECT() *MockAccountsServiceHandlerMockRecorder {
	return m.recorder
}

// CreateAccount mocks base method.
func (m *MockAccountsServiceHandler) CreateAccount(arg0 context.Context, arg1 *connect.Request[v1.CreateAccountRequest]) (*connect.Response[v1.CreateAccountResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.CreateAccountResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockAccountsServiceHandlerMockRecorder) CreateAccount(arg0, arg1 interface{}) *AccountsServiceHandlerCreateAccountCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockAccountsServiceHandler)(nil).CreateAccount), arg0, arg1)
	return &AccountsServiceHandlerCreateAccountCall{Call: call}
}

// AccountsServiceHandlerCreateAccountCall wrap *gomock.Call
type AccountsServiceHandlerCreateAccountCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *AccountsServiceHandlerCreateAccountCall) Return(arg0 *connect.Response[v1.CreateAccountResponse], arg1 error) *AccountsServiceHandlerCreateAccountCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *AccountsServiceHandlerCreateAccountCall) Do(f func(context.Context, *connect.Request[v1.CreateAccountRequest]) (*connect.Response[v1.CreateAccountResponse], error)) *AccountsServiceHandlerCreateAccountCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *AccountsServiceHandlerCreateAccountCall) DoAndReturn(f func(context.Context, *connect.Request[v1.CreateAccountRequest]) (*connect.Response[v1.CreateAccountResponse], error)) *AccountsServiceHandlerCreateAccountCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListAccounts mocks base method.
func (m *MockAccountsServiceHandler) ListAccounts(arg0 context.Context, arg1 *connect.Request[v1.ListAccountsRequest]) (*connect.Response[v1.ListAccountsResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccounts", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.ListAccountsResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccounts indicates an expected call of ListAccounts.
func (mr *MockAccountsServiceHandlerMockRecorder) ListAccounts(arg0, arg1 interface{}) *AccountsServiceHandlerListAccountsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccounts", reflect.TypeOf((*MockAccountsServiceHandler)(nil).ListAccounts), arg0, arg1)
	return &AccountsServiceHandlerListAccountsCall{Call: call}
}

// AccountsServiceHandlerListAccountsCall wrap *gomock.Call
type AccountsServiceHandlerListAccountsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *AccountsServiceHandlerListAccountsCall) Return(arg0 *connect.Response[v1.ListAccountsResponse], arg1 error) *AccountsServiceHandlerListAccountsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *AccountsServiceHandlerListAccountsCall) Do(f func(context.Context, *connect.Request[v1.ListAccountsRequest]) (*connect.Response[v1.ListAccountsResponse], error)) *AccountsServiceHandlerListAccountsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *AccountsServiceHandlerListAccountsCall) DoAndReturn(f func(context.Context, *connect.Request[v1.ListAccountsRequest]) (*connect.Response[v1.ListAccountsResponse], error)) *AccountsServiceHandlerListAccountsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
