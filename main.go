package main

import (
	"TableToStruct/config"
	"TableToStruct/parsers"
	"TableToStruct/parsers/mysql"
	"TableToStruct/parsers/postgresql"
	"database/sql"
	"fmt"
	"os"
)

func main() {
	config.Get()
	config.Logger()
	defer config.FileLogger().Close()

	result := selectQuery()

	file, err := os.OpenFile(fmt.Sprintf("%s.go", config.Get().TableName), os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		config.Logger().Fatal(err.Error())
	}
	defer file.Close()

	_, err = file.Write([]byte(result))
	if err != nil {
		config.Logger().Fatal(err.Error())
	}
}

func selectQuery() string {
	var ddl parsers.SQLTableParser
	var rows *sql.Rows
	switch config.Get().Dialect {
	case "mysql":
		ddl = &mysql.MysqlDDL{}

		query := fmt.Sprintf("show create table %s", config.Get().TableName)
		rows = parsers.InitQuery(query)
	case "postgres":
		ddl = &postgresql.PostgresDDL{}
		query := fmt.Sprintf(`select column_name,is_nullable,data_type 
			from information_schema.columns 
			where table_name = '%s';`, config.Get().TableName)
		rows = parsers.InitQuery(query)
	}
	return ddl.Parse(ddl.GetTable(rows))
}
