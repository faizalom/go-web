package routers

import (
	"net/http"

	"github.com/faizalom/go-web/config"
	"github.com/faizalom/go-web/controllers"
	"github.com/faizalom/go-web/middleware"
	"github.com/julienschmidt/httprouter"
)

func SetRoutes() *httprouter.Router {
	router := httprouter.New()
	router.ServeFiles("/public/*filepath", http.Dir(config.PublicPath))

	router.GET("/", middleware.AuthMiddleware(controllers.IndexController))
	router.GET("/login", controllers.LoginIndexController)
	router.POST("/login", controllers.LoginSubmitController)
	return router
}
