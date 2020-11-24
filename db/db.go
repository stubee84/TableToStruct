package db

import (
	"TableToStruct/config"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func MysqlConnect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.Get().ConnString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db. %s", err.Error())
	}
	return db, nil
}

func PostgresConnect() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.Get().ConnString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db. %s", err.Error())
	}
	return db, nil
}
