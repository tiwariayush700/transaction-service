package impl

import (
	"context"
	"fmt"
	"time"
	"transaction-service/pkg/logger"
	"transaction-service/pkg/models"
	"transaction-service/pkg/repository"
	"transaction-service/pkg/service"
)

type transactionServiceImpl struct {
	transactionRepository repository.TransactionRepository
}

func (t *transactionServiceImpl) CreateAccount(ctx context.Context, account *models.Account) error {
	return t.transactionRepository.SaveAccount(ctx, account)
}

func (t *transactionServiceImpl) GetAccountByID(ctx context.Context, accountID int) (*models.Account, error) {
	return t.transactionRepository.FetchAccountByID(ctx, accountID)
}

func (t *transactionServiceImpl) CreateTransaction(ctx context.Context, transaction *models.Transaction) error {
	// Check if the account exists
	account := &models.Account{}
	err := t.transactionRepository.Get(ctx, account, "account_id", transaction.AccountID)
	if err != nil {
		logger.WithCtx(ctx).Errorf("Failed to get account by id: %v", err)
		return fmt.Errorf("account with id %d does not exist", transaction.AccountID)
	}

	operationType := &models.OperationType{}
	err = t.transactionRepository.Get(ctx, operationType, "operation_type_id", transaction.OperationTypeID)
	if err != nil {
		logger.WithCtx(ctx).Errorf("Failed to get operation type by id: %v", err)
		return err
	}

	switch operationType.OperationTypeID {
	case 1, 2, 3:
		transaction.Amount = -transaction.Amount
	case 4:
		transaction.Amount = transaction.Amount
	}
	transaction.EventDate = time.Now()
	return t.transactionRepository.SaveTransaction(ctx, transaction)
}

func NewTransactionServiceImpl(transactionRepository repository.TransactionRepository) service.TransactionService {
	return &transactionServiceImpl{transactionRepository: transactionRepository}
}
