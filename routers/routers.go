// Package routers create your routes
package routers

import (
	"net/http"
)

/*
Web Routes

Here is where you can register web routes for your application.
Add New Routes inside this function.
*/
func SetRoutes() http.Handler {
	// Call func to define your routes
	mux := http.NewServeMux()
	mux.Handle("/", WebRouters())
	mux.Handle("/api/", http.StripPrefix("/api", APIRouters()))

	return mux
}
