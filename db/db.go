package db

import (
	"TableToStruct/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open(config.Get().Dialect, config.Get().ConnString)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("failed to connect to db. %s", err.Error())
	}
	return db, nil
}
