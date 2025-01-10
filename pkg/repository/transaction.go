package repository

import (
	"context"
	"transaction-service/pkg/models"
)

type TransactionRepository interface {
	Repository
	SaveAccount(ctx context.Context, account *models.Account) error
	FetchAccountByID(ctx context.Context, accountID int) (*models.Account, error)
	SaveTransaction(ctx context.Context, transaction *models.Transaction) error
}
