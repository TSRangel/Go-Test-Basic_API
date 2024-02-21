package database

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/TSRangel/Go-Test-Basic_API/internal/entities/user"
	"github.com/TSRangel/Go-Test-Basic_API/pkg/tools"
)

func TestCreateUser(t *testing.T) {
	as := assert.New(t)
	newUser, err := user.NewUser("Thiago", "slipknothiago@gmail.com", "12345")
	as.Nil(err)
	db, err := tools.DBConnection()
	as.Nil(err)
	defer db.Close()
	err = tools.CreateUsersTable(db)
	as.Nil(err)
	newUserConnection := NewUserConnection(db)
	err = newUserConnection.Create(newUser)
	as.Nil(err)
	stmt, err := db.Prepare("select name, email, password from users where id = ?")
	as.Nil(err)
	defer stmt.Close()
	var testedUser user.User 
	err = stmt.QueryRow(newUser.ID).Scan(&testedUser.Name, &testedUser.Email, &testedUser.Password)
	as.Nil(err)
	as.Equal(newUser.Name, testedUser.Name)
	as.Equal(newUser.Email, testedUser.Email)
	as.Equal(newUser.Password, testedUser.Password)
}

func TestFindByEmail(t *testing.T) {
	as := assert.New(t)
	db, err := tools.DBConnection()
	as.Nil(err)
	defer db.Close()
	err = tools.CreateUsersTable(db)
	as.Nil(err)
	newUser, err := user.NewUser("Thiago", "slipknothiago@gmail.com", "12345")
	as.Nil(err)
	newUserConnection := NewUserConnection(db)
	err = newUserConnection.Create(newUser)
	as.Nil(err)
	foundUser, err := newUserConnection.FindByEmail("slipknothiago@gmail.com")
	as.Nil(err)
	as.Equal(newUser.ID, foundUser.ID)
	as.Equal(newUser.Name, foundUser.Name)
	as.Equal(newUser.Password, foundUser.Password)
}