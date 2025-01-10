package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"transaction-service/pkg/models"
	"transaction-service/pkg/tserrors"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockTransactionService is a mock implementation of the TransactionService interface
type MockTransactionService struct {
	mock.Mock
}

func (m *MockTransactionService) CreateAccount(ctx context.Context, account *models.Account) error {
	args := m.Called(ctx, account)
	return args.Error(0)
}

func (m *MockTransactionService) GetAccountByID(ctx context.Context, accountID int) (*models.Account, error) {
	args := m.Called(ctx, accountID)
	return args.Get(0).(*models.Account), args.Error(1)
}

func (m *MockTransactionService) CreateTransaction(ctx context.Context, transaction *models.Transaction) error {
	args := m.Called(ctx, transaction)
	return args.Error(0)
}

func TestCreateAccount(t *testing.T) {
	mockService := new(MockTransactionService)
	handler := NewHandler(mockService)

	account := &models.Account{AccountID: 1, DocumentNumber: "Test Account"}
	mockService.On("CreateAccount", mock.Anything, account).Return(nil)

	body, _ := json.Marshal(account)
	req, _ := http.NewRequest("POST", "/accounts", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	handler.CreateAccount(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	mockService.AssertExpectations(t)
}

func TestCreateAccount_InvalidPayload(t *testing.T) {
	mockService := new(MockTransactionService)
	handler := NewHandler(mockService)

	req, _ := http.NewRequest("POST", "/accounts", bytes.NewBuffer([]byte("invalid payload")))
	rr := httptest.NewRecorder()

	handler.CreateAccount(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestGetAccountByID(t *testing.T) {
	mockService := new(MockTransactionService)
	handler := NewHandler(mockService)

	account := &models.Account{AccountID: 1, DocumentNumber: "Test Account"}
	mockService.On("GetAccountByID", mock.Anything, 1).Return(account, nil)

	req, _ := http.NewRequest("GET", "/accounts/1", nil)
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/accounts/{accountId}", handler.GetAccountByID)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mockService.AssertExpectations(t)
}

func TestGetAccountByID_InvalidID(t *testing.T) {
	mockService := new(MockTransactionService)
	handler := NewHandler(mockService)

	req, _ := http.NewRequest("GET", "/accounts/invalid", nil)
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/accounts/{accountId}", handler.GetAccountByID)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestCreateTransaction(t *testing.T) {
	mockService := new(MockTransactionService)
	handler := NewHandler(mockService)

	transaction := &models.Transaction{TransactionID: 1, Amount: 100}
	mockService.On("CreateTransaction", mock.Anything, transaction).Return(nil)

	body, _ := json.Marshal(transaction)
	req, _ := http.NewRequest("POST", "/transactions", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	handler.CreateTransaction(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	mockService.AssertExpectations(t)
}

func TestCreateTransaction_InvalidPayload(t *testing.T) {
	mockService := new(MockTransactionService)
	handler := NewHandler(mockService)

	req, _ := http.NewRequest("POST", "/transactions", bytes.NewBuffer([]byte("invalid payload")))
	rr := httptest.NewRecorder()

	handler.CreateTransaction(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestRespondWithError(t *testing.T) {
	rr := httptest.NewRecorder()
	err := tserrors.New(1, "error")

	respondWithError(rr, err)

	assert.Equal(t, http.StatusOK, rr.Code)
	var response models.Response
	json.NewDecoder(rr.Body).Decode(&response)
	assert.Equal(t, "error", response.Status)
	assert.Equal(t, "test error", response.Message)
	assert.Equal(t, 123, response.Code)
}
