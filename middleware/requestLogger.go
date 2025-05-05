package middleware

import (
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/faizalom/go-web/config"
)

var (
	requestLogger = log.New(log.Writer(), log.Prefix(), log.Flags())
	logLock       sync.Mutex // Mutex to prevent race conditions
)

// This function is run during every request to your application. And stored in a log file
func RequestLogger(targetMux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func(w http.ResponseWriter, r *http.Request) {
			if rec := recover(); rec != nil {
				log.Println("Panic: ", r.URL.Path, rec)
				http.Error(w, rec.(error).Error(), http.StatusInternalServerError)
			}
		}(w, r)

		// Pass the original request to the logging goroutine
		go AccessLog(r)
		targetMux.ServeHTTP(w, r)
	})
}

// Log file store function
func AccessLog(r *http.Request) {
	logLock.Lock() // Lock before accessing shared resources
	defer logLock.Unlock()

	logFile, err := os.OpenFile(config.LogFile.AccessLog, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err)
		return
	}
	defer logFile.Close()

	requestLogger.SetOutput(logFile)

	requestLogger.Printf(
		"%s\t%s\t%s",
		r.Method,
		r.RemoteAddr,
		r.RequestURI,
	)
}
