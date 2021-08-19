package entity_test

import (
	"99-project/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser(t *testing.T) {
	b, err := entity.NewUser("American Gods","053.043.099-10")
	assert.Nil(t, err)
	assert.Equal(t, b.Name, "American Gods")
	assert.NotNil(t, b.ID)
}

func TestUserValidate(t *testing.T) {
	type test struct {
		name    string
		documentNumber string
		want     error
	}

	tests := []test{
		{
			name:    "American Gods",
			documentNumber: "053.043.099-10",
			want:     nil,
		},
		{
			name:    "",
			documentNumber: "053.043.099-10",
			want:     entity.ErrInvalidEntity,
		},
		{
			name:    "Fernando Luís",
			documentNumber: "",
			want:     entity.ErrDocumentNumberInvalid,
		},

	}
	for _, tc := range tests {
		_, err := entity.NewUser(tc.name, tc.documentNumber)
		assert.Equal(t, err, tc.want)
	}

}