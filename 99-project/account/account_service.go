package account

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

func (s *Service) Create(b *entity.Account) (entity.ID, error) {
	b.ID = entity.NewID()
	b.CreatedAt = time.Now()
	b.Status = entity.ACTIVE

	err := b.Validate()
	if err != nil {
		return entity.NewID(), err
	}

	return s.repo.Create(b)
}

func (s *Service) UpdateAmount(id entity.ID, transactionType entity.TransactionType, value float64) error {

	account, err := s.Find(id)
	if err != nil {
		return err
	}
	if transactionType == entity.CREDIT {
		account.Amount += value
	} else if transactionType == entity.DEBIT {
		account.Amount -= value
	}

	err = s.repo.Save(account)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) Find(id entity.ID) (*entity.Account, error) {
	return s.repo.Find(id)
}

func (s *Service) Search(query string) ([]*entity.Account, error) {
	return s.repo.Search(strings.ToLower(query))
}

func (s *Service) FindAll() ([]*entity.Account, error) {
	return s.repo.FindAll()
}

func (s *Service) Delete(id entity.ID) error {
	return s.repo.Delete(id)
}
