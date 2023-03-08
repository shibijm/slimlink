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
	mySqlDB := &MySqlDB{db}
	return mySqlDB, nil
}

func (mySqlDB *MySqlDB) QueryRow(query string, args ...any) *sql.Row {
	return mySqlDB.db.QueryRow(query, args...)
}

func (mySqlDB *MySqlDB) Query(query string, args ...any) (*sql.Rows, error) {
	return mySqlDB.db.Query(query, args...)
}

func (mySqlDB *MySqlDB) Exec(query string, args ...any) (sql.Result, error) {
	return mySqlDB.db.Exec(query, args...)
}
