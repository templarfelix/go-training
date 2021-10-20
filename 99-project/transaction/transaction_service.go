package transaction

import (
	"99-project/account"
	"99-project/entity"
	"fmt"
	"strings"
	"time"
)

type Service struct {
	repo           Repository
	accountService account.Service
}

func NewService(r Repository, a *account.Service) *Service {
	return &Service{
		repo:           r,
		accountService: *a,
	}
}

func (s *Service) Create(b *entity.Transaction) (entity.ID, error) {
	b.ID = entity.NewID()
	b.CreatedAt = time.Now()

	err := b.Validate()
	if err != nil {
		return entity.NewID(), err
	}

	// FIXME REGRA DE ACCOUTNS AQUI

	account, err := s.accountService.Find(b.AccountID)
	if err != nil {
		return entity.NewID(), err
	}

	fmt.Println(account)

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
