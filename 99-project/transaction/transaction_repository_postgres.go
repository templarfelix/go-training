package transaction

import (
	"99-project/entity"
	"gorm.io/gorm"
)

type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		db: db,
	}
}

func (repository *GormRepository) Find(id entity.ID) (*entity.Transaction, error) {
	result := entity.Transaction{}
	tx := repository.db.First(&result, "id", id.String())
	return &result, tx.Error
}

func (repository *GormRepository) Create(Transaction *entity.Transaction) (entity.ID, error) {
	tx := repository.db.Create(Transaction)
	return Transaction.ID, tx.Error
}

func (repository *GormRepository) FindAll() ([]*entity.Transaction, error) {
	var Transactions []*entity.Transaction
	tx := repository.db.Find(&Transactions)
	return Transactions, tx.Error
}

func (repository *GormRepository) Search(query string) ([]*entity.Transaction, error) {
	return nil, nil
}
