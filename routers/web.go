package routers

import (
	"net/http"

	"github.com/faizalom/go-web/config"
	"github.com/faizalom/go-web/controllers/frontend"
	"github.com/gorilla/csrf"
)

// func noDirListing(h http.Handler) http.HandlerFunc {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// if strings.HasSuffix(r.URL.Path, "/") {
// 		// 	fileNotFoundHandler(w, r)
// 		// 	return
// 		// }
// 		h.ServeHTTP(w, r)
// 	})
// }

func defaultRoutes(mux *http.ServeMux) {
	mux.Handle("GET /", http.FileServer(http.Dir(config.Path.Public)))

	mux.HandleFunc("/", fileNotFoundHandler)
	mux.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("----- METHOD NOT ALLOWED -----"))
	})
}

func WebRouters() http.Handler {
	web := http.NewServeMux()
	defaultRoutes(web)

	web.HandleFunc("GET /login", frontend.Login)
	// router.GET("/google-user/login", frontend.GoogleLogin)

	web.HandleFunc("GET /register", frontend.Register)
	web.HandleFunc("GET /register/{jwtToken}", frontend.CompleteRegister)

	web.HandleFunc("GET /{$}", frontend.Dashboard)
	web.HandleFunc("GET /profile", frontend.Profile)

	CSRF := csrf.Protect([]byte(config.Cipher), csrf.ErrorHandler(http.HandlerFunc(pageExpiredHandler)))
	return CSRF(web)
}
