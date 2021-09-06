package user

import (
	"99-project/entity"
	"gorm.io/gorm"
)

type GormRepository struct {
	db   *gorm.DB
}

func NewGormRepository(db   *gorm.DB) *GormRepository {
	return &GormRepository{
		db:   db,
	}
}

func (r *GormRepository) Find(id entity.ID) (*entity.User, error) {
	result := entity.User{}

	r.db.First(result, id)

	return &result, nil

}

func (r *GormRepository) Store(b *entity.User) (entity.ID, error) {
	return entity.NewID(), nil
}

func (r *GormRepository) FindAll() ([]*entity.User, error) {
	return nil, nil
}

func (r *GormRepository) Search(query string) ([]*entity.User, error) {
	return nil, nil
}

func (r *GormRepository) Delete(id entity.ID) error {
	return nil
}