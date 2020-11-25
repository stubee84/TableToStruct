package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
)

type Config struct {
	// Dialect    string     `json:"dialect"`
	// ConnString string     `json:"connString"`
	// TableName  string     `json:"tableName"`
	LogFile   string     `json:"logFile"`
	Databases []Database `json:"databases"`
}

type Database struct {
	Dialect    string  `json:"dialect"`
	ConnString string  `json:"connString"`
	Tables     []Table `json:"tables"`
	Query      string
}

type Table struct {
	Name string `json:"name"`
}

func InitConfig() {
	if !FlagParser() {
		cfg := &Config{}
		cfgBody, err := ioutil.ReadFile(cfgFile)
		if err != nil {
			FlagParser()
			Logger().Fatal(err.Error())
		}
		err = json.Unmarshal(cfgBody, cfg)
		if err != nil {
			Logger().Fatal(err.Error())
		}

		logFile = cfg.LogFile
		databases = cfg.Databases

		// for _, db := range cfg.Databases {
		// 	for _, table := range db.Tables {
		// 		tables = append(tables, table)
		// 	}
		// }
	}
}

func FlagParser() bool {
	flag.StringVar(&cfgFile, "file", "config/config.json", "Config file for table access")
	flag.StringVar(&dialect, "d", "mysql", "Specify SQL dialect. Default is mysql")
	flag.StringVar(&connString, "c", "", "Specify the connection string. Example: user:password@tcp(127.0.0.1:3306)/database")
	flag.StringVar(&tableName, "t", "", "The database tablename to parse.")
	flag.StringVar(&logFile, "l", "config/TableToStruct", "The log file to which logs are written")

	flag.Parse()

	if dialect == "" || connString == "" || tableName == "" {
		log.Print("empty values for -d, -c or -t. using config.json")
		return false
	}
	return true
}

func GetDBs() []Database {
	return databases
}

var dialect string
var connString string
var tableName string
var logFile string
var cfgFile string

// var tables []Table = []Table{}
var databases []Database
