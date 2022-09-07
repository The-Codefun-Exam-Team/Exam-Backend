package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	*sql.DB
}

func New(connection string) (*DB, error) {
	db_handle, err := sql.Open("mysql", connection)
	if err != nil {
		return nil, err
	}

	db := DB{
		db_handle,
	}

	return &db, nil
}