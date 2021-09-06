package user

import (
	"99-project/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStore(t *testing.T) {
	repo := NewMemorymRepository()
	service := NewService(repo)
	b := &entity.User{
		Name:        "Test",
		DocumentNumber: "212.712.270-45",
	}
	id, err := service.Create(b)
	assert.Nil(t, err)
	assert.NotNil(t, id)
	assert.False(t, b.CreatedAt.IsZero())
}

func TestSearchAndFindAll(t *testing.T) {
	repo := NewMemorymRepository()
	service := NewService(repo)
	b := &entity.User{
		Name:        "Test",
		DocumentNumber: "212.712.270-45",
	}
	b2 := &entity.User{
		Name:        "Google",
		DocumentNumber: "212.712.270-45",
	}
	bID, _ := service.Create(b)
	_, _ = service.Create(b2)

	t.Run("search", func(t *testing.T) {
		c, err := service.Search("Test")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(c))
		assert.Equal(t, "Test", c[0].Name)

		c, err = service.Search("bing")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, c)
	})
	t.Run("find all", func(t *testing.T) {
		all, err := service.FindAll()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("find", func(t *testing.T) {
		saved, err := service.Find(bID)
		assert.Nil(t, err)
		assert.Equal(t, b.Name, saved.Name)
	})
}

func TestDelete(t *testing.T) {
	repo := NewMemorymRepository()
	service := NewService(repo)
	b := &entity.User{
		Name:        "Test",
		DocumentNumber: "212.712.270-45",
	}
	b2 := &entity.User{
		Name:        "Google",
		DocumentNumber: "212.712.270-45",
	}
	bID, _ := service.Create(b)
	b2ID, _ := service.Create(b2)

	err := service.Delete(bID)
	assert.Equal(t, nil, err)

	err = service.Delete(b2ID)
	assert.Nil(t, err)
	_, err = service.Find(b2ID)
	assert.Equal(t, entity.ErrNotFound, err)

}