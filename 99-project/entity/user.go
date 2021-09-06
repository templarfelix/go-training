package entity

import (
	"github.com/paemuri/brdoc"
	"time"
)

type User struct {
	ID             ID        `json:"id", gorm:"column:id"`
	Name           string    `json:"name", gorm:"column:name"`
	DocumentNumber string    `json:"document_number", gorm:"column:document_number"`
	CreatedAt      time.Time `json:"created_at", gorm:"column:created_at"`
	UpdatedAt      time.Time `json:"updated_at", gorm:"column:updated_at"`
}

func (entity *User) Validate() error {
	if entity.Name == "" {
		return ErrInvalidEntity
	}
	if !brdoc.IsCPF(entity.DocumentNumber) {
		return ErrDocumentNumberInvalid
	}
	return nil
}
