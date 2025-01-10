package models

type Account struct {
	AccountID      int    `json:"account_id" gorm:"primaryKey"`
	DocumentNumber string `json:"document_number"`
}
