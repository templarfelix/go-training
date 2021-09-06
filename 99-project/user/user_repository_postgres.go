package user

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

func (repository *GormRepository) Find(id entity.ID) (*entity.User, error) {
	result := entity.User{}
	tx := repository.db.First(&result, "id", id.String())
	return &result, tx.Error
}

func (repository *GormRepository) Create(user *entity.User) (entity.ID, error) {
	tx := repository.db.Create(user)
	return user.ID, tx.Error
}

func (repository *GormRepository) FindAll() ([]*entity.User, error) {
	var users []*entity.User
	tx := repository.db.Find(&users)
	return users, tx.Error
}

// FIXME IMPLEMENTAR
func (repository *GormRepository) Search(query string) ([]*entity.User, error) {
	return nil, nil
}

// FIXME IMPLEMENTAR
func (repository *GormRepository) Delete(id entity.ID) error {
	return nil
}
