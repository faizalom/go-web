package routers

import (
	"net/http"

	"github.com/faizalom/go-web/controllers"
	"github.com/faizalom/go-web/middleware"
	"github.com/julienschmidt/httprouter"
)

func SetRoutes() *httprouter.Router {
	router := httprouter.New()
	router.ServeFiles("/public/*filepath", http.Dir("public"))

	router.GET("/", middleware.AuthMiddleware(controllers.IndexContoller))
	router.GET("/login", controllers.LoginIndexContoller)
	router.POST("/login", controllers.LoginSubmitContoller)
	return router
}
