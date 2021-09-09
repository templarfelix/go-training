package account

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

func (repository *GormRepository) Find(id entity.ID) (*entity.Account, error) {
	result := entity.Account{}
	tx := repository.db.First(&result, "id", id.String())
	return &result, tx.Error
}

func (repository *GormRepository) Create(Account *entity.Account) (entity.ID, error) {
	tx := repository.db.Create(Account)
	return Account.ID, tx.Error
}

func (repository *GormRepository) FindAll() ([]*entity.Account, error) {
	var Accounts []*entity.Account
	tx := repository.db.Find(&Accounts)
	return Accounts, tx.Error
}

// FIXME IMPLEMENTAR
func (repository *GormRepository) Search(query string) ([]*entity.Account, error) {
	return nil, nil
}

func (repository *GormRepository) Delete(id entity.ID) error {
	var Accounts *entity.Account
	tx := repository.db.Delete(&Accounts, id)
	return tx.Error
}
