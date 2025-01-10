package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"transaction-service/pkg/logger"
	"transaction-service/pkg/models"
	"transaction-service/pkg/service"
	"transaction-service/pkg/tserrors"

	"github.com/gorilla/mux"
)

type Handler struct {
	transactionService service.TransactionService
}

func NewHandler(transactionService service.TransactionService) *Handler {
	return &Handler{transactionService: transactionService}
}

func (h *Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var account models.Account
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		respondWithJSON(w, http.StatusBadRequest, models.Response{Status: "error", Message: "Invalid request payload"})
		return
	}

	if err := h.transactionService.CreateAccount(ctx, &account); err != nil {
		logger.WithCtx(ctx).Errorf("Failed to create account: %v", err)
		respondWithError(w, err)
		return
	}

	respondWithJSON(w, http.StatusCreated, models.Response{Status: "success", Message: "Account created successfully", Data: account})
}

func (h *Handler) GetAccountByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	accountID, err := strconv.Atoi(vars["accountId"])
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, models.Response{Status: "error", Message: "Invalid account ID"})
		return
	}

	account, err := h.transactionService.GetAccountByID(ctx, accountID)
	if err != nil {
		logger.WithCtx(ctx).Errorf("Failed to get account by ID: %v", err)
		respondWithError(w, err)
		return
	}

	respondWithJSON(w, http.StatusOK, models.Response{Status: "success", Message: "Account retrieved successfully", Data: account})
}

func (h *Handler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var transaction models.Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		respondWithJSON(w, http.StatusBadRequest, models.Response{Status: "error", Message: "Invalid request payload"})
		return
	}

	if err := h.transactionService.CreateTransaction(ctx, &transaction); err != nil {
		logger.WithCtx(ctx).Errorf("Failed to create transaction: %v", err)
		respondWithError(w, err)
		return
	}

	respondWithJSON(w, http.StatusCreated, models.Response{Status: "success", Message: "Transaction created successfully", Data: transaction})
}

func respondWithJSON(w http.ResponseWriter, status int, response models.Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

func respondWithError(w http.ResponseWriter, err error) {
	if tsErr, ok := err.(*tserrors.Error); ok {
		respondWithJSON(w, tsErr.Code, models.Response{Status: "error", Message: tsErr.Error()})
	} else {
		respondWithJSON(w, http.StatusInternalServerError, models.Response{Status: "error", Message: err.Error()})
	}
}
