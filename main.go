//if you don't accept a definite you never a failure

package main

import (
	"log"
	"net/http"
	"os"
	"theme"

	"github.com/faizalom/go-web/controllers"
	"github.com/faizalom/go-web/controllers/apis"
	"github.com/faizalom/go-web/middleware"

	"github.com/julienschmidt/httprouter"
)

func main() {
	logFile, err := os.OpenFile("error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	//defer logFile.Close()

	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	router := httprouter.New()
	router.ServeFiles("/public/*filepath", http.Dir("public"))
	router.ServeFiles("/images/*filepath", http.Dir("images"))
	router.ServeFiles("/static/*filepath", http.Dir(theme.CoreUIPath))

	router.GET("/login", middleware.Logger(controllers.Login))
	router.POST("/login", middleware.Logger(controllers.LoginSubmit))
	router.GET("/logout", middleware.Logger(controllers.Logout))
	//router.GET("/u/*filepath", middleware.AuthMiddleware(controllers.CoreUI))
	router.GET("/u/*filepath", middleware.AuthMiddleware(controllers.CoreUI))

	router.GET("/api/staff", middleware.AuthMiddleware(apis.GetStaff))
	router.GET("/api/staff/:id", middleware.AuthMiddleware(apis.GetStaffByID))

	log.Fatal(http.ListenAndServe(":8181", router))
}
