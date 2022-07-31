package middleware

import (
	"log"
	"net/http"
	"os"

	"github.com/faizalom/go-web/config"
	"github.com/julienschmidt/httprouter"
)

// This function is run during every request to your application. And stored in a log file
func RequestLogger(targetMux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		targetMux.ServeHTTP(w, r)
		AccessLog(r)
	})
}

// Log file store function
func AccessLog(r *http.Request) {
	l := log.New(log.Writer(), log.Prefix(), log.Flags())

	logFile, err := os.OpenFile(config.AccessLogFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err)
	}
	defer logFile.Close()

	l.SetOutput(logFile)
	l.Printf(
		"%s\t%s\t%s",
		r.Method,
		r.RemoteAddr,
		r.RequestURI,
	)
}

// This function is called from AuthMiddleware
// This will run at user tries to access auth routes with invalid or Expired auth session or without login
// User can use this fucntion in controllers
func redirectToLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
