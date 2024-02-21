package tools

import (
	"database/sql"
	_ "modernc.org/sqlite"
)

func DBConnection() (*sql.DB, error){
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func DBConnectionWithDataSource(dataSource string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dataSource)
	if err != nil {
		return nil, err
	}
	return db, err
}

func CreateUsersTable(db *sql.DB) error {
	_, err := db.Exec("create table users (id uuid primary key, name varchar(150), email varchar(150), password varchar(250))")
	if err != nil {
		return err
	}
	return nil
}

func CreateProductsTable(db *sql.DB) error {
	_, err := db.Exec("create table products (id uuid primary key, name varchar(150), price decimal, created_at timestamp)")
	if err != nil {
		return err
	}
	return nil
}