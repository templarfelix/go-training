package transaction

import "99-project/entity"

type Reader interface {
	Find(id entity.ID) (*entity.Transaction, error)
	Search(query string) ([]*entity.Transaction, error)
	FindAll() ([]*entity.Transaction, error)
}

type Writer interface {
	Create(b *entity.Transaction) (entity.ID, error)
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	Reader
	Writer
}
