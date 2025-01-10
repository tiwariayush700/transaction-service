package models

type OperationType struct {
	OperationTypeID int    `json:"operation_type_id" gorm:"primaryKey"`
	Description     string `json:"description"`
}
