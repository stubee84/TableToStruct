package postgresql

import (
	"TableToStruct/config"
	"fmt"
	"strconv"
	"strings"
)

func Postgres(tableMap map[string]string) string {
	tableName := config.Get().TableName
	sqlStruct := strings.Builder{}

	config.Logger().Info("Parsing the table DDL")
	sqlStruct.WriteString(fmt.Sprintf("package %s \n\ntype %s struct {\n", tableName, tableName))

	for column, value := range tableMap {
		splitValue := strings.Split(value, ":")
		isNull, err := strconv.ParseBool(splitValue[0])
		if err != nil {
			config.Logger().Error(err.Error())
		}
		dataType := splitValue[1]

		if isNull {
			postgresToGoNotNull(column, dataType, &sqlStruct)
		} else {
			postgresToGoNull(column, dataType, &sqlStruct)
		}

	}

	sqlStruct.WriteString(fmt.Sprintf("}\n\nfunc (%s *%s) TableName() string {\nreturn \"%s\"\n}", tableName, tableName, tableName))
	return sqlStruct.String()
}

func postgresToGoNotNull(columnName, columnType string, line *strings.Builder) {
	line.WriteString(fmt.Sprintf("%s ", columnName))

	switch {
	case strings.Contains(columnType, "integer"):
		line.WriteString("int64 ")
	case strings.Contains(columnType, "bigint"):
		line.WriteString("int64 ")
	case strings.Contains(columnType, "smallint"):
		line.WriteString("int32 ")
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
	line.WriteString(fmt.Sprintf("%s ", columnName))

	switch {
	case strings.Contains(columnType, "integer"):
		line.WriteString("sql.NullInt64 ")
	case strings.Contains(columnType, "bigint"):
		line.WriteString("sql.NullInt64 ")
	case strings.Contains(columnType, "smallint"):
		line.WriteString("sql.NullInt32 ")
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
