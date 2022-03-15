//if you don't accept a definite you never a failure

package main

import (
	"fmt"
	"log"
	"net/http"
	"theme"

	"github.com/faizalom/go-web/controllers"
	"github.com/faizalom/go-web/middleware"

	"github.com/julienschmidt/httprouter"
)

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	router := httprouter.New()
	router.ServeFiles("/public/*filepath", http.Dir("public"))
	router.ServeFiles("/images/*filepath", http.Dir("images"))
	router.ServeFiles("/static/*filepath", http.Dir(theme.CoreUIPath))

	router.GET("/login", middleware.Logger(controllers.Login))
	router.POST("/login", middleware.Logger(controllers.LoginSubmit))

	router.GET("/", middleware.AuthMiddleware(controllers.CoreUI))
	router.GET("/profile", middleware.AuthMiddleware(controllers.Profile))

	router.GET("/hello/:name", Hello)

	log.Fatal(http.ListenAndServe(":8282", router))
}
