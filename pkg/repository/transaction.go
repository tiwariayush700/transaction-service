package repository

import (
	"context"
	"transaction-service/pkg/models"
)

type TransactionRepository interface {
	Repository
	CreateAccount(ctx context.Context, account *models.Account) error
	GetAccountByID(ctx context.Context, accountID int) (*models.Account, error)
	CreateTransaction(ctx context.Context, transaction *models.Transaction) error
}
