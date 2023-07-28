package data

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type MySqlDB struct {
	client *sql.DB
}

func NewMySqlDB(dsn string) (*MySqlDB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &MySqlDB{db}, nil
}

func (db *MySqlDB) EnsureCreated() error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS `links` (`id` VARCHAR(64) NOT NULL, `url` VARCHAR(2048) NOT NULL, PRIMARY KEY (`id`))")
	return err
}

func (db *MySqlDB) QueryRow(query string, args ...any) *sql.Row {
	return db.client.QueryRow(query, args...)
}

func (db *MySqlDB) Exec(query string, args ...any) (sql.Result, error) {
	return db.client.Exec(query, args...)
}
