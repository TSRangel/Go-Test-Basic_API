package user

import (
	"github.com/TSRangel/Go-Test-Basic_API/pkg/tools"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID tools.ID `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"-"`
}

func NewUser(name, email, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		ID: tools.NewID(),
		Name: name,
		Email: email,
		Password: string(hash),
	}, nil
}

func (u *User)ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}