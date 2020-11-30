package parsers

import (
	"TableToStruct/config"
	"TableToStruct/db"
	"database/sql"
)

type SQLTableParser interface {
	Parse(map[string]string) string
	GetTable(*sql.Rows) map[string]string
}

func InitQuery(database *config.Database) *sql.Rows {
	var DB *sql.DB
	var err error
	switch database.Dialect {
	case "mysql":
		DB, err = db.MysqlConnect(database)
	case "postgres":
		DB, err = db.PostgresConnect(database)
	}
	if err != nil {
		config.Logger().Error(err.Error())
	}

	rows, err := DB.Query(database.Query)
	if err != nil {
		config.Logger().Fatal(err.Error())
	}
	return rows
}
