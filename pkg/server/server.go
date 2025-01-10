package server

import (
	"context"
	"gorm.io/gorm"
	"net/http"
	"transaction-service/pkg/config"
	"transaction-service/pkg/handler"
	"transaction-service/pkg/logger"
	"transaction-service/pkg/middleware"
	"transaction-service/pkg/models"
	"transaction-service/pkg/repository/impl"
	impl2 "transaction-service/pkg/service/impl"

	"github.com/gorilla/mux"
)

func Start(ctx context.Context) {
	appConfig := config.GetAppConfiguration()

	r := mux.NewRouter()
	r.Use(middleware.RequestIDMiddleware)

	repositoryImpl, ddb, err := impl.NewRepositoryImpl(appConfig.PGConfig)
	if err != nil {
		logger.WithCtx(ctx).Fatalf("Failed to create repositoryImpl: %v", err)
	}

	err = autoMigrate(ddb)
	if err != nil {
		logger.WithCtx(ctx).Fatalf("Failed to auto migrate: %v", err)
	}
	transactionRepository := impl.NewTransactionImpl(repositoryImpl)
	transactionService := impl2.NewTransactionServiceImpl(transactionRepository)
	h := handler.NewHandler(transactionService)

	r.HandleFunc("/accounts", h.CreateAccount).Methods("POST")
	r.HandleFunc("/accounts/{accountId}", h.GetAccountByID).Methods("GET")
	r.HandleFunc("/transactions", h.CreateTransaction).Methods("POST")

	http.ListenAndServe(":"+appConfig.Port, r)
	logger.WithCtx(ctx).Info("-----Server started------")
}

func autoMigrate(ddb *gorm.DB) error {
	if err := ddb.AutoMigrate(&models.Account{}); err != nil {
		return err
	}

	if err := ddb.AutoMigrate(&models.Transaction{}); err != nil {
		return err
	}

	if err := ddb.AutoMigrate(&models.OperationType{}); err != nil {
		return err
	}

	operationType1 := &models.OperationType{
		OperationTypeID: 1,
		Description:     "Normal Purchase",
	}
	if err := ddb.FirstOrCreate(&operationType1, "operation_type_id = ?", operationType1.OperationTypeID).Error; err != nil {
		return err
	}

	operationType2 := &models.OperationType{
		OperationTypeID: 2,
		Description:     "Installment Purchase",
	}
	if err := ddb.FirstOrCreate(&operationType2, "operation_type_id = ?", operationType2.OperationTypeID).Error; err != nil {
		return err
	}

	operationType3 := &models.OperationType{
		OperationTypeID: 3,
		Description:     "Withdraw",
	}
	if err := ddb.FirstOrCreate(&operationType3, "operation_type_id = ?", operationType3.OperationTypeID).Error; err != nil {
		return err
	}

	operationType4 := &models.OperationType{
		OperationTypeID: 4,
		Description:     "Payment",
	}
	if err := ddb.FirstOrCreate(&operationType4, "operation_type_id = ?", operationType4.OperationTypeID).Error; err != nil {
		return err
	}

	return nil
}
