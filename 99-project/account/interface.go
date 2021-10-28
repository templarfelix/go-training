package account

import "99-project/entity"

type Reader interface {
	Find(id entity.ID) (*entity.Account, error)
	Search(query string) ([]*entity.Account, error)
	FindAll() ([]*entity.Account, error)
}

type Writer interface {
	Create(b *entity.Account) (entity.ID, error)
	Save(b *entity.Account) error
	Delete(id entity.ID) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	Reader
	Writer
}
