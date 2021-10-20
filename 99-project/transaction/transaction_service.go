package transaction

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

func (s *Service) Create(b *entity.Transaction) (entity.ID, error) {
	b.ID = entity.NewID()
	b.CreatedAt = time.Now()

	err := b.Validate()
	if err != nil {
		return entity.NewID(), err
	}

	return s.repo.Create(b)
}

func (s *Service) Find(id entity.ID) (*entity.Transaction, error) {
	return s.repo.Find(id)
}

func (s *Service) Search(query string) ([]*entity.Transaction, error) {
	return s.repo.Search(strings.ToLower(query))
}

func (s *Service) FindAll() ([]*entity.Transaction, error) {
	return s.repo.FindAll()
}
