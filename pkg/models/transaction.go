package models

import "time"

type Transaction struct {
	TransactionID   int       `json:"transaction_id" gorm:"primaryKey;autoIncrement"`
	AccountID       int       `json:"account_id" gorm:"foreignKey:account_id;references:account_id"`
	OperationTypeID int       `json:"operation_type_id" gorm:"foreignKey:operation_type_id;references:operation_type_id"`
	Amount          float64   `json:"amount"`
	EventDate       time.Time `json:"event_date"`
}
