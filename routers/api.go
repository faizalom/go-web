package routers

import (
	"net/http"

	"github.com/faizalom/go-web/controllers/api/usercontroller"
	"github.com/faizalom/go-web/middleware"
)

func APIRouters() http.Handler {
	api := http.NewServeMux()
	api.HandleFunc("POST /login", usercontroller.Login)
	// router.GET("/api/google-user/login", usercontroller.GoogleLogin)

	api.HandleFunc("POST /register", usercontroller.Register)
	api.HandleFunc("GET /register/{jwtToken}", usercontroller.CompleteRegister)

	api.HandleFunc("GET /profile", middleware.ApiAuthMiddleware(usercontroller.Profile))
	api.HandleFunc("PATCH /profile", middleware.ApiAuthMiddleware(usercontroller.Update))
	// router.GET("/api/logout", middleware.ApiAuthMiddleware(usercontroller.Profile))

	return api
}
