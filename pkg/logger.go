package pkg

import (
	"log"
	"os"
)

var FileLogger *log.Logger

func InitLogger() {
	logFile, err := os.OpenFile("fraud.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Gagal membuka file log: %v", err)
	}
	FileLogger = log.New(logFile, "", log.LstdFlags)
}
