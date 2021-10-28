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

func (repository *GormRepository) Save(Account *entity.Account) error {
	tx := repository.db.Save(Account)
	return tx.Error
}

func (repository *GormRepository) FindAll() ([]*entity.Account, error) {
	var Accounts []*entity.Account
	tx := repository.db.Where("status", entity.ACTIVE).Find(&Accounts)
	return Accounts, tx.Error
}

// FIXME IMPLEMENTAR
func (repository *GormRepository) Search(query string) ([]*entity.Account, error) {
	return nil, nil
}

func (repository *GormRepository) Delete(id entity.ID) error {
	Account, err := repository.Find(id)

	if err != nil {
		return err
	}

	Account.Status = entity.DELETED
	tx := repository.db.Save(&Account)

	return tx.Error
}
