package routers

import (
	"fmt"
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

	router.NotFound = http.HandlerFunc(unauthorizedHandler)
	router.MethodNotAllowed = http.HandlerFunc(unauthorizedHandler)

	CSRF := csrf.Protect([]byte("32-byte-long-auth-key"), csrf.ErrorHandler(router.MethodNotAllowed))
	return CSRF(router)
}

func unauthorizedHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, fmt.Sprintf("%s - %s", http.StatusText(http.StatusForbidden), "ewrew"), http.StatusForbidden)
}
