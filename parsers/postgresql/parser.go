package postgresql

import (
	"TableToStruct/config"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
)

type PostgresDDL struct {
	columnName string
	isNullable string
	dataType   string
}

func (p *PostgresDDL) Parse(tableMap map[string]string) string {
	tableName := tableMap["tableName"]
	delete(tableMap, "tableName")
	sqlStruct := strings.Builder{}

	config.Logger().Info("Parsing the table DDL")
	sqlStruct.WriteString(fmt.Sprintf("package %s \n\ntype %s struct {\n", tableName, tableName))

	for i := 1; i <= len(tableMap); i++ {
		splitValue := strings.Split(tableMap[fmt.Sprint(i)], ":")
		column := splitValue[0]
		isNull, err := strconv.ParseBool(splitValue[1])
		if err != nil {
			config.Logger().Error(err.Error())
		}
		dataType := splitValue[2]

		if isNull {
			postgresToGoNull(column, dataType, &sqlStruct)
		} else {
			postgresToGoNotNull(column, dataType, &sqlStruct)
		}
	}

	sqlStruct.WriteString(fmt.Sprintf("}\n\nfunc (%s *%s) TableName() string {\nreturn \"%s\"\n}", tableName, tableName, tableName))
	config.Logger().Info("Finished parsing")

	return sqlStruct.String()
}

func (p *PostgresDDL) GetTable(rows *sql.Rows) map[string]string {
	defer rows.Close()

	tableMap := make(map[string]string)
	index := 0
	for rows.Next() {
		index++
		err := rows.Scan(&p.columnName, &p.isNullable, &p.dataType)
		if err != nil {
			config.Logger().Fatal(err.Error())
		}

		isNullable := false
		if p.isNullable == "YES" {
			isNullable = true
		}
		//provide the index as the key because maps do not retain order
		tableMap[strconv.Itoa(index)] = fmt.Sprintf("%s:%t:%s", p.columnName, isNullable, p.dataType)
	}
	tableMap["tableName"] = config.Get().TableName

	return tableMap
}

func splitter(r rune) bool {
	return r == '_' || r == '-'
}

func postgresToGoNotNull(columnName, columnType string, line *strings.Builder) {
	colName := ""
	columnNames := strings.FieldsFunc(columnName, splitter)
	for _, name := range columnNames {
		colName += strings.Title(name)
	}
	line.WriteString(fmt.Sprintf("%s ", colName))

	switch {
	case strings.Contains(columnType, "integer"):
		line.WriteString("int64 ")
	case strings.Contains(columnType, "bigint"):
		line.WriteString("int64 ")
	case strings.Contains(columnType, "smallint"):
		line.WriteString("int32 ")
	case strings.Contains(columnType, "boolean"):
		line.WriteString("bool ")
	case strings.Contains(columnType, "double precision"):
		line.WriteString("float64 ")
	case strings.Contains(columnType, "real"):
		line.WriteString("float32 ")
	case strings.Contains(columnType, "text"):
		line.WriteString("string ")
	case strings.Contains(columnType, "character"):
		line.WriteString("string ")
	case strings.Contains(columnType, "timestamp"):
		line.WriteString("time.Time ")
	case strings.Contains(columnType, "time"):
		line.WriteString("time.Time ")
	default:
		line.WriteString("string ")
	}
	line.WriteString(fmt.Sprintf("`gorm:\"column:%s\"`\n", columnName))
}

func postgresToGoNull(columnName, columnType string, line *strings.Builder) {
	colName := ""
	columnNames := strings.FieldsFunc(columnName, splitter)
	for _, name := range columnNames {
		colName += strings.Title(name)
	}
	line.WriteString(fmt.Sprintf("%s ", colName))

	switch {
	case strings.Contains(columnType, "integer"):
		line.WriteString("sql.NullInt64 ")
	case strings.Contains(columnType, "bigint"):
		line.WriteString("sql.NullInt64 ")
	case strings.Contains(columnType, "smallint"):
		line.WriteString("sql.NullInt32 ")
	case strings.Contains(columnType, "boolean"):
		line.WriteString("sql.NullBool ")
	case strings.Contains(columnType, "double precision"):
		line.WriteString("sql.NullFloat64 ")
	case strings.Contains(columnType, "real"):
		line.WriteString("sql.NullFloat32 ")
	case strings.Contains(columnType, "text"):
		line.WriteString("sql.NullString ")
	case strings.Contains(columnType, "character"):
		line.WriteString("sql.NullString ")
	case strings.Contains(columnType, "timestamp"):
		line.WriteString("sql.NullTime ")
	case strings.Contains(columnType, "time"):
		line.WriteString("sql.NullTime ")
	default:
		line.WriteString("sql.NullString ")
	}
	line.WriteString(fmt.Sprintf("`gorm:\"column:%s\"`\n", columnName))

}
