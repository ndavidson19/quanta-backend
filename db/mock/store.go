// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ndavidson19/quanta-backend/db (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	db "github.com/ndavidson19/quanta-backend/db"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateAccount mocks base method.
func (m *MockStore) CreateAccount(arg0 context.Context, arg1 db.CreateAccountParams) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockStoreMockRecorder) CreateAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockStore)(nil).CreateAccount), arg0, arg1)
}

// CreateDeposit mocks base method.
func (m *MockStore) CreateDeposit(arg0 context.Context, arg1 db.CreateDepositParams) (db.Deposit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDeposit", arg0, arg1)
	ret0, _ := ret[0].(db.Deposit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDeposit indicates an expected call of CreateDeposit.
func (mr *MockStoreMockRecorder) CreateDeposit(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDeposit", reflect.TypeOf((*MockStore)(nil).CreateDeposit), arg0, arg1)
}

// CreateLogs mocks base method.
func (m *MockStore) CreateLogs(arg0 context.Context, arg1 sql.NullString) (db.AuditLog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLogs", arg0, arg1)
	ret0, _ := ret[0].(db.AuditLog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLogs indicates an expected call of CreateLogs.
func (mr *MockStoreMockRecorder) CreateLogs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLogs", reflect.TypeOf((*MockStore)(nil).CreateLogs), arg0, arg1)
}

// CreateTrade mocks base method.
func (m *MockStore) CreateTrade(arg0 context.Context, arg1 db.CreateTradeParams) (db.Trade, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTrade", arg0, arg1)
	ret0, _ := ret[0].(db.Trade)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTrade indicates an expected call of CreateTrade.
func (mr *MockStoreMockRecorder) CreateTrade(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTrade", reflect.TypeOf((*MockStore)(nil).CreateTrade), arg0, arg1)
}

// CreateTx mocks base method.
func (m *MockStore) CreateTx(arg0 context.Context, arg1 db.CreateTxParams) (db.CreateTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTx", arg0, arg1)
	ret0, _ := ret[0].(db.CreateTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTx indicates an expected call of CreateTx.
func (mr *MockStoreMockRecorder) CreateTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTx", reflect.TypeOf((*MockStore)(nil).CreateTx), arg0, arg1)
}

// DeleteAccount mocks base method.
func (m *MockStore) DeleteAccount(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccount", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccount indicates an expected call of DeleteAccount.
func (mr *MockStoreMockRecorder) DeleteAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccount", reflect.TypeOf((*MockStore)(nil).DeleteAccount), arg0, arg1)
}

// GetAccount mocks base method.
func (m *MockStore) GetAccount(arg0 context.Context, arg1 int64) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockStoreMockRecorder) GetAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockStore)(nil).GetAccount), arg0, arg1)
}

// GetAccountForUpdate mocks base method.
func (m *MockStore) GetAccountForUpdate(arg0 context.Context, arg1 int64) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountForUpdate", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountForUpdate indicates an expected call of GetAccountForUpdate.
func (mr *MockStoreMockRecorder) GetAccountForUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountForUpdate", reflect.TypeOf((*MockStore)(nil).GetAccountForUpdate), arg0, arg1)
}

// GetTrade mocks base method.
func (m *MockStore) GetTrade(arg0 context.Context, arg1 int64) (db.Trade, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTrade", arg0, arg1)
	ret0, _ := ret[0].(db.Trade)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTrade indicates an expected call of GetTrade.
func (mr *MockStoreMockRecorder) GetTrade(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTrade", reflect.TypeOf((*MockStore)(nil).GetTrade), arg0, arg1)
}

// ListAccounts mocks base method.
func (m *MockStore) ListAccounts(arg0 context.Context, arg1 db.ListAccountsParams) ([]db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccounts", arg0, arg1)
	ret0, _ := ret[0].([]db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccounts indicates an expected call of ListAccounts.
func (mr *MockStoreMockRecorder) ListAccounts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccounts", reflect.TypeOf((*MockStore)(nil).ListAccounts), arg0, arg1)
}

// ListLogs mocks base method.
func (m *MockStore) ListLogs(arg0 context.Context, arg1 db.ListLogsParams) ([]db.AuditLog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLogs", arg0, arg1)
	ret0, _ := ret[0].([]db.AuditLog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListLogs indicates an expected call of ListLogs.
func (mr *MockStoreMockRecorder) ListLogs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLogs", reflect.TypeOf((*MockStore)(nil).ListLogs), arg0, arg1)
}

// ListTrades mocks base method.
func (m *MockStore) ListTrades(arg0 context.Context, arg1 db.ListTradesParams) ([]db.Trade, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTrades", arg0, arg1)
	ret0, _ := ret[0].([]db.Trade)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTrades indicates an expected call of ListTrades.
func (mr *MockStoreMockRecorder) ListTrades(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTrades", reflect.TypeOf((*MockStore)(nil).ListTrades), arg0, arg1)
}

// UpdateAccount mocks base method.
func (m *MockStore) UpdateAccount(arg0 context.Context, arg1 db.UpdateAccountParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccount", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAccount indicates an expected call of UpdateAccount.
func (mr *MockStoreMockRecorder) UpdateAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccount", reflect.TypeOf((*MockStore)(nil).UpdateAccount), arg0, arg1)
}
