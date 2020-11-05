package config

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type Config struct {
	Dialect    string
	ConnString string
	TableName  string
}

func Get() *Config {
	if dialect == "" && connString == "" && tableName == "" {
		FlagParser()
	}

	return &Config{
		Dialect:    dialect,
		ConnString: connString,
		TableName:  tableName,
	}
}

func FlagParser() {
	flag.StringVar(&dialect, "d", "mysql", "Specify SQL dialect. Default is mysql")
	flag.StringVar(&connString, "c", "", "Specify the connection string. Example: user:password@tcp(127.0.0.1:3306)/database")
	flag.StringVar(&tableName, "t", "", "The database tablename to parse.")
	flag.StringVar(&logFile, "l", fmt.Sprintf("TableToStruct_%d ", time.Now().Unix()), "The log file to which logs are written")

	flag.Parse()

	if dialect == "" || connString == "" || tableName == "" {
		log.Fatalf("Cannot have empty values for -d, -c or -t")
	}
}

func ReadInput(stdOut string) string {
	fmt.Printf(stdOut)
	scanner := bufio.NewScanner(os.Stdin)

	for !scanner.Scan() {
	}

	return scanner.Text()
}

var dialect string
var connString string
var tableName string
var logFile string

// var dialect = ReadInput("Enter the sql dialect (mysql, postgresql): ")
// var connString = ReadInput("\nEnter the connection string ('user:password@tcp(127.0.0.1:3306)/database'): ")
// var tableName = ReadInput("\nEnter the table name: ")
