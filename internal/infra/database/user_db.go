package database

import (
	"github.com/TSRangel/Go-Test-Basic_API/internal/entities/user"
	"database/sql"
)

type UserConnection struct {
	DB *sql.DB
}

func NewUserConnection(db *sql.DB) *UserConnection {
	return &UserConnection{DB: db}
}

func (u *UserConnection) Create(user *user.User) error {
	stmt, err := u.DB.Prepare("insert into users (id, name, email, password) values (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.ID, user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserConnection) FindByEmail(email string) (*user.User, error) {
	stmt, err := u.DB.Prepare("select id, name, password from users where email = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var searchedUser user.User
	err = stmt.QueryRow(email).Scan(&searchedUser.ID, &searchedUser.Name, &searchedUser.Password)
	if err != nil {
		return nil, err
	}
	return &searchedUser, nil
}