// Package routers create your routes
package routers

import (
	"net/http"

	"github.com/faizalom/go-web/config"
	"github.com/faizalom/go-web/controllers"
	"github.com/faizalom/go-web/middleware"
	"github.com/gorilla/csrf"
	"github.com/julienschmidt/httprouter"
)

func SetRoutes() http.Handler {
	router := httprouter.New()
	router.ServeFiles("/public/*filepath", http.Dir(config.PublicPath))

	router.GET("/", middleware.AuthMiddleware(controllers.IndexController))
	router.GET("/login", controllers.LoginIndexController)
	router.POST("/login", controllers.LoginSubmitController)
	router.GET("/logout", controllers.LogoutController)
	router.GET("/register", controllers.RegisterController)
	router.POST("/register", controllers.RegisterSubmitController)

	// router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "resources/views/404.html")
	// })

	router.NotFound = http.HandlerFunc(fileNotFoundHandler)
	router.MethodNotAllowed = http.HandlerFunc(methodNotAllowedHandler)

	CSRF := csrf.Protect([]byte("32-byte-long-auth-key"), csrf.ErrorHandler(http.HandlerFunc(pageExpiredHandler)))
	return CSRF(router)
}
