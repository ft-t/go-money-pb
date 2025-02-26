// Code generated by MockGen. DO NOT EDIT.
// Source: ./gomoneypb/currency/v1/currencyv1connect/currency.connect.go

// Package currencyv1connect is a generated GoMock package.
package currencyv1connect

import (
	context "context"
	reflect "reflect"

	connect "connectrpc.com/connect"
	v1 "github.com/ft-t/go-money-pb/gen/gomoneypb/currency/v1"
	gomock "github.com/golang/mock/gomock"
)

// MockCurrencyServiceClient is a mock of CurrencyServiceClient interface.
type MockCurrencyServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockCurrencyServiceClientMockRecorder
}

// MockCurrencyServiceClientMockRecorder is the mock recorder for MockCurrencyServiceClient.
type MockCurrencyServiceClientMockRecorder struct {
	mock *MockCurrencyServiceClient
}

// NewMockCurrencyServiceClient creates a new mock instance.
func NewMockCurrencyServiceClient(ctrl *gomock.Controller) *MockCurrencyServiceClient {
	mock := &MockCurrencyServiceClient{ctrl: ctrl}
	mock.recorder = &MockCurrencyServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCurrencyServiceClient) EXPECT() *MockCurrencyServiceClientMockRecorder {
	return m.recorder
}

// CreateCurrency mocks base method.
func (m *MockCurrencyServiceClient) CreateCurrency(arg0 context.Context, arg1 *connect.Request[v1.CreateCurrencyRequest]) (*connect.Response[v1.CreateCurrencyResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCurrency", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.CreateCurrencyResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCurrency indicates an expected call of CreateCurrency.
func (mr *MockCurrencyServiceClientMockRecorder) CreateCurrency(arg0, arg1 interface{}) *CurrencyServiceClientCreateCurrencyCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCurrency", reflect.TypeOf((*MockCurrencyServiceClient)(nil).CreateCurrency), arg0, arg1)
	return &CurrencyServiceClientCreateCurrencyCall{Call: call}
}

// CurrencyServiceClientCreateCurrencyCall wrap *gomock.Call
type CurrencyServiceClientCreateCurrencyCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CurrencyServiceClientCreateCurrencyCall) Return(arg0 *connect.Response[v1.CreateCurrencyResponse], arg1 error) *CurrencyServiceClientCreateCurrencyCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CurrencyServiceClientCreateCurrencyCall) Do(f func(context.Context, *connect.Request[v1.CreateCurrencyRequest]) (*connect.Response[v1.CreateCurrencyResponse], error)) *CurrencyServiceClientCreateCurrencyCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CurrencyServiceClientCreateCurrencyCall) DoAndReturn(f func(context.Context, *connect.Request[v1.CreateCurrencyRequest]) (*connect.Response[v1.CreateCurrencyResponse], error)) *CurrencyServiceClientCreateCurrencyCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DeleteCurrency mocks base method.
func (m *MockCurrencyServiceClient) DeleteCurrency(arg0 context.Context, arg1 *connect.Request[v1.DeleteCurrencyRequest]) (*connect.Response[v1.DeleteCurrencyResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCurrency", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.DeleteCurrencyResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteCurrency indicates an expected call of DeleteCurrency.
func (mr *MockCurrencyServiceClientMockRecorder) DeleteCurrency(arg0, arg1 interface{}) *CurrencyServiceClientDeleteCurrencyCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCurrency", reflect.TypeOf((*MockCurrencyServiceClient)(nil).DeleteCurrency), arg0, arg1)
	return &CurrencyServiceClientDeleteCurrencyCall{Call: call}
}

// CurrencyServiceClientDeleteCurrencyCall wrap *gomock.Call
type CurrencyServiceClientDeleteCurrencyCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CurrencyServiceClientDeleteCurrencyCall) Return(arg0 *connect.Response[v1.DeleteCurrencyResponse], arg1 error) *CurrencyServiceClientDeleteCurrencyCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CurrencyServiceClientDeleteCurrencyCall) Do(f func(context.Context, *connect.Request[v1.DeleteCurrencyRequest]) (*connect.Response[v1.DeleteCurrencyResponse], error)) *CurrencyServiceClientDeleteCurrencyCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CurrencyServiceClientDeleteCurrencyCall) DoAndReturn(f func(context.Context, *connect.Request[v1.DeleteCurrencyRequest]) (*connect.Response[v1.DeleteCurrencyResponse], error)) *CurrencyServiceClientDeleteCurrencyCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Exchange mocks base method.
func (m *MockCurrencyServiceClient) Exchange(arg0 context.Context, arg1 *connect.Request[v1.ExchangeRequest]) (*connect.Response[v1.ExchangeResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exchange", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.ExchangeResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exchange indicates an expected call of Exchange.
func (mr *MockCurrencyServiceClientMockRecorder) Exchange(arg0, arg1 interface{}) *CurrencyServiceClientExchangeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exchange", reflect.TypeOf((*MockCurrencyServiceClient)(nil).Exchange), arg0, arg1)
	return &CurrencyServiceClientExchangeCall{Call: call}
}

// CurrencyServiceClientExchangeCall wrap *gomock.Call
type CurrencyServiceClientExchangeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CurrencyServiceClientExchangeCall) Return(arg0 *connect.Response[v1.ExchangeResponse], arg1 error) *CurrencyServiceClientExchangeCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CurrencyServiceClientExchangeCall) Do(f func(context.Context, *connect.Request[v1.ExchangeRequest]) (*connect.Response[v1.ExchangeResponse], error)) *CurrencyServiceClientExchangeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CurrencyServiceClientExchangeCall) DoAndReturn(f func(context.Context, *connect.Request[v1.ExchangeRequest]) (*connect.Response[v1.ExchangeResponse], error)) *CurrencyServiceClientExchangeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetCurrencies mocks base method.
func (m *MockCurrencyServiceClient) GetCurrencies(arg0 context.Context, arg1 *connect.Request[v1.GetCurrenciesRequest]) (*connect.Response[v1.GetCurrenciesResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrencies", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.GetCurrenciesResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrencies indicates an expected call of GetCurrencies.
func (mr *MockCurrencyServiceClientMockRecorder) GetCurrencies(arg0, arg1 interface{}) *CurrencyServiceClientGetCurrenciesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrencies", reflect.TypeOf((*MockCurrencyServiceClient)(nil).GetCurrencies), arg0, arg1)
	return &CurrencyServiceClientGetCurrenciesCall{Call: call}
}

// CurrencyServiceClientGetCurrenciesCall wrap *gomock.Call
type CurrencyServiceClientGetCurrenciesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CurrencyServiceClientGetCurrenciesCall) Return(arg0 *connect.Response[v1.GetCurrenciesResponse], arg1 error) *CurrencyServiceClientGetCurrenciesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CurrencyServiceClientGetCurrenciesCall) Do(f func(context.Context, *connect.Request[v1.GetCurrenciesRequest]) (*connect.Response[v1.GetCurrenciesResponse], error)) *CurrencyServiceClientGetCurrenciesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CurrencyServiceClientGetCurrenciesCall) DoAndReturn(f func(context.Context, *connect.Request[v1.GetCurrenciesRequest]) (*connect.Response[v1.GetCurrenciesResponse], error)) *CurrencyServiceClientGetCurrenciesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdateCurrency mocks base method.
func (m *MockCurrencyServiceClient) UpdateCurrency(arg0 context.Context, arg1 *connect.Request[v1.UpdateCurrencyRequest]) (*connect.Response[v1.UpdateCurrencyResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCurrency", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.UpdateCurrencyResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCurrency indicates an expected call of UpdateCurrency.
func (mr *MockCurrencyServiceClientMockRecorder) UpdateCurrency(arg0, arg1 interface{}) *CurrencyServiceClientUpdateCurrencyCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCurrency", reflect.TypeOf((*MockCurrencyServiceClient)(nil).UpdateCurrency), arg0, arg1)
	return &CurrencyServiceClientUpdateCurrencyCall{Call: call}
}

// CurrencyServiceClientUpdateCurrencyCall wrap *gomock.Call
type CurrencyServiceClientUpdateCurrencyCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CurrencyServiceClientUpdateCurrencyCall) Return(arg0 *connect.Response[v1.UpdateCurrencyResponse], arg1 error) *CurrencyServiceClientUpdateCurrencyCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CurrencyServiceClientUpdateCurrencyCall) Do(f func(context.Context, *connect.Request[v1.UpdateCurrencyRequest]) (*connect.Response[v1.UpdateCurrencyResponse], error)) *CurrencyServiceClientUpdateCurrencyCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CurrencyServiceClientUpdateCurrencyCall) DoAndReturn(f func(context.Context, *connect.Request[v1.UpdateCurrencyRequest]) (*connect.Response[v1.UpdateCurrencyResponse], error)) *CurrencyServiceClientUpdateCurrencyCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockCurrencyServiceHandler is a mock of CurrencyServiceHandler interface.
type MockCurrencyServiceHandler struct {
	ctrl     *gomock.Controller
	recorder *MockCurrencyServiceHandlerMockRecorder
}

// MockCurrencyServiceHandlerMockRecorder is the mock recorder for MockCurrencyServiceHandler.
type MockCurrencyServiceHandlerMockRecorder struct {
	mock *MockCurrencyServiceHandler
}

// NewMockCurrencyServiceHandler creates a new mock instance.
func NewMockCurrencyServiceHandler(ctrl *gomock.Controller) *MockCurrencyServiceHandler {
	mock := &MockCurrencyServiceHandler{ctrl: ctrl}
	mock.recorder = &MockCurrencyServiceHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCurrencyServiceHandler) EXPECT() *MockCurrencyServiceHandlerMockRecorder {
	return m.recorder
}

// CreateCurrency mocks base method.
func (m *MockCurrencyServiceHandler) CreateCurrency(arg0 context.Context, arg1 *connect.Request[v1.CreateCurrencyRequest]) (*connect.Response[v1.CreateCurrencyResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCurrency", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.CreateCurrencyResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCurrency indicates an expected call of CreateCurrency.
func (mr *MockCurrencyServiceHandlerMockRecorder) CreateCurrency(arg0, arg1 interface{}) *CurrencyServiceHandlerCreateCurrencyCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCurrency", reflect.TypeOf((*MockCurrencyServiceHandler)(nil).CreateCurrency), arg0, arg1)
	return &CurrencyServiceHandlerCreateCurrencyCall{Call: call}
}

// CurrencyServiceHandlerCreateCurrencyCall wrap *gomock.Call
type CurrencyServiceHandlerCreateCurrencyCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CurrencyServiceHandlerCreateCurrencyCall) Return(arg0 *connect.Response[v1.CreateCurrencyResponse], arg1 error) *CurrencyServiceHandlerCreateCurrencyCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CurrencyServiceHandlerCreateCurrencyCall) Do(f func(context.Context, *connect.Request[v1.CreateCurrencyRequest]) (*connect.Response[v1.CreateCurrencyResponse], error)) *CurrencyServiceHandlerCreateCurrencyCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CurrencyServiceHandlerCreateCurrencyCall) DoAndReturn(f func(context.Context, *connect.Request[v1.CreateCurrencyRequest]) (*connect.Response[v1.CreateCurrencyResponse], error)) *CurrencyServiceHandlerCreateCurrencyCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DeleteCurrency mocks base method.
func (m *MockCurrencyServiceHandler) DeleteCurrency(arg0 context.Context, arg1 *connect.Request[v1.DeleteCurrencyRequest]) (*connect.Response[v1.DeleteCurrencyResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCurrency", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.DeleteCurrencyResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteCurrency indicates an expected call of DeleteCurrency.
func (mr *MockCurrencyServiceHandlerMockRecorder) DeleteCurrency(arg0, arg1 interface{}) *CurrencyServiceHandlerDeleteCurrencyCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCurrency", reflect.TypeOf((*MockCurrencyServiceHandler)(nil).DeleteCurrency), arg0, arg1)
	return &CurrencyServiceHandlerDeleteCurrencyCall{Call: call}
}

// CurrencyServiceHandlerDeleteCurrencyCall wrap *gomock.Call
type CurrencyServiceHandlerDeleteCurrencyCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CurrencyServiceHandlerDeleteCurrencyCall) Return(arg0 *connect.Response[v1.DeleteCurrencyResponse], arg1 error) *CurrencyServiceHandlerDeleteCurrencyCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CurrencyServiceHandlerDeleteCurrencyCall) Do(f func(context.Context, *connect.Request[v1.DeleteCurrencyRequest]) (*connect.Response[v1.DeleteCurrencyResponse], error)) *CurrencyServiceHandlerDeleteCurrencyCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CurrencyServiceHandlerDeleteCurrencyCall) DoAndReturn(f func(context.Context, *connect.Request[v1.DeleteCurrencyRequest]) (*connect.Response[v1.DeleteCurrencyResponse], error)) *CurrencyServiceHandlerDeleteCurrencyCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Exchange mocks base method.
func (m *MockCurrencyServiceHandler) Exchange(arg0 context.Context, arg1 *connect.Request[v1.ExchangeRequest]) (*connect.Response[v1.ExchangeResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exchange", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.ExchangeResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exchange indicates an expected call of Exchange.
func (mr *MockCurrencyServiceHandlerMockRecorder) Exchange(arg0, arg1 interface{}) *CurrencyServiceHandlerExchangeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exchange", reflect.TypeOf((*MockCurrencyServiceHandler)(nil).Exchange), arg0, arg1)
	return &CurrencyServiceHandlerExchangeCall{Call: call}
}

// CurrencyServiceHandlerExchangeCall wrap *gomock.Call
type CurrencyServiceHandlerExchangeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CurrencyServiceHandlerExchangeCall) Return(arg0 *connect.Response[v1.ExchangeResponse], arg1 error) *CurrencyServiceHandlerExchangeCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CurrencyServiceHandlerExchangeCall) Do(f func(context.Context, *connect.Request[v1.ExchangeRequest]) (*connect.Response[v1.ExchangeResponse], error)) *CurrencyServiceHandlerExchangeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CurrencyServiceHandlerExchangeCall) DoAndReturn(f func(context.Context, *connect.Request[v1.ExchangeRequest]) (*connect.Response[v1.ExchangeResponse], error)) *CurrencyServiceHandlerExchangeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetCurrencies mocks base method.
func (m *MockCurrencyServiceHandler) GetCurrencies(arg0 context.Context, arg1 *connect.Request[v1.GetCurrenciesRequest]) (*connect.Response[v1.GetCurrenciesResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrencies", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.GetCurrenciesResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrencies indicates an expected call of GetCurrencies.
func (mr *MockCurrencyServiceHandlerMockRecorder) GetCurrencies(arg0, arg1 interface{}) *CurrencyServiceHandlerGetCurrenciesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrencies", reflect.TypeOf((*MockCurrencyServiceHandler)(nil).GetCurrencies), arg0, arg1)
	return &CurrencyServiceHandlerGetCurrenciesCall{Call: call}
}

// CurrencyServiceHandlerGetCurrenciesCall wrap *gomock.Call
type CurrencyServiceHandlerGetCurrenciesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CurrencyServiceHandlerGetCurrenciesCall) Return(arg0 *connect.Response[v1.GetCurrenciesResponse], arg1 error) *CurrencyServiceHandlerGetCurrenciesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CurrencyServiceHandlerGetCurrenciesCall) Do(f func(context.Context, *connect.Request[v1.GetCurrenciesRequest]) (*connect.Response[v1.GetCurrenciesResponse], error)) *CurrencyServiceHandlerGetCurrenciesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CurrencyServiceHandlerGetCurrenciesCall) DoAndReturn(f func(context.Context, *connect.Request[v1.GetCurrenciesRequest]) (*connect.Response[v1.GetCurrenciesResponse], error)) *CurrencyServiceHandlerGetCurrenciesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdateCurrency mocks base method.
func (m *MockCurrencyServiceHandler) UpdateCurrency(arg0 context.Context, arg1 *connect.Request[v1.UpdateCurrencyRequest]) (*connect.Response[v1.UpdateCurrencyResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCurrency", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.UpdateCurrencyResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCurrency indicates an expected call of UpdateCurrency.
func (mr *MockCurrencyServiceHandlerMockRecorder) UpdateCurrency(arg0, arg1 interface{}) *CurrencyServiceHandlerUpdateCurrencyCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCurrency", reflect.TypeOf((*MockCurrencyServiceHandler)(nil).UpdateCurrency), arg0, arg1)
	return &CurrencyServiceHandlerUpdateCurrencyCall{Call: call}
}

// CurrencyServiceHandlerUpdateCurrencyCall wrap *gomock.Call
type CurrencyServiceHandlerUpdateCurrencyCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CurrencyServiceHandlerUpdateCurrencyCall) Return(arg0 *connect.Response[v1.UpdateCurrencyResponse], arg1 error) *CurrencyServiceHandlerUpdateCurrencyCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CurrencyServiceHandlerUpdateCurrencyCall) Do(f func(context.Context, *connect.Request[v1.UpdateCurrencyRequest]) (*connect.Response[v1.UpdateCurrencyResponse], error)) *CurrencyServiceHandlerUpdateCurrencyCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CurrencyServiceHandlerUpdateCurrencyCall) DoAndReturn(f func(context.Context, *connect.Request[v1.UpdateCurrencyRequest]) (*connect.Response[v1.UpdateCurrencyResponse], error)) *CurrencyServiceHandlerUpdateCurrencyCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
