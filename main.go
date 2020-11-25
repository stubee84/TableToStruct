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
	config.Logger()
	defer config.FileLogger().Close()

	config.InitConfig()

	for _, database := range config.GetDBs() {
		selectQuery(&database)
	}
}

func selectQuery(database *config.Database) {
	var ddl parsers.SQLTableParser
	var rows *sql.Rows

	for _, table := range database.Tables {
		config.Logger().Info(fmt.Sprintf("Querying the table %s", table.Name))
		switch database.Dialect {
		case "mysql":
			ddl = &mysql.MysqlDDL{}

			database.Query = fmt.Sprintf("show create table %s", table.Name)
			rows = parsers.InitQuery(database)
		case "postgres":
			ddl = &postgresql.PostgresDDL{
				Table: table.Name,
			}
			database.Query = fmt.Sprintf(`select column_name,is_nullable,data_type 
				from information_schema.columns 
				where table_name = '%s';`, table.Name)

			rows = parsers.InitQuery(database)
		}

		file, err := os.OpenFile(fmt.Sprintf("%s.go", table.Name), os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			config.Logger().Fatal(err.Error())
		}
		defer file.Close()

		_, err = file.Write([]byte(ddl.Parse(ddl.GetTable(rows))))
		if err != nil {
			config.Logger().Fatal(err.Error())
		}
	}
}
