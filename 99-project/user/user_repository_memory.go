package user

import (
	"99-project/entity"
	"strings"
)

type IRepo struct {
	m map[string]*entity.User
}

func NewMemorymRepository() *IRepo {
	var m = map[string]*entity.User{}
	return &IRepo{
		m: m,
	}
}

func (r *IRepo) Store(a *entity.User) (entity.ID, error) {
	r.m[a.ID.String()] = a
	return a.ID, nil
}

func (r *IRepo) Find(id entity.ID) (*entity.User, error) {
	if r.m[id.String()] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id.String()], nil
}

func (r *IRepo) Search(query string) ([]*entity.User, error) {
	var d []*entity.User
	for _, j := range r.m {
		if strings.Contains(strings.ToLower(j.Name), query) {
			d = append(d, j)
		}
	}
	if len(d) == 0 {
		return nil, entity.ErrNotFound
	}

	return d, nil
}

func (r *IRepo) FindAll() ([]*entity.User, error) {
	var d []*entity.User
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

func (r *IRepo) Delete(id entity.ID) error {
	if r.m[id.String()] == nil {
		return entity.ErrNotFound
	}
	r.m[id.String()] = nil
	return nil
}