package entity

import (
	"context"
	"time"
)

type User struct {
	ID        ID
	Name     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRepository interface {
	GetByID(ctx context.Context, id ID) (User, error)
}

func NewUser(name string) (*User, error) {
	b := &User{
		ID:        NewID(),
		Name:     name,
		CreatedAt: time.Now(),
	}
	err := b.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return b, nil
}

func (b *User) Validate() error {
	if b.Name == "" {
		return ErrInvalidEntity
	}
	return nil
}