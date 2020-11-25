package db

import (
	"TableToStruct/config"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func MysqlConnect(dbConfig *config.Database) (*sql.DB, error) {
	db, err := sql.Open("mysql", dbConfig.ConnString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db. %s", err.Error())
	}
	return db, nil
}

func PostgresConnect(dbConfig *config.Database) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbConfig.ConnString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db. %s", err.Error())
	}
	return db, nil
}
