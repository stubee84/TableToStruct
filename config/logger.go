package config

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Logs struct {
	FileLogger *log.Logger
}

func FileLogger() *os.File {
	year, month, day := time.Now().Local().Date()
	file, err := os.OpenFile(fmt.Sprintf("%s_%d-%s-%d.log", logFile, year, month.String(), day), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err)
	}
	return file
}

func Logger() *Logs {
	log.SetOutput(os.Stdout)
	log.SetPrefix(logFile)
	log.SetFlags(log.Ldate | log.Ltime)
	fileLogger := log.New(FileLogger(), logFile, log.Ldate|log.Ltime)

	return &Logs{
		FileLogger: fileLogger,
	}
}

func (l *Logs) Fatal(fatal string) {
	l.FileLogger.Printf("Fatal: %s\n", fatal)
	log.Fatal(fatal)
}

func (l *Logs) Info(info string) {
	l.FileLogger.Printf("Info: %s\n", info)
	log.Printf("Info: %s\n", info)
}

func (l *Logs) Error(err string) {
	l.FileLogger.Printf("Error %s\n", err)
	log.Printf("Error %s\n", err)
}
