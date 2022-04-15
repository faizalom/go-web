package middleware

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/faizalom/go-web/lib"
	"github.com/julienschmidt/httprouter"
)

var l = log.New(log.Writer(), log.Prefix(), log.Flags())

func init() {
	logFile, err := os.OpenFile("access.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	//defer logFile.Close()
	l.SetOutput(logFile)
	//l.SetFlags(log.LstdFlags | log.Lshortfile)
}

func Logger(f func(http.ResponseWriter, *http.Request, httprouter.Params)) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		start := time.Now().UnixNano()
		defer l.Println(r.RemoteAddr, r.RequestURI, time.Now().UnixNano()-start)
		//fmt.Println(now() + "before")
		//defer fmt.Println(now() + "after")
		f(w, r, ps) // original function call
	}
}

func AuthMiddleware(f func(http.ResponseWriter, *http.Request, httprouter.Params)) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		start := time.Now().UnixNano()
		defer l.Println(r.RemoteAddr, r.RequestURI, time.Now().UnixNano()-start)
		auth, _ := lib.Auth(r)
		if auth, ok := auth.Values["authenticated"].(bool); !ok || !auth {
			redirect(w, r, ps)
		} else {
			f(w, r, ps) // original function call
		}
	}
}

func redirect(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
