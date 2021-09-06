package user

import "99-project/entity"

type Reader interface {
	Find(id entity.ID) (*entity.User, error)
	Search(query string) ([]*entity.User, error)
	FindAll() ([]*entity.User, error)
}

type Writer interface {
	Create(b *entity.User) (entity.ID, error)
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
