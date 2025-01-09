package impl

import (
	"context"
	"gorm.io/gorm"
	"transaction-service/pkg/logger"
	"transaction-service/pkg/models"
	"transaction-service/pkg/repository"
	"transaction-service/pkg/tserrors"
)

type transactionImpl struct {
	repositoryImpl
	db *gorm.DB
}

func (t *transactionImpl) CreateAccount(ctx context.Context, account *models.Account) error {
	return t.repositoryImpl.Create(ctx, account)
}

func (t *transactionImpl) GetAccountByID(ctx context.Context, accountID int) (*models.Account, error) {
	account := &models.Account{}
	err := t.repositoryImpl.Get(ctx, account, accountID)
	if err != nil {
		logger.WithCtx(ctx).Errorf("Failed to get account by id: %v", err)
		return nil, tserrors.New(tserrors.DBError.Code, err.Error())
	}

	return account, nil
}

func (t *transactionImpl) CreateTransaction(ctx context.Context, transaction *models.Transaction) error {
	return t.repositoryImpl.Create(ctx, transaction)
}

func NewCardRepositoryImpl(db *gorm.DB) repository.TransactionRepository {
	return &transactionImpl{db: db}
}
