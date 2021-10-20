package entity

import (
	"time"
)

type Account struct {
	ID            ID        `json:"id" gorm:"primaryKey, column:id"`
	UserID        ID        `json:"user_id" gorm:"column:user_id"`
	AccountType   string    `json:"account_type" gorm:"column:account_type"`
	AccountNumber string    `json:"account_number" gorm:"column:account_number"`
	Amount        float64   `json:"amount" gorm:"column:account_amount"`
	Limit         float64   `json:"limit" gorm:"column:account_limit"`
	CreatedAt     time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"column:updated_at"`
	Status        Status    `json:"-" gorm:"column:status"`
}

func (entity *Account) Validate() error {
	if entity.AccountType == "" {
		return ErrInvalidEntity
	}
	if entity.AccountNumber == "" {
		return ErrInvalidEntity
	}
	if entity.Limit <= 1000 {
		return ErrInvalidEntity
	}
	return nil
}
