package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type DBStorage struct {
	db *sql.DB
}

func NewDBStorage(serveName string) *DBStorage {
	// TODO Put all this stuff in the secure place
	dataSourceName := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		serveName, `vadim`, `gocoursepwd`, `go-course`)

	db, _ := sql.Open("pgx", dataSourceName)
	// TODO Close but when?
	//defer db.Close()

	return &DBStorage{
		db: db,
	}
}

func (s *DBStorage) Ping() error {
	return s.db.Ping()
}
