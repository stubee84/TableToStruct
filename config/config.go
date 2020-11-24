package config

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Dialect    string `json:"dialect"`
	ConnString string `json:"connString"`
	TableName  string `json:"tableName"`
	LogFile    string `json:"logFile"`
}

func Get() *Config {
	if cfgFile == "" || (dialect == "" && connString == "" && tableName == "") {
		FlagParser()
	}

	return &Config{
		Dialect:    dialect,
		ConnString: connString,
		TableName:  tableName,
	}
}

func GetConfig() {
	cfg := &Config{}
	cfgBody, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		Logger().Fatal(err.Error())
	}
	err = json.Unmarshal(cfgBody, cfg)
	if err != nil {
		Logger().Fatal(err.Error())
	}

	dialect = cfg.Dialect
	connString = cfg.ConnString
	tableName = cfg.TableName
	logFile = cfg.LogFile
}

func FlagParser() {
	flag.StringVar(&cfgFile, "file", "config/config.json", "Config file for table access")
	flag.StringVar(&dialect, "d", "mysql", "Specify SQL dialect. Default is mysql")
	flag.StringVar(&connString, "c", "", "Specify the connection string. Example: user:password@tcp(127.0.0.1:3306)/database")
	flag.StringVar(&tableName, "t", "", "The database tablename to parse.")
	flag.StringVar(&logFile, "l", "config/TableToStruct", "The log file to which logs are written")

	flag.Parse()

	if dialect == "" || connString == "" || tableName == "" {
		log.Print("empty values for -d, -c or -t. using config.json")
		GetConfig()
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
var cfgFile string
