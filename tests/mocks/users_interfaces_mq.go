// Code generated by MockGen. DO NOT EDIT.
// Source: domain/users/interfaces.go

// Package mock is a generated GoMock package.
package mock

import (
	users "bux-wallet/domain/users"
	reflect "reflect"
	time "time"

	buxmodels "github.com/BuxOrg/bux-models"
	transports "github.com/BuxOrg/go-buxclient/transports"
	gomock "github.com/golang/mock/gomock"
	bip32 "github.com/libsv/go-bk/bip32"
	datastore "github.com/mrz1836/go-datastore"
)

// MockAccKey is a mock of AccKey interface.
type MockAccKey struct {
        ctrl     *gomock.Controller
        recorder *MockAccKeyMockRecorder
}

// MockAccKeyMockRecorder is the mock recorder for MockAccKey.
type MockAccKeyMockRecorder struct {
        mock *MockAccKey
}

// NewMockAccKey creates a new mock instance.
func NewMockAccKey(ctrl *gomock.Controller) *MockAccKey {
        mock := &MockAccKey{ctrl: ctrl}
        mock.recorder = &MockAccKeyMockRecorder{mock}
        return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccKey) EXPECT() *MockAccKeyMockRecorder {
        return m.recorder
}

// GetAccessKey mocks base method.
func (m *MockAccKey) GetAccessKey() string {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetAccessKey")
        ret0, _ := ret[0].(string)
        return ret0
}

// GetAccessKey indicates an expected call of GetAccessKey.
func (mr *MockAccKeyMockRecorder) GetAccessKey() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessKey", reflect.TypeOf((*MockAccKey)(nil).GetAccessKey))
}

// GetAccessKeyId mocks base method.
func (m *MockAccKey) GetAccessKeyId() string {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetAccessKeyId")
        ret0, _ := ret[0].(string)
        return ret0
}

// GetAccessKeyId indicates an expected call of GetAccessKeyId.
func (mr *MockAccKeyMockRecorder) GetAccessKeyId() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessKeyId", reflect.TypeOf((*MockAccKey)(nil).GetAccessKeyId))
}

// MockPubKey is a mock of PubKey interface.
type MockPubKey struct {
        ctrl     *gomock.Controller
        recorder *MockPubKeyMockRecorder
}

// MockPubKeyMockRecorder is the mock recorder for MockPubKey.
type MockPubKeyMockRecorder struct {
        mock *MockPubKey
}

// NewMockPubKey creates a new mock instance.
func NewMockPubKey(ctrl *gomock.Controller) *MockPubKey {
        mock := &MockPubKey{ctrl: ctrl}
        mock.recorder = &MockPubKeyMockRecorder{mock}
        return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPubKey) EXPECT() *MockPubKeyMockRecorder {
        return m.recorder
}

// GetCurrentBalance mocks base method.
func (m *MockPubKey) GetCurrentBalance() uint64 {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetCurrentBalance")
        ret0, _ := ret[0].(uint64)
        return ret0
}

// GetCurrentBalance indicates an expected call of GetCurrentBalance.
func (mr *MockPubKeyMockRecorder) GetCurrentBalance() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentBalance", reflect.TypeOf((*MockPubKey)(nil).GetCurrentBalance))
}

// GetId mocks base method.
func (m *MockPubKey) GetId() string {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetId")
        ret0, _ := ret[0].(string)
        return ret0
}

// GetId indicates an expected call of GetId.
func (mr *MockPubKeyMockRecorder) GetId() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetId", reflect.TypeOf((*MockPubKey)(nil).GetId))
}

// GetXPub mocks base method.
func (m *MockPubKey) GetXPub() string {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetXPub")
        ret0, _ := ret[0].(string)
        return ret0
}

// GetXPub indicates an expected call of GetXPub.
func (mr *MockPubKeyMockRecorder) GetXPub() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetXPub", reflect.TypeOf((*MockPubKey)(nil).GetXPub))
}

// MockTransaction is a mock of Transaction interface.
type MockTransaction struct {
        ctrl     *gomock.Controller
        recorder *MockTransactionMockRecorder
}

// MockTransactionMockRecorder is the mock recorder for MockTransaction.
type MockTransactionMockRecorder struct {
        mock *MockTransaction
}

// NewMockTransaction creates a new mock instance.
func NewMockTransaction(ctrl *gomock.Controller) *MockTransaction {
        mock := &MockTransaction{ctrl: ctrl}
        mock.recorder = &MockTransactionMockRecorder{mock}
        return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransaction) EXPECT() *MockTransactionMockRecorder {
        return m.recorder
}

// GetTransactionCreatedDate mocks base method.
func (m *MockTransaction) GetTransactionCreatedDate() time.Time {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetTransactionCreatedDate")
        ret0, _ := ret[0].(time.Time)
        return ret0
}

// GetTransactionCreatedDate indicates an expected call of GetTransactionCreatedDate.
func (mr *MockTransactionMockRecorder) GetTransactionCreatedDate() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionCreatedDate", reflect.TypeOf((*MockTransaction)(nil).GetTransactionCreatedDate))
}

// GetTransactionDirection mocks base method.
func (m *MockTransaction) GetTransactionDirection() string {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetTransactionDirection")
        ret0, _ := ret[0].(string)
        return ret0
}

// GetTransactionDirection indicates an expected call of GetTransactionDirection.
func (mr *MockTransactionMockRecorder) GetTransactionDirection() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionDirection", reflect.TypeOf((*MockTransaction)(nil).GetTransactionDirection))
}

// GetTransactionFee mocks base method.
func (m *MockTransaction) GetTransactionFee() uint64 {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetTransactionFee")
        ret0, _ := ret[0].(uint64)
        return ret0
}

// GetTransactionFee indicates an expected call of GetTransactionFee.
func (mr *MockTransactionMockRecorder) GetTransactionFee() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionFee", reflect.TypeOf((*MockTransaction)(nil).GetTransactionFee))
}

// GetTransactionId mocks base method.
func (m *MockTransaction) GetTransactionId() string {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetTransactionId")
        ret0, _ := ret[0].(string)
        return ret0
}

// GetTransactionId indicates an expected call of GetTransactionId.
func (mr *MockTransactionMockRecorder) GetTransactionId() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionId", reflect.TypeOf((*MockTransaction)(nil).GetTransactionId))
}

// GetTransactionReceiver mocks base method.
func (m *MockTransaction) GetTransactionReceiver() string {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetTransactionReceiver")
        ret0, _ := ret[0].(string)
        return ret0
}

// GetTransactionReceiver indicates an expected call of GetTransactionReceiver.
func (mr *MockTransactionMockRecorder) GetTransactionReceiver() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionReceiver", reflect.TypeOf((*MockTransaction)(nil).GetTransactionReceiver))
}

// GetTransactionSender mocks base method.
func (m *MockTransaction) GetTransactionSender() string {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetTransactionSender")
        ret0, _ := ret[0].(string)
        return ret0
}

// GetTransactionSender indicates an expected call of GetTransactionSender.
func (mr *MockTransactionMockRecorder) GetTransactionSender() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionSender", reflect.TypeOf((*MockTransaction)(nil).GetTransactionSender))
}

// GetTransactionStatus mocks base method.
func (m *MockTransaction) GetTransactionStatus() string {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetTransactionStatus")
        ret0, _ := ret[0].(string)
        return ret0
}

// GetTransactionStatus indicates an expected call of GetTransactionStatus.
func (mr *MockTransactionMockRecorder) GetTransactionStatus() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionStatus", reflect.TypeOf((*MockTransaction)(nil).GetTransactionStatus))
}

// GetTransactionTotalValue mocks base method.
func (m *MockTransaction) GetTransactionTotalValue() uint64 {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetTransactionTotalValue")
        ret0, _ := ret[0].(uint64)
        return ret0
}

// GetTransactionTotalValue indicates an expected call of GetTransactionTotalValue.
func (mr *MockTransactionMockRecorder) GetTransactionTotalValue() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionTotalValue", reflect.TypeOf((*MockTransaction)(nil).GetTransactionTotalValue))
}

// MockFullTransaction is a mock of FullTransaction interface.
type MockFullTransaction struct {
        ctrl     *gomock.Controller
        recorder *MockFullTransactionMockRecorder
}

// MockFullTransactionMockRecorder is the mock recorder for MockFullTransaction.
type MockFullTransactionMockRecorder struct {
        mock *MockFullTransaction
}

// NewMockFullTransaction creates a new mock instance.
func NewMockFullTransaction(ctrl *gomock.Controller) *MockFullTransaction {
        mock := &MockFullTransaction{ctrl: ctrl}
        mock.recorder = &MockFullTransactionMockRecorder{mock}
        return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFullTransaction) EXPECT() *MockFullTransactionMockRecorder {
        return m.recorder
}

// GetTransactionBlockHash mocks base method.
func (m *MockFullTransaction) GetTransactionBlockHash() string {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetTransactionBlockHash")
        ret0, _ := ret[0].(string)
        return ret0
}

// GetTransactionBlockHash indicates an expected call of GetTransactionBlockHash.
func (mr *MockFullTransactionMockRecorder) GetTransactionBlockHash() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionBlockHash", reflect.TypeOf((*MockFullTransaction)(nil).GetTransactionBlockHash))
}

// GetTransactionBlockHeight mocks base method.
func (m *MockFullTransaction) GetTransactionBlockHeight() uint64 {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetTransactionBlockHeight")
        ret0, _ := ret[0].(uint64)
        return ret0
}

// GetTransactionBlockHeight indicates an expected call of GetTransactionBlockHeight.
func (mr *MockFullTransactionMockRecorder) GetTransactionBlockHeight() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionBlockHeight", reflect.TypeOf((*MockFullTransaction)(nil).GetTransactionBlockHeight))
}

// GetTransactionCreatedDate mocks base method.
func (m *MockFullTransaction) GetTransactionCreatedDate() time.Time {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetTransactionCreatedDate")
        ret0, _ := ret[0].(time.Time)
        return ret0
}

// GetTransactionCreatedDate indicates an expected call of GetTransactionCreatedDate.
func (mr *MockFullTransactionMockRecorder) GetTransactionCreatedDate() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionCreatedDate", reflect.TypeOf((*MockFullTransaction)(nil).GetTransactionCreatedDate))
}

// GetTransactionDirection mocks base method.
func (m *MockFullTransaction) GetTransactionDirection() string {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetTransactionDirection")
        ret0, _ := ret[0].(string)
        return ret0
}

// GetTransactionDirection indicates an expected call of GetTransactionDirection.
func (mr *MockFullTransactionMockRecorder) GetTransactionDirection() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionDirection", reflect.TypeOf((*MockFullTransaction)(nil).GetTransactionDirection))
}

// GetTransactionFee mocks base method.
func (m *MockFullTransaction) GetTransactionFee() uint64 {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetTransactionFee")
        ret0, _ := ret[0].(uint64)
        return ret0
}

// GetTransactionFee indicates an expected call of GetTransactionFee.
func (mr *MockFullTransactionMockRecorder) GetTransactionFee() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionFee", reflect.TypeOf((*MockFullTransaction)(nil).GetTransactionFee))
}

// GetTransactionId mocks base method.
func (m *MockFullTransaction) GetTransactionId() string {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetTransactionId")
        ret0, _ := ret[0].(string)
        return ret0
}

// GetTransactionId indicates an expected call of GetTransactionId.
func (mr *MockFullTransactionMockRecorder) GetTransactionId() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionId", reflect.TypeOf((*MockFullTransaction)(nil).GetTransactionId))
}

// GetTransactionNumberOfInputs mocks base method.
func (m *MockFullTransaction) GetTransactionNumberOfInputs() uint32 {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetTransactionNumberOfInputs")
        ret0, _ := ret[0].(uint32)
        return ret0
}

// GetTransactionNumberOfInputs indicates an expected call of GetTransactionNumberOfInputs.
func (mr *MockFullTransactionMockRecorder) GetTransactionNumberOfInputs() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionNumberOfInputs", reflect.TypeOf((*MockFullTransaction)(nil).GetTransactionNumberOfInputs))
}

// GetTransactionNumberOfOutputs mocks base method.
func (m *MockFullTransaction) GetTransactionNumberOfOutputs() uint32 {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetTransactionNumberOfOutputs")
        ret0, _ := ret[0].(uint32)
        return ret0
}

// GetTransactionNumberOfOutputs indicates an expected call of GetTransactionNumberOfOutputs.
func (mr *MockFullTransactionMockRecorder) GetTransactionNumberOfOutputs() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionNumberOfOutputs", reflect.TypeOf((*MockFullTransaction)(nil).GetTransactionNumberOfOutputs))
}

// GetTransactionReceiver mocks base method.
func (m *MockFullTransaction) GetTransactionReceiver() string {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetTransactionReceiver")
        ret0, _ := ret[0].(string)
        return ret0
}

// GetTransactionReceiver indicates an expected call of GetTransactionReceiver.
func (mr *MockFullTransactionMockRecorder) GetTransactionReceiver() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionReceiver", reflect.TypeOf((*MockFullTransaction)(nil).GetTransactionReceiver))
}

// GetTransactionSender mocks base method.
func (m *MockFullTransaction) GetTransactionSender() string {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetTransactionSender")
        ret0, _ := ret[0].(string)
        return ret0
}

// GetTransactionSender indicates an expected call of GetTransactionSender.
func (mr *MockFullTransactionMockRecorder) GetTransactionSender() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionSender", reflect.TypeOf((*MockFullTransaction)(nil).GetTransactionSender))
}

// GetTransactionStatus mocks base method.
func (m *MockFullTransaction) GetTransactionStatus() string {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetTransactionStatus")
        ret0, _ := ret[0].(string)
        return ret0
}

// GetTransactionStatus indicates an expected call of GetTransactionStatus.
func (mr *MockFullTransactionMockRecorder) GetTransactionStatus() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionStatus", reflect.TypeOf((*MockFullTransaction)(nil).GetTransactionStatus))
}

// GetTransactionTotalValue mocks base method.
func (m *MockFullTransaction) GetTransactionTotalValue() uint64 {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetTransactionTotalValue")
        ret0, _ := ret[0].(uint64)
        return ret0
}

// GetTransactionTotalValue indicates an expected call of GetTransactionTotalValue.
func (mr *MockFullTransactionMockRecorder) GetTransactionTotalValue() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionTotalValue", reflect.TypeOf((*MockFullTransaction)(nil).GetTransactionTotalValue))
}

// MockDraftTransaction is a mock of DraftTransaction interface.
type MockDraftTransaction struct {
        ctrl     *gomock.Controller
        recorder *MockDraftTransactionMockRecorder
}

// MockDraftTransactionMockRecorder is the mock recorder for MockDraftTransaction.
type MockDraftTransactionMockRecorder struct {
        mock *MockDraftTransaction
}

// NewMockDraftTransaction creates a new mock instance.
func NewMockDraftTransaction(ctrl *gomock.Controller) *MockDraftTransaction {
        mock := &MockDraftTransaction{ctrl: ctrl}
        mock.recorder = &MockDraftTransactionMockRecorder{mock}
        return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDraftTransaction) EXPECT() *MockDraftTransactionMockRecorder {
        return m.recorder
}

// GetDraftTransactionHex mocks base method.
func (m *MockDraftTransaction) GetDraftTransactionHex() string {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetDraftTransactionHex")
        ret0, _ := ret[0].(string)
        return ret0
}

// GetDraftTransactionHex indicates an expected call of GetDraftTransactionHex.
func (mr *MockDraftTransactionMockRecorder) GetDraftTransactionHex() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDraftTransactionHex", reflect.TypeOf((*MockDraftTransaction)(nil).GetDraftTransactionHex))
}

// GetDraftTransactionId mocks base method.
func (m *MockDraftTransaction) GetDraftTransactionId() string {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetDraftTransactionId")
        ret0, _ := ret[0].(string)
        return ret0
}

// GetDraftTransactionId indicates an expected call of GetDraftTransactionId.
func (mr *MockDraftTransactionMockRecorder) GetDraftTransactionId() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDraftTransactionId", reflect.TypeOf((*MockDraftTransaction)(nil).GetDraftTransactionId))
}

// MockUserBuxClient is a mock of UserBuxClient interface.
type MockUserBuxClient struct {
        ctrl     *gomock.Controller
        recorder *MockUserBuxClientMockRecorder
}

// MockUserBuxClientMockRecorder is the mock recorder for MockUserBuxClient.
type MockUserBuxClientMockRecorder struct {
        mock *MockUserBuxClient
}

// NewMockUserBuxClient creates a new mock instance.
func NewMockUserBuxClient(ctrl *gomock.Controller) *MockUserBuxClient {
        mock := &MockUserBuxClient{ctrl: ctrl}
        mock.recorder = &MockUserBuxClientMockRecorder{mock}
        return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserBuxClient) EXPECT() *MockUserBuxClientMockRecorder {
        return m.recorder
}

// CreateAccessKey mocks base method.
func (m *MockUserBuxClient) CreateAccessKey() (users.AccKey, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "CreateAccessKey")
        ret0, _ := ret[0].(users.AccKey)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// CreateAccessKey indicates an expected call of CreateAccessKey.
func (mr *MockUserBuxClientMockRecorder) CreateAccessKey() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccessKey", reflect.TypeOf((*MockUserBuxClient)(nil).CreateAccessKey))
}

// CreateAndFinalizeTransaction mocks base method.
func (m *MockUserBuxClient) CreateAndFinalizeTransaction(recipients []*transports.Recipients, metadata *buxmodels.Metadata) (users.DraftTransaction, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "CreateAndFinalizeTransaction", recipients, metadata)
        ret0, _ := ret[0].(users.DraftTransaction)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// CreateAndFinalizeTransaction indicates an expected call of CreateAndFinalizeTransaction.
func (mr *MockUserBuxClientMockRecorder) CreateAndFinalizeTransaction(recipients, metadata interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAndFinalizeTransaction", reflect.TypeOf((*MockUserBuxClient)(nil).CreateAndFinalizeTransaction), recipients, metadata)
}

// GetAccessKey mocks base method.
func (m *MockUserBuxClient) GetAccessKey(accessKeyId string) (users.AccKey, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetAccessKey", accessKeyId)
        ret0, _ := ret[0].(users.AccKey)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetAccessKey indicates an expected call of GetAccessKey.
func (mr *MockUserBuxClientMockRecorder) GetAccessKey(accessKeyId interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessKey", reflect.TypeOf((*MockUserBuxClient)(nil).GetAccessKey), accessKeyId)
}

// GetTransaction mocks base method.
func (m *MockUserBuxClient) GetTransaction(transactionId, userPaymail string) (users.FullTransaction, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetTransaction", transactionId, userPaymail)
        ret0, _ := ret[0].(users.FullTransaction)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetTransaction indicates an expected call of GetTransaction.
func (mr *MockUserBuxClientMockRecorder) GetTransaction(transactionId, userPaymail interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransaction", reflect.TypeOf((*MockUserBuxClient)(nil).GetTransaction), transactionId, userPaymail)
}

// GetTransactions mocks base method.
func (m *MockUserBuxClient) GetTransactions(queryParam datastore.QueryParams, userPaymail string) ([]users.Transaction, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetTransactions", queryParam, userPaymail)
        ret0, _ := ret[0].([]users.Transaction)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetTransactions indicates an expected call of GetTransactions.
func (mr *MockUserBuxClientMockRecorder) GetTransactions(queryParam, userPaymail interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactions", reflect.TypeOf((*MockUserBuxClient)(nil).GetTransactions), queryParam, userPaymail)
}

// GetTransactionsCount mocks base method.
func (m *MockUserBuxClient) GetTransactionsCount() (int64, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetTransactionsCount")
        ret0, _ := ret[0].(int64)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetTransactionsCount indicates an expected call of GetTransactionsCount.
func (mr *MockUserBuxClientMockRecorder) GetTransactionsCount() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionsCount", reflect.TypeOf((*MockUserBuxClient)(nil).GetTransactionsCount))
}

// GetXPub mocks base method.
func (m *MockUserBuxClient) GetXPub() (users.PubKey, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetXPub")
        ret0, _ := ret[0].(users.PubKey)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetXPub indicates an expected call of GetXPub.
func (mr *MockUserBuxClientMockRecorder) GetXPub() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetXPub", reflect.TypeOf((*MockUserBuxClient)(nil).GetXPub))
}

// RecordTransaction mocks base method.
func (m *MockUserBuxClient) RecordTransaction(hex, draftTxId string, metadata *buxmodels.Metadata) {
        m.ctrl.T.Helper()
        m.ctrl.Call(m, "RecordTransaction", hex, draftTxId, metadata)
}

// RecordTransaction indicates an expected call of RecordTransaction.
func (mr *MockUserBuxClientMockRecorder) RecordTransaction(hex, draftTxId, metadata interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordTransaction", reflect.TypeOf((*MockUserBuxClient)(nil).RecordTransaction), hex, draftTxId, metadata)
}

// RevokeAccessKey mocks base method.
func (m *MockUserBuxClient) RevokeAccessKey(accessKeyId string) (users.AccKey, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "RevokeAccessKey", accessKeyId)
        ret0, _ := ret[0].(users.AccKey)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// RevokeAccessKey indicates an expected call of RevokeAccessKey.
func (mr *MockUserBuxClientMockRecorder) RevokeAccessKey(accessKeyId interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeAccessKey", reflect.TypeOf((*MockUserBuxClient)(nil).RevokeAccessKey), accessKeyId)
}

// SendToRecipients mocks base method.
func (m *MockUserBuxClient) SendToRecipients(recipients []*transports.Recipients, senderPaymail string) (users.Transaction, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "SendToRecipients", recipients, senderPaymail)
        ret0, _ := ret[0].(users.Transaction)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// SendToRecipients indicates an expected call of SendToRecipients.
func (mr *MockUserBuxClientMockRecorder) SendToRecipients(recipients, senderPaymail interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendToRecipients", reflect.TypeOf((*MockUserBuxClient)(nil).SendToRecipients), recipients, senderPaymail)
}

// MockAdmBuxClient is a mock of AdmBuxClient interface.
type MockAdmBuxClient struct {
        ctrl     *gomock.Controller
        recorder *MockAdmBuxClientMockRecorder
}

// MockAdmBuxClientMockRecorder is the mock recorder for MockAdmBuxClient.
type MockAdmBuxClientMockRecorder struct {
        mock *MockAdmBuxClient
}

// NewMockAdmBuxClient creates a new mock instance.
func NewMockAdmBuxClient(ctrl *gomock.Controller) *MockAdmBuxClient {
        mock := &MockAdmBuxClient{ctrl: ctrl}
        mock.recorder = &MockAdmBuxClientMockRecorder{mock}
        return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdmBuxClient) EXPECT() *MockAdmBuxClientMockRecorder {
        return m.recorder
}

// RegisterPaymail mocks base method.
func (m *MockAdmBuxClient) RegisterPaymail(alias, xpub string) (string, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "RegisterPaymail", alias, xpub)
        ret0, _ := ret[0].(string)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// RegisterPaymail indicates an expected call of RegisterPaymail.
func (mr *MockAdmBuxClientMockRecorder) RegisterPaymail(alias, xpub interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterPaymail", reflect.TypeOf((*MockAdmBuxClient)(nil).RegisterPaymail), alias, xpub)
}

// RegisterXpub mocks base method.
func (m *MockAdmBuxClient) RegisterXpub(xpriv *bip32.ExtendedKey) (string, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "RegisterXpub", xpriv)
        ret0, _ := ret[0].(string)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// RegisterXpub indicates an expected call of RegisterXpub.
func (mr *MockAdmBuxClientMockRecorder) RegisterXpub(xpriv interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterXpub", reflect.TypeOf((*MockAdmBuxClient)(nil).RegisterXpub), xpriv)
}

// MockBuxClientFactory is a mock of BuxClientFactory interface.
type MockBuxClientFactory struct {
        ctrl     *gomock.Controller
        recorder *MockBuxClientFactoryMockRecorder
}

// MockBuxClientFactoryMockRecorder is the mock recorder for MockBuxClientFactory.
type MockBuxClientFactoryMockRecorder struct {
        mock *MockBuxClientFactory
}

// NewMockBuxClientFactory creates a new mock instance.
func NewMockBuxClientFactory(ctrl *gomock.Controller) *MockBuxClientFactory {
        mock := &MockBuxClientFactory{ctrl: ctrl}
        mock.recorder = &MockBuxClientFactoryMockRecorder{mock}
        return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuxClientFactory) EXPECT() *MockBuxClientFactoryMockRecorder {
        return m.recorder
}

// CreateAdminBuxClient mocks base method.
func (m *MockBuxClientFactory) CreateAdminBuxClient() (users.AdmBuxClient, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "CreateAdminBuxClient")
        ret0, _ := ret[0].(users.AdmBuxClient)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// CreateAdminBuxClient indicates an expected call of CreateAdminBuxClient.
func (mr *MockBuxClientFactoryMockRecorder) CreateAdminBuxClient() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAdminBuxClient", reflect.TypeOf((*MockBuxClientFactory)(nil).CreateAdminBuxClient))
}

// CreateWithAccessKey mocks base method.
func (m *MockBuxClientFactory) CreateWithAccessKey(accessKey string) (users.UserBuxClient, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "CreateWithAccessKey", accessKey)
        ret0, _ := ret[0].(users.UserBuxClient)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// CreateWithAccessKey indicates an expected call of CreateWithAccessKey.
func (mr *MockBuxClientFactoryMockRecorder) CreateWithAccessKey(accessKey interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWithAccessKey", reflect.TypeOf((*MockBuxClientFactory)(nil).CreateWithAccessKey), accessKey)
}

// CreateWithXpriv mocks base method.
func (m *MockBuxClientFactory) CreateWithXpriv(xpriv string) (users.UserBuxClient, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "CreateWithXpriv", xpriv)
        ret0, _ := ret[0].(users.UserBuxClient)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// CreateWithXpriv indicates an expected call of CreateWithXpriv.
func (mr *MockBuxClientFactoryMockRecorder) CreateWithXpriv(xpriv interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWithXpriv", reflect.TypeOf((*MockBuxClientFactory)(nil).CreateWithXpriv), xpriv)
}
