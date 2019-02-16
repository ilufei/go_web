package utils

import (
	"os"
	"encoding/json"
	"fmt"
	"log"
)

type Configuration struct {
    Address      string
    Static       string
}

var Config Configuration
var logger *log.Logger
var systemLogger *log.Logger

func init() {
	initAppConfig()
	initLogger()
	initSystemLogger()
}

func initAppConfig() {
	file, _ := os.Open("config/app.config.json")
	defer file.Close()

	jsonDecoder := json.NewDecoder(file)
	Config = Configuration{}
	err := jsonDecoder.Decode(&Config)
	if err != nil {
    	fmt.Println("config.json content decode failed")
    }
}

func initSystemLogger() {
	file, err := os.OpenFile("logs/system.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Failed to open log file", err)
	}
	systemLogger = log.New(file, "INFO ", log.Ldate|log.Ltime)
}

func initLogger() {
	file, err := os.OpenFile("logs/common.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Failed to open log file", err)
	}
	logger = log.New(file, "INFO ", log.Ldate|log.Ltime)
}

// for logging
func Slog(args ...interface{}) {
	systemLogger.SetPrefix("[SYSTEM] ")
	systemLogger.Println(args...)
}

// for logging
func Info(args ...interface{}) {
	logger.SetPrefix("[INFO] ")
	logger.Println(args...)
}

func Error(args ...interface{}) {
	logger.SetPrefix("[ERROR] ")
	logger.Println(args...)
}

func Warning(args ...interface{}) {
	logger.SetPrefix("[WARNING] ")
	logger.Println(args...)
}
