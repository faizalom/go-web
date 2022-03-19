package middleware

import (
	"log"
	"net/http"

	"github.com/faizalom/go-web/lib"
	"github.com/julienschmidt/httprouter"
)

// func NoAuth() {
// 	http.HandleFunc("/coreui", Logger(controllers.CoreUI))
// }

func Logger(f func(http.ResponseWriter, *http.Request, httprouter.Params)) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		//fmt.Printf("%#v", r)
		defer log.Println(r.RemoteAddr, r.RequestURI)
		//fmt.Println(now() + "before")
		//defer fmt.Println(now() + "after")
		f(w, r, ps) // original function call
	}
}

func AuthMiddleware(f func(http.ResponseWriter, *http.Request, httprouter.Params)) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		defer log.Println(r.RemoteAddr, r.RequestURI)
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
