package storage

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type DBStorage struct {
	db *sql.DB
}

func NewDBStorage(serveName string) *DBStorage {
	db, _ := sql.Open("pgx", serveName)
	defer db.Close()

	return &DBStorage{
		db: db,
	}
}
