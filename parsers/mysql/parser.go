package mysql

import (
	"TableToStruct/config"
	"database/sql"
	"fmt"
	"regexp"
	"strings"
)

type MysqlDDL struct {
	table       string
	createTable string
}

//`id`
var columnNameReg *regexp.Regexp = regexp.MustCompile("(\\`\\w+\\`)\\s.*,$")

//int(11) or timestamp
var columnTypeWSizeReg *regexp.Regexp = regexp.MustCompile("\\s(\\w+\\(\\w+\\))\\s.*,$")
var columnTypeWoutSizeReg *regexp.Regexp = regexp.MustCompile("\\s(\\w+)\\s.*,$")

//NOT NULL or DEFAULT NULL
var columnNotNullReg *regexp.Regexp = regexp.MustCompile("NOT NULL")
var columnNullReg *regexp.Regexp = regexp.MustCompile("DEFAULT NULL")

func (m *MysqlDDL) Parse(tableMap map[string]string) string {
	var tableName string
	var ddl string
	for tableName, ddl = range tableMap {
	}
	// tableName := config.Get().TableName
	sqlStruct := &strings.Builder{}
	config.Logger().Info("Parsing the table DDL")
	sqlStruct.WriteString(fmt.Sprintf("package %s \n\ntype %s struct {\n", tableName, tableName))

	for i, line := range strings.Split(ddl, "\n") {
		//first line is the create table line from show create table
		if i == 0 {
			continue
		}
		var columnName string
		if columnNameReg.MatchString(line) {
			for _, value := range columnNameReg.FindStringSubmatch(line) {
				columnName = value
			}
			line = strings.Trim(line, columnName)
			columnName = strings.Trim(columnName, "\\`")

		} else {
			//if no column name is found then continue to the next line
			continue
		}
		var columnType string
		if columnTypeWSizeReg.MatchString(line) {
			for _, value := range columnTypeWSizeReg.FindStringSubmatch(line) {
				columnType = value
			}
		} else if columnTypeWoutSizeReg.MatchString(line) {
			for _, value := range columnTypeWoutSizeReg.FindStringSubmatch(line) {
				columnType = value
			}
		}

		line = strings.Trim(strings.TrimSpace(line), columnType)
		if columnNotNullReg.MatchString(line) {
			mysqlToGoNotNull(columnName, columnType, sqlStruct)
		} else if columnNullReg.MatchString(line) {
			mysqlToGoNull(columnName, columnType, sqlStruct)
		}
	}
	//this is the table name method
	sqlStruct.WriteString(fmt.Sprintf("}\n\nfunc (%s *%s) TableName() string {\nreturn \"%s\"\n}", tableName, tableName, tableName))
	config.Logger().Info("Finished parsing")

	return sqlStruct.String()
}

func (m *MysqlDDL) GetTable(rows *sql.Rows) map[string]string {
	defer rows.Close()

	tableMap := make(map[string]string)
	for rows.Next() {
		err := rows.Scan(&m.table, &m.createTable)
		if err != nil {
			config.Logger().Fatal(err.Error())
		}
		tableMap[m.table] = m.createTable
	}
	return tableMap
}

func splitter(r rune) bool {
	return r == '_' || r == '-'
}

//NOT NULL types
func mysqlToGoNotNull(columnName, columnType string, line *strings.Builder) {
	colName := ""
	columnNames := strings.FieldsFunc(columnName, splitter)
	for _, name := range columnNames {
		colName += strings.Title(name)
	}
	line.WriteString(fmt.Sprintf("%s ", colName))

	switch {
	case strings.Contains(columnType, "int"):
		line.WriteString("int64 ")
	case strings.Contains(columnType, "TINYINT"):
		line.WriteString("int32 ")
	case strings.Contains(columnType, "decimal"):
		line.WriteString("float64 ")
	case strings.Contains(columnType, "varchar"):
		line.WriteString("string ")
	case strings.Contains(columnType, "timestamp"):
		line.WriteString("time.Time ")
	default:
		line.WriteString("string ")
	}
	line.WriteString(fmt.Sprintf("`gorm:\"column:%s\"`\n", columnName))
}

//DEFAULT NULL types
func mysqlToGoNull(columnName, columnType string, line *strings.Builder) {
	colName := ""
	columnNames := strings.FieldsFunc(columnName, splitter)
	for _, name := range columnNames {
		colName += strings.Title(name)
	}
	line.WriteString(fmt.Sprintf("%s ", colName))

	switch {
	case strings.Contains(columnType, "int"):
		line.WriteString("sql.NullInt64 ")
	case strings.Contains(columnType, "TINYINT"):
		line.WriteString("sql.NullInt32 ")
	case strings.Contains(columnType, "decimal"):
		line.WriteString("sql.NullFloat64 ")
	case strings.Contains(columnType, "varchar"):
		line.WriteString("sql.NullString ")
	case strings.Contains(columnType, "timestamp"):
		line.WriteString("sql.NullTime ")
	default:
		line.WriteString("sql.NullString ")
	}
	line.WriteString(fmt.Sprintf("`gorm:\"column:%s\"`\n", columnName))

}
