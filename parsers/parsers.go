package parsers

import (
	"TableToStruct/config"
	"TableToStruct/db"
	"database/sql"
	"fmt"
)

type SQLTableParser interface {
	Parse(map[string]string) string
	GetTable(*sql.Rows) map[string]string
}

func InitQuery(query string) *sql.Rows {
	var DB *sql.DB
	var err error
	switch config.Get().Dialect {
	case "mysql":
		DB, err = db.MysqlConnect()
	case "postgres":
		DB, err = db.PostgresConnect()
	}
	if err != nil {
		config.Logger().Error(err.Error())
	}

	config.Logger().Info(fmt.Sprintf("Querying the table %s", config.Get().TableName))
	rows, err := DB.Query(query)
	if err != nil {
		config.Logger().Fatal(err.Error())
	}
	return rows
}
