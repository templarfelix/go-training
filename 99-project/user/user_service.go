package user

import (
	"99-project/entity"
	"strings"
	"time"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) Store(b *entity.User) (entity.ID, error) {
	b.ID = entity.NewID()
	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()

	err := b.Validate()
	if err != nil {
		return entity.NewID(), err
	}

	return s.repo.Store(b)
}

func (s *Service) Find(id entity.ID) (*entity.User, error) {
	return s.repo.Find(id)
}

func (s *Service) Search(query string) ([]*entity.User, error) {
	return s.repo.Search(strings.ToLower(query))
}

func (s *Service) FindAll() ([]*entity.User, error) {
	return s.repo.FindAll()
}

func (s *Service) Delete(id entity.ID) error {
	return s.repo.Delete(id)
}
