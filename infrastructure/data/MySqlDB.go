package data

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type MySqlDB struct {
	db *sql.DB
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

func (mySqlDB *MySqlDB) EnsureCreated() error {
	_, err := mySqlDB.Exec("CREATE TABLE IF NOT EXISTS `links` (`id` VARCHAR(64) NOT NULL, `url` VARCHAR(2048) NOT NULL, PRIMARY KEY (`id`))")
	return err
}

func (mySqlDB *MySqlDB) QueryRow(query string, args ...any) *sql.Row {
	return mySqlDB.db.QueryRow(query, args...)
}

func (mySqlDB *MySqlDB) Exec(query string, args ...any) (sql.Result, error) {
	return mySqlDB.db.Exec(query, args...)
}
