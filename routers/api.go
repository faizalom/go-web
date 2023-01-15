package routers

import (
	"github.com/faizalom/go-web/controllers/api/usercontroller"
	"github.com/faizalom/go-web/middleware"
)

func apiRoute() {
	router.POST("/api/login", usercontroller.Login)
	router.GET("/api/google-user/login", usercontroller.GoogleLogin)

	router.POST("/api/register", usercontroller.Register)
	router.GET("/api/register/:jwtToken", usercontroller.CompleteRegister)

	router.GET("/api/profile", middleware.ApiAuthMiddleware(usercontroller.Profile))
	router.PATCH("/api/profile", middleware.ApiAuthMiddleware(usercontroller.Update))
	router.GET("/api/logout", middleware.ApiAuthMiddleware(usercontroller.Profile))
}
