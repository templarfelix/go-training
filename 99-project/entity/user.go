package entity

import (
	"context"
	"github.com/paemuri/brdoc"
	"time"
)

// https://bancodofernando.com/user/123e4567-e89b-12d3-a456-426614174000
// https://bancodofernando.com/user/2
// https://bancodofernando.com/user/3
// https://bancodofernando.com/user/4
/*
	ID, NAME, DOCUMENT_NUMBER, CREATED_AT, UPDATED_AT
     1, FERNANDO, 0000, xx,xx
     2, ZE, 0000, xx,xx
     3, ZE, 0000, xx,xx
     123e4567-e89b-12d3-a456-426614174000,  FERNANDO, 0000, xx, xx
*/
type User struct {
	ID        ID
	Name     string
	DocumentNumber string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(name string, documentNumber string) (*User, error) {
	b := &User{
		ID:        NewID(),
		Name:     name,
		DocumentNumber: documentNumber,
		CreatedAt: time.Now(),
	}
	err := b.Validate()
	if err != nil {
		return nil, err
	}
	return b, nil
}

// validação das regras de negocio do usuario
func (b *User) Validate() error {
	if b.Name == "" {
		return ErrInvalidEntity
	}
	if !brdoc.IsCPF(b.DocumentNumber) {
		return ErrDocumentNumberInvalid
	}
	return nil
}

type UserRepository interface {
	GetByID(ctx context.Context, id ID) (User, error)
}
