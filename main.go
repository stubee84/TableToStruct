package main

import (
	"TableToStruct/config"
	"TableToStruct/db"
	"TableToStruct/parsers/mysql"
	"fmt"
	"log"
	"os"
)

func main() {
	config.Get()
	config.Logger()
	defer config.FileLogger().Close()

	result := mysql.Mysql(selectQuery(), config.Get().TableName)

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
	var query string
	switch {
	case config.Get().Dialect == "mysql":
		query = fmt.Sprintf("show create table %s", config.Get().TableName)
		return getMysqlTable(query)
	case config.Get().Dialect == "postgresql":
		query = fmt.Sprintf(`select column_name,data_type 
			from information_schema.columns 
			where table_name = '%s';`, config.Get().TableName)
	}
	return ""
}

func getMysqlTable(query string) string {
	db, err := db.Connect()
	if err != nil {
		log.Println(err)
	}

	type columns struct {
		table       string
		createTable string
	}
	cols := columns{}

	config.Logger().Info(fmt.Sprintf("Querying the table %s", config.Get().TableName))
	rows, err := db.Query(query)
	if err != nil {
		config.Logger().Fatal(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&cols.table, &cols.createTable)
		if err != nil {
			config.Logger().Fatal(err.Error())
		}
	}
	return cols.createTable
}
