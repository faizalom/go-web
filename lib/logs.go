package lib

import (
	"log"
	"os"
)

// This function is set a log file
// All errors will be store in errorLogFile
func LogErrors(errorLogFile string) {
	logFile, err := os.OpenFile(errorLogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
