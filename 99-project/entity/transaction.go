package entity

import (
	"time"
)

type Transaction struct {
	ID              ID              `json:"id" gorm:"primaryKey, column:id"`
	AccountID       ID              `json:"account_id" gorm:"column:account_id"`
	TransactionType TransactionType `json:"type" gorm:"column:type"`
	Value           float64         `json:"value" gorm:"column:value"`
	CreatedAt       time.Time       `json:"created_at" gorm:"column:created_at"`
	UpdatedAt       time.Time       `json:"updated_at" gorm:"column:updated_at"`
}

func (entity *Transaction) Validate() error {
	if entity.Value == 0 {
		return ErrValueInvalidEntity
	}
	return nil
}
