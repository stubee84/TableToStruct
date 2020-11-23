package main

import (
	"TableToStruct/config"
	"TableToStruct/db"
	"TableToStruct/parsers/mysql"
	"database/sql"
	"fmt"
	"log"
	"os"
)

func main() {
	config.Get()
	config.Logger()
	defer config.FileLogger().Close()

	result := mysql.Mysql(selectQuery())

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

func selectQuery() map[string]string {
	var query string
	var tableMap map[string]string
	switch config.Get().Dialect {
	case "mysql":
		query = fmt.Sprintf("show create table %s", config.Get().TableName)
		tableMap = getMysqlTable(initQuery(query))
	case "postgresql":
		query = fmt.Sprintf(`select column_name,is_nullable,data_type 
			from information_schema.columns 
			where table_name = '%s';`, config.Get().TableName)
		tableMap = getPostgresTable(initQuery(query))
	}
	return tableMap
}

func initQuery(query string) *sql.Rows {
	db, err := db.Connect()
	if err != nil {
		log.Println(err)
	}

	config.Logger().Info(fmt.Sprintf("Querying the table %s", config.Get().TableName))
	rows, err := db.Query(query)
	if err != nil {
		config.Logger().Fatal(err.Error())
	}
	return rows
}

func getMysqlTable(rows *sql.Rows) map[string]string {
	defer rows.Close()

	type columns struct {
		table       string
		createTable string
	}
	cols := columns{}
	tableMap := make(map[string]string)
	for rows.Next() {
		err := rows.Scan(&cols.table, &cols.createTable)
		if err != nil {
			config.Logger().Fatal(err.Error())
		}
		tableMap[cols.table] = cols.createTable
	}
	return tableMap
}

func getPostgresTable(rows *sql.Rows) map[string]string {
	defer rows.Close()

	type columns struct {
		columnName string
		isNullable string
		dataType   string
	}
	cols := columns{}
	tableMap := make(map[string]string)
	for rows.Next() {
		err := rows.Scan(&cols.columnName, &cols.isNullable, &cols.dataType)
		if err != nil {
			config.Logger().Fatal(err.Error())
		}

		isNullable := true
		if cols.isNullable != "YES" {
			isNullable = false
		}
		tableMap[cols.columnName] = fmt.Sprintf("%t:%s", isNullable, cols.dataType)
	}

	return tableMap
}
