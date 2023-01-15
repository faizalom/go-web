package routers

import (
	"github.com/faizalom/go-web/controllers/frontend"
)

func web() {

	router.GET("/login", frontend.Login)
	router.GET("/google-user/login", frontend.GoogleLogin)

	router.GET("/register", frontend.Register)
	router.GET("/register/:jwtToken", frontend.CompleteRegister)

	router.GET("/", frontend.Dashboard)
	router.GET("/profile", frontend.Profile)
}
