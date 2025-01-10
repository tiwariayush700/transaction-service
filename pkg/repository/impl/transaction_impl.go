package impl

import (
	"context"
	"transaction-service/pkg/logger"
	"transaction-service/pkg/models"
	"transaction-service/pkg/repository"
	"transaction-service/pkg/tserrors"
)

type transactionImpl struct {
	// repositoryImpl is the implementation of the repository interface
	// overrides the generic repository for inheriting crud operation
	repository.Repository
}

func (t *transactionImpl) SaveAccount(ctx context.Context, account *models.Account) error {
	return t.Repository.Create(ctx, account)
}

func (t *transactionImpl) FetchAccountByID(ctx context.Context, accountID int) (*models.Account, error) {
	account := &models.Account{}
	err := t.Repository.Get(ctx, account, "account_id", accountID)
	if err != nil {
		logger.WithCtx(ctx).Errorf("Failed to get account by id: %v", err)
		return nil, tserrors.New(tserrors.DBError.Code, err.Error())
	}

	return account, nil
}

func (t *transactionImpl) SaveTransaction(ctx context.Context, transaction *models.Transaction) error {
	return t.Repository.Create(ctx, transaction)
}

func NewTransactionImpl(impl repository.Repository) repository.TransactionRepository {
	return &transactionImpl{Repository: impl}
}
