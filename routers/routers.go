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

/*
Web Routes

Here is where you can register web routes for your application.
Add New Routes inside this function.
*/
func SetRoutes() http.Handler {
	router := httprouter.New()
	router.ServeFiles("/public/*filepath", http.Dir(config.PublicPath))

	router.GET("/login", controllers.LoginIndexController)
	router.POST("/login", controllers.LoginSubmitController)
	router.GET("/logout", controllers.LogoutController)
	router.GET("/register", controllers.RegisterController)
	router.POST("/register", controllers.RegisterSubmitController)

	// This route only accessible when user login or user with valid auth session
	router.GET("/", middleware.AuthMiddleware(controllers.IndexController))

	// router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "resources/views/404.html")
	// })

	router.NotFound = http.HandlerFunc(fileNotFoundHandler)
	router.MethodNotAllowed = http.HandlerFunc(methodNotAllowedHandler)

	CSRF := csrf.Protect([]byte(config.Cipher), csrf.ErrorHandler(http.HandlerFunc(pageExpiredHandler)))
	return CSRF(router)
}
