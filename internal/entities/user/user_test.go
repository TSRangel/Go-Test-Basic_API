package user

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	as := assert.New(t)
	user, err := NewUser("Thiago", "slipknothiago@gmail.com", "12345")
	as.Nil(err)
	as.NotNil(user.ID)
	as.Equal("Thiago", user.Name)
	as.Equal("slipknothiago@gmail.com", user.Email)
	as.NotEmpty(user.Password)
}

func TestUserValidatePassword(t *testing.T) {
	as := assert.New(t)
	user, err := NewUser("Thiago", "slipknothiago@gmail.com", "12345")
	as.Nil(err)
	as.True(user.ValidatePassword("12345"))
	as.NotEqual("12345", user.Password)
}
