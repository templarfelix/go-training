package entity_test

import (
	"99-project/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser(t *testing.T) {
	b, err := entity.NewUser("American Gods")
	assert.Nil(t, err)
	assert.Equal(t, b.Name, "American Gods")
	assert.NotNil(t, b.ID)
}

func TestUserValidate(t *testing.T) {
	type test struct {
		name    string
		want     error
	}

	tests := []test{
		{
			name:    "American Gods",
			want:     nil,
		},
		{
			name:    "",
			want:     entity.ErrInvalidEntity,
		},
	}
	for _, tc := range tests {

		_, err := entity.NewUser(tc.name)
		assert.Equal(t, err, tc.want)
	}

}