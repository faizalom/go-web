package routers

import (
	"net/http"

	"github.com/faizalom/go-web/config"
	"github.com/faizalom/go-web/controllers/frontend"
	"github.com/faizalom/go-web/middleware"
	"github.com/gorilla/csrf"
)

func defaultRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", fileNotFoundHandler)
	mux.Handle("GET /public/", http.StripPrefix("/public", http.FileServer(http.Dir(config.Path.Public))))
}

func WebRouters() http.Handler {
	web := http.NewServeMux()
	defaultRoutes(web)

	web.HandleFunc("GET /login", frontend.Login)
	web.HandleFunc("POST /login", frontend.LoginSubmit)
	// Google redirect URL
	// Add this url in you google console OAuth Authorised redirect URIs
	web.HandleFunc("GET /google-user/login", frontend.GoogleLogin)
	web.HandleFunc("GET /logout", frontend.Logout)

	web.HandleFunc("GET /register", frontend.Register)
	web.HandleFunc("POST /register", frontend.RegisterSubmit)
	web.HandleFunc("GET /register/{jwtToken}", frontend.CompleteRegister)

	// This route only accessible when user login or user with valid auth session
	web.HandleFunc("GET /{$}", middleware.AuthMiddleware(frontend.Dashboard))

	// This route protects all routes with CSRF
	CSRF := csrf.Protect([]byte(config.Cipher), csrf.ErrorHandler(http.HandlerFunc(pageExpiredHandler)))
	return CSRF(web)
}

func NoCSRFRouters() http.Handler {
	web := http.NewServeMux()
	defaultRoutes(web)

	web.HandleFunc("POST /loginsubmit", frontend.LoginSubmit)
	return web
}
