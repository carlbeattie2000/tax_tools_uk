package logger

import (
	"log"
	"os"
	"sync"
)

var (
	instance *log.Logger
	once     sync.Once
)

func GetLogger() *log.Logger {
	once.Do(func() {
		file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("Failed to open log file: %v", err)
		}
		instance = log.New(file, "", log.LstdFlags)
	})
	return instance
}
